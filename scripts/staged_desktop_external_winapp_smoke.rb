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
DEFAULT_APP_ID = "org.xnix.external.desktop-notepad"
DEFAULT_APP_NAME = "External Desktop Notepad"
DEFAULT_OPERATOR_APP_ID = "org.xnix.external.operator-run"
DEFAULT_OPERATOR_APP_NAME = "External Windows App"
NOTEPAD_FILE_ARGUMENT_APP_ID = "org.xnix.external.desktop-notepad-file-argument"
NOTEPAD_FILE_ARGUMENT_APP_NAME = "External Desktop Notepad File Argument"
NOTEPAD_FILE_ARGUMENT_WINDOW_MATCH = "sample-document.txt"
DEFAULT_IMAGE = "xnix-wine-smoke:local"
DEFAULT_RUN_ROOT = Pathname.new("/tmp/xnix-staged-desktop-external-winapp-smoke-#{VERSION}")
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
ALLOW_LOCAL_GO_COMPILE_ENV = "XNIX_ALLOW_LOCAL_GO_COMPILE"
REMOTE_GO_BUILD_HINT = "scripts/remote_go_build.rb --execute"
RUNTIME_GO_BIN_ENV = "XNIX_RUNTIME_GO_BIN"
COMPAT_LAUNCH_BIN_ENV = "XNIX_COMPAT_LAUNCH_BIN"
RUNTIME_IMAGE_ENV = "XNIX_RUNTIME_IMAGE"
RUNTIME_GO_IMAGE_PATH = "/usr/local/bin/xnix-runtime-go"
COMPAT_LAUNCH_IMAGE_PATH = "/usr/local/bin/xnix-compat-launch"
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
  launch_packet_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-launch-packet.json").to_s,
  one_shot_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-one-shot.json").to_s,
  one_shot_launch_packet_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-one-shot-launch-packet.json").to_s,
  runtime_packet_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-real-gui-packet.json").to_s,
  kde_page_output: DEFAULT_RUN_ROOT.join("staged-desktop-external-winapp-kde-page.json").to_s,
  executable: "",
  bundle_root: "",
  executable_relative_path: "",
  import_record: "",
  app_id: "",
  display_name: "",
  window_match: "",
  desktop_argument_mode: ENV.fetch("XNIX_EXTERNAL_APP_DESKTOP_ARGUMENT_MODE", "file-uri"),
  fixture: "notepad",
  image: ENV.fetch("XNIX_WINE_IMAGE", DEFAULT_IMAGE),
  runtime_image: ENV.fetch(RUNTIME_IMAGE_ENV, "xnix-builder:#{VERSION}"),
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
  parser.on("--launch-packet-output PATH", "Go Runtime desktop external Windows app launch packet output path") { |value| options[:launch_packet_output] = value }
  parser.on("--one-shot-output PATH", "Go Runtime one-shot import-stage-and-launch output path") { |value| options[:one_shot_output] = value }
  parser.on("--one-shot-launch-packet-output PATH", "Go Runtime one-shot desktop launch packet output path") { |value| options[:one_shot_launch_packet_output] = value }
  parser.on("--runtime-packet-output PATH", "Go Runtime real Windows GUI packet output path") { |value| options[:runtime_packet_output] = value }
  parser.on("--kde-page-output PATH", "KDE center page JSON output path") { |value| options[:kde_page_output] = value }
  parser.on("--executable PATH", "Optional existing Windows GUI executable; defaults to Wine Notepad from the local image") { |value| options[:executable] = value }
  parser.on("--bundle-root PATH", "Optional portable Windows app directory to import for the external fixture") { |value| options[:bundle_root] = value }
  parser.on("--executable-relative-path PATH", "Portable bundle relative path to the Windows .exe") { |value| options[:executable_relative_path] = value }
  parser.on("--import-record PATH", "Optional existing Runtime external Windows app import record to stage and launch") { |value| options[:import_record] = value }
  parser.on("--app-id ID", "Application id for an operator-supplied external executable") { |value| options[:app_id] = value }
  parser.on("--display-name NAME", "Display name for an operator-supplied external executable") { |value| options[:display_name] = value }
  parser.on("--window-match TEXT", "Optional observed-window text required for the external executable") { |value| options[:window_match] = value }
  parser.on("--desktop-argument-mode MODE", "Desktop argument mode: file-uri or none") { |value| options[:desktop_argument_mode] = value }
  parser.on("--fixture NAME", "Executable fixture: notepad, notepad-file-argument, or external") { |value| options[:fixture] = value }
  parser.on("--image IMAGE", "Local Wine GUI smoke image") { |value| options[:image] = value }
  parser.on("--runtime-image IMAGE", "Runtime image with prebuilt Xnix command binaries") { |value| options[:runtime_image] = value }
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

def command_available?(name)
  !resolve_executable(name).nil?
end

def runtime_go_command
  configured = ENV.fetch(RUNTIME_GO_BIN_ENV, "").strip
  configured_path = resolve_executable(configured)
  return [configured_path] unless configured_path.nil?

  return ["xnix-runtime-go"] if command_available?("xnix-runtime-go")
  return ["go", "run", "./cmd/xnix-runtime-go"] if ENV.fetch(ALLOW_LOCAL_GO_COMPILE_ENV, "") == "1" && command_available?("go")

  nil
end

def launcher_binary
  configured = ENV.fetch(COMPAT_LAUNCH_BIN_ENV, "").strip
  configured_path = resolve_executable(configured)
  return configured_path unless configured_path.nil?

  resolve_executable("xnix-compat-launch")
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

def copy_image_binary(docker_bin, image, source_path, destination_path)
  stdout, _stderr, status = Open3.capture3(docker_bin, "create", "--network", "none", "--entrypoint", "sh", image, "-lc", "sleep 30")
  return false unless status.success?

  container_id = stdout.strip
  begin
    _cp_stdout, _cp_stderr, cp_status = Open3.capture3(docker_bin, "cp", "#{container_id}:#{source_path}", destination_path.to_s)
    return false unless cp_status.success?

    FileUtils.chmod(0o755, destination_path)
    true
  ensure
    Open3.capture3(docker_bin, "rm", "-f", container_id) unless container_id.empty?
  end
end

def desktop_exec_from(path)
  line = path.each_line.find { |entry| entry.start_with?("Exec=") }
  abort "staged desktop entry does not contain Exec" if line.nil?

  line.delete_prefix("Exec=").strip
end

def state_root_from_import_record_path(record_path, record)
  relative_path = Pathname.new(record.fetch("record_relative_path")).cleanpath
  abort "external import record has an unsafe record_relative_path" if relative_path.absolute? || relative_path.each_filename.any? { |part| part == ".." }

  state_root = record_path.dirname
  relative_path.dirname.each_filename { |_part| state_root = state_root.dirname }
  expected = state_root.join(relative_path).cleanpath
  abort "external import record path does not match record_relative_path" unless expected == record_path.cleanpath

  state_root
end

def write_skip_report(report_output, markdown_output, reason, app_id, app_name)
  packet = {
    "schema_version" => SCHEMA_VERSION,
    "status" => "skipped",
    "app_id" => app_id,
    "display_name" => app_name,
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
      "- App: #{app_name} (`#{app_id}`)",
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
      "- Fixture: #{packet.fetch("fixture")}",
      "- App: #{packet.fetch("display_name")} (`#{packet.fetch("app_id")}`)",
      "- Version: #{packet.fetch("version")}",
      "- Desktop Exec uses external app handle: #{packet.fetch("desktop_exec_uses_external_app_handle")}",
      "- External app desktop handle ready: #{packet.fetch("external_app_desktop_handle_ready")}",
      "- Activation receipt desktop handle ready: #{packet.fetch("activation_receipt_external_app_desktop_handle_ready")}",
      "- Desktop Exec invocation exact: #{packet.fetch("desktop_exec_invocation_exact")}",
      "- Launcher context from environment: #{packet.fetch("launcher_context_from_environment")}",
      "- Launcher extra arguments appended: #{packet.fetch("launcher_extra_arguments_appended")}",
      "- External desktop argument count: #{packet.fetch("external_desktop_argument_count")}",
      "- External file URI arguments accepted: #{packet.fetch("external_file_uri_arguments_accepted")}",
      "- External file-open requested: #{packet.fetch("external_file_open_requested")}",
      "- External file bridge copy enabled: #{packet.fetch("external_file_bridge_copy_enabled")}",
      "- External file bridge copied count: #{packet.fetch("external_file_bridge_copied_count")}",
      "- External file bridge arguments passed: #{packet.fetch("external_file_bridge_arguments_passed")}",
      "- External file bridge argument observed count: #{packet.fetch("external_file_bridge_argument_observed_count")}",
      "- External file bridge Wine path translated: #{packet.fetch("external_file_bridge_winepath_translated")}",
      "- External file bridge Wine path translated count: #{packet.fetch("external_file_bridge_winepath_translated_count")}",
      "- External file bridge ready: #{packet.fetch("external_file_bridge_ready")}",
      "- Windows process file-argument window observed: #{packet.fetch("windows_process_file_argument_window_observed")}",
      "- External file bridge mount enabled: #{packet.fetch("external_file_bridge_mount_enabled")}",
      "- Raw file URI arguments exposed: #{packet.fetch("raw_file_uri_arguments_exposed")}",
      "- Desktop launch packet ready: #{packet.fetch("desktop_launch_packet_ready")}",
      "- Desktop launch packet safe for KDE: #{packet.fetch("desktop_launch_packet_safe_for_kde")}",
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
      "- Desktop launch packet: #{packet.fetch("desktop_launch_packet_path")}",
      "- One-shot Runtime launch packet: #{packet.fetch("one_shot_path")}",
      "- One-shot Runtime launch passed: #{packet.fetch("one_shot_status")}",
      "- One-shot staged launcher invoked: #{packet.fetch("one_shot_staged_launcher_invoked")}",
      "- One-shot desktop launch packet written: #{packet.fetch("one_shot_desktop_launch_packet_written")}",
      "- KDE page: #{packet.fetch("kde_page_path")}",
      "- Runtime image: #{packet.fetch("runtime_image")}",
      "- Runtime binary from Runtime image: #{packet.fetch("runtime_binary_from_runtime_image")}",
      "- Launcher binary from Runtime image: #{packet.fetch("launcher_binary_from_runtime_image")}",
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
launch_packet_output = absolute_path(options.fetch(:launch_packet_output))
one_shot_output = absolute_path(options.fetch(:one_shot_output))
one_shot_launch_packet_output = absolute_path(options.fetch(:one_shot_launch_packet_output))
runtime_packet_output = absolute_path(options.fetch(:runtime_packet_output))
kde_page_output = absolute_path(options.fetch(:kde_page_output))
build_root = run_root.join("build")
stage_root = run_root.join("stage")
state_root = run_root.join("state")
one_shot_stage_root = run_root.join("one-shot-stage")
one_shot_state_root = run_root.join("one-shot-state")
sample_document_path = run_root.join("sample-document.txt")
sample_document_uri = "file://#{sample_document_path}"
launcher_bin = build_root.join("xnix-compat-launch")
staged_launcher = stage_root.join("usr/local/bin/xnix-compat-launch")
fixture = options.fetch(:fixture).to_s.strip
bundle_root_option = options.fetch(:bundle_root).to_s.strip
executable_relative_path = options.fetch(:executable_relative_path).to_s.strip
bundle_mode = !bundle_root_option.empty? || !executable_relative_path.empty?
import_record_option = options.fetch(:import_record).to_s.strip
record_mode = !import_record_option.empty?
app_id = DEFAULT_APP_ID
app_name = DEFAULT_APP_NAME
window_match = ""
case fixture
when "notepad"
  abort "notepad fixture does not accept --import-record" if record_mode
  abort "notepad fixture does not accept --bundle-root" if bundle_mode

  executable_path = options.fetch(:executable).to_s.strip.empty? ? run_root.join("notepad.exe") : absolute_path(options.fetch(:executable))
when "notepad-file-argument"
  abort "notepad-file-argument fixture does not accept --import-record" if record_mode
  abort "notepad-file-argument fixture does not accept --bundle-root" if bundle_mode

  app_id = NOTEPAD_FILE_ARGUMENT_APP_ID
  app_name = NOTEPAD_FILE_ARGUMENT_APP_NAME
  window_match = NOTEPAD_FILE_ARGUMENT_WINDOW_MATCH
  executable_path = options.fetch(:executable).to_s.strip.empty? ? run_root.join("notepad.exe") : absolute_path(options.fetch(:executable))
when "external"
  executable_supplied = !options.fetch(:executable).to_s.strip.empty?
  source_count = [record_mode, executable_supplied, bundle_mode].count(true)
  abort "external fixture accepts exactly one source: --import-record, --executable, or --bundle-root with --executable-relative-path" if source_count > 1
  abort "external fixture requires --import-record, --executable, or --bundle-root with --executable-relative-path" if source_count.zero?
  abort "external fixture portable bundle import requires --bundle-root" if bundle_mode && bundle_root_option.empty?
  abort "external fixture portable bundle import requires --executable-relative-path" if bundle_mode && executable_relative_path.empty?

  app_id = DEFAULT_OPERATOR_APP_ID
  app_name = DEFAULT_OPERATOR_APP_NAME
  import_record_path = absolute_path(import_record_option) if record_mode
  bundle_root_path = absolute_path(bundle_root_option) if bundle_mode
  executable_path = if record_mode
                      nil
                    elsif bundle_mode
                      bundle_root_path.join(executable_relative_path).cleanpath
                    else
                      absolute_path(options.fetch(:executable))
                    end
else
  abort "unsupported fixture #{fixture.inspect}; expected notepad, notepad-file-argument, or external"
end
app_id = options.fetch(:app_id).to_s.strip unless options.fetch(:app_id).to_s.strip.empty?
app_name = options.fetch(:display_name).to_s.strip unless options.fetch(:display_name).to_s.strip.empty?
window_match = options.fetch(:window_match).to_s.strip unless options.fetch(:window_match).to_s.strip.empty?
desktop_argument_mode = options.fetch(:desktop_argument_mode).to_s.strip
abort "desktop argument mode must be file-uri or none" unless %w[file-uri none].include?(desktop_argument_mode)
file_open_lane = desktop_argument_mode == "file-uri"
desktop_arguments = file_open_lane ? [sample_document_uri] : []
docker_bin = resolve_executable(options.fetch(:docker))

go_env = {
  "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
  "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
  "GOTMPDIR" => GO_CACHE_ROOT.join("tmp").to_s
}

FileUtils.mkdir_p(build_root)
FileUtils.mkdir_p(stage_root)
FileUtils.mkdir_p(state_root)
FileUtils.mkdir_p(one_shot_stage_root)
FileUtils.mkdir_p(one_shot_state_root)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("tmp"))
File.write(sample_document_path, "Xnix external Windows app file-open smoke document\n") if file_open_lane
FileUtils.mkdir_p(report_output.dirname)
FileUtils.mkdir_p(markdown_output.dirname)
FileUtils.mkdir_p(delegated_output.dirname)
FileUtils.mkdir_p(activation_status_output.dirname)
FileUtils.mkdir_p(launch_packet_output.dirname)
FileUtils.mkdir_p(one_shot_output.dirname)
FileUtils.mkdir_p(one_shot_launch_packet_output.dirname)
FileUtils.mkdir_p(runtime_packet_output.dirname)
FileUtils.mkdir_p(kde_page_output.dirname)

runtime_command = runtime_go_command
runtime_binary_from_runtime_image = false
if runtime_command.nil? && !docker_bin.nil? && docker_image_available?(docker_bin, options.fetch(:runtime_image))
  runtime_go_bin = build_root.join("xnix-runtime-go")
  runtime_binary_from_runtime_image = copy_image_binary(docker_bin, options.fetch(:runtime_image), RUNTIME_GO_IMAGE_PATH, runtime_go_bin)
  runtime_command = [runtime_go_bin.to_s] if runtime_binary_from_runtime_image && runtime_go_bin.file? && runtime_go_bin.executable?
end
unless runtime_command
  write_skip_report(
    report_output,
    markdown_output,
    "xnix-runtime-go is unavailable; local Go compilation is disabled by default, set #{RUNTIME_GO_BIN_ENV} to a prebuilt Runtime binary, provide #{RUNTIME_IMAGE_ENV}=xnix-builder:#{VERSION}, or build on q4 with #{REMOTE_GO_BUILD_HINT}",
    app_id,
    app_name
  )
  exit 0
end

managed_launcher_bin = launcher_binary
launcher_binary_from_runtime_image = false
if managed_launcher_bin.nil? && !docker_bin.nil? && docker_image_available?(docker_bin, options.fetch(:runtime_image))
  runtime_launcher_bin = build_root.join("xnix-compat-launch")
  launcher_binary_from_runtime_image = copy_image_binary(docker_bin, options.fetch(:runtime_image), COMPAT_LAUNCH_IMAGE_PATH, runtime_launcher_bin)
  managed_launcher_bin = runtime_launcher_bin.to_s if launcher_binary_from_runtime_image && runtime_launcher_bin.file? && runtime_launcher_bin.executable?
end
if managed_launcher_bin.nil? && ENV.fetch(ALLOW_LOCAL_GO_COMPILE_ENV, "") == "1" && command_available?("go")
  build_stdout, build_stderr, build_status = run_command(
    go_env,
    "go", "build", "-o", launcher_bin.to_s, "./cmd/xnix-compat-launch"
  )
  abort "go build failed:\n#{build_stderr}\n#{build_stdout}" unless build_status.success?
  managed_launcher_bin = launcher_bin.to_s
end

unless managed_launcher_bin
  write_skip_report(
    report_output,
    markdown_output,
    "xnix-compat-launch is unavailable; local Go compilation is disabled by default, set #{COMPAT_LAUNCH_BIN_ENV} to a prebuilt launcher, provide #{RUNTIME_IMAGE_ENV}=xnix-builder:#{VERSION}, or build on q4 with #{REMOTE_GO_BUILD_HINT}",
    app_id,
    app_name
  )
  exit 0
end

if docker_bin.nil?
  write_skip_report(report_output, markdown_output, "Docker runner #{options.fetch(:docker)} is unavailable", app_id, app_name)
  exit 0
end

unless docker_image_available?(docker_bin, options.fetch(:image))
  write_skip_report(report_output, markdown_output, "Docker image #{options.fetch(:image)} is unavailable", app_id, app_name)
  exit 0
end

if options.fetch(:executable).to_s.strip.empty? && %w[notepad notepad-file-argument].include?(fixture)
  notepad_path = container_notepad_path(docker_bin, options.fetch(:image))
  if notepad_path.nil?
    write_skip_report(report_output, markdown_output, "Wine Notepad was not found in #{options.fetch(:image)}", app_id, app_name)
    exit 0
  end
  copy_container_notepad(docker_bin, options.fetch(:image), notepad_path, executable_path)
end

assert(import_record_path.file?, "external Windows import record must exist") if record_mode
assert(executable_path.file?, "external Windows executable must exist") unless record_mode
if bundle_mode
  assert(bundle_root_path.directory?, "external Windows portable bundle root must exist")
  assert(executable_path.to_s.start_with?("#{bundle_root_path}/"), "external Windows portable executable must stay inside the bundle root")
end

if record_mode
  import_record = JSON.parse(import_record_path.read)
  app_id = import_record.fetch("application_id")
  app_name = import_record.fetch("display_name")
  state_root = state_root_from_import_record_path(import_record_path, import_record)
  import_stdout = JSON.pretty_generate(import_record)
  import_stdout_forbidden = [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, import_record_path.to_s]
else
  import_command = [
    *runtime_command,
    bundle_mode ? "external-winapp-bundle-import-record" : "external-winapp-import-record",
    "--state-root", state_root.to_s
  ]
  if bundle_mode
    import_command.concat([
      "--bundle-root", bundle_root_path.to_s,
      "--executable-relative-path", executable_relative_path
    ])
  else
    import_command.concat(["--executable", executable_path.to_s])
  end
  import_command.concat([
    "--app-id", app_id,
    "--display-name", app_name
  ])
  import_record, import_stdout = run_json(go_env, *import_command)
  import_record_path = state_root.join(import_record.fetch("record_relative_path"))
  import_stdout_forbidden = [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : ""]
end
assert(import_record.fetch("import_recorded") == true, "external app import record must be persisted")
assert(import_record.fetch("artifact_copied") == true, "external app import must copy the artifact")
assert(import_record.fetch("windows_executable_validated") == true, "external app import must validate an MZ executable")
if bundle_mode || import_record.fetch("artifact_kind", "") == "portable-directory"
  assert(import_record.fetch("request_type") == "external-winapp-bundle-import-record", "external portable bundle import must use the bundle import record request type")
  assert(import_record.fetch("artifact_kind") == "portable-directory", "external portable bundle import must record portable-directory artifact kind")
  assert(import_record.fetch("bundle_manifest_sha256").to_s.match?(/\A[0-9a-f]{64}\z/), "external portable bundle import must record the bundle manifest digest")
  executable_relative_path = import_record.fetch("executable_relative_path", executable_relative_path)
  assert(!executable_relative_path.empty?, "external portable bundle import must preserve the executable relative path")
else
  assert(import_record.fetch("request_type") == "external-winapp-import-record", "external single executable import must use the executable import record request type")
end
assert(import_record.fetch("host_root_modified") == false, "external app import must not mutate the host root")
assert_no_forbidden(import_stdout, import_stdout_forbidden, "external import output")
stage, stage_stdout = run_json(
  go_env,
  *runtime_command,
  "desktop-activation-stage",
  "--external-app-import-record", import_record_path.to_s,
  "--mode", "development",
  "--staging-root", stage_root.to_s,
  "--managed-launcher-bin", managed_launcher_bin
)
written_ids = stage.fetch("written_file_ids")
%w[desktop-entry managed-launcher-artifact managed-launcher-executable].each do |id|
  assert(written_ids.include?(id), "desktop activation stage must write #{id}")
end
assert(stage.fetch("application_id") == app_id, "desktop activation stage must target the imported app")
assert(stage.fetch("external_app_handle") == app_id, "desktop activation stage must expose the opaque external app handle")
assert(stage.fetch("desktop_exec_uses_external_app_handle") == true, "desktop activation stage must prove the desktop Exec uses the external app handle")
assert(stage.fetch("external_app_desktop_handle_ready") == true, "desktop activation stage must mark external app desktop handle readiness")
assert(stage.fetch("desktop_exec_uses_raw_import_record") == false, "desktop activation stage must not expose raw import-record Exec routing")
assert(stage.fetch("desktop_exec_uses_state_root") == false, "desktop activation stage must not expose state-root Exec routing")
assert(stage.fetch("launch_enabled") == false, "desktop activation stage must keep launch gated")
assert(stage.fetch("execution_started") == false, "desktop activation stage must not start execution")
assert(stage.fetch("host_root_modified") == false, "desktop activation stage must not mutate the host root")
assert_no_forbidden(stage_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : ""], "desktop activation stage output")

activation_status, status_stdout = run_json(
  go_env,
  *runtime_command,
  "desktop-activation-status-preview",
  "--external-app-import-record", import_record_path.to_s,
  "--mode", "development",
  "--activation-root", stage_root.to_s
)
File.write(activation_status_output, JSON.pretty_generate(activation_status) + "\n")
receipt_evidence = activation_status.fetch("receipt_evidence")
assert(receipt_evidence.fetch("external_app_handle") == app_id, "activation receipt evidence must preserve the opaque external app handle")
assert(receipt_evidence.fetch("desktop_exec_uses_external_app_handle") == true, "activation receipt evidence must prove external app handle Exec routing")
assert(receipt_evidence.fetch("external_app_desktop_handle_ready") == true, "activation receipt evidence must mark external desktop handle readiness")
assert(receipt_evidence.fetch("desktop_exec_uses_raw_import_record") == false, "activation receipt evidence must not persist raw import-record Exec routing")
assert(receipt_evidence.fetch("desktop_exec_uses_state_root") == false, "activation receipt evidence must not persist state-root Exec routing")
assert(receipt_evidence.fetch("safe_for_kde") == true, "activation receipt evidence must stay safe for KDE consumption")
assert(activation_status_output.file?, "activation status file must be written")
assert_no_forbidden(status_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : ""], "desktop activation status output")

desktop_artifact = stage.fetch("written_files").find { |entry| entry.fetch("id") == "desktop-entry" }
desktop_path = stage_root.join(desktop_artifact.fetch("relative_path"))
assert(desktop_path.file?, "staged desktop entry must exist")
assert(staged_launcher.file?, "staged managed launcher must exist")

desktop_exec = desktop_exec_from(desktop_path)
desktop_tokens = Shellwords.split(desktop_exec).reject { |token| token == "%U" }
assert(desktop_tokens == ["xnix-compat-launch", "--external-app-handle", app_id], "desktop Exec must expose only the external app handle")
assert_no_forbidden(desktop_exec, [state_root.to_s, import_record_path.to_s, executable_path.to_s, "notepad.exe", "wine ", "docker", "qemu-system", "--state-root", "--external-app-import-record"], "desktop Exec")

launcher_env = {
  "XNIX_EXTERNAL_APP_STATE_ROOT" => state_root.to_s,
  "XNIX_DOCKER_BIN" => docker_bin,
  "XNIX_WINE_IMAGE" => options.fetch(:image),
  "XNIX_COMPAT_LAUNCH_TIMEOUT" => options.fetch(:timeout),
  "XNIX_EXTERNAL_APP_DESKTOP_ACTIVATION_ROOT" => stage_root.to_s,
  "XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT" => launch_packet_output.to_s,
  "XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_MODE" => "development"
}
launcher_env["XNIX_EXTERNAL_APP_WINDOW_MATCH"] = window_match unless window_match.empty?
launcher_argv = [
  staged_launcher.to_s,
  *desktop_tokens.drop(1),
  *desktop_arguments
]
assert(launcher_argv == [staged_launcher.to_s, "--external-app-handle", app_id, *desktop_arguments], "staged launcher invocation must match the desktop Exec handle route plus the configured KDE desktop arguments without extra launcher options")

launcher_stdout, launcher_stderr, launcher_status = run_command(launcher_env, *launcher_argv)
abort "staged external app launcher failed:\n#{launcher_stderr}\n#{launcher_stdout}" unless launcher_status.success?

payload = JSON.parse(launcher_stdout)
File.write(delegated_output, JSON.pretty_generate(payload) + "\n")
assert(payload.fetch("request_type") == "windows-external-app-run", "launcher must enter the external app Runtime run")
assert(payload.fetch("status") == "passed", "launcher smoke must pass")
assert(payload.fetch("application_id") == app_id, "launcher smoke must preserve imported app id")
assert(payload.fetch("external_app_import_record_consumed") == true, "launcher smoke must consume the import record")
assert(payload.fetch("external_app_handle_consumed") == true, "launcher smoke must consume the desktop handle")
assert(payload.fetch("external_desktop_argument_count") == desktop_arguments.length, "launcher smoke must preserve the configured KDE desktop argument count")
assert(payload.fetch("external_file_uri_arguments_accepted") == file_open_lane, "launcher smoke must report file URI argument acceptance only for the file-open lane")
assert(payload.fetch("external_file_open_requested") == file_open_lane, "launcher smoke must report file-open requests only for the file-open lane")
assert(payload.fetch("external_file_bridge_copy_enabled") == file_open_lane, "launcher smoke must enable copy-only file bridging only for the file-open lane")
assert(payload.fetch("external_file_bridge_copied_count") == desktop_arguments.length, "launcher smoke must copy exactly the configured file-open arguments")
assert(payload.fetch("external_file_bridge_argument_observed_count") == desktop_arguments.length, "launcher smoke must observe exactly the configured file-open arguments")
assert(payload.fetch("external_file_bridge_winepath_translated_count") == desktop_arguments.length, "launcher smoke must translate exactly the configured file-open arguments")
if file_open_lane
  assert(payload.fetch("external_file_bridge_arguments_passed") == true, "launcher smoke must observe file-open arguments at the container launch boundary")
  assert(payload.fetch("external_file_bridge_winepath_translated") == true, "launcher smoke must translate copied file arguments to Wine paths")
  assert(payload.fetch("external_file_bridge_ready") == true, "launcher smoke must mark the copied, translated, and passed file bridge ready")
else
  assert(payload.fetch("external_file_bridge_arguments_passed") == false, "launcher smoke must not claim file argument passing when no file-open argument was supplied")
  assert(payload.fetch("external_file_bridge_winepath_translated") == false, "launcher smoke must not claim Wine path translation when no file-open argument was supplied")
  assert(payload.fetch("external_file_bridge_ready") == false, "launcher smoke must not claim file bridge readiness when no file-open argument was supplied")
end
assert(payload.fetch("external_file_bridge_mount_enabled") == false, "launcher smoke must keep file bridge mounts disabled before policy")
assert(payload.fetch("raw_file_uri_arguments_exposed") == false, "launcher smoke must not expose raw file URI arguments")
assert(payload.fetch("imported_artifact_digest_verified") == true, "launcher smoke must verify the imported artifact digest")
if bundle_mode
  runtime_payload = payload.fetch("runtime_payload")
  assert(payload.fetch("artifact_kind") == "portable-directory", "launcher smoke must preserve portable-directory artifact kind")
  assert(payload.fetch("executable_relative_path") == executable_relative_path, "launcher smoke must preserve the portable executable relative path")
  assert(payload.fetch("application_workspace_copied") == true, "launcher smoke must copy the portable workspace into the container")
  assert(payload.fetch("application_workspace_mode") == "portable-directory", "launcher smoke must record portable-directory workspace mode")
  assert(runtime_payload.fetch("application_name") == "/app/#{executable_relative_path}", "launcher smoke must launch the portable executable from /app")
  assert(runtime_payload.fetch("application_workspace_copied") == true, "runtime payload must copy the portable workspace")
  assert(runtime_payload.fetch("application_workspace_mode") == "portable-directory", "runtime payload must record portable-directory workspace mode")
end
assert(payload.fetch("x_window_observed") == true, "launcher smoke must observe a Windows GUI X window")
assert(payload.fetch("window_observed") == true, "launcher smoke must expose generic observed-window evidence")
if !window_match.empty?
  runtime_payload = payload.fetch("runtime_payload")
  assert(runtime_payload.fetch("window_match") == window_match, "launcher smoke must use the configured window match")
  assert(payload.fetch("window_evidence_summary").include?(window_match), "launcher smoke must observe the configured Windows process window title")
end
assert(payload.fetch("container_network_mode") == "none", "launcher smoke must disable container networking")
assert(payload.fetch("container_host_mount_count") == 0, "launcher smoke must not mount host directories")
assert(payload.fetch("docker_socket_mounted") == false, "launcher smoke must not mount the Docker socket")
assert(payload.fetch("host_networking_required") == false, "launcher smoke must not require host networking")
assert(payload.fetch("broad_host_mount_required") == false, "launcher smoke must not require broad host mounts")
assert(payload.fetch("host_root_modified") == false, "launcher smoke must not mutate the host root")
assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : "", sample_document_path.to_s, sample_document_uri, docker_bin], "staged launcher output")

assert(launch_packet_output.file?, "staged launcher must write the desktop launch packet sidecar")
launch_packet_text = launch_packet_output.read
launch_packet = JSON.parse(launch_packet_text)
assert(launch_packet.fetch("request_type") == "desktop-external-winapp-launch-packet-preview", "desktop launch packet must use the external launch packet request type")
assert(launch_packet.fetch("status") == "passed", "desktop launch packet must pass")
assert(launch_packet.fetch("application_id") == app_id, "desktop launch packet must target the imported app")
assert(launch_packet.fetch("external_app_handle") == app_id, "desktop launch packet must preserve the opaque external app handle")
assert(launch_packet.fetch("activation_receipt_backed") == true, "desktop launch packet must consume activation receipt evidence")
assert(launch_packet.fetch("activation_receipt_safe_for_kde") == true, "desktop launch packet must keep receipt evidence safe for KDE")
assert(launch_packet.fetch("desktop_exec_uses_external_app_handle") == true, "desktop launch packet must prove handle-only desktop Exec routing")
assert(launch_packet.fetch("external_app_desktop_handle_ready") == true, "desktop launch packet must preserve desktop handle readiness")
assert(launch_packet.fetch("desktop_exec_uses_raw_import_record") == false, "desktop launch packet must keep raw import-record Exec routing closed")
assert(launch_packet.fetch("desktop_exec_uses_state_root") == false, "desktop launch packet must keep state-root Exec routing closed")
assert(launch_packet.fetch("run_record_consumed") == true, "desktop launch packet must consume the actual launcher run record")
assert(launch_packet.fetch("external_app_run_record_consumed") == true, "desktop launch packet must mark the external run record consumed")
assert(launch_packet.fetch("external_app_import_record_consumed") == true, "desktop launch packet must preserve import-record consumption")
assert(launch_packet.fetch("external_app_handle_consumed") == true, "desktop launch packet must preserve handle consumption")
assert(launch_packet.fetch("external_desktop_argument_count") == desktop_arguments.length, "desktop launch packet must preserve the configured KDE desktop argument count")
assert(launch_packet.fetch("external_file_uri_arguments_accepted") == file_open_lane, "desktop launch packet must preserve file URI argument acceptance only for the file-open lane")
assert(launch_packet.fetch("external_file_open_requested") == file_open_lane, "desktop launch packet must preserve file-open request evidence only for the file-open lane")
assert(launch_packet.fetch("external_file_bridge_copy_enabled") == file_open_lane, "desktop launch packet must preserve copy-only file bridge evidence only for the file-open lane")
assert(launch_packet.fetch("external_file_bridge_copied_count") == desktop_arguments.length, "desktop launch packet must preserve copied file-open count")
assert(launch_packet.fetch("external_file_bridge_argument_observed_count") == desktop_arguments.length, "desktop launch packet must preserve observed file-open argument count")
assert(launch_packet.fetch("external_file_bridge_winepath_translated_count") == desktop_arguments.length, "desktop launch packet must preserve Wine path translation count")
if file_open_lane
  assert(launch_packet.fetch("external_file_bridge_arguments_passed") == true, "desktop launch packet must preserve observed file-open argument passing")
  assert(launch_packet.fetch("external_file_bridge_winepath_translated") == true, "desktop launch packet must preserve Wine path translation evidence")
  assert(launch_packet.fetch("external_file_bridge_ready") == true, "desktop launch packet must preserve copied, translated, and passed file bridge readiness")
else
  assert(launch_packet.fetch("external_file_bridge_arguments_passed") == false, "desktop launch packet must not claim file-open argument passing when no argument was supplied")
  assert(launch_packet.fetch("external_file_bridge_winepath_translated") == false, "desktop launch packet must not claim Wine path translation when no argument was supplied")
  assert(launch_packet.fetch("external_file_bridge_ready") == false, "desktop launch packet must not claim file bridge readiness when no argument was supplied")
end
assert(launch_packet.fetch("external_file_bridge_mount_enabled") == false, "desktop launch packet must keep file bridge mounts disabled before policy")
assert(launch_packet.fetch("raw_file_uri_arguments_exposed") == false, "desktop launch packet must not expose raw file URI arguments")
assert(launch_packet.fetch("imported_artifact_digest_verified") == true, "desktop launch packet must preserve imported artifact digest verification")
assert(launch_packet.fetch("runtime_launch_executed") == true, "desktop launch packet must prove Runtime launch execution")
assert(launch_packet.fetch("window_observed") == true, "desktop launch packet must preserve generic observed-window evidence")
assert(launch_packet.fetch("x_window_observed") == true, "desktop launch packet must preserve X observed-window evidence")
assert(launch_packet.fetch("container_network_mode") == "none", "desktop launch packet must preserve container network isolation")
assert(launch_packet.fetch("container_host_mount_count") == 0, "desktop launch packet must preserve zero host mounts")
assert(launch_packet.fetch("runtime_launch_authority") == true, "desktop launch packet must keep launch authority in Runtime")
assert(launch_packet.fetch("kde_launch_authority") == false, "desktop launch packet must not give launch authority to KDE")
assert(launch_packet.fetch("desktop_launch_packet_ready") == true, "desktop launch packet must be ready")
assert(launch_packet.fetch("safe_for_kde") == true, "desktop launch packet must be safe for KDE")
assert(launch_packet.fetch("unsafe_reason_ids").empty?, "desktop launch packet must not report unsafe reasons")
assert(launch_packet.fetch("backend_details_exposed") == false, "desktop launch packet must not expose backend details")
assert(launch_packet.fetch("raw_import_record_path_exposed") == false, "desktop launch packet must not expose raw import-record paths")
assert(launch_packet.fetch("raw_state_root_path_exposed") == false, "desktop launch packet must not expose raw state-root paths")
assert(launch_packet.fetch("raw_executable_path_exposed") == false, "desktop launch packet must not expose raw executable paths")
assert(launch_packet.fetch("host_root_modified") == false, "desktop launch packet must not mutate the host root")
assert(launch_packet.fetch("docker_socket_mounted") == false, "desktop launch packet must not mount the Docker socket")
assert(launch_packet.fetch("broad_host_mount_required") == false, "desktop launch packet must not require broad host mounts")
assert_no_forbidden(launch_packet_text, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, delegated_output.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : "", sample_document_path.to_s, sample_document_uri, docker_bin, "notepad.exe", "wine ", "docker run", "/var/run/docker.sock"], "desktop launch packet output")

one_shot_command = [
  *runtime_command,
  "external-winapp-import-stage-and-launch"
]
if record_mode
  one_shot_command.concat(["--external-app-import-record", import_record_path.to_s])
elsif bundle_mode
  one_shot_command.concat(["--state-root", one_shot_state_root.to_s])
  one_shot_command.concat([
    "--bundle-root", bundle_root_path.to_s,
    "--executable-relative-path", executable_relative_path
  ])
else
  one_shot_command.concat(["--state-root", one_shot_state_root.to_s])
  one_shot_command.concat(["--executable", executable_path.to_s])
end
one_shot_command.concat([
  "--app-id", app_id,
  "--display-name", app_name,
  "--mode", "development",
  "--staging-root", one_shot_stage_root.to_s,
  "--managed-launcher-bin", managed_launcher_bin,
  "--desktop-launch-packet-output", one_shot_launch_packet_output.to_s,
  "--image", options.fetch(:image),
  "--docker", docker_bin,
  "--timeout", options.fetch(:timeout)
])
one_shot_command.concat(["--window-match", window_match]) unless window_match.empty?
one_shot_command.concat(desktop_arguments)
one_shot, one_shot_stdout = run_json(go_env, *one_shot_command)
File.write(one_shot_output, JSON.pretty_generate(one_shot) + "\n")
assert(one_shot.fetch("request_type") == "external-winapp-import-stage-and-launch", "one-shot Runtime command must use the import-stage-and-launch request type")
assert(one_shot.fetch("status") == "passed", "one-shot Runtime command must pass")
assert(one_shot.fetch("application_id") == app_id, "one-shot Runtime command must target the imported app")
assert(one_shot.fetch("external_app_handle") == app_id, "one-shot Runtime command must preserve the opaque desktop handle")
assert(one_shot.fetch("import_recorded") == true, "one-shot Runtime command must persist the import record")
assert(one_shot.fetch("existing_import_record_consumed") == record_mode, "one-shot Runtime command must report whether it consumed an existing import record")
assert(one_shot.fetch("desktop_activation_staged") == true, "one-shot Runtime command must stage desktop activation artifacts")
assert(one_shot.fetch("staged_launcher_invoked") == true, "one-shot Runtime command must invoke the staged launcher")
assert(one_shot.fetch("staged_launcher_from_activation_root") == true, "one-shot Runtime command must run the launcher staged under the activation root")
assert(one_shot.fetch("managed_launcher_executable_staged") == true, "one-shot Runtime command must stage the managed launcher executable")
assert(one_shot.fetch("desktop_exec_uses_external_app_handle") == true, "one-shot Runtime command must preserve handle-only desktop Exec routing")
assert(one_shot.fetch("external_app_desktop_handle_ready") == true, "one-shot Runtime command must preserve desktop handle readiness")
assert(one_shot.fetch("desktop_launch_packet_requested") == true, "one-shot Runtime command must request a desktop launch packet")
assert(one_shot.fetch("desktop_launch_packet_written") == true, "one-shot Runtime command must write the desktop launch packet")
assert(one_shot.fetch("launcher_request_type") == "windows-external-app-run", "one-shot Runtime command must invoke the external app launcher request")
assert(one_shot.fetch("launcher_status") == "passed", "one-shot Runtime command must receive a passed launcher result")
assert(one_shot.fetch("external_app_import_record_consumed") == true, "one-shot Runtime command must consume the import record through the launcher")
assert(one_shot.fetch("external_app_handle_consumed") == true, "one-shot Runtime command must consume the desktop handle through the launcher")
assert(one_shot.fetch("external_desktop_argument_count") == desktop_arguments.length, "one-shot Runtime command must preserve the configured KDE desktop argument count")
assert(one_shot.fetch("external_file_uri_arguments_accepted") == file_open_lane, "one-shot Runtime command must accept KDE file URI arguments only for the file-open lane")
assert(one_shot.fetch("external_file_bridge_ready") == file_open_lane, "one-shot Runtime command must prove file bridging only for the file-open lane")
assert(one_shot.fetch("imported_artifact_digest_verified") == true, "one-shot Runtime command must verify the imported artifact digest")
if bundle_mode
  assert(one_shot.fetch("launcher_result").fetch("artifact_kind") == "portable-directory", "one-shot Runtime command must preserve portable-directory artifact kind")
  assert(one_shot.fetch("launcher_result").fetch("application_workspace_copied") == true, "one-shot Runtime command must copy the portable workspace")
  assert(one_shot.fetch("launcher_result").fetch("application_workspace_mode") == "portable-directory", "one-shot Runtime command must preserve portable-directory workspace mode")
end
assert(one_shot.fetch("runtime_launch_executed") == true, "one-shot Runtime command must execute the Runtime launch path")
assert(one_shot.fetch("window_observed") == true, "one-shot Runtime command must observe a Windows GUI window")
assert(one_shot.fetch("x_window_observed") == true, "one-shot Runtime command must preserve X window evidence")
assert(one_shot.fetch("launch_enabled") == false, "one-shot Runtime command must keep staged desktop launch gated")
assert(one_shot.fetch("backend_launch_enabled") == false, "one-shot Runtime command must keep backend launch gated at staging")
assert(one_shot.fetch("execution_started") == true, "one-shot Runtime command must prove execution started through the managed launcher")
assert(one_shot.fetch("host_root_modified") == false, "one-shot Runtime command must not mutate the host root")
assert(one_shot.fetch("docker_socket_mounted") == false, "one-shot Runtime command must not mount the Docker socket")
assert(one_shot.fetch("broad_host_mount_required") == false, "one-shot Runtime command must not require broad host mounts")
assert(one_shot.fetch("raw_import_record_path_exposed") == false, "one-shot Runtime command must not expose raw import-record paths")
assert(one_shot.fetch("raw_state_root_path_exposed") == false, "one-shot Runtime command must not expose raw state-root paths")
assert(one_shot.fetch("raw_executable_path_exposed") == false, "one-shot Runtime command must not expose raw executable paths")
assert(one_shot.fetch("raw_launcher_path_exposed") == false, "one-shot Runtime command must not expose raw launcher paths")
assert(one_shot.fetch("raw_launcher_output_exposed") == false, "one-shot Runtime command must not expose raw launcher output")
assert(one_shot_output.file?, "one-shot Runtime command output file must be written")
assert(one_shot_launch_packet_output.file?, "one-shot Runtime command desktop launch packet must be written")
assert_no_forbidden(one_shot_stdout, [PROJECT_ROOT.to_s, one_shot_stage_root.to_s, one_shot_state_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : "", sample_document_path.to_s, sample_document_uri, docker_bin, "docker run", "/var/run/docker.sock"], "one-shot Runtime command output")

runtime_packet, = run_json(
  go_env,
  *runtime_command,
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
  *runtime_command,
  "kde-center-page-preview",
  "--external-app-evidence-file", runtime_packet_output.to_s,
  "--decision", "approved"
)
File.write(kde_page_output, JSON.pretty_generate(kde_page) + "\n")
cards = kde_page.fetch("known_app_gui_evidence_cards")
assert(kde_page.fetch("application_id") == app_id, "KDE page must target the imported app")
assert(kde_page.fetch("known_app_gui_evidence_count") == 1, "KDE page must consume one real GUI evidence item")
assert(cards.length == 1, "KDE page must render one real GUI evidence card")
assert(cards.first.fetch("app_id") == app_id, "KDE GUI evidence card must target the imported app")
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
assert_no_forbidden(kde_stdout, [PROJECT_ROOT.to_s, run_root.to_s, stage_root.to_s, state_root.to_s, import_record_path.to_s, executable_path.to_s, bundle_mode ? bundle_root_path.to_s : ""], "KDE page output")

packet = {
  "schema_version" => SCHEMA_VERSION,
  "status" => "passed",
  "fixture" => fixture,
  "app_id" => app_id,
  "display_name" => app_name,
  "version" => VERSION,
  "desktop_argument_mode" => desktop_argument_mode,
  "desktop_file_open_lane" => file_open_lane,
  "portable_directory_external_app" => import_record.fetch("artifact_kind", "") == "portable-directory",
  "portable_directory_bundle_import_recorded" => import_record.fetch("request_type") == "external-winapp-bundle-import-record",
  "existing_import_record_consumed" => record_mode,
  "record_first_launch_path" => record_mode,
  "artifact_kind" => import_record.fetch("artifact_kind", ""),
  "bundle_manifest_sha256_present" => import_record.fetch("bundle_manifest_sha256", "").to_s.match?(/\A[0-9a-f]{64}\z/),
  "executable_relative_path" => import_record.fetch("executable_relative_path", ""),
  "application_workspace_copied" => payload.fetch("application_workspace_copied", false),
  "application_workspace_mode" => payload.fetch("application_workspace_mode", ""),
  "desktop_exec_uses_external_app_handle" => stage.fetch("desktop_exec_uses_external_app_handle"),
  "external_app_desktop_handle_ready" => stage.fetch("external_app_desktop_handle_ready"),
  "activation_receipt_external_app_desktop_handle_ready" => receipt_evidence.fetch("external_app_desktop_handle_ready"),
  "activation_receipt_safe_for_kde" => receipt_evidence.fetch("safe_for_kde"),
  "desktop_exec_invocation_exact" => launcher_argv == [staged_launcher.to_s, "--external-app-handle", app_id, *desktop_arguments],
  "launcher_context_from_environment" => true,
  "launcher_extra_arguments_appended" => false,
  "external_desktop_argument_count" => payload.fetch("external_desktop_argument_count"),
  "external_file_uri_arguments_accepted" => payload.fetch("external_file_uri_arguments_accepted"),
  "external_file_open_requested" => payload.fetch("external_file_open_requested"),
  "external_file_bridge_copy_enabled" => payload.fetch("external_file_bridge_copy_enabled"),
  "external_file_bridge_copied_count" => payload.fetch("external_file_bridge_copied_count"),
  "external_file_bridge_arguments_passed" => payload.fetch("external_file_bridge_arguments_passed"),
  "external_file_bridge_argument_observed_count" => payload.fetch("external_file_bridge_argument_observed_count"),
  "external_file_bridge_winepath_translated" => payload.fetch("external_file_bridge_winepath_translated"),
  "external_file_bridge_winepath_translated_count" => payload.fetch("external_file_bridge_winepath_translated_count"),
  "external_file_bridge_ready" => payload.fetch("external_file_bridge_ready"),
  "windows_process_file_argument_window_observed" => file_open_lane && !window_match.empty? && payload.fetch("window_evidence_summary").include?(window_match),
  "external_file_bridge_mount_enabled" => payload.fetch("external_file_bridge_mount_enabled"),
  "raw_file_uri_arguments_exposed" => payload.fetch("raw_file_uri_arguments_exposed"),
  "desktop_launch_packet_output_written" => launch_packet_output.file?,
  "desktop_launch_packet_ready" => launch_packet.fetch("desktop_launch_packet_ready"),
  "desktop_launch_packet_safe_for_kde" => launch_packet.fetch("safe_for_kde"),
  "desktop_launch_packet_external_app_handle_consumed" => launch_packet.fetch("external_app_handle_consumed"),
  "desktop_launch_packet_window_observed" => launch_packet.fetch("window_observed"),
  "desktop_launch_packet_x_window_observed" => launch_packet.fetch("x_window_observed"),
  "one_shot_output_written" => one_shot_output.file?,
  "one_shot_path" => one_shot_output.to_s,
  "one_shot_launch_packet_path" => one_shot_launch_packet_output.to_s,
  "one_shot_status" => one_shot.fetch("status"),
  "one_shot_import_recorded" => one_shot.fetch("import_recorded"),
  "one_shot_desktop_activation_staged" => one_shot.fetch("desktop_activation_staged"),
  "one_shot_staged_launcher_invoked" => one_shot.fetch("staged_launcher_invoked"),
  "one_shot_staged_launcher_from_activation_root" => one_shot.fetch("staged_launcher_from_activation_root"),
  "one_shot_managed_launcher_executable_staged" => one_shot.fetch("managed_launcher_executable_staged"),
  "one_shot_desktop_exec_uses_external_app_handle" => one_shot.fetch("desktop_exec_uses_external_app_handle"),
  "one_shot_external_app_desktop_handle_ready" => one_shot.fetch("external_app_desktop_handle_ready"),
  "one_shot_desktop_launch_packet_written" => one_shot.fetch("desktop_launch_packet_written"),
  "one_shot_external_app_handle_consumed" => one_shot.fetch("external_app_handle_consumed"),
  "one_shot_external_file_bridge_ready" => one_shot.fetch("external_file_bridge_ready"),
  "one_shot_imported_artifact_digest_verified" => one_shot.fetch("imported_artifact_digest_verified"),
  "one_shot_runtime_launch_executed" => one_shot.fetch("runtime_launch_executed"),
  "one_shot_window_observed" => one_shot.fetch("window_observed"),
  "one_shot_x_window_observed" => one_shot.fetch("x_window_observed"),
  "one_shot_host_root_modified" => one_shot.fetch("host_root_modified"),
  "one_shot_docker_socket_mounted" => one_shot.fetch("docker_socket_mounted"),
  "one_shot_raw_paths_exposed" => one_shot.fetch("raw_import_record_path_exposed") || one_shot.fetch("raw_state_root_path_exposed") || one_shot.fetch("raw_executable_path_exposed") || one_shot.fetch("raw_launcher_path_exposed"),
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
  "desktop_launch_packet_path" => launch_packet_output.to_s,
  "kde_page_path" => kde_page_output.to_s,
  "runtime_image" => options.fetch(:runtime_image),
  "runtime_binary_from_runtime_image" => runtime_binary_from_runtime_image,
  "launcher_binary_from_runtime_image" => launcher_binary_from_runtime_image,
  "report_path" => report_output.to_s,
  "markdown_path" => markdown_output.to_s
}

File.write(report_output, JSON.pretty_generate(packet) + "\n")
write_markdown_report(markdown_output, packet)

puts "PASS: staged desktop external Windows app smoke"
puts JSON.pretty_generate(packet)
