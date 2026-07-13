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

assert(core_header.read.include?("#define XNIX_RUNTIME_VERSION \"0.2.59\""), "C Runtime core must expose the current version")
assert(core_header.read.include?("XNIX_RUNTIME_WRITE_ERROR"), "C Runtime core must define the write-method D-Bus error")
assert(core_header.read.include?("XnixRuntimeApplication"), "C Runtime core must define application catalog records")
assert(core_header.read.include?("XnixRuntimeEngine"), "C Runtime core must define compatibility engine records")
assert(core_header.read.include?("XnixRuntimePortalPolicy"), "C Runtime core must define Portal access policy records")
assert(core_source.read.include?("xnix_runtime_write_gate"), "C Runtime core must implement write gate logic")
assert(core_source.read.include?("blocked-until-production-backend"), "C Runtime core must keep write gates blocked")
assert(core_source.read.include?("xnix_runtime_find_application"), "C Runtime core must query application catalog records")
assert(core_source.read.include?("xnix_runtime_select_engine_for_mode"), "C Runtime core must select engine records by recipe mode")
assert(core_source.read.include?("xnix_runtime_find_portal_policy"), "C Runtime core must query Portal access policies")
assert(core_cli.read.include?("business_logic_runtime"), "C Runtime core CLI must identify the business runtime language")
assert(core_cli.read.include?("list-applications"), "C Runtime core CLI must expose application catalog listing")
assert(core_cli.read.include?("list-engines"), "C Runtime core CLI must expose engine catalog listing")
assert(core_cli.read.include?("list-portal-policies"), "C Runtime core CLI must expose Portal policy listing")

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
  assert(probe["version"] == "0.2.59", "C Runtime core probe must expose the current version")
  assert(probe["core_language"] == "c", "C Runtime core must report C as its implementation language")
  assert(probe["business_logic_runtime"] == "c", "important Runtime business logic must be C-owned")
  assert(probe["ruby_role"] == "tests-and-development-tools", "Ruby must be limited to tests and development tooling")
  assert(probe["bus_name"] == "org.xnix.Compatibility1", "C Runtime core must keep the stable bus name")
  assert(probe["runtime_owned"], "C Runtime core must expose Runtime ownership")
  assert(!probe["kde_policy_owner"], "KDE must not own Runtime policy")
  assert(probe["application_catalog_owner"] == "c", "C Runtime core must own the first application catalog layer")
  assert(probe["application_count"] == 1, "C Runtime core must count registered applications")
  assert(probe["engine_catalog_owner"] == "c", "C Runtime core must own the first engine catalog layer")
  assert(probe["engine_count"] == 3, "C Runtime core must count registered engines")
  assert(probe["portal_policy_owner"] == "c", "C Runtime core must own the first Portal policy layer")
  assert(probe["portal_policy_count"] == 7, "C Runtime core must count registered Portal policies")
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

  stdout, stderr, status = Open3.capture3(binary.to_s, "list-engines")
  assert(status.success?, "C Runtime core engine catalog must exit successfully: #{stderr}")
  engine_catalog = JSON.parse(stdout)
  assert(engine_catalog["version"] == "0.2.59", "C Runtime core engine catalog must expose the current version")
  assert(engine_catalog["catalog_type"] == "compatibility-engine", "C Runtime core engine catalog must identify the catalog type")
  assert(engine_catalog["runtime_policy_owner"], "C Runtime core must own engine policy")
  assert(!engine_catalog["desktop_shell_policy_owner"], "C Runtime core must not make the desktop shell the engine policy owner")
  assert(engine_catalog["catalog_owner"] == "c", "C Runtime core engine catalog must be C-owned")
  assert(!engine_catalog["backend_details_exposed"], "C Runtime core engine catalog must hide backend details")
  assert(engine_catalog["engine_count"] == 3, "C Runtime core engine catalog must count engines")
  engines = engine_catalog.fetch("engines")
  assert(engines.map { |engine| engine.fetch("id") } == %w[automatic-managed local-compatibility-engine isolated-compatibility-engine], "C Runtime core must keep engine order stable")
  assert(engines.all? { |engine| !engine.fetch("ready") }, "C Runtime core must not claim engine readiness")
  assert(engines.all? { |engine| !engine.fetch("launch_enabled") }, "C Runtime core must not claim launch enablement")
  assert(engines.all? { |engine| engine.fetch("user_visible") }, "C Runtime core must expose engines as user-visible Runtime strategies")
  assert(engines.all? { |engine| !engine.fetch("backend_details_exposed") }, "C Runtime core engines must hide backend details")
  assert(!stdout.match?(/prefix|\.wine|proton|virtual machine/i), "C Runtime core engine catalog must not expose backend implementation terms")

  stdout, stderr, status = Open3.capture3(binary.to_s, "select-engine", "automatic")
  assert(status.success?, "C Runtime core automatic engine selection must exit successfully: #{stderr}")
  automatic = JSON.parse(stdout)
  assert(automatic["engine_id"] == "automatic-managed", "automatic mode must select the automatic-managed engine")
  assert(!automatic["launch_enabled"], "automatic mode must not enable launch yet")

  stdout, stderr, status = Open3.capture3(binary.to_s, "select-engine", "wine")
  assert(status.success?, "C Runtime core local engine selection must exit successfully: #{stderr}")
  local = JSON.parse(stdout)
  assert(local["engine_id"] == "local-compatibility-engine", "local recipe mode must select the local compatibility engine")
  assert(!local["backend_details_exposed"], "local engine selection must hide backend details")

  stdout, stderr, status = Open3.capture3(binary.to_s, "select-engine", "vm")
  assert(status.success?, "C Runtime core isolated engine selection must exit successfully: #{stderr}")
  isolated = JSON.parse(stdout)
  assert(isolated["engine_id"] == "isolated-compatibility-engine", "isolated recipe mode must select the isolated compatibility engine")

  _stdout, stderr, status = Open3.capture3(binary.to_s, "select-engine", "unknown")
  assert(!status.success?, "C Runtime core must reject unknown recipe modes")
  assert(stderr.include?("mode must be one of"), "C Runtime core must explain recipe mode validation")

  stdout, stderr, status = Open3.capture3(binary.to_s, "list-portal-policies")
  assert(status.success?, "C Runtime core Portal policy catalog must exit successfully: #{stderr}")
  portal_catalog = JSON.parse(stdout)
  assert(portal_catalog["version"] == "0.2.59", "C Runtime core Portal policy catalog must expose the current version")
  assert(portal_catalog["catalog_type"] == "portal-access-policy", "C Runtime core Portal policy catalog must identify the catalog type")
  assert(portal_catalog["runtime_policy_owner"], "C Runtime core must own Portal policy")
  assert(!portal_catalog["desktop_shell_policy_owner"], "C Runtime core must not make the desktop shell the Portal policy owner")
  assert(portal_catalog["catalog_owner"] == "c", "C Runtime core Portal policy catalog must be C-owned")
  assert(!portal_catalog["backend_details_exposed"], "C Runtime core Portal policy catalog must hide backend details")
  assert(portal_catalog["policy_count"] == 7, "C Runtime core Portal policy catalog must count policies")
  policies = portal_catalog.fetch("policies")
  assert(policies.map { |policy| policy.fetch("operation") } == %w[file-open uri-open print screenshot clipboard camera remote-desktop], "C Runtime core must keep Portal policy order stable")
  assert(policies.all? { |policy| policy.fetch("portal_required") }, "C Runtime core policies must require portals")
  assert(policies.all? { |policy| policy.fetch("user_mediation_required") }, "C Runtime core policies must require user mediation")
  assert(policies.all? { |policy| !policy.fetch("direct_access_allowed") }, "C Runtime core policies must deny direct desktop access")
  assert(policies.all? { |policy| policy.fetch("runtime_policy_owner") }, "C Runtime core policies must be Runtime-owned")
  assert(policies.all? { |policy| !policy.fetch("desktop_shell_policy_owner") }, "C Runtime core policies must not be KDE-owned")
  assert(policies.all? { |policy| !policy.fetch("backend_details_exposed") }, "C Runtime core policies must hide backend details")
  assert(!stdout.match?(/wine|prefix|\.wine|proton|virtual machine/i), "C Runtime core Portal policy catalog must not expose backend implementation terms")

  stdout, stderr, status = Open3.capture3(binary.to_s, "portal-policy", "org.xnix.sample.notepad", "file-open")
  assert(status.success?, "C Runtime core file Portal policy must exit successfully: #{stderr}")
  file_policy = JSON.parse(stdout)
  assert(file_policy["policy_type"] == "portal-access", "C Runtime core Portal policy must identify the policy type")
  assert(file_policy["desktop"] == "KDE Plasma", "C Runtime core Portal policy must target KDE Plasma")
  assert(file_policy["application_id"] == "org.xnix.sample.notepad", "C Runtime core Portal policy must preserve the application id")
  assert(file_policy["operation"] == "file-open", "C Runtime core Portal policy must preserve the operation")
  assert(file_policy["decision"] == "ask", "C Runtime core must ask for file access")
  assert(file_policy["portal_interface"] == "org.freedesktop.portal.FileChooser", "C Runtime core must select the file chooser Portal")
  assert(file_policy["resources"].include?("selected-files"), "C Runtime core must scope selected files")
  assert(file_policy["request_flow"]["dbus_api"] == "XDG Desktop Portal", "C Runtime core Portal policy must use XDG Desktop Portal")
  assert(file_policy["request_flow"]["runtime_policy_owner"], "C Runtime core Portal request flow must be Runtime-owned")
  assert(!file_policy["request_flow"]["desktop_shell_policy_owner"], "C Runtime core Portal request flow must not be KDE-owned")
  assert(!file_policy["direct_access_allowed"], "C Runtime core Portal policy must deny direct access")
  assert(!file_policy["backend_details_exposed"], "C Runtime core Portal policy must hide backend details")

  stdout, stderr, status = Open3.capture3(binary.to_s, "portal-policy", "org.xnix.sample.notepad", "camera")
  assert(status.success?, "C Runtime core camera Portal policy must exit successfully: #{stderr}")
  camera_policy = JSON.parse(stdout)
  assert(camera_policy["decision"] == "deny", "C Runtime core must deny camera access by default")

  _stdout, stderr, status = Open3.capture3(binary.to_s, "portal-policy", "org.xnix.unknown", "file-open")
  assert(!status.success?, "C Runtime core must reject Portal policy for unknown applications")
  assert(stderr.include?("application is not registered"), "C Runtime core must explain unknown Portal application validation")

  _stdout, stderr, status = Open3.capture3(binary.to_s, "portal-policy", "org.xnix.sample.notepad", "unknown")
  assert(!status.success?, "C Runtime core must reject unknown Portal operations")
  assert(stderr.include?("registered sensitive desktop operation"), "C Runtime core must explain unknown Portal operation validation")

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
