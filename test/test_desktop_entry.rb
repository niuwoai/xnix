#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/compatibility/application_recipe"
require_relative "../lib/xnix/compatibility/desktop_entry"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

recipe = Xnix::Compatibility::ApplicationRecipe.new(
  id: "org.example.ledger",
  name: "Example Ledger",
  icon: "office-chart-area",
  mode: "wine",
  supported_extensions: [".abc", ".xls"]
)
entry = Xnix::Compatibility::DesktopEntry.new(recipe)
contents = entry.render

assert(entry.file_name == "xnix-org.example.ledger.desktop", "desktop file must use the stable application id")
assert(contents.include?("Exec=xnix-compat-launch --app org.example.ledger %U"), "desktop entry must call the managed launcher")
assert(contents.include?("MimeType=application/x-xnix-abc;application/x-xnix-xls;"), "desktop entry must declare recipe MIME types")
assert(!contents.downcase.include?("wine"), "desktop entry must not expose the Wine implementation")
assert(!contents.downcase.include?("prefix"), "desktop entry must not expose implementation storage")
assert(!contents.include?(".exe"), "desktop entry must not expose a Windows executable path")

container_gui_recipe = Xnix::Compatibility::ApplicationRecipe.new(
  id: "org.xnix.sample.notepad",
  name: "Sample Notepad",
  icon: "accessories-text-editor",
  mode: "automatic",
  supported_extensions: [".txt"],
  container_gui_smoke: {
    "app" => "notepad.exe",
    "window_match" => "notepad.exe"
  }
)
container_gui_entry = Xnix::Compatibility::DesktopEntry.new(container_gui_recipe).render
assert(container_gui_entry.include?("Exec=xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U"), "container GUI desktop entry must route through the packaged recipe registry")
assert(!container_gui_entry.include?("notepad.exe"), "container GUI desktop entry must not expose the raw Windows executable")
assert(!container_gui_entry.downcase.include?("wine"), "container GUI desktop entry must not expose the backend")

puts "PASS: compatibility desktop entry unit tests"
