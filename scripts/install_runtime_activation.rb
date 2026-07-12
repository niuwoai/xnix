#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
SOURCE_LIBEXEC = PROJECT_ROOT.join("libexec/xnix/compatd")
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
  SOURCE_DBUS_SERVICE => root.join("usr/share/dbus-1/system-services/org.xnix.Compatibility1.service"),
  SOURCE_SYSTEMD_UNIT => root.join("usr/lib/systemd/system/xnix-compatd.service")
}

install_map.each do |source, destination|
  FileUtils.mkdir_p(destination.dirname)
  FileUtils.install(source, destination, mode: source.executable? ? 0o755 : 0o644)
end

puts "Installed Xnix Runtime activation files under #{root}"
