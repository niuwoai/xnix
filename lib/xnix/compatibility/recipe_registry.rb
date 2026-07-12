# frozen_string_literal: true

require "digest"
require "json"
require "optparse"
require "pathname"
require_relative "application_recipe"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class RecipeRegistry
      SCHEMA_VERSION = 1
      DIGEST_PATTERN = /\A[0-9a-f]{64}\z/
      SIGNATURE_STATUSES = %w[development-only verified].freeze

      attr_reader :path

      def initialize(path:)
        @path = Pathname.new(path)
      end

      def verify
        validate_manifest!

        entries = recipes.map { |entry| verify_entry(entry) }

        {
          "version" => RuntimeDaemon::VERSION,
          "schema_version" => SCHEMA_VERSION,
          "registry_name" => manifest.fetch("registry_name"),
          "registry_path" => path.to_s,
          "recipe_count" => entries.length,
          "recipes" => entries,
          "trust" => {
            "digest_verified" => entries.all? { |entry| entry.fetch("digest_verified") },
            "signed_recipe_validation" => entries.all? { |entry| entry.fetch("signature_status") == "verified" },
            "development_registry" => entries.any? { |entry| entry.fetch("signature_status") == "development-only" }
          }
        }
      end

      def verify!
        verify
      end

      def recipes
        manifest.fetch("recipes")
      end

      private

      def validate_manifest!
        raise ArgumentError, "missing recipe registry: #{path}" unless path.file?
        raise ArgumentError, "recipe registry must be an object" unless manifest.is_a?(Hash)
        raise ArgumentError, "unsupported registry schema version" unless manifest.fetch("schema_version") == SCHEMA_VERSION
        raise ArgumentError, "registry_name must be a non-empty string" unless non_empty_string?(manifest.fetch("registry_name"))
        raise ArgumentError, "recipes must be an array" unless recipes.is_a?(Array)

        ids = recipes.map { |entry| entry.fetch("id") }
        raise ArgumentError, "recipe ids must be unique" unless ids.uniq == ids
      rescue KeyError => e
        raise ArgumentError, "recipe registry missing #{e.key}"
      end

      def verify_entry(entry)
        validate_entry!(entry)
        recipe_path = safe_recipe_path(entry.fetch("path"))
        actual_sha256 = Digest::SHA256.file(recipe_path).hexdigest
        expected_sha256 = entry.fetch("sha256")
        raise ArgumentError, "recipe digest mismatch: #{entry.fetch("id")}" unless actual_sha256 == expected_sha256

        {
          "id" => entry.fetch("id"),
          "path" => entry.fetch("path"),
          "sha256" => actual_sha256,
          "digest_verified" => true,
          "signature_status" => entry.fetch("signature_status")
        }
      end

      def validate_entry!(entry)
        raise ArgumentError, "recipe registry entry must be an object" unless entry.is_a?(Hash)

        id = entry.fetch("id")
        relative_path = entry.fetch("path")
        sha256 = entry.fetch("sha256")
        signature_status = entry.fetch("signature_status")

        raise ArgumentError, "recipe id must be a reverse-DNS identifier" unless id.is_a?(String) && id.match?(ApplicationRecipe::ID_PATTERN)
        raise ArgumentError, "recipe path must be a relative JSON file" unless valid_relative_json_path?(relative_path)
        raise ArgumentError, "recipe digest must be a SHA-256 hex string" unless sha256.is_a?(String) && sha256.match?(DIGEST_PATTERN)
        raise ArgumentError, "unsupported recipe signature status" unless SIGNATURE_STATUSES.include?(signature_status)
      rescue KeyError => e
        raise ArgumentError, "recipe registry entry missing #{e.key}"
      end

      def manifest
        @manifest ||= JSON.parse(path.read)
      rescue JSON::ParserError => e
        raise ArgumentError, "invalid recipe registry JSON: #{e.message}"
      end

      def safe_recipe_path(relative_path)
        destination = path.dirname.join(relative_path).cleanpath
        unless destination.to_s.start_with?("#{path.dirname.cleanpath}#{File::SEPARATOR}")
          raise ArgumentError, "recipe path must stay inside the recipe directory"
        end
        raise ArgumentError, "missing recipe file: #{relative_path}" unless destination.file?

        destination
      end

      def valid_relative_json_path?(relative_path)
        return false unless relative_path.is_a?(String)
        return false unless relative_path.end_with?(".json")
        return false if Pathname.new(relative_path).absolute?
        return false if relative_path.split(File::SEPARATOR).include?("..")

        true
      end

      def non_empty_string?(value)
        value.is_a?(String) && !value.empty? && !value.match?(/[\r\n]/)
      end

      class CLI
        DEFAULT_REGISTRY = RuntimeDaemon::PROJECT_ROOT.join("runtime/recipes/registry.json").to_s

        def initialize(argv)
          @argv = argv.dup
          @registry_path = DEFAULT_REGISTRY
        end

        def run
          parser.order!(@argv)
          command = @argv.shift || "verify"
          registry = RecipeRegistry.new(path: @registry_path)

          case command
          when "verify"
            puts JSON.pretty_generate(registry.verify)
          else
            warn "unknown command: #{command}"
            return 64
          end

          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-recipe-registry: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-recipe-registry [--registry PATH] verify"
            options.on("--registry PATH", "Read recipe registry metadata from PATH") do |value|
              @registry_path = value
            end
          end
        end
      end
    end
  end
end
