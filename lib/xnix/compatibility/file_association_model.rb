# frozen_string_literal: true

require "json"
require "optparse"
require_relative "desktop_entry"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class FileAssociationModel
      MIMEAPPS_RELATIVE_PATH = "usr/share/applications/mimeapps.list"
      FILE_OPEN_COMMAND = "xnix-compat-open"

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "association_type" => "desktop-file-association",
          "desktop" => "KDE Plasma",
          "application" => application,
          "mimeapps" => {
            "path" => MIMEAPPS_RELATIVE_PATH,
            "contents" => render_mimeapps
          },
          "associations" => associations,
          "safety" => safety
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      def render_mimeapps
        default_lines = associations.map do |association|
          "#{association.fetch("mime_type")}=#{desktop_file}"
        end
        added_lines = associations.map do |association|
          "#{association.fetch("mime_type")}=#{desktop_file};"
        end

        ([
          "[Default Applications]",
          *default_lines,
          "",
          "[Added Associations]",
          *added_lines
        ]).join("\n") + "\n"
      end

      private

      attr_reader :recipe

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "desktop_file" => desktop_file,
          "supported_extensions" => recipe.supported_extensions,
          "mime_types" => recipe.mime_types
        }
      end

      def associations
        recipe.mime_types.map do |mime_type|
          {
            "mime_type" => mime_type,
            "desktop_file" => desktop_file,
            "default_application" => true,
            "file_open" => {
              "argv" => [FILE_OPEN_COMMAND, "%U"],
              "portal_required" => true
            }
          }
        end
      end

      def safety
        {
          "standard_mimeapps_list" => true,
          "staged_root_only" => true,
          "overwrite_existing_mimeapps" => false,
          "portal_required_for_file_open" => true,
          "backend_details_exposed" => false
        }
      end

      def desktop_file
        @desktop_file ||= DesktopEntry.new(recipe).file_name
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts FileAssociationModel.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-file-association-model: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-file-association-model --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build file association model for APP_ID") do |value|
              @application_id = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
          end
        end
      end
    end
  end
end
