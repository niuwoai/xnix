# frozen_string_literal: true

require "minitest/autorun"
require "pathname"

ROOT = Pathname.new(__dir__).join("..").realpath

class KDEControlledLaunchActionStubTest < Minitest::Test
  def test_desktop_action_declares_evidence_only_runtime_route
    contents = ROOT.join("kde/actions/xnix-runtime-status-controlled-launch.desktop").read

    assert_includes contents, "Name=Xnix Runtime controlled launch"
    assert_includes contents, "X-Xnix-KDE-Action-ID=xnix.runtime-status.controlled-launch"
    assert_includes contents, "X-Xnix-Runtime-Preview=xnix-runtime-go kde-controlled-launch-action-preview"
    assert_includes contents, "X-Xnix-Restricted-Smoke-Plan=xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview"
    assert_includes contents, "X-Xnix-DBus-Service=org.xnix.Compatibility1"
    assert_includes contents, "X-Xnix-DBus-Object-Path=/org/xnix/Compatibility1"
    assert_includes contents, "X-Xnix-DBus-Method=org.xnix.Compatibility1.ShowRuntimeControlledLaunch"
    assert_includes contents, "X-Xnix-Forwarded-Argument=evidence-relative-path"
    assert_includes contents, "X-Xnix-Forwards-Only-Evidence-Handle=true"
    assert_includes contents, "X-Xnix-KDE-Policy-Owner=false"
    assert_includes contents, "X-Xnix-Owner-Service-Args-Exposed-To-KDE=false"
    assert_includes contents, "X-Xnix-State-Root-Access=false"
    assert_includes contents, "X-Xnix-Receipt-Reconstruction=false"
    assert_includes contents, "X-Xnix-Backend-Launch-Enabled=false"
    assert_includes contents, "X-Xnix-Execution-Started=false"
    assert_includes contents, "X-Xnix-Host-Root-Modified=false"
    refute_includes contents, "--state-root"
    refute_includes contents, "owner_service_call_args"
  end
end
