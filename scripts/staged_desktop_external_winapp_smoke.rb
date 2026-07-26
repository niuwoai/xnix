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
SCHEMA_VERSION = "xnix.scripts.staged_desktop_external_winapp_smoke.v1"
APP_ID = "org.xnix.external.desktop-notepad"
APP_NAME = "External Desktop Notepad"
DEFAULT_IMAGE = "xnix-wine-smoke:local"
DEFAULT_RUN_ROOT = Pathname.new("/tmp/xnix-staged-desktop-external-winapp-smoke-#{VERSION}")
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
CONTAINER_NOTEPAD_CANDIDATES = [
  "/usr/lib/wine/x86_64-windows/notepad.exe",
  "/usr/lib/wine/i386-windows/notepad.exe",
  "/usr/lib/aarch64-linux-gnu/wine/aarch64-windows/notepad.exe",
  "/usr/lib/x86_64-linux-gnu/wine/x86_64-windows/notepad.exe",
  "/usr/lib/i386-linux-gnu/wine/i386-windows/notepad.exe",
  "/usr/lib/wine/notepad.exe",
  "/usr/share/wine/notepad.exe"
].freeze

options = {
  run_root: DEFAULT_RUN_ROOT.join(RUN_ID).to_s,
  report_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-smoke.json").to_s,
  markdown_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-smoke.md").to_s,
  delegated_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-delegated-launcher.json").to_s,
  activation_status_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-activation-status.json").to_s,
  runtime_packet_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-real-gui-packet.json").to_s,
  kde_page_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-kde-page.json").to_s,
  executable: "",
  image: ENV.fetch("XNIX_WINE_IMAGE", DEFAULT_IMAGE),
  docker: ENV.fetch("XNIX_DOCKER_BIN", "docker"),
  timeout: "120s"
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/staged_desktop_external_winapp_smoke.rb [options]"
  parser.on("--run-root PATH", "Temporary run root") { |value| options[:run_root] = value }
  parser.on("--report-output PATH", "JSON packet output path") { |value| options[:report_output] = value }
  parser.on("--markdown-output PATH", "Markdown packet output path") { |value| options[:markdown_output] = value }
  parser.on("--delegated-output PATH", "Delegated launcher JSON output path") { |value| options[:delegated_output] = value }
  parser.on("--activation-status-output PATH", "Go Runtime activation status JSON output path") { |value| options[:activation_status_output] = value }
  parser.on("--runtime-packet-output PATH", "Go Runtime real Windows GUI packet output path") { |value| options[:runtime_packet_output] = value }
  parser.on("--kde-page-output PATH", "KDE center page JSON output path") { |value| options[:kde_page_output] = value }
  parser.on("--executable PATH", "Optional existing Windows GUI executable; defaults to Wine Notepad from the local image") { |value| options[:executable] = value }
  parser.on("--image IMAGE", "Local Wine GUI smoke image") { |value| options[:image] = value }
  parser.on("--docker PATH", "Docker runner path") { |value| options[:docker] = value }
  parser.on("--timeout DURATION", "GUI smoke timeout") { |value| options[:timeout] = value }
end.parse!

abort "staged desktop external Windows app smoke does not accept positional arguments" unless ARGV.empty?

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

def docker_image_available?(docker_bin, image)
  system(docker_bin, "image", "inspect", image, out: File::NULL, err: File::NULL)
rescue SystemCallError
  false
end

def container_notepad_path(docker_bin, image)
  script = CONTAINER_NOTEPAD_CANDIDATES.map { |path| "test -f #{Shellwords.escape(path)} && printf '%s\\n' #{Shellwords.escape(path)} && exit 0" }.join("; ")
  stdout, _stderr, status = Open3.capture3(docker_bin, "run", "--rm", "--network", "none", "--entrypoint", "sh", image, "-lc", "#{script}; exit 1")
  return nil unless status.success?

  stdout.lines.first&.strip
end

def copy_container_notepad(docker_bin, image, source_path, destination_path)
  stdout, stderr, status = Open3.capture3(docker_bin, "create", "--network", "none", "--entrypoint", "sh", image, "-lc", "sleep 30")
  abort "docker create for Notepad extraction failed:\n#{stderr}\n#{stdout}" unless status.success?

  container_id = stdout.strip
  begin
    _cp_stdout, cp_stderr, cp_status = Open3.capture3(docker_bin, "cp", "#{container_id}:#{source_path}", destination_path.to_s)
    abort "docker cp for Notepad extraction failed:\n#{cp_stderr}" unless cp_status.success?
  ensure
    Open3.capture3(docker_bin, "rm", "-f", container_id) unless container_id.empty?
  end
end

def desktop_exec_from(path)
  line = path.each_line.find { |entry| entry.start_with?("Exec=") }
  abort "staged desktop entry does not contain Exec" if line.nil?

  line.delete_prefix("Exec=").strip
end

def write_skip_report(report_output, markdown_output, reason)
  packet = {
    "schema_version" => SCHEMA_VERSION,
    "status" => "skipped",
    "app_id" => APP_ID,
    "display_name" => APP_NAME,
    "version" => VERSION,
    "skip_reason" => reason,
    "report_path" => report_output.to_s,
    "markdown_path" => markdown_output.to_s
  }
  File.write(report_output, JSON.pretty_generate(packet) + "\n")
  File.write(
    markdown_output,
    [
      "# Staged Desktop External Windows App Smoke",
      "",
      "- Status: skipped",
      "- App: #{APP_NAME} (`#{APP_ID}`)",
      "- Version: #{VERSION}",
      "- Skip reason: #{reason}",
      ""
    ].join("\n")
  )
  puts "SKIP: staged desktop external Windows app smoke (#{reason})"
end

def write_markdown_report(markdown_output, packet)
  File.write(
    markdown_output,
    [
      "# Staged Desktop External Windows App Smoke",
      "",
      "- Status: #{packet.fetch("status")}",
      "- App: #{packet.fetch("display_name")} (`#{packet.fetch("app_id")}`)",
      "- Version: #{packet.fetch("version")}",
      "- Desktop Exec uses external app handle: #{packet.fetch("desktop_exec_uses_external_app_handle")}",
      "- External app desktop handle ready: #{packet.fetch("external_app_desktop_handle_ready")}",
      "- Activation receipt desktop handle ready: #{packet.fetch("activation_receipt_external_app_desktop_handle_ready")}",
      "- External import record consumed: #{packet.fetch("external_app_import_record_consumed")}",
      "- External app handle consumed: #{packet.fetch("external_app_handle_consumed")}",
      "- Runtime packet external app handle consumed: #{packet.fetch("runtime_packet_external_app_handle_consumed")}",
      "- KDE card external app handle consumed: #{packet.fetch("kde_page_card_external_app_handle_consumed")}",
      "- Imported artifact digest verified: #{packet.fetch("imported_artifact_digest_verified")}",
      "- X window observed: #{packet.fetch("window_observed")}",
      "- Container network mode: #{packet.fetch("container_network_mode")}",
      "- Container host mount count: #{packet.fetch("container_host_mount_count")}",
      "- Docker socket mounted: #{packet.fetch("docker_socket_mounted")}",
      "- Broad host mount required: #{packet.fetch("broad_host_mount_required")}",
      "- Host root modified: #{packet.fetch("host_root_modified")}",
      "- Runtime packet consumed report: #{packet.fetch("runtime_packet_consumed_report")}",
      "- KDE GUI evidence count: #{packet.fetch("kde_page_known_app_gui_evidence_count")}",
      "- Delegated launcher payload: #{packet.fetch("delegated_launcher_payload_path")}",
      "- Runtime packet: #{packet.fetch("runtime_packet_path")}",
      "- Activation status: #{packet.fetch("activation_status_path")}",
      "- KDE page: #{packet.fetch("kde_page_path")}",
      "- KDE card window observed: #{packet.fetch("kde_page_card_window_observed")}",
      "- KDE card X window observed: #{packet.fetch("kde_page_card_x_window_observed")}",
      ""
    ].join("\n")
  )
end

run_root = absolute_path(options.fetch(:run_root))
report_output = absolute_path(options.fetch(:report_output))
markdown_output = absolute_path(options.fetch(:markdown_output))
delegated_output = absolute_path(options.fetch(:delegated_output))
activation_status_output = absolute_path(options.fetch(:activation_status_output))
runtime_packet_output = absolute_path(options.fetch(:runtime_packet_output))
kde_page_output = absolute_path(options.fetch(:kde_page_output))
build_root = run_root.join("build")
stage_root = run_root.join("stage")
state_root = run_root.join("state")
launcher_bin = build_root.join("xnix-compat-launch")
staged_launcher = stage_root.join("usr/local/bin/xnix-compat-launch")
executable_path = options.fetch(:executable).to_s.strip.empty? ? run_root.join("notepad.exe") : absolute_path(options.fetch(:executable))
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
FileUtils.mkdir_p(activation_status_output.dirname)
FileUtils.mkdir_p(runtime_packet_output.dirname)
FileUtils.mkdir_p(kde_page_output.dirname)

if docker_bin.nil?
  write_skip_report(report_output, markdown_output, "Docker runner #{options.fetch(:docker)} is unavailable")
  exit 0
end

unless docker_image_available?(docker_bin, options.fetch(:image))
  write_skip_report(report_output, markdown_output, "Docker image #{options.fetch(:image)} is unavailable")
  exit 0
end

if options.fetch(:executable).to_s.strip.empty?
  notepad_path = container_notepad_path(docker_bin, options.fetch(:image))
  if notepad_path.nil?
    write_skip_report(report_output, markdown_output, "Wine Notepad was not found in #{options.fetch(:image)}")
    exit 0
  end
  copy_container_notepad(docker_bin, options.fetch(:image), notepad_path, executable_path)
end

assert(executable_path.file?, "external Windows executable must exist")

build_stdout, build_stderr, build_status = run_command(
  go_env,
  "go", "build", "-o", launcher_bin.to_s, "./cmd/xnix-compat-launch"
)
abort "go build failed:\n#{build_stderr}\n#{build_stdout}" unless build_status.success?

import_record, import_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "external-winapp-import-record",
  "--state-root", state_root.to_s,
  "--executable", executable_path.to_s,
  "--app-id", APP_ID,
  "--display-name", APP_NAME
)
assert(import_record.fetch("import_recorded") == true, "external app import record must be persisted")
assert(import_record.fetch("artifact_copied") == true, "external app import must copy the artifact")
assert(import_record.fetch("windows_executable_validated") == true, "external app import must validate an MZ executable")
assert(import_record.fetch("host_root_modified") == false, "external app import must not mutate the host root")
assert_no_forbidden(import_stdout, [state_root.to_s, executable_path.to_s], "external import output")

import_record_path = state_root.join(import_record.fetch("record_relative_path"))
stage, stage_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "desktop-activation-stage",
  "--external-app-import-record", import_record_path.to_s,
  "--mode", "development",
  "--staging-root", stage_root.to_s,
  "--managed-launcher-bin", launcher_bin.to_s
)
written_ids = stage.fetch("written_file_ids")
%w[desktop-entry managed-launcher-artifact managed-launcher-executable].each do |id|
  assert(written_ids.include?(id), "desktop activation stage must write #{id}")
end
assert(stage.fetch("application_id") == APP_ID, "desktop activation stage must target the imported app")
assert(stage.fetch("external_app_handle") == APP_ID, "desktop activation stage must expose the opaque external app handle")
assert(stage.fetch("desktop_exec_uses_external_app_handle") == true, "desktop activation stage must prove the desktop Exec uses the external app handle")
assert(stage.fetch("external_app_desktop_handle_ready") == true, "desktop activation stage must mark external app desktop handle readiness")
assert(stage.fetch("desktop_exec_uses_raw_import_record") == false, "desktop activation stage must not expose raw import-record Exec routing")
assert(stage.fetch("desktop_exec_uses_state_root") == false, "desktop activation stage must not expose state-root Exec routing")
assert(stage.fetch("launch_enabled") == false, "desktop activation stage must keep launch gated")
assert(stage.fetch("execution_started") == false, "desktop activation stage must not start execution")
assert(stage.fetch("host_root_modified") == false, "desktop activation stage must not mutate the host root")
assert_no_forbidden(stage_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s], "desktop activation stage output")

activation_status, status_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "desktop-activation-status-preview",
  "--external-app-import-record", import_record_path.to_s,
  "--mode", "development",
  "--activation-root", stage_root.to_s
)
File.write(activation_status_output, JSON.pretty_generate(activation_status) + "\n")
receipt_evidence = activation_status.fetch("receipt_evidence")
assert(receipt_evidence.fetch("external_app_handle") == APP_ID, "activation receipt evidence must preserve the opaque external app handle")
assert(receipt_evidence.fetch("desktop_exec_uses_external_app_handle") == true, "activation receipt evidence must prove external app handle Exec routing")
assert(receipt_evidence.fetch("external_app_desktop_handle_ready") == true, "activation receipt evidence must mark external desktop handle readiness")
assert(receipt_evidence.fetch("desktop_exec_uses_raw_import_record") == false, "activation receipt evidence must not persist raw import-record Exec routing")
assert(receipt_evidence.fetch("desktop_exec_uses_state_root") == false, "activation receipt evidence must not persist state-root Exec routing")
assert(receipt_evidence.fetch("safe_for_kde") == true, "activation receipt evidence must stay safe for KDE consumption")
assert(activation_status_output.file?, "activation status file must be written")
assert_no_forbidden(status_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s], "desktop activation status output")

desktop_artifact = stage.fetch("written_files").find { |entry| entry.fetch("id") == "desktop-entry" }
desktop_path = stage_root.join(desktop_artifact.fetch("relative_path"))
assert(desktop_path.file?, "staged desktop entry must exist")
assert(staged_launcher.file?, "staged managed launcher must exist")

desktop_exec = desktop_exec_from(desktop_path)
desktop_tokens = Shellwords.split(desktop_exec).reject { |token| token == "%U" }
assert(desktop_tokens == ["xnix-compat-launch", "--external-app-handle", APP_ID], "desktop Exec must expose only the external app handle")
assert_no_forbidden(desktop_exec, [state_root.to_s, import_record_path.to_s, executable_path.to_s, "notepad.exe", "wine ", "docker", "qemu-system", "--state-root", "--external-app-import-record"], "desktop Exec")

launcher_env = {
  "XNIX_EXTERNAL_APP_STATE_ROOT" => state_root.to_s,
  "XNIX_DOCKER_BIN" => docker_bin
}
launcher_argv = [
  staged_launcher.to_s,
  *desktop_tokens.drop(1),
  "--image", options.fetch(:image),
  "--docker", docker_bin,
  "--timeout", options.fetch(:timeout)
]

launcher_stdout, launcher_stderr, launcher_status = run_command(launcher_env, *launcher_argv)
abort "staged external app launcher failed:\n#{launcher_stderr}\n#{launcher_stdout}" unless launcher_status.success?

payload = JSON.parse(launcher_stdout)
File.write(delegated_output, JSON.pretty_generate(payload) + "\n")
assert(payload.fetch("request_type") == "windows-external-app-run", "launcher must enter the external app Runtime run")
assert(payload.fetch("status") == "passed", "launcher smoke must pass")
assert(payload.fetch("application_id") == APP_ID, "launcher smoke must preserve imported app id")
assert(payload.fetch("external_app_import_record_consumed") == true, "launcher smoke must consume the import record")
assert(payload.fetch("external_app_handle_consumed") == true, "launcher smoke must consume the desktop handle")
assert(payload.fetch("imported_artifact_digest_verified") == true, "launcher smoke must verify the imported artifact digest")
assert(payload.fetch("x_window_observed") == true, "launcher smoke must observe a Windows GUI X window")
assert(payload.fetch("window_observed") == true, "launcher smoke must expose generic observed-window evidence")
assert(payload.fetch("container_network_mode") == "none", "launcher smoke must disable container networking")
assert(payload.fetch("container_host_mount_count") == 0, "launcher smoke must not mount host directories")
assert(payload.fetch("docker_socket_mounted") == false, "launcher smoke must not mount the Docker socket")
assert(payload.fetch("host_networking_required") == false, "launcher smoke must not require host networking")
assert(payload.fetch("broad_host_mount_required") == false, "launcher smoke must not require broad host mounts")
assert(payload.fetch("host_root_modified") == false, "launcher smoke must not mutate the host root")
assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, docker_bin], "staged launcher output")

runtime_packet, = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "real-winapp-gui-evidence-packet-preview",
  "--gui-smoke-report", delegated_output.to_s,
  "--output", runtime_packet_output.to_s
)
assert(runtime_packet.fetch("request_type") == "real-winapp-gui-evidence-packet-preview", "Runtime packet must use the real GUI packet request type")
assert(runtime_packet.fetch("report_consumed") == true, "Runtime packet must consume the delegated report")
assert(runtime_packet.fetch("external_app_run_record_consumed") == true, "Runtime packet must consume the external app run record")
assert(runtime_packet.fetch("external_app_handle_consumed") == true, "Runtime packet must preserve external app handle consumption")
assert(runtime_packet.fetch("external_app_import_record_consumed") == true, "Runtime packet must preserve import-record consumption")
assert(runtime_packet.fetch("imported_artifact_digest_verified") == true, "Runtime packet must preserve imported artifact digest verification")
assert(runtime_packet.fetch("known_app_gui_evidence_verified_count") == 1, "Runtime packet must verify one GUI evidence item")
assert(runtime_packet.fetch("container_network_mode") == "none", "Runtime packet must preserve network isolation")
assert(runtime_packet.fetch("container_host_mount_count") == 0, "Runtime packet must preserve zero host mounts")
assert(runtime_packet.fetch("backend_launch_enabled") == false, "Runtime packet must not enable backend launch")
assert(runtime_packet.fetch("host_root_modified") == false, "Runtime packet must not mutate the host root")
assert(runtime_packet_output.file?, "Runtime packet file must be written")

kde_page, kde_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "kde-center-page-preview",
  "--external-app-evidence-file", runtime_packet_output.to_s,
  "--decision", "approved"
)
File.write(kde_page_output, JSON.pretty_generate(kde_page) + "\n")
cards = kde_page.fetch("known_app_gui_evidence_cards")
assert(kde_page.fetch("application_id") == APP_ID, "KDE page must target the imported app")
assert(kde_page.fetch("known_app_gui_evidence_count") == 1, "KDE page must consume one real GUI evidence item")
assert(cards.length == 1, "KDE page must render one real GUI evidence card")
assert(cards.first.fetch("app_id") == APP_ID, "KDE GUI evidence card must target the imported app")
assert(cards.first.fetch("external_app_run_record_consumed") == true, "KDE GUI evidence card must preserve external run consumption")
assert(cards.first.fetch("external_app_handle_consumed") == true, "KDE GUI evidence card must preserve external app handle consumption")
assert(cards.first.fetch("external_app_import_record_consumed") == true, "KDE GUI evidence card must preserve import-record consumption")
assert(cards.first.fetch("imported_artifact_digest_verified") == true, "KDE GUI evidence card must preserve digest verification")
assert(cards.first.fetch("window_observed") == true, "KDE GUI evidence card must preserve generic observed-window evidence")
assert(cards.first.fetch("x_window_observed") == true, "KDE GUI evidence card must preserve X observed-window evidence")
assert(cards.first.fetch("desktop_launch_enabled") == false, "KDE GUI evidence card must not enable desktop launch")
assert(cards.first.fetch("backend_launch_enabled") == false, "KDE GUI evidence card must not enable backend launch")
assert(cards.first.fetch("backend_details_exposed") == false, "KDE GUI evidence card must not expose backend details")
assert(cards.first.fetch("host_root_modified") == false, "KDE GUI evidence card must not mutate the host root")
assert(kde_page_output.file?, "KDE page file must be written")
assert_no_forbidden(kde_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s], "KDE page output")

packet = {
  "schema_version" => SCHEMA_VERSION,
  "status" => "passed",
  "app_id" => APP_ID,
  "display_name" => APP_NAME,
  "version" => VERSION,
  "desktop_exec_uses_external_app_handle" => stage.fetch("desktop_exec_uses_external_app_handle"),
  "external_app_desktop_handle_ready" => stage.fetch("external_app_desktop_handle_ready"),
  "activation_receipt_external_app_desktop_handle_ready" => receipt_evidence.fetch("external_app_desktop_handle_ready"),
  "activation_receipt_safe_for_kde" => receipt_evidence.fetch("safe_for_kde"),
  "external_app_import_record_consumed" => payload.fetch("external_app_import_record_consumed"),
  "external_app_handle_consumed" => payload.fetch("external_app_handle_consumed"),
  "imported_artifact_digest_verified" => payload.fetch("imported_artifact_digest_verified"),
  "window_observed" => payload.fetch("window_observed"),
  "x_window_observed" => payload.fetch("x_window_observed"),
  "container_network_mode" => payload.fetch("container_network_mode"),
  "container_host_mount_count" => payload.fetch("container_host_mount_count"),
  "host_root_modified" => payload.fetch("host_root_modified"),
  "docker_socket_mounted" => payload.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => payload.fetch("broad_host_mount_required"),
  "runtime_packet_output_written" => runtime_packet_output.file?,
  "runtime_packet_consumed_report" => runtime_packet.fetch("report_consumed"),
  "runtime_packet_external_app_run_record_consumed" => runtime_packet.fetch("external_app_run_record_consumed"),
  "runtime_packet_external_app_handle_consumed" => runtime_packet.fetch("external_app_handle_consumed"),
  "runtime_packet_known_app_gui_evidence_verified_count" => runtime_packet.fetch("known_app_gui_evidence_verified_count"),
  "kde_page_output_written" => kde_page_output.file?,
  "kde_page_known_app_gui_evidence_count" => kde_page.fetch("known_app_gui_evidence_count"),
  "kde_page_card_external_app_run_record_consumed" => cards.first.fetch("external_app_run_record_consumed"),
  "kde_page_card_external_app_handle_consumed" => cards.first.fetch("external_app_handle_consumed"),
  "kde_page_card_window_observed" => cards.first.fetch("window_observed"),
  "kde_page_card_x_window_observed" => cards.first.fetch("x_window_observed"),
  "delegated_launcher_payload_path" => delegated_output.to_s,
  "runtime_packet_path" => runtime_packet_output.to_s,
  "activation_status_path" => activation_status_output.to_s,
  "kde_page_path" => kde_page_output.to_s,
  "report_path" => report_output.to_s,
  "markdown_path" => markdown_output.to_s
}

File.write(report_output, JSON.pretty_generate(packet) + "\n")
write_markdown_report(markdown_output, packet)

puts "PASS: staged desktop external Windows app smoke"
puts JSON.pretty_generate(packet)
