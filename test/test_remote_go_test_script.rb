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
script = project_root.join("scripts/remote_go_test.rb")
source = script.read

assert(source.include?("root@q4"), "remote Go test must default to q4")
assert(source.include?("DEFAULT_PACKAGES"), "remote Go test must declare default targeted packages")
assert(source.include?("DEFAULT_RUN_REGEX"), "remote Go test must declare a default targeted regex")
assert(source.include?("./internal/runtime/owner"), "remote Go test must include Runtime owner tests by default")
assert(source.include?("./cmd/xnix-runtime-owner"), "remote Go test must include Runtime owner CLI tests by default")
assert(source.include?("./cmd/xnix-compat-open"), "remote Go test must include compat-open tests by default")
assert(source.include?("./cmd/xnix-compat-launch"), "remote Go test must include compat-launch tests by default")
assert(source.include?("./internal/runtime/appidentity"), "remote Go test must include appidentity tests by default")
assert(source.include?("./cmd/xnix-runtime-go"), "remote Go test must include Runtime Go command tests by default")
assert(source.include?("docs/claude-code-implementation-packages.md"), "remote Go test source sync must exclude the protected Claude package document")
assert(source.include?("ensure_remote_xnix_path!"), "remote Go test must constrain remote writable paths")
assert(source.include?("ensure_go_package!"), "remote Go test must validate package targets")
assert(source.include?("ensure_go_test_regex!"), "remote Go test must validate test regex")
assert(source.include?("host_compilation_avoided"), "remote Go test must make host compilation avoidance observable")
assert(source.include?("targeted_test_required"), "remote Go test must make targeted small-version testing observable")
assert(source.include?("ssh_command"), "remote Go test must use bounded SSH options")
assert(source.include?("BatchMode=yes"), "remote Go test must use non-interactive SSH mode")
assert(source.include?("ServerAliveInterval=15"), "remote Go test must use SSH keepalive")
assert(source.include?("remote shell operation timed out"), "remote Go test must bound remote shell operations")
assert(source.include?("GOCACHE="), "remote Go test must keep Go cache on q4")
assert(source.include?("GOMODCACHE="), "remote Go test must keep module cache on q4")
assert(source.include?("GOTMPDIR="), "remote Go test must keep temporary files on q4")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "remote Go test plan must succeed: #{stderr}")

payload = JSON.parse(stdout)
assert(payload["schema_version"] == "xnix.scripts.remote_go_test.v1", "remote Go test must expose a stable schema")
assert(payload["request_type"] == "remote-go-test", "remote Go test must expose request type")
assert(payload["status"] == "planned", "remote Go test must be planned by default")
assert(payload["execute"] == false, "remote Go test must not execute without --execute")
assert(payload["remote_host"] == "root@q4", "remote Go test must default to q4")
assert(payload["source_sync_planned"] == true, "remote Go test must sync source by default")
assert(payload["source_sync_mode"] == "runtime", "remote Go test must default to Runtime source sync")
assert(payload["source_sync_entry_count"] == 12, "remote Go test Runtime sync must expose the source entry count")
%w[.dockerignore Dockerfile VERSION go.mod CLAUDE.md cmd internal runtime scripts lib kde docs].each do |entry|
  assert(payload["source_sync_entries"].include?(entry), "remote Go test Runtime sync must include #{entry}")
end
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "remote Go test source root must stay under /home/xnix-*")
assert(payload["remote_source_root"].include?("xnix-remote-go-test-runtime-"), "remote Go test source root must include sync mode")
assert(payload["remote_build_root"].start_with?("/home/xnix-"), "remote Go test build root must stay under /home/xnix-*")
assert(payload["package_count"] == 6, "remote Go test must expose default package count")
assert(payload["packages"] == %w[./internal/runtime/owner ./cmd/xnix-runtime-owner ./cmd/xnix-compat-open ./cmd/xnix-compat-launch ./internal/runtime/appidentity ./cmd/xnix-runtime-go], "remote Go test must expose default package order")
assert(payload["run_regex"].include?("TestCompatOpen"), "remote Go test must expose default test regex")
assert(payload["run_regex"].include?("TestPreviewRealWinAppRunReceiptSummary"), "remote Go test must include real run receipt summary model tests")
assert(payload["run_regex"].include?("TestRealWinAppRunReceiptSummaryPreviewCommand"), "remote Go test must include real run receipt summary CLI tests")
assert(payload["run_regex"].include?("TestLocalGoCompilePolicyPreviewRequiresQ4ByDefault"), "remote Go test must include local Go compile policy model tests")
assert(payload["run_regex"].include?("TestLocalGoCompilePolicyPreviewCommand"), "remote Go test must include local Go compile policy CLI tests")
assert(payload["run_regex"].include?("TestPreviewQ4KnownPortableWinAppRunPlanConsumesRuntimeCatalog"), "remote Go test must include q4 known portable catalog-backed plan model tests")
assert(payload["run_regex"].include?("TestPreviewQ4KnownPortableWinAppRunPlanSelectsPuttySingleExecutableLane"), "remote Go test must include q4 known portable PuTTY single-executable plan model tests")
assert(payload["run_regex"].include?("TestPreviewQ4KnownPortableWinAppRunPlanRejectsKnownNonRunnableCatalogApp"), "remote Go test must include q4 known portable non-runnable rejection model tests")
assert(payload["run_regex"].include?("TestQ4KnownPortableWinAppRunPlanPreviewCommand"), "remote Go test must include q4 known portable catalog-backed plan CLI tests")
assert(payload["run_regex"].include?("TestQ4KnownPortableWinAppRunPlanPreviewCommandRendersPuttySingleExecutableLane"), "remote Go test must include q4 known portable PuTTY single-executable plan CLI tests")
assert(payload["count"] == 1, "remote Go test must default to count 1")
assert(payload["go_test_timeout"] == "5m", "remote Go test must default to a bounded Go test timeout")
assert(payload["remote_test_planned"] == true, "remote Go test must plan remote test")
assert(payload["targeted_test_required"] == true, "remote Go test must mark targeted test intent")
assert(payload["host_compilation_avoided"] == true, "remote Go test must avoid host compilation")
assert(payload["host_root_modified"] == false, "remote Go test must not mutate host root")
assert(payload["privileged_container_required"] == false, "remote Go test must not require privileged containers")
assert(payload["host_networking_required"] == false, "remote Go test must not require host networking")
assert(payload["docker_socket_mounted"] == false, "remote Go test must not mount Docker socket")
assert(payload["broad_host_mount_required"] == false, "remote Go test must not require broad host mounts")

custom_stdout, custom_stderr, custom_status = Open3.capture3(
  "ruby", script.to_s,
  "--source-sync-mode", "full",
  "--remote-source-root", "/tmp/xnix-remote-test/full",
  "--remote-build-root", "/tmp/xnix-remote-cache",
  "--package", "./internal/runtime/owner",
  "--run", "TestServiceCallDispatchesShowRuntimeControlledLaunch",
  "--count", "2",
  "--go-test-timeout", "30s",
  chdir: project_root.to_s
)
assert(custom_status.success?, "remote Go test custom plan must succeed: #{custom_stderr}")
custom_payload = JSON.parse(custom_stdout)
assert(custom_payload["source_sync_mode"] == "full", "remote Go test must expose full source sync mode")
assert(custom_payload["source_sync_entries"] == ["."], "remote Go test full sync must use the full checkout")
assert(custom_payload["remote_source_root"].start_with?("/tmp/xnix-"), "remote Go test must allow constrained /tmp/xnix-* source roots")
assert(custom_payload["remote_build_root"].start_with?("/tmp/xnix-"), "remote Go test must allow constrained /tmp/xnix-* build roots")
assert(custom_payload["package_count"] == 1, "remote Go test must expose custom package count")
assert(custom_payload["packages"].first == "./internal/runtime/owner", "remote Go test must expose custom package")
assert(custom_payload["run_regex"] == "TestServiceCallDispatchesShowRuntimeControlledLaunch", "remote Go test must expose custom regex")
assert(custom_payload["count"] == 2, "remote Go test must expose custom count")
assert(custom_payload["go_test_timeout"] == "30s", "remote Go test must expose custom Go timeout")

bad_package_stdout, bad_package_stderr, bad_package_status = Open3.capture3(
  "ruby", script.to_s,
  "--package", "../internal/runtime/owner",
  chdir: project_root.to_s
)
assert(!bad_package_status.success?, "remote Go test must reject parent-traversal packages")
assert((bad_package_stdout + bad_package_stderr).include?("repository-relative"), "remote Go test must explain unsafe package targets")

bad_root_stdout, bad_root_stderr, bad_root_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-build-root", "/var/tmp/xnix-cache",
  chdir: project_root.to_s
)
assert(!bad_root_status.success?, "remote Go test must reject remote build roots outside constrained prefixes")
assert((bad_root_stdout + bad_root_stderr).include?("remote build root must stay under /home/xnix-* or /tmp/xnix-*"), "remote Go test must explain unsafe build roots")

bad_regex_stdout, bad_regex_stderr, bad_regex_status = Open3.capture3(
  "ruby", script.to_s,
  "--run", "",
  chdir: project_root.to_s
)
assert(!bad_regex_status.success?, "remote Go test must reject empty regex")
assert((bad_regex_stdout + bad_regex_stderr).include?("Go test regex must be non-empty"), "remote Go test must explain empty regex")

puts "PASS: remote Go test script is q4-first and targeted"
