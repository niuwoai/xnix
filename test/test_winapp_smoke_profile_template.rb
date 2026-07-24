#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
template_path = project_root.join("docs/examples/windows-app-smoke-profile.template.json")
runbook_path = project_root.join("docs/windows-app-smoke-profile-runbook.md")

template = JSON.parse(template_path.read)
runbook = runbook_path.read

assert(template.fetch("schema_version") == "xnix.runtime.windows_app_smoke_profile.v1", "profile template must use the smoke profile schema")
assert(template.fetch("executable_path") == "path/to/application.exe", "profile template must use a placeholder executable path")
assert(template.fetch("working_directory") == "path/to/application-directory", "profile template must include a working directory placeholder")
assert(template.fetch("runner_path") == "path/to/wine", "profile template must include a runner placeholder")
assert(template.fetch("state_root") == ".local/xnix/winapp-smoke/profile-state", "profile template must keep state under the project-local smoke root")
assert(template.fetch("timeout") == "30s", "profile template must provide a bounded timeout")
assert(template.fetch("expected_marker") == "XNIX_WINAPP_SMOKE_OK", "profile template must default to the Xnix smoke marker")
assert(template.fetch("success_mode") == "marker", "profile template must default to marker mode")
assert(template.fetch("redact_output"), "profile template must request redacted output")
assert(template.fetch("runner_arguments").is_a?(Array), "profile template runner arguments must be an array")
assert(template.fetch("arguments").is_a?(Array), "profile template app arguments must be an array")

[
  "docs/examples/windows-app-smoke-profile.template.json",
  "windows-app-smoke-profile-preflight",
  "ruby scripts/winapp_smoke.rb --format json --profile",
  "go run ./cmd/xnix-runtime-go windows-app-run-smoke --profile",
  "profile_supplied",
  "working_directory_mode",
  "runner_argument_count",
  "startup-window"
].each do |token|
  assert(runbook.include?(token), "profile runbook must include #{token}")
end

forbidden_terms = [
  "--privileged",
  "--network host",
  "docker.sock"
]
forbidden_terms.each do |term|
  assert(!template_path.read.include?(term), "profile template must not include unsafe term #{term}")
  assert(!runbook.include?(term), "profile runbook must not include unsafe term #{term}")
end
