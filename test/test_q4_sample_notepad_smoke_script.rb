# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tempfile"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/q4_sample_notepad_smoke.rb")
source = script.read

assert(source.include?("scripts/remote_wine_guest_gui_smoke.rb"), "q4 Sample Notepad smoke must delegate to the maintained q4 GUI smoke")
assert(source.include?("\"--launch-mode\", \"owner-controlled-launch\""), "q4 Sample Notepad smoke must use owner-controlled launch")
assert(source.include?("\"--file-open-entrypoint\""), "q4 Sample Notepad smoke must use the desktop file-open entrypoint")
assert(source.include?("\"--known-app-id\", \"org.xnix.sample.notepad\""), "q4 Sample Notepad smoke must select Sample Notepad")
assert(source.include?("\"--sample-file-argument\""), "q4 Sample Notepad smoke must create a sample document argument")
assert(source.include?("\"--window-match\""), "q4 Sample Notepad smoke must require observed window text")
assert(source.include?("--output PATH"), "q4 Sample Notepad smoke must support writing reusable evidence")
assert(source.include?("ensure_local_output_path!"), "q4 Sample Notepad smoke must constrain local evidence output paths")
assert(source.include?("real_run_acceptance_ready"), "q4 Sample Notepad smoke must require real run acceptance")
assert(source.include?("host_compilation_avoided"), "q4 Sample Notepad smoke must document host compilation avoidance")
assert(!source.include?("go build"), "q4 Sample Notepad smoke wrapper must not compile locally")
assert(!source.include?("go test"), "q4 Sample Notepad smoke wrapper must not test locally")
assert(!source.include?("docs/claude-code-implementation-packages.md"), "q4 Sample Notepad smoke wrapper must not touch Claude implementation packages")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "q4 Sample Notepad smoke plan must succeed: #{stderr}")
payload = JSON.parse(stdout)

assert(payload["schema_version"] == "xnix.scripts.q4_sample_notepad_smoke.v1", "q4 Sample Notepad smoke must expose a stable schema")
assert(payload["request_type"] == "q4-sample-notepad-smoke", "q4 Sample Notepad smoke must expose request type")
assert(payload["status"] == "planned", "q4 Sample Notepad smoke must be planned by default")
assert(payload["execute"] == false, "q4 Sample Notepad smoke must not execute without --execute")
assert(payload["remote_host"] == "root@q4", "q4 Sample Notepad smoke must default to q4")
assert(payload["delegated_script"] == "scripts/remote_wine_guest_gui_smoke.rb", "q4 Sample Notepad smoke must expose delegated script")
assert(payload["output_path"] == "", "q4 Sample Notepad smoke must not write output by default")
assert(payload.fetch("delegated_command").include?("scripts/remote_wine_guest_gui_smoke.rb"), "q4 Sample Notepad smoke must expose a relative delegated command")
assert(payload.fetch("delegated_command").none? { |part| part.include?(project_root.to_s) }, "q4 Sample Notepad smoke must not expose local absolute paths")
assert(payload["source_sync_planned"] == true, "q4 Sample Notepad smoke must sync source by default")
assert(payload["app_id"] == "org.xnix.sample.notepad", "q4 Sample Notepad smoke must target Sample Notepad")
assert(payload["display_name"] == "Sample Notepad", "q4 Sample Notepad smoke must expose display name")
assert(payload["sample_file_argument"] == "sample-document.txt", "q4 Sample Notepad smoke must use the default sample document")
assert(payload["window_match"] == "sample-document.txt", "q4 Sample Notepad smoke must match the default sample document window title")
assert(payload["launch_mode"] == "owner-controlled-launch", "q4 Sample Notepad smoke must use owner-controlled launch")
assert(payload["file_open_entrypoint_requested"] == true, "q4 Sample Notepad smoke must request file-open entrypoint")
assert(payload["real_run_acceptance_required"] == true, "q4 Sample Notepad smoke must require acceptance")
assert(payload["q4_compile_required"] == true, "q4 Sample Notepad smoke must declare q4 compilation")
assert(payload["host_compilation_avoided"] == true, "q4 Sample Notepad smoke must avoid host compilation")
%w[--launch-mode owner-controlled-launch --file-open-entrypoint --known-app-id org.xnix.sample.notepad --sample-file-argument sample-document.txt --window-match sample-document.txt --execute].each do |token|
  assert(payload.fetch("delegated_command").include?(token) == (token != "--execute"), "plan command must include #{token} only when expected")
end
assert(payload["host_root_modified"] == false, "q4 Sample Notepad smoke must not modify host root")
assert(payload["privileged_container_required"] == false, "q4 Sample Notepad smoke must not require privileged containers")
assert(payload["host_networking_required"] == false, "q4 Sample Notepad smoke must not require host networking")
assert(payload["docker_socket_mounted"] == false, "q4 Sample Notepad smoke must not mount Docker socket")
assert(payload["broad_host_mount_required"] == false, "q4 Sample Notepad smoke must not require broad host mounts")

custom_stdout, custom_stderr, custom_status = Open3.capture3(
  "ruby", script.to_s,
  "--no-sync-source",
  "--remote", "root@q4",
  "--sample-file", "daily-note.txt",
  "--window-match", "daily-note.txt",
  "--remote-timeout-seconds", "333",
  chdir: project_root.to_s
)
assert(custom_status.success?, "q4 Sample Notepad custom plan must succeed: #{custom_stderr}")
custom_payload = JSON.parse(custom_stdout)
assert(custom_payload["source_sync_planned"] == false, "custom plan must forward no-sync-source")
assert(custom_payload["sample_file_argument"] == "daily-note.txt", "custom plan must expose sample file")
assert(custom_payload["window_match"] == "daily-note.txt", "custom plan must expose window match")
assert(custom_payload.fetch("delegated_command").include?("--no-sync-source"), "custom plan must forward --no-sync-source")
assert(custom_payload.fetch("delegated_command").include?("333"), "custom plan must forward timeout")

output_file = Tempfile.new(["xnix-q4-sample-notepad", ".json"], "/tmp")
output_path = output_file.path.sub("/tmp/", "/tmp/xnix-")
output_file.close
File.unlink(output_file.path)
output_stdout, output_stderr, output_status = Open3.capture3(
  "ruby", script.to_s,
  "--output", output_path,
  chdir: project_root.to_s
)
assert(output_status.success?, "q4 Sample Notepad output plan must succeed: #{output_stderr}")
output_payload = JSON.parse(output_stdout)
written_payload = JSON.parse(Pathname.new(output_path).read)
assert(output_payload == written_payload, "q4 Sample Notepad smoke must write the same JSON it prints")
assert(written_payload["output_path"] == output_path, "q4 Sample Notepad smoke must expose constrained output path")
File.unlink(output_path)

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby", script.to_s,
  "--sample-file", "../escape.txt",
  chdir: project_root.to_s
)
assert(!bad_status.success?, "q4 Sample Notepad smoke must reject path-like sample files")
assert((bad_stdout + bad_stderr).include?("sample file must be a simple file name"), "q4 Sample Notepad smoke must explain unsafe sample file names")

bad_output_stdout, bad_output_stderr, bad_output_status = Open3.capture3(
  "ruby", script.to_s,
  "--output", "/var/tmp/not-xnix/q4.json",
  chdir: project_root.to_s
)
assert(!bad_output_status.success?, "q4 Sample Notepad smoke must reject broad output paths")
assert((bad_output_stdout + bad_output_stderr).include?("output path must stay under this checkout or /tmp/xnix-*"), "q4 Sample Notepad smoke must explain unsafe output paths")

puts "PASS: q4 Sample Notepad smoke script is execute-gated and q4-first"
