#!/usr/bin/env ruby
# frozen_string_literal: true

require "tmpdir"
require "json"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

Dir.mktmpdir("xnix-recipes") do |dir|
  recipe_path = File.join(dir, "org.example.ledger.json")
  registry_path = File.join(dir, "registry.json")
  File.write(
    recipe_path,
    JSON.pretty_generate(
      "id" => "org.example.ledger",
      "name" => "Example Ledger",
      "icon" => "office-chart-area",
      "mode" => "automatic",
      "supported_extensions" => [".abc"]
    )
  )
  File.write(
    registry_path,
    JSON.pretty_generate(
      "schema_version" => 1,
      "registry_name" => "test",
      "recipes" => []
    )
  )

  store = Xnix::Compatibility::RecipeStore.new(path: dir)
  recipes = store.all

  assert(recipes.length == 1, "recipe store must load JSON recipes")
  assert(recipes.first.id == "org.example.ledger", "recipe store must preserve recipe ids")
  assert(store.find("org.example.ledger").name == "Example Ledger", "recipe store must find recipes by id")
  assert(store.find("org.example.missing").nil?, "recipe store must return nil for missing recipes")
end

puts "PASS: compatibility recipe store unit tests"
