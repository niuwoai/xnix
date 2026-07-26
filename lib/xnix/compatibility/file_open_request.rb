# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "uri"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class FileOpenRequest
      SOURCE = "dolphin-service-menu"

      attr_reader :recipe_store

      def initialize(recipe_store:)
        @recipe_store = recipe_store
      end

      def build(file_uris:, application_id: nil)
        uris = normalize_file_uris(file_uris)
        recipe = application_id ? require_recipe(application_id) : select_recipe_for(uris)

        {
          "request_type" => "open-file",
          "source" => SOURCE,
          "application_id" => recipe.id,
          "application_name" => recipe.name,
          "runtime_method" => "Launch",
          "portal_required" => true,
          "file_count" => uris.length,
          "file_uris" => uris
        }
      end

      private

      def normalize_file_uris(file_uris)
        raise ArgumentError, "at least one file URI is required" if file_uris.empty?

        file_uris.map { |file_uri| normalize_file_uri(file_uri) }
      end

      def normalize_file_uri(file_uri)
        uri = URI.parse(file_uri)
        raise ArgumentError, "only file URIs are accepted" unless uri.scheme == "file"
        raise ArgumentError, "file URI must include an absolute path" if uri.path.nil? || uri.path.empty? || !uri.path.start_with?("/")

        uri.to_s
      rescue URI::InvalidURIError
        raise ArgumentError, "invalid file URI: #{file_uri.inspect}"
      end

      def select_recipe_for(file_uris)
        extension = File.extname(URI.parse(file_uris.first).path).downcase
        raise ArgumentError, "selected file must have an extension" if extension.empty?

        matches = recipe_store.all.select { |recipe| recipe.supported_extensions.map(&:downcase).include?(extension) }
        raise ArgumentError, "no compatible application is registered for #{extension}" if matches.empty?

        matches.first
      end

      def require_recipe(application_id)
        recipe = recipe_store.find(application_id)
        return recipe if recipe

        raise ArgumentError, "unknown application: #{application_id}"
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @application_id = nil
          @preview_engine = "auto"
          @runtime_bin = ENV.fetch("XNIX_RUNTIME_GO", "xnix-runtime-go")
          @runtime_registry = nil
        end

        def run
          parser.parse!(@argv)
          runtime_result = run_go_runtime_preview
          if runtime_result
            puts runtime_result
            return 0
          end

          request = FileOpenRequest.new(recipe_store: RecipeStore.new(path: @recipe_dir))
                                   .build(file_uris: @argv, application_id: @application_id)
          puts JSON.pretty_generate(request)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-open: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-open [--recipe-dir PATH] [--app APP_ID] FILE_URI [FILE_URI ...]"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--app APP_ID", "Open with a specific Runtime application") do |value|
              @application_id = value
            end
            options.on("--preview-engine ENGINE", "Preview engine: auto, go, or ruby") do |value|
              @preview_engine = value
            end
            options.on("--runtime-bin PATH", "Go Runtime preview command") do |value|
              @runtime_bin = value
            end
            options.on("--runtime-registry PATH", "Go Runtime recipe registry") do |value|
              @runtime_registry = value
            end
          end
        end

        def run_go_runtime_preview
          raise ArgumentError, "preview engine must be auto, go, or ruby" unless %w[auto go ruby].include?(@preview_engine)
          return nil if @preview_engine == "ruby"

          registry = runtime_registry_path
          runtime = resolve_runtime_bin
          if runtime.nil? || registry.nil?
            raise ArgumentError, "Go Runtime file-open preview is unavailable" if @preview_engine == "go"

            return nil
          end

          args = [runtime, "file-open-preview", "--registry", registry.to_s]
          args += ["--app", @application_id] if @application_id
          args += @argv
          stdout, status = Open3.capture2(*args)
          raise ArgumentError, "Go Runtime file-open preview failed" unless status.success?

          stdout
        rescue Errno::ENOENT
          raise ArgumentError, "Go Runtime file-open preview is unavailable" if @preview_engine == "go"

          nil
        end

        def runtime_registry_path
          explicit = @runtime_registry && Pathname.new(@runtime_registry)
          return explicit if explicit && explicit.file?

          registry = Pathname.new(@recipe_dir).join("registry.json")
          return registry if registry.file?

          nil
        end

        def resolve_runtime_bin
          path = Pathname.new(@runtime_bin)
          return path.to_s if path.absolute? && path.executable?
          return path.to_s if path.dirname.to_s != "." && path.executable?

          ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).each do |dir|
            candidate = Pathname.new(dir).join(@runtime_bin)
            return candidate.to_s if candidate.executable?
          end
          nil
        end
      end
    end
  end
end
