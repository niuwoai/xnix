# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_recipe"
require_relative "recipe_registry"
require_relative "recipe_trust_policy"

module Xnix
  module Compatibility
    class RecipeInstallGate
      MODES = %w[production development].freeze

      attr_reader :registry_report, :application_id, :mode

      def initialize(registry_report:, application_id:, mode: "production")
        @registry_report = registry_report
        @application_id = application_id
        @mode = mode
        validate!
      end

      def to_h
        {
          "version" => RecipeRegistry::VERSION,
          "gate_type" => "recipe-install",
          "application_id" => application_id,
          "mode" => mode,
          "decision" => decision,
          "policy_decision" => policy.fetch("decision"),
          "matched_recipe" => matched_recipe_summary,
          "blocking_reasons" => blocking_reasons,
          "requirements" => requirements,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def validate!
        unless application_id.is_a?(String) && application_id.match?(ApplicationRecipe::ID_PATTERN)
          raise ArgumentError, "application id must be a reverse-DNS identifier"
        end
        raise ArgumentError, "mode must be one of: #{MODES.join(", ")}" unless MODES.include?(mode)
      end

      def decision
        blocking_reasons.empty? ? "allow" : "block"
      end

      def matched_recipe
        @matched_recipe ||= registry_report.fetch("recipes").find { |recipe| recipe.fetch("id") == application_id }
      end

      def matched_recipe_summary
        return nil unless matched_recipe

        {
          "id" => matched_recipe.fetch("id"),
          "signature_status" => matched_recipe.fetch("signature_status"),
          "digest_verified" => matched_recipe.fetch("digest_verified")
        }
      end

      def policy
        @policy ||= RecipeTrustPolicy.new(registry_report: registry_report).to_h
      end

      def blocking_reasons
        reasons = []
        reasons << "recipe is not registered" unless matched_recipe
        reasons << "recipe digest is not verified" if matched_recipe && !matched_recipe.fetch("digest_verified")
        reasons.concat(production_blocking_reasons) if mode == "production"
        reasons.concat(development_blocking_reasons) if mode == "development"
        reasons.uniq
      end

      def production_blocking_reasons
        return [] if policy.fetch("decision") == "production-trusted" &&
                     matched_recipe &&
                     matched_recipe.fetch("signature_status") == "verified"

        policy.fetch("blocking_reasons")
      end

      def development_blocking_reasons
        return [] if matched_recipe && matched_recipe.fetch("digest_verified") &&
                     %w[development-only production-trusted].include?(policy.fetch("decision"))

        policy.fetch("blocking_reasons")
      end

      def requirements
        return [] if decision == "allow"

        if mode == "production"
          [
            "Use a registry with a production signed source and verified recipe signatures.",
            "Keep SHA-256 digest verification enabled before activation.",
            "Do not promote development-only recipes into production installation."
          ]
        else
          [
            "Register the application recipe in the selected registry.",
            "Keep SHA-256 digest verification enabled for development staging."
          ]
        end
      end

      def desktop_safe_summary
        return "Recipe installation may continue." if decision == "allow"

        "Recipe installation is blocked until trust requirements are satisfied."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @registry_path = RecipeRegistry::CLI::DEFAULT_REGISTRY
          @application_id = nil
          @mode = "production"
        end

        def run
          parser.order!(@argv)
          command = @argv.shift || "evaluate"
          raise ArgumentError, "--app is required" unless @application_id

          case command
          when "evaluate"
            report = RecipeRegistry.new(path: @registry_path).verify
            puts RecipeInstallGate.new(
              registry_report: report,
              application_id: @application_id,
              mode: @mode
            ).to_json
          else
            warn "unknown command: #{command}"
            return 64
          end

          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-recipe-install-gate: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-recipe-install-gate --app ID [--registry PATH] [--mode MODE] evaluate"
            options.on("--registry PATH", "Read recipe registry metadata from PATH") do |value|
              @registry_path = value
            end
            options.on("--app ID", "Evaluate the installation gate for application recipe ID") do |value|
              @application_id = value
            end
            options.on("--mode MODE", MODES, "Evaluate production or development installation policy") do |value|
              @mode = value
            end
          end
        end
      end
    end
  end
end
