#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
DEFAULT_GO_BIN = Pathname.new("/opt/homebrew/bin/go").executable? ? "/opt/homebrew/bin/go" : "go"
DEFAULT_REGISTRY = "runtime/recipes/registry.json"
DEFAULT_APP = "org.xnix.sample.notepad"
DEFAULT_DECISION = "approved"
DEFAULT_MODE = "development"
DEFAULT_FILE_URI = "file:///home/xnix/Documents/sample.txt"
EXPECTED_ENTRYPOINTS = %w[
  launcher
  task-manager
  file-manager
  system-tray
  notifications
  compatibility-center
  settings
].freeze
EXPECTED_STAGED_FILES = %w[
  desktop-entry
  dolphin-service-menu
  mimeapps-list
  desktop-integration-manifest
  desktop-activation-receipt
].freeze
SAFETY_FALSE_KEYS = %w[
  backend_details_exposed
  backend_launch_enabled
  backend_process_started
  compatibility_storage_exposed
  direct_host_file_access
  direct_file_read_enabled
  execution_started
  file_content_read
  file_paths_exposed
  file_writes_performed
  host_permission_changed
  host_root_modified
  ai_provider_call_enabled
  launch_allowed
  launch_enabled
  live_tray_bridge_enabled
  network_required
  notifications_sent
  permission_granted
  privileged_container_required
  raw_command_exposed
  raw_executable_exposed
  raw_windows_executable_exposed
  request_object_created
  request_objects_created
  settings_persisted
  task_manager_entry_active
].freeze

class SmokeFailure < StandardError; end

def parse_options(argv)
  options = {
    go_bin: DEFAULT_GO_BIN,
    registry: DEFAULT_REGISTRY,
    app: DEFAULT_APP,
    decision: DEFAULT_DECISION,
    mode: DEFAULT_MODE,
    format: "text"
  }

  parser = OptionParser.new do |opts|
    opts.banner = "Usage: ruby scripts/kde_first_presence_smoke.rb [options] [file-uri]"
    opts.on("--go-bin PATH", "Go binary to use") { |value| options[:go_bin] = value }
    opts.on("--registry PATH", "Recipe registry path") { |value| options[:registry] = value }
    opts.on("--app ID", "Application id") { |value| options[:app] = value }
    opts.on("--decision DECISION", "Review decision") { |value| options[:decision] = value }
    opts.on("--mode MODE", "Activation mode") { |value| options[:mode] = value }
    opts.on("--format FORMAT", "Output format: text, json, or markdown") { |value| options[:format] = value }
  end
  parser.parse!(argv)
  options[:file_uri] = argv.first || DEFAULT_FILE_URI
  options
end

def go_preview(options, command, *arguments)
  stdout = go_preview_output(options, command, *arguments)
  JSON.parse(stdout)
rescue JSON::ParserError => e
  raise SmokeFailure, "#{command} did not return JSON: #{e.message}"
end

def go_preview_text(options, command, *arguments)
  go_preview_output(options, command, *arguments)
end

def go_preview_output(options, command, *arguments)
  env = {
    "GOCACHE" => PROJECT_ROOT.join(".cache", "go-build").to_s
  }
  command_line = [
    options.fetch(:go_bin),
    "run",
    "./cmd/xnix-runtime-go",
    command,
    *arguments
  ]
  stdout, stderr, status = Open3.capture3(env, *command_line, chdir: PROJECT_ROOT.to_s)
  raise SmokeFailure, "#{command} failed: #{stderr.strip}" unless status.success?

  stdout
end

def base_args(options)
  ["--registry", options.fetch(:registry), "--app", options.fetch(:app)]
end

def decision_args(options)
  [*base_args(options), "--decision", options.fetch(:decision), options.fetch(:file_uri)]
end

def mode_args(options)
  [*base_args(options), "--mode", options.fetch(:mode)]
end

def collect_payloads(options)
  {
    "desktop-safety-policy-preview" => go_preview(options, "desktop-safety-policy-preview"),
    "desktop-identity-plan" => go_preview(options, "desktop-identity-plan", *base_args(options)),
    "desktop-entry-preview" => go_preview_text(options, "desktop-entry-preview", *base_args(options)),
    "mimeapps-preview" => go_preview_text(options, "mimeapps-preview", *base_args(options)),
    "desktop-activation-bundle-preview" => go_preview(options, "desktop-activation-bundle-preview", *base_args(options)),
    "desktop-activation-staging-preview" => go_preview(options, "desktop-activation-staging-preview", *mode_args(options)),
    "desktop-activation-transaction-preview" => go_preview(options, "desktop-activation-transaction-preview", *mode_args(options)),
    "desktop-activation-status-preview" => go_preview(options, "desktop-activation-status-preview", *mode_args(options)),
    "kde-entrypoints-preview" => go_preview(options, "kde-entrypoints-preview", *decision_args(options)),
    "kde-action-card-deck-preview" => go_preview(options, "kde-action-card-deck-preview", *decision_args(options)),
    "kde-action-dependency-graph-preview" => go_preview(options, "kde-action-dependency-graph-preview", *decision_args(options)),
    "kde-journey-evidence-preview" => go_preview(options, "kde-journey-evidence-preview", *decision_args(options)),
    "compatibility-onboarding-checklist-preview" => go_preview(options, "compatibility-onboarding-checklist-preview", *base_args(options), "--root", "."),
    "support-bundle-manifest-preview" => go_preview(options, "support-bundle-manifest-preview", *base_args(options), "--runtime-root", "."),
    "multi-application-install-queue-preview" => go_preview(options, "multi-application-install-queue-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), "--mode", options.fetch(:mode)),
    "runtime-policy-explanation-cards-preview" => go_preview(options, "runtime-policy-explanation-cards-preview", *base_args(options), "--runtime-root", ".", "--mode", options.fetch(:mode), "--issue", "portal-approval-required"),
    "kde-center-page-preview" => go_preview(options, "kde-center-page-preview", *decision_args(options)),
    "kde-center-page-sections-preview" => go_preview(options, "kde-center-page-sections-preview", *decision_args(options)),
    "kde-center-page-section-detail-preview" => go_preview(options, "kde-center-page-section-detail-preview", *base_args(options), "--section", "diagnostics", "--decision", options.fetch(:decision), options.fetch(:file_uri)),
    "file-open-preview" => go_preview(options, "file-open-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "dolphin-drop-preview" => go_preview(options, "dolphin-drop-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "dolphin-ai-analysis-preview" => go_preview(options, "dolphin-ai-analysis-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "window-identity-preview" => go_preview(options, "window-identity-preview", *base_args(options)),
    "tray-status-preview" => go_preview(options, "tray-status-preview", *base_args(options)),
    "notification-preview" => go_preview(options, "notification-preview", *base_args(options), "--event", "approval-required"),
    "settings-preview" => go_preview(options, "settings-preview", *base_args(options)),
    "runtime-owner-route-manifest-preview" => go_preview(options, "runtime-owner-route-manifest-preview", "--root", "."),
    "runtime-route-convergence-preview" => go_preview(options, "runtime-route-convergence-preview", "--root", "."),
    "runtime-method-parity-manifest-preview" => go_preview(options, "runtime-method-parity-manifest-preview", "--root", "."),
    "runtime-write-gate-preview" => go_preview(options, "runtime-write-gate-preview", "--root", ".", "--method", "Launch"),
    "state-root-preview" => go_preview(options, "state-root-preview", *base_args(options)),
    "snapshot-plan-preview" => go_preview(options, "snapshot-plan-preview", "--app", options.fetch(:app), "--reason", "before-repair"),
    "portal-access-policy-preview" => go_preview(options, "portal-access-policy-preview", "--app", options.fetch(:app), "--operation", "file-open"),
    "ai-diagnostic-input-preview" => go_preview(options, "ai-diagnostic-input-preview", *base_args(options)),
    "ai-diagnostic-recommendation-preview" => go_preview(options, "ai-diagnostic-recommendation-preview", *base_args(options)),
    "ai-repair-approval-gate-preview" => go_preview(options, "ai-repair-approval-gate-preview", *base_args(options))
  }
end

def assert(condition, check_id, message)
  raise SmokeFailure, "#{check_id}: #{message}" unless condition
end

def assert_text_artifacts(desktop_entry, mimeapps, desktop_file)
  assert(desktop_entry.include?("[Desktop Entry]"), "start_menu.desktop_entry", "desktop entry text must be rendered")
  assert(desktop_entry.include?("Exec=xnix-compat-launch --app"), "start_menu.launcher", "desktop entry must route launch through the Runtime")
  assert(desktop_entry.include?("X-Xnix-RuntimeOwned=true"), "start_menu.runtime_owned", "desktop entry must identify Runtime ownership")
  assert(mimeapps.include?(desktop_file), "file_manager.mimeapps", "MIME apps output must target the generated desktop file")
end

def deep_each(value, path = [], &block)
  yield(path, value)
  case value
  when Hash
    value.each { |key, child| deep_each(child, [*path, key], &block) }
  when Array
    value.each_with_index { |child, index| deep_each(child, [*path, index], &block) }
  end
end

def safe_payload?(payload, allowed_file_uri)
  raise SmokeFailure, "global_safety.policy: desktop safety policy must be supplied" unless @desktop_safety_policy

  safety_false_keys = (@desktop_safety_policy.fetch("safety_false_keys") + SAFETY_FALSE_KEYS).uniq
  forbidden_patterns = forbidden_text_patterns(@desktop_safety_policy)
  deep_each(payload) do |path, value|
    key = path.last.to_s
    if safety_false_keys.include?(key) && value == true
      raise SmokeFailure, "global_safety.#{key}: expected false, got true"
    end
    next unless value.is_a?(String)
    next if path.include?("forbidden_user_terms")
    next if path.include?("safety_false_keys")
    next if value == allowed_file_uri
    next if value.start_with?("applications:")
    next if value.start_with?("usr/share/")

    forbidden_patterns.each do |pattern|
      raise SmokeFailure, "global_safety.text: forbidden value at #{path.join(".")}: #{value}" if value.match?(pattern)
    end
  end
end

def forbidden_text_patterns(policy)
  policy.fetch("forbidden_user_terms").map do |term|
    case term
    when "prefix", "bottle", "wine", "proton"
      /\b#{Regexp.escape(term)}\b/i
    when ".exe"
      /\.exe\b/i
    else
      /#{Regexp.escape(term)}/i
    end
  end
end

def entry_point_ids(payload)
  payload.fetch("entry_point_ids") do
    payload.fetch("entry_points", []).map { |entry| entry.fetch("id") }
  end
end

def assert_entrypoints(payload)
  ids = entry_point_ids(payload)
  missing = EXPECTED_ENTRYPOINTS - ids
  assert(missing.empty?, "kde_entrypoints.complete", "missing entrypoints: #{missing.join(", ")}")
  assert(payload.fetch("surface_type") == "kde-first-release-entrypoints", "kde_entrypoints.surface", "unexpected surface type")
  assert(payload.fetch("normal_application_surface") == true, "kde_entrypoints.normal_surface", "normal application surface must be true")
  assert(payload.fetch("kde_policy_owner") == false, "kde_entrypoints.runtime_boundary", "KDE must not own policy")
end

def assert_staging(payload)
  ids = payload.fetch("planned_file_ids")
  missing = EXPECTED_STAGED_FILES - ids
  assert(missing.empty?, "activation_staging.files", "missing staged files: #{missing.join(", ")}")
  assert(payload.fetch("file_writes_performed") == false, "activation_staging.non_executing", "staging preview must not write files")
  assert(payload.fetch("rollback_receipt_planned") == true, "activation_staging.rollback", "rollback receipt must be planned")
end

def assert_transaction(payload)
  assert(payload.fetch("transaction_committed") == false, "activation_transaction.commit", "activation commit must remain disabled")
  assert(payload.fetch("write_gate").fetch("dispatch_enabled") == false, "activation_transaction.write_gate", "activation write dispatch must remain disabled")
  receipt = payload.fetch("receipt_evidence")
  assert(receipt.fetch("evidence_state") == "planned-runtime-gated", "activation_transaction.receipt_state", "receipt evidence must remain planned and Runtime gated")
  assert(receipt.fetch("commit_receipt_planned") == true, "activation_transaction.commit_receipt", "commit receipt must be planned")
  assert(receipt.fetch("rollback_receipt_planned") == true, "activation_transaction.rollback_receipt", "rollback receipt must be planned")
  assert(receipt.fetch("commit_available") == false, "activation_transaction.commit_available", "commit must remain unavailable")
  assert(receipt.fetch("rollback_available") == false, "activation_transaction.rollback_available", "rollback must remain unavailable before commit")
  assert(receipt.fetch("host_root_modified") == false, "activation_transaction.receipt_host", "receipt evidence must not mutate the host root")
  assert(receipt.fetch("backend_details_exposed") == false, "activation_transaction.receipt_backend", "receipt evidence must not expose backend details")
  rollback_steps = payload.fetch("rollback_steps", [])
  assert(!rollback_steps.empty?, "activation_transaction.rollback", "rollback steps must be present")
end

def assert_file_manager(file_open, drop, ai_analysis)
  assert(file_open.fetch("source") == "dolphin-service-menu", "file_manager.source", "file-open source must be Dolphin service menu")
  assert(file_open.fetch("portal_required") == true, "file_manager.portal", "file-open must require Portal review")
  assert(drop.fetch("source") == "dolphin-drag-and-drop", "file_manager.drop", "drop source must be Dolphin drag-and-drop")
  assert(ai_analysis.fetch("source") == "dolphin-ai-action", "file_manager.ai_source", "AI analysis source must be Dolphin AI action")
  assert(ai_analysis.fetch("file_content_read") == false, "file_manager.ai_privacy", "AI analysis must not read file contents")
  assert(ai_analysis.fetch("file_paths_exposed") == false, "file_manager.ai_paths", "AI analysis must not expose file paths")
end

def assert_center(page, sections, diagnostics)
  assert(page.fetch("request_type") == "kde-center-page-preview", "center.page", "center page preview required")
  assert(page.fetch("action_deck").fetch("request_type") == "kde-action-card-deck-preview", "center.deck", "center page must include action deck")
  assert(page.fetch("action_dependency_graph").fetch("request_type") == "kde-action-dependency-graph-preview", "center.action_graph", "center page must include action dependency graph")
  assert(page.fetch("action_dependency_graph").fetch("blocked_action_count") == 7, "center.action_graph_blocked", "center page action graph must keep all actions blocked")
  assert(page.fetch("action_dependency_graph").fetch("dependency_graph_persisted") == false, "center.action_graph_persistence", "center page action graph must not persist state")
  assert(page.fetch("action_dependency_graph").fetch("execution_started") == false, "center.action_graph_execution", "center page action graph must not start execution")
  section_ids = sections.fetch("sections").map { |section| section.fetch("id") }
  %w[overview backend activation execution launch window files tray notifications actions settings diagnostics].each do |section|
    assert(section_ids.include?(section), "center.section.#{section}", "center sections must include #{section}")
  end
  assert(diagnostics.fetch("section_id") == "diagnostics", "center.diagnostics", "diagnostics detail required")
end

def assert_action_dependency_graph(payload)
  assert(payload.fetch("request_type") == "kde-action-dependency-graph-preview", "action_graph.request_type", "action dependency graph preview required")
  assert(payload.fetch("graph_type") == "compatibility-center-action-dependency-graph", "action_graph.graph_type", "action dependency graph type required")
  assert(payload.fetch("action_node_count") == 7, "action_graph.actions", "action dependency graph must cover seven actions")
  assert(payload.fetch("missing_evidence_count") > 0, "action_graph.missing_evidence", "action dependency graph must expose missing evidence")
  assert(payload.fetch("blocked_action_count") == 7, "action_graph.blocked_actions", "action dependency graph must keep all actions blocked")
  assert(payload.fetch("dependency_graph_persisted") == false, "action_graph.persistence", "dependency graph preview must not persist state")
  assert(payload.fetch("request_objects_created") == false, "action_graph.requests", "dependency graph preview must not create request objects")
  assert(payload.fetch("permission_grant_created") == false, "action_graph.permissions", "dependency graph preview must not grant permissions")
  assert(payload.fetch("settings_persisted") == false, "action_graph.settings", "dependency graph preview must not persist settings")
  assert(payload.fetch("execution_started") == false, "action_graph.execution", "dependency graph preview must not start execution")
  assert(payload.fetch("host_root_modified") == false, "action_graph.host_root", "dependency graph preview must not mutate host root")
  assert(payload.fetch("backend_details_exposed") == false, "action_graph.backend", "dependency graph preview must not expose backend details")
  validation = payload.fetch("receipt_validation")
  assert(validation.fetch("rejects_mismatched_app_id") == true, "action_graph.receipt_app", "receipt validation must reject mismatched app ids")
  assert(validation.fetch("rejects_malformed_operation_id") == true, "action_graph.receipt_operation", "receipt validation must reject malformed operation ids")
  assert(validation.fetch("rejects_path_escape_evidence") == true, "action_graph.receipt_path", "receipt validation must reject path escape evidence")
  assert(validation.fetch("rejects_unsafe_side_effects") == true, "action_graph.receipt_side_effects", "receipt validation must reject unsafe side effects")
end

def assert_journey_evidence(payload)
  assert(payload.fetch("request_type") == "kde-journey-evidence-preview", "journey.request_type", "KDE journey evidence preview required")
  assert(payload.fetch("journey_type") == "kde-seven-entrypoint-runtime-evidence", "journey.type", "KDE journey evidence type required")
  assert(payload.fetch("entry_point_count") == 7, "journey.entrypoints", "journey evidence must cover seven KDE entrypoints")
  ids = payload.fetch("entry_point_ids")
  %w[launcher task-manager file-manager system-tray notification-center ai-compatibility-center unified-settings].each do |entrypoint|
    assert(ids.include?(entrypoint), "journey.entrypoint.#{entrypoint}", "journey evidence must include #{entrypoint}")
  end
  assert(payload.fetch("shared_readiness_status") == "not-ready", "journey.readiness", "journey evidence must use shared readiness state")
  assert(payload.fetch("missing_evidence_count") > 0, "journey.missing_evidence", "journey evidence must expose missing evidence")
  assert(payload.fetch("blocked_action_count") == 7, "journey.blocked_actions", "journey evidence must keep seven actions blocked")
  assert(payload.fetch("journey_evidence_persisted") == false, "journey.persistence", "journey evidence must not persist state")
  assert(payload.fetch("runtime_write_methods_enabled") == false, "journey.write_methods", "journey evidence must not enable write methods")
  assert(payload.fetch("request_objects_created") == false, "journey.requests", "journey evidence must not create request objects")
  assert(payload.fetch("permission_grant_created") == false, "journey.permissions", "journey evidence must not grant permissions")
  assert(payload.fetch("settings_persisted") == false, "journey.settings", "journey evidence must not persist settings")
  assert(payload.fetch("launch_enabled") == false, "journey.launch", "journey evidence must not enable launch")
  assert(payload.fetch("execution_started") == false, "journey.execution", "journey evidence must not start execution")
  assert(payload.fetch("kwin_rule_applied") == false, "journey.kwin", "journey evidence must not apply KWin rules")
  assert(payload.fetch("tray_bridge_activated") == false, "journey.tray", "journey evidence must not activate tray bridge")
  assert(payload.fetch("notification_sent") == false, "journey.notification", "journey evidence must not send notifications")
  assert(payload.fetch("host_root_modified") == false, "journey.host_root", "journey evidence must not mutate host root")
  assert(payload.fetch("backend_details_exposed") == false, "journey.backend", "journey evidence must not expose backend details")
  agreement = payload.fetch("agreement")
  assert(agreement.fetch("app_id_consistent") == true, "journey.app_consistent", "journey evidence must keep app id consistent")
  assert(agreement.fetch("display_name_consistent") == true, "journey.name_consistent", "journey evidence must keep display name consistent")
  assert(agreement.fetch("disabled_action_state_shared") == true, "journey.disabled_shared", "journey evidence must share disabled state")
  read_models = payload.fetch("cross_linked_read_models")
  %w[application-readiness-preview kde-action-dependency-graph-preview task-manager-identity-preview kwin-window-rule-preview tray-status-preview notification-preview kde-center-page-preview settings-preview].each do |read_model|
    assert(read_models.include?(read_model), "journey.read_model.#{read_model}", "journey evidence must cross-link #{read_model}")
  end
end

def assert_onboarding_checklist(payload)
  assert(payload.fetch("request_type") == "compatibility-onboarding-checklist-preview", "onboarding.request_type", "compatibility onboarding checklist preview required")
  assert(payload.fetch("checklist_type") == "first-run-compatibility-onboarding", "onboarding.type", "first-run onboarding checklist type required")
  assert(payload.fetch("section_count") == 9, "onboarding.sections", "onboarding checklist must expose nine sections")
  assert(payload.fetch("runtime_owned") == true, "onboarding.runtime_owned", "Runtime must own onboarding policy")
  assert(payload.fetch("go_runtime_backed") == true, "onboarding.go_backed", "onboarding checklist must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "onboarding.kde_owner", "KDE must not own onboarding policy")
  assert(payload.fetch("ready") == false, "onboarding.ready", "first-run onboarding must not mark the app ready without evidence")
  states = payload.fetch("states")
  %w[ready needs_review missing_evidence blocked not_yet_implemented].each do |state_key|
    assert(states.fetch(state_key) > 0, "onboarding.state.#{state_key}", "onboarding checklist must include #{state_key}")
  end
  ids = payload.fetch("section_ids")
  %w[runtime-owner-readiness recipe-trust artifact-staging backend-lifecycle portal-review snapshot-baseline diagnostics-privacy kde-entry-points production-activation].each do |section_id|
    assert(ids.include?(section_id), "onboarding.section.#{section_id}", "onboarding checklist must include #{section_id}")
  end
  %w[runtime_write_methods_enabled request_objects_created permission_grants_created artifact_staged settings_persisted backend_process_started launch_enabled network_required host_package_manager_invoked host_root_modified privileged_container_required ai_provider_called state_root_path_exposed raw_executable_exposed raw_command_exposed file_content_read backend_details_exposed].each do |key|
    assert(payload.fetch(key) == false, "onboarding.disabled.#{key}", "#{key} must remain false")
  end
  payload.fetch("sections").each do |section|
    assert(section.fetch("runtime_owned") == true, "onboarding.section_runtime.#{section.fetch("id")}", "section must be Runtime owned")
    assert(section.fetch("go_runtime_backed") == true, "onboarding.section_go.#{section.fetch("id")}", "section must be Go Runtime backed")
    assert(section.fetch("kde_policy_owner") == false, "onboarding.section_kde.#{section.fetch("id")}", "section must not make KDE a policy owner")
    assert(section.fetch("side_effects_enabled") == false, "onboarding.section_side_effects.#{section.fetch("id")}", "section must not enable side effects")
    assert(section.fetch("host_root_modified") == false, "onboarding.section_host.#{section.fetch("id")}", "section must not mutate host root")
    assert(section.fetch("backend_details_exposed") == false, "onboarding.section_backend.#{section.fetch("id")}", "section must not expose backend details")
  end
end

def assert_support_bundle_manifest(payload)
  assert(payload.fetch("request_type") == "support-bundle-manifest-preview", "support_bundle.request_type", "support bundle manifest preview required")
  assert(payload.fetch("manifest_type") == "redacted-offline-support-bundle-manifest", "support_bundle.type", "redacted offline support bundle manifest required")
  assert(payload.fetch("section_count") == 7, "support_bundle.sections", "support bundle manifest must expose seven sections")
  assert(payload.fetch("runtime_owned") == true, "support_bundle.runtime_owned", "Runtime must own support bundle manifest policy")
  assert(payload.fetch("go_runtime_backed") == true, "support_bundle.go_backed", "support bundle manifest must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "support_bundle.kde_owner", "KDE must not own support bundle manifest policy")
  assert(payload.fetch("user_visible") == true, "support_bundle.user_visible", "support bundle manifest must be user visible")
  assert(payload.fetch("offline_only") == true, "support_bundle.offline", "support bundle manifest must remain offline-only")
  ids = payload.fetch("section_ids")
  %w[runtime-version application-identity diagnostic-summaries failing-signals repair-recommendations ai-diagnostic-boundary omitted-evidence].each do |section_id|
    assert(ids.include?(section_id), "support_bundle.section.#{section_id}", "support bundle manifest must include #{section_id}")
  end
  %w[archive_created file_content_read file_paths_exposed ai_provider_called ai_provider_call_enabled auto_repair_requested auto_repair_executed backend_process_started host_root_modified state_root_path_exposed raw_executable_exposed raw_command_exposed backend_details_exposed network_required privileged_container_required].each do |key|
    assert(payload.fetch(key) == false, "support_bundle.disabled.#{key}", "#{key} must remain false")
  end
  redaction = payload.fetch("privacy_redaction")
  assert(redaction.fetch("redaction_status") == "redacted-preview-only", "support_bundle.redaction.status", "support bundle manifest must be redacted")
  %w[user_documents_included host_paths_included state_root_paths_included environment_variables_included token_shaped_values_included command_shaped_values_included usernames_included secrets_included network_calls_allowed].each do |key|
    assert(redaction.fetch(key) == false, "support_bundle.redaction.#{key}", "#{key} must remain false")
  end
  omitted = payload.fetch("omitted_evidence")
  %w[file_contents host_paths environment_variables token_shaped_values command_shaped_values usernames].each do |key|
    assert(omitted.fetch(key) > 0, "support_bundle.omitted.#{key}", "support bundle manifest must count omitted #{key}")
  end
  payload.fetch("sections").each do |section|
    assert(section.fetch("side_effects_enabled") == false, "support_bundle.section_side_effects.#{section.fetch("id")}", "section must not enable side effects")
    assert(section.fetch("host_root_modified") == false, "support_bundle.section_host.#{section.fetch("id")}", "section must not mutate host root")
    assert(section.fetch("backend_details_exposed") == false, "support_bundle.section_backend.#{section.fetch("id")}", "section must not expose backend details")
  end
end

def assert_multi_application_install_queue(payload)
  assert(payload.fetch("request_type") == "multi-application-install-queue-preview", "multi_app_install_queue.request_type", "multi-application install queue preview required")
  assert(payload.fetch("queue_type") == "review-only-multi-application-install-queue", "multi_app_install_queue.type", "review-only queue type required")
  assert(payload.fetch("schema_version") == "xnix.runtime.multi_application_install_queue.v1", "multi_app_install_queue.schema", "multi-application install queue schema required")
  assert(payload.fetch("runtime_method") == "GetMultiApplicationInstallQueue", "multi_app_install_queue.runtime_method", "Runtime queue method required")
  assert(payload.fetch("read_method") == "GetMultiApplicationInstallQueuePreview", "multi_app_install_queue.read_method", "Runtime queue read method required")
  assert(payload.fetch("runtime_owned") == true, "multi_app_install_queue.runtime_owned", "Runtime must own the queue")
  assert(payload.fetch("go_runtime_backed") == true, "multi_app_install_queue.go_backed", "queue preview must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "multi_app_install_queue.kde_owner", "KDE must not own queue policy")
  assert(payload.fetch("user_visible") == true, "multi_app_install_queue.user_visible", "queue preview must be user visible")
  assert(payload.fetch("review_only") == true, "multi_app_install_queue.review_only", "queue preview must remain review-only")
  assert(payload.fetch("application_count") >= 1, "multi_app_install_queue.application_count", "queue preview must include at least one app")
  assert(payload.fetch("queue_status") == "missing-evidence", "multi_app_install_queue.status", "fixture queue must remain missing evidence")
  assert(payload.fetch("ready_for_review") == true, "multi_app_install_queue.ready_for_review", "queue preview must be ready for review")
  counts = payload.fetch("counts")
  assert(counts.fetch("total") == payload.fetch("application_count"), "multi_app_install_queue.counts.total", "queue counts must match application count")
  assert(counts.fetch("missing_evidence") >= 1, "multi_app_install_queue.counts.missing", "queue must surface missing evidence")
  %w[queue_persisted request_objects_created artifacts_staged artifacts_downloaded network_request_created host_package_manager_invoked desktop_activation_started install_started backend_process_started launch_enabled execution_started settings_persisted host_root_modified privileged_container_required state_root_path_exposed raw_executable_exposed raw_command_exposed backend_details_exposed].each do |key|
    assert(payload.fetch(key) == false, "multi_app_install_queue.disabled.#{key}", "#{key} must remain false")
  end
  payload.fetch("applications").each do |application|
    assert(application.fetch("queue_state") == "missing-evidence", "multi_app_install_queue.item_state.#{application.fetch("application_id")}", "fixture application must remain missing evidence")
    assert(application.fetch("user_review_required") == true, "multi_app_install_queue.item_review.#{application.fetch("application_id")}", "fixture application must require review")
    %w[request_object_created artifact_staged install_started backend_process_started launch_enabled execution_started host_root_modified state_root_path_exposed raw_command_exposed backend_details_exposed].each do |key|
      assert(application.fetch(key) == false, "multi_app_install_queue.item_disabled.#{key}", "#{key} must remain false")
    end
  end
end

def assert_runtime_policy_explanation_cards(payload)
  assert(payload.fetch("schema_version") == "xnix.runtime.policy_explanation_cards.v1", "policy_cards.schema", "Runtime policy explanation card schema required")
  assert(payload.fetch("request_type") == "runtime-policy-explanation-cards-preview", "policy_cards.request_type", "Runtime policy explanation card preview required")
  assert(payload.fetch("card_deck_type") == "kde-runtime-policy-explanation-card-deck", "policy_cards.type", "KDE Runtime policy explanation card deck required")
  assert(payload.fetch("runtime_method") == "GetRuntimePolicyExplanationCards", "policy_cards.runtime_method", "Runtime policy explanation method required")
  assert(payload.fetch("read_method") == "GetRuntimePolicyExplanationCardsPreview", "policy_cards.read_method", "Runtime policy explanation read method required")
  assert(payload.fetch("runtime_owned") == true, "policy_cards.runtime_owned", "Runtime must own policy explanations")
  assert(payload.fetch("go_runtime_backed") == true, "policy_cards.go_backed", "policy explanations must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "policy_cards.kde_owner", "KDE must not own policy explanations")
  assert(payload.fetch("review_only") == true, "policy_cards.review_only", "policy explanations must remain review-only")
  assert(payload.fetch("card_count") == 11, "policy_cards.count", "policy explanation deck must include 11 cards")
  expected_cards = %w[install launch execution portal-permission snapshot diagnostics repair settings desktop-activation backend-readiness unsupported-production-route]
  missing_cards = expected_cards - payload.fetch("card_ids")
  assert(missing_cards.empty?, "policy_cards.ids", "missing policy explanation cards: #{missing_cards.join(", ")}")
  counts = payload.fetch("counts")
  %w[blocked review_only missing_evidence not_yet_implemented].each do |key|
    assert(counts.fetch(key) >= 1, "policy_cards.counts.#{key}", "policy card count #{key} must be present")
  end
  %w[cards_persisted action_enablement_changed request_objects_created permission_grants_created settings_persisted ai_provider_called ai_provider_call_enabled backend_process_started launch_enabled execution_started host_root_modified network_required privileged_container_required state_root_path_exposed raw_executable_exposed raw_command_exposed backend_details_exposed].each do |key|
    assert(payload.fetch(key) == false, "policy_cards.disabled.#{key}", "#{key} must remain false")
  end
  payload.fetch("cards").each do |card|
    %w[action_enabled card_persisted request_object_created permission_granted settings_persisted ai_provider_called backend_process_started launch_enabled execution_started host_root_modified state_root_path_exposed raw_command_exposed backend_details_exposed].each do |key|
      assert(card.fetch(key) == false, "policy_cards.card_disabled.#{card.fetch("id")}.#{key}", "#{key} must remain false")
    end
    summary = card.fetch("user_facing_summary").downcase
    %w[backend prefix wine proton .exe qemu].each do |term|
      assert(!summary.include?(term), "policy_cards.user_summary.#{card.fetch("id")}.#{term}", "user-facing policy summary must hide #{term}")
    end
  end
end

def assert_desktop_safety_policy(payload)
  assert(payload.fetch("schema_version") == "xnix.runtime.desktop_safety_policy.v1", "desktop_safety.schema", "desktop safety policy schema required")
  assert(payload.fetch("request_type") == "desktop-safety-policy-preview", "desktop_safety.request_type", "desktop safety policy preview required")
  assert(payload.fetch("policy_type") == "kde-first-user-facing-safety-policy", "desktop_safety.policy_type", "KDE-first policy type required")
  assert(payload.fetch("runtime_owned") == true, "desktop_safety.runtime_owned", "Runtime must own desktop safety policy")
  assert(payload.fetch("go_runtime_backed") == true, "desktop_safety.go_backed", "desktop safety policy must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "desktop_safety.kde_owner", "KDE must not own desktop safety policy")
  assert(payload.fetch("backend_terminology_hidden") == true, "desktop_safety.hidden_terms", "backend terminology must be hidden")
  assert(payload.fetch("entrypoint_count") == 7, "desktop_safety.entrypoint_count", "desktop safety policy must cover seven entrypoints")
  missing_entrypoints = EXPECTED_ENTRYPOINTS - payload.fetch("entrypoints")
  assert(missing_entrypoints.empty?, "desktop_safety.entrypoints", "missing entrypoints: #{missing_entrypoints.join(", ")}")
  missing_settings = %w[mode preference documents downloads camera network snapshots] - payload.fetch("settings_field_ids")
  assert(missing_settings.empty?, "desktop_safety.settings", "missing settings fields: #{missing_settings.join(", ")}")
  %w[prefix bottle wine proton .exe docker.sock].each do |term|
    assert(payload.fetch("forbidden_user_terms").include?(term), "desktop_safety.forbidden.#{term}", "missing forbidden user term: #{term}")
  end
  %w[backend_launch_enabled execution_started host_root_modified real_portal_transport_enabled request_object_created settings_persisted].each do |key|
    assert(payload.fetch("safety_false_keys").include?(key), "desktop_safety.false_key.#{key}", "missing false safety key: #{key}")
  end
  %w[write_methods_enabled backend_launch_enabled execution_enabled real_portal_transport_enabled ai_provider_call_enabled host_root_modified network_required privileged_container_required].each do |key|
    assert(payload.fetch(key) == false, "desktop_safety.disabled.#{key}", "#{key} must remain false")
  end
end

def assert_settings(payload, policy)
  assert(payload.fetch("request_type") == "settings-preview", "settings.request_type", "settings preview required")
  assert(payload.fetch("kde_policy_owner") == false, "settings.runtime_boundary", "KDE must not own settings policy")
  assert(payload.fetch("settings_persisted") == false, "settings.persistence", "settings preview must not persist settings")
  sections = payload.fetch("sections")
  field_ids = sections.flat_map { |section| section.fetch("fields").map { |field| field.fetch("id") } }
  missing = policy.fetch("settings_field_ids") - field_ids
  assert(missing.empty?, "settings.user_fields", "missing user-facing settings fields: #{missing.join(", ")}")
  section_titles = sections.map { |section| section.fetch("title") }
  %w[Run File Devices Network Snapshots].each do |term|
    assert(section_titles.any? { |title| title.include?(term) }, "settings.section.#{term.downcase}", "settings must expose #{term} section")
  end
end

def assert_route_baseline(payload)
  counts = payload.fetch("route_counts")
  checks = payload.fetch("checks").to_h { |check| [check.fetch("id"), check.fetch("status")] }
  assert(counts.fetch("ruby_legacy") == 0, "routes.ruby_legacy", "Ruby legacy routes must be zero")
  assert(checks.fetch("method-parity") == "pass", "routes.method_parity", "method parity must pass")
  assert(checks.fetch("write-route-gate") == "pass", "routes.write_gate", "write route gate must pass")
  assert(checks.fetch("host-safety-boundary") == "pass", "routes.host_safety", "host safety boundary must pass")
end

def assert_route_convergence(payload)
  counts = payload.fetch("classification_counts")
  checks = payload.fetch("checks").to_h { |check| [check.fetch("id"), check.fetch("status")] }
  assert(payload.fetch("request_type") == "runtime-route-convergence-preview", "route_convergence.request_type", "route convergence preview required")
  assert(payload.fetch("all_routes_classified") == true, "route_convergence.classified", "all Runtime read routes must be classified")
  assert(payload.fetch("native_go_coverage_ready") == true, "route_convergence.go_coverage", "Runtime read routes must have Go coverage")
  assert(payload.fetch("production_owner_ready") == false, "route_convergence.production_owner", "route convergence must not enable production ownership")
  assert(payload.fetch("write_methods_enabled") == false, "route_convergence.write_methods", "route convergence must not enable write methods")
  assert(payload.fetch("host_root_modified") == false, "route_convergence.host_root", "route convergence must not mutate host root")
  assert(counts.fetch("unclassified") == 0, "route_convergence.unclassified", "unclassified Runtime routes must be zero")
  assert(checks.fetch("write-gate-disabled") == "pass", "route_convergence.write_gate", "write gate must remain disabled")
  assert(checks.fetch("host-safety-boundary") == "pass", "route_convergence.host_safety", "host safety boundary must pass")
end

def assert_method_parity(payload)
  assert(payload.fetch("read_only_method_parity_ready") == true, "method_parity.ready", "read-only method parity must be ready")
  assert(payload.fetch("write_method_dispatch_enabled") == false, "method_parity.write_dispatch", "write dispatch must remain disabled")
end

def assert_write_gate(payload)
  assert(payload.fetch("request_type") == "runtime-write-gate-preview", "write_gate.request_type", "write gate preview required")
  assert(payload.fetch("runtime_method") == "GetRuntimeWriteGate", "write_gate.runtime_method", "write gate must map to Runtime method")
  assert(payload.fetch("method_name") == "Launch", "write_gate.method", "Launch gate must be evaluated")
  assert(payload.fetch("write_method_enabled") == false, "write_gate.disabled", "Launch must remain disabled")
  assert(payload.fetch("dispatch_enabled") == false, "write_gate.dispatch", "write dispatch must remain disabled")
  assert(payload.fetch("request_object_created") == false, "write_gate.request_object", "write gate preview must not create requests")
  assert(payload.fetch("execution_started") == false, "write_gate.execution", "write gate preview must not start execution")
  assert(payload.fetch("denial_error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "write_gate.error", "write gate error must be explicit")
  assert(payload.fetch("go_runtime_backed") == true, "write_gate.go_owner", "write gate must be Go Runtime backed")
end

def assert_runtime_safety_substrate(state_root, snapshot_plan, portal_policy)
  assert(state_root.fetch("request_type") == "state-root-preview", "state_root.request_type", "state root preview required")
  assert(state_root.fetch("directories_created") == false, "state_root.no_dirs", "state root preview must not create directories")
  assert(state_root.fetch("user_documents_included") == false, "state_root.no_documents", "state root must exclude user documents")
  assert(state_root.fetch("portal_required_for_user_files") == true, "state_root.portal", "state root must require Portal grants for user files")
  assert(snapshot_plan.fetch("request_type") == "snapshot-plan-preview", "snapshot.request_type", "snapshot plan preview required")
  assert(snapshot_plan.fetch("snapshot_created") == false, "snapshot.not_created", "snapshot preview must not create snapshots")
  assert(snapshot_plan.fetch("restore_executed") == false, "snapshot.no_restore", "snapshot preview must not execute restores")
  assert(snapshot_plan.fetch("user_documents_included") == false, "snapshot.no_documents", "snapshot must exclude user documents")
  assert(portal_policy.fetch("request_type") == "portal-access-policy-preview", "portal_policy.request_type", "Portal access policy preview required")
  assert(portal_policy.fetch("portal_required") == true, "portal_policy.required", "Portal policy must require Portal mediation")
  assert(portal_policy.fetch("direct_access_allowed") == false, "portal_policy.no_direct", "Portal policy must deny direct access")
  assert(portal_policy.fetch("request_object_created") == false, "portal_policy.no_request", "Portal policy preview must not create request objects")
end

def assert_ai_diagnostics(input, recommendation, gate)
  assert(input.fetch("request_type") == "ai-diagnostic-input-preview", "ai.input", "AI diagnostic input preview required")
  assert(input.fetch("runtime_method") == "GetAIDiagnosticInput", "ai.input_method", "AI input must map to Runtime method")
  assert(input.fetch("ai_provider_call_enabled") == false, "ai.no_provider", "AI input must not call a provider")
  assert(input.fetch("file_content_read") == false, "ai.no_file_content", "AI input must not read file contents")
  assert(input.fetch("file_paths_exposed") == false, "ai.no_file_paths", "AI input must not expose file paths")
  assert(recommendation.fetch("request_type") == "ai-diagnostic-recommendation-preview", "ai.recommendation", "AI recommendation preview required")
  assert(recommendation.fetch("auto_execution_allowed") == false, "ai.no_auto_exec", "AI recommendations must not auto-execute")
  assert(gate.fetch("request_type") == "ai-repair-approval-gate-preview", "ai.repair_gate", "AI repair approval gate preview required")
  assert(gate.fetch("gate_decision") == "blocked-until-approval", "ai.repair_blocked", "AI repair must remain approval-gated")
  assert(gate.fetch("repair_executed") == false, "ai.repair_not_executed", "AI repair gate must not execute repairs")
end

def build_success_report(options, payloads)
  route_counts = payloads.fetch("runtime-owner-route-manifest-preview").fetch("route_counts")
  policy = payloads.fetch("desktop-safety-policy-preview")
  {
    "version" => payloads.fetch("runtime-write-gate-preview").fetch("version"),
    "schema_version" => "xnix.kde_first_presence_smoke.v1",
    "report_type" => "kde-first-presence-smoke",
    "application_id" => options.fetch(:app),
    "registry" => options.fetch(:registry),
    "mode" => options.fetch(:mode),
    "decision" => options.fetch(:decision),
    "entrypoints" => EXPECTED_ENTRYPOINTS,
    "entrypoint_count" => EXPECTED_ENTRYPOINTS.length,
    "settings_field_ids" => policy.fetch("settings_field_ids"),
    "planned_staged_files" => EXPECTED_STAGED_FILES,
    "forbidden_user_terms" => policy.fetch("forbidden_user_terms"),
    "preview_commands" => payloads.keys,
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "normal_application_surface" => true,
    "route_baseline" => {
      "total" => route_counts.fetch("total"),
      "go_routed" => route_counts.fetch("go_routed"),
      "c_core_backed" => route_counts.fetch("c_core_backed"),
      "ruby_legacy" => route_counts.fetch("ruby_legacy")
    },
    "write_methods_enabled" => false,
    "execution_enabled" => false,
    "backend_launch_enabled" => false,
    "host_root_modified" => false,
    "network_required" => false,
    "privileged_container_required" => false,
    "real_portal_transport_enabled" => false,
    "ai_provider_call_enabled" => false,
    "desktop_safe_summary" => "Digest-verified KDE application presence is proven across the seven first-release entry points while Runtime execution and host mutation gates remain closed."
  }
end

def render_text(report)
  [
    "PASS: KDE-first presence smoke",
    "application: #{report.fetch("application_id")}",
    "entrypoints: #{report.fetch("entrypoints").join(", ")}",
    "route-baseline: #{report.dig("route_baseline", "total")} total, #{report.dig("route_baseline", "go_routed")} Go-routed, #{report.dig("route_baseline", "c_core_backed")} C-backed, #{report.dig("route_baseline", "ruby_legacy")} Ruby legacy",
    "execution: disabled by Runtime gates",
    "host-root: unchanged"
  ].join("\n")
end

def render_markdown(report)
  lines = []
  lines << "# KDE-First Presence Smoke"
  lines << ""
  lines << "- Version: #{report.fetch("version")}"
  lines << "- Application: `#{report.fetch("application_id")}`"
  lines << "- Entry points: #{report.fetch("entrypoints").join(", ")}"
  lines << "- Settings fields: #{report.fetch("settings_field_ids").join(", ")}"
  lines << "- Forbidden user terms checked: #{report.fetch("forbidden_user_terms").length}"
  lines << "- Preview commands: #{report.fetch("preview_commands").length}"
  lines << "- Runtime owned: #{report.fetch("runtime_owned")}"
  lines << "- Go Runtime backed: #{report.fetch("go_runtime_backed")}"
  lines << "- KDE policy owner: #{report.fetch("kde_policy_owner")}"
  lines << "- Execution enabled: #{report.fetch("execution_enabled")}"
  lines << "- Backend launch enabled: #{report.fetch("backend_launch_enabled")}"
  lines << "- Host root modified: #{report.fetch("host_root_modified")}"
  lines << "- Network required: #{report.fetch("network_required")}"
  lines << "- Privileged container required: #{report.fetch("privileged_container_required")}"
  lines << ""
  lines << "## Route Baseline"
  lines << ""
  lines << "- Total: #{report.dig("route_baseline", "total")}"
  lines << "- Go-routed: #{report.dig("route_baseline", "go_routed")}"
  lines << "- C-backed: #{report.dig("route_baseline", "c_core_backed")}"
  lines << "- Ruby legacy: #{report.dig("route_baseline", "ruby_legacy")}"
  lines << ""
  lines << "## Summary"
  lines << ""
  lines << report.fetch("desktop_safe_summary")
  lines.join("\n")
end

def print_report(report, format)
  case format
  when "text"
    puts render_text(report)
  when "json"
    puts JSON.pretty_generate(report)
  when "markdown"
    puts render_markdown(report)
  else
    raise SmokeFailure, "unsupported format: #{format}"
  end
end

def run_smoke(options)
  payloads = collect_payloads(options)
  @desktop_safety_policy = payloads.fetch("desktop-safety-policy-preview")
  assert_desktop_safety_policy(@desktop_safety_policy)
  payloads.each_value { |payload| safe_payload?(payload, options.fetch(:file_uri)) }

  desktop_file = payloads.fetch("desktop-identity-plan").fetch("desktop_file")
  assert_text_artifacts(payloads.fetch("desktop-entry-preview"), payloads.fetch("mimeapps-preview"), desktop_file)
  assert_entrypoints(payloads.fetch("kde-entrypoints-preview"))
  assert_staging(payloads.fetch("desktop-activation-staging-preview"))
  assert_transaction(payloads.fetch("desktop-activation-transaction-preview"))
  assert_file_manager(
    payloads.fetch("file-open-preview"),
    payloads.fetch("dolphin-drop-preview"),
    payloads.fetch("dolphin-ai-analysis-preview")
  )
  assert_center(
    payloads.fetch("kde-center-page-preview"),
    payloads.fetch("kde-center-page-sections-preview"),
    payloads.fetch("kde-center-page-section-detail-preview")
  )
  assert_action_dependency_graph(payloads.fetch("kde-action-dependency-graph-preview"))
  assert_journey_evidence(payloads.fetch("kde-journey-evidence-preview"))
  assert_onboarding_checklist(payloads.fetch("compatibility-onboarding-checklist-preview"))
  assert_support_bundle_manifest(payloads.fetch("support-bundle-manifest-preview"))
  assert_multi_application_install_queue(payloads.fetch("multi-application-install-queue-preview"))
  assert_runtime_policy_explanation_cards(payloads.fetch("runtime-policy-explanation-cards-preview"))
  assert_settings(payloads.fetch("settings-preview"), @desktop_safety_policy)
  assert_route_baseline(payloads.fetch("runtime-owner-route-manifest-preview"))
  assert_route_convergence(payloads.fetch("runtime-route-convergence-preview"))
  assert_method_parity(payloads.fetch("runtime-method-parity-manifest-preview"))
  assert_write_gate(payloads.fetch("runtime-write-gate-preview"))
  assert_runtime_safety_substrate(
    payloads.fetch("state-root-preview"),
    payloads.fetch("snapshot-plan-preview"),
    payloads.fetch("portal-access-policy-preview")
  )
  assert_ai_diagnostics(
    payloads.fetch("ai-diagnostic-input-preview"),
    payloads.fetch("ai-diagnostic-recommendation-preview"),
    payloads.fetch("ai-repair-approval-gate-preview")
  )

  build_success_report(options, payloads)
end

begin
  options = parse_options(ARGV)
  print_report(run_smoke(options), options.fetch(:format))
rescue SmokeFailure => e
  warn "FAIL: KDE-first presence smoke"
  warn "reason: #{e.message}"
  exit 1
end
