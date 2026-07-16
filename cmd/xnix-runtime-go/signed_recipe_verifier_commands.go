package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/recipe"
)

func runSignedRecipeVerifierPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("signed-recipe-verifier-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	recipePath := flags.String("recipe", "", "path to a local recipe JSON file")
	metadataPath := flags.String("metadata", "", "path to signed recipe metadata JSON")
	publicKeyPath := flags.String("public-key", "", "path to a base64-encoded Ed25519 public key")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *recipePath == "" || *metadataPath == "" || *publicKeyPath == "" {
		return errors.New("signed-recipe-verifier-preview requires --recipe, --metadata, and --public-key")
	}
	if flags.NArg() != 0 {
		return errors.New("signed-recipe-verifier-preview does not accept positional arguments")
	}

	recipeData, err := os.ReadFile(*recipePath)
	if err != nil {
		return err
	}
	metadataData, err := os.ReadFile(*metadataPath)
	if err != nil {
		return err
	}
	var metadata recipe.SignedMetadata
	if err := json.Unmarshal(metadataData, &metadata); err != nil {
		return err
	}
	publicKeyData, err := os.ReadFile(*publicKeyPath)
	if err != nil {
		return err
	}
	publicKey, err := base64.StdEncoding.DecodeString(string(publicKeyData))
	if err != nil {
		return errors.New("signed recipe public key must be base64 encoded")
	}
	entry := recipe.Entry{ID: metadata.ApplicationID, SHA256: metadata.RecipeSHA256, SignatureStatus: recipe.SignatureSigned}
	preview := recipe.NewSignedRecipeVerificationPreview(entry, recipeData, metadata, ed25519.PublicKey(publicKey))
	return encodeIndentedJSON(stdout, preview)
}
