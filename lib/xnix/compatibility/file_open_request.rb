# frozen_string_literal: true

require "json"
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
        end

        def run
          parser.parse!(@argv)
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
          end
        end
      end
    end
  end
end
