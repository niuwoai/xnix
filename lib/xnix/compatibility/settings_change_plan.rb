# frozen_string_literal: true

require "json"
require "optparse"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class SettingsChangePlan
      FIELD_RULES = {
        "run-mode" => {
          "mode" => %w[automatic performance compatibility],
          "preference" => %w[performance compatibility]
        },
        "resource-access" => {
          "documents" => %w[allow ask deny],
          "downloads" => %w[allow ask deny]
        },
        "devices" => {
          "camera" => %w[allow ask deny]
        },
        "network" => {
          "network" => %w[allow ask deny]
        },
        "snapshots" => {
          "snapshots" => %w[enabled disabled]
        }
      }.freeze

      PORTAL_REVIEW_SECTIONS = %w[resource-access devices network].freeze
      SNAPSHOT_RECOMMENDED_SECTIONS = %w[run-mode snapshots].freeze
      CONFIRMATION_SECTIONS = %w[run-mode resource-access devices network snapshots].freeze

      def initialize(application_id:, section_id:, field_id:, value:)
        @application_id = application_id
        @section_id = section_id
        @field_id = field_id
        @value = value
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "settings-change-plan",
          "application_id" => application_id,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "section_id" => section_id,
          "field_id" => field_id,
          "requested_value" => value,
          "change_state" => "planned",
          "apply_enabled" => false,
          "settings_persisted" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "user_confirmation_required" => user_confirmation_required?,
          "snapshot_recommended" => snapshot_recommended?,
          "portal_policy_review_required" => portal_policy_review_required?,
          "runtime_restart_required" => false,
          "affected_policy" => affected_policy,
          "steps" => steps,
          "blocked_actions" => blocked_actions,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :application_id, :section_id, :field_id, :value

      def validate!
        unless application_id.is_a?(String) && application_id.match?(/\A[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+\z/)
          raise ArgumentError, "application id must be a reverse-DNS identifier"
        end

        fields = FIELD_RULES.fetch(section_id) do
          raise ArgumentError, "unknown settings section: #{section_id}"
        end

        options = fields.fetch(field_id) do
          raise ArgumentError, "unknown settings field: #{field_id}"
        end

        raise ArgumentError, "unsupported settings value: #{value}" unless options.include?(value)
      end

      def affected_policy
        {
          "section" => section_id,
          "field" => field_id,
          "value" => value,
          "options" => FIELD_RULES.fetch(section_id).fetch(field_id)
        }
      end

      def steps
        [
          step("validate-setting", "pass", "The requested settings value is valid for the Runtime settings schema."),
          step("review-user-confirmation", user_confirmation_required? ? "required" : "pass", "KDE must present the change for user review before persistence."),
          step("review-portal-policy", portal_policy_review_required? ? "required" : "pass", "Runtime Portal policy must be reviewed before desktop resource access changes."),
          step("prepare-restore-point", snapshot_recommended? ? "recommended" : "pass", "Runtime should prepare a restore point before risky compatibility settings changes."),
          step("persist-runtime-setting", "pending", "Runtime persistence is not enabled in this version.")
        ]
      end

      def step(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def blocked_actions
        [
          "persist compatibility settings before Runtime confirmation",
          "grant desktop resources without Portal policy review",
          "modify host root while planning settings changes",
          "expose backend implementation settings to KDE"
        ]
      end

      def user_confirmation_required?
        CONFIRMATION_SECTIONS.include?(section_id)
      end

      def snapshot_recommended?
        SNAPSHOT_RECOMMENDED_SECTIONS.include?(section_id)
      end

      def portal_policy_review_required?
        PORTAL_REVIEW_SECTIONS.include?(section_id)
      end

      def desktop_safe_summary
        "Compatibility settings change is planned and waiting for Runtime persistence support."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @section_id = nil
          @field_id = nil
          @value = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--section is required" unless @section_id
          raise ArgumentError, "--field is required" unless @field_id
          raise ArgumentError, "--value is required" unless @value

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts SettingsChangePlan.new(
            application_id: recipe.id,
            section_id: @section_id,
            field_id: @field_id,
            value: @value
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-settings-change: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-settings-change --app APP_ID --section SECTION --field FIELD --value VALUE [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime-owned settings change plan for APP_ID") do |value|
              @application_id = value
            end
            options.on("--section SECTION", "Settings section identifier") do |value|
              @section_id = value
            end
            options.on("--field FIELD", "Settings field identifier") do |value|
              @field_id = value
            end
            options.on("--value VALUE", "Requested settings value") do |value|
              @value = value
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
