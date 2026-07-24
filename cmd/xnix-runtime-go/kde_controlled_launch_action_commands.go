package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEControlledLaunchActionPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KDEControlledLaunchActionRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root containing the Runtime-status launch evidence handoff")
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	kdeGUICardFile := flags.String("kde-gui-card-file", "", "safe KDE GUI evidence card JSON file")
	kdeCenterPageFile := flags.String("kde-center-page-file", "", "safe KDE center page JSON file")
	appID := flags.String("app", "", "application id to select from the KDE center page")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("kde-controlled-launch-action-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("kde-controlled-launch-action-preview requires --state-root")
	}
	if strings.TrimSpace(*kdeGUICardFile) != "" && strings.TrimSpace(*kdeCenterPageFile) != "" {
		return errors.New("kde-controlled-launch-action-preview accepts only one of --kde-gui-card-file or --kde-center-page-file")
	}
	kdeGUICard, err := readKDEControlledLaunchActionGUICard(*kdeGUICardFile)
	if err != nil {
		return err
	}
	kdeCenterPage, err := readKDEControlledLaunchActionCenterPage(*kdeCenterPageFile)
	if err != nil {
		return err
	}
	preview, err := appidentity.PreviewKDEControlledLaunchAction(appidentity.KDEControlledLaunchActionRequest{
		StateRoot:            *stateRoot,
		EvidenceID:           *evidenceID,
		EvidenceRelativePath: *evidenceRelativePath,
		KDECenterGUICard:     kdeGUICard,
		KDECenterPage:        kdeCenterPage,
		KDECenterPageAppID:   *appID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func readKDEControlledLaunchActionGUICard(path string) (*appidentity.KDECenterPageKnownAppMatrixCard, error) {
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return nil, nil
	}
	payload, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}
	var card appidentity.KDECenterPageKnownAppMatrixCard
	if err := json.Unmarshal(payload, &card); err != nil {
		return nil, err
	}
	return &card, nil
}

func readKDEControlledLaunchActionCenterPage(path string) (*appidentity.KDECenterPagePreview, error) {
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return nil, nil
	}
	payload, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}
	var page appidentity.KDECenterPagePreview
	if err := json.Unmarshal(payload, &page); err != nil {
		return nil, err
	}
	return &page, nil
}
