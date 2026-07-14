#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/recipe_trust_policy"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
registry_report = Xnix::Compatibility::RecipeRegistry.new(
  path: project_root.join("runtime/recipes/registry.json")
).verify
policy = Xnix::Compatibility::RecipeTrustPolicy.new(registry_report: registry_report).to_h

assert(policy["version"] == "0.2.142", "recipe trust policy must expose the current version")
assert(policy["policy_type"] == "recipe-trust", "recipe trust policy must identify the model type")
assert(policy["decision"] == "development-only", "development registry must not be production-trusted")
assert(policy["recipe_count"] == 1, "recipe trust policy must report recipe count")
assert(policy["checks"].any? { |check| check["id"] == "registry.digest" && check["status"] == "pass" }, "recipe trust policy must pass digest checks")
assert(policy["checks"].any? { |check| check["id"] == "registry.signature" && check["status"] == "pending" }, "recipe trust policy must keep production signatures pending")
assert(policy["blocking_reasons"].include?("production signed recipe validation is not enabled"), "recipe trust policy must block production trust without signatures")
assert(policy["desktop_safe_summary"] == "Recipes are verified for local development only.", "recipe trust policy must expose a desktop-safe summary")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-recipe-trust-policy").to_s, "evaluate")
assert(status.success?, "recipe trust policy CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == policy, "recipe trust policy CLI must emit the policy model")

trusted_report = {
  "version" => "0.2.142",
  "schema_version" => 1,
  "registry_name" => "trusted",
  "registry_path" => "/dev/null",
  "recipe_count" => 1,
  "recipes" => [],
  "trust" => {
    "digest_verified" => true,
    "signed_recipe_validation" => true,
    "development_registry" => false
  }
}
trusted_policy = Xnix::Compatibility::RecipeTrustPolicy.new(registry_report: trusted_report).to_h
assert(trusted_policy["decision"] == "production-trusted", "verified signatures must allow production trust")
assert(trusted_policy["blocking_reasons"].empty?, "production-trusted recipes must not have blocking reasons")

Dir.mktmpdir("xnix-recipe-trust") do |dir|
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

  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-recipe-trust-policy").to_s,
    "--registry",
    registry_path.to_s,
    "evaluate"
  )
  assert(status.success?, "recipe trust policy CLI must evaluate verified registries: #{stderr}")
  assert(JSON.parse(stdout)["decision"] == "production-trusted", "verified registry must evaluate as production-trusted")
end

puts "PASS: compatibility recipe trust policy unit tests"
