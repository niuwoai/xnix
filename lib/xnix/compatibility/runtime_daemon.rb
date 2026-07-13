# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "application_state_root"
require_relative "ai_diagnostic_input"
require_relative "ai_diagnostic_recommendation"
require_relative "ai_repair_approval_gate"
require_relative "compatibility_acquisition_preflight"
require_relative "compatibility_artifact_manifest"
require_relative "compatibility_backend_binding"
require_relative "compatibility_engine_catalog"
require_relative "compatibility_install_plan"
require_relative "compatibility_package_source"
require_relative "compatibility_run_plan"
require_relative "compatibility_repair_plan"
require_relative "compatibility_snapshot_plan"
require_relative "compatibility_test_plan"
require_relative "compatibility_test_result"
require_relative "portal_access_policy"
require_relative "registry_backed_recipe_store"
require_relative "runtime_service_binding"
require_relative "settings_model"
require_relative "settings_change_plan"

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
            "application_state_roots" => true,
            "compatibility_acquisition_preflight" => true,
            "compatibility_artifact_manifests" => true,
            "compatibility_engine_catalog" => true,
            "compatibility_install_planning" => true,
            "compatibility_package_sources" => true,
            "compatibility_backend_binding" => true,
            "compatibility_run_planning" => true,
            "compatibility_repair_planning" => true,
            "compatibility_test_planning" => true,
            "compatibility_test_results" => true,
            "ai_diagnostic_inputs" => true,
            "ai_diagnostic_recommendations" => true,
            "ai_repair_approval_gates" => true,
            "runtime_service_binding" => true,
            "compatibility_settings" => true,
            "compatibility_settings_change_planning" => true,
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
              "id" => "engine.binding",
              "status" => "pending",
              "message" => "Compatibility engine launch binding is not enabled in this version."
            }
          ],
          "test_plan" => test_plan_summary(recipe),
          "test_result" => test_result_summary(recipe),
          "install_plan" => install_plan_summary(recipe),
          "acquisition_preflight" => acquisition_preflight_summary(recipe),
          "artifact_manifest" => artifact_manifest_summary(recipe),
          "package_source" => package_source_summary(recipe),
          "state_root" => state_root_summary(recipe),
          "backend_binding" => backend_binding_summary(recipe),
          "ai_diagnostic_input" => ai_diagnostic_input_summary(recipe),
          "ai_diagnostic_recommendation" => ai_diagnostic_recommendation_summary(recipe),
          "ai_repair_approval_gate" => ai_repair_approval_gate_summary(recipe),
          "runtime_service_binding" => runtime_service_binding_summary,
          "settings" => settings_summary(recipe),
          "settings_change_plan" => settings_change_plan_summary(recipe),
          "repair_plan" => repair_plan_summary(recipe.id, "engine-binding-pending")
        }
      end

      def engine_catalog
        CompatibilityEngineCatalog.new.to_h
      end

      def run_plan(application_id)
        recipe = require_recipe(application_id)
        CompatibilityRunPlan.new(recipe: recipe).to_h
      end

      def state_root(application_id)
        recipe = require_recipe(application_id)
        ApplicationStateRoot.new(recipe: recipe).to_h
      end

      def package_source(application_id)
        recipe = require_recipe(application_id)
        CompatibilityPackageSource.new(recipe: recipe).to_h
      end

      def acquisition_preflight(application_id)
        recipe = require_recipe(application_id)
        CompatibilityAcquisitionPreflight.new(recipe: recipe).to_h
      end

      def artifact_manifest(application_id)
        recipe = require_recipe(application_id)
        CompatibilityArtifactManifest.new(recipe: recipe).to_h
      end

      def install_plan(application_id, environment = "development")
        recipe = require_recipe(application_id)
        CompatibilityInstallPlan.new(recipe: recipe, environment: environment).to_h
      end

      def backend_binding(application_id)
        recipe = require_recipe(application_id)
        CompatibilityBackendBinding.new(recipe: recipe).to_h
      end

      def repair_plan(application_id, issue)
        require_recipe(application_id)
        CompatibilityRepairPlan.new(application_id: application_id, issue: issue).to_h
      end

      def test_plan(application_id, test_type = "preflight")
        recipe = require_recipe(application_id)
        CompatibilityTestPlan.new(recipe: recipe, test_type: test_type).to_h
      end

      def test_result(application_id, test_type = "preflight")
        recipe = require_recipe(application_id)
        CompatibilityTestResult.new(recipe: recipe, test_type: test_type).to_h
      end

      def ai_diagnostic_input(application_id, issue = AIDiagnosticInput::DEFAULT_ISSUE, test_type = "preflight")
        recipe = require_recipe(application_id)
        AIDiagnosticInput.new(recipe: recipe, issue: issue, test_type: test_type).to_h
      end

      def ai_diagnostic_recommendation(application_id, issue = AIDiagnosticInput::DEFAULT_ISSUE, test_type = "preflight")
        recipe = require_recipe(application_id)
        AIDiagnosticRecommendation.new(recipe: recipe, issue: issue, test_type: test_type).to_h
      end

      def ai_repair_approval_gate(application_id, issue = AIDiagnosticInput::DEFAULT_ISSUE, test_type = "preflight")
        recipe = require_recipe(application_id)
        AIRepairApprovalGate.new(recipe: recipe, issue: issue, test_type: test_type).to_h
      end

      def snapshot_plan(application_id, reason)
        require_recipe(application_id)
        CompatibilitySnapshotPlan.new(application_id: application_id, reason: reason).to_h
      end

      def portal_access_policy(application_id, operation)
        require_recipe(application_id)
        PortalAccessPolicy.new(application_id: application_id, operation: operation).to_h
      end

      def runtime_service_binding
        RuntimeServiceBinding.new.to_h
      end

      def settings(application_id)
        recipe = require_recipe(application_id)
        SettingsModel.new(application_id: recipe.id).to_h
      end

      def settings_change_plan(application_id, section_id, field_id, value)
        recipe = require_recipe(application_id)
        SettingsChangePlan.new(
          application_id: recipe.id,
          section_id: section_id,
          field_id: field_id,
          value: value
        ).to_h
      end

      def dispatch(method_name, parameters = [])
        case method_name
        when "ListApplications"
          list_applications
        when "GetApplication"
          get_application(required_parameter(method_name, parameters, 0))
        when "GetDiagnostics"
          diagnostics(required_parameter(method_name, parameters, 0))
        when "GetEngineCatalog"
          engine_catalog
        when "GetRunPlan"
          run_plan(required_parameter(method_name, parameters, 0))
        when "GetApplicationStateRoot"
          state_root(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityPackageSource"
          package_source(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityAcquisitionPreflight"
          acquisition_preflight(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityArtifactManifest"
          artifact_manifest(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityInstallPlan"
          install_plan(
            required_parameter(method_name, parameters, 0),
            parameters[1] || "development"
          )
        when "GetBackendBinding"
          backend_binding(required_parameter(method_name, parameters, 0))
        when "GetRepairPlan"
          repair_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetTestPlan"
          test_plan(
            required_parameter(method_name, parameters, 0),
            parameters[1] || "preflight"
          )
        when "GetTestResult"
          test_result(
            required_parameter(method_name, parameters, 0),
            parameters[1] || "preflight"
          )
        when "GetAIDiagnosticInput"
          ai_diagnostic_input(
            required_parameter(method_name, parameters, 0),
            parameters[1] || AIDiagnosticInput::DEFAULT_ISSUE,
            parameters[2] || "preflight"
          )
        when "GetAIDiagnosticRecommendation"
          ai_diagnostic_recommendation(
            required_parameter(method_name, parameters, 0),
            parameters[1] || AIDiagnosticInput::DEFAULT_ISSUE,
            parameters[2] || "preflight"
          )
        when "GetAIRepairApprovalGate"
          ai_repair_approval_gate(
            required_parameter(method_name, parameters, 0),
            parameters[1] || AIDiagnosticInput::DEFAULT_ISSUE,
            parameters[2] || "preflight"
          )
        when "GetSnapshotPlan"
          snapshot_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetPortalAccessPolicy"
          portal_access_policy(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetRuntimeServiceBinding"
          runtime_service_binding
        when "GetCompatibilitySettings"
          settings(required_parameter(method_name, parameters, 0))
        when "GetCompatibilitySettingsChangePlan"
          settings_change_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1),
            required_parameter(method_name, parameters, 2),
            required_parameter(method_name, parameters, 3)
          )
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

      def repair_plan_summary(application_id, issue)
        plan = CompatibilityRepairPlan.new(application_id: application_id, issue: issue).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "issue" => plan.fetch("issue"),
          "severity" => plan.fetch("severity"),
          "user_approval_required" => plan.fetch("user_approval_required"),
          "snapshot_required" => plan.fetch("snapshot_required"),
          "snapshot_plan" => plan.fetch("snapshot_plan"),
          "notification_event" => plan.fetch("notification_event"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def test_plan_summary(recipe)
        plan = CompatibilityTestPlan.new(recipe: recipe).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "test_type" => plan.fetch("test_type"),
          "step_count" => plan.fetch("steps").length,
          "pending_step_count" => plan.fetch("steps").count { |step| step.fetch("status") == "pending" },
          "blocked" => plan.fetch("blocked"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def test_result_summary(recipe)
        result = CompatibilityTestResult.new(recipe: recipe).to_h
        {
          "result_type" => result.fetch("result_type"),
          "test_type" => result.fetch("test_type"),
          "execution_state" => result.fetch("execution_state"),
          "overall_status" => result.fetch("overall_status"),
          "counts" => result.fetch("counts"),
          "summary" => result.fetch("desktop_safe_summary")
        }
      end

      def backend_binding_summary(recipe)
        binding = CompatibilityBackendBinding.new(recipe: recipe).to_h
        {
          "binding_type" => binding.fetch("binding_type"),
          "selected_strategy" => binding.fetch("selected_strategy"),
          "managed_binding_ready" => binding.fetch("managed_binding_ready"),
          "launch_enabled" => binding.fetch("launch_enabled"),
          "preflight_count" => binding.fetch("required_preflight").length,
          "summary" => binding.fetch("desktop_safe_summary")
        }
      end

      def state_root_summary(recipe)
        state_root = ApplicationStateRoot.new(recipe: recipe).to_h
        {
          "root_type" => state_root.fetch("root_type"),
          "state_namespace" => state_root.fetch("state_namespace"),
          "allocation_state" => state_root.fetch("allocation_state"),
          "managed_scope_count" => state_root.fetch("managed_scopes").length,
          "snapshot_eligible" => state_root.fetch("snapshot_eligible"),
          "user_documents_included" => state_root.fetch("user_documents_included"),
          "summary" => state_root.fetch("desktop_safe_summary")
        }
      end

      def package_source_summary(recipe)
        source = CompatibilityPackageSource.new(recipe: recipe).to_h
        {
          "source_type" => source.fetch("source_type"),
          "selected_strategy" => source.fetch("selected_strategy"),
          "source_selection_state" => source.fetch("source_selection_state"),
          "package_source_ready" => source.fetch("package_source_ready"),
          "install_enabled" => source.fetch("install_enabled"),
          "channel_count" => source.fetch("source_channels").length,
          "preflight_count" => source.fetch("required_preflight").length,
          "summary" => source.fetch("desktop_safe_summary")
        }
      end

      def acquisition_preflight_summary(recipe)
        preflight = CompatibilityAcquisitionPreflight.new(recipe: recipe).to_h
        {
          "preflight_type" => preflight.fetch("preflight_type"),
          "selected_strategy" => preflight.fetch("selected_strategy"),
          "preflight_state" => preflight.fetch("preflight_state"),
          "acquisition_ready" => preflight.fetch("acquisition_ready"),
          "download_enabled" => preflight.fetch("download_enabled"),
          "install_enabled" => preflight.fetch("install_enabled"),
          "check_count" => preflight.fetch("checks").length,
          "summary" => preflight.fetch("desktop_safe_summary")
        }
      end

      def artifact_manifest_summary(recipe)
        manifest = CompatibilityArtifactManifest.new(recipe: recipe).to_h
        {
          "manifest_type" => manifest.fetch("manifest_type"),
          "selected_strategy" => manifest.fetch("selected_strategy"),
          "manifest_state" => manifest.fetch("manifest_state"),
          "manifest_ready" => manifest.fetch("manifest_ready"),
          "signature_verified" => manifest.fetch("signature_verified"),
          "download_enabled" => manifest.fetch("download_enabled"),
          "artifact_group_count" => manifest.fetch("artifact_groups").length,
          "summary" => manifest.fetch("desktop_safe_summary")
        }
      end

      def install_plan_summary(recipe)
        plan = CompatibilityInstallPlan.new(recipe: recipe).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "selected_strategy" => plan.fetch("selected_strategy"),
          "install_state" => plan.fetch("install_state"),
          "install_ready" => plan.fetch("install_ready"),
          "desktop_activation_ready" => plan.fetch("desktop_activation_ready"),
          "download_enabled" => plan.fetch("download_enabled"),
          "install_enabled" => plan.fetch("install_enabled"),
          "phase_count" => plan.fetch("phases").length,
          "blocked_phase_count" => plan.fetch("phases").count { |phase| phase.fetch("status") == "blocked" },
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def ai_diagnostic_input_summary(recipe)
        input = AIDiagnosticInput.new(recipe: recipe).to_h
        {
          "input_type" => input.fetch("input_type"),
          "section_count" => input.fetch("context_sections").length,
          "signal_count" => input.fetch("diagnostic_signals").length,
          "ai_provider_called" => input.fetch("ai_provider_called"),
          "network_required" => input.fetch("network_required"),
          "safe_for_ai_diagnostics" => input.fetch("safe_for_ai_diagnostics"),
          "summary" => input.fetch("desktop_safe_summary")
        }
      end

      def ai_diagnostic_recommendation_summary(recipe)
        recommendation = AIDiagnosticRecommendation.new(recipe: recipe).to_h
        {
          "recommendation_type" => recommendation.fetch("recommendation_type"),
          "recommendation_count" => recommendation.fetch("recommendations").length,
          "approval_required_count" => recommendation.fetch("approval_required_actions").length,
          "ai_provider_called" => recommendation.fetch("ai_provider_called"),
          "network_required" => recommendation.fetch("network_required"),
          "safe_for_ai_diagnostics" => recommendation.fetch("safe_for_ai_diagnostics"),
          "summary" => recommendation.fetch("desktop_safe_summary")
        }
      end

      def ai_repair_approval_gate_summary(recipe)
        gate = AIRepairApprovalGate.new(recipe: recipe).to_h
        {
          "gate_type" => gate.fetch("gate_type"),
          "gate_decision" => gate.fetch("gate_decision"),
          "required_gate_count" => gate.fetch("required_gates").length,
          "approval_required_count" => gate.fetch("approval_required_actions").length,
          "auto_execution_allowed" => gate.fetch("auto_execution_allowed"),
          "repair_executed" => gate.fetch("repair_executed"),
          "summary" => gate.fetch("desktop_safe_summary")
        }
      end

      def runtime_service_binding_summary
        binding = RuntimeServiceBinding.new.to_h
        {
          "binding_type" => binding.fetch("binding_type"),
          "activation_binding_ready" => binding.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => binding.fetch("live_dbus_owner_ready"),
          "counts" => binding.fetch("counts"),
          "summary" => binding.fetch("desktop_safe_summary")
        }
      end

      def settings_summary(recipe)
        model = SettingsModel.new(application_id: recipe.id).to_h
        {
          "request_type" => model.fetch("request_type"),
          "settings_state" => model.fetch("settings_state"),
          "settings_persisted" => model.fetch("settings_persisted"),
          "section_count" => model.fetch("section_count"),
          "host_root_modified" => model.fetch("host_root_modified"),
          "backend_details_exposed" => model.fetch("backend_details_exposed"),
          "summary" => model.fetch("desktop_safe_summary")
        }
      end

      def settings_change_plan_summary(recipe)
        plan = SettingsChangePlan.new(
          application_id: recipe.id,
          section_id: "resource-access",
          field_id: "documents",
          value: "ask"
        ).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "change_state" => plan.fetch("change_state"),
          "apply_enabled" => plan.fetch("apply_enabled"),
          "settings_persisted" => plan.fetch("settings_persisted"),
          "portal_policy_review_required" => plan.fetch("portal_policy_review_required"),
          "snapshot_recommended" => plan.fetch("snapshot_recommended"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
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
          when "test-plan"
            application_id = require_argument(command)
            test_type = @argv.shift || "preflight"
            write_json(runtime.test_plan(application_id, test_type))
          when "test-result"
            application_id = require_argument(command)
            test_type = @argv.shift || "preflight"
            write_json(runtime.test_result(application_id, test_type))
          when "backend-binding"
            write_json(runtime.backend_binding(require_argument(command)))
          when "state-root"
            write_json(runtime.state_root(require_argument(command)))
          when "package-source"
            write_json(runtime.package_source(require_argument(command)))
          when "acquisition-preflight"
            write_json(runtime.acquisition_preflight(require_argument(command)))
          when "artifact-manifest"
            write_json(runtime.artifact_manifest(require_argument(command)))
          when "install-plan"
            application_id = require_argument(command)
            environment = @argv.shift || "development"
            write_json(runtime.install_plan(application_id, environment))
          when "ai-diagnostic-input"
            application_id = require_argument(command)
            issue = @argv.shift || AIDiagnosticInput::DEFAULT_ISSUE
            test_type = @argv.shift || "preflight"
            write_json(runtime.ai_diagnostic_input(application_id, issue, test_type))
          when "ai-diagnostic-recommendation"
            application_id = require_argument(command)
            issue = @argv.shift || AIDiagnosticInput::DEFAULT_ISSUE
            test_type = @argv.shift || "preflight"
            write_json(runtime.ai_diagnostic_recommendation(application_id, issue, test_type))
          when "ai-repair-approval-gate"
            application_id = require_argument(command)
            issue = @argv.shift || AIDiagnosticInput::DEFAULT_ISSUE
            test_type = @argv.shift || "preflight"
            write_json(runtime.ai_repair_approval_gate(application_id, issue, test_type))
          when "service-binding"
            write_json(runtime.runtime_service_binding)
          when "settings"
            write_json(runtime.settings(require_argument(command)))
          when "settings-change"
            application_id = require_argument(command)
            section_id = @argv.shift
            field_id = @argv.shift
            value = @argv.shift
            raise ArgumentError, "settings-change requires section, field, and value" unless section_id && field_id && value

            write_json(runtime.settings_change_plan(application_id, section_id, field_id, value))
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
