#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "fileutils"
require "open3"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
KNOWN_APP_CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
APP_ID = ENV.fetch("XNIX_KNOWN_WINAPP_ID", "7zr")

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)
FileUtils.mkdir_p(KNOWN_APP_CACHE_ROOT)

stdout, stderr, status = run_command(
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
    "GOTMPDIR" => GO_TMP_ROOT.to_s
  },
  "go", "run", "./cmd/xnix-runtime-go", "windows-known-app-fetch",
  "--app", APP_ID,
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--timeout", ENV.fetch("XNIX_KNOWN_WINAPP_FETCH_TIMEOUT", "60s")
)

unless status.zero?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  warn "FAIL: known Windows app fetch command failed"
  exit 1
end

payload = JSON.parse(stdout)
case payload.fetch("status")
when "passed"
  puts "PASS: known Windows app fetched and verified (#{payload.fetch("app_id")} #{payload.fetch("app_version")})"
when "skipped"
  puts "SKIP: known Windows app fetch (#{payload.fetch("skip_reason")})"
else
  warn stdout
  warn "FAIL: known Windows app fetch"
  exit 1
end
