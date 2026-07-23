#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
WORK_ROOT = PROJECT_ROOT.join(".local", "xnix", "winapp-container-smoke")
GO_CACHE_ROOT = PROJECT_ROOT.join(".gocache")
APP_ROOT = WORK_ROOT.join("app")
EXE_PATH = APP_ROOT.join("hello.exe")
STATE_ROOT = WORK_ROOT.join("state")
IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")
MARKER = "XNIX_WINAPP_SMOKE_OK"

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

FileUtils.mkdir_p(APP_ROOT)
FileUtils.mkdir_p(STATE_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))

build_stdout, build_stderr, build_status = run_command(
  {
    "GOOS" => "windows",
    "GOARCH" => "amd64",
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
  },
  "go", "build", "-o", EXE_PATH.to_s, "./test/fixtures/winapp/hello"
)

unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: Windows app fixture build failed"
  exit 1
end

smoke_stdout, smoke_stderr, smoke_status = run_command(
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
  },
  "go", "run", "./cmd/xnix-runtime-go", "windows-app-container-run-smoke",
  "--exe", EXE_PATH.to_s,
  "--state-root", STATE_ROOT.to_s,
  "--image", IMAGE,
  "--timeout", "60s"
)

unless smoke_status.zero?
  warn smoke_stdout unless smoke_stdout.empty?
  warn smoke_stderr unless smoke_stderr.empty?
  warn "FAIL: Windows app container smoke command failed"
  exit 1
end

payload = JSON.parse(smoke_stdout)
case payload.fetch("status")
when "passed"
  if payload["marker_observed"] && payload["stdout"].include?(MARKER)
    puts "PASS: containerized real Windows app smoke"
    exit 0
  end
  warn "FAIL: containerized Windows app smoke marker missing"
  exit 1
when "skipped"
  puts "SKIP: containerized real Windows app smoke (#{payload.fetch("skip_reason")})"
  exit 0
else
  warn smoke_stdout
  warn smoke_stderr unless smoke_stderr.empty?
  warn "FAIL: containerized real Windows app smoke"
  exit 1
end
