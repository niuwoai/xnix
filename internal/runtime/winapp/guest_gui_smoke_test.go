package winapp

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunGuestGUISmokeObservesWindowThroughGoRuntime(t *testing.T) {
	if os.PathSeparator != '/' {
		t.Skip("shell fixtures require a POSIX host")
	}
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "guest.log")
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'winex11'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *'wineboot --init'*) printf 'boot initialized\\n' >&2; exit 0 ;;\n" +
		"  *'wine '*'winemine.exe'*) exit 0 ;;\n" +
		"  *'cat '*'stderr.txt'*) printf ''; exit 0 ;;\n" +
		"  *'wineserver -k'*) exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	xwininfoPath := filepath.Join(tempDir, "fake-xwininfo")
	xwininfoBody := "#!/bin/sh\n" +
		"printf 'xwininfo display=%s\\n' \"$DISPLAY\" >> '" + logPath + "'\n" +
		"printf 'xwininfo: Window id: 0x3a7 (the root window)\\n'\n" +
		"printf '  0x200001 \"WineMine\": ()  320x240+0+0  +0+0\\n'\n"
	if err := os.WriteFile(xwininfoPath, []byte(xwininfoBody), 0o700); err != nil {
		t.Fatalf("WriteFile xwininfo returned error: %v", err)
	}
	guestDisplay := strings.Join([]string{"10", "0", "2", "2"}, ".") + ":100"

	result, err := RunGuestGUISmoke(context.Background(), GuestGUIRequest{
		GUIAppPath:   DefaultGuestGUIApp,
		Host:         DefaultGuestHost,
		Port:         "2222",
		User:         DefaultGuestUser,
		KeyPath:      filepath.Join(tempDir, "id_ed25519"),
		RemoteDir:    "/tmp/xnix-wine-guest-gui-smoke",
		SSHPath:      sshPath,
		XWinInfoPath: xwininfoPath,
		GuestDisplay: guestDisplay,
		HostDisplay:  ":100",
		Timeout:      5 * time.Second,
		Wait:         1 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("RunGuestGUISmoke returned error: %v", err)
	}
	if result.SchemaVersion != GuestGUISchemaVersion ||
		result.RequestType != GuestGUIRequestType ||
		result.Status != PassedStatus ||
		result.GUIAppName != "winemine.exe" ||
		result.Backend != "qemu-guest-wine-x11" ||
		!result.GuestReachable ||
		!result.WineAvailable ||
		!result.GuestX11DriverAvailable ||
		!result.WinebootInvoked ||
		!result.LaunchAttempted ||
		!result.LaunchPIDRecorded ||
		!result.XWinInfoInvoked ||
		!result.XWindowObserved ||
		result.XWindowChildCount != 1 ||
		result.XWindowObservationAttempts == 0 ||
		result.WinebootStderrBytes == 0 ||
		result.LoopbackSSHForwardingOnly != true ||
		result.QEMURequired != true ||
		result.XvfbRequired != true ||
		result.QEMUUserNetworkRestrictDisabledForDisplay != true ||
		result.HostRootModified != false ||
		result.PrivilegedContainerRequired != false ||
		result.HostNetworkingRequired != false ||
		result.DockerSocketMounted != false ||
		result.BroadHostMountRequired != false ||
		result.RawHostPathExposed != false ||
		result.RawGuestGUIAppPathExposed != false ||
		result.RawCommandExposed != false {
		t.Fatalf("unexpected GUI smoke result: %#v", result)
	}
	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "wineboot --init") ||
		!strings.Contains(log, "WINEDEBUG='err+winediag'") ||
		!strings.Contains(log, "WINEDLLOVERRIDES='winemenubuilder.exe=d,mscoree,mshtml='") ||
		!strings.Contains(log, "grep 'appwiz[.]cpl install_mono'") ||
		!strings.Contains(log, "winemine.exe") ||
		!strings.Contains(log, "xwininfo display=:100") {
		t.Fatalf("GUI smoke did not run expected commands: %s", log)
	}
}

func TestRunGuestGUISmokeCopiesLocalGUIExecutableIntoGuest(t *testing.T) {
	if os.PathSeparator != '/' {
		t.Skip("shell fixtures require a POSIX host")
	}
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "guest-copy.log")
	executablePath := filepath.Join(tempDir, "hello-gui.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x014c), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'winex11'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *'wineboot --init'*) printf 'boot initialized\\n' >&2; exit 0 ;;\n" +
		"  *'wine '*'hello-gui.exe'*) exit 0 ;;\n" +
		"  *'cat '*'stderr.txt'*) printf ''; exit 0 ;;\n" +
		"  *'wineserver -k'*) exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	scpPath := filepath.Join(tempDir, "fake-scp")
	scpBody := "#!/bin/sh\n" +
		"printf 'scp %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"exit 0\n"
	if err := os.WriteFile(scpPath, []byte(scpBody), 0o700); err != nil {
		t.Fatalf("WriteFile scp returned error: %v", err)
	}
	xwininfoPath := filepath.Join(tempDir, "fake-xwininfo")
	xwininfoBody := "#!/bin/sh\n" +
		"printf 'xwininfo display=%s\\n' \"$DISPLAY\" >> '" + logPath + "'\n" +
		"printf 'xwininfo: Window id: 0x3a7 (the root window)\\n'\n" +
		"printf '  0x200001 \"Xnix Hello GUI\": ()  320x240+0+0  +0+0\\n'\n"
	if err := os.WriteFile(xwininfoPath, []byte(xwininfoBody), 0o700); err != nil {
		t.Fatalf("WriteFile xwininfo returned error: %v", err)
	}

	result, err := RunGuestGUISmoke(context.Background(), GuestGUIRequest{
		ExecutablePath: executablePath,
		Host:           DefaultGuestHost,
		Port:           "2222",
		User:           DefaultGuestUser,
		KeyPath:        filepath.Join(tempDir, "id_ed25519"),
		RemoteDir:      "/tmp/xnix-wine-guest-gui-smoke",
		SSHPath:        sshPath,
		SCPPath:        scpPath,
		XWinInfoPath:   xwininfoPath,
		GuestDisplay:   "10.0.2.2:100",
		HostDisplay:    ":100",
		Timeout:        5 * time.Second,
		Wait:           1 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("RunGuestGUISmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.GUIAppName != "hello-gui.exe" ||
		!result.ExecutableCopied ||
		!result.LaunchAttempted ||
		!result.GuestX11DriverAvailable ||
		!result.XWindowObserved ||
		result.XWindowObservationAttempts == 0 ||
		result.RawHostPathExposed ||
		result.RawGuestGUIAppPathExposed ||
		result.RawCommandExposed {
		t.Fatalf("unexpected copied GUI smoke result: %#v", result)
	}
	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "scp ") ||
		!strings.Contains(log, executablePath) ||
		!strings.Contains(log, "root@127.0.0.1:/tmp/xnix-wine-guest-gui-smoke/hello-gui.exe") ||
		!strings.Contains(log, "wine '/tmp/xnix-wine-guest-gui-smoke/hello-gui.exe'") {
		t.Fatalf("GUI smoke did not copy and launch the expected executable: %s", log)
	}
}
