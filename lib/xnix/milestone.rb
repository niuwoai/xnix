# frozen_string_literal: true

module Xnix
  class Milestone
    INTERVAL = 10

    def self.full_build_required?(version)
      match = version.match(/\A\d+\.\d+\.(\d+)(?:-rc\d+)?\z/)
      return false if match.nil?

      match[1].to_i.positive? && (match[1].to_i % INTERVAL).zero?
    end
  end
end
