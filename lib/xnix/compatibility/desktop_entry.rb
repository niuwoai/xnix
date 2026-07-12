# frozen_string_literal: true

require_relative "application_recipe"

module Xnix
  module Compatibility
    class DesktopEntry
      LAUNCHER = "xnix-compat-launch"

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
          "Exec=#{LAUNCHER} --app #{@recipe.id} %U",
          "Icon=#{@recipe.icon}",
          "Categories=Utility;",
          "StartupNotify=true",
          "X-Xnix-ApplicationId=#{@recipe.id}"
        ]
        lines << "MimeType=#{@recipe.mime_types.join(";")};" unless @recipe.mime_types.empty?
        "#{lines.join("\n")}\n"
      end
    end
  end
end
