#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
core_source = project_root.join("runtime/core/xnix_runtime_core.c")
core_header = project_root.join("runtime/core/xnix_runtime_core.h")
core_cli = project_root.join("runtime/core/xnix_runtime_core_cli.c")

assert(core_header.read.include?("#define XNIX_RUNTIME_VERSION \"0.2.57\""), "C Runtime core must expose the current version")
assert(core_header.read.include?("XNIX_RUNTIME_WRITE_ERROR"), "C Runtime core must define the write-method D-Bus error")
assert(core_header.read.include?("XnixRuntimeApplication"), "C Runtime core must define application catalog records")
assert(core_source.read.include?("xnix_runtime_write_gate"), "C Runtime core must implement write gate logic")
assert(core_source.read.include?("blocked-until-production-backend"), "C Runtime core must keep write gates blocked")
assert(core_source.read.include?("xnix_runtime_find_application"), "C Runtime core must query application catalog records")
assert(core_cli.read.include?("business_logic_runtime"), "C Runtime core CLI must identify the business runtime language")
assert(core_cli.read.include?("list-applications"), "C Runtime core CLI must expose application catalog listing")

Dir.mktmpdir("xnix-runtime-core") do |dir|
  binary = Pathname.new(dir).join("xnix-runtime-core")
  _stdout, stderr, status = Open3.capture3(
    "cc",
    "-std=c11",
    "-Wall",
    "-Wextra",
    "-Werror",
    core_source.to_s,
    core_cli.to_s,
    "-o",
    binary.to_s
  )
  assert(status.success?, "C Runtime core must compile cleanly: #{stderr}")

  stdout, stderr, status = Open3.capture3(binary.to_s, "probe")
  assert(status.success?, "C Runtime core probe must exit successfully: #{stderr}")
  probe = JSON.parse(stdout)
  assert(probe["version"] == "0.2.57", "C Runtime core probe must expose the current version")
  assert(probe["core_language"] == "c", "C Runtime core must report C as its implementation language")
  assert(probe["business_logic_runtime"] == "c", "important Runtime business logic must be C-owned")
  assert(probe["ruby_role"] == "tests-and-development-tools", "Ruby must be limited to tests and development tooling")
  assert(probe["bus_name"] == "org.xnix.Compatibility1", "C Runtime core must keep the stable bus name")
  assert(probe["runtime_owned"], "C Runtime core must expose Runtime ownership")
  assert(!probe["kde_policy_owner"], "KDE must not own Runtime policy")
  assert(probe["application_catalog_owner"] == "c", "C Runtime core must own the first application catalog layer")
  assert(probe["application_count"] == 1, "C Runtime core must count registered applications")
  assert(probe["write_methods"] == %w[InstallRecipe Launch CreateSnapshot RestoreSnapshot], "C Runtime core must list reserved write methods")
  assert(probe["write_method_count"] == 4, "C Runtime core must count write methods")
  assert(!probe["write_methods_supported"], "C Runtime core must not claim write method support")
  assert(!probe["write_method_dispatch_enabled"], "C Runtime core must not enable write dispatch")

  stdout, stderr, status = Open3.capture3(binary.to_s, "list-applications")
  assert(status.success?, "C Runtime core application list must exit successfully: #{stderr}")
  applications = JSON.parse(stdout)
  assert(applications.length == 1, "C Runtime core must list the sample application")
  application = applications.first
  assert(application["id"] == "org.xnix.sample.notepad", "C Runtime core must preserve the recipe id")
  assert(application["name"] == "Sample Notepad", "C Runtime core must preserve the recipe name")
  assert(application["icon"] == "accessories-text-editor", "C Runtime core must preserve the desktop icon")
  assert(application["runtime_mode"] == "automatic", "C Runtime core must preserve the runtime mode")
  assert(application["desktop_category"] == "Utility", "C Runtime core must expose a safe desktop category")
  assert(application["launcher_command"] == "xnix-compat-launch --app org.xnix.sample.notepad %U", "C Runtime core must expose the managed launcher")
  assert(application["primary_extension"] == ".txt", "C Runtime core must expose the primary file extension")
  assert(application["primary_mime_type"] == "application/x-xnix-txt", "C Runtime core must expose the primary MIME type")
  assert(application["runtime_owned"], "C Runtime core application catalog must be Runtime-owned")
  assert(!application["kde_policy_owner"], "C Runtime core application catalog must not make KDE the policy owner")
  assert(!application["backend_details_exposed"], "C Runtime core application catalog must hide backend details")
  assert(!stdout.match?(/wine|prefix|proton|virtual machine/i), "C Runtime core application catalog must not expose backend details")

  stdout, stderr, status = Open3.capture3(binary.to_s, "get-application", "org.xnix.sample.notepad")
  assert(status.success?, "C Runtime core application lookup must exit successfully: #{stderr}")
  found_application = JSON.parse(stdout)
  assert(found_application == application, "C Runtime core lookup must return the catalog record")

  _stdout, stderr, status = Open3.capture3(binary.to_s, "get-application", "org.xnix.unknown")
  assert(!status.success?, "C Runtime core must reject unknown application ids")
  assert(stderr.include?("not registered"), "C Runtime core must explain unknown application validation")

  stdout, stderr, status = Open3.capture3(binary.to_s, "write-gate", "Launch")
  assert(status.success?, "C Runtime core write-gate must exit successfully: #{stderr}")
  gate = JSON.parse(stdout)
  assert(gate["gate_type"] == "runtime-write-gate", "C Runtime core must emit write gate models")
  assert(gate["method_name"] == "Launch", "C Runtime core write gate must preserve the method")
  assert(gate["gate_decision"] == "blocked-until-production-backend", "C Runtime core must block writes until production backend")
  assert(!gate["write_method_enabled"], "C Runtime core must not enable writes")
  assert(!gate["dispatch_enabled"], "C Runtime core must not enable write dispatch")
  assert(!gate["request_object_created"], "C Runtime core must not create request objects")
  assert(!gate["execution_started"], "C Runtime core must not start execution")
  assert(gate["denial_error_name"] == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "C Runtime core must expose the write denial error")

  _stdout, stderr, status = Open3.capture3(binary.to_s, "write-gate", "Unknown")
  assert(!status.success?, "C Runtime core must reject unknown write gate methods")
  assert(stderr.include?("reserved Runtime write method"), "C Runtime core must explain write gate validation")
end

puts "PASS: C Runtime core unit tests"
