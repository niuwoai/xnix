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
script = project_root.join("scripts/q4_putty_external_winapp_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_putty_external_winapp_smoke.v1"), "q4 PuTTY smoke must expose a stable schema")
assert(source.include?("q4-putty-external-winapp-smoke"), "q4 PuTTY smoke must expose a stable request type")
assert(source.include?("https://the.earth.li/~sgtatham/putty/0.84/w32/putty.exe"), "q4 PuTTY smoke must target the official standalone executable")
assert(source.include?("d5a83cd1233f6da38fa82b14d970dbb2c2705769b5ebabb464918b9b57180bc4"), "q4 PuTTY smoke must pin the official executable checksum")
assert(source.include?("curl"), "q4 PuTTY smoke must download the executable on q4")
assert(source.include?("--retry-all-errors"), "q4 PuTTY smoke must tolerate transient q4 download failures")
assert(source.include?("--continue-at"), "q4 PuTTY smoke must resume interrupted q4 downloads")
assert(source.include?("sha256sum"), "q4 PuTTY smoke must verify the executable on q4")
assert(source.include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "q4 PuTTY smoke must delegate to the staged external Runtime/KDE path")
assert(source.include?("real_third_party_windows_app"), "q4 PuTTY smoke must distinguish real third-party app evidence")
assert(source.include?("single_file_windows_app"), "q4 PuTTY smoke must record that this app fits the current single-file lane")
assert(source.include?("remote_runtime_binary"), "q4 PuTTY smoke must expose q4 Runtime binary for operator follow-up")
assert(source.include?("staged_external_winapp_acceptance_ready"), "q4 PuTTY smoke must expose staged external acceptance readiness")
assert(source.include?("compatibility_evidence_bundle_report"), "q4 PuTTY smoke must expose the Runtime compatibility bundle path for operator follow-up")
assert(source.include?("host_download_avoided"), "q4 PuTTY smoke must avoid host-side download")
assert(source.include?("host_compilation_avoided"), "q4 PuTTY smoke must avoid host-side compilation")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-putty-plan.json",
  "--markdown-output", "/tmp/xnix-q4-putty-plan.md"
)
assert(status.success?, "q4 PuTTY plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_putty_external_winapp_smoke.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-putty-external-winapp-smoke", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("source_kind") == "official-putty-release", "plan must identify official source")
assert(payload.fetch("putty_version") == "0.84", "plan must pin the PuTTY version")
assert(payload.fetch("putty_executable") == "putty.exe", "plan must pin the executable")
assert(payload.fetch("putty_sha256") == "d5a83cd1233f6da38fa82b14d970dbb2c2705769b5ebabb464918b9b57180bc4", "plan must pin the executable checksum")
assert(payload.fetch("putty_download_planned") == true, "plan must download the executable")
assert(payload.fetch("putty_downloaded") == false, "plan must not claim download before execute")
assert(payload.fetch("putty_sha256_verified") == false, "plan must not claim checksum before execute")
assert(payload.fetch("putty_executable_configured") == true, "plan must configure the executable")
assert(payload.fetch("real_third_party_windows_app") == true, "plan must mark the app as a real third-party app")
assert(payload.fetch("single_file_windows_app") == true, "plan must record the current single-file lane")
assert(payload.fetch("delegated_command").include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "plan must delegate to staged external smoke")
assert(payload.fetch("delegated_command").include?("--remote-executable"), "plan must pass the q4 executable")
assert(payload.fetch("delegated_command").include?("org.xnix.external.putty"), "plan must pass app identity")
assert(payload.fetch("delegated_command").include?("PuTTY"), "plan must pass display name and window match")
assert(payload.fetch("runtime_owned") == true, "plan must keep Runtime ownership")
assert(payload.fetch("go_runtime_backed") == true, "plan must keep Go Runtime backing")
assert(payload.fetch("kde_policy_owner") == false, "plan must not make KDE the policy owner")
assert(payload.fetch("q4_download_required") == true, "plan must require q4 download")
assert(payload.fetch("q4_compile_required") == true, "plan must require q4 Runtime compilation")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("host_download_avoided") == true, "plan must avoid host download")
assert(payload.fetch("full_smoke_required") == false, "plan must remain targeted")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-materials-root", "/Users/rocky/not-q4"
)
assert(!status.success?, "q4 PuTTY plan must reject unsafe remote materials roots")
assert(stderr.include?("remote materials root must stay under /home/xnix-* or /tmp/xnix-* on q4"), "q4 PuTTY plan must explain unsafe roots")

puts "PASS: q4 PuTTY external Windows app smoke is official-download and staged-runtime backed"
