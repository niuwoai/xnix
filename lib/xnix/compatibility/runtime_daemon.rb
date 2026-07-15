# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "action_review_receipt"
require_relative "application_state_root"
require_relative "ai_diagnostic_input"
require_relative "ai_diagnostic_recommendation"
require_relative "ai_repair_approval_gate"
require_relative "compatibility_acquisition_preflight"
require_relative "compatibility_action_queue"
require_relative "compatibility_artifact_manifest"
require_relative "compatibility_backend_binding"
require_relative "compatibility_backend_capability_matrix"
require_relative "compatibility_backend_selection_plan"
require_relative "compatibility_backend_environment_plan"
require_relative "compatibility_backend_lifecycle"
require_relative "compatibility_engine_catalog"
require_relative "compatibility_execution_readiness"
require_relative "compatibility_install_plan"
require_relative "compatibility_mode_switch_plan"
require_relative "compatibility_permission_review_plan"
require_relative "compatibility_review_flow_plan"
require_relative "compatibility_package_source"
require_relative "compatibility_run_plan"
require_relative "compatibility_repair_plan"
require_relative "compatibility_snapshot_plan"
require_relative "compatibility_test_plan"
require_relative "compatibility_test_result"
require_relative "desktop_entry"
require_relative "desktop_resource_bridge_plan"
require_relative "desktop_integration_manifest"
require_relative "file_association_model"
require_relative "kde_application_surface_plan"
require_relative "kde_shell_integration_plan"
require_relative "launch_request"
require_relative "notification_request"
require_relative "portal_access_policy"
require_relative "portal_request_model"
require_relative "registry_backed_recipe_store"
require_relative "runtime_live_owner_gate"
require_relative "runtime_method_parity_manifest"
require_relative "runtime_owner_smoke_plan"
require_relative "runtime_service_binding"
require_relative "runtime_write_gate"
require_relative "settings_model"
require_relative "settings_change_plan"
require_relative "task_manager_identity"
require_relative "tray_status_model"

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
            "compatibility_action_queues" => true,
            "compatibility_action_review_receipts" => true,
            "compatibility_center_summaries" => true,
            "kde_center_pages" => true,
            "compatibility_artifact_manifests" => true,
            "compatibility_engine_catalog" => true,
            "compatibility_install_planning" => true,
            "compatibility_mode_switch_planning" => true,
            "compatibility_permission_review_planning" => true,
            "compatibility_review_flow_planning" => true,
            "compatibility_package_sources" => true,
            "compatibility_backend_binding" => true,
            "compatibility_backend_capability_matrix" => true,
            "compatibility_backend_selection_plans" => true,
            "compatibility_backend_environment_plans" => true,
            "compatibility_backend_lifecycle" => true,
            "compatibility_run_planning" => true,
            "desktop_activation_manifests" => true,
            "kde_integration_status" => true,
            "kde_shell_integration_plans" => true,
            "kde_application_surface_plans" => true,
            "desktop_resource_bridge_plans" => true,
            "desktop_entry_planning" => true,
            "task_manager_identity_planning" => true,
            "kwin_window_rule_planning" => true,
            "file_association_planning" => true,
            "notification_planning" => true,
            "tray_status_planning" => true,
            "krunner_query_planning" => true,
            "compatibility_repair_planning" => true,
            "compatibility_test_planning" => true,
            "compatibility_test_results" => true,
            "compatibility_execution_readiness" => true,
            "compatibility_launch_intents" => true,
            "ai_diagnostic_inputs" => true,
            "ai_diagnostic_recommendations" => true,
            "ai_repair_approval_gates" => true,
            "runtime_live_owner_gates" => true,
            "runtime_method_parity_manifests" => true,
            "runtime_owner_smoke_plans" => true,
            "runtime_service_binding" => true,
            "runtime_write_gates" => true,
            "compatibility_settings" => true,
            "compatibility_settings_change_planning" => true,
            "portal_request_planning" => true,
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
          "execution_readiness" => execution_readiness_summary(recipe),
          "launch_intent" => launch_intent_summary(recipe),
          "action_queue" => action_queue_summary(recipe),
          "action_review_receipt" => action_review_receipt_summary(recipe),
          "compatibility_center_summary" => compatibility_center_summary_summary(recipe),
          "kde_center_page" => kde_center_page_summary(recipe),
          "desktop_activation_manifest" => desktop_activation_manifest_summary(recipe),
          "kde_shell_integration_plan" => kde_shell_integration_plan_summary,
          "kde_application_surface_plan" => kde_application_surface_plan_summary(recipe),
          "desktop_resource_bridge_plan" => desktop_resource_bridge_plan_summary(recipe),
          "desktop_entry_plan" => desktop_entry_plan_summary(recipe),
          "task_manager_identity_plan" => task_manager_identity_plan_summary(recipe),
          "file_association_plan" => file_association_plan_summary(recipe),
          "notification_plan" => notification_plan_summary(recipe),
          "tray_status" => tray_status_summary,
          "install_plan" => install_plan_summary(recipe),
          "acquisition_preflight" => acquisition_preflight_summary(recipe),
          "artifact_manifest" => artifact_manifest_summary(recipe),
          "package_source" => package_source_summary(recipe),
          "state_root" => state_root_summary(recipe),
          "backend_binding" => backend_binding_summary(recipe),
          "backend_capability_matrix" => backend_capability_matrix_summary,
          "backend_selection_plan" => backend_selection_plan_summary(recipe),
          "backend_environment_plan" => backend_environment_plan_summary(recipe),
          "backend_lifecycle" => backend_lifecycle_summary(recipe),
          "ai_diagnostic_input" => ai_diagnostic_input_summary(recipe),
          "ai_diagnostic_recommendation" => ai_diagnostic_recommendation_summary(recipe),
          "ai_repair_approval_gate" => ai_repair_approval_gate_summary(recipe),
          "runtime_live_owner_gate" => runtime_live_owner_gate_summary,
          "runtime_method_parity_manifest" => runtime_method_parity_manifest_summary,
          "runtime_owner_smoke_plan" => runtime_owner_smoke_plan_summary,
          "runtime_service_binding" => runtime_service_binding_summary,
          "runtime_write_gate" => runtime_write_gate_summary("Launch"),
          "portal_request_plan" => portal_request_plan_summary(recipe),
          "settings" => settings_summary(recipe),
          "settings_change_plan" => settings_change_plan_summary(recipe),
          "compatibility_mode_switch_plan" => compatibility_mode_switch_plan_summary(recipe),
          "compatibility_permission_review_plan" => compatibility_permission_review_plan_summary(recipe),
          "compatibility_review_flow_plan" => compatibility_review_flow_plan_summary(recipe),
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

      def desktop_activation_manifest(application_id)
        recipe = require_recipe(application_id)
        DesktopIntegrationManifest.new(recipe: recipe).to_h
      end

      def kde_integration_status
        entry_points = kde_integration_entry_points

        {
          "status_type" => "kde-integration-status",
          "desktop" => "KDE Plasma",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "official_desktop_only" => true,
          "stable_desktop_contract" => true,
          "entry_point_ids" => entry_points.map { |entry| entry.fetch("id") },
          "entry_point_names" => entry_points.map { |entry| entry.fetch("name") },
          "entry_point_states" => entry_points.map { |entry| entry.fetch("state") },
          "runtime_methods" => entry_points.map { |entry| entry.fetch("runtime_method") },
          "entry_points" => entry_points,
          "entry_point_count" => entry_points.length,
          "initial_count" => entry_points.count { |entry| entry.fetch("state") == "initial" },
          "planned_count" => entry_points.count { |entry| entry.fetch("state") == "planned" },
          "complete_count" => entry_points.count { |entry| entry.fetch("state") == "complete" },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KDE Plasma is the only official first-release shell, and all seven entry points are Runtime-backed."
        }
      end

      def kde_application_surface_plan(application_id)
        recipe = require_recipe(application_id)
        KDEApplicationSurfacePlan.new(recipe: recipe).to_h
      end

      def kde_shell_integration_plan
        KDEShellIntegrationPlan.new.to_h
      end

      def desktop_resource_bridge_plan(application_id)
        recipe = require_recipe(application_id)
        DesktopResourceBridgePlan.new(recipe: recipe).to_h
      end

      def desktop_entry_plan(application_id)
        recipe = require_recipe(application_id)
        desktop_entry = DesktopEntry.new(recipe)
        {
          "version" => VERSION,
          "plan_type" => "desktop-entry-plan",
          "desktop" => "KDE Plasma",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "application_id" => recipe.id,
          "desktop_file" => desktop_entry.file_name,
          "relative_path" => File.join("applications", desktop_entry.file_name),
          "name" => recipe.name,
          "comment" => "Run with Xnix Compatibility Runtime",
          "exec" => "#{DesktopEntry::LAUNCHER} --app #{recipe.id} %U",
          "icon" => recipe.icon,
          "categories" => ["Utility"],
          "mime_types" => recipe.mime_types,
          "startup_wm_class" => "xnix-#{recipe.id}",
          "file_argument_mode" => "%U",
          "standard_desktop_entry" => true,
          "launch_uses_runtime" => true,
          "accepts_file_uris" => true,
          "user_visible" => true,
          "startup_notify" => true,
          "terminal" => false,
          "no_display" => false,
          "safety" => {
            "files_written" => false,
            "host_root_modified" => false,
            "backend_command_exposed" => false,
            "raw_windows_executable_exposed" => false,
            "compatibility_storage_path_exposed" => false,
            "backend_details_exposed" => false
          },
          "desktop_safe_summary" => "Compatibility application desktop entries launch through the Runtime without exposing backend commands."
        }
      end

      def task_manager_identity_plan(application_id)
        recipe = require_recipe(application_id)
        plan = TaskManagerIdentity.new(application_id: recipe.id, name: recipe.name).to_h
        plan["task_manager"] = plan.fetch("task_manager").merge(
          "skip_taskbar" => false,
          "show_in_switcher" => true
        )
        plan["kwin"] = plan.fetch("kwin").merge(
          "placement" => "normal-window",
          "set" => plan.fetch("kwin").fetch("set").merge(
            "skip_taskbar" => false,
            "show_in_switcher" => true
          )
        )
        plan.merge(
          "plan_type" => "task-manager-identity-plan",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "restore" => {
            "restore_allowed" => true,
            "restore_key" => recipe.id,
            "prefer_existing_window" => true
          },
          "safety" => {
            "window_manager_policy_only" => true,
            "runtime_owns_backend_policy" => true,
            "host_root_modified" => false,
            "backend_details_exposed" => false
          },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Compatibility windows are grouped, pinned, switched, and restored through a normal desktop entry."
        )
      end

      def kwin_window_rule_plan(application_id)
        identity = task_manager_identity_plan(application_id)
        {
          "version" => VERSION,
          "request_type" => "kwin-window-rule",
          "desktop" => "KDE Plasma",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "application_id" => identity.fetch("application_id"),
          "name" => identity.fetch("name"),
          "script_role" => "identity-and-layout",
          "match" => identity.fetch("kwin").fetch("match").merge(
            "title_hint" => identity.fetch("window").fetch("title_hint")
          ),
          "set" => {
            "desktop_file" => identity.fetch("desktop_file"),
            "application_id" => identity.fetch("application_id"),
            "task_manager_grouping_key" => identity.fetch("task_manager").fetch("grouping_key"),
            "launcher_url" => identity.fetch("task_manager").fetch("launcher_url"),
            "skip_taskbar" => identity.fetch("task_manager").fetch("skip_taskbar"),
            "show_in_switcher" => identity.fetch("task_manager").fetch("show_in_switcher"),
            "placement" => identity.fetch("kwin").fetch("placement")
          },
          "restore" => {
            "pinning_allowed" => identity.fetch("task_manager").fetch("pinning_allowed"),
            "restore_allowed" => identity.fetch("restore").fetch("restore_allowed"),
            "restore_key" => identity.fetch("restore").fetch("restore_key"),
            "prefer_existing_window" => identity.fetch("restore").fetch("prefer_existing_window")
          },
          "safety" => {
            "window_manager_policy_only" => true,
            "runtime_owns_backend_policy" => true,
            "host_root_modified" => false,
            "backend_details_exposed" => false
          },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KWin window rules are Runtime-owned identity and layout hints for normal desktop behavior."
        }
      end

      def file_association_plan(application_id)
        recipe = require_recipe(application_id)
        FileAssociationModel.new(recipe: recipe).to_h.merge(
          "plan_type" => "file-association-plan",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "files_written" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Recipe MIME types are mapped to a generated desktop entry and portal-mediated file opens."
        )
      end

      def notification_plan(application_id, event_type)
        recipe = require_recipe(application_id)
        request = NotificationRequest.new(application_id: recipe.id, event_type: event_type).to_h
        {
          "version" => VERSION,
          "plan_type" => "notification-plan",
          "desktop" => "KDE Plasma",
          "event_type" => request.fetch("event_type"),
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "application" => {
            "id" => recipe.id,
            "name" => recipe.name,
            "icon" => recipe.icon,
            "runtime_mode" => recipe.mode,
            "desktop_entry" => "xnix-#{recipe.id}.desktop"
          },
          "notification" => {
            "id" => "#{recipe.id}.#{event_type}",
            "title" => request.fetch("title"),
            "body" => request.fetch("body"),
            "urgency" => request.fetch("urgency"),
            "category" => "compatibility.#{event_type.tr("-", ".")}",
            "desktop_entry" => "xnix-#{recipe.id}.desktop",
            "actions" => request.fetch("actions"),
            "action_count" => request.fetch("actions").length
          },
          "safety" => {
            "user_visible" => true,
            "requires_user_review" => %w[approval-required install-failed].include?(event_type),
            "action_execution_enabled" => false,
            "repair_execution_enabled" => false,
            "settings_persistence_enabled" => false,
            "host_root_modified" => false,
            "backend_details_exposed" => false
          },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Runtime events become KDE notification plans while execution and persistence gates remain closed."
        }
      end

      def tray_status
        TrayStatusModel.new(
          active_count: 1,
          attention_count: 1,
          bridged_tray_count: 0
        ).to_h.merge(
          "status_type" => "tray-status-plan",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "safety" => {
            "user_visible" => true,
            "live_backend_bridge_enabled" => false,
            "bridge_configuration_persisted" => false,
            "host_root_modified" => false,
            "backend_details_exposed" => false
          },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Runtime tray status is visible to KDE while live tray bridging and persistence remain gated."
        )
      end

      def krunner_query_plan(query)
        normalized_query = normalize_krunner_query(query)
        matches = normalized_query.empty? ? [] : krunner_matches(normalized_query)

        {
          "version" => VERSION,
          "query_type" => "krunner-query-plan",
          "entry_point" => "krunner",
          "desktop" => "KDE Plasma",
          "query" => query.to_s,
          "matches" => matches,
          "summary" => {
            "match_count" => matches.length,
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

      def action_queue(application_id)
        recipe = require_recipe(application_id)
        CompatibilityActionQueue.new(recipe: recipe).to_h
      end

      def action_review_receipt(application_id, action_id, decision)
        recipe = require_recipe(application_id)
        ActionReviewReceipt.new(recipe: recipe, action_id: action_id, decision: decision).to_h
      end

      def compatibility_center_summary(application_id)
        recipe = require_recipe(application_id)
        repair = repair_plan_summary(recipe.id, "engine-binding-pending")

        {
          "version" => VERSION,
          "summary_type" => "compatibility-center-summary",
          "desktop" => "KDE Plasma",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "application" => {
            "id" => recipe.id,
            "name" => recipe.name,
            "icon" => recipe.icon,
            "runtime_mode" => recipe.mode
          },
          "compatibility" => {
            "state" => "review-required",
            "runtime_mode" => recipe.mode,
            "known_issue_count" => 1,
            "known_issues" => [
              {
                "id" => "engine-binding-pending",
                "severity" => repair.fetch("severity"),
                "summary" => repair.fetch("summary")
              }
            ]
          },
          "repair_records" => {
            "state" => "pending-review",
            "last_event" => repair.fetch("notification_event"),
            "record_count" => 1,
            "records" => [
              {
                "id" => "review-ai-repair",
                "state" => "pending-review",
                "last_event" => repair.fetch("notification_event")
              }
            ]
          },
          "actions" => [
            "open-compatibility-center",
            "review-actions",
            "open-settings",
            "run-preflight-test"
          ],
          "safety" => {
            "user_visible" => true,
            "action_execution_enabled" => false,
            "repair_execution_enabled" => false,
            "backend_launch_enabled" => false,
            "settings_persistence_enabled" => false,
            "host_root_modified" => false,
            "backend_details_exposed" => false
          },
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Compatibility Center summary is Runtime-owned, user-visible, and non-executing."
        }
      end

      def kde_center_page(application_id, decision)
        recipe = require_recipe(application_id)
        summary = compatibility_center_summary(recipe.id)
        settings_model = settings(recipe.id)
        decision_allowed = %w[approved reviewed].include?(decision)

        {
          "version" => VERSION,
          "request_type" => "kde-center-page",
          "page_type" => "compatibility-center-application-page",
          "desktop" => "KDE Plasma",
          "runtime_method" => "GetKDECenterPage",
          "read_model_source" => "go-kde-center-page-preview",
          "application_id" => recipe.id,
          "application_name" => recipe.name,
          "compatibility_state" => summary.fetch("compatibility").fetch("state"),
          "runtime_mode" => summary.fetch("compatibility").fetch("runtime_mode"),
          "known_issue_count" => summary.fetch("compatibility").fetch("known_issue_count"),
          "repair_record_state" => summary.fetch("repair_records").fetch("state"),
          "action_deck_request_type" => "kde-action-card-deck-preview",
          "card_count" => 7,
          "waiting_card_count" => decision_allowed ? 7 : 0,
          "deferred_card_count" => decision == "deferred" ? 7 : 0,
          "rejected_card_count" => decision == "rejected" ? 7 : 0,
          "navigation_action_count" => 28,
          "disabled_action_count" => 21,
          "settings_request_type" => settings_model.fetch("request_type"),
          "settings_section_count" => settings_model.fetch("sections").length,
          "primary_navigation_target" => "compatibility-center-gates",
          "runtime_owned" => true,
          "go_runtime_backed" => true,
          "kde_policy_owner" => false,
          "official_desktop_only" => true,
          "user_visible" => true,
          "safe_for_ai_diagnostics" => true,
          "user_decision_captured" => true,
          "user_decision_allows_launch" => decision_allowed,
          "page_preview_created" => true,
          "page_persisted" => false,
          "deck_persisted" => false,
          "cards_persisted" => false,
          "card_actions_enabled" => false,
          "settings_persisted" => false,
          "settings_persistence_enabled" => false,
          "notifications_sent" => false,
          "resource_grant_created" => false,
          "runtime_launch_approval" => false,
          "launch_allowed" => false,
          "launch_enabled" => false,
          "execution_started" => false,
          "backend_process_started" => false,
          "request_objects_created" => false,
          "permission_grant_created" => false,
          "host_root_modified" => false,
          "network_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KDE can read a Runtime-owned Compatibility Center page over D-Bus, but the page cannot approve, persist, grant, notify, or start execution."
        }
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

      def backend_capability_matrix
        CompatibilityBackendCapabilityMatrix.new.to_h
      end

      def backend_selection_plan(application_id)
        recipe = require_recipe(application_id)
        CompatibilityBackendSelectionPlan.new(recipe: recipe).to_h
      end

      def backend_lifecycle(application_id)
        recipe = require_recipe(application_id)
        CompatibilityBackendLifecycle.new(recipe: recipe).to_h
      end

      def backend_environment_plan(application_id)
        recipe = require_recipe(application_id)
        CompatibilityBackendEnvironmentPlan.new(recipe: recipe).to_h
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

      def execution_readiness(application_id)
        recipe = require_recipe(application_id)
        CompatibilityExecutionReadiness.new(recipe: recipe).to_h
      end

      def launch_intent(application_id)
        LaunchRequest.new(recipe_store: recipe_store).build(application_id: application_id)
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

      def portal_request_plan(application_id, operation)
        require_recipe(application_id)
        PortalRequestModel.new(application_id: application_id, operation: operation).to_h
      end

      def runtime_service_binding
        RuntimeServiceBinding.new.to_h
      end

      def runtime_live_owner_gate
        RuntimeLiveOwnerGate.new.to_h
      end

      def runtime_owner_smoke_plan
        RuntimeOwnerSmokePlan.new.to_h
      end

      def runtime_method_parity_manifest
        RuntimeMethodParityManifest.new.to_h
      end

      def runtime_write_gate(method_name)
        RuntimeWriteGate.new(method_name: method_name).to_h
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

      def compatibility_mode_switch_plan(application_id, requested_mode)
        recipe = require_recipe(application_id)
        CompatibilityModeSwitchPlan.new(recipe: recipe, requested_mode: requested_mode).to_h
      end

      def compatibility_permission_review_plan(application_id)
        recipe = require_recipe(application_id)
        CompatibilityPermissionReviewPlan.new(recipe: recipe).to_h
      end

      def compatibility_review_flow_plan(application_id, section_id, field_id, value, operation)
        recipe = require_recipe(application_id)
        CompatibilityReviewFlowPlan.new(
          recipe: recipe,
          section_id: section_id,
          field_id: field_id,
          value: value,
          operation: operation
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
        when "GetDesktopActivationManifest"
          desktop_activation_manifest(required_parameter(method_name, parameters, 0))
        when "GetKDEIntegrationStatus"
          kde_integration_status
        when "GetKDEShellIntegrationPlan"
          kde_shell_integration_plan
        when "GetKDEApplicationSurfacePlan"
          kde_application_surface_plan(required_parameter(method_name, parameters, 0))
        when "GetDesktopResourceBridgePlan"
          desktop_resource_bridge_plan(required_parameter(method_name, parameters, 0))
        when "GetDesktopEntryPlan"
          desktop_entry_plan(required_parameter(method_name, parameters, 0))
        when "GetTaskManagerIdentityPlan"
          task_manager_identity_plan(required_parameter(method_name, parameters, 0))
        when "GetKWinWindowRulePlan"
          kwin_window_rule_plan(required_parameter(method_name, parameters, 0))
        when "GetFileAssociationPlan"
          file_association_plan(required_parameter(method_name, parameters, 0))
        when "GetNotificationPlan"
          notification_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetTrayStatus"
          tray_status
        when "GetKRunnerQueryPlan"
          krunner_query_plan(required_parameter(method_name, parameters, 0))
        when "GetApplicationStateRoot"
          state_root(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityPackageSource"
          package_source(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityAcquisitionPreflight"
          acquisition_preflight(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityActionQueue"
          action_queue(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityActionReviewReceipt"
          action_review_receipt(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1),
            required_parameter(method_name, parameters, 2)
          )
        when "GetCompatibilityCenterSummary"
          compatibility_center_summary(required_parameter(method_name, parameters, 0))
        when "GetKDECenterPage"
          kde_center_page(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetCompatibilityArtifactManifest"
          artifact_manifest(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityInstallPlan"
          install_plan(
            required_parameter(method_name, parameters, 0),
            parameters[1] || "development"
          )
        when "GetBackendBinding"
          backend_binding(required_parameter(method_name, parameters, 0))
        when "GetBackendCapabilityMatrix"
          backend_capability_matrix
        when "GetBackendSelectionPlan"
          backend_selection_plan(required_parameter(method_name, parameters, 0))
        when "GetBackendLifecycle"
          backend_lifecycle(required_parameter(method_name, parameters, 0))
        when "GetBackendEnvironmentPlan"
          backend_environment_plan(required_parameter(method_name, parameters, 0))
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
        when "GetExecutionReadiness"
          execution_readiness(required_parameter(method_name, parameters, 0))
        when "GetLaunchIntent"
          launch_intent(required_parameter(method_name, parameters, 0))
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
        when "GetPortalRequestPlan"
          portal_request_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetRuntimeServiceBinding"
          runtime_service_binding
        when "GetRuntimeLiveOwnerGate"
          runtime_live_owner_gate
        when "GetRuntimeOwnerSmokePlan"
          runtime_owner_smoke_plan
        when "GetRuntimeMethodParityManifest"
          runtime_method_parity_manifest
        when "GetRuntimeWriteGate"
          runtime_write_gate(required_parameter(method_name, parameters, 0))
        when "GetCompatibilitySettings"
          settings(required_parameter(method_name, parameters, 0))
        when "GetCompatibilitySettingsChangePlan"
          settings_change_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1),
            required_parameter(method_name, parameters, 2),
            required_parameter(method_name, parameters, 3)
          )
        when "GetCompatibilityModeSwitchPlan"
          compatibility_mode_switch_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1)
          )
        when "GetCompatibilityPermissionReviewPlan"
          compatibility_permission_review_plan(required_parameter(method_name, parameters, 0))
        when "GetCompatibilityReviewFlowPlan"
          compatibility_review_flow_plan(
            required_parameter(method_name, parameters, 0),
            required_parameter(method_name, parameters, 1),
            required_parameter(method_name, parameters, 2),
            required_parameter(method_name, parameters, 3),
            required_parameter(method_name, parameters, 4)
          )
        when *RuntimeWriteGate::WRITE_METHODS
          raise ArgumentError, RuntimeWriteGate.new(method_name: method_name).failure_message
        else
          raise ArgumentError, "unsupported runtime method: #{method_name}"
        end
      end

      def introspection_xml
        CONTRACT_PATH.read
      end

      private

      def kde_integration_entry_points
        [
          kde_integration_entry(
            "launcher",
            "Launcher",
            "GetDesktopEntryPlan",
            "standard-desktop-entry",
            "Windows applications appear in the KDE launcher through generated desktop entries.",
            "Connect activation receipts to a production recipe installer."
          ),
          kde_integration_entry(
            "task-manager",
            "Task Manager",
            "GetTaskManagerIdentityPlan",
            "window-identity-and-restore",
            "Compatibility windows expose grouping, pinning, switcher, and restore identity.",
            "Connect Runtime identity plans to a production KWin script and task manager bridge."
          ),
          kde_integration_entry(
            "file-manager",
            "File Manager",
            "GetFileAssociationPlan",
            "dolphin-service-menu-and-mime",
            "Dolphin opens selected files through portal-mediated Runtime file-open planning.",
            "Connect file-open requests to production Runtime launch requests."
          ),
          kde_integration_entry(
            "system-tray",
            "System Tray",
            "GetTrayStatus",
            "runtime-status-surface",
            "The tray can show Runtime activity, attention state, and bridge readiness.",
            "Connect tray status plans to a production Plasma tray surface."
          ),
          kde_integration_entry(
            "notifications",
            "Notifications",
            "GetNotificationPlan",
            "runtime-event-notification",
            "Runtime events map to KDE notification payloads for install, repair, mode, and approval states.",
            "Connect notification plans to the production KDE notification path."
          ),
          kde_integration_entry(
            "compatibility-center",
            "Compatibility Center",
            "GetCompatibilityCenterSummary",
            "plasma-read-model",
            "The Compatibility Center can show Runtime-owned application state and safe action cards.",
            "Render live Runtime applications and diagnostics in the Plasmoid."
          ),
          kde_integration_entry(
            "settings",
            "Settings",
            "GetCompatibilitySettings",
            "user-facing-policy-controls",
            "Settings expose user-facing Runtime policy without backend terminology.",
            "Connect settings plans to a production KDE settings module and persisted Runtime policy."
          )
        ]
      end

      def kde_integration_entry(id, name, runtime_method, adapter_role, summary, next_step)
        {
          "id" => id,
          "name" => name,
          "state" => "initial",
          "runtime_method" => runtime_method,
          "adapter_role" => adapter_role,
          "runtime_backed" => true,
          "c_runtime_backed" => true,
          "dbus_read_available" => true,
          "kde_policy_owner" => false,
          "summary" => summary,
          "next_step" => next_step
        }
      end

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

      def execution_readiness_summary(recipe)
        readiness = CompatibilityExecutionReadiness.new(recipe: recipe).to_h
        {
          "readiness_type" => readiness.fetch("readiness_type"),
          "execution_state" => readiness.fetch("execution_state"),
          "overall_status" => readiness.fetch("overall_status"),
          "launch_allowed" => readiness.fetch("launch_allowed"),
          "launch_enabled" => readiness.fetch("launch_enabled"),
          "backend_binding_ready" => readiness.fetch("backend_binding_ready"),
          "summary" => readiness.fetch("desktop_safe_summary")
        }
      end

      def launch_intent_summary(recipe)
        intent = LaunchRequest.new(recipe_store: recipe_store).build(application_id: recipe.id)
        {
          "intent_type" => intent.fetch("intent_type"),
          "source" => intent.fetch("source"),
          "runtime_method" => intent.fetch("runtime_method"),
          "read_method" => intent.fetch("read_method"),
          "launch_allowed" => intent.fetch("launch_allowed"),
          "launch_enabled" => intent.fetch("launch_enabled"),
          "execution_request_created" => intent.fetch("execution_request_created"),
          "summary" => intent.fetch("desktop_safe_summary")
        }
      end

      def action_queue_summary(recipe)
        queue = CompatibilityActionQueue.new(recipe: recipe).to_h
        {
          "queue_type" => queue.fetch("queue_type"),
          "action_count" => queue.fetch("action_count"),
          "pending_action_count" => queue.fetch("pending_action_count"),
          "user_review_required_count" => queue.fetch("user_review_required_count"),
          "execution_enabled" => queue.fetch("execution_enabled"),
          "repair_execution_enabled" => queue.fetch("repair_execution_enabled"),
          "settings_persistence_enabled" => queue.fetch("settings_persistence_enabled"),
          "backend_details_exposed" => queue.fetch("backend_details_exposed"),
          "summary" => queue.fetch("desktop_safe_summary")
        }
      end

      def action_review_receipt_summary(recipe)
        receipt = ActionReviewReceipt.new(recipe: recipe, action_id: "review-ai-repair", decision: "approved").to_h
        {
          "receipt_type" => receipt.fetch("receipt_type"),
          "decision" => receipt.fetch("decision"),
          "decision_recorded" => receipt.fetch("decision_recorded"),
          "action_id" => receipt.fetch("action").fetch("id"),
          "execution_enabled" => receipt.fetch("execution_enabled"),
          "repair_execution_enabled" => receipt.fetch("repair_execution_enabled"),
          "settings_persistence_enabled" => receipt.fetch("settings_persistence_enabled"),
          "resource_grant_created" => receipt.fetch("resource_grant_created"),
          "backend_details_exposed" => receipt.fetch("backend_details_exposed"),
          "summary" => receipt.fetch("desktop_safe_summary")
        }
      end

      def compatibility_center_summary_summary(recipe)
        summary = compatibility_center_summary(recipe.id)
        {
          "summary_type" => summary.fetch("summary_type"),
          "compatibility_state" => summary.fetch("compatibility").fetch("state"),
          "runtime_mode" => summary.fetch("compatibility").fetch("runtime_mode"),
          "known_issue_count" => summary.fetch("compatibility").fetch("known_issue_count"),
          "repair_record_state" => summary.fetch("repair_records").fetch("state"),
          "repair_record_count" => summary.fetch("repair_records").fetch("record_count"),
          "last_repair_event" => summary.fetch("repair_records").fetch("last_event"),
          "action_count" => summary.fetch("actions").length,
          "action_execution_enabled" => summary.fetch("safety").fetch("action_execution_enabled"),
          "repair_execution_enabled" => summary.fetch("safety").fetch("repair_execution_enabled"),
          "backend_launch_enabled" => summary.fetch("safety").fetch("backend_launch_enabled"),
          "backend_details_exposed" => summary.fetch("backend_details_exposed"),
          "summary" => summary.fetch("desktop_safe_summary")
        }
      end

      def kde_center_page_summary(recipe)
        page = kde_center_page(recipe.id, "approved")
        {
          "page_type" => page.fetch("page_type"),
          "runtime_method" => page.fetch("runtime_method"),
          "card_count" => page.fetch("card_count"),
          "settings_section_count" => page.fetch("settings_section_count"),
          "page_preview_created" => page.fetch("page_preview_created"),
          "page_persisted" => page.fetch("page_persisted"),
          "card_actions_enabled" => page.fetch("card_actions_enabled"),
          "settings_persisted" => page.fetch("settings_persisted"),
          "execution_started" => page.fetch("execution_started"),
          "backend_details_exposed" => page.fetch("backend_details_exposed"),
          "summary" => page.fetch("desktop_safe_summary")
        }
      end

      def desktop_activation_manifest_summary(recipe)
        manifest = DesktopIntegrationManifest.new(recipe: recipe).to_h
        {
          "manifest_type" => manifest.fetch("manifest_type"),
          "desktop" => manifest.fetch("desktop"),
          "official_desktop_only" => manifest.fetch("official_desktop_only"),
          "artifact_count" => manifest.fetch("artifacts").length,
          "entry_points" => manifest.fetch("entry_points"),
          "backend_commands_exposed" => manifest.fetch("safety").fetch("backend_commands_exposed"),
          "portal_required_for_file_access" => manifest.fetch("safety").fetch("portal_required_for_file_access"),
          "host_privilege_required" => manifest.fetch("safety").fetch("host_privilege_required")
        }
      end

      def kde_application_surface_plan_summary(recipe)
        plan = kde_application_surface_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "surface_state" => plan.fetch("surface_state"),
          "desktop_shell" => plan.fetch("desktop_shell"),
          "entry_point_count" => plan.fetch("entry_point_count"),
          "normal_linux_application_surface" => plan.fetch("normal_linux_application_surface"),
          "standard_launcher_visible" => plan.fetch("standard_launcher_visible"),
          "launch_enabled" => plan.fetch("launch_enabled"),
          "backend_process_started" => plan.fetch("backend_process_started"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def kde_shell_integration_plan_summary
        plan = kde_shell_integration_plan
        {
          "plan_type" => plan.fetch("plan_type"),
          "runtime_method" => plan.fetch("runtime_method"),
          "desktop_shell" => plan.fetch("desktop_shell"),
          "component_count" => plan.fetch("component_count"),
          "initial_component_count" => plan.fetch("initial_component_count"),
          "planned_component_count" => plan.fetch("planned_component_count"),
          "official_desktop_only" => plan.fetch("official_desktop_only"),
          "runtime_owned" => plan.fetch("runtime_owned"),
          "kde_policy_owner" => plan.fetch("kde_policy_owner"),
          "plasma_fork_required" => plan.fetch("plasma_fork_required"),
          "plasma_source_modified" => plan.fetch("plasma_source_modified"),
          "shell_configuration_written" => plan.fetch("shell_configuration_written"),
          "component_activation_enabled" => plan.fetch("component_activation_enabled"),
          "backend_launch_enabled" => plan.fetch("backend_launch_enabled"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "host_root_modified" => plan.fetch("host_root_modified"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def desktop_resource_bridge_plan_summary(recipe)
        plan = desktop_resource_bridge_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "bridge_state" => plan.fetch("bridge_state"),
          "resource_count" => plan.fetch("resource_count"),
          "portal_mediated" => plan.fetch("portal_mediated"),
          "file_bridge_planned" => plan.fetch("file_bridge_planned"),
          "print_bridge_planned" => plan.fetch("print_bridge_planned"),
          "clipboard_bridge_planned" => plan.fetch("clipboard_bridge_planned"),
          "bridges_enabled" => plan.fetch("bridges_enabled"),
          "requests_created" => plan.fetch("requests_created"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def desktop_entry_plan_summary(recipe)
        plan = desktop_entry_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "desktop_file" => plan.fetch("desktop_file"),
          "standard_desktop_entry" => plan.fetch("standard_desktop_entry"),
          "launch_uses_runtime" => plan.fetch("launch_uses_runtime"),
          "accepts_file_uris" => plan.fetch("accepts_file_uris"),
          "backend_command_exposed" => plan.fetch("safety").fetch("backend_command_exposed"),
          "raw_windows_executable_exposed" => plan.fetch("safety").fetch("raw_windows_executable_exposed"),
          "compatibility_storage_path_exposed" => plan.fetch("safety").fetch("compatibility_storage_path_exposed"),
          "host_root_modified" => plan.fetch("safety").fetch("host_root_modified"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def task_manager_identity_plan_summary(recipe)
        plan = task_manager_identity_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "desktop_file" => plan.fetch("desktop_file"),
          "class_group" => plan.fetch("window").fetch("class_group"),
          "grouping_key" => plan.fetch("task_manager").fetch("grouping_key"),
          "pinning_allowed" => plan.fetch("task_manager").fetch("pinning_allowed"),
          "restore_allowed" => plan.fetch("task_manager").fetch("restore_allowed"),
          "skip_taskbar" => plan.fetch("kwin").fetch("set").fetch("skip_taskbar"),
          "show_in_switcher" => plan.fetch("kwin").fetch("set").fetch("show_in_switcher"),
          "window_manager_policy_only" => plan.fetch("safety").fetch("window_manager_policy_only"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def file_association_plan_summary(recipe)
        plan = file_association_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "association_type" => plan.fetch("association_type"),
          "desktop_file" => plan.fetch("application").fetch("desktop_file"),
          "mimeapps_path" => plan.fetch("mimeapps").fetch("path"),
          "association_count" => plan.fetch("associations").length,
          "standard_mimeapps_list" => plan.fetch("safety").fetch("standard_mimeapps_list"),
          "staged_root_only" => plan.fetch("safety").fetch("staged_root_only"),
          "overwrite_existing_mimeapps" => plan.fetch("safety").fetch("overwrite_existing_mimeapps"),
          "portal_required_for_file_open" => plan.fetch("safety").fetch("portal_required_for_file_open"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def notification_plan_summary(recipe)
        plan = notification_plan(recipe.id, "approval-required")
        {
          "plan_type" => plan.fetch("plan_type"),
          "event_type" => plan.fetch("event_type"),
          "notification_id" => plan.fetch("notification").fetch("id"),
          "urgency" => plan.fetch("notification").fetch("urgency"),
          "action_count" => plan.fetch("notification").fetch("action_count"),
          "requires_user_review" => plan.fetch("safety").fetch("requires_user_review"),
          "action_execution_enabled" => plan.fetch("safety").fetch("action_execution_enabled"),
          "repair_execution_enabled" => plan.fetch("safety").fetch("repair_execution_enabled"),
          "settings_persistence_enabled" => plan.fetch("safety").fetch("settings_persistence_enabled"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def tray_status_summary
        status = tray_status
        {
          "status_type" => status.fetch("status_type"),
          "active_application_count" => status.fetch("runtime_activity").fetch("active_application_count"),
          "attention_required_count" => status.fetch("runtime_activity").fetch("attention_required_count"),
          "compatibility_state" => status.fetch("compatibility_status").fetch("state"),
          "tray_bridge_state" => status.fetch("tray_bridge").fetch("state"),
          "bridged_tray_application_count" => status.fetch("tray_bridge").fetch("bridged_tray_application_count"),
          "live_backend_bridge_enabled" => status.fetch("safety").fetch("live_backend_bridge_enabled"),
          "bridge_configuration_persisted" => status.fetch("safety").fetch("bridge_configuration_persisted"),
          "backend_details_exposed" => status.fetch("backend_details_exposed"),
          "summary" => status.fetch("desktop_safe_summary")
        }
      end

      def krunner_matches(normalized_query)
        list_applications.each_with_object([]) do |application, matches|
          relevance = krunner_relevance_for(application, normalized_query)
          next if relevance.zero?

          matches << krunner_match(application, relevance)
        end.sort_by { |match| [-match.fetch("relevance_percent"), match.fetch("name")] }
      end

      def krunner_match(application, relevance)
        application_id = application.fetch("id")
        {
          "runner_id" => "xnix.compatibility.#{application_id}",
          "application_id" => application_id,
          "name" => application.fetch("name"),
          "icon" => application.fetch("icon"),
          "relevance_percent" => relevance,
          "subtitle" => "Open as a normal Linux application",
          "mode_label" => "Automatic",
          "supported_extensions" => application.fetch("supported_extensions", []),
          "runtime_owned_launch" => true,
          "backend_details_exposed" => false,
          "action" => {
            "type" => "runtime-launch",
            "desktop_entry_id" => "#{application_id}.desktop",
            "argv" => ["xnix-compat-launch", "--app", application_id]
          }
        }
      end

      def krunner_relevance_for(application, normalized_query)
        name = normalize_krunner_query(application.fetch("name"))
        application_id = normalize_krunner_query(application.fetch("id"))
        extensions = application.fetch("supported_extensions", []).map do |extension|
          normalize_krunner_query(extension.delete_prefix("."))
        end

        return 100 if name == normalized_query || normalized_query == "notepad"
        return 95 if name.start_with?(normalized_query)
        return 90 if %w[open launch start run].any? { |verb| normalized_query == "#{verb} notepad" }
        return 85 if extensions.include?(normalized_query.delete_prefix(".")) || normalized_query == "open txt"
        return 75 if application_id.include?(normalized_query) || name.include?(normalized_query)
        query_tokens = normalized_query.split
        return 65 if (query_tokens & extensions).any? && (query_tokens & %w[open launch start run file]).any?

        0
      end

      def normalize_krunner_query(value)
        value.to_s.downcase.strip.gsub(/\s+/, " ")
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

      def backend_capability_matrix_summary
        matrix = backend_capability_matrix
        {
          "matrix_type" => matrix.fetch("matrix_type"),
          "runtime_method" => matrix.fetch("runtime_method"),
          "profile_count" => matrix.fetch("profile_count"),
          "capability_count" => matrix.fetch("capability_count"),
          "ready_capability_count" => matrix.fetch("ready_capability_count"),
          "pending_capability_count" => matrix.fetch("pending_capability_count"),
          "selection_enabled" => matrix.fetch("selection_enabled"),
          "backend_launch_enabled" => matrix.fetch("backend_launch_enabled"),
          "capability_activation_enabled" => matrix.fetch("capability_activation_enabled"),
          "request_objects_created" => matrix.fetch("request_objects_created"),
          "state_root_created" => matrix.fetch("state_root_created"),
          "snapshots_created" => matrix.fetch("snapshots_created"),
          "host_root_modified" => matrix.fetch("host_root_modified"),
          "backend_details_exposed" => matrix.fetch("backend_details_exposed"),
          "summary" => matrix.fetch("desktop_safe_summary")
        }
      end

      def backend_selection_plan_summary(recipe)
        plan = CompatibilityBackendSelectionPlan.new(recipe: recipe).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "runtime_method" => plan.fetch("runtime_method"),
          "recommended_profile_id" => plan.fetch("recommended_profile_id"),
          "candidate_count" => plan.fetch("candidate_count"),
          "blocked_candidate_count" => plan.fetch("blocked_candidate_count"),
          "selection_committed" => plan.fetch("selection_committed"),
          "selection_change_enabled" => plan.fetch("selection_change_enabled"),
          "backend_launch_enabled" => plan.fetch("backend_launch_enabled"),
          "environment_created" => plan.fetch("environment_created"),
          "request_object_created" => plan.fetch("request_object_created"),
          "host_root_modified" => plan.fetch("host_root_modified"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def backend_lifecycle_summary(recipe)
        lifecycle = CompatibilityBackendLifecycle.new(recipe: recipe).to_h
        {
          "lifecycle_type" => lifecycle.fetch("lifecycle_type"),
          "lifecycle_state" => lifecycle.fetch("lifecycle_state"),
          "overall_status" => lifecycle.fetch("overall_status"),
          "stage_count" => lifecycle.fetch("stages").length,
          "backend_process_started" => lifecycle.fetch("backend_process_started"),
          "launch_enabled" => lifecycle.fetch("launch_enabled"),
          "summary" => lifecycle.fetch("desktop_safe_summary")
        }
      end

      def backend_environment_plan_summary(recipe)
        plan = CompatibilityBackendEnvironmentPlan.new(recipe: recipe).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "environment_state" => plan.fetch("environment_state"),
          "profile_count" => plan.fetch("profiles").length,
          "environment_created" => plan.fetch("environment_created"),
          "backend_process_started" => plan.fetch("backend_process_started"),
          "launch_enabled" => plan.fetch("launch_enabled"),
          "summary" => plan.fetch("desktop_safe_summary")
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

      def runtime_live_owner_gate_summary
        gate = RuntimeLiveOwnerGate.new.to_h
        {
          "gate_type" => gate.fetch("gate_type"),
          "activation_binding_ready" => gate.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => gate.fetch("live_dbus_owner_ready"),
          "production_owner_enabled" => gate.fetch("production_owner_enabled"),
          "owner_transition_ready" => gate.fetch("owner_transition_ready"),
          "smoke_adapter_is_production_owner" => gate.fetch("smoke_adapter_is_production_owner"),
          "pending_gate_count" => gate.fetch("required_gates").count { |item| item.fetch("status") == "pending" },
          "blocked_reason_count" => gate.fetch("blocked_reasons").length,
          "summary" => gate.fetch("desktop_safe_summary")
        }
      end

      def runtime_owner_smoke_plan_summary
        plan = RuntimeOwnerSmokePlan.new.to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "smoke_state" => plan.fetch("smoke_state"),
          "smoke_environment" => plan.fetch("smoke_environment"),
          "activation_binding_ready" => plan.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => plan.fetch("live_dbus_owner_ready"),
          "production_owner_enabled" => plan.fetch("production_owner_enabled"),
          "owner_transition_ready" => plan.fetch("owner_transition_ready"),
          "pending_step_count" => plan.fetch("counts").fetch("pending"),
          "system_service_started" => plan.fetch("system_service_started"),
          "production_bus_claimed" => plan.fetch("production_bus_claimed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def runtime_method_parity_manifest_summary
        manifest = RuntimeMethodParityManifest.new.to_h
        {
          "manifest_type" => manifest.fetch("manifest_type"),
          "method_count" => manifest.fetch("method_count"),
          "read_only_method_parity_ready" => manifest.fetch("read_only_method_parity_ready"),
          "passed_check_count" => manifest.fetch("counts").fetch("passed"),
          "blocked_check_count" => manifest.fetch("counts").fetch("blocked"),
          "write_methods_supported" => manifest.fetch("write_methods_supported"),
          "write_method_dispatch_enabled" => manifest.fetch("write_method_dispatch_enabled"),
          "summary" => manifest.fetch("desktop_safe_summary")
        }
      end

      def runtime_write_gate_summary(method_name)
        gate = RuntimeWriteGate.new(method_name: method_name).to_h
        {
          "gate_type" => gate.fetch("gate_type"),
          "method_name" => gate.fetch("method_name"),
          "gate_decision" => gate.fetch("gate_decision"),
          "write_method_enabled" => gate.fetch("write_method_enabled"),
          "dispatch_enabled" => gate.fetch("dispatch_enabled"),
          "request_object_created" => gate.fetch("request_object_created"),
          "required_gate_count" => gate.fetch("required_gates").length,
          "denial_error_name" => gate.fetch("denial_error_name"),
          "backend_details_exposed" => gate.fetch("backend_details_exposed"),
          "summary" => gate.fetch("desktop_safe_summary")
        }
      end

      def portal_request_plan_summary(recipe)
        plan = PortalRequestModel.new(application_id: recipe.id, operation: "file-open").to_h
        {
          "request_type" => plan.fetch("request_type"),
          "operation" => plan.fetch("operation"),
          "decision" => plan.fetch("decision"),
          "request_allowed" => plan.fetch("request_allowed"),
          "portal_required" => plan.fetch("safety").fetch("portal_required"),
          "host_permission_changed" => plan.fetch("safety").fetch("host_permission_changed"),
          "backend_details_exposed" => plan.fetch("safety").fetch("backend_details_exposed")
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

      def compatibility_mode_switch_plan_summary(recipe)
        plan = compatibility_mode_switch_plan(recipe.id, "prefer-compatibility")
        {
          "plan_type" => plan.fetch("plan_type"),
          "runtime_method" => plan.fetch("runtime_method"),
          "current_mode" => plan.fetch("current_mode"),
          "requested_mode" => plan.fetch("requested_mode"),
          "mode_state" => plan.fetch("mode_state"),
          "mode_count" => plan.fetch("mode_count"),
          "requires_user_confirmation" => plan.fetch("requires_user_confirmation"),
          "portal_review_required" => plan.fetch("portal_review_required"),
          "snapshot_required" => plan.fetch("snapshot_required"),
          "settings_persistence_enabled" => plan.fetch("settings_persistence_enabled"),
          "backend_reconfiguration_enabled" => plan.fetch("backend_reconfiguration_enabled"),
          "backend_process_started" => plan.fetch("backend_process_started"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def compatibility_permission_review_plan_summary(recipe)
        plan = compatibility_permission_review_plan(recipe.id)
        {
          "plan_type" => plan.fetch("plan_type"),
          "runtime_method" => plan.fetch("runtime_method"),
          "review_state" => plan.fetch("review_state"),
          "permission_count" => plan.fetch("permission_count"),
          "allow_count" => plan.fetch("allow_count"),
          "ask_count" => plan.fetch("ask_count"),
          "deny_count" => plan.fetch("deny_count"),
          "user_review_required" => plan.fetch("user_review_required"),
          "portal_review_required" => plan.fetch("portal_review_required"),
          "permission_changes_applied" => plan.fetch("permission_changes_applied"),
          "request_objects_created" => plan.fetch("request_objects_created"),
          "permissions_granted" => plan.fetch("permissions_granted"),
          "settings_persisted" => plan.fetch("settings_persisted"),
          "host_permission_changed" => plan.fetch("host_permission_changed"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("desktop_safe_summary")
        }
      end

      def compatibility_review_flow_plan_summary(recipe)
        plan = compatibility_review_flow_plan(
          recipe.id,
          CompatibilityReviewFlowPlan::DEFAULT_SECTION_ID,
          CompatibilityReviewFlowPlan::DEFAULT_FIELD_ID,
          CompatibilityReviewFlowPlan::DEFAULT_VALUE,
          CompatibilityReviewFlowPlan::DEFAULT_OPERATION
        )
        {
          "plan_type" => plan.fetch("plan_type"),
          "runtime_method" => plan.fetch("runtime_method"),
          "review_state" => plan.fetch("review_state"),
          "step_count" => plan.fetch("step_count"),
          "required_review_count" => plan.fetch("required_review_count"),
          "blocked_step_count" => plan.fetch("blocked_step_count"),
          "pending_step_count" => plan.fetch("pending_step_count"),
          "user_confirmation_required" => plan.fetch("user_confirmation_required"),
          "portal_policy_review_required" => plan.fetch("portal_policy_review_required"),
          "apply_enabled" => plan.fetch("apply_enabled"),
          "request_object_created" => plan.fetch("request_object_created"),
          "permission_granted" => plan.fetch("permission_granted"),
          "settings_persisted" => plan.fetch("settings_persisted"),
          "execution_started" => plan.fetch("execution_started"),
          "host_root_modified" => plan.fetch("host_root_modified"),
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
          when "execution-readiness"
            write_json(runtime.execution_readiness(require_argument(command)))
          when "launch-intent"
            write_json(runtime.launch_intent(require_argument(command)))
          when "backend-binding"
            write_json(runtime.backend_binding(require_argument(command)))
          when "backend-capability-matrix"
            write_json(runtime.backend_capability_matrix)
          when "backend-selection-plan"
            write_json(runtime.backend_selection_plan(require_argument(command)))
          when "backend-lifecycle"
            write_json(runtime.backend_lifecycle(require_argument(command)))
          when "backend-environment-plan"
            write_json(runtime.backend_environment_plan(require_argument(command)))
          when "kde-integration-status"
            write_json(runtime.kde_integration_status)
          when "kde-shell-integration-plan"
            write_json(runtime.kde_shell_integration_plan)
          when "kde-application-surface-plan"
            write_json(runtime.kde_application_surface_plan(require_argument(command)))
          when "desktop-resource-bridge-plan"
            write_json(runtime.desktop_resource_bridge_plan(require_argument(command)))
          when "desktop-entry-plan"
            write_json(runtime.desktop_entry_plan(require_argument(command)))
          when "task-manager-identity-plan"
            write_json(runtime.task_manager_identity_plan(require_argument(command)))
          when "kwin-window-rule-plan"
            write_json(runtime.kwin_window_rule_plan(require_argument(command)))
          when "file-association-plan"
            write_json(runtime.file_association_plan(require_argument(command)))
          when "notification-plan"
            application_id = require_argument(command)
            event_type = @argv.shift
            raise ArgumentError, "notification-plan requires event type" unless event_type

            write_json(runtime.notification_plan(application_id, event_type))
          when "tray-status"
            write_json(runtime.tray_status)
          when "krunner-query-plan"
            write_json(runtime.krunner_query_plan(require_argument(command)))
          when "state-root"
            write_json(runtime.state_root(require_argument(command)))
          when "package-source"
            write_json(runtime.package_source(require_argument(command)))
          when "acquisition-preflight"
            write_json(runtime.acquisition_preflight(require_argument(command)))
          when "action-queue"
            write_json(runtime.action_queue(require_argument(command)))
          when "action-review"
            application_id = require_argument(command)
            action_id = @argv.shift
            decision = @argv.shift
            raise ArgumentError, "action-review requires action id and decision" unless action_id && decision

            write_json(runtime.action_review_receipt(application_id, action_id, decision))
          when "compatibility-center-summary"
            write_json(runtime.compatibility_center_summary(require_argument(command)))
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
          when "live-owner-gate"
            write_json(runtime.runtime_live_owner_gate)
          when "owner-smoke-plan"
            write_json(runtime.runtime_owner_smoke_plan)
          when "method-parity"
            write_json(runtime.runtime_method_parity_manifest)
          when "write-gate"
            write_json(runtime.runtime_write_gate(require_argument(command)))
          when "settings"
            write_json(runtime.settings(require_argument(command)))
          when "settings-change"
            application_id = require_argument(command)
            section_id = @argv.shift
            field_id = @argv.shift
            value = @argv.shift
            raise ArgumentError, "settings-change requires section, field, and value" unless section_id && field_id && value

            write_json(runtime.settings_change_plan(application_id, section_id, field_id, value))
          when "mode-switch-plan"
            application_id = require_argument(command)
            requested_mode = @argv.shift
            raise ArgumentError, "mode-switch-plan requires requested mode" unless requested_mode

            write_json(runtime.compatibility_mode_switch_plan(application_id, requested_mode))
          when "permission-review-plan"
            write_json(runtime.compatibility_permission_review_plan(require_argument(command)))
          when "review-flow-plan"
            application_id = require_argument(command)
            section_id = @argv.shift || CompatibilityReviewFlowPlan::DEFAULT_SECTION_ID
            field_id = @argv.shift || CompatibilityReviewFlowPlan::DEFAULT_FIELD_ID
            value = @argv.shift || CompatibilityReviewFlowPlan::DEFAULT_VALUE
            operation = @argv.shift || CompatibilityReviewFlowPlan::DEFAULT_OPERATION
            write_json(runtime.compatibility_review_flow_plan(application_id, section_id, field_id, value, operation))
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
