# frozen_string_literal: true

require "fileutils"
require "digest"
require "json"
require "open3"
require "optparse"
require "pathname"
require_relative "desktop_entry"
require_relative "desktop_integration_manifest"
require_relative "dolphin_service_menu"
require_relative "file_association_model"
require_relative "recipe_install_gate"
require_relative "recipe_store"
require_relative "registry_backed_recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DesktopActivationInstaller
      APPLICATIONS_DIR = "usr/share/applications"
      SERVICE_MENUS_DIR = "usr/share/kio/servicemenus"
      MANIFESTS_DIR = "usr/share/xnix/compatibility/manifests"
      RECEIPTS_DIR = "usr/share/xnix/compatibility/activation-receipts"

      def initialize(root:, recipe:, install_gate: nil, desktop_entry_renderer: nil)
        @root = Pathname.new(root)
        @recipe = recipe
        @install_gate = install_gate
        @desktop_entry_renderer = desktop_entry_renderer || RubyDesktopEntryRenderer.new
      end

      def install
        validate_root!
        validate_preflight!
        validate_file_association_target!

        installed = [
          install_desktop_entry,
          install_dolphin_service_menu,
          install_file_associations,
          install_manifest
        ]

        receipt = install_receipt(installed)

        {
          "version" => RuntimeDaemon::VERSION,
          "application_id" => recipe.id,
          "root" => root.to_s,
          "preflight" => preflight_report,
          "installed" => installed,
          "receipt" => receipt,
          "activated_entry_points" => manifest.to_h.fetch("entry_points"),
          "safety" => {
            "staging_root_required" => true,
            "host_root_modified" => false,
            "backend_commands_exposed" => false,
            "desktop_entry_source" => desktop_entry_renderer.source,
            "recipe_install_gate_enforced" => !install_gate.nil?,
            "rollback_receipt_written" => true
          }
        }
      end

      private

      attr_reader :root, :recipe, :install_gate, :desktop_entry_renderer

      def validate_root!
        raise ArgumentError, "root must not be /" if root.cleanpath.to_s == "/"
      end

      def validate_preflight!
        return unless install_gate

        return if preflight_report.fetch("decision") == "allow"

        reason = preflight_report.fetch("blocking_reasons").join("; ")
        raise ArgumentError, "recipe install gate blocked activation: #{reason}"
      end

      def validate_file_association_target!
        destination = safe_destination(FileAssociationModel::MIMEAPPS_RELATIVE_PATH)
        raise ArgumentError, "refusing to overwrite existing mimeapps list" if destination.exist?
      end

      def install_desktop_entry
        desktop_entry = DesktopEntry.new(recipe)
        install_file(
          relative_path: File.join(APPLICATIONS_DIR, desktop_entry.file_name),
          contents: desktop_entry_renderer.render(recipe),
          mode: 0o644,
          kind: "desktop-entry",
          entry_point: "launcher"
        )
      end

      def install_dolphin_service_menu
        service_menu = DolphinServiceMenu.new
        install_file(
          relative_path: File.join(SERVICE_MENUS_DIR, service_menu.file_name),
          contents: service_menu.render,
          mode: 0o644,
          kind: "dolphin-service-menu",
          entry_point: "file-manager"
        )
      end

      def install_file_associations
        model = FileAssociationModel.new(recipe: recipe)
        install_file(
          relative_path: FileAssociationModel::MIMEAPPS_RELATIVE_PATH,
          contents: model.render_mimeapps,
          mode: 0o644,
          kind: "mimeapps-list",
          entry_point: "file-manager"
        )
      end

      def install_manifest
        install_file(
          relative_path: File.join(MANIFESTS_DIR, "#{recipe.id}.json"),
          contents: "#{JSON.pretty_generate(manifest.to_h)}\n",
          mode: 0o644,
          kind: "desktop-integration-manifest",
          entry_point: "all"
        )
      end

      def install_receipt(installed)
        receipt = {
          "version" => RuntimeDaemon::VERSION,
          "application_id" => recipe.id,
          "installed" => installed,
          "rollback" => {
            "command" => "xnix-rollback-desktop-integration",
            "requires_matching_sha256" => true
          }
        }

        install_file(
          relative_path: File.join(RECEIPTS_DIR, "#{recipe.id}.json"),
          contents: "#{JSON.pretty_generate(receipt)}\n",
          mode: 0o644,
          kind: "desktop-activation-receipt",
          entry_point: "rollback"
        )
      end

      def install_file(relative_path:, contents:, mode:, kind:, entry_point:)
        destination = safe_destination(relative_path)
        FileUtils.mkdir_p(destination.dirname)
        File.write(destination, contents)
        File.chmod(mode, destination)

        {
          "entry_point" => entry_point,
          "kind" => kind,
          "path" => relative_path,
          "mode" => format("%04o", mode),
          "sha256" => Digest::SHA256.file(destination).hexdigest
        }
      end

      def safe_destination(relative_path)
        raise ArgumentError, "install path must be relative" if Pathname.new(relative_path).absolute?
        raise ArgumentError, "install path must not escape root" if relative_path.split(File::SEPARATOR).include?("..")

        root.join(relative_path)
      end

      def manifest
        @manifest ||= DesktopIntegrationManifest.new(recipe: recipe)
      end

      def preflight_report
        @preflight_report ||= if install_gate
                                install_gate.to_h
                              else
                                {
                                  "gate_type" => "recipe-install",
                                  "decision" => "not-enforced",
                                  "blocking_reasons" => []
                                }
                              end
      end

      class RubyDesktopEntryRenderer
        def source
          "ruby"
        end

        def render(recipe)
          DesktopEntry.new(recipe).render
        end
      end

      class RuntimeGoDesktopEntryRenderer
        def initialize(command:, registry_path:, application_id:)
          @command = command
          @registry_path = registry_path
          @application_id = application_id
        end

        def source
          "runtime-go"
        end

        def render(recipe)
          raise ArgumentError, "runtime-go renderer application mismatch" unless recipe.id == @application_id

          output, status = Open3.capture2(
            @command,
            "desktop-entry-preview",
            "--registry",
            @registry_path,
            "--app",
            @application_id
          )
          raise ArgumentError, "runtime-go desktop entry renderer failed" unless status.success?

          validate_output!(output, recipe)
          output
        rescue SystemCallError => e
          raise ArgumentError, "runtime-go desktop entry renderer unavailable: #{e.message}"
        end

        private

        def validate_output!(output, recipe)
          required = [
            "[Desktop Entry]\n",
            "Exec=xnix-compat-launch --app #{recipe.id} %U\n",
            "X-Xnix-ApplicationId=#{recipe.id}\n"
          ]
          missing = required.reject { |fragment| output.include?(fragment) }
          raise ArgumentError, "runtime-go desktop entry renderer returned an incomplete desktop entry" unless missing.empty?
          raise ArgumentError, "runtime-go desktop entry renderer exposed backend details" if output.match?(/wine|prefix|\.exe|proton|qemu-system|program files/i)
        end
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @root = nil
          @mode = "production"
          @desktop_entry_source = "ruby"
          @runtime_go_bin = "xnix-runtime-go"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--root is required" unless @root
          raise ArgumentError, "--app is required" unless @application_id

          recipe_store = RegistryBackedRecipeStore.for_path(@recipe_dir)
          recipe = recipe_store.find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          install_gate = build_install_gate(recipe_store)
          result = DesktopActivationInstaller.new(
            root: @root,
            recipe: recipe,
            install_gate: install_gate,
            desktop_entry_renderer: build_desktop_entry_renderer
          ).install
          puts JSON.pretty_generate(result)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-install-desktop-integration: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-install-desktop-integration --root PATH --app APP_ID [--recipe-dir PATH] [--mode MODE]"
            options.on("--root PATH", "Install desktop activation files under PATH") do |value|
              @root = value
            end
            options.on("--app APP_ID", "Install desktop activation files for APP_ID") do |value|
              @application_id = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--mode MODE", RecipeInstallGate::MODES, "Evaluate production or development install policy") do |value|
              @mode = value
            end
            options.on("--desktop-entry-source SOURCE", %w[ruby runtime-go], "Render desktop entries with ruby or runtime-go") do |value|
              @desktop_entry_source = value
            end
            options.on("--runtime-go-bin PATH", "Path to xnix-runtime-go when --desktop-entry-source runtime-go is used") do |value|
              @runtime_go_bin = value
            end
          end
        end

        def build_desktop_entry_renderer
          return RubyDesktopEntryRenderer.new if @desktop_entry_source == "ruby"

          registry_path = File.join(@recipe_dir, "registry.json")
          RuntimeGoDesktopEntryRenderer.new(
            command: @runtime_go_bin,
            registry_path: registry_path,
            application_id: @application_id
          )
        end

        def build_install_gate(recipe_store)
          unless recipe_store.respond_to?(:registry_report)
            raise ArgumentError, "recipe install gate requires a verified recipe registry"
          end

          RecipeInstallGate.new(
            registry_report: recipe_store.registry_report,
            application_id: @application_id,
            mode: @mode
          )
        end
      end
    end
  end
end
