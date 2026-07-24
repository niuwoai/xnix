# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

ROOT = Pathname.new(__dir__).join("..").realpath
SCRIPT = ROOT.join("scripts", "remote_known_winapp_matrix_smoke.rb")

def assert(condition, message)
  raise message unless condition
end

script = SCRIPT.read

assert(script.include?("\"windows-known-app-run\""), "matrix smoke must call the unified known app run entrypoint")
assert(script.include?("\"--backend\", \"guest-wine\""), "matrix smoke must select the guest-wine backend")
assert(script.include?("\"--start-qemu\""), "matrix smoke must let the Go Runtime start QEMU")
assert(script.include?("\"--port\", \"auto\""), "matrix smoke must avoid fixed host SSH ports")
assert(script.include?("\"--redact-output\""), "matrix smoke must keep real app output product-safe")
assert(script.include?("\"--report-output\", app_plan.fetch(\"report_output\")"), "matrix smoke must persist per-app JSON evidence")
assert(script.include?("\"--qemu-serial-log\", app_plan.fetch(\"serial_log_output\")"), "matrix smoke must persist per-app serial logs")
assert(script.include?("DEFAULT_MATRIX_REPORT_OUTPUT"), "matrix smoke must define a default aggregate report output")
assert(script.include?("\"matrix_report_output_written\" => true"), "matrix smoke must mark successful aggregate report persistence")
assert(script.include?("\"scp\", file.path"), "matrix smoke must upload the aggregate matrix report to the remote state directory")
assert(script.include?("DEFAULT_APP_IDS"), "matrix smoke must define default app ids")
assert(script.include?("\"7zr,busybox-w32\""), "matrix smoke must default to the first two real known apps")
assert(script.include?("%w[go.mod cmd internal runtime]"), "matrix smoke must default to the lightweight Runtime source set")
assert(script.include?("docs/claude-code-implementation-packages.md"), "matrix smoke sync must exclude the protected Claude implementation package document")
assert(script.include?("passed_count"), "matrix smoke must report aggregate pass counts")
assert(script.include?("failed_count"), "matrix smoke must report aggregate failure counts")
assert(!script.include?("\"windows-known-app-dispatch-smoke\""), "matrix smoke must not bypass the unified known app run entrypoint")
assert(!script.include?("\"--guest-boundary\""), "matrix smoke must not rely on the older dispatch boundary flag")

stdout, stderr, status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, chdir: ROOT.to_s)
abort stderr unless status.success?

payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.remote_known_windows_app_matrix_smoke.v1", "dry-run matrix must expose the script schema")
assert(payload.fetch("request_type") == "remote-known-winapp-matrix-smoke", "dry-run matrix must expose the request type")
assert(payload.fetch("status") == "planned", "dry-run matrix must not execute by default")
assert(payload.fetch("execute") == false, "dry-run matrix must require explicit execution")
assert(payload.fetch("source_sync_mode") == "runtime", "dry-run matrix must default to lightweight Runtime source sync")
assert(payload.fetch("source_sync_entries") == %w[go.mod cmd internal runtime], "dry-run matrix must expose lightweight source entries")
assert(payload.fetch("remote_source_root").include?("xnix-runtime-source-matrix-runtime-"), "dry-run matrix must use a versioned matrix source root")
assert(payload.fetch("matrix_report_output").include?("known-run-matrix-"), "dry-run matrix must expose the aggregate matrix report path")
assert(payload.fetch("matrix_report_output_written") == false, "dry-run matrix must not claim aggregate report persistence")
assert(payload.fetch("app_count") == 2, "dry-run matrix must include both default apps")
assert(payload.fetch("app_ids") == %w[7zr busybox-w32], "dry-run matrix must default to 7zr and BusyBox-w32")
assert(payload.fetch("backend") == "guest-wine", "dry-run matrix must select the guest backend")
assert(payload.fetch("start_qemu") == true, "dry-run matrix must use Runtime-owned QEMU startup")
assert(payload.fetch("guest_port") == "auto", "dry-run matrix must use automatic loopback port allocation")
assert(payload.fetch("redact_output") == true, "dry-run matrix must redact raw app output")
assert(payload.fetch("host_root_modified") == false, "dry-run matrix must not claim host-root mutation")
assert(payload.fetch("privileged_container_required") == false, "dry-run matrix must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "dry-run matrix must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "dry-run matrix must not mount a Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "dry-run matrix must not require broad host mounts")

apps = payload.fetch("apps")
assert(apps.map { |app| app.fetch("app_id") } == %w[7zr busybox-w32], "dry-run matrix must expose per-app plans")
apps.each do |app|
  assert(app.fetch("report_output").include?("known-run-"), "each app must have a report output path")
  assert(app.fetch("serial_log_output").include?("qemu-serial-"), "each app must have a serial log path")
end

single_stdout, single_stderr, single_status = Open3.capture3({ "XNIX_LOCAL_SHELL" => "/bin/zsh" }, "ruby", SCRIPT.to_s, "--app", "busybox-w32", chdir: ROOT.to_s)
abort single_stderr unless single_status.success?

single_payload = JSON.parse(single_stdout)
assert(single_payload.fetch("app_ids") == ["busybox-w32"], "explicit --app must replace the default matrix")
assert(single_payload.fetch("app_count") == 1, "explicit --app matrix must use the requested app count")

custom_stdout, custom_stderr, custom_status = Open3.capture3(
  { "XNIX_LOCAL_SHELL" => "/bin/zsh" },
  "ruby",
  SCRIPT.to_s,
  "--remote-materials-root",
  "/home/xnix-custom-materials",
  chdir: ROOT.to_s
)
abort custom_stderr unless custom_status.success?

custom_payload = JSON.parse(custom_stdout)
assert(custom_payload.fetch("matrix_report_output").start_with?("/home/xnix-custom-materials/state/known-run-matrix-"), "default aggregate report output must follow the selected materials root")

puts "PASS: remote known Windows app matrix smoke script is execute-gated and Runtime-owned"
