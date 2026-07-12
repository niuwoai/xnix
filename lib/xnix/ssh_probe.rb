# frozen_string_literal: true

module Xnix
  class SshProbe
    HOST = "127.0.0.1"
    PORT = "2222"
    USER = "root"

    def command
      [
        "ssh",
        "-p", PORT,
        "-o", "BatchMode=yes",
        "-o", "ConnectTimeout=5",
        "-o", "StrictHostKeyChecking=no",
        "-o", "UserKnownHostsFile=/dev/null",
        "#{USER}@#{HOST}",
        "true"
      ]
    end
  end
end
