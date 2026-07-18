#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
DEFAULT_GO_BIN = Pathname.new("/opt/homebrew/bin/go").executable? ? "/opt/homebrew/bin/go" : "go"
EXPECTED_SHAPES = %w[
  document-editor
  game
  installer
  launcher
  network-heavy
  tray-heavy
  unsupported
].freeze
SAFETY_FALSE_KEYS = %w[
  network_fetch_enabled
  package_manager_invoked
  artifact_staging_enabled
  backend_launch_enabled
  docker_required
  qemu_required
  request_objects_created
  request_object_created
  settings_persisted
  file_content_read
  state_root_path_exposed
  raw_executable_exposed
  raw_command_exposed
  backend_details_exposed
  host_root_modified
  privileged_container_required
  backend_process_started
  launch_enabled
  execution_started
].freeze

class FixtureMatrixFailure < StandardError; end

def parse_options(argv)
  options = {
    go_bin: DEFAULT_GO_BIN,
    runtime_root: ".",
    artifact_receipt_root: nil,
    snapshot_state_root: nil,
    format: "json",
    shapes: []
  }

  parser = OptionParser.new do |opts|
    opts.banner = "Usage: ruby scripts/offline_application_fixture_matrix.rb [options]"
    opts.on("--go-bin PATH", "Go binary to use") { |value| options[:go_bin] = value }
    opts.on("--runtime-root PATH", "Project root used for read-only Runtime evidence checks") { |value| options[:runtime_root] = value }
    opts.on("--artifact-receipt-root PATH", "Controlled root for read-only artifact stage receipt evidence") { |value| options[:artifact_receipt_root] = value }
    opts.on("--snapshot-state-root PATH", "Controlled state root for read-only snapshot baseline evidence") { |value| options[:snapshot_state_root] = value }
    opts.on("--shape ID", "Fixture shape id to include; may be repeated") { |value| options[:shapes] << value }
    opts.on("--format FORMAT", "Output format: json or markdown") { |value| options[:format] = value }
  end
  parser.parse!(argv)
  raise FixtureMatrixFailure, "unexpected positional arguments: #{argv.join(" ")}" unless argv.empty?

  options
end

def collect_matrix(options)
  env = {
    "GOCACHE" => PROJECT_ROOT.join(".cache", "go-build").to_s
  }
  command = [
    options.fetch(:go_bin),
    "run",
    "./cmd/xnix-runtime-go",
    "offline-application-fixture-matrix-preview",
    "--runtime-root",
    options.fetch(:runtime_root)
  ]
  options.fetch(:shapes).each do |shape|
    command.concat(["--shape", shape])
  end
  if options[:artifact_receipt_root]
    command.concat(["--artifact-receipt-root", options.fetch(:artifact_receipt_root)])
  end
  if options[:snapshot_state_root]
    command.concat(["--snapshot-state-root", options.fetch(:snapshot_state_root)])
  end
  stdout, stderr, status = Open3.capture3(env, *command, chdir: PROJECT_ROOT.to_s)
  raise FixtureMatrixFailure, "offline fixture matrix failed: #{stderr.strip}" unless status.success?

  JSON.parse(stdout)
rescue JSON::ParserError => e
  raise FixtureMatrixFailure, "offline fixture matrix did not return JSON: #{e.message}"
end

def assert_matrix!(payload)
  assert(payload.fetch("schema_version") == "xnix.runtime.offline_application_fixture_matrix.v1", "schema version mismatch")
  assert(payload.fetch("request_type") == "offline-application-fixture-matrix-preview", "request type mismatch")
  assert(payload.fetch("review_only") == true, "matrix must be review-only")
  assert(payload.fetch("offline_default") == true, "matrix must default to offline")
  assert(payload.fetch("runtime_owned") == true, "matrix must be Runtime-owned")
  assert(payload.fetch("go_runtime_backed") == true, "matrix must be Go Runtime backed")
  assert(payload.fetch("kde_policy_owner") == false, "KDE must not own fixture policy")

  shape_ids = payload.fetch("shape_ids")
  missing = EXPECTED_SHAPES - shape_ids
  assert(missing.empty?, "missing expected shapes: #{missing.join(", ")}") if payload.fetch("missing_shape_ids").empty?
  assert(payload.fetch("rows").all? { |row| row.fetch("kde_journey_entry_point_count") == 7 }, "each row must cover the seven KDE entry points")
  assert(payload.fetch("rows").any? { |row| row.fetch("shape_id") == "unsupported" && row.fetch("matrix_state") == "blocked-unsupported" }, "unsupported shape must be explicitly blocked")
  assert_safety_false!(payload)
end

def assert_safety_false!(payload)
  deep_each(payload) do |path, value|
    key = path.last.to_s
    next unless SAFETY_FALSE_KEYS.include?(key)

    raise FixtureMatrixFailure, "#{path.join(".")} expected false, got true" if value == true
  end
end

def deep_each(value, path = [], &block)
  yield(path, value)
  case value
  when Hash
    value.each { |key, child| deep_each(child, [*path, key], &block) }
  when Array
    value.each_with_index { |child, index| deep_each(child, [*path, index], &block) }
  end
end

def assert(condition, message)
  raise FixtureMatrixFailure, message unless condition
end

def render_json(payload)
  JSON.pretty_generate(payload)
end

def render_markdown(payload)
  lines = []
  lines << "# Offline Application Fixture Matrix"
  lines << ""
  lines << "- Schema: `#{payload.fetch("schema_version")}`"
  lines << "- Status: `#{payload.fetch("matrix_status")}`"
  lines << "- Shapes: #{payload.fetch("counts").fetch("total")}"
  lines << "- Missing fixtures: #{payload.fetch("counts").fetch("missing_fixture")}"
  lines << "- Unsafe actions enabled: `false`"
  lines << ""
  lines << "| Shape | Application | Profile | Artifact | Snapshot | Portal needs | KDE entry points | State |"
  lines << "| --- | --- | --- | --- | --- | --- | ---: | --- |"
  payload.fetch("rows").each do |row|
    lines << [
      row.fetch("shape_id"),
      row.fetch("application_name"),
      row.fetch("backend_profile_mapping"),
      row.fetch("artifact_readiness"),
      row.fetch("snapshot_readiness"),
      row.fetch("portal_needs").join(", "),
      row.fetch("kde_journey_entry_point_count"),
      row.fetch("matrix_state")
    ].join(" | ").then { |body| "| #{body} |" }
  end
  lines << ""
  lines << "Blocked unsafe actions:"
  payload.fetch("blocked_unsafe_actions").each do |action|
    lines << "- #{action}"
  end
  lines.join("\n")
end

def main
  options = parse_options(ARGV)
  payload = collect_matrix(options)
  assert_matrix!(payload)
  case options.fetch(:format)
  when "json"
    puts render_json(payload)
  when "markdown"
    puts render_markdown(payload)
  else
    raise FixtureMatrixFailure, "unknown format: #{options.fetch(:format)}"
  end
end

begin
  main
rescue FixtureMatrixFailure => e
  warn "FAIL: #{e.message}"
  exit 1
end
