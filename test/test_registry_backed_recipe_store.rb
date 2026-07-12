#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "json"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/registry_backed_recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
store = Xnix::Compatibility::RegistryBackedRecipeStore.for_path(project_root.join("runtime/recipes"))
recipes = store.all

assert(store.is_a?(Xnix::Compatibility::RegistryBackedRecipeStore), "default recipe directory must use the registry-backed store")
assert(store.registry_report["trust"]["digest_verified"], "registry-backed store must verify digests before loading recipes")
assert(recipes.length == 1, "registry-backed store must load registered recipes")
assert(recipes.first.id == "org.xnix.sample.notepad", "registry-backed store must expose registered recipe ids")
assert(store.find("org.xnix.sample.notepad").name == "Sample Notepad", "registry-backed store must find registered recipes")

Dir.mktmpdir("xnix-registry-store") do |dir|
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
  File.write(
    Pathname.new(dir).join("registry.json"),
    JSON.pretty_generate(
      "schema_version" => 1,
      "registry_name" => "test",
      "recipes" => [
        {
          "id" => "org.example.good",
          "path" => "org.example.good.json",
          "sha256" => Digest::SHA256.file(recipe_path).hexdigest,
          "signature_status" => "development-only"
        }
      ]
    )
  )

  registered_store = Xnix::Compatibility::RegistryBackedRecipeStore.for_path(dir)
  assert(registered_store.all.first.id == "org.example.good", "registry-backed store must load valid registered recipes")
end

Dir.mktmpdir("xnix-registry-store") do |dir|
  recipe_path = Pathname.new(dir).join("org.example.loose.json")
  File.write(
    recipe_path,
    JSON.pretty_generate(
      "id" => "org.example.loose",
      "name" => "Loose Example",
      "icon" => "application-x-executable",
      "mode" => "automatic",
      "supported_extensions" => [".loose"]
    )
  )

  fallback_store = Xnix::Compatibility::RegistryBackedRecipeStore.for_path(dir)
  assert(fallback_store.is_a?(Xnix::Compatibility::RecipeStore), "missing registry must keep the development fallback store")
  assert(fallback_store.all.first.id == "org.example.loose", "development fallback store must still load local test recipes")
end

Dir.mktmpdir("xnix-registry-store") do |dir|
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
  File.write(
    Pathname.new(dir).join("registry.json"),
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

  bad_store = Xnix::Compatibility::RegistryBackedRecipeStore.for_path(dir)
  begin
    bad_store.all
    assert(false, "registry-backed store must reject digest mismatches")
  rescue ArgumentError => e
    assert(e.message.include?("recipe digest mismatch"), "registry-backed store must explain digest mismatches")
  end
end

puts "PASS: compatibility registry-backed recipe store unit tests"
