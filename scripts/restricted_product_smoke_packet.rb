#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

options = {
  format: "json",
  repo_root: PROJECT_ROOT.to_s,
  manifest: "image/kinoite/manifest.json"
}

OptionParser.new do |parser|
  parser.on("--format FORMAT", %w[json markdown]) { |value| options[:format] = value }
  parser.on("--repo-root PATH") { |value| options[:repo_root] = value }
  parser.on("--manifest PATH") { |value| options[:manifest] = value }
end.parse!

command = [
  "go", "run", "./cmd/xnix-runtime-go", "restricted-product-smoke-packet-preview",
  "--repo-root", options.fetch(:repo_root), "--manifest", options.fetch(:manifest)
]
environment = {"GOCACHE" => ENV.fetch("GOCACHE", "/tmp/xnix-go-cache")}
stdout, stderr, status = Open3.capture3(environment, *command, chdir: PROJECT_ROOT.to_s)
abort "Restricted product smoke packet failed: #{stderr.strip}" unless status.success?

packet = JSON.parse(stdout)

def render_markdown(packet)
  lines = [
    "# Restricted Product Smoke Packet",
    "",
    "- Version: #{File.read(File.join(PROJECT_ROOT, "VERSION")).strip}",
    "- Schema: #{packet.fetch("schema_version")}",
    "- Packet prepared: #{packet.fetch("packet_prepared")}",
    "- Ready for authorized smoke: #{packet.fetch("ready_for_authorized_smoke")}",
    "- Human authorization required: #{packet.fetch("human_authorization_required")}",
    "- Docker executed: #{packet.fetch("docker_executed")}",
    "- QEMU executed: #{packet.fetch("qemu_executed")}",
    "- Serial log persisted: #{packet.fetch("serial_log_persisted")}",
    "- Loopback-only networking: #{packet.fetch("loopback_only_networking")}",
    "- Docker socket mounted: #{packet.fetch("docker_socket_mounted")}",
    "- Host networking enabled: #{packet.fetch("host_network_enabled")}",
    "- Broad host mount enabled: #{packet.fetch("broad_host_mount_enabled")}",
    "- Host root modified: #{packet.fetch("host_root_modified")}",
    "",
    "## Evidence",
    ""
  ]
  packet.fetch("evidence").each do |evidence|
    lines << "- #{evidence.fetch("id")}: #{evidence.fetch("status")} (#{evidence.fetch("evidence_count")} source checks)"
  end
  lines.concat(["", packet.fetch("desktop_safe_summary"), ""])
  lines.join("\n")
end

if options.fetch(:format) == "markdown"
  puts render_markdown(packet)
else
  puts JSON.pretty_generate(packet)
end
