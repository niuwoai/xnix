# frozen_string_literal: true

require "fileutils"
require "digest"
require "json"
require "optparse"
require "pathname"
require_relative "desktop_entry"
require_relative "desktop_integration_manifest"
require_relative "dolphin_service_menu"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DesktopActivationInstaller
      APPLICATIONS_DIR = "usr/share/applications"
      SERVICE_MENUS_DIR = "usr/share/kio/servicemenus"
      MANIFESTS_DIR = "usr/share/xnix/compatibility/manifests"
      RECEIPTS_DIR = "usr/share/xnix/compatibility/activation-receipts"

      def initialize(root:, recipe:)
        @root = Pathname.new(root)
        @recipe = recipe
      end

      def install
        validate_root!

        installed = [
          install_desktop_entry,
          install_dolphin_service_menu,
          install_manifest
        ]

        receipt = install_receipt(installed)

        {
          "version" => RuntimeDaemon::VERSION,
          "application_id" => recipe.id,
          "root" => root.to_s,
          "installed" => installed,
          "receipt" => receipt,
          "activated_entry_points" => manifest.to_h.fetch("entry_points"),
          "safety" => {
            "staging_root_required" => true,
            "host_root_modified" => false,
            "backend_commands_exposed" => false,
            "rollback_receipt_written" => true
          }
        }
      end

      private

      attr_reader :root, :recipe

      def validate_root!
        raise ArgumentError, "root must not be /" if root.cleanpath.to_s == "/"
      end

      def install_desktop_entry
        desktop_entry = DesktopEntry.new(recipe)
        install_file(
          relative_path: File.join(APPLICATIONS_DIR, desktop_entry.file_name),
          contents: desktop_entry.render,
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

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @root = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--root is required" unless @root
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          result = DesktopActivationInstaller.new(root: @root, recipe: recipe).install
          puts JSON.pretty_generate(result)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-install-desktop-integration: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-install-desktop-integration --root PATH --app APP_ID [--recipe-dir PATH]"
            options.on("--root PATH", "Install desktop activation files under PATH") do |value|
              @root = value
            end
            options.on("--app APP_ID", "Install desktop activation files for APP_ID") do |value|
              @application_id = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
          end
        end
      end
    end
  end
end
