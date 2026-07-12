# frozen_string_literal: true

require_relative "application_recipe"
require_relative "recipe_registry"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class RegistryBackedRecipeStore
      REGISTRY_FILE = "registry.json"

      attr_reader :path

      def self.for_path(path)
        recipe_dir = Pathname.new(path)
        registry_path = recipe_dir.join(REGISTRY_FILE)
        return new(path: recipe_dir, registry_path: registry_path) if registry_path.file?

        RecipeStore.new(path: recipe_dir)
      end

      def initialize(path:, registry_path:)
        @path = Pathname.new(path)
        @registry_path = Pathname.new(registry_path)
      end

      def all
        registry_report.fetch("recipes").map do |entry|
          load_recipe(path.join(entry.fetch("path")))
        end
      end

      def find(application_id)
        all.find { |recipe| recipe.id == application_id }
      end

      def registry_report
        @registry_report ||= RecipeRegistry.new(path: @registry_path).verify!
      end

      private

      def load_recipe(recipe_path)
        ApplicationRecipe.from_hash(JSON.parse(recipe_path.read))
      rescue JSON::ParserError => e
        raise ArgumentError, "invalid registry recipe JSON in #{recipe_path.basename}: #{e.message}"
      rescue KeyError, ArgumentError => e
        raise ArgumentError, "invalid registry recipe in #{recipe_path.basename}: #{e.message}"
      end
    end
  end
end
