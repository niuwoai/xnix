//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

const (
	messageBoxTitle = "Xnix Windows GUI Smoke"
	messageBoxBody  = "XNIX_WINAPP_GUI_OK"
)

func main() {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(mustUTF16Ptr(messageBoxBody))),
		uintptr(unsafe.Pointer(mustUTF16Ptr(messageBoxTitle))),
		0,
	)
}

func mustUTF16Ptr(value string) *uint16 {
	pointer, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		panic(err)
	}
	return pointer
}
