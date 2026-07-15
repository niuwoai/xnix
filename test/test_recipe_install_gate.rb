#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/recipe_install_gate"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def write_verified_registry(dir)
  recipe_path = Pathname.new(dir).join("org.example.good.json")
  File.write(
    recipe_path,
    JSON.pretty_generate(
      "id" => "org.example.good",
      "name" => "Good Example",
      "icon" => "application-x-executable",
      "mode" => "automatic",
      "supported_extensions" => [".good"]
    )
  )

  registry_path = Pathname.new(dir).join("registry.json")
  File.write(
    registry_path,
    JSON.pretty_generate(
      "schema_version" => 1,
      "registry_name" => "test",
      "recipes" => [
        {
          "id" => "org.example.good",
          "path" => "org.example.good.json",
          "sha256" => Digest::SHA256.file(recipe_path).hexdigest,
          "signature_status" => "verified"
        }
      ]
    )
  )

  registry_path
end

project_root = Pathname.new(__dir__).join("..").realpath
registry_report = Xnix::Compatibility::RecipeRegistry.new(
  path: project_root.join("runtime/recipes/registry.json")
).verify

production_gate = Xnix::Compatibility::RecipeInstallGate.new(
  registry_report: registry_report,
  application_id: "org.xnix.sample.notepad"
).to_h
assert(production_gate["version"] == "0.2.212", "recipe install gate must expose the current version")
assert(production_gate["gate_type"] == "recipe-install", "recipe install gate must identify the model type")
assert(production_gate["mode"] == "production", "recipe install gate must default to production mode")
assert(production_gate["decision"] == "block", "production mode must block development-only registries")
assert(production_gate["policy_decision"] == "development-only", "production mode must surface the trust policy decision")
assert(production_gate["matched_recipe"]["digest_verified"], "recipe install gate must expose digest verification")
assert(production_gate["blocking_reasons"].include?("production signed recipe validation is not enabled"), "production mode must require signed recipe validation")
assert(production_gate["blocking_reasons"].include?("registry contains development-only recipes"), "production mode must block development-only recipes")

development_gate = Xnix::Compatibility::RecipeInstallGate.new(
  registry_report: registry_report,
  application_id: "org.xnix.sample.notepad",
  mode: "development"
).to_h
assert(development_gate["decision"] == "allow", "development mode must allow digest-verified local recipes")
assert(development_gate["requirements"].empty?, "allowed development recipes must not list requirements")

unknown_gate = Xnix::Compatibility::RecipeInstallGate.new(
  registry_report: registry_report,
  application_id: "org.xnix.unknown"
).to_h
assert(unknown_gate["decision"] == "block", "unknown application recipes must be blocked")
assert(unknown_gate["matched_recipe"].nil?, "unknown application recipes must not report a match")
assert(unknown_gate["blocking_reasons"].include?("recipe is not registered"), "unknown application recipes must explain the registry miss")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-recipe-install-gate").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "evaluate"
)
assert(status.success?, "recipe install gate CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == production_gate, "recipe install gate CLI must emit the production gate model")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-recipe-install-gate").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--mode",
  "development",
  "evaluate"
)
assert(status.success?, "recipe install gate CLI must evaluate development mode: #{stderr}")
assert(JSON.parse(stdout)["decision"] == "allow", "recipe install gate CLI must allow development staging")

Dir.mktmpdir("xnix-recipe-install-gate") do |dir|
  registry_path = write_verified_registry(dir)
  report = Xnix::Compatibility::RecipeRegistry.new(path: registry_path).verify
  trusted_gate = Xnix::Compatibility::RecipeInstallGate.new(
    registry_report: report,
    application_id: "org.example.good"
  ).to_h
  assert(trusted_gate["decision"] == "allow", "verified recipe registry must allow production installation")
  assert(trusted_gate["policy_decision"] == "production-trusted", "verified registry must surface production trust")

  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-recipe-install-gate").to_s,
    "--registry",
    registry_path.to_s,
    "--app",
    "org.example.good",
    "evaluate"
  )
  assert(status.success?, "recipe install gate CLI must evaluate verified registries: #{stderr}")
  assert(JSON.parse(stdout)["decision"] == "allow", "verified registry CLI evaluation must allow production installation")
end

desktop_json = JSON.pretty_generate(development_gate)
%w[wine prefix .wine proton virtual\ machine].each do |term|
  assert(!desktop_json.downcase.include?(term.delete("\\")), "recipe install gate desktop model must not expose #{term}")
end

puts "PASS: compatibility recipe install gate unit tests"
