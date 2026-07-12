# frozen_string_literal: true

require "json"
require "optparse"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityEngineCatalog
      ENGINES = [
        {
          "id" => "automatic-managed",
          "label" => "Automatic",
          "kind" => "orchestrator",
          "ready" => false,
          "launch_enabled" => false,
          "user_visible" => true,
          "summary" => "Xnix will choose the best available compatibility path."
        },
        {
          "id" => "local-compatibility-engine",
          "label" => "Local compatibility engine",
          "kind" => "local",
          "ready" => false,
          "launch_enabled" => false,
          "user_visible" => true,
          "summary" => "Local compatibility engine binding is pending."
        },
        {
          "id" => "isolated-compatibility-engine",
          "label" => "Isolated compatibility engine",
          "kind" => "isolated",
          "ready" => false,
          "launch_enabled" => false,
          "user_visible" => true,
          "summary" => "Isolated compatibility engine binding is pending."
        }
      ].freeze

      MODE_TO_ENGINE = {
        "automatic" => "automatic-managed",
        "wine" => "local-compatibility-engine",
        "vm" => "isolated-compatibility-engine"
      }.freeze

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "catalog_type" => "compatibility-engine",
          "runtime_policy_owner" => true,
          "desktop_shell_policy_owner" => false,
          "backend_details_exposed" => false,
          "engines" => ENGINES
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      def select_for_mode(mode)
        engine_id = MODE_TO_ENGINE.fetch(mode)
        engine = ENGINES.find { |entry| entry.fetch("id") == engine_id }
        {
          "engine_id" => engine.fetch("id"),
          "label" => engine.fetch("label"),
          "kind" => engine.fetch("kind"),
          "ready" => engine.fetch("ready"),
          "launch_enabled" => engine.fetch("launch_enabled"),
          "backend_details_exposed" => false,
          "summary" => engine.fetch("summary")
        }
      rescue KeyError
        raise ArgumentError, "mode must be one of: #{MODE_TO_ENGINE.keys.join(", ")}"
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @mode = nil
        end

        def run
          parser.parse!(@argv)
          catalog = CompatibilityEngineCatalog.new
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          value = @mode ? catalog.select_for_mode(@mode) : catalog.to_h
          puts JSON.pretty_generate(value)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-engine-catalog: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-engine-catalog [--mode MODE]"
            options.on("--mode MODE", MODE_TO_ENGINE.keys, "Select the engine for a recipe mode") do |value|
              @mode = value
            end
          end
        end
      end
    end
  end
end
