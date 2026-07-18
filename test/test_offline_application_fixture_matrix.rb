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
script_path = project_root.join("scripts/offline_application_fixture_matrix.rb")
script = script_path.read

%w[
  offline-application-fixture-matrix-preview
  --format
  --shape
  --artifact-receipt-root
  --snapshot-state-root
  JSON.pretty_generate
  render_markdown
  GOCACHE
  .cache
  network_fetch_enabled
  package_manager_invoked
  artifact_staging_enabled
  backend_launch_enabled
  docker_required
  qemu_required
  host_root_modified
  file_content_read
  state_root_path_exposed
  raw_command_exposed
  backend_details_exposed
].each do |token|
  assert(script.include?(token), "offline fixture matrix script must include #{token}")
end

%w[
  document-editor
  game
  installer
  launcher
  network-heavy
  tray-heavy
  unsupported
].each do |shape|
  assert(script.include?(shape), "offline fixture matrix script must expect #{shape}")
end

assert(!script.include?("scripts/container.rb"), "offline fixture matrix must not run Docker container smokes")
assert(!script.include?("boot-system"), "offline fixture matrix must not boot machine smoke checks")
assert(!script.include?("docs/claude-code-implementation-packages.md"), "offline fixture matrix must not touch protected Claude Code package work")

env = {
  "GOCACHE" => project_root.join(".cache", "go-build").to_s
}
json_stdout, json_stderr, json_status = Open3.capture3(env, "ruby", script_path.to_s, "--format", "json", chdir: project_root.to_s)
assert(json_status.success?, "offline fixture matrix JSON run failed: #{json_stderr}")
payload = JSON.parse(json_stdout)
assert(payload.fetch("schema_version") == "xnix.runtime.offline_application_fixture_matrix.v1", "unexpected schema")
assert(payload.fetch("request_type") == "offline-application-fixture-matrix-preview", "unexpected request type")
assert(payload.fetch("counts").fetch("total") == 7, "matrix must include seven shapes")
assert(payload.fetch("counts").fetch("unsupported") == 1, "matrix must include one unsupported shape")
assert(payload.fetch("missing_shape_ids").empty?, "default matrix must not miss fixtures")
assert(payload.fetch("rows").all? { |row| row.fetch("kde_journey_entry_point_count") == 7 }, "each fixture row must cover seven KDE entry points")
assert(payload.fetch("rows").any? { |row| row.fetch("shape_id") == "unsupported" && row.fetch("matrix_state") == "blocked-unsupported" }, "unsupported row must be blocked")

payload_text = json_stdout.downcase
%w[
  network_fetch_enabled
  package_manager_invoked
  artifact_staging_enabled
  backend_launch_enabled
  docker_required
  qemu_required
  host_root_modified
].each do |key|
  assert(payload_text.include?(%("#{key}": false)), "#{key} must remain false")
end

markdown_stdout, markdown_stderr, markdown_status = Open3.capture3(env, "ruby", script_path.to_s, "--format", "markdown", chdir: project_root.to_s)
assert(markdown_status.success?, "offline fixture matrix Markdown run failed: #{markdown_stderr}")
assert(markdown_stdout.include?("# Offline Application Fixture Matrix"), "Markdown output must include title")
assert(markdown_stdout.include?("| Shape | Application | Profile | Artifact | Snapshot | Portal needs | KDE entry points | State |"), "Markdown output must include matrix table")
assert(markdown_stdout.include?("unsupported"), "Markdown output must include unsupported shape")
assert(markdown_stdout.include?("blocked-unsupported"), "Markdown output must include blocked unsupported state")

filtered_stdout, filtered_stderr, filtered_status = Open3.capture3(env, "ruby", script_path.to_s, "--format", "json", "--shape", "document-editor", "--shape", "unsupported", chdir: project_root.to_s)
assert(filtered_status.success?, "offline fixture matrix filtered run failed: #{filtered_stderr}")
filtered = JSON.parse(filtered_stdout)
assert(filtered.fetch("counts").fetch("total") == 2, "filtered matrix must include two rows")
assert(filtered.fetch("counts").fetch("missing_fixture") == 5, "filtered matrix must report five missing fixtures")

puts "PASS: offline application fixture matrix script unit tests"
