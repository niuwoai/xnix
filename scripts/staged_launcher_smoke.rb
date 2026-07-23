#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
RUN_ROOT = PROJECT_ROOT.join(".cache/xnix/staged-launcher-smoke/#{Process.pid}")
BUILD_DIR = RUN_ROOT.join("build")
STAGE_ROOT = RUN_ROOT.join("stage")
KNOWN_APP_CACHE = RUN_ROOT.join("known-app-cache")
LAUNCHER_BIN = BUILD_DIR.join("xnix-compat-launch")
STAGED_LAUNCHER = STAGE_ROOT.join("usr/local/bin/xnix-compat-launch")
GO_ENV = {
  "GOCACHE" => PROJECT_ROOT.join(".cache/go-build").to_s,
  "GOMODCACHE" => PROJECT_ROOT.join(".cache/go-mod").to_s
}.freeze
SMOKE_NAME = "staged managed launcher smoke"
GO_BUILD_STEP = "go build"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def run_command(env, *command)
  stdout, stderr, status = Open3.capture3(env, *command, chdir: PROJECT_ROOT.to_s)
  assert(status.success?, "#{command.join(" ")} must exit successfully: #{stderr}\n#{stdout}")
  stdout
end

def parse_json(stdout, label)
  JSON.parse(stdout)
rescue JSON::ParserError => e
  warn stdout
  warn "FAIL: #{label} must emit JSON: #{e.message}"
  exit 1
end

def assert_no_forbidden(text, forbidden_terms, label)
  downcased = text.downcase
  forbidden_terms.each do |term|
    next if term.to_s.empty?

    assert(!downcased.include?(term.to_s.downcase), "#{label} must not expose #{term}")
  end
end

FileUtils.mkdir_p(BUILD_DIR)
FileUtils.mkdir_p(KNOWN_APP_CACHE)

run_command(GO_ENV, "go", "build", "-o", LAUNCHER_BIN.to_s, "./cmd/xnix-compat-launch")

stage_stdout = run_command(
  GO_ENV,
  "go", "run", "./cmd/xnix-runtime-go",
  "desktop-activation-stage",
  "--registry", "runtime/recipes/registry.json",
  "--app", "org.xnix.sample.notepad",
  "--mode", "development",
  "--staging-root", STAGE_ROOT.to_s,
  "--managed-launcher-bin", LAUNCHER_BIN.to_s
)
stage = parse_json(stage_stdout, "desktop activation stage")

assert(stage["schema_version"] == "xnix.runtime.desktop_activation_stage.v1", "desktop activation stage schema must match")
assert(stage["request_type"] == "desktop-activation-stage", "desktop activation stage request type must match")
assert(stage["written_file_count"] == 7, "desktop activation stage must write the managed launcher executable")
assert(stage.fetch("written_file_ids").include?("managed-launcher-executable"), "desktop activation stage must report managed-launcher-executable")
%w[
  runtime_owned
  go_runtime_backed
  staging_root_required
  file_writes_performed
  desktop_files_written
  mimeapps_written
  manifest_written
  receipt_written
  rollback_receipt_written
].each do |field|
  assert(stage[field] == true, "desktop activation stage must set #{field}=true")
end
%w[
  staging_root_path_exposed
  host_root_allowed
  launch_enabled
  backend_launch_enabled
  execution_started
  host_root_modified
  network_required
  privileged_container_required
  backend_details_exposed
].each do |field|
  assert(stage[field] == false, "desktop activation stage must set #{field}=false")
end
assert_no_forbidden(stage_stdout, [STAGE_ROOT.to_s, LAUNCHER_BIN.to_s, PROJECT_ROOT.to_s], "desktop activation stage output")

assert(STAGED_LAUNCHER.file?, "staged launcher executable must exist")
assert((STAGED_LAUNCHER.stat.mode & 0o777) == 0o755, "staged launcher executable must use mode 0755")

bridge_stdout = run_command({}, STAGED_LAUNCHER.to_s, "--app", "7zr", "--cache-root", KNOWN_APP_CACHE.to_s)
bridge = parse_json(bridge_stdout, "windows-known-app-launch-bridge-preview")

assert(bridge["schema_version"] == "xnix.runtime.known_windows_app_launch_bridge.v1", "launch bridge schema must match")
assert(bridge["request_type"] == "windows-known-app-launch-bridge-preview", "launch bridge request type must match")
assert(bridge["managed_launcher"] == "xnix-compat-launch --app 7zr", "launch bridge must preserve managed launcher")
assert(bridge["launcher_argv_accepted"] == true, "launch bridge must accept the staged launcher argv")
assert(bridge["runtime_owned_bridge"] == true, "launch bridge must be Runtime-owned")
assert(bridge["kde_presentation_only"] == true, "launch bridge must keep KDE presentation-only")
assert(bridge["dry_run"] == true, "launch bridge must remain dry-run")
%w[
  dispatch_started
  execution_started
  backend_process_started
  host_root_modified
  host_networking_required
  docker_socket_mounted
  broad_host_mount_required
  raw_host_path_exposed
  raw_executable_path_exposed
  raw_command_exposed
  backend_details_exposed
].each do |field|
  assert(bridge[field] == false, "launch bridge must set #{field}=false")
end
assert_no_forbidden(
  bridge_stdout,
  [RUN_ROOT.to_s, PROJECT_ROOT.to_s, ".exe", "wine ", "wine/", ".wine", "qemu-system", "program files"],
  "launch bridge output"
)

puts "PASS: #{SMOKE_NAME}"
