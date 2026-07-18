#!/usr/bin/env ruby
# frozen_string_literal: true

# Boot smoke for a produced Xnix KDE Plasma disk image. Boots the image
# under QEMU, signs in through the serial console, and actively probes the
# graphical target and login manager declared in the manifest. Requires a
# real disk image and UEFI firmware; it never fakes a boot.
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
require_relative "../lib/xnix/image/serial_probe"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def arg_value(flag)
  index = ARGV.index(flag)
  index ? ARGV[index + 1] : nil
end

image = Xnix::Image::KdeImage.new(project_root: PROJECT_ROOT.to_s)
smoke = Xnix::Image::BootSmoke.new(image)
disk_build = Xnix::Image::DiskBuild.new(project_root: PROJECT_ROOT.to_s)

# Default the disk path to what the disk-build step produces, so the
# pipeline connects end-to-end without repeating paths.
default_output = Pathname.new(arg_value("--output") || PROJECT_ROOT.join("output")).expand_path
disk = arg_value("--disk") ||
       disk_build.output_path("qcow2", default_output.to_s)
firmware = arg_value("--firmware") || "/usr/share/edk2/ovmf/OVMF_CODE.fd"
serial_log = Pathname.new(arg_value("--serial-log") || default_output.join("serial-smoke.log")).expand_path
acceleration = File.readable?("/dev/kvm") && File.writable?("/dev/kvm") ? "kvm" : "tcg"
account = Array(disk_build.blueprint.dig("customizations", "user")).first

abort "FAIL: disk image not found: #{disk}" unless File.file?(disk)
abort "FAIL: UEFI firmware not found: #{firmware}" unless File.file?(firmware)
abort "FAIL: disk blueprint has no serial smoke account" unless account

command = smoke.boot_command(disk_path: disk, firmware_path: firmware, acceleration: acceleration)
puts "RUN: #{command.join(' ')}"
puts "LOG: #{serial_log}"

FileUtils.mkdir_p(serial_log.dirname)
serial = +""
probe = Xnix::Image::SerialProbe.new(
  username: account.fetch("name"),
  password: account.fetch("password"),
  command: smoke.probe_command
)
deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + smoke.timeout_seconds
File.open(serial_log, "w") do |log|
  Open3.popen2e(*command) do |stdin, output, wait_thread|
    begin
      while wait_thread.alive? && Process.clock_gettime(Process::CLOCK_MONOTONIC) < deadline
        next unless IO.select([output], nil, nil, 1)

        chunk = output.read_nonblock(4096)
        serial << chunk
        log.write(chunk)
        log.flush
        print chunk
        input = probe.next_input(serial)
        stdin.write(input) if input
        stdin.flush if input
        break if smoke.booted?(serial) || serial.include?(smoke.fail_marker)
      end
    rescue EOFError
      nil
    ensure
      begin
        Process.kill("TERM", wait_thread.pid) if wait_thread.alive?
      rescue Errno::ESRCH
        nil
      end
      wait_thread.value
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
