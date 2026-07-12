# frozen_string_literal: true

require_relative "ssh_test_key"

module Xnix
  class SshProbe
    HOST = "127.0.0.1"
    PORT = "2222"
    USER = "root"

    def command(key_path: SshTestKey::PRIVATE_KEY_PATH)
      [
        "ssh",
        "-p", PORT,
        "-i", key_path,
        "-o", "BatchMode=yes",
        "-o", "ConnectTimeout=5",
        "-o", "IdentitiesOnly=yes",
        "-o", "StrictHostKeyChecking=no",
        "-o", "UserKnownHostsFile=/dev/null",
        "#{USER}@#{HOST}",
        "true"
      ]
    end
  end
end
