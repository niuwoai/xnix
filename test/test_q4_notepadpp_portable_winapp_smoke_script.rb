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
script = project_root.join("scripts/q4_notepadpp_portable_winapp_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_notepadpp_portable_winapp_smoke.v1"), "q4 Notepad++ smoke must expose a stable schema")
assert(source.include?("q4-notepadpp-portable-winapp-smoke"), "q4 Notepad++ smoke must expose a stable request type")
assert(source.include?("https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.9.7/npp.8.9.7.portable.zip"), "q4 Notepad++ smoke must target the official portable zip")
assert(source.include?("ce0690fac91c1fc5d61dcdf5b09733ff0d143a61d0a27c6cb9f4003ea92765bb"), "q4 Notepad++ smoke must pin the official portable zip checksum")
assert(source.include?("curl"), "q4 Notepad++ smoke must download the portable zip on q4")
assert(source.include?("--retry-all-errors"), "q4 Notepad++ smoke must tolerate transient q4 download failures")
assert(source.include?("--continue-at"), "q4 Notepad++ smoke must resume interrupted q4 downloads")
assert(source.include?("sha256sum"), "q4 Notepad++ smoke must verify the portable zip on q4")
assert(source.include?("unzip"), "q4 Notepad++ smoke must extract the portable zip on q4")
assert(source.include?("python3"), "q4 Notepad++ smoke must keep a q4 zip extraction fallback")
assert(source.include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "q4 Notepad++ smoke must delegate to the staged external Runtime/KDE path")
assert(source.include?("--remote-bundle-root"), "q4 Notepad++ smoke must delegate a q4 portable bundle root")
assert(source.include?("--executable-relative-path"), "q4 Notepad++ smoke must delegate the selected portable executable")
assert(source.include?("notepad++.exe"), "q4 Notepad++ smoke must select the portable Notepad++ executable")
assert(source.include?("portable_directory_external_app"), "q4 Notepad++ smoke must record portable-directory evidence")
assert(source.include?("portable_directory_bundle_import_required"), "q4 Notepad++ smoke must require the bundle import path")
assert(source.include?("real_third_party_windows_app"), "q4 Notepad++ smoke must distinguish real third-party app evidence")
assert(source.include?("host_download_avoided"), "q4 Notepad++ smoke must avoid host-side download")
assert(source.include?("host_compilation_avoided"), "q4 Notepad++ smoke must avoid host-side compilation")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-notepadpp-plan.json",
  "--markdown-output", "/tmp/xnix-q4-notepadpp-plan.md"
)
assert(status.success?, "q4 Notepad++ plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_notepadpp_portable_winapp_smoke.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-notepadpp-portable-winapp-smoke", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("source_kind") == "official-notepad-plus-plus-github-release", "plan must identify official source")
assert(payload.fetch("notepadpp_version") == "8.9.7", "plan must pin the Notepad++ version")
assert(payload.fetch("notepadpp_zip") == "npp.8.9.7.portable.zip", "plan must pin the portable zip")
assert(payload.fetch("notepadpp_sha256") == "ce0690fac91c1fc5d61dcdf5b09733ff0d143a61d0a27c6cb9f4003ea92765bb", "plan must pin the portable zip checksum")
assert(payload.fetch("notepadpp_download_planned") == true, "plan must download the portable zip")
assert(payload.fetch("notepadpp_downloaded") == false, "plan must not claim download before execute")
assert(payload.fetch("notepadpp_sha256_verified") == false, "plan must not claim checksum before execute")
assert(payload.fetch("notepadpp_extracted") == false, "plan must not claim extraction before execute")
assert(payload.fetch("notepadpp_executable_configured") == true, "plan must configure the executable")
assert(payload.fetch("notepadpp_executable_relative_path") == "notepad++.exe", "plan must select the portable executable")
assert(payload.fetch("portable_directory_external_app") == true, "plan must mark the app as portable-directory")
assert(payload.fetch("portable_directory_bundle_import_required") == true, "plan must require bundle import")
assert(payload.fetch("real_third_party_windows_app") == true, "plan must mark the app as a real third-party app")
assert(payload.fetch("single_file_windows_app") == false, "plan must not claim the app is single-file")
assert(payload.fetch("delegated_command").include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "plan must delegate to staged external smoke")
assert(payload.fetch("delegated_command").include?("--remote-bundle-root"), "plan must pass the q4 portable bundle root")
assert(payload.fetch("delegated_command").include?("--executable-relative-path"), "plan must pass the selected executable")
assert(payload.fetch("delegated_command").include?("org.xnix.external.notepadplusplus"), "plan must pass app identity")
assert(payload.fetch("delegated_command").include?("Notepad++ Portable"), "plan must pass display name")
assert(payload.fetch("runtime_owned") == true, "plan must keep Runtime ownership")
assert(payload.fetch("go_runtime_backed") == true, "plan must keep Go Runtime backing")
assert(payload.fetch("kde_policy_owner") == false, "plan must not make KDE the policy owner")
assert(payload.fetch("q4_download_required") == true, "plan must require q4 download")
assert(payload.fetch("q4_extract_required") == true, "plan must require q4 extraction")
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
assert(!status.success?, "q4 Notepad++ plan must reject unsafe remote materials roots")
assert(stderr.include?("remote materials root must stay under /home/xnix-* or /tmp/xnix-* on q4"), "q4 Notepad++ plan must explain unsafe roots")

puts "PASS: q4 Notepad++ Portable Windows app smoke is official-download and q4-bundle-runtime backed"
