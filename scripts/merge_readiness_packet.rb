#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

TOOL_DEFINITIONS = {
  "layout" => {
    title: "Layout verifier",
    command: ["ruby", "scripts/verify_layout.rb"],
    parser: "text",
    required: true
  },
  "implementation" => {
    title: "Implementation evidence report",
    command: ["ruby", "scripts/implementation_evidence_report.rb", "--format", "json"],
    parser: "json",
    required: true
  },
  "contract_drift" => {
    title: "Runtime contract drift report",
    command: ["ruby", "scripts/runtime_contract_drift_report.rb", "--format", "json"],
    parser: "json",
    required: true
  },
  "kde_smoke" => {
    title: "KDE-first presence smoke",
    command: ["ruby", "scripts/kde_first_presence_smoke.rb", "--format", "json"],
    parser: "json",
    required: true
  },
  "mainline_review" => {
    title: "Mainline integration review",
    command: ["ruby", "scripts/mainline_integration_review.rb", "--format", "json"],
    parser: "json",
    required: true
  },
  "release_evidence" => {
    title: "Release evidence index",
    command: ["ruby", "scripts/release_evidence_index.rb", "--format", "json"],
    parser: "json",
    required: false
  },
  "full_checkpoint_promotion" => {
    title: "Full checkpoint promotion packet",
    command: ["ruby", "scripts/full_checkpoint_promotion_packet.rb", "--format", "json"],
    parser: "json",
    required: false
  },
  "desktop_trigger_request_preflight_smoke" => {
    title: "Desktop-trigger request preflight smoke evidence",
    command: ["ruby", "scripts/desktop_trigger_request_preflight_smoke.rb", "--format", "json"],
    parser: "json",
    required: false,
    fixture_only: true
  },
  "q4_sample_notepad_smoke" => {
    title: "q4 Sample Notepad real Windows app acceptance smoke",
    command: ["ruby", "scripts/q4_sample_notepad_smoke.rb", "--execute"],
    parser: "json",
    required: false,
    fixture_only: true
  },
  "q4_messagebox_smoke" => {
    title: "q4 MessageBox real external Windows app document smoke",
    command: ["ruby", "scripts/q4_messagebox_smoke.rb", "--execute"],
    parser: "json",
    required: false,
    fixture_only: true
  },
  "offline_fixture_matrix" => {
    title: "Offline application fixture matrix",
    command: ["ruby", "scripts/offline_application_fixture_matrix.rb", "--format", "json"],
    parser: "json",
    required: false
  }
}.freeze

FIXTURE_OPTIONS = {
  "layout" => :layout_report,
  "implementation" => :implementation_report,
  "contract_drift" => :contract_drift_report,
  "kde_smoke" => :kde_smoke_report,
  "mainline_review" => :mainline_review,
  "release_evidence" => :release_evidence_index,
  "full_checkpoint_promotion" => :full_checkpoint_promotion,
  "desktop_trigger_request_preflight_smoke" => :desktop_trigger_request_preflight_smoke,
  "q4_sample_notepad_smoke" => :q4_sample_notepad_smoke,
  "q4_messagebox_smoke" => :q4_messagebox_smoke,
  "offline_fixture_matrix" => :fixture_matrix_report
}.freeze

UNSAFE_KEYS = %w[
  automatic_release_tagging_enabled
  automatic_staging_enabled
  backend_launch_enabled
  backend_process_started
  docker_executed
  docker_required
  execution_started
  host_permission_changed
  host_root_modified
  network_checks_run
  network_fetch_enabled
  network_required
  package_manager_invoked
  privileged_container_required
  qemu_executed
  qemu_required
  settings_persisted
].freeze

RELEASE_ONLY_BLOCKERS = %w[
  restricted-docker-or-qemu-smoke-requires-human-authorization
  full-checkpoint-promotion-not-allowed
  desktop-trigger-request-preflight-smoke-not-passed
  q4-sample-notepad-acceptance-smoke-not-passed
  q4-messagebox-document-content-smoke-not-passed
  production-runtime-and-windows-execution-remain-disabled
].freeze

class PacketError < StandardError; end

def parse_options(argv)
  options = {
    format: "json",
    offline_only: true,
    fixtures: {},
    command_overrides: {},
    skipped_tools: []
  }

  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/merge_readiness_packet.rb [--format json|markdown] [--offline-only] [fixtures]"
    parser.on("--format FORMAT", "Output format: json or markdown") { |value| options[:format] = value }
    parser.on("--offline-only", "Keep Docker, QEMU, network, package-manager, staging, tagging, and host mutation disabled") { options[:offline_only] = true }
    parser.on("--skip-tool TOOL", "Mark one input tool as skipped") { |value| options[:skipped_tools] << value }
    parser.on("--tool-command TOOL=COMMAND", "Override one tool command for tests or review rehearsal") do |value|
      tool, command = value.split("=", 2)
      raise PacketError, "--tool-command requires TOOL=COMMAND" if tool.to_s.empty? || command.to_s.empty?

      options[:command_overrides][tool] = Shellwords.split(command)
    end
    parser.on("--layout-report PATH", "Use an existing layout verifier text report") { |value| options[:fixtures][:layout_report] = value }
    parser.on("--implementation-report PATH", "Use an existing implementation evidence JSON report") { |value| options[:fixtures][:implementation_report] = value }
    parser.on("--contract-drift-report PATH", "Use an existing Runtime contract drift JSON report") { |value| options[:fixtures][:contract_drift_report] = value }
    parser.on("--kde-smoke-report PATH", "Use an existing KDE-first presence smoke JSON report") { |value| options[:fixtures][:kde_smoke_report] = value }
    parser.on("--mainline-review PATH", "Use an existing mainline integration review JSON report") { |value| options[:fixtures][:mainline_review] = value }
    parser.on("--release-evidence-index PATH", "Use an existing release evidence index JSON report") { |value| options[:fixtures][:release_evidence_index] = value }
    parser.on("--full-checkpoint-promotion PATH", "Use an existing full checkpoint promotion packet JSON report") { |value| options[:fixtures][:full_checkpoint_promotion] = value }
    parser.on("--desktop-trigger-request-preflight-smoke PATH", "Use an existing desktop-trigger request preflight smoke JSON report") { |value| options[:fixtures][:desktop_trigger_request_preflight_smoke] = value }
    parser.on("--q4-sample-notepad-smoke PATH", "Use an existing q4 Sample Notepad acceptance smoke JSON report") { |value| options[:fixtures][:q4_sample_notepad_smoke] = value }
    parser.on("--q4-messagebox-smoke PATH", "Use an existing q4 MessageBox external Windows app document smoke JSON report") { |value| options[:fixtures][:q4_messagebox_smoke] = value }
    parser.on("--fixture-matrix-report PATH", "Use an existing offline fixture matrix JSON report") { |value| options[:fixtures][:fixture_matrix_report] = value }
  end.parse!(argv)

  raise PacketError, "merge_readiness_packet supports --format json or --format markdown" unless %w[json markdown].include?(options[:format])

  options
end

def fixture_path_for(tool_id, options)
  fixture_key = FIXTURE_OPTIONS.fetch(tool_id)
  options[:fixtures][fixture_key]
end

def command_for(tool_id, definition, options)
  options[:command_overrides].fetch(tool_id, definition.fetch(:command))
end

def safe_command_string(command)
  command.join(" ")
end

def load_tool(tool_id, definition, options)
  command = command_for(tool_id, definition, options)
  base = {
    "id" => tool_id,
    "title" => definition.fetch(:title),
    "required" => definition.fetch(:required),
    "parser" => definition.fetch(:parser),
    "command" => safe_command_string(command),
    "fixture_used" => false,
    "exit_status" => nil,
    "status" => "pending",
    "summary" => "",
    "error" => nil,
    "data" => nil
  }

  if options[:skipped_tools].include?(tool_id)
    return base.merge(
      "status" => "skipped",
      "summary" => "Tool was skipped by request; no command was executed.",
      "error" => "tool-skipped"
    )
  end

  fixture_path = fixture_path_for(tool_id, options)
  if fixture_path
    text = Pathname.new(fixture_path).read
    return parse_tool_output(base.merge("fixture_used" => true, "command" => "fixture:#{tool_id}"), text)
  end

  if definition.fetch(:fixture_only, false)
    return base.merge(
      "status" => "skipped",
      "summary" => "Optional evidence was not supplied; no command was executed.",
      "error" => "fixture-not-supplied"
    )
  end

  stdout, stderr, status = Open3.capture3(*command, chdir: PROJECT_ROOT.to_s)
  return parse_tool_output(base.merge("exit_status" => status.exitstatus), stdout) if status.success?

  base.merge(
    "status" => "fail",
    "exit_status" => status.exitstatus,
    "summary" => "Tool command failed.",
    "error" => stderr.strip.empty? ? "command-failed" : stderr.lines.first.to_s.strip
  )
rescue Errno::ENOENT
  base.merge(
    "status" => "missing-command",
    "summary" => "Tool command could not be found.",
    "error" => "missing-command"
  )
rescue JSON::ParserError => e
  base.merge(
    "status" => "malformed-json",
    "summary" => "Tool emitted malformed JSON.",
    "error" => e.message
  )
end

def parse_tool_output(base, stdout)
  if base.fetch("parser") == "json"
    data = JSON.parse(stdout)
    return base.merge(
      "status" => "pass",
      "summary" => json_tool_summary(base.fetch("id"), data),
      "data" => data
    )
  end

  passed = stdout.include?("PASS:")
  base.merge(
    "status" => passed ? "pass" : "fail",
    "summary" => passed ? stdout.lines.find { |line| line.include?("PASS:") }.to_s.strip : "Layout verifier did not emit PASS.",
    "data" => { "text_summary" => stdout.lines.first(3).map(&:strip).reject(&:empty?) }
  )
end

def json_tool_summary(tool_id, data)
  case tool_id
  when "implementation"
    domains = data.fetch("domains", [])
    "#{domains.length} implementation domain(s) reported."
  when "contract_drift"
    data.fetch("drift_detected", false) ? "Runtime contract drift reported." : "No Runtime contract drift reported."
  when "kde_smoke"
    "#{data.fetch("entrypoint_count", 0)} KDE entry point(s) reported."
  when "mainline_review"
    "#{data.fetch("changed_file_count", 0)} changed file(s), #{data.fetch("unclassified_file_count", 0)} unclassified."
  when "release_evidence"
    "#{data.fetch("claim_count", data.fetch("claims", []).length)} release evidence claim(s) reported."
  when "full_checkpoint_promotion"
    "Full checkpoint promotion decision: #{data.fetch("promotion_decision", "unknown")}."
  when "desktop_trigger_request_preflight_smoke"
    "Desktop-trigger request preflight smoke state: #{data.fetch("preflight_smoke_state", "unknown")}."
  when "q4_sample_notepad_smoke"
    "q4 Sample Notepad acceptance state: #{data.fetch("status", "unknown")}."
  when "q4_messagebox_smoke"
    "q4 MessageBox document smoke state: #{data.fetch("status", "unknown")}."
  when "offline_fixture_matrix"
    counts = data.fetch("counts", {})
    row_count = data.fetch("row_count", data.fetch("rows", []).length)
    blocked_count = counts.fetch("blocked", counts.fetch("unsupported", 0))
    "#{row_count} fixture row(s), #{blocked_count} blocked or unsupported."
  else
    "JSON report parsed."
  end
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

def unsafe_findings(tools)
  findings = []
  tools.each do |tool|
    data = tool.fetch("data")
    next unless data

    deep_each(data) do |path, value|
      key = path.last.to_s
      next unless UNSAFE_KEYS.include?(key) && value == true

      findings << {
        "tool" => tool.fetch("id"),
        "key" => key,
        "path" => path.join("."),
        "reason" => "unsafe-operation-flag"
      }
    end
  end
  findings
end

def tool_blockers(tools)
  tools.filter_map do |tool|
    next if tool.fetch("status") == "pass"
    next if !tool.fetch("required") && tool.fetch("status") == "skipped"
    next if tool.fetch("id") == "desktop_trigger_request_preflight_smoke"

    "#{tool.fetch("id")}:#{tool.fetch("status")}"
  end
end

def mainline_data(tools)
  tool_data(tools, "mainline_review")
end

def tool_data(tools, id)
  tool = tools.find { |candidate| candidate.fetch("id") == id }
  tool && tool.fetch("data")
end

def changed_file_counts(mainline)
  lanes = mainline.fetch("lanes", [])
  classified_count = lanes.reject { |lane| %w[unclassified blocked-protected-claude-owned-file].include?(lane.fetch("id")) }.sum { |lane| lane.fetch("file_count", 0).to_i }
  {
    "total" => mainline.fetch("changed_file_count", 0).to_i,
    "classified" => classified_count,
    "unclassified" => mainline.fetch("unclassified_file_count", 0).to_i,
    "protected" => mainline.fetch("protected_claude_file_modified", false) ? 1 : 0,
    "risky" => mainline.fetch("risky_file_count", 0).to_i,
    "code" => mainline.fetch("code_change_count", 0).to_i
  }
end

def lane_classification(mainline)
  mainline.fetch("lanes", []).map do |lane|
    {
      "id" => lane.fetch("id"),
      "workstream" => lane.fetch("workstream"),
      "title" => lane.fetch("title"),
      "file_count" => lane.fetch("file_count", 0),
      "statuses" => lane.fetch("statuses", []),
      "required_verification" => lane.fetch("required_verification", []),
      "safety_guards" => lane.fetch("safety_guards", [])
    }
  end
end

def required_follow_up_commands(mainline)
  lane_classification(mainline)
    .flat_map { |lane| lane.fetch("required_verification") }
    .uniq
    .sort
end

def authorized_product_smoke_complete?(release_evidence)
  claim = release_evidence.fetch("claims", []).find { |item| item.fetch("id", "") == "product-image-qemu-acceptance" }
  claim && claim.fetch("evidence_level", "") == "implemented" && claim.fetch("blockers", []).empty?
end

def checkpoint_promotion_allowed?(full_checkpoint_promotion)
  full_checkpoint_promotion.fetch("promotion_allowed", false) == true &&
    full_checkpoint_promotion.fetch("formal_release_ready", false) == true
end

def desktop_trigger_request_preflight_smoke_passed?(tools)
  tool = tools.find { |candidate| candidate.fetch("id") == "desktop_trigger_request_preflight_smoke" }
  return nil unless tool
  return nil if tool.fetch("status") == "skipped"
  return false unless tool.fetch("status") == "pass"

  data = tool.fetch("data") || {}
  data.fetch("smoke_passed", false) == true &&
    data.fetch("preflight_smoke_state", "") == "passed" &&
    data.fetch("blocked_preflight_state", "") == "blocked-missing-promotion" &&
    data.fetch("ready_preflight_state", "") == "ready-for-operator-request" &&
    data.fetch("owner_service_call_shape_verified", false) == true &&
    data.fetch("operator_request_ready", false) == true &&
    data.fetch("service_call_dispatched", true) == false &&
    data.fetch("dbus_called", true) == false &&
    data.fetch("backend_launch_enabled", true) == false &&
    data.fetch("host_root_modified", true) == false
end

def q4_sample_notepad_smoke_passed?(tools)
  tool = tools.find { |candidate| candidate.fetch("id") == "q4_sample_notepad_smoke" }
  return nil unless tool
  return nil if tool.fetch("status") == "skipped"
  return false unless tool.fetch("status") == "pass"

  data = tool.fetch("data") || {}
  data.fetch("schema_version", "") == "xnix.scripts.q4_sample_notepad_smoke.v1" &&
    data.fetch("request_type", "") == "q4-sample-notepad-smoke" &&
    data.fetch("status", "") == "passed" &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_schema", "") == "xnix.runtime.q4_sample_notepad_acceptance.v1" &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_request_type", "") == "q4-sample-notepad-acceptance-preview" &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_ready", false) == true &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_consumed", false) == true &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_path_exposed", true) == false &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_remote_host_exposed", true) == false &&
    data.fetch("go_owned_q4_sample_notepad_acceptance_delegated_command_exposed", true) == false &&
    data.fetch("host_root_modified", true) == false &&
    data.fetch("privileged_container_required", true) == false &&
    data.fetch("host_networking_required", true) == false &&
    data.fetch("docker_socket_mounted", true) == false &&
    data.fetch("broad_host_mount_required", true) == false
end

def q4_messagebox_smoke_passed?(tools)
  tool = tools.find { |candidate| candidate.fetch("id") == "q4_messagebox_smoke" }
  return nil unless tool
  return nil if tool.fetch("status") == "skipped"
  return false unless tool.fetch("status") == "pass"

  data = tool.fetch("data") || {}
  data.fetch("schema_version", "") == "xnix.scripts.q4_messagebox_smoke.v1" &&
    data.fetch("request_type", "") == "q4-messagebox-smoke" &&
    data.fetch("status", "") == "passed" &&
    data.fetch("app_id", "") == "org.xnix.apps.messagebox" &&
    data.fetch("window_match", "") == "Xnix document opened by Windows app" &&
    data.fetch("document_content_marker_observation_required", false) == true &&
    data.fetch("document_content_marker_observed", false) == true &&
    data.fetch("real_run_receipt_summary_ready", false) == true &&
    data.fetch("real_run_receipt_summary_file_open_verified", false) == true &&
    data.fetch("real_run_receipt_summary_document_content_marker_observed", false) == true &&
    data.fetch("real_run_acceptance_ready", false) == true &&
    data.fetch("real_run_acceptance_document_content_marker_observed", false) == true &&
    data.fetch("real_run_acceptance_center_projection_consumed", false) == true &&
    data.fetch("real_run_acceptance_kde_page_projection_consumed", false) == true &&
    data.fetch("owner_file_open_entrypoint_invoked", false) == true &&
    data.fetch("runtime_evidence_owner_file_open_entrypoint_invoked", false) == true &&
    data.fetch("go_owned_q4_winapp_acceptance_schema", "") == "xnix.runtime.q4_winapp_acceptance.v1" &&
    data.fetch("go_owned_q4_winapp_acceptance_request_type", "") == "q4-winapp-acceptance-preview" &&
    data.fetch("go_owned_q4_winapp_acceptance_ready", false) == true &&
    data.fetch("go_owned_q4_winapp_acceptance_consumed", false) == true &&
    data.fetch("go_owned_q4_winapp_acceptance_document_content_marker_observed", false) == true &&
    data.fetch("go_owned_q4_winapp_acceptance_path_exposed", true) == false &&
    data.fetch("go_owned_q4_winapp_acceptance_remote_host_exposed", true) == false &&
    data.fetch("go_owned_q4_winapp_acceptance_delegated_command_exposed", true) == false &&
    data.fetch("remote_executable_path_exposed", true) == false &&
    data.fetch("host_compilation_avoided", false) == true &&
    data.fetch("host_root_modified", true) == false &&
    data.fetch("privileged_container_required", true) == false &&
    data.fetch("host_networking_required", true) == false &&
    data.fetch("docker_socket_mounted", true) == false &&
    data.fetch("broad_host_mount_required", true) == false
end

def release_blocking_reasons(tools, mainline, contract_drift, kde_smoke, release_evidence, full_checkpoint_promotion, unsafe_findings, preflight_smoke_passed, q4_sample_notepad_smoke_passed, q4_messagebox_smoke_passed)
  reasons = tool_blockers(tools)
  reasons << "protected-claude-file-modified" if mainline.fetch("protected_claude_file_modified", false)
  reasons << "unclassified-files-present" if mainline.fetch("unclassified_file_count", 0).to_i.positive?
  reasons << "risky-files-present" if mainline.fetch("risky_file_count", 0).to_i.positive?
  reasons << "runtime-contract-drift-detected" if contract_drift.fetch("drift_detected", false)
  reasons << "kde-seven-entrypoint-smoke-incomplete" if kde_smoke.fetch("entrypoint_count", 0).to_i != 7
  reasons << "unsafe-operation-detected" unless unsafe_findings.empty?
  reasons << "restricted-docker-or-qemu-smoke-requires-human-authorization" unless authorized_product_smoke_complete?(release_evidence)
  reasons << "full-checkpoint-promotion-not-allowed" unless checkpoint_promotion_allowed?(full_checkpoint_promotion)
  reasons << "desktop-trigger-request-preflight-smoke-not-passed" if preflight_smoke_passed == false
  reasons << "q4-sample-notepad-acceptance-smoke-not-passed" unless q4_sample_notepad_smoke_passed == true
  reasons << "q4-messagebox-document-content-smoke-not-passed" unless q4_messagebox_smoke_passed == true
  reasons << "production-runtime-and-windows-execution-remain-disabled"
  reasons.uniq
end

def build_packet(options)
  tools = TOOL_DEFINITIONS.map do |tool_id, definition|
    load_tool(tool_id, definition, options)
  end

  mainline = mainline_data(tools) || {}
  contract_drift = tool_data(tools, "contract_drift") || {}
  kde_smoke = tool_data(tools, "kde_smoke") || {}
  release_evidence = tool_data(tools, "release_evidence") || {}
  full_checkpoint_promotion = tool_data(tools, "full_checkpoint_promotion") || {}
  unsafe = unsafe_findings(tools)
  preflight_smoke_passed = desktop_trigger_request_preflight_smoke_passed?(tools)
  q4_sample_notepad_smoke_passed = q4_sample_notepad_smoke_passed?(tools)
  q4_messagebox_smoke_passed = q4_messagebox_smoke_passed?(tools)
  release_blockers = release_blocking_reasons(tools, mainline, contract_drift, kde_smoke, release_evidence, full_checkpoint_promotion, unsafe, preflight_smoke_passed, q4_sample_notepad_smoke_passed, q4_messagebox_smoke_passed)
  merge_blockers = release_blockers - RELEASE_ONLY_BLOCKERS

  {
    "version" => VERSION,
    "schema_version" => "xnix.runtime.merge_readiness_packet.v1",
    "report_type" => "merge-readiness-packet",
    "offline_only" => options.fetch(:offline_only),
    "runtime_owned" => true,
    "ruby_report_only" => true,
    "kde_policy_owner" => false,
    "merge_ready" => merge_blockers.empty?,
    "release_ready" => false,
    "tool_statuses" => tools.map { |tool| tool.reject { |key, _| key == "data" } },
    "tool_status_counts" => tools.group_by { |tool| tool.fetch("status") }.transform_values(&:length),
    "changed_file_counts" => changed_file_counts(mainline),
    "lane_classification" => lane_classification(mainline),
    "protected_file_status" => {
      "protected_file" => "docs/claude-code-implementation-packages.md",
      "modified" => mainline.fetch("protected_claude_file_modified", false),
      "status" => mainline.fetch("protected_claude_file_modified", false) ? "blocked" : "clean"
    },
    "unsafe_operation_status" => {
      "docker_executed" => false,
      "qemu_executed" => false,
      "network_checks_run" => false,
      "package_manager_invoked" => false,
      "backend_launch_enabled" => false,
      "host_root_modified" => false,
      "automatic_staging_enabled" => false,
      "automatic_commit_enabled" => false,
      "automatic_release_tagging_enabled" => false,
      "automatic_push_enabled" => false,
      "detected_unsafe_findings" => unsafe
    },
    "required_follow_up_commands" => required_follow_up_commands(mainline),
    "merge_blocking_reasons" => merge_blockers,
    "release_blocking_reasons" => release_blockers,
    "full_checkpoint_promotion_status" => {
      "promotion_allowed" => checkpoint_promotion_allowed?(full_checkpoint_promotion),
      "promotion_decision" => full_checkpoint_promotion.fetch("promotion_decision", "missing"),
      "full_smoke_state" => full_checkpoint_promotion.fetch("full_smoke_state", "missing"),
      "formal_release_ready" => full_checkpoint_promotion.fetch("formal_release_ready", false),
      "operator_required_command" => full_checkpoint_promotion.fetch("operator_required_command", "ruby scripts/full_smoke.rb")
    },
    "desktop_trigger_request_preflight_smoke_status" => {
      "evidence_supplied" => !preflight_smoke_passed.nil?,
      "smoke_passed" => preflight_smoke_passed == true,
      "status" => preflight_smoke_passed.nil? ? "not-supplied" : (preflight_smoke_passed ? "passed" : "blocked"),
      "release_blocking_reason" => preflight_smoke_passed == false ? "desktop-trigger-request-preflight-smoke-not-passed" : nil
    },
    "q4_sample_notepad_smoke_status" => {
      "evidence_supplied" => !q4_sample_notepad_smoke_passed.nil?,
      "smoke_passed" => q4_sample_notepad_smoke_passed == true,
      "status" => q4_sample_notepad_smoke_passed.nil? ? "not-supplied" : (q4_sample_notepad_smoke_passed ? "passed" : "blocked"),
      "release_blocking_reason" => q4_sample_notepad_smoke_passed == true ? nil : "q4-sample-notepad-acceptance-smoke-not-passed"
    },
    "q4_messagebox_smoke_status" => {
      "evidence_supplied" => !q4_messagebox_smoke_passed.nil?,
      "smoke_passed" => q4_messagebox_smoke_passed == true,
      "status" => q4_messagebox_smoke_passed.nil? ? "not-supplied" : (q4_messagebox_smoke_passed ? "passed" : "blocked"),
      "release_blocking_reason" => q4_messagebox_smoke_passed == true ? nil : "q4-messagebox-document-content-smoke-not-passed"
    },
    "desktop_safe_summary" => "Merge readiness is aggregated offline from local reports; staging, committing, tagging, pushing, Docker, QEMU, network fetch, package managers, backend launch, and host-root mutation remain disabled."
  }
end

def render_markdown(packet)
  lines = []
  lines << "# Merge Readiness Packet"
  lines << ""
  lines << "- Version: #{packet.fetch("version")}"
  lines << "- Merge ready: #{packet.fetch("merge_ready")}"
  lines << "- Release ready: #{packet.fetch("release_ready")}"
  lines << "- Offline only: #{packet.fetch("offline_only")}"
  lines << "- Protected Claude file: #{packet.fetch("protected_file_status").fetch("status")}"
  lines << "- Full checkpoint promotion: #{packet.fetch("full_checkpoint_promotion_status").fetch("promotion_decision")}"
  lines << "- Desktop-trigger request preflight smoke: #{packet.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status")}"
  lines << "- q4 Sample Notepad acceptance smoke: #{packet.fetch("q4_sample_notepad_smoke_status").fetch("status")}"
  lines << "- q4 MessageBox document smoke: #{packet.fetch("q4_messagebox_smoke_status").fetch("status")}"
  lines << ""
  lines << "## Tool Statuses"
  lines << ""
  lines << "| Tool | Status | Command | Summary |"
  lines << "| --- | --- | --- | --- |"
  packet.fetch("tool_statuses").each do |tool|
    lines << "| #{tool.fetch("id")} | #{tool.fetch("status")} | `#{tool.fetch("command")}` | #{tool.fetch("summary")} |"
  end
  lines << ""
  lines << "## Changed Files"
  lines << ""
  packet.fetch("changed_file_counts").each do |key, value|
    lines << "- #{key}: #{value}"
  end
  lines << ""
  lines << "## Lane Classification"
  lines << ""
  packet.fetch("lane_classification").each do |lane|
    next if lane.fetch("file_count").to_i.zero?

    lines << "- #{lane.fetch("id")}: #{lane.fetch("file_count")} file(s)"
  end
  lines << ""
  lines << "## Blocking Reasons"
  lines << ""
  if packet.fetch("release_blocking_reasons").empty?
    lines << "- None"
  else
    packet.fetch("release_blocking_reasons").each { |reason| lines << "- #{reason}" }
  end
  lines << ""
  lines << "## Required Follow-up Commands"
  lines << ""
  packet.fetch("required_follow_up_commands").each { |command| lines << "- `#{command}`" }
  lines.join("\n")
end

begin
  options = parse_options(ARGV)
  packet = build_packet(options)
  case options.fetch(:format)
  when "json"
    puts JSON.pretty_generate(packet)
  when "markdown"
    puts render_markdown(packet)
  end
rescue PacketError => e
  abort e.message
end
