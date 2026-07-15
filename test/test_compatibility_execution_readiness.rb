# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

ROOT = Pathname.new(__dir__).join("..").expand_path
COMMAND = ROOT.join("bin/xnix-compat-execution-readiness").to_s

stdout, stderr, status = Open3.capture3(
  COMMAND,
  "--app",
  "org.xnix.sample.notepad",
  "--recipe-dir",
  ROOT.join("runtime/recipes").to_s
)

abort stderr unless status.success?

readiness = JSON.parse(stdout)
gate_ids = readiness.fetch("gates").map { |gate| gate.fetch("id") }

def assert(condition, message)
  raise message unless condition
end

assert(readiness["version"] == "0.2.240", "execution readiness must expose the current version")
assert(readiness["readiness_type"] == "compatibility-execution-readiness", "execution readiness must identify the record type")
assert(readiness["runtime_method"] == "GetExecutionReadiness", "execution readiness must expose the Runtime method")
assert(readiness.fetch("application").fetch("id") == "org.xnix.sample.notepad", "execution readiness must preserve the application id")
assert(readiness.fetch("application").fetch("launcher_command").include?("xnix-compat-launch"), "execution readiness must preserve the desktop-safe launcher command")
assert(readiness.fetch("engine").fetch("engine_id") == "automatic-managed", "execution readiness must preserve the selected engine")
assert(readiness["runtime_owned"], "execution readiness must be Runtime-owned")
assert(readiness["c_runtime_backed"], "execution readiness must be C Runtime-backed")
assert(!readiness["kde_policy_owner"], "execution readiness must not be KDE-owned")
assert(readiness["compatibility_center_card"], "execution readiness must be suitable for Compatibility Center cards")
assert(readiness["safe_for_ai_diagnostics"], "execution readiness must be safe for AI diagnostics")
assert(readiness["desktop_entry_launch_visible"], "execution readiness must allow KDE to show the desktop entry")
assert(readiness["execution_state"] == "blocked", "execution readiness must stay blocked before backend binding")
assert(readiness["overall_status"] == "not-ready", "execution readiness must not claim execution readiness")
assert(!readiness["launch_allowed"], "execution readiness must not allow launch")
assert(!readiness["launch_enabled"], "execution readiness must not enable launch")
assert(!readiness["execution_request_created"], "execution readiness must not create execution requests")
assert(!readiness["backend_binding_ready"], "execution readiness must not claim backend binding readiness")
assert(readiness["portal_policy_required"], "execution readiness must require Portal review")
assert(readiness["snapshot_required"], "execution readiness must require a snapshot baseline")
assert(readiness["user_action_required"], "execution readiness must require user or Runtime action while blocked")
assert(gate_ids == %w[recipe-validation portal-policy-review snapshot-baseline backend-binding runtime-launch-write-gate], "execution readiness gates must be stable and ordered")
assert(readiness.fetch("gates").any? { |gate| gate.fetch("status") == "blocked" }, "execution readiness must expose the blocked Launch gate")
assert(readiness.fetch("blocked_actions").include?("launch compatibility backend"), "execution readiness must block backend launch")
assert(readiness.fetch("blocked_actions").include?("expose raw backend command to desktop shell"), "execution readiness must hide raw backend commands")
assert(!readiness["host_root_modified"], "execution readiness must not mutate the host root")
assert(!readiness["network_required"], "execution readiness must not require network access")
assert(!readiness["backend_details_exposed"], "execution readiness must hide backend details")
assert(readiness["desktop_safe_summary"].include?("not ready"), "execution readiness must provide a desktop-safe blocked summary")

_stdout, stderr, status = Open3.capture3(
  COMMAND,
  "--app",
  "org.xnix.unknown",
  "--recipe-dir",
  ROOT.join("runtime/recipes").to_s
)
assert(!status.success?, "execution readiness must reject unknown applications")
assert(stderr.include?("unknown application"), "execution readiness must explain unknown application failures")

puts "PASS: compatibility execution readiness unit tests"
