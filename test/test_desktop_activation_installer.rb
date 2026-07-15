#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "fileutils"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/desktop_activation_installer"
require_relative "../lib/xnix/compatibility/recipe_install_gate"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe_store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
recipe = recipe_store.find("org.xnix.sample.notepad")
registry_report = Xnix::Compatibility::RecipeRegistry.new(
  path: project_root.join("runtime/recipes/registry.json")
).verify
development_gate = Xnix::Compatibility::RecipeInstallGate.new(
  registry_report: registry_report,
  application_id: "org.xnix.sample.notepad",
  mode: "development"
)

Dir.mktmpdir("xnix-desktop-root") do |root|
  result = Xnix::Compatibility::DesktopActivationInstaller.new(
    root: root,
    recipe: recipe,
    install_gate: development_gate
  ).install
  root_path = Pathname.new(root)
  desktop_entry = root_path.join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  mimeapps_list = root_path.join("usr/share/applications/mimeapps.list")
  service_menu = root_path.join("usr/share/kio/servicemenus/xnix-open-with-compatibility.desktop")
  manifest_path = root_path.join("usr/share/xnix/compatibility/manifests/org.xnix.sample.notepad.json")

  assert(result["version"] == "0.2.209", "desktop activation installer must expose the current version")
  assert(result["application_id"] == "org.xnix.sample.notepad", "desktop activation installer must identify the application")
  assert(result["root"] == root, "desktop activation installer must report the staging root")
  assert(result["preflight"]["decision"] == "allow", "desktop activation installer must report allowed preflight")
  assert(result["preflight"]["mode"] == "development", "desktop activation installer must preserve install gate mode")
  assert(result["installed"].length == 4, "desktop activation installer must install four staged files")
  assert(result["receipt"]["kind"] == "desktop-activation-receipt", "desktop activation installer must write a rollback receipt")
  assert(result["receipt"]["sha256"].match?(/\A[0-9a-f]{64}\z/), "desktop activation receipt must include a SHA-256 digest")
  assert(result["activated_entry_points"].length == 7, "desktop activation installer must activate all manifest entry points")
  assert(result["safety"]["staging_root_required"], "desktop activation installer must require a staging root")
  assert(!result["safety"]["host_root_modified"], "desktop activation installer must not modify the host root")
  assert(!result["safety"]["backend_commands_exposed"], "desktop activation installer must not expose backend commands")
  assert(result["safety"]["desktop_entry_source"] == "ruby", "desktop activation installer must report the default desktop entry source")
  assert(result["safety"]["file_association_source"] == "ruby", "desktop activation installer must report the default file association source")
  assert(result["safety"]["recipe_install_gate_enforced"], "desktop activation installer must enforce the recipe install gate when supplied")
  assert(result["safety"]["rollback_receipt_written"], "desktop activation installer must report rollback receipt writing")

  assert(desktop_entry.file?, "desktop activation installer must install the generated desktop entry")
  assert(desktop_entry.read.include?("Exec=xnix-compat-launch --app org.xnix.sample.notepad %U"), "installed desktop entry must call the managed launcher")
  assert((desktop_entry.stat.mode & 0o777) == 0o644, "installed desktop entry must be non-executable")

  assert(mimeapps_list.file?, "desktop activation installer must install the MIME association list")
  assert(mimeapps_list.read.include?("application/x-xnix-txt=xnix-org.xnix.sample.notepad.desktop"), "installed MIME list must map text files to the generated desktop entry")
  assert(mimeapps_list.read.include?("application/x-xnix-log=xnix-org.xnix.sample.notepad.desktop;"), "installed MIME list must add log file associations")
  assert((mimeapps_list.stat.mode & 0o777) == 0o644, "installed MIME list must be non-executable")

  assert(service_menu.file?, "desktop activation installer must install the Dolphin service menu")
  assert(service_menu.read.include?("Exec=xnix-compat-open %U"), "installed Dolphin service menu must call the managed file-open command")
  assert((service_menu.stat.mode & 0o777) == 0o644, "installed service menu must be non-executable")

  assert(manifest_path.file?, "desktop activation installer must persist the desktop integration manifest")
  manifest = JSON.parse(manifest_path.read)
  assert(manifest["manifest_type"] == "desktop-integration", "installed manifest must keep its manifest type")
  assert(manifest["entry_points"] == result["activated_entry_points"], "installed manifest must match activated entry points")
  assert(!manifest_path.read.match?(/wine|prefix|\.wine|proton|virtual machine/i), "installed manifest must not expose backend implementation terms")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  result = Xnix::Compatibility::DesktopActivationInstaller.new(
    root: root,
    recipe: recipe,
    install_gate: development_gate,
    desktop_entry_renderer: Xnix::Compatibility::DesktopActivationInstaller::RuntimeGoDesktopEntryRenderer.new(
      command: project_root.join("test/fixtures/xnix-runtime-go-fake").to_s,
      registry_path: project_root.join("runtime/recipes/registry.json").to_s,
      application_id: recipe.id
    ),
    file_association_renderer: Xnix::Compatibility::DesktopActivationInstaller::RuntimeGoFileAssociationRenderer.new(
      command: project_root.join("test/fixtures/xnix-runtime-go-fake").to_s,
      registry_path: project_root.join("runtime/recipes/registry.json").to_s,
      application_id: recipe.id
    )
  ).install
  desktop_entry = Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  mimeapps_list = Pathname.new(root).join("usr/share/applications/mimeapps.list")

  assert(result["safety"]["desktop_entry_source"] == "runtime-go", "desktop activation installer must report Runtime Go desktop entry source")
  assert(result["safety"]["file_association_source"] == "runtime-go", "desktop activation installer must report Runtime Go file association source")
  assert(desktop_entry.read.include?("X-Xnix-Renderer=fake-runtime-go"), "Runtime Go renderer output must be staged as the desktop entry")
  assert(desktop_entry.read.include?("Exec=xnix-compat-launch --app org.xnix.sample.notepad %U"), "Runtime Go renderer output must keep the managed launcher")
  assert(!desktop_entry.read.match?(/wine|prefix|\.exe|proton|qemu-system|program files/i), "Runtime Go renderer output must not expose backend details")
  assert(mimeapps_list.read.include?("[Default Applications]\n"), "Runtime Go file association renderer output must include default associations")
  assert(mimeapps_list.read.include?("application/x-xnix-txt=xnix-org.xnix.sample.notepad.desktop"), "Runtime Go file association renderer output must map text files")
  assert(mimeapps_list.read.include?("application/x-xnix-log=xnix-org.xnix.sample.notepad.desktop;"), "Runtime Go file association renderer output must add log files")
  assert(!mimeapps_list.read.match?(/wine|prefix|\.exe|proton|qemu-system|program files/i), "Runtime Go file association renderer output must not expose backend details")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-install-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad",
    "--mode",
    "development"
  )
  assert(status.success?, "desktop activation installer CLI must exit successfully: #{stderr}")
  result = JSON.parse(stdout)
  assert(result["application_id"] == "org.xnix.sample.notepad", "desktop activation installer CLI must emit install result")
  assert(result["preflight"]["decision"] == "allow", "desktop activation installer CLI must enforce an allowed preflight")
  assert(Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop").file?, "desktop activation installer CLI must install the desktop entry")
  assert(Pathname.new(root).join("usr/share/applications/mimeapps.list").file?, "desktop activation installer CLI must install file associations")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-install-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad",
    "--mode",
    "development",
    "--desktop-entry-source",
    "runtime-go",
    "--file-association-source",
    "runtime-go",
    "--runtime-go-bin",
    project_root.join("test/fixtures/xnix-runtime-go-fake").to_s
  )
  assert(status.success?, "desktop activation installer CLI must use the Runtime Go renderer: #{stderr}")
  result = JSON.parse(stdout)
  desktop_entry = Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  mimeapps_list = Pathname.new(root).join("usr/share/applications/mimeapps.list")
  assert(result["safety"]["desktop_entry_source"] == "runtime-go", "desktop activation installer CLI must report Runtime Go rendering")
  assert(result["safety"]["file_association_source"] == "runtime-go", "desktop activation installer CLI must report Runtime Go file association rendering")
  assert(desktop_entry.read.include?("X-Xnix-Renderer=fake-runtime-go"), "desktop activation installer CLI must stage Runtime Go renderer output")
  assert(mimeapps_list.read.include?("application/x-xnix-txt=xnix-org.xnix.sample.notepad.desktop"), "desktop activation installer CLI must stage Runtime Go MIME association output")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  mimeapps_list = Pathname.new(root).join("usr/share/applications/mimeapps.list")
  FileUtils.mkdir_p(mimeapps_list.dirname)
  File.write(mimeapps_list, "[Default Applications]\ntext/plain=existing.desktop\n")

  _stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-install-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad",
    "--mode",
    "development"
  )

  assert(!status.success?, "desktop activation installer CLI must reject existing mimeapps overwrite")
  assert(stderr.include?("refusing to overwrite existing mimeapps list"), "desktop activation installer CLI must explain mimeapps overwrite refusal")
  assert(mimeapps_list.read.include?("existing.desktop"), "desktop activation installer must preserve existing mimeapps files")
  assert(!Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop").exist?, "mimeapps overwrite refusal must not leave partial desktop entries")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  _stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-install-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(!status.success?, "desktop activation installer CLI must block development-only registries in production mode")
  assert(stderr.include?("recipe install gate blocked activation"), "desktop activation installer CLI must explain preflight blocking")
  assert(!Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop").exist?, "blocked desktop activation must not install files")
end

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-install-desktop-integration").to_s,
  "--root",
  "/",
  "--app",
  "org.xnix.sample.notepad"
)
assert(!status.success?, "desktop activation installer CLI must reject host root installation")
assert(stderr.include?("root must not be /"), "desktop activation installer CLI must explain host root rejection")

puts "PASS: desktop activation installer unit tests"
