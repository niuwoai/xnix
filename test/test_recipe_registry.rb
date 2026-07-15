#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/recipe_registry"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
registry_path = project_root.join("runtime/recipes/registry.json")
report = Xnix::Compatibility::RecipeRegistry.new(path: registry_path).verify

assert(report["version"] == "0.2.199", "recipe registry must expose the current version")
assert(report["schema_version"] == 1, "recipe registry must expose the supported schema version")
assert(report["registry_name"] == "xnix-local-development", "recipe registry must identify the development registry")
assert(report["recipe_count"] == 1, "recipe registry must report the sample recipe")
assert(report["trust"]["digest_verified"], "recipe registry must verify recipe digests")
assert(!report["trust"]["signed_recipe_validation"], "recipe registry must not claim signed recipe validation yet")
assert(report["trust"]["development_registry"], "recipe registry must mark the current registry as development-only")

recipe = report.fetch("recipes").first
assert(recipe["id"] == "org.xnix.sample.notepad", "recipe registry must preserve recipe ids")
assert(recipe["digest_verified"], "recipe registry entries must verify digests")
assert(recipe["sha256"] == Digest::SHA256.file(project_root.join("runtime/recipes/org.xnix.sample.notepad.json")).hexdigest, "recipe registry must report the actual digest")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-recipe-registry").to_s, "verify")
assert(status.success?, "recipe registry CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == report, "recipe registry CLI must emit the verification report")

Dir.mktmpdir("xnix-recipe-registry") do |dir|
  recipe_path = Pathname.new(dir).join("org.example.bad.json")
  File.write(
    recipe_path,
    JSON.pretty_generate(
      "id" => "org.example.bad",
      "name" => "Bad Example",
      "icon" => "application-x-executable",
      "mode" => "automatic",
      "supported_extensions" => [".bad"]
    )
  )
  registry = Pathname.new(dir).join("registry.json")
  File.write(
    registry,
    JSON.pretty_generate(
      "schema_version" => 1,
      "registry_name" => "test",
      "recipes" => [
        {
          "id" => "org.example.bad",
          "path" => "org.example.bad.json",
          "sha256" => "0" * 64,
          "signature_status" => "development-only"
        }
      ]
    )
  )

  _stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-recipe-registry").to_s,
    "--registry",
    registry.to_s,
    "verify"
  )
  assert(!status.success?, "recipe registry CLI must reject digest mismatches")
  assert(stderr.include?("recipe digest mismatch"), "recipe registry CLI must explain digest mismatches")
end

Dir.mktmpdir("xnix-recipe-registry") do |dir|
  registry = Pathname.new(dir).join("registry.json")
  File.write(
    registry,
    JSON.pretty_generate(
      "schema_version" => 1,
      "registry_name" => "test",
      "recipes" => [
        {
          "id" => "org.example.escape",
          "path" => "../escape.json",
          "sha256" => "0" * 64,
          "signature_status" => "development-only"
        }
      ]
    )
  )

  _stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-recipe-registry").to_s,
    "--registry",
    registry.to_s,
    "verify"
  )
  assert(!status.success?, "recipe registry CLI must reject escaping paths")
  assert(stderr.include?("recipe path must be a relative JSON file"), "recipe registry CLI must explain unsafe paths")
end

puts "PASS: compatibility recipe registry unit tests"
