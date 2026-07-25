# frozen_string_literal: true

module Xnix
  module Compatibility
    class ApplicationRecipe
      MODES = %w[automatic wine vm].freeze
      ID_PATTERN = /\A[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+\z/
      EXTENSION_PATTERN = /\A\.[a-z0-9]{1,16}\z/i

      attr_reader :id, :name, :icon, :mode, :supported_extensions, :container_gui_smoke

      def self.from_hash(data)
        raise ArgumentError, "recipe must be an object" unless data.is_a?(Hash)

        new(
          id: data.fetch("id"),
          name: data.fetch("name"),
          icon: data.fetch("icon"),
          mode: data.fetch("mode"),
          supported_extensions: data.fetch("supported_extensions", []),
          container_gui_smoke: data.fetch("container_gui_smoke", nil)
        )
      end

      def initialize(id:, name:, icon:, mode:, supported_extensions: [], container_gui_smoke: nil)
        @id = id
        @name = name
        @icon = icon
        @mode = mode
        @supported_extensions = supported_extensions
        @container_gui_smoke = container_gui_smoke
        validate!
      end

      def container_gui_smoke?
        !container_gui_smoke.nil?
      end

      def mime_types
        supported_extensions.map do |extension|
          "application/x-xnix-#{extension.delete_prefix(".").downcase}"
        end
      end

      private

      def validate!
        validate_identifier!
        validate_display_value!(name, "name")
        validate_display_value!(icon, "icon")
        raise ArgumentError, "mode must be one of: #{MODES.join(", ")}" unless MODES.include?(mode)
        raise ArgumentError, "supported_extensions must be an array" unless supported_extensions.is_a?(Array)

        supported_extensions.each do |extension|
          unless extension.is_a?(String) && extension.match?(EXTENSION_PATTERN)
            raise ArgumentError, "invalid file extension: #{extension.inspect}"
          end
        end
        validate_container_gui_smoke!
      end

      def validate_identifier!
        return if id.is_a?(String) && id.match?(ID_PATTERN)

        raise ArgumentError, "id must be a reverse-DNS identifier"
      end

      def validate_display_value!(value, label)
        return if value.is_a?(String) && !value.empty? && !value.match?(/[\r\n]/)

        raise ArgumentError, "#{label} must be a non-empty single-line string"
      end

      def validate_container_gui_smoke!
        return if container_gui_smoke.nil?

        raise ArgumentError, "container_gui_smoke must be an object" unless container_gui_smoke.is_a?(Hash)

        app = container_gui_smoke.fetch("app", nil)
        window_match = container_gui_smoke.fetch("window_match", nil)
        validate_display_value!(app, "container_gui_smoke app")
        validate_display_value!(window_match, "container_gui_smoke window_match")
        raise ArgumentError, "container_gui_smoke app must be a container application name, not a path" if app.match?(/[\\\/]/)
        raise ArgumentError, "container_gui_smoke app must name a Windows executable" unless app.downcase.end_with?(".exe")
      end
    end
  end
end
