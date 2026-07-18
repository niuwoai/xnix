# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "kde_image"

module Xnix
  module Image
    # DiskBuild turns the xnix-kinoite bootc container image into a
    # bootable disk image (qcow2/raw/iso) using bootc-image-builder. It is
    # the structural link between "build the container" and "boot the disk":
    #
    #   manifest.json -> Containerfile -> (podman build) container image
    #     -> disk-config.json -> (bootc-image-builder) disk image -> boot smoke
    #
    # Like KdeImage, everything here is pure data + command rendering. The
    # actual disk build needs a privileged podman host (bootc-image-builder
    # runs privileged and reads the container store); the driver detects
    # this and fails honestly rather than faking a build.
    class DiskBuild
      CLI_COMMAND = "xnix-kde-disk"
      LOCAL_REGISTRY = "localhost"

      REQUIRED_KEYS = %w[
        schema source_image builder_image rootfs output_basename
        output_types supported_output_types blueprint
      ].freeze

      SUPPORTED_ROOTFS_TYPES = %w[btrfs ext4 xfs].freeze

      # bootc-image-builder writes each type to a known sub-path under the
      # output directory.
      OUTPUT_LAYOUT = {
        "qcow2" => "qcow2/disk.qcow2",
        "raw" => "image/disk.raw",
        "iso" => "bootiso/install.iso"
      }.freeze

      class ConfigError < StandardError; end

      def self.default_config_path(project_root)
        Pathname.new(project_root).join("image/kinoite/disk-config.json")
      end

      def initialize(project_root:, config_path: nil, kde_image: nil)
        @project_root = Pathname.new(project_root)
        @config_path = Pathname.new(config_path || self.class.default_config_path(@project_root))
        @config = load_config
        @kde_image = kde_image || KdeImage.new(project_root: @project_root.to_s)
      end

      attr_reader :config

      def version
        @kde_image.version
      end

      def source_reference
        "#{LOCAL_REGISTRY}/#{@config.fetch('source_image')}:#{version}"
      end

      def output_types
        Array(@config["output_types"])
      end

      def blueprint
        @config.fetch("blueprint")
      end

      def problems
        issues = []
        issues.concat(REQUIRED_KEYS.reject { |k| @config.key?(k) }.map { |k| "disk config missing key: #{k}" })
        issues.concat(output_type_problems)
        issues.concat(rootfs_problems)
        issues.concat(source_consistency_problems)
        issues
      end

      def valid?
        problems.empty?
      end

      # Output artifact path for a given type, relative to the output dir.
      def output_path(type, output_dir)
        layout = OUTPUT_LAYOUT.fetch(type) { raise ConfigError, "unknown output type: #{type}" }
        Pathname.new(output_dir).join(layout).to_s
      end

      # The privileged podman invocation that runs bootc-image-builder for
      # one output type. blueprint_path is a file the driver writes from the
      # config's blueprint block; output_dir receives the disk image.
      def builder_command(type:, output_dir:, blueprint_path:, storage_root: nil, runroot: nil,
                          storage_config_path: nil)
        raise ConfigError, "unsupported output type: #{type}" unless supported_type?(type)

        command = ["podman"]
        if storage_root
          command.concat(["--root", storage_root])
          command.concat(["--runroot", runroot]) if runroot
        end
        command.concat([
          "run", "--rm",
          "--privileged",
          "--security-opt", "label=type:unconfined_t"
        ])
        if storage_root
          raise ConfigError, "storage_config_path is required with storage_root" unless storage_config_path

          command.concat([
            "-e", "CONTAINERS_STORAGE_CONF=/xnix-storage.conf",
            "-v", "#{storage_root}:#{storage_root}",
            "-v", "#{storage_root}:/var/lib/containers/storage",
            "-v", "#{storage_config_path}:/xnix-storage.conf:ro"
          ])
        else
          command.concat(["-v", "/var/lib/containers/storage:/var/lib/containers/storage"])
        end
        command.concat([
          "-v", "#{output_dir}:/output",
          "-v", "#{blueprint_path}:/config.json:ro",
          @config.fetch("builder_image"),
          "--type", type,
          "--rootfs", @config.fetch("rootfs"),
          "--local",
          "--config", "/config.json",
          source_reference
        ])
        command
      end

      def to_h
        {
          "version" => version,
          "artifact_type" => "kde-plasma-disk-image",
          "source_image" => @config.fetch("source_image"),
          "source_reference" => source_reference,
          "builder_image" => @config.fetch("builder_image"),
          "rootfs" => @config.fetch("rootfs"),
          "output_types" => output_types,
          "output_basename" => @config.fetch("output_basename"),
          "privileged_build_required" => true,
          "valid" => valid?,
          "problems" => problems
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :project_root, :config_path

      def load_config
        raise ConfigError, "disk config not found: #{config_path}" unless config_path.file?

        JSON.parse(config_path.read)
      rescue JSON::ParserError => e
        raise ConfigError, "disk config is not valid JSON (#{config_path}): #{e.message}"
      end

      def supported_type?(type)
        Array(@config["supported_output_types"]).include?(type)
      end

      def output_type_problems
        types = output_types
        return ["disk config must request at least one output type"] if types.empty?

        types.reject { |t| supported_type?(t) }
             .map { |t| "requested output type '#{t}' is not in supported_output_types" }
      end

      def rootfs_problems
        rootfs = @config["rootfs"]
        return [] if SUPPORTED_ROOTFS_TYPES.include?(rootfs)

        ["rootfs '#{rootfs}' is not supported (expected one of: #{SUPPORTED_ROOTFS_TYPES.join(', ')})"]
      end

      def source_consistency_problems
        manifest_name = @kde_image.manifest.fetch("image_name")
        return [] if @config.fetch("source_image", nil) == manifest_name

        ["source_image '#{@config['source_image']}' must match image manifest name '#{manifest_name}'"]
      end
    end
  end
end

if $PROGRAM_NAME == __FILE__
  options = { project_root: Pathname.new(__dir__).join("../../..").realpath.to_s }
  subcommand = ARGV.shift || "validate"

  OptionParser.new do |parser|
    parser.banner = "Usage: #{Xnix::Image::DiskBuild::CLI_COMMAND} {validate|report} [options]"
    parser.on("--project-root PATH") { |v| options[:project_root] = v }
    parser.on("--config PATH") { |v| options[:config_path] = v }
  end.parse!(ARGV)

  disk = Xnix::Image::DiskBuild.new(
    project_root: options[:project_root],
    config_path: options[:config_path]
  )

  case subcommand
  when "validate"
    problems = disk.problems
    if problems.empty?
      puts "PASS: #{disk.source_reference} disk build config is consistent"
    else
      problems.each { |problem| warn "FAIL: #{problem}" }
      exit 1
    end
  when "report"
    puts disk.to_json
  else
    warn "Unknown subcommand: #{subcommand}"
    exit 2
  end
end
