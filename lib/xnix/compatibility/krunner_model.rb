# frozen_string_literal: true

require "json"
require "optparse"
require_relative "dbus_runtime_client"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class KRunnerModel
      SOURCES = %w[auto local dbus].freeze
      MODE_LABELS = {
        "automatic" => "Automatic",
        "wine" => "Managed compatibility",
        "vm" => "Isolated environment"
      }.freeze

      attr_reader :runtime, :query

      def initialize(runtime:, query:)
        @runtime = runtime
        @query = query.to_s
      end

      def to_h
        resolved_matches = matches

        {
          "version" => RuntimeDaemon::VERSION,
          "entry_point" => "krunner",
          "desktop" => "KDE Plasma",
          "query" => query,
          "source" => source_metadata,
          "matches" => resolved_matches,
          "summary" => summary(resolved_matches)
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def source_metadata
        return runtime.source_metadata if runtime.respond_to?(:source_metadata)

        {
          "kind" => "runtime-local-read-model",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def matches
        normalized_query = normalize(query)
        return [] if normalized_query.empty?

        runtime.list_applications.filter_map do |application|
          relevance = relevance_for(application, normalized_query)
          next if relevance.zero?

          match_for(application, relevance)
        end.sort_by { |match| [-match.fetch("relevance"), match.fetch("name")] }
      end

      def relevance_for(application, normalized_query)
        name = normalize(application.fetch("name"))
        application_id = normalize(application.fetch("id"))
        extensions = extensions_for(application).map { |extension| normalize(extension.delete_prefix(".")) }

        return 1.0 if name == normalized_query || application_id == normalized_query
        return 0.95 if name.start_with?(normalized_query)
        return 0.9 if name.include?(normalized_query)
        return 0.85 if extensions.include?(normalized_query.delete_prefix("."))
        return 0.75 if application_id.include?(normalized_query)
        return 0.65 if natural_launch_query?(normalized_query, name)
        return 0.55 if extension_launch_query?(normalized_query, extensions)

        0.0
      end

      def natural_launch_query?(normalized_query, normalized_name)
        %w[open launch start run].any? do |verb|
          normalized_query == "#{verb} #{normalized_name}" || normalized_query.end_with?(" #{normalized_name}")
        end
      end

      def extension_launch_query?(normalized_query, normalized_extensions)
        tokens = normalized_query.split
        return false unless (tokens & %w[open launch start run file]).any?

        normalized_extensions.any? { |extension| tokens.include?(extension) }
      end

      def match_for(application, relevance)
        application_id = application.fetch("id")

        {
          "runner_id" => "xnix.compatibility.#{application_id}",
          "application_id" => application_id,
          "name" => application.fetch("name"),
          "icon" => application.fetch("icon"),
          "relevance" => relevance,
          "subtitle" => "Open as a normal Linux application",
          "mode_label" => MODE_LABELS.fetch(application.fetch("mode"), "Automatic"),
          "supported_extensions" => extensions_for(application),
          "action" => {
            "type" => "runtime-launch",
            "desktop_entry_id" => "#{application_id}.desktop",
            "argv" => ["xnix-compat-launch", "--app", application_id]
          }
        }
      end

      def extensions_for(application)
        value = application.fetch("supported_extensions", [])
        return value if value.is_a?(Array)

        []
      end

      def summary(resolved_matches)
        {
          "match_count" => resolved_matches.length,
          "official_desktop" => "KDE Plasma",
          "runtime_owned_launch" => true,
          "backend_details_exposed" => false
        }
      end

      def normalize(value)
        value.to_s.downcase.strip.gsub(/\s+/, " ")
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @source = "auto"
          @query = ""
        end

        def run
          parser.parse!(@argv)
          puts KRunnerModel.new(runtime: runtime_source, query: @query).to_json
          0
        rescue OptionParser::ParseError, KeyError, ArgumentError, DBusRuntimeClient::Error => e
          warn "xnix-krunner-model: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-krunner-model [--query TEXT] [--source SOURCE] [--recipe-dir PATH]"
            options.on("--query TEXT", "Resolve a KRunner query through Runtime applications") do |value|
              @query = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--source SOURCE", "Read from auto, local, or dbus") do |value|
              raise OptionParser::InvalidArgument, "source must be one of: #{SOURCES.join(", ")}" unless SOURCES.include?(value)

              @source = value
            end
          end
        end

        def runtime_source
          case @source
          when "local"
            local_runtime
          when "dbus"
            DBusRuntimeClient.new
          when "auto"
            DBusRuntimeClient.available? ? DBusRuntimeClient.new : local_runtime
          end
        end

        def local_runtime
          RuntimeDaemon.new(recipe_store: RecipeStore.new(path: @recipe_dir))
        end
      end
    end
  end
end
