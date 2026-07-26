//go:build windows

package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

const (
	messageBoxBaseTitle = "Xnix Windows GUI Smoke"
	messageBoxBody      = "XNIX_WINAPP_GUI_OK"
	maxTitleMarkerBytes = 96
)

func main() {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(mustUTF16Ptr(messageBoxBody))),
		uintptr(unsafe.Pointer(mustUTF16Ptr(messageBoxTitle()))),
		0,
	)
}

func messageBoxTitle() string {
	filePath := openedFileArgument()
	if filePath == "" {
		return messageBoxBaseTitle
	}
	marker := fileContentMarker(filePath)
	if marker == "" {
		marker = filepath.Base(filePath)
	}
	return messageBoxBaseTitle + " - " + marker
}

func openedFileArgument() string {
	if len(os.Args) < 2 {
		return ""
	}
	return strings.TrimSpace(os.Args[1])
}

func fileContentMarker(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	marker := strings.TrimSpace(strings.SplitN(string(content), "\n", 2)[0])
	if len(marker) > maxTitleMarkerBytes {
		marker = marker[:maxTitleMarkerBytes]
	}
	return marker
}

func mustUTF16Ptr(value string) *uint16 {
	pointer, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		panic(err)
	}
	return pointer
}
