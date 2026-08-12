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
script = project_root.join("scripts/q4_full_smoke.rb")
source = script.read
container_script = project_root.join("scripts/container.rb").read

assert(source.include?("xnix.scripts.q4_full_smoke.v1"), "q4 full smoke must expose a stable schema")
assert(source.include?("root@q4"), "q4 full smoke must default to q4")
assert(source.include?("ruby scripts/full_smoke.rb"), "q4 full smoke must run the existing full smoke remotely")
assert(source.include?("XNIX_ALLOW_NON_COLIMA_DOCKER=1"), "q4 full smoke must opt in to non-Colima Docker only on q4")
assert(source.include?("XNIX_TOOLS_BASE_IMAGE"), "q4 full smoke must pass an explicit tools base image to remote Docker builds")
assert(source.include?("XNIX_WINE_BASE_IMAGE"), "q4 full smoke must pass an explicit cached Wine GUI smoke base image to remote Docker builds")
assert(source.include?("python:3.12-slim"), "q4 full smoke must default to a q4-cached slim base image")
assert(source.include?("XNIX_CONTAINER_MEMORY_LIMIT"), "q4 full smoke must pass an explicit remote container memory limit")
assert(source.include?("XNIX_CONTAINER_CPU_LIMIT"), "q4 full smoke must pass an explicit remote container CPU limit")
assert(source.include?("12g"), "q4 full smoke must default to a q4-sized memory limit")
assert(source.include?("2.0"), "q4 full smoke must default to a q4-sized CPU limit")
assert(source.include?("XNIX_KNOWN_WINAPP_FETCH_TIMEOUT"), "q4 full smoke must pass an explicit known Windows app fetch timeout")
assert(source.include?("300s"), "q4 full smoke must default to an extended known Windows app fetch timeout")
assert(source.include?("XNIX_Q4_FULL_SMOKE_PREFLIGHT_TIMEOUT_SECONDS"), "q4 full smoke must expose a bounded q4 preflight timeout")
assert(source.include?("remote_preflight_script"), "q4 full smoke must run an SSH preflight before source sync")
assert(source.include?("XNIX_Q4_FULL_SMOKE_REMOTE_READY=true"), "q4 full smoke preflight must emit a readiness marker")
assert(source.include?("command -v ruby"), "q4 full smoke preflight must check Ruby availability")
assert(source.include?("command -v rsync"), "q4 full smoke preflight must check rsync availability")
assert(source.include?("command -v docker"), "q4 full smoke preflight must check Docker availability")
assert(source.include?("q4 full smoke remote preflight failed"), "q4 full smoke must fail explicitly when q4 is unavailable before rsync")
assert(source.include?("ensure_docker_image_ref!"), "q4 full smoke must validate the tools base image reference")
assert(source.include?("ensure_docker_limit!"), "q4 full smoke must validate Docker resource limits")
assert(source.include?("ensure_duration!"), "q4 full smoke must validate duration arguments")
assert(source.include?("docs/claude-code-implementation-packages.md"), "q4 full smoke source sync must exclude the protected Claude package document")
assert(source.include?("--delete"), "q4 full smoke must keep the remote full source mirror reproducible")
assert(source.include?("full-smoke-report.json"), "q4 full smoke must fetch the JSON full smoke report")
assert(source.include?("full-smoke-report.md"), "q4 full smoke must fetch the Markdown full smoke report")
assert(source.include?("serial.log"), "q4 full smoke must fetch the serial log")
assert(source.include?("host_compilation_avoided"), "q4 full smoke must make host compilation avoidance observable")
assert(source.include?("q4_compile_required"), "q4 full smoke must make q4 compilation ownership observable")
assert(source.include?("ensure_remote_xnix_path!"), "q4 full smoke must constrain remote writable paths")
assert(source.include?("ensure_local_xnix_path!"), "q4 full smoke must constrain local report paths")
assert(source.include?("BatchMode=yes"), "q4 full smoke must use non-interactive SSH mode")
assert(source.include?("ServerAliveInterval=15"), "q4 full smoke must use SSH keepalive")
assert(source.include?("rsync_ssh_transport"), "q4 full smoke rsync must reuse bounded SSH transport options")
assert(source.include?("source_sync_ssh_transport"), "q4 full smoke must expose source-sync SSH transport policy")
assert(source.include?("source_sync_exit_code"), "q4 full smoke source-sync failures must expose the rsync exit code")
assert(source.include?("q4 full smoke shell operation timed out"), "q4 full smoke must bound shell operations")
assert(source.include?("Xnix::Milestone.full_build_required?"), "q4 full smoke must honor the full-build milestone gate")

assert(container_script.include?("XNIX_ALLOW_NON_COLIMA_DOCKER"), "container CLI must expose the explicit trusted-remote non-Colima Docker override")
assert(container_script.include?("trusted remote build host such as q4"), "container CLI must explain the non-Colima Docker override scope")
assert(container_script.include?("colima_context? || non_colima_docker_allowed?"), "container CLI must keep Colima as the local default while allowing q4 opt-in")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "q4 full smoke plan must succeed: #{stderr}")
payload = JSON.parse(stdout)
assert(payload["schema_version"] == "xnix.scripts.q4_full_smoke.v1", "q4 full smoke must expose schema")
assert(payload["request_type"] == "q4-full-smoke", "q4 full smoke must expose request type")
assert(payload["status"] == "planned", "q4 full smoke must be planned by default")
assert(payload["execute"] == false, "q4 full smoke must not execute without --execute")
assert(payload["remote_host"] == "root@q4", "q4 full smoke must default to q4")
assert(payload["source_sync_mode"] == "full", "q4 full smoke must sync the full checkout")
assert(payload["source_sync_entries"] == ["."], "q4 full smoke must expose the full source sync entry")
assert(payload["source_sync_ssh_transport"] == "BatchMode+ConnectTimeout+ServerAlive", "q4 full smoke must expose bounded rsync SSH transport")
assert(payload["source_sync_excludes"].include?("docs/claude-code-implementation-packages.md"), "q4 full smoke must exclude the protected Claude file")
assert(payload["protected_claude_file_excluded"] == true, "q4 full smoke must make protected-file exclusion observable")
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "q4 full smoke remote source root must stay constrained")
assert(payload["report_root"].start_with?("output/q4-full-smoke-"), "q4 full smoke reports must default under output")
assert(payload["remote_report_files"].include?("output/full-smoke-report.json"), "q4 full smoke must plan JSON report fetch")
assert(payload["remote_report_files"].include?("output/full-smoke-report.md"), "q4 full smoke must plan Markdown report fetch")
assert(payload["remote_report_files"].include?("output/serial.log"), "q4 full smoke must plan serial log fetch")
assert(payload["remote_full_smoke_command"] == "ruby scripts/full_smoke.rb", "q4 full smoke must expose the remote full smoke command")
assert(payload["remote_docker_context_override"] == "XNIX_ALLOW_NON_COLIMA_DOCKER=1", "q4 full smoke must expose the q4 Docker context override")
assert(payload["non_colima_docker_opt_in_env"] == "XNIX_ALLOW_NON_COLIMA_DOCKER", "q4 full smoke must expose the opt-in env")
assert(payload["tools_base_image_env"] == "XNIX_TOOLS_BASE_IMAGE", "q4 full smoke must expose the tools base image env")
assert(payload["tools_base_image"] == "python:3.12-slim", "q4 full smoke must default to the q4 cached slim base image")
assert(payload["q4_cached_tools_base_image_preferred"] == true, "q4 full smoke must prefer q4 cached base images")
assert(payload["wine_base_image_env"] == "XNIX_WINE_BASE_IMAGE", "q4 full smoke must expose the Wine GUI smoke base image env")
assert(payload["wine_base_image"] == "python:3.12-slim", "q4 full smoke must default the Wine GUI smoke base to the q4 cached slim image")
assert(payload["q4_cached_wine_base_image_preferred"] == true, "q4 full smoke must prefer the q4 cached Wine base image")
assert(payload["container_memory_limit_env"] == "XNIX_CONTAINER_MEMORY_LIMIT", "q4 full smoke must expose the memory limit env")
assert(payload["container_memory_limit"] == "12g", "q4 full smoke must default to a larger q4 memory limit")
assert(payload["container_cpu_limit_env"] == "XNIX_CONTAINER_CPU_LIMIT", "q4 full smoke must expose the CPU limit env")
assert(payload["container_cpu_limit"] == "2.0", "q4 full smoke must default to a q4 CPU limit")
assert(payload["q4_heavy_smoke_resource_override"] == true, "q4 full smoke must make the heavy-smoke resource override observable")
assert(payload["known_winapp_fetch_timeout_env"] == "XNIX_KNOWN_WINAPP_FETCH_TIMEOUT", "q4 full smoke must expose the known app fetch timeout env")
assert(payload["known_winapp_fetch_timeout"] == "300s", "q4 full smoke must default to a longer known app fetch timeout")
assert(payload["q4_known_winapp_fetch_timeout_extended"] == true, "q4 full smoke must make the known app fetch timeout extension observable")
assert(payload["remote_full_smoke_planned"] == true, "q4 full smoke must plan remote execution")
assert(payload["remote_preflight_planned"] == true, "q4 full smoke must plan a remote preflight")
assert(payload["remote_preflight_command"] == "ruby+rsync+docker availability over SSH", "q4 full smoke must describe its preflight")
assert(payload["remote_preflight_timeout_seconds"] == 45, "q4 full smoke must expose the default preflight timeout")
assert(payload["q4_compile_required"] == true, "q4 full smoke must require q4 compilation")
assert(payload["host_compilation_avoided"] == true, "q4 full smoke must avoid host compilation")
assert(payload["host_root_modified"] == false, "q4 full smoke must not mutate host root")
assert(payload["privileged_container_required"] == false, "q4 full smoke must not require privileged containers")
assert(payload["host_networking_required"] == false, "q4 full smoke must not require host networking")
assert(payload["docker_socket_mounted"] == false, "q4 full smoke must not mount Docker socket")
assert(payload["broad_host_mount_required"] == false, "q4 full smoke must not require broad host mounts")

custom_stdout, custom_stderr, custom_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-source-root", "/tmp/xnix-q4-full-smoke/custom",
  "--report-root", "/tmp/xnix-q4-full-smoke-report",
  "--tools-base-image", "xnix-builder-tools:0.2.640-rc219-dbus-amd64",
  "--wine-base-image", "python:3.12-slim",
  "--container-memory-limit", "10g",
  "--container-cpu-limit", "1.5",
  "--known-winapp-fetch-timeout", "5m",
  "--remote-preflight-timeout-seconds", "17",
  "--timeout-seconds", "123",
  chdir: project_root.to_s
)
assert(custom_status.success?, "q4 full smoke custom plan must succeed: #{custom_stderr}")
custom_payload = JSON.parse(custom_stdout)
assert(custom_payload["remote_source_root"].start_with?("/tmp/xnix-"), "q4 full smoke must allow constrained /tmp/xnix-* remote roots")
assert(custom_payload["report_root"].start_with?("/tmp/xnix-"), "q4 full smoke must allow constrained /tmp/xnix-* report roots")
assert(custom_payload["tools_base_image"] == "xnix-builder-tools:0.2.640-rc219-dbus-amd64", "q4 full smoke must expose custom tools base image")
assert(custom_payload["wine_base_image"] == "python:3.12-slim", "q4 full smoke must expose custom Wine GUI smoke base image")
assert(custom_payload["container_memory_limit"] == "10g", "q4 full smoke must expose custom memory limit")
assert(custom_payload["container_cpu_limit"] == "1.5", "q4 full smoke must expose custom CPU limit")
assert(custom_payload["known_winapp_fetch_timeout"] == "5m", "q4 full smoke must expose custom known app fetch timeout")
assert(custom_payload["remote_preflight_timeout_seconds"] == 17, "q4 full smoke must expose custom preflight timeout")
assert(custom_payload["remote_timeout_seconds"] == 123, "q4 full smoke must expose custom timeout")

bad_root_stdout, bad_root_stderr, bad_root_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-source-root", "/var/tmp/xnix-q4-full-smoke",
  chdir: project_root.to_s
)
assert(!bad_root_status.success?, "q4 full smoke must reject unsafe remote roots")
assert((bad_root_stdout + bad_root_stderr).include?("remote source root must stay under /home/xnix-* or /tmp/xnix-*"), "q4 full smoke must explain unsafe remote roots")

bad_report_stdout, bad_report_stderr, bad_report_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--report-root", "/var/tmp/xnix-q4-full-smoke",
  chdir: project_root.to_s
)
assert(!bad_report_status.success?, "q4 full smoke must reject unsafe local report roots")
assert((bad_report_stdout + bad_report_stderr).include?("report root must stay under this checkout or /tmp/xnix-*"), "q4 full smoke must explain unsafe local report roots")

bad_base_stdout, bad_base_stderr, bad_base_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--tools-base-image", "bad image",
  chdir: project_root.to_s
)
assert(!bad_base_status.success?, "q4 full smoke must reject whitespace in Docker image refs")
assert((bad_base_stdout + bad_base_stderr).include?("tools base image must not contain shell whitespace"), "q4 full smoke must explain unsafe Docker image refs")

bad_memory_stdout, bad_memory_stderr, bad_memory_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--container-memory-limit", "bad limit",
  chdir: project_root.to_s
)
assert(!bad_memory_status.success?, "q4 full smoke must reject whitespace in Docker memory limits")
assert((bad_memory_stdout + bad_memory_stderr).include?("container memory limit must not contain shell whitespace"), "q4 full smoke must explain unsafe Docker memory limits")

bad_duration_stdout, bad_duration_stderr, bad_duration_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--known-winapp-fetch-timeout", "300 seconds",
  chdir: project_root.to_s
)
assert(!bad_duration_status.success?, "q4 full smoke must reject whitespace in known app fetch timeouts")
assert((bad_duration_stdout + bad_duration_stderr).include?("known Windows app fetch timeout must not contain shell whitespace"), "q4 full smoke must explain unsafe known app fetch timeouts")

bad_duration_shape_stdout, bad_duration_shape_stderr, bad_duration_shape_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--known-winapp-fetch-timeout", "300",
  chdir: project_root.to_s
)
assert(!bad_duration_shape_status.success?, "q4 full smoke must reject duration values without units")
assert((bad_duration_shape_stdout + bad_duration_shape_stderr).include?("known Windows app fetch timeout must look like a Go duration"), "q4 full smoke must explain invalid known app fetch timeout shapes")

puts "PASS: q4 full smoke script is remote-first and execute-gated"
