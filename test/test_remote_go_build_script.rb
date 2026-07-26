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
script = project_root.join("scripts/remote_go_build.rb")
source = script.read

assert(source.include?("root@q4"), "remote Go build must default to q4")
assert(source.include?("DEFAULT_PACKAGES"), "remote Go build must declare default command packages")
assert(source.include?("./cmd/xnix-runtime-go"), "remote Go build must build the Go Runtime by default")
assert(source.include?("./cmd/xnix-runtime-owner"), "remote Go build must build the Runtime owner by default")
assert(source.include?("./cmd/xnix-compat-launch"), "remote Go build must build the managed launcher by default")
assert(source.include?("./cmd/xnix-compat-open"), "remote Go build must build the managed file-open command by default")
assert(source.include?("docs/claude-code-implementation-packages.md"), "remote Go build source sync must exclude the protected Claude package document")
assert(source.include?("ensure_remote_xnix_path!"), "remote Go build must constrain remote writable paths")
assert(source.include?("host_compilation_avoided"), "remote Go build must make host compilation avoidance observable")
assert(source.include?("ssh_command"), "remote Go build must use bounded SSH options")
assert(source.include?("BatchMode=yes"), "remote Go build must use non-interactive SSH mode")
assert(source.include?("ServerAliveInterval=15"), "remote Go build must use SSH keepalive")
assert(source.include?("remote shell operation timed out"), "remote Go build must bound remote shell operations")
assert(source.include?("GOOS="), "remote Go build must set GOOS remotely")
assert(source.include?("GOARCH="), "remote Go build must set GOARCH remotely")
assert(source.include?("GOCACHE="), "remote Go build must keep Go cache on q4")
assert(source.include?("GOMODCACHE="), "remote Go build must keep module cache on q4")
assert(source.include?("GOTMPDIR="), "remote Go build must keep temporary files on q4")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "remote Go build plan must succeed: #{stderr}")

payload = JSON.parse(stdout)
assert(payload["schema_version"] == "xnix.scripts.remote_go_build.v1", "remote Go build must expose a stable schema")
assert(payload["request_type"] == "remote-go-build", "remote Go build must expose request type")
assert(payload["status"] == "planned", "remote Go build must be planned by default")
assert(payload["execute"] == false, "remote Go build must not execute without --execute")
assert(payload["remote_host"] == "root@q4", "remote Go build must default to q4")
assert(payload["source_sync_planned"] == true, "remote Go build must sync source by default")
assert(payload["source_sync_mode"] == "runtime", "remote Go build must default to Runtime source sync")
assert(payload["source_sync_entry_count"] == 10, "remote Go build Runtime sync must expose the source entry count")
%w[.dockerignore Dockerfile VERSION go.mod cmd internal runtime scripts lib kde].each do |entry|
  assert(payload["source_sync_entries"].include?(entry), "remote Go build Runtime sync must include #{entry}")
end
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "remote source root must stay under /home/xnix-*")
assert(payload["remote_source_root"].include?("xnix-remote-go-build-runtime-"), "remote source root must include sync mode")
assert(payload["remote_build_root"].start_with?("/home/xnix-"), "remote build root must stay under /home/xnix-*")
assert(payload["remote_output_root"].end_with?("/bin/linux-amd64"), "remote output root must include target platform")
assert(payload["goos"] == "linux", "remote Go build must default GOOS to linux")
assert(payload["goarch"] == "amd64", "remote Go build must default GOARCH to amd64")
assert(payload["package_count"] == 4, "remote Go build must build four core command packages by default")
assert(payload["build_targets"].map { |entry| entry.fetch("package") } == %w[./cmd/xnix-runtime-go ./cmd/xnix-runtime-owner ./cmd/xnix-compat-launch ./cmd/xnix-compat-open], "remote Go build must expose default package order")
assert(payload["build_targets"].all? { |entry| entry.fetch("remote_output").start_with?(payload["remote_output_root"]) }, "remote Go build outputs must stay under the remote output root")
assert(payload["remote_build_planned"] == true, "remote Go build must plan remote build")
assert(payload["host_compilation_avoided"] == true, "remote Go build must avoid host compilation")
assert(payload["host_root_modified"] == false, "remote Go build must not mutate host root")
assert(payload["privileged_container_required"] == false, "remote Go build must not require privileged containers")
assert(payload["host_networking_required"] == false, "remote Go build must not require host networking")
assert(payload["docker_socket_mounted"] == false, "remote Go build must not mount Docker socket")
assert(payload["broad_host_mount_required"] == false, "remote Go build must not require broad host mounts")

custom_stdout, custom_stderr, custom_status = Open3.capture3(
  "ruby", script.to_s,
  "--source-sync-mode", "full",
  "--remote-source-root", "/tmp/xnix-remote-build/full",
  "--remote-build-root", "/tmp/xnix-remote-cache",
  "--package", "./cmd/xnix-runtime-go",
  "--goos", "windows",
  "--goarch", "amd64",
  chdir: project_root.to_s
)
assert(custom_status.success?, "remote Go build custom plan must succeed: #{custom_stderr}")
custom_payload = JSON.parse(custom_stdout)
assert(custom_payload["source_sync_mode"] == "full", "remote Go build must expose full source sync mode")
assert(custom_payload["source_sync_entries"] == ["."], "remote Go build full sync must use the full checkout")
assert(custom_payload["remote_source_root"].start_with?("/tmp/xnix-"), "remote Go build must allow constrained /tmp/xnix-* source roots")
assert(custom_payload["remote_build_root"].start_with?("/tmp/xnix-"), "remote Go build must allow constrained /tmp/xnix-* build roots")
assert(custom_payload["goos"] == "windows", "remote Go build must expose custom GOOS")
assert(custom_payload["goarch"] == "amd64", "remote Go build must expose custom GOARCH")
assert(custom_payload["package_count"] == 1, "remote Go build must expose custom package count")
assert(custom_payload["build_targets"].first.fetch("package") == "./cmd/xnix-runtime-go", "remote Go build must expose custom package")
assert(custom_payload["remote_output_root"].end_with?("/bin/windows-amd64"), "remote Go build custom output root must include target platform")

bad_package_stdout, bad_package_stderr, bad_package_status = Open3.capture3(
  "ruby", script.to_s,
  "--package", "../cmd/xnix-runtime-go",
  chdir: project_root.to_s
)
assert(!bad_package_status.success?, "remote Go build must reject parent-traversal packages")
assert((bad_package_stdout + bad_package_stderr).include?("repository-relative command package"), "remote Go build must explain unsafe package targets")

bad_root_stdout, bad_root_stderr, bad_root_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-build-root", "/var/tmp/xnix-cache",
  chdir: project_root.to_s
)
assert(!bad_root_status.success?, "remote Go build must reject remote build roots outside constrained prefixes")
assert((bad_root_stdout + bad_root_stderr).include?("remote build root must stay under /home/xnix-* or /tmp/xnix-*"), "remote Go build must explain unsafe build roots")

darwin_stdout, darwin_stderr, darwin_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--goos", "darwin",
  "--goarch", "arm64",
  "--package", "./cmd/xnix-runtime-go"
)
assert(darwin_status.success?, "remote Go build must plan q4 cross-compiled Darwin Runtime binaries: #{darwin_stderr}")
darwin_payload = JSON.parse(darwin_stdout)
assert(darwin_payload["goos"] == "darwin", "remote Go build must expose Darwin GOOS")
assert(darwin_payload["goarch"] == "arm64", "remote Go build must expose Darwin arm64 GOARCH")
assert(darwin_payload["remote_output_root"].end_with?("/bin/darwin-arm64"), "remote Go build Darwin output root must include target platform")

bad_platform_stdout, bad_platform_stderr, bad_platform_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--goos", "freebsd"
)
assert(!bad_platform_status.success?, "remote Go build must reject unsupported target GOOS")
assert((bad_platform_stdout + bad_platform_stderr).include?("GOOS must be linux, windows, or darwin"), "remote Go build must explain unsupported GOOS")

puts "PASS: remote Go build script is q4-first and execute-gated"
