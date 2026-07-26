# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

ROOT = Pathname.new(__dir__).join("..").realpath
SCRIPT = ROOT.join("scripts", "remote_known_winapp_guest_wine_smoke.rb")

def assert(condition, message)
  raise message unless condition
end

script = SCRIPT.read

assert(script.include?("\"windows-known-app-run\""), "remote smoke must call the unified known app run entrypoint")
assert(script.include?("\"--backend\", \"guest-wine\""), "remote smoke must select the guest-wine backend")
assert(script.include?("\"--start-qemu\""), "remote smoke must let the Go Runtime start QEMU")
assert(script.include?("\"--port\", \"auto\""), "remote smoke must avoid a fixed host SSH port")
assert(script.include?("\"--redact-output\""), "remote smoke must keep real app output product-safe")
assert(script.include?("\"--report-output\", remote_report"), "remote smoke must persist JSON evidence through the Runtime")
assert(script.include?("\"--qemu-serial-log\", remote_serial_log"), "remote smoke must persist a QEMU serial log")
assert(script.include?("DEFAULT_REMOTE_GO"), "remote smoke must use an explicit remote Go toolchain")
assert(script.include?("DEFAULT_LOCAL_SHELL"), "remote smoke must use an explicit local shell for ssh alias resolution")
assert(script.include?("DEFAULT_SOURCE_SYNC_MODE"), "remote smoke must expose an explicit source sync mode")
assert(script.include?("%w[VERSION go.mod cmd internal runtime]"), "remote smoke must default to the lightweight Runtime source set")
assert(script.include?("run_shell(options.fetch(:local_shell), shell_join(rsync_args))"), "remote smoke source sync must use the configured local shell")
assert(script.include?("run_shell(options.fetch(:local_shell), shell_join([\"ssh\", remote_host, remote_command]))"), "remote smoke execution must use the configured local shell")
assert(script.include?("\"known-existing-winapp-acceptance-preview\""), "remote smoke must be able to emit Go-owned acceptance JSON")
assert(script.include?("\"known-app-verified-catalog-run-acceptance-preview\""), "remote smoke must be able to emit Go-owned verified catalog acceptance JSON")
assert(script.include?("\"--known-winapp-run\", remote_report"), "remote smoke acceptance must consume the persisted known app run report")
assert(script.include?("\"--run-plan\", remote_run_plan"), "remote smoke verified catalog acceptance must consume a q4-synced run plan")
assert(script.include?("--verified-catalog-run-plan"), "remote smoke must expose a verified catalog run-plan flag")
assert(script.include?("--verified-catalog-acceptance-json"), "remote smoke must expose a verified catalog acceptance mode")
assert(script.include?("--desktop-consumption-json"), "remote smoke must expose a desktop consumption packet mode")
assert(script.include?("\"compatibility-center-preview\""), "remote smoke desktop packet must consume acceptance through Compatibility Center")
assert(script.include?("\"kde-center-page-preview\""), "remote smoke desktop packet must consume acceptance through KDE Center page")
assert(script.include?("remote-known-winapp-verified-catalog-desktop-consumption"), "remote smoke desktop packet must expose a stable request type")
assert(script.include?("/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go"), "remote smoke must default to the q4 private Go toolchain")
assert(script.include?("ensure_remote_xnix_path!"), "remote smoke must restrict writable remote roots")
assert(script.include?("\"/home/xnix-run-materials\""), "remote smoke must use the managed q4 run materials root")
assert(script.include?("\"/home/xnix-build-cache\""), "remote smoke must use the managed q4 build cache root")
assert(script.include?("xnix-runtime-source-\#{DEFAULT_SOURCE_SYNC_MODE}-\#{VERSION}"), "remote smoke must default to a versioned lightweight source root")
assert(script.include?("\"--exclude\", \"docs/claude-code-implementation-packages.md\""), "remote smoke source sync must not touch the protected Claude implementation package document")
assert(!script.include?("\"windows-known-app-dispatch-smoke\""), "remote smoke must not bypass the unified known app run entrypoint")
assert(!script.include?("\"--guest-boundary\""), "remote smoke must not rely on the older dispatch boundary flag")

stdout, stderr, status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, chdir: ROOT.to_s)
abort stderr unless status.success?

payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.remote_known_windows_app_guest_wine_smoke.v1", "dry-run plan must expose the script schema")
assert(payload.fetch("request_type") == "remote-known-winapp-guest-wine-smoke", "dry-run plan must expose the request type")
assert(payload.fetch("status") == "planned", "dry-run plan must not execute by default")
assert(payload.fetch("execute") == false, "dry-run plan must require explicit execution")
assert(payload.fetch("local_shell") == "/bin/zsh", "dry-run plan must expose the local shell used for remote command resolution")
assert(payload.fetch("source_sync_planned") == true, "dry-run plan must sync source by default")
assert(payload.fetch("source_sync_mode") == "runtime", "dry-run plan must default to lightweight Runtime source sync")
assert(payload.fetch("source_sync_entry_count") == 5, "dry-run plan must expose the lightweight source entry count")
assert(payload.fetch("source_sync_entries") == %w[VERSION go.mod cmd internal runtime], "dry-run plan must expose the lightweight source entries")
assert(payload.fetch("remote_source_root").include?("xnix-runtime-source-runtime-"), "dry-run plan must use a versioned lightweight source root")
assert(payload.fetch("acceptance_json_planned") == false, "dry-run plan must keep acceptance JSON disabled by default")
assert(payload.fetch("verified_catalog_acceptance_json_planned") == false, "dry-run plan must keep verified catalog acceptance JSON disabled by default")
assert(payload.fetch("verified_catalog_run_plan_sync_planned") == false, "dry-run plan must not sync a run plan by default")
assert(payload.fetch("desktop_consumption_json_planned") == false, "dry-run plan must keep desktop consumption disabled by default")
assert(payload.fetch("desktop_consumption_request_type") == "remote-known-winapp-verified-catalog-desktop-consumption", "dry-run plan must name the desktop consumption request type")
assert(payload.fetch("desktop_app_id") == "org.xnix.sample.notepad", "dry-run plan must expose the default KDE page host app")
assert(payload.fetch("acceptance_request_type") == "known-existing-winapp-acceptance-preview", "dry-run plan must name the Go acceptance request type")
assert(payload.fetch("verified_catalog_acceptance_request_type") == "known-app-verified-catalog-run-acceptance-preview", "dry-run plan must name the verified catalog acceptance request type")
assert(payload.fetch("backend") == "guest-wine", "dry-run plan must select the guest backend")
assert(payload.fetch("start_qemu") == true, "dry-run plan must use Runtime-owned QEMU startup")
assert(payload.fetch("guest_port") == "auto", "dry-run plan must use automatic loopback port allocation")
assert(payload.fetch("redact_output") == true, "dry-run plan must redact raw app output")
assert(payload.fetch("host_root_modified") == false, "dry-run plan must not claim host-root mutation")
assert(payload.fetch("privileged_container_required") == false, "dry-run plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "dry-run plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "dry-run plan must not mount a Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "dry-run plan must not require broad host mounts")

full_stdout, full_stderr, full_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--source-sync-mode", "full", chdir: ROOT.to_s)
abort full_stderr unless full_status.success?

full_payload = JSON.parse(full_stdout)
assert(full_payload.fetch("source_sync_mode") == "full", "full dry-run plan must expose full source sync mode")
assert(full_payload.fetch("source_sync_entries") == ["."], "full dry-run plan must sync the full checkout when requested")
assert(full_payload.fetch("remote_source_root").include?("xnix-runtime-source-full-"), "full dry-run plan must use a versioned full source root")

acceptance_stdout, acceptance_stderr, acceptance_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--acceptance-json", chdir: ROOT.to_s)
abort acceptance_stderr unless acceptance_status.success?

acceptance_payload = JSON.parse(acceptance_stdout)
assert(acceptance_payload.fetch("status") == "planned", "acceptance dry-run plan must not execute")
assert(acceptance_payload.fetch("acceptance_json_planned") == true, "acceptance dry-run plan must expose acceptance JSON mode")
assert(acceptance_payload.fetch("acceptance_request_type") == "known-existing-winapp-acceptance-preview", "acceptance dry-run plan must route to Go-owned acceptance")
assert(acceptance_payload.fetch("host_root_modified") == false, "acceptance dry-run plan must keep host-root mutation closed")

run_plan = ROOT.join("test", "fixtures", "known_app_verified_catalog_run_plan_7zr.json")
verified_stdout, verified_stderr, verified_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--verified-catalog-run-plan", run_plan.to_s, "--verified-catalog-acceptance-json", chdir: ROOT.to_s)
abort verified_stderr unless verified_status.success?

verified_payload = JSON.parse(verified_stdout)
assert(verified_payload.fetch("status") == "planned", "verified catalog dry-run plan must not execute")
assert(verified_payload.fetch("acceptance_json_planned") == false, "verified catalog dry-run plan must not use the generic acceptance mode")
assert(verified_payload.fetch("verified_catalog_acceptance_json_planned") == true, "verified catalog dry-run plan must expose verified catalog acceptance JSON mode")
assert(verified_payload.fetch("verified_catalog_run_plan_sync_planned") == true, "verified catalog dry-run plan must sync a run plan")
assert(verified_payload.fetch("desktop_consumption_json_planned") == false, "verified catalog dry-run plan must keep desktop consumption disabled unless requested")
assert(verified_payload.fetch("verified_catalog_run_plan_remote_path").include?("/home/xnix-run-materials/state/known-run-plan-"), "verified catalog dry-run plan must use the managed q4 state directory for the run plan")
assert(verified_payload.fetch("verified_catalog_acceptance_request_type") == "known-app-verified-catalog-run-acceptance-preview", "verified catalog dry-run plan must route to Go-owned verified catalog acceptance")
assert(verified_payload.fetch("host_root_modified") == false, "verified catalog dry-run plan must keep host-root mutation closed")

desktop_stdout, desktop_stderr, desktop_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--verified-catalog-run-plan", run_plan.to_s, "--verified-catalog-acceptance-json", "--desktop-consumption-json", chdir: ROOT.to_s)
abort desktop_stderr unless desktop_status.success?

desktop_payload = JSON.parse(desktop_stdout)
assert(desktop_payload.fetch("status") == "planned", "desktop consumption dry-run plan must not execute")
assert(desktop_payload.fetch("verified_catalog_acceptance_json_planned") == true, "desktop consumption dry-run plan must include verified catalog acceptance")
assert(desktop_payload.fetch("desktop_consumption_json_planned") == true, "desktop consumption dry-run plan must expose desktop consumption mode")
assert(desktop_payload.fetch("desktop_consumption_request_type") == "remote-known-winapp-verified-catalog-desktop-consumption", "desktop consumption dry-run plan must route to the stable packet")
assert(desktop_payload.fetch("desktop_app_id") == "org.xnix.sample.notepad", "desktop consumption dry-run plan must expose the KDE page host app")
assert(desktop_payload.fetch("host_root_modified") == false, "desktop consumption dry-run plan must keep host-root mutation closed")

desktop_without_verified_stdout, desktop_without_verified_stderr, desktop_without_verified_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--desktop-consumption-json", chdir: ROOT.to_s)
assert(!desktop_without_verified_status.success?, "remote smoke must reject desktop consumption without verified catalog acceptance")
assert(desktop_without_verified_stderr.include?("requires --verified-catalog-acceptance-json") || desktop_without_verified_stdout.include?("requires --verified-catalog-acceptance-json"), "desktop consumption dependency error must be explicit")

mixed_stdout, mixed_stderr, mixed_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--acceptance-json", "--verified-catalog-run-plan", run_plan.to_s, "--verified-catalog-acceptance-json", chdir: ROOT.to_s)
assert(!mixed_status.success?, "remote smoke must reject mixed acceptance JSON modes")
assert(mixed_stderr.include?("accepts only one acceptance JSON mode") || mixed_stdout.include?("accepts only one acceptance JSON mode"), "mixed acceptance JSON mode error must be explicit")

puts "PASS: remote known Windows app guest Wine smoke script is execute-gated and Runtime-owned"
