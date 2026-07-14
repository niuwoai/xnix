# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "fileutils"
require_relative "../lib/xnix/compatibility/compatibility_backend_capability_matrix"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def assert(condition, message)
  raise "FAIL: #{message}" unless condition
end

def assert_no_backend_terms(value, message)
  serialized = JSON.generate(value)
  assert(!serialized.match?(/wine|prefix|\.wine|proton|virtual machine/i), message)
end

matrix = Xnix::Compatibility::CompatibilityBackendCapabilityMatrix.new.to_h

assert(matrix["version"] == "0.2.137", "backend capability matrix must expose the current version")
assert(matrix["matrix_type"] == "compatibility-backend-capability-matrix", "backend capability matrix must identify the matrix type")
assert(matrix["runtime_method"] == "GetBackendCapabilityMatrix", "backend capability matrix must expose the Runtime method")
assert(matrix["profile_count"] == 2, "backend capability matrix must expose two compatibility profiles")
assert(matrix["capability_count"] == 7, "backend capability matrix must expose seven capabilities")
assert(matrix["ready_capability_count"] == 2, "backend capability matrix must count ready diagnostics capabilities across profiles")
assert(matrix["pending_capability_count"] == 12, "backend capability matrix must count pending gated capabilities across profiles")
assert(matrix["profiles"].map { |profile| profile.fetch("id") } == %w[local-compatibility isolated-compatibility], "backend capability matrix must preserve profile order")
assert(matrix["profiles"].all? { |profile| profile.fetch("capabilities").map { |capability| capability.fetch("id") }.include?("application-launch") }, "backend capability matrix must include launch capabilities")
assert(matrix["profiles"].all? { |profile| profile.fetch("capabilities").map { |capability| capability.fetch("id") }.include?("file-bridge") }, "backend capability matrix must include file bridge capabilities")
assert(matrix["profiles"].all? { |profile| profile.fetch("capabilities").map { |capability| capability.fetch("id") }.include?("clipboard-bridge") }, "backend capability matrix must include clipboard bridge capabilities")
assert(matrix["profiles"].all? { |profile| profile.fetch("capabilities").map { |capability| capability.fetch("id") }.include?("print-bridge") }, "backend capability matrix must include print bridge capabilities")
assert(matrix["runtime_owned"], "backend capability matrix must be Runtime-owned")
assert(matrix["c_runtime_backed"], "backend capability matrix must be C Runtime-backed")
assert(!matrix["kde_policy_owner"], "backend capability matrix must not make KDE the policy owner")
assert(!matrix["selection_enabled"], "backend capability matrix must not enable backend selection")
assert(!matrix["backend_launch_enabled"], "backend capability matrix must not enable backend launch")
assert(!matrix["capability_activation_enabled"], "backend capability matrix must not enable capability activation")
assert(!matrix["request_objects_created"], "backend capability matrix must not create request objects")
assert(!matrix["state_root_created"], "backend capability matrix must not create state roots")
assert(!matrix["snapshots_created"], "backend capability matrix must not create snapshots")
assert(!matrix["host_root_modified"], "backend capability matrix must not mutate the host root")
assert(!matrix["privileged_container_required"], "backend capability matrix must not require privileged containers")
assert(!matrix["backend_details_exposed"], "backend capability matrix must not expose backend details")
assert_no_backend_terms(matrix, "backend capability matrix must avoid backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", PROJECT_ROOT.join("bin/xnix-compat-backend-capability-matrix").to_s)
assert(status.success?, "backend capability matrix CLI must exit successfully: #{stderr}")
cli_matrix = JSON.parse(stdout)
assert(cli_matrix["matrix_type"] == "compatibility-backend-capability-matrix", "backend capability matrix CLI must emit the matrix")
assert(cli_matrix["profile_count"] == 2, "backend capability matrix CLI must preserve profile count")
assert(!cli_matrix["backend_launch_enabled"], "backend capability matrix CLI must keep backend launch gated")
assert_no_backend_terms(cli_matrix, "backend capability matrix CLI must avoid backend implementation terms")

core_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core.c")
core_cli_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core_cli.c")
core_binary = PROJECT_ROOT.join("tmp/xnix-runtime-core-backend-capability-matrix-test")
FileUtils.mkdir_p(core_binary.dirname)

_stdout, stderr, status = Open3.capture3(
  "cc",
  "-std=c99",
  "-Wall",
  "-Wextra",
  "-I",
  PROJECT_ROOT.join("runtime/core").to_s,
  core_source.to_s,
  core_cli_source.to_s,
  "-o",
  core_binary.to_s
)
assert(status.success?, "C Runtime core CLI must compile for backend capability matrix: #{stderr}")

stdout, stderr, status = Open3.capture3(core_binary.to_s, "backend-capability-matrix")
assert(status.success?, "C Runtime core CLI must emit backend capability matrix: #{stderr}")
c_matrix = JSON.parse(stdout)
assert(c_matrix["runtime_method"] == "GetBackendCapabilityMatrix", "C Runtime core CLI must expose the Runtime method")
assert(c_matrix["profile_count"] == 2, "C Runtime core CLI must expose two profiles")
assert(c_matrix["capability_count"] == 7, "C Runtime core CLI must expose seven capabilities")
assert(c_matrix["ready_capability_count"] == 2, "C Runtime core CLI must count ready capabilities")
assert(c_matrix["pending_capability_count"] == 12, "C Runtime core CLI must count pending capabilities")
assert(!c_matrix["selection_enabled"], "C Runtime core CLI must not enable backend selection")
assert(!c_matrix["backend_launch_enabled"], "C Runtime core CLI must not enable backend launch")
assert(!c_matrix["capability_activation_enabled"], "C Runtime core CLI must not enable capability activation")
assert(!c_matrix["request_objects_created"], "C Runtime core CLI must not create request objects")
assert(!c_matrix["state_root_created"], "C Runtime core CLI must not create state roots")
assert(!c_matrix["snapshots_created"], "C Runtime core CLI must not create snapshots")
assert(!c_matrix["host_root_modified"], "C Runtime core CLI must not mutate the host root")
assert(!c_matrix["privileged_container_required"], "C Runtime core CLI must not require privileged containers")
assert(!c_matrix["backend_details_exposed"], "C Runtime core CLI must not expose backend details")
assert_no_backend_terms(c_matrix, "C Runtime core CLI must avoid backend implementation terms")

FileUtils.rm_f(core_binary)

puts "PASS: compatibility backend capability matrix unit tests"
