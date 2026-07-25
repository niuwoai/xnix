# frozen_string_literal: true

require_relative "application_recipe"

module Xnix
  module Compatibility
    class DesktopEntry
      LAUNCHER = "xnix-compat-launch"
      PACKAGED_RECIPE_REGISTRY = "/usr/share/xnix/compatibility/recipes/registry.json"

      def initialize(recipe)
        @recipe = recipe
      end

      def file_name
        "xnix-#{@recipe.id}.desktop"
      end

      def render
        lines = [
          "[Desktop Entry]",
          "Type=Application",
          "Version=1.0",
          "Name=#{@recipe.name}",
          "Comment=Run with Xnix Compatibility Runtime",
          "Exec=#{exec_command}",
          "Icon=#{@recipe.icon}",
          "Categories=Utility;",
          "StartupNotify=true",
          "X-Xnix-ApplicationId=#{@recipe.id}"
        ]
        lines << "MimeType=#{@recipe.mime_types.join(";")};" unless @recipe.mime_types.empty?
        "#{lines.join("\n")}\n"
      end

      def exec_command
        command = "#{LAUNCHER} --app #{@recipe.id}"
        command = "#{command} --registry #{PACKAGED_RECIPE_REGISTRY}" if @recipe.container_gui_smoke?
        "#{command} %U"
      end
    end
  end
end
