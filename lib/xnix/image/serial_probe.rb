# frozen_string_literal: true

module Xnix
  module Image
    # Drives the development image's serial login without logging credentials.
    class SerialProbe
      def initialize(username:, password:, command:)
        @username = username
        @password = password
        @command = command
        @state = :login
      end

      def next_input(serial_contents)
        case @state
        when :login
          advance(:password, "#{@username}\n") if serial_contents.include?("login:")
        when :password
          advance(:shell, "#{@password}\n") if serial_contents.include?("Password:")
        when :shell
          advance(:result, "#{@command}\n") if serial_contents.include?("Welcome to Fedora")
        end
      end

      private

      def advance(next_state, input)
        @state = next_state
        input
      end
    end
  end
end
