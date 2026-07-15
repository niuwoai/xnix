#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "find"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
SOURCE_LIBEXEC = PROJECT_ROOT.join("libexec/xnix/compatd")
SOURCE_RUNTIME_LIB = PROJECT_ROOT.join("lib/xnix")
SOURCE_VERSION = PROJECT_ROOT.join("VERSION")
SOURCE_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes")
SOURCE_DBUS_CONTRACT = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.xml")
SOURCE_DBUS_SMOKE = PROJECT_ROOT.join("runtime/dbus/xnix_compatd_smoke.c")
SOURCE_DBUS_SMOKE_INTROSPECTION = PROJECT_ROOT.join("runtime/dbus/xnix_compatd_introspection.inc")
SOURCE_DBUS_SMOKE_KDE_CENTER = PROJECT_ROOT.join("runtime/dbus/xnix_compatd_kde_center.inc")
SOURCE_SESSION_SMOKE = PROJECT_ROOT.join("scripts/dbus_session_smoke.rb")
SOURCE_DBUS_SERVICE = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.service")
SOURCE_SYSTEMD_UNIT = PROJECT_ROOT.join("runtime/systemd/xnix-compatd.service")

options = {
  root: nil
}

OptionParser.new do |parser|
  parser.banner = "Usage: install_runtime_activation.rb --root PATH"
  parser.on("--root PATH", "Install activation files under PATH") do |value|
    options[:root] = Pathname.new(value)
  end
end.parse!

abort "Usage: install_runtime_activation.rb --root PATH" unless options[:root]

root = options[:root]
install_map = {
  SOURCE_LIBEXEC => root.join("usr/libexec/xnix/compatd"),
  SOURCE_VERSION => root.join("usr/VERSION"),
  SOURCE_DBUS_CONTRACT => root.join("usr/runtime/dbus/org.xnix.Compatibility1.xml"),
  SOURCE_DBUS_SMOKE => root.join("usr/runtime/dbus/xnix_compatd_smoke.c"),
  SOURCE_DBUS_SMOKE_INTROSPECTION => root.join("usr/runtime/dbus/xnix_compatd_introspection.inc"),
  SOURCE_DBUS_SMOKE_KDE_CENTER => root.join("usr/runtime/dbus/xnix_compatd_kde_center.inc"),
  SOURCE_SESSION_SMOKE => root.join("usr/scripts/dbus_session_smoke.rb"),
  SOURCE_DBUS_SERVICE => root.join("usr/share/dbus-1/system-services/org.xnix.Compatibility1.service"),
  SOURCE_SYSTEMD_UNIT => root.join("usr/lib/systemd/system/xnix-compatd.service")
}

install_map.each do |source, destination|
  FileUtils.mkdir_p(destination.dirname)
  FileUtils.install(source, destination, mode: source.executable? ? 0o755 : 0o644)
end

Find.find(SOURCE_RUNTIME_LIB) do |source|
  source_path = Pathname.new(source)
  next if source_path.directory?

  relative_path = source_path.relative_path_from(SOURCE_RUNTIME_LIB)
  destination = root.join("usr/lib/xnix", relative_path)
  FileUtils.mkdir_p(destination.dirname)
  FileUtils.install(source_path, destination, mode: source_path.executable? ? 0o755 : 0o644)
end

Find.find(SOURCE_RECIPE_DIR) do |source|
  source_path = Pathname.new(source)
  next if source_path.directory?

  relative_path = source_path.relative_path_from(SOURCE_RECIPE_DIR)
  destination = root.join("usr/runtime/recipes", relative_path)
  FileUtils.mkdir_p(destination.dirname)
  FileUtils.install(source_path, destination, mode: 0o644)
end

puts "Installed Xnix Runtime activation files under #{root}"
