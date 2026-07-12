# frozen_string_literal: true

require "json"
require "pathname"
require_relative "application_recipe"

module Xnix
  module Compatibility
    class RecipeStore
      RECIPE_EXTENSION = ".json"

      attr_reader :path

      def initialize(path:)
        @path = Pathname.new(path)
      end

      def all
        return [] unless path.directory?

        path.children
            .select { |entry| entry.file? && entry.extname == RECIPE_EXTENSION }
            .sort_by(&:basename)
            .map { |entry| load_recipe(entry) }
      end

      def find(application_id)
        all.find { |recipe| recipe.id == application_id }
      end

      private

      def load_recipe(entry)
        data = JSON.parse(entry.read)
        ApplicationRecipe.from_hash(data)
      rescue JSON::ParserError => e
        raise ArgumentError, "invalid recipe JSON in #{entry.basename}: #{e.message}"
      rescue KeyError, ArgumentError => e
        raise ArgumentError, "invalid recipe in #{entry.basename}: #{e.message}"
      end
    end
  end
end
