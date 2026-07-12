# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "registry_backed_recipe_store"

module Xnix
  module Compatibility
    class RuntimeDaemon
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s
      CONTRACT_PATH = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.xml")
      BUS_NAME = "org.xnix.Compatibility1"
      OBJECT_PATH = "/org/xnix/Compatibility1"
      INTERFACE = "org.xnix.Compatibility1"

      attr_reader :recipe_store

      def initialize(recipe_store:)
        @recipe_store = recipe_store
      end

      def probe
        {
          "version" => VERSION,
          "bus_name" => BUS_NAME,
          "object_path" => OBJECT_PATH,
          "interface" => INTERFACE,
          "recipe_count" => list_applications.length,
          "recipe_trust" => recipe_trust,
          "capabilities" => {
            "recipe_store" => true,
            "registry_backed_recipe_store" => registry_backed_recipe_store?,
            "application_listing" => true,
            "compatibility_run_planning" => true,
            "diagnostics" => true,
            "dbus_method_dispatch" => true,
            "dbus_binding" => false,
            "wine_backend" => false,
            "vm_backend" => false
          }
        }
      end

      def list_applications
        recipe_store.all.map { |recipe| recipe_to_hash(recipe) }
      end

      def get_application(application_id)
        recipe = require_recipe(application_id)
        recipe_to_hash(recipe)
      end

      def diagnostics(application_id)
        recipe = require_recipe(application_id)
        {
          "application_id" => recipe.id,
          "status" => "known",
          "runtime_mode" => recipe.mode,
          "checks" => [
            {
              "id" => "recipe.validation",
              "status" => "pass",
              "message" => "Recipe is valid and can be exposed through desktop integration."
            },
            {
              "id" => "backend.binding",
              "status" => "pending",
              "message" => "Wine and VM launch backends are not enabled in this version."
            }
          ]
        }
      end

      def dispatch(method_name, parameters = [])
        case method_name
        when "ListApplications"
          list_applications
        when "GetApplication"
          get_application(required_parameter(method_name, parameters, 0))
        when "GetDiagnostics"
          diagnostics(required_parameter(method_name, parameters, 0))
        else
          raise ArgumentError, "unsupported runtime method: #{method_name}"
        end
      end

      def introspection_xml
        CONTRACT_PATH.read
      end

      private

      def required_parameter(method_name, parameters, index)
        value = parameters[index]
        raise ArgumentError, "#{method_name} requires parameter #{index + 1}" if value.nil?

        value
      end

      def require_recipe(application_id)
        recipe = recipe_store.find(application_id)
        return recipe if recipe

        raise ArgumentError, "unknown application: #{application_id}"
      end

      def recipe_to_hash(recipe)
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon,
          "mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions,
          "mime_types" => recipe.mime_types
        }
      end

      def registry_backed_recipe_store?
        recipe_store.respond_to?(:registry_report)
      end

      def recipe_trust
        unless registry_backed_recipe_store?
          return {
            "registry_backed" => false,
            "digest_verified" => false,
            "signed_recipe_validation" => false,
            "development_fallback" => true
          }
        end

        trust = recipe_store.registry_report.fetch("trust")
        {
          "registry_backed" => true,
          "digest_verified" => trust.fetch("digest_verified"),
          "signed_recipe_validation" => trust.fetch("signed_recipe_validation"),
          "development_registry" => trust.fetch("development_registry")
        }
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = DEFAULT_RECIPE_DIR
        end

        def run
          parser.order!(@argv)
          command = @argv.shift || "probe"
          runtime = RuntimeDaemon.new(recipe_store: RegistryBackedRecipeStore.for_path(@recipe_dir))

          case command
          when "probe"
            write_json(runtime.probe)
          when "list-applications"
            write_json(runtime.list_applications)
          when "get-application"
            write_json(runtime.get_application(require_argument(command)))
          when "diagnostics"
            write_json(runtime.diagnostics(require_argument(command)))
          when "dispatch"
            method_name = require_argument(command)
            parameters = @argv.empty? ? [] : JSON.parse(@argv.shift)
            raise ArgumentError, "dispatch parameters must be an array" unless parameters.is_a?(Array)

            write_json(runtime.dispatch(method_name, parameters))
          when "introspect"
            puts runtime.introspection_xml
          else
            warn "unknown command: #{command}"
            return 64
          end

          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compatd: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compatd [--recipe-dir PATH] COMMAND"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
          end
        end

        def require_argument(command)
          value = @argv.shift
          raise ArgumentError, "#{command} requires an application id" unless value

          value
        end

        def write_json(value)
          puts JSON.pretty_generate(value)
        end
      end
    end
  end
end
