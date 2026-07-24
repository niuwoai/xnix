package winapp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	DefaultQEMUBinary      = "qemu-system-i386"
	DefaultQEMUMemory      = "1024M"
	DefaultQEMUCPUCount    = "2"
	DefaultQEMUCPUModel    = "qemu32"
	DefaultQEMUBootTimeout = 180 * time.Second
)

type QEMUStartedGuestRequest struct {
	Binary        string
	KernelImage   string
	Memory        string
	CPUCount      string
	CPUModel      string
	Host          string
	Port          string
	User          string
	KeyPath       string
	SSHPath       string
	BootTimeout   time.Duration
	SerialLogPath string
}

type QEMUStartedGuest struct {
	command       *exec.Cmd
	serialBuffer  *bytes.Buffer
	serialLogPath string
	waitDone      chan error
	stopped       bool
}

func StartQEMUStartedGuest(ctx context.Context, request QEMUStartedGuestRequest) (*QEMUStartedGuest, error) {
	qemuPath, err := resolveTool(request.Binary, DefaultQEMUBinary)
	if err != nil {
		return nil, fmt.Errorf("qemu guest runner unavailable: %w", err)
	}
	kernelImage, err := validateQEMUKernelImage(request.KernelImage)
	if err != nil {
		return nil, err
	}
	sshPath, err := resolveTool(request.SSHPath, "ssh")
	if err != nil {
		return nil, fmt.Errorf("guest ssh transport unavailable: %w", err)
	}

	timeout := request.BootTimeout
	if timeout <= 0 {
		timeout = DefaultQEMUBootTimeout
	}
	bootCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var serial bytes.Buffer
	command := exec.CommandContext(ctx, qemuPath, qemuGuestArgs(request, kernelImage)...)
	command.Stdout = &serial
	command.Stderr = &serial
	if err := command.Start(); err != nil {
		return nil, fmt.Errorf("qemu guest start failed: %w", err)
	}

	guest := &QEMUStartedGuest{
		command:       command,
		serialBuffer:  &serial,
		serialLogPath: request.SerialLogPath,
		waitDone:      make(chan error, 1),
	}

	go func() {
		guest.waitDone <- command.Wait()
	}()

	probeRequest := GuestRequest{
		Host:    request.Host,
		Port:    request.Port,
		User:    request.User,
		KeyPath: request.KeyPath,
	}
	sshBase := guestSSHBaseArgs(probeRequest)
	target := guestTarget(probeRequest)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		if err := runGuestSSH(bootCtx, sshPath, sshBase, target, "true", nil, nil); err == nil {
			return guest, nil
		}

		select {
		case err := <-guest.waitDone:
			guest.writeSerialLog()
			return nil, fmt.Errorf("qemu guest exited before ssh became ready: %w", err)
		case <-bootCtx.Done():
			guest.Stop()
			return nil, errors.New("qemu guest timed out waiting for ssh")
		case <-ticker.C:
		}
	}
}

func (guest *QEMUStartedGuest) Stop() {
	if guest == nil || guest.command == nil || guest.command.Process == nil || guest.stopped {
		return
	}
	guest.stopped = true
	_ = guest.command.Process.Signal(syscall.SIGTERM)
	select {
	case <-guest.waitDone:
	case <-time.After(5 * time.Second):
		_ = guest.command.Process.Kill()
		<-guest.waitDone
	}
	guest.writeSerialLog()
}

func (guest *QEMUStartedGuest) writeSerialLog() {
	if guest == nil || strings.TrimSpace(guest.serialLogPath) == "" || guest.serialBuffer == nil || guest.serialBuffer.Len() == 0 {
		return
	}
	path, err := filepath.Abs(guest.serialLogPath)
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(path), 0o700)
	_ = os.WriteFile(path, guest.serialBuffer.Bytes(), 0o600)
}

func validateQEMUKernelImage(path string) (string, error) {
	if strings.TrimSpace(path) == "" {
		return "", errors.New("qemu guest kernel unavailable")
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", errors.New("qemu guest kernel unavailable")
	}
	if info.IsDir() {
		return "", errors.New("qemu guest kernel unavailable")
	}
	return absolutePath, nil
}

func qemuGuestArgs(request QEMUStartedGuestRequest, kernelImage string) []string {
	memory := stringDefault(request.Memory, DefaultQEMUMemory)
	cpuCount := stringDefault(request.CPUCount, DefaultQEMUCPUCount)
	cpuModel := stringDefault(request.CPUModel, DefaultQEMUCPUModel)
	return []string{
		"-machine", "q35,accel=tcg",
		"-cpu", cpuModel,
		"-m", memory,
		"-smp", cpuCount,
		"-nographic",
		"-serial", "mon:stdio",
		"-no-reboot",
		"-kernel", kernelImage,
		"-append", "console=ttyS0,115200 panic=-1",
		"-netdev", "user,id=net0,restrict=on,hostfwd=tcp:" + stringDefault(request.Host, DefaultGuestHost) + ":" + guestPort(GuestRequest{Port: request.Port}) + "-:22",
		"-device", "e1000,netdev=net0",
	}
}

func stringDefault(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
