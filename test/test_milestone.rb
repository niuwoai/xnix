#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/milestone"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

assert(!Xnix::Milestone.full_build_required?("0.1.9"), "ordinary small versions must not run a full build")
assert(!Xnix::Milestone.full_build_required?("0.2.590"), "ordinary non-twentieth versions must not run a full build")
assert(Xnix::Milestone.full_build_required?("0.2.580"), "the twentieth version must run a full build")
assert(Xnix::Milestone.full_build_required?("1.4.20"), "later twentieth versions must run a full build")
assert(Xnix::Milestone.full_build_required?("0.2.580-rc1"), "twentieth-version release candidates must rerun the full build after fixes")

puts "PASS: milestone version unit tests"
