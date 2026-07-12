#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/compatibility/application_recipe"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

recipe = Xnix::Compatibility::ApplicationRecipe.new(
  id: "org.example.ledger",
  name: "Example Ledger",
  icon: "office-chart-area",
  mode: "automatic",
  supported_extensions: [".abc", ".xls"]
)

assert(recipe.mime_types == ["application/x-xnix-abc", "application/x-xnix-xls"], "recipe must derive scoped MIME types")

invalid_ids = ["ledger", "org.example.Ledger", "org.example.ledger;rm"]
invalid_ids.each do |identifier|
  begin
    Xnix::Compatibility::ApplicationRecipe.new(
      id: identifier,
      name: "Example",
      icon: "application-x-executable",
      mode: "automatic"
    )
    assert(false, "invalid identifier must be rejected: #{identifier}")
  rescue ArgumentError
    nil
  end
end

begin
  Xnix::Compatibility::ApplicationRecipe.new(
    id: "org.example.invalid-extension",
    name: "Example",
    icon: "application-x-executable",
    mode: "automatic",
    supported_extensions: [".abc;rm"]
  )
  assert(false, "unsafe file extension must be rejected")
rescue ArgumentError
  nil
end

puts "PASS: compatibility application recipe unit tests"
