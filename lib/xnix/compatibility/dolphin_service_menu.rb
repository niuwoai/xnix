# frozen_string_literal: true

module Xnix
  module Compatibility
    class DolphinServiceMenu
      COMMAND = "xnix-compat-open"
      FILE_NAME = "xnix-open-with-compatibility.desktop"

      def file_name
        FILE_NAME
      end

      def render
        [
          "[Desktop Entry]",
          "Type=Service",
          "MimeType=application/octet-stream;text/plain;",
          "Actions=openWithXnixCompatibility;",
          "X-KDE-ServiceTypes=KonqPopupMenu/Plugin",
          "X-KDE-Priority=TopLevel",
          "",
          "[Desktop Action openWithXnixCompatibility]",
          "Name=Open with Xnix Compatibility",
          "Icon=preferences-desktop",
          "Exec=#{COMMAND} %U"
        ].join("\n") + "\n"
      end
    end
  end
end
