#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/q4_known_portable_winapp_run.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_known_portable_winapp_run.v1"), "known portable run must expose a stable schema")
assert(source.include?("q4-known-portable-winapp-run"), "known portable run must expose a stable request type")
assert(source.include?("org.xnix.external.notepadplusplus"), "known portable run must target the Notepad++ catalog app")
assert(source.include?("scripts/q4_notepadpp_portable_winapp_smoke.rb"), "known portable run must delegate to the proven q4 Notepad++ path")
assert(source.include?("q4-known-portable-winapp-run-plan-preview"), "known portable run must advertise its Go Runtime plan")
assert(source.include?("go_runtime_plan_catalog_consumed"), "known portable run must consume Go Runtime catalog-backed plan metadata")
assert(source.include?("supported_known_portable_app_ids"), "known portable run must preserve Go Runtime supported portable app ids")
assert(source.include?("catalog_artifact_kind"), "known portable run must preserve catalog artifact kind")
assert(source.include?("download_artifact_name"), "known portable run must preserve catalog download artifact name")
assert(source.include?("executable_relative_path"), "known portable run must preserve executable relative path")
assert(source.include?("q4-known-portable-winapp-run-acceptance-preview"), "known portable run must require Go Runtime operator acceptance")
assert(source.include?("known_portable_bundle_acceptance_ready"), "known portable run must require Go-owned acceptance")
assert(source.include?("known_portable_bundle_acceptance_runtime_accepted_chain_verified"), "known portable run must require accepted chain verification")
assert(source.include?("operator_run_acceptance_ready"), "known portable run must record Go-owned operator acceptance readiness")
assert(source.include?("Operator acceptance ready"), "known portable run markdown must include operator acceptance readiness")
assert(source.include?("operator_run_acceptance_artifact_fetched"), "known portable run must fetch operator acceptance artifact")
assert(source.include?("operator_run_application_detail_acceptance_consumed"), "known portable run must generate an operator-accepted application detail")
assert(source.include?("operator_run_kde_page_acceptance_consumed"), "known portable run must generate a KDE page from the operator-accepted detail")
assert(source.include?("--q4-known-portable-winapp-run-acceptance"), "known portable run must pass operator acceptance into application detail")
assert(source.include?("runtime-accepted-real-app-run"), "known portable run must require accepted Runtime/KDE state")
assert(source.include?("write_markdown"), "known portable run must write a Markdown run report")
assert(source.include?("host_compilation_avoided"), "known portable run must avoid host compilation")
assert(source.include?("host_download_avoided"), "known portable run must avoid host download")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--app", "org.xnix.external.notepadplusplus",
  "--output", "/tmp/xnix-known-portable-run-plan.json",
  "--markdown-output", "/tmp/xnix-known-portable-run-plan.md"
)
assert(status.success?, "known portable run plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_known_portable_winapp_run.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-known-portable-winapp-run", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("app_id") == "org.xnix.external.notepadplusplus", "plan must preserve app id")
assert(payload.fetch("display_name") == "Notepad++ Portable", "plan must expose display name")
assert(payload.fetch("known_catalog_app") == true, "plan must identify a known catalog app")
assert(payload.fetch("portable_directory_external_app") == true, "plan must require portable directory handling")
assert(payload.fetch("official_download_required") == true, "plan must require official download")
assert(payload.fetch("pinned_checksum_required") == true, "plan must require pinned checksum")
assert(payload.fetch("go_runtime_plan_required") == true, "plan must require Go Runtime planning")
assert(payload.fetch("go_runtime_plan_request_type") == "q4-known-portable-winapp-run-plan-preview", "plan must advertise Go Runtime request")
assert(payload.fetch("go_runtime_plan_status") == "planned", "plan must mark Go Runtime plan as planned")
assert(payload.fetch("go_runtime_plan_catalog_consumed") == false, "plan mode must not claim Go Runtime catalog consumption before q4 execution")
assert(payload.fetch("operator_run_acceptance_planned") == true, "plan must require operator run acceptance")
assert(payload.fetch("operator_run_acceptance_request_type") == "q4-known-portable-winapp-run-acceptance-preview", "plan must advertise operator acceptance request")
assert(payload.fetch("operator_run_acceptance_status") == "planned", "plan must mark operator acceptance as planned")
assert(payload.fetch("operator_run_application_detail_planned") == true, "plan must generate operator-accepted application detail")
assert(payload.fetch("operator_run_application_detail_status") == "planned", "plan must mark operator detail as planned")
assert(payload.fetch("operator_run_kde_page_planned") == true, "plan must generate operator-accepted KDE page")
assert(payload.fetch("operator_run_kde_page_status") == "planned", "plan must mark operator KDE page as planned")
assert(payload.fetch("delegated_request_type") == "q4-notepadpp-portable-winapp-smoke", "plan must delegate to q4 Notepad++ runner")
assert(payload.fetch("delegated_command").include?("scripts/q4_notepadpp_portable_winapp_smoke.rb"), "plan must include delegated script")
assert(payload.fetch("q4_download_required") == true, "plan must require q4 download")
assert(payload.fetch("q4_extract_required") == true, "plan must require q4 extraction")
assert(payload.fetch("q4_compile_required") == true, "plan must require q4 compile")
assert(payload.fetch("q4_execution_required") == true, "plan must require q4 execution")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("host_download_avoided") == true, "plan must avoid host download")
assert(payload.fetch("full_smoke_required") == false, "plan must remain targeted")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")
assert(Pathname.new("/tmp/xnix-known-portable-run-plan.md").file?, "plan must write markdown output")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--app", "org.example.unsupported"
)
assert(!status.success?, "known portable run must reject unsupported apps")
assert(stderr.include?("unsupported known portable app"), "known portable run must explain unsupported app rejection")

puts "PASS: q4 known portable Windows app run script is Runtime-backed"
