#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
guarded_scripts = [
  ["staged launcher smoke", project_root.join("scripts/staged_launcher_smoke.rb")],
  ["staged launcher dispatch smoke", project_root.join("scripts/staged_launcher_dispatch_smoke.rb")],
  ["staged desktop Notepad smoke", project_root.join("scripts/staged_desktop_notepad_smoke.rb")],
  ["staged desktop external Windows app smoke", project_root.join("scripts/staged_desktop_external_winapp_smoke.rb")]
]

guarded_scripts.each do |label, script|
  source = script.read
  assert(source.include?("XNIX_ALLOW_LOCAL_GO_COMPILE"), "#{label} must expose the explicit local Go compile override")
  assert(source.include?("local Go compilation is disabled by default"), "#{label} must default away from local Go compilation")
  assert(source.include?("scripts/remote_go_build.rb --execute"), "#{label} must point compile-heavy work to q4")
end

staged_external_source = project_root.join("scripts/staged_desktop_external_winapp_smoke.rb").read
assert(
  staged_external_source.include?('"windows_process_file_argument_window_observed" => file_open_lane &&'),
  "staged desktop external smoke must only claim file-argument window evidence on file-open lanes"
)

minimal_env = {
  "PATH" => "/usr/bin:/bin",
  "XNIX_ALLOW_LOCAL_GO_COMPILE" => nil,
  "XNIX_COMPAT_LAUNCH_BIN" => nil
}

guarded_scripts.each do |label, script|
  stdout, stderr, status = Open3.capture3(minimal_env, "ruby", script.to_s)
  assert(status.success?, "#{label} must skip cleanly without local Go compile override: #{stderr}")
  assert(stdout.include?("SKIP:"), "#{label} must print a SKIP marker without local Go compile override")
  assert(stdout.include?("local Go compilation is disabled by default"), "#{label} must explain the q4-first compile guard")
end

puts "PASS: local Go compile guarded smoke scripts"
