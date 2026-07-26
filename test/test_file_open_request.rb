#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/file_open_request"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
request = Xnix::Compatibility::FileOpenRequest.new(recipe_store: store).build(
  file_uris: ["file:///home/test/Documents/example.txt"]
)

assert(request["request_type"] == "open-file", "file open request must identify the request type")
assert(request["source"] == "dolphin-service-menu", "file open request must identify the Dolphin source")
assert(request["application_id"] == "org.xnix.sample.notepad", "file open request must resolve applications by extension")
assert(request["runtime_method"] == "Launch", "file open request must target the Runtime launch method")
assert(request["portal_required"], "file open request must require portal-mediated access")
assert(request["file_count"] == 1, "file open request must count selected files")
assert(request["file_uris"].first == "file:///home/test/Documents/example.txt", "file open request must preserve file URIs")

json = JSON.pretty_generate(request)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "file open request must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-open").to_s,
  "--preview-engine",
  "ruby",
  "file:///home/test/Documents/example.txt"
)
assert(status.success?, "file open CLI must exit successfully: #{stderr}")
cli_request = JSON.parse(stdout)
assert(cli_request == request, "file open CLI must emit the request model")

fake_runtime = Pathname.new(Dir.mktmpdir("xnix-file-open-runtime")).join("xnix-runtime-go")
fake_runtime.write(<<~RUBY)
  #!/usr/bin/env ruby
  # frozen_string_literal: true

  require "json"

  puts JSON.pretty_generate(
    "schema_version" => "xnix.runtime.file_open.v1",
    "request_type" => "file-open-preview",
    "source" => "dolphin-service-menu",
    "runtime_args" => ARGV,
    "runtime_owned" => true,
    "backend_launch_enabled" => false,
    "host_root_modified" => false
  )
RUBY
fake_runtime.chmod(0o700)
stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-open").to_s,
  "--runtime-bin",
  fake_runtime.to_s,
  "--runtime-registry",
  project_root.join("runtime/recipes/registry.json").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "file:///home/test/Documents/example.txt"
)
assert(status.success?, "file open CLI must delegate to Go Runtime preview when available: #{stderr}")
go_request = JSON.parse(stdout)
assert(go_request["request_type"] == "file-open-preview", "file open CLI must emit the Go Runtime file-open preview")
assert(go_request["runtime_args"].include?("file-open-preview"), "file open CLI must call the Go Runtime file-open preview command")
assert(go_request["runtime_args"].include?("--registry"), "file open CLI must pass the Runtime registry to Go")
assert(go_request["runtime_args"].include?("--app"), "file open CLI must pass explicit app ids to Go")
assert(go_request["runtime_args"].include?("file:///home/test/Documents/example.txt"), "file open CLI must pass file URIs to Go")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-open").to_s, "--preview-engine", "ruby", "https://example.invalid/file.txt")
assert(!status.success?, "file open CLI must reject non-file URIs")
assert(stderr.include?("only file URIs"), "file open CLI must explain rejected URI schemes")

puts "PASS: compatibility file open request unit tests"
