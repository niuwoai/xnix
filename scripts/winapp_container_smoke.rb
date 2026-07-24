#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")
PLATFORM = ENV.fetch("XNIX_WINE_PLATFORM", "linux/amd64")

exec(
  "ruby",
  PROJECT_ROOT.join("scripts", "winapp_smoke.rb").to_s,
  "--backend", "container",
  "--format", "text",
  "--image", IMAGE,
  "--platform", PLATFORM,
  "--timeout", "420s",
  "--bootstrap-timeout", "300s"
)
