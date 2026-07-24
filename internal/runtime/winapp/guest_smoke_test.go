package winapp

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestRunGuestSmokeCopiesExecutableAndObservesMarker(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	logPath := filepath.Join(tempDir, "guest.log")
	sshPath := writeFakeGuestSSH(t, tempDir, logPath, true)
	scpPath := writeFakeGuestSCP(t, tempDir, logPath)

	result, err := RunGuestSmoke(context.Background(), GuestRequest{
		ExecutablePath: executablePath,
		Host:           "127.0.0.1",
		Port:           "2222",
		User:           "root",
		KeyPath:        filepath.Join(tempDir, "id_ed25519"),
		RemoteDir:      "/tmp/xnix-winapp-smoke",
		SSHPath:        sshPath,
		SCPPath:        scpPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunGuestSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.ExecutableName != "hello.exe" ||
		result.GuestTransport != "loopback-ssh" ||
		!result.GuestReachable ||
		!result.WineAvailable ||
		!result.ExecutableCopied ||
		!result.MarkerObserved ||
		result.ExitCode != 0 ||
		result.ExpectedMarker != DefaultMarker {
		t.Fatalf("unexpected guest smoke result: %#v", result)
	}
	if !result.LoopbackOnlyNetworking ||
		!result.QEMURequired ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed {
		t.Fatalf("unexpected safety flags: %#v", result)
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	for _, token := range []string{
		"ssh -p 2222",
		"root@127.0.0.1 true",
		"command -v wine",
		"mkdir -p '/tmp/xnix-winapp-smoke'",
		"scp -P 2222",
		"hello.exe root@127.0.0.1:/tmp/xnix-winapp-smoke/hello.exe",
		"WINEPREFIX='/tmp/xnix-winapp-smoke/wineprefix' WINEDEBUG=-all wine '/tmp/xnix-winapp-smoke/hello.exe'",
	} {
		if !strings.Contains(log, token) {
			t.Fatalf("guest log missing %q: %s", token, log)
		}
	}
}

func TestRunGuestSmokeSkipsWhenGuestWineUnavailable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	logPath := filepath.Join(tempDir, "guest.log")
	sshPath := writeFakeGuestSSH(t, tempDir, logPath, false)
	scpPath := writeFakeGuestSCP(t, tempDir, logPath)

	result, err := RunGuestSmoke(context.Background(), GuestRequest{
		ExecutablePath: executablePath,
		SSHPath:        sshPath,
		SCPPath:        scpPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunGuestSmoke returned error: %v", err)
	}
	if result.Status != SkippedStatus ||
		result.SkipReason != "guest wine runner unavailable" ||
		!result.GuestReachable ||
		result.WineAvailable ||
		result.ExecutableCopied {
		t.Fatalf("unexpected missing wine result: %#v", result)
	}
}

func TestResolveToolAcceptsPathCommandNames(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell PATH fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	toolPath := filepath.Join(tempDir, "fake-tool")
	if err := os.WriteFile(toolPath, []byte("#!/bin/sh\nexit 0\n"), 0o700); err != nil {
		t.Fatalf("WriteFile fake tool returned error: %v", err)
	}
	t.Setenv("PATH", tempDir)

	resolved, err := resolveTool("fake-tool", "fallback-tool")
	if err != nil {
		t.Fatalf("resolveTool returned error: %v", err)
	}
	if resolved != toolPath {
		t.Fatalf("resolveTool resolved %q, want %q", resolved, toolPath)
	}
}

func TestParseGuestPortForQEMUAcceptsAutoOnlyWhenStartingQEMU(t *testing.T) {
	port, err := ParseGuestPortForQEMU("auto", true)
	if err != nil {
		t.Fatalf("ParseGuestPortForQEMU returned error: %v", err)
	}
	if port != AutoGuestPort {
		t.Fatalf("ParseGuestPortForQEMU returned %q, want %q", port, AutoGuestPort)
	}
	if _, err := ParseGuestPortForQEMU("auto", false); err == nil {
		t.Fatalf("ParseGuestPortForQEMU accepted auto without QEMU start")
	}
}

func writeFakeGuestSSH(t *testing.T, tempDir string, logPath string, wineAvailable bool) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-ssh")
	wineExit := "1"
	if wineAvailable {
		wineExit = "0"
	}
	body := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit " + wineExit + " ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *' wine '*'hello.exe'*) printf 'XNIX_WINAPP_SMOKE_OK\\n'; exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile fake ssh returned error: %v", err)
	}
	return path
}

func writeFakeGuestSCP(t *testing.T, tempDir string, logPath string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-scp")
	body := "#!/bin/sh\n" +
		"printf 'scp %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"exit 0\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile fake scp returned error: %v", err)
	}
	return path
}
