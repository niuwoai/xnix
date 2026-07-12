#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "fileutils"
require "open-uri"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
LOCK_PATH = PROJECT_ROOT.join("buildroot", "sources.lock")
CACHE_ROOT = PROJECT_ROOT.join(".cache", "buildroot")

def read_lock
  LOCK_PATH.each_line.each_with_object({}) do |line, lock|
    key, value = line.strip.split("=", 2)
    lock[key] = value unless key.nil? || value.nil? || key.empty? || value.empty?
  end
end

def validate_lock(lock)
  required_keys = %w[buildroot.version buildroot.url buildroot.sha256]
  missing_keys = required_keys.reject { |key| lock.key?(key) }
  abort "Missing source-lock keys: #{missing_keys.join(", ")}" unless missing_keys.empty?
  abort "Buildroot URL must use HTTPS" unless lock.fetch("buildroot.url").start_with?("https://")
  abort "Buildroot SHA-256 is invalid" unless lock.fetch("buildroot.sha256").match?(/\A[0-9a-f]{64}\z/)
end

def running_in_container?
  File.exist?("/.dockerenv") || File.read("/proc/1/cgroup").include?("docker")
rescue Errno::ENOENT
  false
end

def download_archive(url, archive_path)
  URI.open(url) do |source|
    File.open(archive_path, "wb") { |destination| IO.copy_stream(source, destination) }
  end
end

def verify_archive(archive_path, expected_sha256)
  actual_sha256 = Digest::SHA256.file(archive_path).hexdigest
  return if actual_sha256 == expected_sha256

  FileUtils.rm_f(archive_path)
  abort "Buildroot SHA-256 mismatch: expected #{expected_sha256}, got #{actual_sha256}"
end

def extract_archive(archive_path, destination)
  success = system("tar", "--extract", "--file", archive_path.to_s, "--directory", destination.to_s)
  abort "Unable to extract Buildroot archive" unless success
end

lock = read_lock
validate_lock(lock)

if ARGV == ["--verify-lock"]
  puts "PASS: Buildroot #{lock.fetch("buildroot.version")} source lock is valid"
  exit 0
end

abort "Usage: ruby scripts/fetch_buildroot.rb [--verify-lock]" unless ARGV.empty?
abort "Buildroot retrieval must run inside the restricted build container" unless running_in_container?

FileUtils.mkdir_p(CACHE_ROOT)
archive_path = CACHE_ROOT.join("buildroot-#{lock.fetch("buildroot.version")}.tar.xz")
source_directory = CACHE_ROOT.join("buildroot-#{lock.fetch("buildroot.version")}")

unless archive_path.file?
  download_archive(lock.fetch("buildroot.url"), archive_path)
end

verify_archive(archive_path, lock.fetch("buildroot.sha256"))
extract_archive(archive_path, CACHE_ROOT) unless source_directory.directory?
puts source_directory
