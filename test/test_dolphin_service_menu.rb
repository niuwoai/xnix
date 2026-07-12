#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/compatibility/dolphin_service_menu"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
menu = Xnix::Compatibility::DolphinServiceMenu.new
rendered = menu.render
packaged = project_root.join("kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop").read

assert(menu.file_name == "xnix-open-with-compatibility.desktop", "Dolphin service menu must use a stable file name")
assert(rendered == packaged, "Dolphin service menu renderer must match the packaged file")
assert(rendered.include?("X-KDE-ServiceTypes=KonqPopupMenu/Plugin"), "Dolphin service menu must target Dolphin popup integration")
assert(rendered.include?("Exec=xnix-compat-open %U"), "Dolphin service menu must delegate to the Runtime file-open entry point")
assert(!rendered.downcase.include?("wine"), "Dolphin service menu must not expose backend implementation names")
assert(!rendered.downcase.include?("prefix"), "Dolphin service menu must not expose backend storage names")
assert(!rendered.include?(".exe"), "Dolphin service menu must not expose Windows executable paths")

puts "PASS: Dolphin service menu unit tests"
