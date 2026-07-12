# frozen_string_literal: true

module Xnix
  class SerialLog
    BOOT_MARKERS = ["Welcome to Buildroot", "xnix login:"].freeze

    def initialize(contents)
      @contents = contents
    end

    def booted?
      BOOT_MARKERS.all? { |marker| @contents.include?(marker) }
    end

    def missing_markers
      BOOT_MARKERS.reject { |marker| @contents.include?(marker) }
    end

    def summary
      return "boot completed" if booted?

      "missing boot markers: #{missing_markers.join(", ")}"
    end
  end
end
