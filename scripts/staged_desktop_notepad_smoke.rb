#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "securerandom"
require "shellwords"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
SCHEMA_VERSION = "xnix.scripts.staged_desktop_notepad_smoke.v1"
APP_ID = "org.xnix.sample.notepad"
APP_NAME = "Sample Notepad"
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
PACKAGED_REGISTRY = "/usr/share/xnix/compatibility/recipes/registry.json"
DEFAULT_IMAGE = "xnix-wine-smoke:local"
DEFAULT_RUN_ROOT = Pathname.new("/tmp/xnix-staged-desktop-notepad-smoke-#{VERSION}")
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")

options = {
  run_root: DEFAULT_RUN_ROOT.join(RUN_ID).to_s,
  report_output: DEFAULT_RUN_ROOT.join("staged-desktop-notepad-smoke.json").to_s,
  markdown_output: DEFAULT_RUN_ROOT.join("staged-desktop-notepad-smoke.md").to_s,
  delegated_output: DEFAULT_RUN_ROOT.join("staged-desktop-notepad-delegated-launcher.json").to_s,
  registry: PROJECT_ROOT.join("runtime/recipes/registry.json").to_s,
  image: ENV.fetch("XNIX_WINE_IMAGE", DEFAULT_IMAGE),
  docker: ENV.fetch("XNIX_DOCKER_BIN", "docker"),
  timeout: "120s"
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/staged_desktop_notepad_smoke.rb [options]"
  parser.on("--run-root PATH", "Temporary run root") { |value| options[:run_root] = value }
  parser.on("--report-output PATH", "JSON packet output path") { |value| options[:report_output] = value }
  parser.on("--markdown-output PATH", "Markdown packet output path") { |value| options[:markdown_output] = value }
  parser.on("--delegated-output PATH", "Delegated launcher JSON output path") { |value| options[:delegated_output] = value }
  parser.on("--registry PATH", "Development recipe registry path") { |value| options[:registry] = value }
  parser.on("--image IMAGE", "Local Wine GUI smoke image") { |value| options[:image] = value }
  parser.on("--docker PATH", "Docker runner path") { |value| options[:docker] = value }
  parser.on("--timeout DURATION", "GUI smoke timeout") { |value| options[:timeout] = value }
end.parse!

abort "staged desktop Notepad smoke does not accept positional arguments" unless ARGV.empty?

def absolute_path(path)
  candidate = Pathname.new(path)
  candidate = PROJECT_ROOT.join(candidate) unless candidate.absolute?
  candidate.cleanpath
end

def run_command(env, *argv)
  Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
end

def run_json(env, *argv)
  stdout, stderr, status = run_command(env, *argv)
  abort "#{argv.join(" ")} failed:\n#{stderr}\n#{stdout}" unless status.success?

  [JSON.parse(stdout), stdout]
rescue JSON::ParserError => e
  abort "#{argv.join(" ")} returned malformed JSON: #{e.message}\n#{stdout}"
end

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def assert_no_forbidden(text, forbidden_terms, label)
  downcased = text.downcase
  forbidden_terms.each do |term|
    next if term.to_s.empty?

    assert(!downcased.include?(term.to_s.downcase), "#{label} must not expose #{term}")
  end
end

def docker_image_available?(docker_bin, image)
  system(docker_bin, "image", "inspect", image, out: File::NULL, err: File::NULL)
rescue SystemCallError
  false
end

def resolve_executable(path)
  value = path.to_s.strip
  return nil if value.empty?

  if value.include?(File::SEPARATOR)
    candidate = Pathname.new(value)
    candidate = PROJECT_ROOT.join(candidate) unless candidate.absolute?
    candidate = candidate.cleanpath
    return candidate.to_s if candidate.file? && candidate.executable?

    return nil
  end

  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).each do |directory|
    candidate = File.join(directory, value)
    return candidate if File.file?(candidate) && File.executable?(candidate)
  end

  nil
end

def desktop_exec_from(path)
  line = path.each_line.find { |entry| entry.start_with?("Exec=") }
  abort "staged desktop entry does not contain Exec" if line.nil?

  line.delete_prefix("Exec=").strip
end

run_root = absolute_path(options.fetch(:run_root))
report_output = absolute_path(options.fetch(:report_output))
markdown_output = absolute_path(options.fetch(:markdown_output))
delegated_output = absolute_path(options.fetch(:delegated_output))
build_root = run_root.join("build")
stage_root = run_root.join("stage")
state_root = run_root.join("state")
launcher_bin = build_root.join("xnix-compat-launch")
staged_launcher = stage_root.join("usr/local/bin/xnix-compat-launch")
docker_bin = resolve_executable(options.fetch(:docker))

go_env = {
  "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
  "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
  "GOTMPDIR" => GO_CACHE_ROOT.join("tmp").to_s
}

FileUtils.mkdir_p(build_root)
FileUtils.mkdir_p(stage_root)
FileUtils.mkdir_p(state_root)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("tmp"))
FileUtils.mkdir_p(report_output.dirname)
FileUtils.mkdir_p(markdown_output.dirname)
FileUtils.mkdir_p(delegated_output.dirname)

if docker_bin.nil?
  puts "SKIP: staged desktop Notepad smoke (Docker runner #{options.fetch(:docker)} is unavailable)"
  exit 0
end

unless docker_image_available?(docker_bin, options.fetch(:image))
  puts "SKIP: staged desktop Notepad smoke (Docker image #{options.fetch(:image)} is unavailable)"
  exit 0
end

build_stdout, build_stderr, build_status = run_command(
  go_env,
  "go", "build", "-o", launcher_bin.to_s, "./cmd/xnix-compat-launch"
)
abort "go build failed:\n#{build_stderr}\n#{build_stdout}" unless build_status.success?

stage, stage_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "desktop-activation-stage",
  "--registry", options.fetch(:registry),
  "--app", APP_ID,
  "--mode", "development",
  "--staging-root", stage_root.to_s,
  "--managed-launcher-bin", launcher_bin.to_s
)

written_ids = stage.fetch("written_file_ids")
%w[application-recipe desktop-entry managed-launcher-executable recipe-registry].each do |id|
  assert(written_ids.include?(id), "desktop activation stage must write #{id}")
end
assert(stage.fetch("application_id") == APP_ID, "desktop activation stage must target Notepad")
assert(stage.fetch("host_root_modified") == false, "desktop activation stage must not mutate the host root")
assert(stage.fetch("launch_enabled") == false, "desktop activation stage must keep launch gated")
assert(stage.fetch("execution_started") == false, "desktop activation stage must not start execution")
assert_no_forbidden(stage_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s], "desktop activation stage output")

desktop_artifact = stage.fetch("written_files").find { |entry| entry.fetch("id") == "desktop-entry" }
desktop_path = stage_root.join(desktop_artifact.fetch("relative_path"))
assert(desktop_path.file?, "staged desktop entry must exist")
assert(staged_launcher.file?, "staged managed launcher must exist")

desktop_exec = desktop_exec_from(desktop_path)
desktop_tokens = Shellwords.split(desktop_exec).reject { |token| token == "%U" }
assert(desktop_tokens.first == "xnix-compat-launch", "desktop Exec must use the managed launcher name")
assert(desktop_tokens.include?("--registry"), "desktop Exec must carry the packaged registry argument")
assert(desktop_tokens.include?(PACKAGED_REGISTRY), "desktop Exec must use the packaged recipe registry")

launch_receipt, = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-launch-authorization-receipt-preview",
  "--app", APP_ID,
  "--state-root", state_root.to_s,
  "--authorize", "review-launch-authorization"
)
assert(launch_receipt.fetch("launch_authorization_recorded") == true, "launch authorization receipt must be recorded")
assert(launch_receipt.fetch("host_root_modified") == false, "launch authorization receipt must not mutate the host root")

session_record, = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-controlled-execution-session-record",
  "--app", APP_ID,
  "--state-root", state_root.to_s,
  "--receipt-id", launch_receipt.fetch("receipt_id"),
  "--cache-root", run_root.join("known-app-cache").to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(session_record.fetch("record_state") == "persisted", "controlled execution session must be persisted")
assert(session_record.fetch("execution_session_handoff_created") == true, "controlled execution session must create handoff evidence")
assert(session_record.fetch("host_root_modified") == false, "controlled execution session must not mutate the host root")

review_receipt, = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-session-gated-launch-review-receipt-record",
  "--app", APP_ID,
  "--state-root", state_root.to_s,
  "--session-id", session_record.fetch("execution_session_id"),
  "--action", "review-session-gated-dispatch",
  "--decision", "approved"
)
assert(review_receipt.fetch("review_receipt_recorded") == true, "session-gated review receipt must be recorded")
assert(review_receipt.fetch("session_digest_verified") == true, "session-gated review receipt must verify session evidence")
assert(review_receipt.fetch("host_root_modified") == false, "session-gated review receipt must not mutate the host root")

launcher_env = {
  "XNIX_STAGING_ROOT" => stage_root.to_s,
  "XNIX_DOCKER_BIN" => docker_bin
}
launcher_argv = [
  staged_launcher.to_s,
  *desktop_tokens.drop(1),
  "--guest-boundary", GUEST_BOUNDARY,
  "--state-root", state_root.to_s,
  "--receipt-id", launch_receipt.fetch("receipt_id"),
  "--review-receipt-id", review_receipt.fetch("receipt_id"),
  "--session-id", session_record.fetch("execution_session_id"),
  "--image", options.fetch(:image),
  "--docker", docker_bin,
  "--timeout", options.fetch(:timeout)
]

launcher_stdout, launcher_stderr, launcher_status = run_command(launcher_env, *launcher_argv)
abort "staged launcher failed:\n#{launcher_stderr}\n#{launcher_stdout}" unless launcher_status.success?

payload = JSON.parse(launcher_stdout)
File.write(delegated_output, JSON.pretty_generate(payload) + "\n")

unless payload.fetch("status") == "passed"
  failure_packet = {
    "schema_version" => SCHEMA_VERSION,
    "status" => "failed",
    "app_id" => APP_ID,
    "display_name" => APP_NAME,
    "version" => VERSION,
    "desktop_exec_uses_packaged_registry" => desktop_tokens.include?(PACKAGED_REGISTRY),
    "staged_registry_resolved" => true,
    "delegated_status" => payload.fetch("status"),
    "delegated_skip_reason" => payload["skip_reason"],
    "delegated_launcher_payload_path" => delegated_output.to_s,
    "report_path" => report_output.to_s,
    "markdown_path" => markdown_output.to_s
  }
  File.write(report_output, JSON.pretty_generate(failure_packet) + "\n")
  File.write(
    markdown_output,
    [
      "# Staged Desktop Notepad Smoke",
      "",
      "- Status: failed",
      "- App: #{APP_NAME} (`#{APP_ID}`)",
      "- Version: #{VERSION}",
      "- Delegated status: #{payload.fetch("status")}",
      "- Delegated skip reason: #{payload["skip_reason"]}",
      "- Delegated payload: #{delegated_output}",
      ""
    ].join("\n")
  )
  warn "FAIL: launcher smoke must pass; delegated payload written to #{delegated_output}"
  exit 1
end

assert(payload.fetch("request_type") == "windows-app-container-x-gui-smoke", "launcher must enter the container X GUI smoke")
assert(payload.fetch("application_id") == APP_ID, "launcher smoke must preserve the Notepad app id")
assert(payload.fetch("display_name") == APP_NAME, "launcher smoke must preserve the Notepad display name")
assert(payload.fetch("app_version") == VERSION, "launcher smoke must preserve the current app version")
assert(payload.fetch("recipe_backed") == true, "launcher smoke must be recipe-backed")
assert(payload.fetch("x_window_observed") == true, "launcher smoke must observe a Notepad X window")
assert(payload.fetch("network_mode") == "none", "launcher smoke must disable container networking")
assert(payload.fetch("host_mount_count") == 0, "launcher smoke must not mount host directories")
assert(payload.fetch("docker_socket_mounted") == false, "launcher smoke must not mount the Docker socket")
assert(payload.fetch("host_networking_required") == false, "launcher smoke must not require host networking")
assert(payload.fetch("broad_host_mount_required") == false, "launcher smoke must not require broad host mounts")
assert(payload.fetch("host_root_modified") == false, "launcher smoke must not mutate the host root")
assert(payload.fetch("session_gated_controlled_dispatch_consumed") == true, "launcher smoke must consume the session-gated dispatch")
assert(payload.fetch("controlled_execution_session_consumed") == true, "launcher smoke must consume the controlled execution session")
assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s], "staged launcher output")

packet = {
  "schema_version" => SCHEMA_VERSION,
  "status" => "passed",
  "app_id" => APP_ID,
  "display_name" => APP_NAME,
  "version" => VERSION,
  "desktop_exec_uses_packaged_registry" => desktop_tokens.include?(PACKAGED_REGISTRY),
  "staged_registry_resolved" => true,
  "window_observed" => payload.fetch("x_window_observed"),
  "container_platform" => payload.fetch("container_platform"),
  "network_mode" => payload.fetch("network_mode"),
  "host_mount_count" => payload.fetch("host_mount_count"),
  "host_root_modified" => payload.fetch("host_root_modified"),
  "docker_socket_mounted" => payload.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => payload.fetch("broad_host_mount_required"),
  "delegated_launcher_payload_path" => delegated_output.to_s,
  "report_path" => report_output.to_s,
  "markdown_path" => markdown_output.to_s
}

File.write(report_output, JSON.pretty_generate(packet) + "\n")
File.write(
  markdown_output,
  [
    "# Staged Desktop Notepad Smoke",
    "",
    "- Status: passed",
    "- App: #{APP_NAME} (`#{APP_ID}`)",
    "- Version: #{VERSION}",
    "- Desktop Exec uses packaged registry: true",
    "- X window observed: #{payload.fetch("x_window_observed")}",
    "- Container platform: #{payload.fetch("container_platform")}",
    "- Network mode: #{payload.fetch("network_mode")}",
    "- Host mount count: #{payload.fetch("host_mount_count")}",
    "- Docker socket mounted: #{payload.fetch("docker_socket_mounted")}",
    "- Host root modified: #{payload.fetch("host_root_modified")}",
    ""
  ].join("\n")
)

puts "PASS: staged desktop Notepad smoke"
puts JSON.pretty_generate(packet)
