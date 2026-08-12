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
script = project_root.join("scripts/q4_staged_desktop_external_winapp_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1"), "q4 staged external smoke must expose a stable schema")
assert(source.include?("q4-staged-desktop-external-winapp-smoke"), "q4 staged external smoke must expose a stable request type")
assert(source.include?("scripts/remote_go_build.rb"), "q4 staged external smoke must build Runtime binaries on q4")
assert(source.include?("scripts/staged_desktop_external_winapp_smoke.rb"), "q4 staged external smoke must delegate to the maintained staged desktop smoke")
assert(source.include?("./cmd/xnix-runtime-go"), "q4 staged external smoke must build xnix-runtime-go")
assert(source.include?("./cmd/xnix-compat-launch"), "q4 staged external smoke must build xnix-compat-launch")
assert(source.include?("XNIX_RUNTIME_GO_BIN"), "q4 staged external smoke must pass the q4-built Runtime binary")
assert(source.include?("XNIX_COMPAT_LAUNCH_BIN"), "q4 staged external smoke must pass the q4-built launcher binary")
assert(source.include?("XNIX_ALLOW_LOCAL_GO_COMPILE"), "q4 staged external smoke must explicitly disable local Go compilation")
assert(source.include?("one_shot_status"), "q4 staged external smoke must report one-shot Runtime command status")
assert(source.include?("one_shot_staged_launcher_invoked"), "q4 staged external smoke must require one-shot staged launcher invocation")
assert(source.include?("one_shot_runtime_launch_executed"), "q4 staged external smoke must require one-shot Runtime launch execution")
assert(source.include?("one_shot_raw_paths_exposed"), "q4 staged external smoke must preserve one-shot path redaction evidence")
assert(source.include?("host_compilation_avoided"), "q4 staged external smoke must report host compilation avoidance")
assert(source.include?("desktop_launch_packet_safe_for_kde"), "q4 staged external smoke must preserve KDE-safe launch packet evidence")
assert(!source.include?("go build"), "q4 staged external smoke wrapper must not compile locally")
assert(!source.include?("go test"), "q4 staged external smoke wrapper must not test locally")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--fixture", "notepad-file-argument",
  "--image", "xnix-wine-smoke:local",
  "--remote", "root@q4",
  "--remote-source-root", "/home/xnix-build/xnix-q4-staged-external-winapp-test",
  "--remote-build-root", "/home/xnix-build-cache",
  "--remote-run-root", "/tmp/xnix-q4-staged-external-winapp-test",
  "--output", "/tmp/xnix-q4-staged-external-plan.json",
  "--markdown-output", "/tmp/xnix-q4-staged-external-plan.md"
)
assert(status.success?, "q4 staged external plan must exit successfully: #{stderr}")
plan = JSON.parse(stdout)
assert(plan.fetch("schema_version") == "xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1", "plan must expose schema")
assert(plan.fetch("request_type") == "q4-staged-desktop-external-winapp-smoke", "plan must expose request type")
assert(plan.fetch("status") == "planned", "plan must not execute by default")
assert(plan.fetch("execute") == false, "plan must keep execution disabled by default")
assert(plan.fetch("remote_host") == "root@q4", "plan must preserve remote host")
assert(plan.fetch("remote_source_root") == "/home/xnix-build/xnix-q4-staged-external-winapp-test", "plan must preserve q4 source root")
assert(plan.fetch("remote_build_root") == "/home/xnix-build-cache", "plan must preserve q4 build root")
assert(plan.fetch("remote_run_root") == "/tmp/xnix-q4-staged-external-winapp-test", "plan must preserve q4 run root")
assert(plan.fetch("remote_runtime_binary") == "/home/xnix-build-cache/bin/linux-amd64/xnix-runtime-go", "plan must point at q4 Runtime binary")
assert(plan.fetch("remote_launcher_binary") == "/home/xnix-build-cache/bin/linux-amd64/xnix-compat-launch", "plan must point at q4 launcher binary")
assert(plan.fetch("delegated_script") == "scripts/staged_desktop_external_winapp_smoke.rb", "plan must delegate staged smoke")
assert(plan.fetch("delegated_command").include?("--fixture"), "plan must forward fixture")
assert(plan.fetch("delegated_command").include?("notepad-file-argument"), "plan must default to file-argument Notepad fixture")
assert(plan.fetch("q4_compile_required") == true, "plan must require q4 compile")
assert(plan.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(plan.fetch("targeted_smoke_required") == true, "plan must identify targeted smoke")
assert(plan.fetch("full_smoke_required") == false, "plan must not claim full smoke is required")
assert(plan.fetch("desktop_shell") == "KDE Plasma", "plan must keep KDE as the desktop shell")
assert(plan.fetch("runtime_owned") == true, "plan must keep Runtime ownership")
assert(plan.fetch("go_runtime_backed") == true, "plan must keep Go Runtime backing")
assert(plan.fetch("kde_policy_owner") == false, "plan must not give policy ownership to KDE")
assert(plan.fetch("host_root_modified") == false, "plan must not mutate the host root")
assert(plan.fetch("privileged_container_required") == false, "plan must avoid privileged containers")
assert(plan.fetch("host_networking_required") == false, "plan must avoid host networking")
assert(plan.fetch("docker_socket_mounted") == false, "plan must avoid Docker socket mounts")
assert(plan.fetch("broad_host_mount_required") == false, "plan must avoid broad host mounts")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-source-root", "/Users/rocky/Sites/xnix",
  "--remote-run-root", "/tmp/xnix-q4-staged-external-winapp-test"
)
assert(!status.success?, "q4 staged external smoke must reject non-q4 source roots")
assert(stderr.include?("remote source root must stay under /home/xnix-* or /tmp/xnix-* on q4"), "q4 staged external smoke must explain unsafe source root")

puts "PASS: q4 staged desktop external Windows app smoke script is q4-first"
