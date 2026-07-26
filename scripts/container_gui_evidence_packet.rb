#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
GO_CACHE_ROOT = PROJECT_ROOT.join(".gocache")
DEFAULT_PACKET_ROOT = Pathname.new("/tmp/xnix-container-gui-evidence-packet-#{VERSION}")
SCHEMA_VERSION = "xnix.scripts.container_gui_evidence_packet.v1"

options = {
  report_output: DEFAULT_PACKET_ROOT.join("winapp-container-x-gui-report.json").to_s,
  evidence_output: DEFAULT_PACKET_ROOT.join("winapp-container-x-gui-evidence.json").to_s,
  runtime_packet_output: DEFAULT_PACKET_ROOT.join("real-winapp-gui-evidence-packet.json").to_s,
  kde_page_output: DEFAULT_PACKET_ROOT.join("winapp-container-x-gui-kde-page.json").to_s,
  registry: PROJECT_ROOT.join("runtime/recipes/registry.json").to_s,
  recipe_app: "org.xnix.sample.notepad",
  app_id: "org.xnix.sample.notepad",
  display_name: "Sample Notepad",
  app_version: VERSION,
  decision: "approved",
  image: ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local"),
  platform: ENV.fetch("XNIX_WINE_PLATFORM", "linux/amd64"),
  gui_app: ENV.fetch("XNIX_WINE_GUI_CONTAINER_APP", "notepad.exe"),
  window_match: ENV.fetch("XNIX_WINE_GUI_CONTAINER_WINDOW_MATCH", "notepad.exe"),
  timeout: "120s",
  docker: nil
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/container_gui_evidence_packet.rb [options]"
  parser.on("--report-output PATH", "JSON smoke report output path") { |value| options[:report_output] = value }
  parser.on("--evidence-output PATH", "Runtime GUI evidence output path") { |value| options[:evidence_output] = value }
  parser.on("--runtime-packet-output PATH", "Go Runtime real Windows app GUI evidence packet output path") { |value| options[:runtime_packet_output] = value }
  parser.on("--kde-page-output PATH", "KDE center page JSON output path") { |value| options[:kde_page_output] = value }
  parser.on("--registry PATH", "Recipe registry path") { |value| options[:registry] = value }
  parser.on("--recipe-app ID", "Registered app id with container GUI smoke hints") { |value| options[:recipe_app] = value }
  parser.on("--app-id ID", "Application id for evidence and KDE page") { |value| options[:app_id] = value }
  parser.on("--display-name NAME", "Display name for evidence") { |value| options[:display_name] = value }
  parser.on("--app-version VERSION", "Application version for evidence") { |value| options[:app_version] = value }
  parser.on("--decision DECISION", "KDE page decision") { |value| options[:decision] = value }
  parser.on("--image IMAGE", "Local Wine GUI smoke image") { |value| options[:image] = value }
  parser.on("--platform PLATFORM", "Container platform") { |value| options[:platform] = value }
  parser.on("--gui-app APP", "Windows GUI app inside the container") { |value| options[:gui_app] = value }
  parser.on("--window-match TEXT", "X window match text") { |value| options[:window_match] = value }
  parser.on("--timeout DURATION", "GUI smoke timeout") { |value| options[:timeout] = value }
  parser.on("--docker PATH", "Explicit Docker runner path") { |value| options[:docker] = value }
end.parse!

abort "container GUI evidence packet does not accept positional arguments" unless ARGV.empty?
abort "container GUI evidence packet requires --app-id" if options.fetch(:app_id).to_s.strip.empty?
abort "container GUI evidence packet requires --recipe-app" if options.fetch(:recipe_app).to_s.strip.empty?
abort "container GUI evidence packet requires --display-name" if options.fetch(:display_name).to_s.strip.empty?
abort "container GUI evidence packet requires --app-version" if options.fetch(:app_version).to_s.strip.empty?
abort "container GUI evidence packet decision must be reviewed, approved, deferred, or rejected" unless %w[reviewed approved deferred rejected].include?(options.fetch(:decision))

def resolve_output_path(path)
  output = Pathname.new(path)
  output = PROJECT_ROOT.join(output) unless output.absolute?
  output.cleanpath
end

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def load_json(stdout, label)
  JSON.parse(stdout)
rescue JSON::ParserError
  warn "FAIL: #{label} returned malformed JSON"
  exit 1
end

def fail_command(label, stdout, stderr)
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  warn "FAIL: #{label}"
  exit 1
end

report_output = resolve_output_path(options.fetch(:report_output))
evidence_output = resolve_output_path(options.fetch(:evidence_output))
runtime_packet_output = resolve_output_path(options.fetch(:runtime_packet_output))
kde_page_output = resolve_output_path(options.fetch(:kde_page_output))

go_env = {
  "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
  "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
}
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(report_output.dirname)
FileUtils.mkdir_p(evidence_output.dirname)
FileUtils.mkdir_p(runtime_packet_output.dirname)
FileUtils.mkdir_p(kde_page_output.dirname)

smoke_command = [
  "ruby", "scripts/winapp_smoke.rb",
  "--backend", "container-x-gui",
  "--format", "json",
  "--report-output", report_output.to_s,
  "--image", options.fetch(:image),
  "--platform", options.fetch(:platform),
  "--gui-app", options.fetch(:gui_app),
  "--window-match", options.fetch(:window_match),
  "--registry", options.fetch(:registry),
  "--recipe-app", options.fetch(:recipe_app),
  "--timeout", options.fetch(:timeout)
]
smoke_command.concat(["--docker", options.fetch(:docker)]) unless options[:docker].to_s.strip.empty?

smoke_stdout, smoke_stderr, smoke_status = run_command(go_env, *smoke_command)
fail_command("container X GUI smoke failed", smoke_stdout, smoke_stderr) unless smoke_status.zero?
smoke_report = load_json(smoke_stdout, "container X GUI smoke")

unless smoke_report.fetch("status") == "passed" &&
       smoke_report.fetch("report_output_written") &&
       smoke_report.fetch("x_window_observed") &&
       smoke_report.fetch("container_recipe_backed") &&
       smoke_report.fetch("container_application_id") == options.fetch(:recipe_app) &&
       smoke_report.fetch("container_payload").fetch("network_mode") == "none" &&
       smoke_report.fetch("container_payload").fetch("host_mount_count") == 0
  warn "FAIL: container X GUI smoke report did not pass the evidence packet gate"
  exit 1
end

evidence_command = [
  "go", "run", "./cmd/xnix-runtime-go", "gui-smoke-evidence-preview",
  "--gui-smoke-report", report_output.to_s,
  "--app-id", options.fetch(:app_id),
  "--display-name", options.fetch(:display_name),
  "--app-version", options.fetch(:app_version),
  "--output", evidence_output.to_s
]
evidence_stdout, evidence_stderr, evidence_status = run_command(go_env, *evidence_command)
fail_command("GUI smoke evidence projection failed", evidence_stdout, evidence_stderr) unless evidence_status.zero?
evidence = load_json(evidence_stdout, "GUI smoke evidence projection")

unless evidence.fetch("report_consumed") &&
       evidence.fetch("compatibility_center_projection_ready") &&
       evidence.fetch("kde_center_projection_ready") &&
       !evidence.fetch("report_path_exposed") &&
       !evidence.fetch("backend_details_exposed") &&
       !evidence.fetch("host_root_modified")
  warn "FAIL: GUI smoke evidence projection did not pass the packet gate"
  exit 1
end

runtime_packet_command = [
  "go", "run", "./cmd/xnix-runtime-go", "real-winapp-gui-evidence-packet-preview",
  "--gui-smoke-report", report_output.to_s,
  "--app-id", options.fetch(:app_id),
  "--display-name", options.fetch(:display_name),
  "--app-version", options.fetch(:app_version),
  "--output", runtime_packet_output.to_s
]
runtime_packet_stdout, runtime_packet_stderr, runtime_packet_status = run_command(go_env, *runtime_packet_command)
fail_command("real Windows app GUI evidence packet projection failed", runtime_packet_stdout, runtime_packet_stderr) unless runtime_packet_status.zero?
runtime_packet = load_json(runtime_packet_stdout, "real Windows app GUI evidence packet")

unless runtime_packet.fetch("schema_version") == "xnix.runtime.real_winapp_gui_evidence_packet.v1" &&
       runtime_packet.fetch("request_type") == "real-winapp-gui-evidence-packet-preview" &&
       runtime_packet.fetch("report_consumed") &&
       runtime_packet.fetch("compatibility_center_projection_ready") &&
       runtime_packet.fetch("kde_center_projection_ready") &&
       runtime_packet.fetch("evidence_source") == "winapp-smoke-container-x-gui" &&
       runtime_packet.fetch("recipe_backed") &&
       runtime_packet.fetch("recipe_app_id") == options.fetch(:recipe_app) &&
       runtime_packet.fetch("container_network_mode") == "none" &&
       runtime_packet.fetch("container_host_mount_count") == 0 &&
       !runtime_packet.fetch("report_path_exposed") &&
       !runtime_packet.fetch("backend_details_exposed") &&
       !runtime_packet.fetch("host_root_modified")
  warn "FAIL: real Windows app GUI Runtime packet did not pass the packet gate"
  exit 1
end

kde_command = [
  "go", "run", "./cmd/xnix-runtime-go", "kde-center-page-preview",
  "--registry", options.fetch(:registry),
  "--app", options.fetch(:app_id),
  "--decision", options.fetch(:decision),
  "--known-app-evidence-file", evidence_output.to_s
]
kde_stdout, kde_stderr, kde_status = run_command(go_env, *kde_command)
fail_command("KDE center page projection failed", kde_stdout, kde_stderr) unless kde_status.zero?
kde_page = load_json(kde_stdout, "KDE center page projection")
File.write(kde_page_output, JSON.pretty_generate(kde_page) + "\n")

cards = kde_page.fetch("known_app_gui_evidence_cards")
unless kde_page.fetch("known_app_gui_evidence_count") == 1 &&
       cards.length == 1 &&
       cards.first.fetch("app_id") == options.fetch(:app_id) &&
       cards.first.fetch("evidence_source") == "winapp-smoke-container-x-gui" &&
       cards.first.fetch("recipe_backed") &&
       cards.first.fetch("recipe_app_id") == options.fetch(:recipe_app) &&
       cards.first.fetch("compatibility_state") == "real-gui-container-wine-verified" &&
       !cards.first.fetch("desktop_launch_enabled") &&
       !cards.first.fetch("backend_launch_enabled") &&
       !cards.first.fetch("backend_details_exposed") &&
       !cards.first.fetch("host_root_modified")
  warn "FAIL: KDE center page did not pass the container GUI evidence packet gate"
  exit 1
end

packet = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => "container-gui-evidence-packet",
  "version" => VERSION,
  "status" => "passed",
  "app_id" => options.fetch(:app_id),
  "recipe_app" => options.fetch(:recipe_app),
  "display_name" => options.fetch(:display_name),
  "app_version" => options.fetch(:app_version),
  "gui_app" => options.fetch(:gui_app),
  "window_match" => options.fetch(:window_match),
  "container_platform" => smoke_report.fetch("container_platform"),
  "container_recipe_backed" => smoke_report.fetch("container_recipe_backed"),
  "container_application_id" => smoke_report.fetch("container_application_id"),
  "kde_card_recipe_backed" => cards.first.fetch("recipe_backed"),
  "kde_card_recipe_app_id" => cards.first.fetch("recipe_app_id"),
  "report_output_written" => report_output.file?,
  "evidence_output_written" => evidence_output.file?,
  "runtime_packet_output_written" => runtime_packet_output.file?,
  "kde_page_output_written" => kde_page_output.file?,
  "report_path_exposed" => false,
  "evidence_path_exposed" => false,
  "runtime_packet_path_exposed" => false,
  "kde_page_path_exposed" => false,
  "smoke_status" => smoke_report.fetch("status"),
  "x_window_observed" => smoke_report.fetch("x_window_observed"),
  "evidence_source" => cards.first.fetch("evidence_source"),
  "runtime_packet_schema_version" => runtime_packet.fetch("schema_version"),
  "runtime_packet_request_type" => runtime_packet.fetch("request_type"),
  "runtime_packet_consumed_report" => runtime_packet.fetch("report_consumed"),
  "runtime_packet_compatibility_state" => runtime_packet.fetch("compatibility_state"),
  "runtime_packet_known_app_gui_evidence_count" => runtime_packet.fetch("known_app_gui_evidence_count"),
  "runtime_packet_known_app_gui_evidence_verified_count" => runtime_packet.fetch("known_app_gui_evidence_verified_count"),
  "compatibility_state" => cards.first.fetch("compatibility_state"),
  "known_app_gui_evidence_count" => kde_page.fetch("known_app_gui_evidence_count"),
  "desktop_launch_enabled" => cards.first.fetch("desktop_launch_enabled"),
  "backend_launch_enabled" => cards.first.fetch("backend_launch_enabled"),
  "backend_details_exposed" => cards.first.fetch("backend_details_exposed"),
  "host_root_modified" => cards.first.fetch("host_root_modified"),
  "privileged_container_required" => smoke_report.fetch("privileged_container_required"),
  "host_networking_required" => smoke_report.fetch("host_networking_required"),
  "docker_socket_mounted" => smoke_report.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => smoke_report.fetch("broad_host_mount_required"),
  "container_network_mode" => smoke_report.fetch("container_payload").fetch("network_mode"),
  "container_host_mount_count" => smoke_report.fetch("container_payload").fetch("host_mount_count")
}

puts JSON.pretty_generate(packet)
