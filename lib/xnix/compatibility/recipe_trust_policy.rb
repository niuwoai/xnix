# frozen_string_literal: true

require "json"
require "optparse"
require_relative "recipe_registry"

module Xnix
  module Compatibility
    class RecipeTrustPolicy
      attr_reader :registry_report

      def initialize(registry_report:)
        @registry_report = registry_report
      end

      def to_h
        {
          "version" => RecipeRegistry::VERSION,
          "policy_type" => "recipe-trust",
          "decision" => decision,
          "recipe_count" => registry_report.fetch("recipe_count"),
          "checks" => checks,
          "blocking_reasons" => blocking_reasons,
          "next_requirements" => next_requirements,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def trust
        registry_report.fetch("trust")
      end

      def decision
        return "production-trusted" if trust.fetch("digest_verified") && trust.fetch("signed_recipe_validation")
        return "development-only" if trust.fetch("digest_verified") && trust.fetch("development_registry")

        "untrusted"
      end

      def checks
        [
          {
            "id" => "registry.digest",
            "status" => trust.fetch("digest_verified") ? "pass" : "fail",
            "message" => "Recipe digests are verified against registry metadata."
          },
          {
            "id" => "registry.signature",
            "status" => trust.fetch("signed_recipe_validation") ? "pass" : "pending",
            "message" => "Production signed recipe validation is required before external recipes are trusted."
          },
          {
            "id" => "registry.development",
            "status" => trust.fetch("development_registry") ? "warn" : "pass",
            "message" => "Development-only recipes are allowed for local testing but not for production trust."
          }
        ]
      end

      def blocking_reasons
        reasons = []
        reasons << "recipe digests are not verified" unless trust.fetch("digest_verified")
        reasons << "production signed recipe validation is not enabled" unless trust.fetch("signed_recipe_validation")
        reasons << "registry contains development-only recipes" if trust.fetch("development_registry")
        reasons
      end

      def next_requirements
        [
          "Add production trust roots for recipe signatures.",
          "Require verified signatures before accepting external recipe registries.",
          "Bind recipe trust decisions to install and activation requests."
        ]
      end

      def desktop_safe_summary
        case decision
        when "production-trusted"
          "Recipes are verified for production use."
        when "development-only"
          "Recipes are verified for local development only."
        else
          "Recipes are not trusted."
        end
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @registry_path = RecipeRegistry::CLI::DEFAULT_REGISTRY
        end

        def run
          parser.order!(@argv)
          command = @argv.shift || "evaluate"

          case command
          when "evaluate"
            report = RecipeRegistry.new(path: @registry_path).verify
            puts RecipeTrustPolicy.new(registry_report: report).to_json
          else
            warn "unknown command: #{command}"
            return 64
          end

          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-recipe-trust-policy: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-recipe-trust-policy [--registry PATH] evaluate"
            options.on("--registry PATH", "Read recipe registry metadata from PATH") do |value|
              @registry_path = value
            end
          end
        end
      end
    end
  end
end
