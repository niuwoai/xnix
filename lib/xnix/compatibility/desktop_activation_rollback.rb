# frozen_string_literal: true

require "digest"
require "fileutils"
require "json"
require "optparse"
require "pathname"
require_relative "desktop_activation_installer"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DesktopActivationRollback
      def initialize(root:, application_id:)
        @root = Pathname.new(root)
        @application_id = application_id
      end

      def rollback
        validate_root!
        receipt = load_receipt
        raise ArgumentError, "receipt application mismatch" unless receipt.fetch("application_id") == application_id

        removed = receipt.fetch("installed").map { |entry| remove_entry(entry) }
        receipt_entry = remove_receipt

        {
          "version" => RuntimeDaemon::VERSION,
          "application_id" => application_id,
          "root" => root.to_s,
          "removed" => removed,
          "receipt_removed" => receipt_entry,
          "safety" => {
            "staging_root_required" => true,
            "host_root_modified" => false,
            "sha256_verified_before_remove" => true
          }
        }
      rescue KeyError, JSON::ParserError => e
        raise ArgumentError, "invalid activation receipt: #{e.message}"
      end

      private

      attr_reader :root, :application_id

      def validate_root!
        raise ArgumentError, "root must not be /" if root.cleanpath.to_s == "/"
      end

      def load_receipt
        raise ArgumentError, "missing activation receipt: #{receipt_relative_path}" unless receipt_path.file?

        JSON.parse(receipt_path.read)
      end

      def remove_entry(entry)
        relative_path = entry.fetch("path")
        expected_sha256 = entry.fetch("sha256")
        destination = safe_destination(relative_path)

        return removal_result(entry, "missing") unless destination.file?

        actual_sha256 = Digest::SHA256.file(destination).hexdigest
        raise ArgumentError, "refusing to remove changed file: #{relative_path}" unless actual_sha256 == expected_sha256

        FileUtils.rm_f(destination)
        removal_result(entry, "removed")
      end

      def remove_receipt
        entry = {
          "entry_point" => "rollback",
          "kind" => "desktop-activation-receipt",
          "path" => receipt_relative_path,
          "sha256" => Digest::SHA256.file(receipt_path).hexdigest
        }
        FileUtils.rm_f(receipt_path)
        removal_result(entry, "removed")
      end

      def removal_result(entry, status)
        {
          "entry_point" => entry.fetch("entry_point"),
          "kind" => entry.fetch("kind"),
          "path" => entry.fetch("path"),
          "status" => status
        }
      end

      def receipt_relative_path
        File.join(DesktopActivationInstaller::RECEIPTS_DIR, "#{application_id}.json")
      end

      def receipt_path
        safe_destination(receipt_relative_path)
      end

      def safe_destination(relative_path)
        raise ArgumentError, "rollback path must be relative" if Pathname.new(relative_path).absolute?
        raise ArgumentError, "rollback path must not escape root" if relative_path.split(File::SEPARATOR).include?("..")

        root.join(relative_path)
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @root = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--root is required" unless @root
          raise ArgumentError, "--app is required" unless @application_id

          result = DesktopActivationRollback.new(root: @root, application_id: @application_id).rollback
          puts JSON.pretty_generate(result)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-rollback-desktop-integration: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-rollback-desktop-integration --root PATH --app APP_ID"
            options.on("--root PATH", "Rollback desktop activation files under PATH") do |value|
              @root = value
            end
            options.on("--app APP_ID", "Rollback desktop activation files for APP_ID") do |value|
              @application_id = value
            end
          end
        end
      end
    end
  end
end
