# frozen_string_literal: true

require "pathname"

ROOT = Pathname.new(__dir__).join("..").realpath
script = ROOT.join("scripts", "known_winapp_guest_wine_smoke.rb").read

def assert(condition, message)
  raise message unless condition
end

assert(script.include?("\"windows-known-app-run\""), "known app guest smoke must call the unified known app run entrypoint")
assert(script.include?("\"--backend\", \"guest-wine\""), "known app guest smoke must select the guest-wine backend")
assert(script.include?("\"--key\", Xnix::SshTestKey::PRIVATE_KEY_PATH"), "known app guest smoke must pass the loopback SSH key to the Runtime")
assert(script.include?("payload.fetch(\"backend\") == \"guest-wine\""), "known app guest smoke must validate the returned backend")
assert(!script.include?("\"windows-known-app-dispatch-smoke\""), "known app guest smoke must not bypass the unified known app run entrypoint")
assert(!script.include?("\"--guest-boundary\""), "known app guest smoke must not rely on the older dispatch boundary flag")

puts "PASS: known Windows app guest Wine smoke script uses unified RunKnownApp guest-wine entrypoint"
