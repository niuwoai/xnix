#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/launch_request"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
request = Xnix::Compatibility::LaunchRequest.new(recipe_store: store).build(
  application_id: "org.xnix.sample.notepad"
)

assert(request["request_type"] == "launch-application", "launch request must identify the request type")
assert(request["source"] == "desktop-launcher", "launch request must identify the desktop source")
assert(request["application_id"] == "org.xnix.sample.notepad", "launch request must keep the Runtime application id")
assert(request["runtime_method"] == "Launch", "launch request must target the Runtime launch method")
assert(!request["portal_required"], "plain launcher requests must not require file portal access")
assert(request["file_count"] == 0, "plain launcher requests must not include files")

file_request = Xnix::Compatibility::LaunchRequest.new(recipe_store: store).build(
  application_id: "org.xnix.sample.notepad",
  file_uris: ["file:///home/test/Documents/example.txt"]
)
assert(file_request["portal_required"], "launcher file requests must require portal-mediated access")
assert(file_request["file_count"] == 1, "launcher file requests must count selected files")
assert(file_request["file_uris"].first == "file:///home/test/Documents/example.txt", "launcher file requests must preserve file URIs")

json = JSON.pretty_generate(file_request)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "launch request must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-launch").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "file:///home/test/Documents/example.txt"
)
assert(status.success?, "launch CLI must exit successfully: #{stderr}")
cli_request = JSON.parse(stdout)
assert(cli_request == file_request, "launch CLI must emit the request model")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-launch").to_s)
assert(!status.success?, "launch CLI must require an application id")
assert(stderr.include?("--app is required"), "launch CLI must explain missing application ids")

puts "PASS: compatibility launch request unit tests"
