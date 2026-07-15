#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/kde_shell_integration_plan"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
plan = Xnix::Compatibility::KDEShellIntegrationPlan.new.to_h
component_ids = plan.fetch("components").map { |component| component.fetch("id") }
runtime_methods = plan.fetch("components").map { |component| component.fetch("runtime_method") }

assert(plan["version"] == "0.2.167", "KDE shell integration plan must expose the current version")
assert(plan["plan_type"] == "kde-shell-integration-plan", "KDE shell integration plan must identify the plan type")
assert(plan["runtime_method"] == "GetKDEShellIntegrationPlan", "KDE shell integration plan must expose the Runtime method")
assert(plan["desktop_shell"] == "KDE Plasma", "KDE shell integration plan must target KDE Plasma")
assert(plan["official_desktop_only"], "KDE shell integration plan must keep KDE as the official first shell")
assert(!plan["fallback_desktops_supported"], "KDE shell integration plan must not claim fallback desktop support")
assert(plan["runtime_owned"], "KDE shell integration plan must be Runtime-owned")
assert(plan["c_runtime_backed"], "KDE shell integration plan must be backed by the C Runtime")
assert(!plan["kde_policy_owner"], "KDE must not own Runtime policy")
assert(!plan["plasma_fork_required"], "KDE shell integration plan must not require a Plasma fork")
assert(!plan["plasma_source_modified"], "KDE shell integration plan must not modify Plasma source")
assert(!plan["shell_configuration_written"], "KDE shell integration plan must not write shell configuration")
assert(!plan["component_activation_enabled"], "KDE shell integration plan must keep component activation disabled")
assert(!plan["backend_launch_enabled"], "KDE shell integration plan must not launch backends")
assert(!plan["backend_details_exposed"], "KDE shell integration plan must hide backend details")
assert(!plan["host_root_modified"], "KDE shell integration plan must not mutate the host root")
assert(!plan["privileged_container_required"], "KDE shell integration plan must not need privileged containers")
assert(plan["component_count"] == 9, "KDE shell integration plan must count shell components")
assert(plan["initial_component_count"] == 8, "KDE shell integration plan must count initial components")
assert(plan["planned_component_count"] == 1, "KDE shell integration plan must count planned components")
assert(component_ids == %w[start-menu task-manager file-manager system-tray notification-center compatibility-center unified-settings krunner-search kwin-window-management], "KDE shell integration plan must expose expected components")
assert(runtime_methods.include?("GetDesktopEntryPlan"), "KDE shell integration plan must reference desktop entry planning")
assert(runtime_methods.include?("GetCompatibilitySettings"), "KDE shell integration plan must reference settings planning")
assert(plan.fetch("components").all? { |component| component.fetch("runtime_owned") }, "KDE shell components must be Runtime-owned")
assert(plan.fetch("components").all? { |component| !component.fetch("kde_policy_owner") }, "KDE shell components must not be KDE policy-owned")
assert(plan.fetch("components").all? { |component| !component.fetch("shell_writes_enabled") }, "KDE shell components must not write shell configuration")
assert(plan.fetch("components").all? { |component| !component.fetch("backend_launch_enabled") }, "KDE shell components must not launch backends")
assert(!JSON.generate(plan).match?(/wine|prefix|\.wine|proton|virtual machine/i), "KDE shell integration plan must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-kde-shell-integration-plan").to_s)
assert(status.success?, "KDE shell integration plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "KDE shell integration plan CLI must emit the plan model")

Dir.mktmpdir("xnix-kde-shell-integration-plan") do |dir|
  binary = Pathname.new(dir).join("xnix-runtime-core")
  _stdout, stderr, status = Open3.capture3(
    "cc",
    "-std=c11",
    "-Wall",
    "-Wextra",
    "-Werror",
    "-Iruntime/core",
    "runtime/core/xnix_runtime_core.c",
    "runtime/core/xnix_runtime_core_cli.c",
    "-o",
    binary.to_s
  )
  assert(status.success?, "C Runtime core must compile cleanly: #{stderr}")

  stdout, stderr, status = Open3.capture3(binary.to_s, "kde-shell-integration-plan")
  assert(status.success?, "C Runtime core KDE shell integration plan must exit successfully: #{stderr}")
  c_plan = JSON.parse(stdout)
  assert(c_plan["version"] == plan["version"], "C Runtime core KDE shell integration plan must expose the current version")
  assert(c_plan["plan_type"] == plan["plan_type"], "C Runtime core KDE shell integration plan must identify the plan type")
  assert(c_plan["runtime_method"] == plan["runtime_method"], "C Runtime core KDE shell integration plan must expose the Runtime method")
  assert(c_plan["component_count"] == 9, "C Runtime core KDE shell integration plan must count components")
  assert(c_plan["initial_component_count"] == 8, "C Runtime core KDE shell integration plan must count initial components")
  assert(c_plan["planned_component_count"] == 1, "C Runtime core KDE shell integration plan must count planned components")
  assert(c_plan.fetch("components").map { |component| component.fetch("id") } == component_ids, "C Runtime core KDE shell integration plan must expose expected components")
  assert(!c_plan["plasma_fork_required"], "C Runtime core KDE shell integration plan must not require a Plasma fork")
  assert(!c_plan["shell_configuration_written"], "C Runtime core KDE shell integration plan must not write shell configuration")
  assert(!c_plan["component_activation_enabled"], "C Runtime core KDE shell integration plan must keep activation disabled")
  assert(!c_plan["backend_launch_enabled"], "C Runtime core KDE shell integration plan must not launch backends")
  assert(!c_plan["backend_details_exposed"], "C Runtime core KDE shell integration plan must hide backend details")
  assert(!stdout.match?(/wine|prefix|\.wine|proton|virtual machine/i), "C Runtime core KDE shell integration plan must not expose backend implementation terms")
end

puts "PASS: KDE shell integration plan unit tests"
