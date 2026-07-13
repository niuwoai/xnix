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
        plan = runtime_query_plan
        resolved_matches = matches_from_plan(plan)

        {
          "version" => RuntimeDaemon::VERSION,
          "query_type" => plan.fetch("query_type", "krunner-query-plan"),
          "entry_point" => plan.fetch("entry_point", "krunner"),
          "desktop" => plan.fetch("desktop", "KDE Plasma"),
          "query" => plan.fetch("query", query),
          "source" => source_metadata,
          "runtime_owned" => plan.fetch("runtime_owned", true),
          "kde_policy_owner" => plan.fetch("kde_policy_owner", false),
          "matches" => resolved_matches,
          "summary" => summary_from_plan(plan, resolved_matches),
          "host_root_modified" => plan.fetch("host_root_modified", false),
          "backend_details_exposed" => plan.fetch("backend_details_exposed", false),
          "desktop_safe_summary" => plan.fetch(
            "desktop_safe_summary",
            "KRunner query planning is Runtime-owned and returns safe launcher actions only."
          )
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

      def runtime_query_plan
        return runtime.krunner_query_plan(query) if runtime.respond_to?(:krunner_query_plan)

        fallback_query_plan
      end

      def fallback_query_plan
        normalized_query = normalize(query)
        resolved_matches = if normalized_query.empty?
                             []
                           else
                             runtime.list_applications.filter_map do |application|
                               relevance = relevance_for(application, normalized_query)
                               next if relevance.zero?

                               match_for(application, relevance)
                             end.sort_by { |match| [-match.fetch("relevance"), match.fetch("name")] }
                           end

        {
          "query_type" => "krunner-query-plan",
          "entry_point" => "krunner",
          "desktop" => "KDE Plasma",
          "query" => query,
          "matches" => resolved_matches,
          "summary" => {
            "match_count" => resolved_matches.length,
            "official_desktop" => "KDE Plasma",
            "runtime_owned_launch" => true,
            "query_execution_enabled" => false,
            "backend_launch_enabled" => false,
            "backend_details_exposed" => false
          },
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KRunner query planning is Runtime-owned and returns safe launcher actions only."
        }
      end

      def matches_from_plan(plan)
        plan_matches = plan.fetch("matches", nil)
        return plan_matches.map { |match| normalized_plan_match(match) } if plan_matches.is_a?(Array)

        return [] if plan.fetch("match_count", 0).to_i.zero?

        [normalized_flat_plan_match(plan)]
      end

      def normalized_plan_match(match)
        relevance_percent = match.fetch("relevance_percent", (match.fetch("relevance", 0.0).to_f * 100).round)
        application_id = match.fetch("application_id")
        action = match.fetch("action", {})

        {
          "runner_id" => match.fetch("runner_id", "xnix.compatibility.#{application_id}"),
          "application_id" => application_id,
          "name" => match.fetch("name"),
          "icon" => match.fetch("icon", "application-x-executable"),
          "relevance" => relevance_percent.to_f / 100.0,
          "relevance_percent" => relevance_percent,
          "subtitle" => match.fetch("subtitle", "Open as a normal Linux application"),
          "mode_label" => match.fetch("mode_label", "Automatic"),
          "supported_extensions" => match.fetch("supported_extensions", []),
          "runtime_owned_launch" => match.fetch("runtime_owned_launch", true),
          "backend_details_exposed" => match.fetch("backend_details_exposed", false),
          "action" => {
            "type" => action.fetch("type", match.fetch("action_type", "runtime-launch")),
            "desktop_entry_id" => action.fetch("desktop_entry_id", match.fetch("desktop_entry_id", "#{application_id}.desktop")),
            "argv" => action.fetch("argv", ["xnix-compat-launch", "--app", application_id])
          }
        }
      end

      def normalized_flat_plan_match(plan)
        application_id = plan.fetch("top_application_id")
        relevance_percent = plan.fetch("top_relevance_percent", 0).to_i

        normalized_plan_match(
          "runner_id" => "xnix.compatibility.#{application_id}",
          "application_id" => application_id,
          "name" => plan.fetch("top_name"),
          "relevance_percent" => relevance_percent,
          "action_type" => plan.fetch("action_type", "runtime-launch"),
          "desktop_entry_id" => plan.fetch("desktop_entry_id", "#{application_id}.desktop"),
          "runtime_owned_launch" => plan.fetch("runtime_owned_launch", true),
          "backend_details_exposed" => plan.fetch("backend_details_exposed", false)
        )
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

      def summary_from_plan(plan, resolved_matches)
        plan_summary = plan.fetch("summary", {})

        {
          "match_count" => resolved_matches.length,
          "official_desktop" => plan_summary.fetch("official_desktop", plan.fetch("desktop", "KDE Plasma")),
          "runtime_owned_launch" => plan_summary.fetch("runtime_owned_launch", plan.fetch("runtime_owned_launch", true)),
          "query_execution_enabled" => plan_summary.fetch("query_execution_enabled", plan.fetch("query_execution_enabled", false)),
          "backend_launch_enabled" => plan_summary.fetch("backend_launch_enabled", plan.fetch("backend_launch_enabled", false)),
          "backend_details_exposed" => plan_summary.fetch("backend_details_exposed", plan.fetch("backend_details_exposed", false))
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
