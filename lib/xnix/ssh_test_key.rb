# frozen_string_literal: true

module Xnix
  class SshTestKey
    DIRECTORY = "/workspace/.cache/xnix-ssh-test"
    PRIVATE_KEY_PATH = "#{DIRECTORY}/id_ed25519"
    PUBLIC_KEY_PATH = "#{PRIVATE_KEY_PATH}.pub"
    COMMENT = "xnix-qemu-test"
  end
end
