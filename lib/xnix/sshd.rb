# frozen_string_literal: true

module Xnix
  class Sshd
    STARTED_MARKER = "Starting sshd: OK"

    def initialize(serial_output)
      @serial_output = serial_output
    end

    def started?
      @serial_output.include?(STARTED_MARKER)
    end
  end
end
