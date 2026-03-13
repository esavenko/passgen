package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/app"
	"github.com/esavenko/passgen/internal/constants"
	flag "github.com/spf13/pflag"
)

var (
	version = "dev"
	commit  = "none" //nolint:unused // injected via ldflags by goreleaser
	date    = "unknown"
)

// commandExecutor runs an external command, writing its output to the provided writers.
type commandExecutor func(stdout, stderr io.Writer) error

func defaultExecutor() commandExecutor {
	return func(stdout, stderr io.Writer) error {
		cmd := exec.Command("bash", "-c", "curl -fsSL https://raw.githubusercontent.com/esavenko/passgen/master/install.sh | sudo bash")
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		return cmd.Run()
	}
}

func run(args []string, stdout, stderr io.Writer, executor commandExecutor) int {
	fs := flag.NewFlagSet("passgen", flag.ContinueOnError)
	fs.SetOutput(stderr)

	showVersion := fs.BoolP("version", "v", false, "Print version information")
	runUpdate := fs.BoolP("update", "u", false, "Update passgen to the latest version")

	fs.Usage = func() {
		fmt.Fprintln(stdout, constants.Logo)
		fmt.Fprintln(stdout)
		fmt.Fprintln(stdout, "A terminal-based password generator with an interactive TUI.")
		fmt.Fprintln(stdout)
		fmt.Fprintln(stdout, "Usage:")
		fmt.Fprintln(stdout, "  passgen [flags]")
		fmt.Fprintln(stdout)
		fmt.Fprintln(stdout, "Flags:")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return 1
	}

	switch {
	case *showVersion:
		fmt.Fprintf(stdout, "passgen %s ( built: %s)\n", version, date)
	case *runUpdate:
		fmt.Fprintln(stdout, "Updating passgen to the latest version...")
		if err := executor(stdout, stderr); err != nil {
			fmt.Fprintf(stderr, "Update failed: %v\n", err)
			return 1
		}
	default:
		p := tea.NewProgram(app.GetAppModel())
		if _, err := p.Run(); err != nil {
			fmt.Fprintln(stderr, "Could not start program", err)
			return 1
		}
	}
	return 0
}

func main() {
	code := run(os.Args[1:], os.Stdout, os.Stderr, defaultExecutor())
	if code != 0 {
		os.Exit(code)
	}
}
