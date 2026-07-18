#!/usr/bin/env ruby
# frozen_string_literal: true

# Boot smoke for a produced Xnix KDE Plasma disk image. Boots the image
# under QEMU and checks the serial log for the graphical-login markers
# declared in the manifest. Requires a real disk image and UEFI firmware;
# it never fakes a boot.
#
# Usage:
#   ruby scripts/boot_kde_image.rb --disk out/xnix-kinoite.qcow2 \
#                                  --firmware /usr/share/edk2/ovmf/OVMF_CODE.fd

require "open3"
require "fileutils"
require "pathname"
require_relative "../lib/xnix/image/kde_image"
require_relative "../lib/xnix/image/boot_smoke"
require_relative "../lib/xnix/image/disk_build"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def arg_value(flag)
  index = ARGV.index(flag)
  index ? ARGV[index + 1] : nil
end

image = Xnix::Image::KdeImage.new(project_root: PROJECT_ROOT.to_s)
smoke = Xnix::Image::BootSmoke.new(image)

# Default the disk path to what the disk-build step produces, so the
# pipeline connects end-to-end without repeating paths.
default_output = Pathname.new(arg_value("--output") || PROJECT_ROOT.join("output")).expand_path
disk = arg_value("--disk") ||
       Xnix::Image::DiskBuild.new(project_root: PROJECT_ROOT.to_s).output_path("qcow2", default_output.to_s)
firmware = arg_value("--firmware") || "/usr/share/edk2/ovmf/OVMF_CODE.fd"
serial_log = Pathname.new(arg_value("--serial-log") || default_output.join("serial-smoke.log")).expand_path
acceleration = File.readable?("/dev/kvm") && File.writable?("/dev/kvm") ? "kvm" : "tcg"

abort "FAIL: disk image not found: #{disk}" unless File.file?(disk)
abort "FAIL: UEFI firmware not found: #{firmware}" unless File.file?(firmware)

command = [
  "timeout", "#{smoke.timeout_seconds}s",
  *smoke.boot_command(disk_path: disk, firmware_path: firmware, acceleration: acceleration)
]
puts "RUN: #{command.join(' ')}"
puts "LOG: #{serial_log}"

FileUtils.mkdir_p(serial_log.dirname)
serial = +""
File.open(serial_log, "w") do |log|
  Open3.popen2e(*command) do |stdin, output, _wait_thread|
    stdin.close
    output.each do |chunk|
      serial << chunk
      log.write(chunk)
      log.flush
      print chunk
    end
  end
end

if smoke.booted?(serial)
  puts "PASS: #{smoke.summary(serial)}"
  exit 0
else
  warn "FAIL: #{smoke.summary(serial)}"
  exit 1
end
