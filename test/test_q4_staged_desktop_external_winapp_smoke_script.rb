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
assert(source.include?("remote_run_root_cleaned"), "q4 staged external smoke must report cleaning its own remote run root")
assert(source.include?('shell_join(["rm", "-rf", remote_run_root])'), "q4 staged external smoke must clean only its validated remote run root")
assert(source.include?("one_shot_status"), "q4 staged external smoke must report one-shot Runtime command status")
assert(source.include?("one_shot_staged_launcher_invoked"), "q4 staged external smoke must require one-shot staged launcher invocation")
assert(source.include?("one_shot_runtime_launch_executed"), "q4 staged external smoke must require one-shot Runtime launch execution")
assert(source.include?("one_shot_raw_paths_exposed"), "q4 staged external smoke must preserve one-shot path redaction evidence")
assert(source.include?("host_compilation_avoided"), "q4 staged external smoke must report host compilation avoidance")
assert(source.include?("desktop_launch_packet_safe_for_kde"), "q4 staged external smoke must preserve KDE-safe launch packet evidence")
assert(source.include?("fetch_remote_artifact"), "q4 staged external smoke must fetch structured desktop evidence artifacts")
assert(source.include?("artifact_fetch_count"), "q4 staged external smoke must report fetched artifact count")
assert(source.include?("kde-external-app-page.json"), "q4 staged external smoke must fetch the KDE external app page artifact")
assert(source.include?("one-shot-import-stage-launch.json"), "q4 staged external smoke must fetch the one-shot Runtime artifact")
assert(source.include?("external-winapp-compatibility-evidence-bundle-preview"), "q4 staged external smoke must generate the Runtime compatibility evidence bundle on q4")
assert(source.include?("compatibility-evidence-bundle.json"), "q4 staged external smoke must fetch the compatibility evidence bundle artifact")
assert(source.include?("external-winapp-application-detail-preview"), "q4 staged external smoke must generate the Runtime application detail on q4")
assert(source.include?("external-winapp-application-detail.json"), "q4 staged external smoke must fetch the Runtime application detail artifact")
assert(source.include?("kde-center-page-preview"), "q4 staged external smoke must generate a KDE page from the Runtime application detail")
assert(source.include?("kde-page-from-application-detail.json"), "q4 staged external smoke must fetch the KDE page from application detail artifact")
assert(source.include?('cd #{Shellwords.escape(remote_source_root)} &&'), "q4 staged external smoke must render follow-up Runtime artifacts inside the q4 source root")
assert(source.include?("shell_join(kde_page_from_detail_command)"), "q4 staged external smoke must shell-join the KDE page from detail command")
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
assert(plan.fetch("artifact_output_root") == "/tmp/artifacts", "plan must derive artifact output root from the JSON output path")
assert(plan.fetch("artifact_output_count") == 10, "plan must announce all desktop evidence artifacts")
artifact_outputs = plan.fetch("artifact_outputs")
assert(artifact_outputs.fetch("one_shot_artifact_output_path") == "/tmp/artifacts/one-shot-import-stage-launch.json", "plan must expose one-shot artifact output")
assert(artifact_outputs.fetch("runtime_packet_artifact_output_path") == "/tmp/artifacts/runtime-gui-evidence-packet.json", "plan must expose Runtime packet artifact output")
assert(artifact_outputs.fetch("kde_page_artifact_output_path") == "/tmp/artifacts/kde-external-app-page.json", "plan must expose KDE page artifact output")
assert(artifact_outputs.fetch("compatibility_evidence_bundle_artifact_output_path") == "/tmp/artifacts/compatibility-evidence-bundle.json", "plan must expose compatibility evidence bundle artifact output")
assert(artifact_outputs.fetch("application_detail_artifact_output_path") == "/tmp/artifacts/external-winapp-application-detail.json", "plan must expose Runtime application detail artifact output")
assert(artifact_outputs.fetch("kde_page_from_application_detail_artifact_output_path") == "/tmp/artifacts/kde-page-from-application-detail.json", "plan must expose KDE page from Runtime application detail artifact output")
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
