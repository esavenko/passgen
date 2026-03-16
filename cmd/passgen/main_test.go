package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/esavenko/passgen/internal/constants"
)

// successExecutor simulates a successful update command.
func successExecutor() commandExecutor {
	return func(stdout, stderr io.Writer) error {
		return nil
	}
}

// failExecutor simulates a failed update command.
func failExecutor(errMsg string) commandExecutor {
	return func(stdout, stderr io.Writer) error {
		return errors.New(errMsg)
	}
}

// nopExecutor is a no-op executor for tests that should not reach the update path.
func nopExecutor() commandExecutor {
	return func(_, _ io.Writer) error {
		return nil
	}
}

// --- Version flag tests ---

func TestVersionFlagLong(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version"}, &stdout, &stderr, nopExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	want := fmt.Sprintf("passgen %s ( built: %s)", version, date)
	if !strings.Contains(out, want) {
		t.Errorf("stdout = %q, want it to contain %q", out, want)
	}
}

func TestVersionFlagShort(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"-v"}, &stdout, &stderr, nopExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	want := fmt.Sprintf("passgen %s ( built: %s)", version, date)
	if !strings.Contains(out, want) {
		t.Errorf("stdout = %q, want it to contain %q", out, want)
	}
}

func TestVersionOutputFormat(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version"}, &stdout, &stderr, nopExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	// Should contain version and date values
	if !strings.Contains(out, version) {
		t.Errorf("version output should contain version %q, got %q", version, out)
	}
	if !strings.Contains(out, date) {
		t.Errorf("version output should contain date %q, got %q", date, out)
	}
	// Should not write to stderr
	if stderr.Len() != 0 {
		t.Errorf("stderr should be empty, got %q", stderr.String())
	}
}

// --- Update flag tests ---

func TestUpdateFlagLongSuccess(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--update"}, &stdout, &stderr, successExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	if !strings.Contains(out, "Updating passgen to the latest version...") {
		t.Errorf("stdout should contain update message, got %q", out)
	}
}

func TestUpdateFlagShortSuccess(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"-u"}, &stdout, &stderr, successExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	if !strings.Contains(out, "Updating passgen") {
		t.Errorf("stdout should contain update message, got %q", out)
	}
}

func TestUpdateFlagFailure(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--update"}, &stdout, &stderr, failExecutor("command not found"))

	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}

	errOut := stderr.String()
	if !strings.Contains(errOut, "Update failed") {
		t.Errorf("stderr should contain 'Update failed', got %q", errOut)
	}
	if !strings.Contains(errOut, "command not found") {
		t.Errorf("stderr should contain error message, got %q", errOut)
	}
}

func TestUpdateFlagFailureWritesToStderr(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"-u"}, &stdout, &stderr, failExecutor("network error"))

	// The update progress message goes to stdout
	if !strings.Contains(stdout.String(), "Updating passgen") {
		t.Errorf("stdout should contain update progress message")
	}
	// The error goes to stderr
	if !strings.Contains(stderr.String(), "Update failed: network error") {
		t.Errorf("stderr should contain failure details, got %q", stderr.String())
	}
}

// --- Help / Usage tests ---

func TestHelpFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"--help"}, &stdout, &stderr, nopExecutor())

	// The usage function writes to stdout, but pflag may also write to stderr
	combined := stdout.String() + stderr.String()
	if !strings.Contains(combined, "passgen") {
		t.Errorf("help output should mention 'passgen', got stdout=%q stderr=%q", stdout.String(), stderr.String())
	}
}

// --- Invalid flag tests ---

func TestInvalidFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--nonexistent"}, &stdout, &stderr, nopExecutor())

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for invalid flag", code)
	}
}

func TestInvalidShortFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"-z"}, &stdout, &stderr, nopExecutor())

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for invalid flag", code)
	}
}

// --- Flag precedence tests ---

func TestVersionTakesPrecedenceOverUpdate(t *testing.T) {
	// When both --version and --update are passed, version should be checked first
	// because it appears first in the switch statement
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version", "--update"}, &stdout, &stderr, nopExecutor())

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}

	out := stdout.String()
	if !strings.Contains(out, "passgen") {
		t.Errorf("should print version info, got %q", out)
	}
	if strings.Contains(out, "Updating") {
		t.Errorf("should not print update message when version flag is also set")
	}
}

// --- No flags (default path) tests ---

func TestNoFlagsSkipped(t *testing.T) {
	t.Skip("requires a terminal to run bubbletea")
}

// --- Usage output content validation ---

func TestUsageContainsLogo(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"--help"}, &stdout, &stderr, nopExecutor())

	combined := stdout.String() + stderr.String()
	if !strings.Contains(combined, constants.Logo) {
		t.Errorf("usage output should contain the logo")
	}
}

func TestUsageContainsDescription(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"--help"}, &stdout, &stderr, nopExecutor())

	combined := stdout.String() + stderr.String()
	if !strings.Contains(combined, "terminal-based password generator") {
		t.Errorf("usage should contain description, got %q", combined)
	}
}

func TestUsageContainsFlagDescriptions(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"--help"}, &stdout, &stderr, nopExecutor())

	combined := stdout.String() + stderr.String()

	for _, want := range []string{"version", "update"} {
		if !strings.Contains(combined, want) {
			t.Errorf("usage should mention %q flag, got %q", want, combined)
		}
	}
}

// --- Executor interaction tests ---

func TestUpdateExecutorIsCalled(t *testing.T) {
	called := false
	executor := func(_, _ io.Writer) error {
		called = true
		return nil
	}

	var stdout, stderr bytes.Buffer
	run([]string{"--update"}, &stdout, &stderr, executor)

	if !called {
		t.Error("executor should have been called for --update flag")
	}
}

func TestVersionDoesNotCallExecutor(t *testing.T) {
	called := false
	executor := func(_, _ io.Writer) error {
		called = true
		return nil
	}

	var stdout, stderr bytes.Buffer
	run([]string{"--version"}, &stdout, &stderr, executor)

	if called {
		t.Error("executor should not be called for --version flag")
	}
}

// --- Table-driven comprehensive flag tests ---

func TestFlagBehaviors(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		wantCode      int
		wantStdout    string
		wantStderr    string
		executorFails bool
	}{
		{
			name:       "long version flag",
			args:       []string{"--version"},
			wantCode:   0,
			wantStdout: "passgen",
		},
		{
			name:       "short version flag",
			args:       []string{"-v"},
			wantCode:   0,
			wantStdout: "passgen",
		},
		{
			name:       "long update flag success",
			args:       []string{"--update"},
			wantCode:   0,
			wantStdout: "Updating passgen",
		},
		{
			name:       "short update flag success",
			args:       []string{"-u"},
			wantCode:   0,
			wantStdout: "Updating passgen",
		},
		{
			name:          "update flag failure",
			args:          []string{"--update"},
			wantCode:      1,
			wantStderr:    "Update failed",
			executorFails: true,
		},
		{
			name:     "unknown flag",
			args:     []string{"--foo"},
			wantCode: 1,
		},
		{
			name:     "unknown short flag",
			args:     []string{"-x"},
			wantCode: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			var executor commandExecutor
			if tt.executorFails {
				executor = failExecutor("simulated failure")
			} else {
				executor = nopExecutor()
			}

			code := run(tt.args, &stdout, &stderr, executor)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d (stdout=%q, stderr=%q)", code, tt.wantCode, stdout.String(), stderr.String())
			}
			if tt.wantStdout != "" && !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, want it to contain %q", stdout.String(), tt.wantStdout)
			}
			if tt.wantStderr != "" && !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, want it to contain %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}
