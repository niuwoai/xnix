#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "tmpdir"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/container_gui_evidence_packet.rb")

Dir.mktmpdir("xnix-container-gui-evidence-packet-test") do |dir|
  temp_root = Pathname.new(dir)
  fake_go = temp_root.join("go")
  fake_go.write(<<~'RUBY')
    #!/usr/bin/env ruby
    # frozen_string_literal: true

    require "fileutils"
    require "json"

    args = ARGV.dup
    File.open(ENV.fetch("XNIX_FAKE_GO_LOG"), "a") { |file| file.puts(args.join("\u0001")) } if ENV["XNIX_FAKE_GO_LOG"]

    if args[0] == "run" && args.include?("windows-app-container-x-gui-smoke")
      recipe_backed = args.include?("--recipe-app")
      recipe_app = recipe_backed ? args[args.index("--recipe-app") + 1] : ""
      payload = {
        "schema_version" => "xnix.runtime.windows_app_container_x_gui_smoke.v1",
        "request_type" => "windows-app-container-x-gui-smoke",
        "status" => "passed",
        "application_id" => recipe_app,
        "display_name" => recipe_backed ? "Sample Notepad" : "",
        "app_version" => recipe_backed ? "test-version" : "",
        "recipe_backed" => recipe_backed,
        "application_name" => args[args.index("--app") + 1],
        "window_match" => args[args.index("--window-match") + 1],
        "container_image" => args[args.index("--image") + 1],
        "container_platform" => args[args.index("--platform") + 1],
        "container_state_mode" => "tmpfs",
        "pull_policy" => "never",
        "network_mode" => "none",
        "desktop_display" => "Xvfb",
        "x_server_started" => true,
        "wine_bootstrap_attempted" => true,
        "runner_available" => true,
        "image_available" => true,
        "x_window_observed" => true,
        "window_evidence_summary" => "0x800001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\")",
        "exit_code" => 0,
        "duration_millis" => 3,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false,
        "host_mount_count" => 0
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("gui-smoke-evidence-preview")
      report_path = args[args.index("--gui-smoke-report") + 1]
      report = JSON.parse(File.read(report_path))
      app_id = args[args.index("--app-id") + 1]
      display_name = args[args.index("--display-name") + 1]
      app_version = args[args.index("--app-version") + 1]
      payload = {
        "schema_version" => "xnix.runtime.gui_smoke_evidence_preview.v1",
        "request_type" => "gui-smoke-evidence-preview",
        "source" => "winapp-smoke-container-x-gui+runtime-evidence-consumer",
        "status" => "passed",
        "report_consumed" => true,
        "report_path_exposed" => false,
        "app_id" => app_id,
        "display_name" => display_name,
        "app_version" => app_version,
        "gui_app_name" => report.fetch("container_gui_app"),
        "wineboot_invoked" => true,
        "x_window_observed" => true,
        "compatibility_center_projection_ready" => true,
        "kde_center_projection_ready" => true,
        "known_app_smoke_evidence" => {
          "app_id" => app_id,
          "display_name" => display_name,
          "app_version" => app_version,
          "smoke_status" => "passed",
          "evidence_source" => "winapp-smoke-container-x-gui",
          "compatibility_state" => "real-gui-container-wine-verified",
          "desktop_launch_enabled" => false,
          "backend_launch_enabled" => false,
          "backend_details_exposed" => false,
          "host_root_modified" => false
        },
        "backend_details_exposed" => false,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false
      }
      if args.include?("--output")
        output_path = args[args.index("--output") + 1]
        FileUtils.mkdir_p(File.dirname(output_path))
        File.write(output_path, JSON.pretty_generate(payload) + "\n")
      end
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("kde-center-page-preview")
      evidence_path = args[args.index("--known-app-evidence-file") + 1]
      evidence = JSON.parse(File.read(evidence_path))
      card = evidence.fetch("known_app_smoke_evidence")
      payload = {
        "schema_version" => "xnix.runtime.kde_center_page.v1",
        "request_type" => "kde-center-page-preview",
        "status" => "passed",
        "app_id" => args[args.index("--app") + 1],
        "known_app_gui_evidence_count" => 1,
        "known_app_gui_evidence_cards" => [card.merge("center_card_state" => "real-gui-container-wine-verified")],
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    warn "unexpected fake go invocation: #{args.join(" ")}"
    exit 2
  RUBY
  fake_go.chmod(0o700)

  report_output = temp_root.join("outputs/report.json")
  evidence_output = temp_root.join("outputs/evidence.json")
  kde_page_output = temp_root.join("outputs/kde-page.json")
  fake_go_log = temp_root.join("fake-go.log")
  env = {
    "PATH" => "#{temp_root}:#{ENV.fetch("PATH")}",
    "XNIX_FAKE_GO_LOG" => fake_go_log.to_s
  }

  stdout, stderr, status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--report-output", report_output.to_s,
    "--evidence-output", evidence_output.to_s,
    "--kde-page-output", kde_page_output.to_s,
    "--registry", temp_root.join("registry.json").to_s,
    "--recipe-app", "org.xnix.sample.notepad",
    "--platform", "linux/arm64",
    "--image", "local/wine-x-gui:test",
    "--gui-app", "notepad.exe",
    "--window-match", "notepad.exe",
    "--app-id", "org.xnix.sample.notepad",
    "--display-name", "Notepad",
    "--app-version", "test-version"
  )
  assert(status.success?, "container GUI evidence packet must pass: #{stderr}")

  packet = JSON.parse(stdout)
  assert(packet.fetch("schema_version") == "xnix.scripts.container_gui_evidence_packet.v1", "packet must expose schema")
  assert(packet.fetch("request_type") == "container-gui-evidence-packet", "packet must expose request type")
  assert(packet.fetch("status") == "passed", "packet must pass")
  assert(packet.fetch("report_output_written"), "packet must write the smoke report")
  assert(packet.fetch("evidence_output_written"), "packet must write projected evidence")
  assert(packet.fetch("kde_page_output_written"), "packet must write the KDE page")
  assert(packet.fetch("x_window_observed"), "packet must carry observed X window evidence")
  assert(packet.fetch("recipe_app") == "org.xnix.sample.notepad", "packet must preserve recipe app id")
  assert(packet.fetch("container_recipe_backed"), "packet must carry recipe-backed container evidence")
  assert(packet.fetch("container_application_id") == "org.xnix.sample.notepad", "packet must preserve container app id")
  assert(packet.fetch("evidence_source") == "winapp-smoke-container-x-gui", "packet must preserve container X GUI source")
  assert(packet.fetch("compatibility_state") == "real-gui-container-wine-verified", "packet must preserve compatibility state")
  assert(packet.fetch("known_app_gui_evidence_count") == 1, "packet must expose one GUI evidence card")
  assert(!packet.fetch("desktop_launch_enabled"), "packet must not enable desktop launch")
  assert(!packet.fetch("backend_launch_enabled"), "packet must not enable backend launch")
  assert(!packet.fetch("backend_details_exposed"), "packet must not expose backend details")
  assert(!packet.fetch("host_root_modified"), "packet must not mutate the host root")
  assert(!packet.fetch("privileged_container_required"), "packet must not require privileged containers")
  assert(!packet.fetch("host_networking_required"), "packet must not require host networking")
  assert(!packet.fetch("docker_socket_mounted"), "packet must not mount the Docker socket")
  assert(!packet.fetch("broad_host_mount_required"), "packet must not require broad host mounts")
  assert(packet.fetch("container_network_mode") == "none", "packet must keep container networking disabled")
  assert(packet.fetch("container_host_mount_count").zero?, "packet must not mount host paths")
  assert(report_output.file?, "smoke report file must exist")
  assert(evidence_output.file?, "evidence file must exist")
  assert(kde_page_output.file?, "KDE page file must exist")
  assert(!stdout.include?(report_output.to_s), "packet stdout must not expose report output path")
  assert(!stdout.include?(evidence_output.to_s), "packet stdout must not expose evidence output path")
  assert(!stdout.include?(kde_page_output.to_s), "packet stdout must not expose KDE page output path")

  report = JSON.parse(report_output.read)
  assert(report.fetch("backend") == "container-x-gui", "persisted report must be container X GUI evidence")
  assert(report.fetch("report_output_written"), "persisted report must record output writing")
  evidence = JSON.parse(evidence_output.read)
  assert(evidence.fetch("report_consumed"), "persisted evidence must consume the report")
  kde_page = JSON.parse(kde_page_output.read)
  assert(kde_page.fetch("known_app_gui_evidence_cards").first.fetch("app_id") == "org.xnix.sample.notepad", "KDE page must bind the app id")

  invocations = fake_go_log.read.lines.map { |line| line.split("\u0001").map(&:chomp) }
  assert(invocations.any? { |argv| argv.include?("windows-app-container-x-gui-smoke") && argv.include?("--recipe-app") }, "packet must run the recipe-backed container X GUI smoke")
  assert(invocations.any? { |argv| argv.include?("gui-smoke-evidence-preview") && argv.include?("--output") }, "packet must persist the evidence projection")
  assert(invocations.any? { |argv| argv.include?("kde-center-page-preview") && argv.include?("--known-app-evidence-file") }, "packet must render the KDE page from evidence")
end
