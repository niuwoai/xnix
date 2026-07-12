#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/milestone"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

assert(!Xnix::Milestone.full_build_required?("0.1.9"), "ordinary small versions must not run a full build")
assert(Xnix::Milestone.full_build_required?("0.1.10"), "the tenth version must run a full build")
assert(Xnix::Milestone.full_build_required?("1.4.20"), "later tenth versions must run a full build")
assert(Xnix::Milestone.full_build_required?("0.1.10-rc1"), "tenth-version release candidates must rerun the full build after fixes")

puts "PASS: milestone version unit tests"
