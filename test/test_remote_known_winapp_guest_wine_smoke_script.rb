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
assert(script.include?("%w[go.mod cmd internal runtime]"), "remote smoke must default to the lightweight Runtime source set")
assert(script.include?("run_shell(options.fetch(:local_shell), shell_join(rsync_args))"), "remote smoke source sync must use the configured local shell")
assert(script.include?("run_shell(options.fetch(:local_shell), shell_join([\"ssh\", remote_host, remote_command]))"), "remote smoke execution must use the configured local shell")
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
assert(payload.fetch("source_sync_entry_count") == 4, "dry-run plan must expose the lightweight source entry count")
assert(payload.fetch("source_sync_entries") == %w[go.mod cmd internal runtime], "dry-run plan must expose the lightweight source entries")
assert(payload.fetch("remote_source_root").include?("xnix-runtime-source-runtime-"), "dry-run plan must use a versioned lightweight source root")
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

puts "PASS: remote known Windows app guest Wine smoke script is execute-gated and Runtime-owned"
