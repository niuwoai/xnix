# frozen_string_literal: true

require "pathname"

ROOT = Pathname.new(__dir__).join("..").realpath
script = ROOT.join("scripts", "known_winapp_guest_wine_smoke.rb").read

def assert(condition, message)
  raise message unless condition
end

assert(script.include?("\"windows-known-app-run\""), "known app guest smoke must call the unified known app run entrypoint")
assert(script.include?("\"--backend\", \"guest-wine\""), "known app guest smoke must select the guest-wine backend")
assert(script.include?("\"--start-qemu\""), "known app guest smoke must let the Go Runtime start the QEMU guest")
assert(script.include?("\"--port\", \"auto\""), "known app guest smoke must avoid a fixed host SSH port for Go-started QEMU")
assert(script.include?("\"--qemu-kernel\", qemu.kernel_image"), "known app guest smoke must pass the managed QEMU kernel to the Runtime")
assert(script.include?("\"--key\", Xnix::SshTestKey::PRIVATE_KEY_PATH"), "known app guest smoke must pass the loopback SSH key to the Runtime")
assert(script.include?("payload.fetch(\"backend\") == \"guest-wine\""), "known app guest smoke must validate the returned backend")
assert(script.include?("payload.fetch(\"guest_start_mode\") == \"go-qemu\""), "known app guest smoke must validate Go-owned QEMU start mode")
assert(!script.include?("Open3.popen2e(*qemu.boot_command"), "known app guest smoke must not start QEMU from Ruby")
assert(!script.include?("\"windows-known-app-dispatch-smoke\""), "known app guest smoke must not bypass the unified known app run entrypoint")
assert(!script.include?("\"--guest-boundary\""), "known app guest smoke must not rely on the older dispatch boundary flag")

puts "PASS: known Windows app guest Wine smoke script uses unified RunKnownApp guest-wine entrypoint"
