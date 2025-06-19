package workers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	Bash = "bash"
	Zsh  = "zsh"
	Fish = "fish"
)

type System struct {
	Shell       string
	HistoryFile string
	Commands    []string
}

func (s *System) getShell() error {
	pid := os.Getppid()

	for {
		cmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "comm=")
		output, err := cmd.Output()
		if err != nil {
			return err
		}

		parent := strings.TrimSpace(string(output))

		if parent == Bash {
			s.Shell = parent
			s.HistoryFile = ".bash_history"
			return nil
		} else if parent == Zsh {
			s.Shell = parent
			s.HistoryFile = ".zsh_history"
			return nil
		} else if parent == Fish {
			s.Shell = parent
			s.HistoryFile = ".config/fish/fish_history"
			return nil
		}

		// Get next parent PID
		cmd = exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "ppid=")
		output, err = cmd.Output()
		if err != nil {
			return err
		}

		pidStr := strings.TrimSpace(string(output))
		pid, err = strconv.Atoi(pidStr)
		if err != nil {
			return err
		}
	}
}

// ReadHistory reads the bash history file from the user's home directory and returns a slice of commands.
func (s *System) ReadHistory() error {
	var (
		err error
	)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error to get home directory: %w", err)
	}

	if err = s.getShell(); err != nil {
		return fmt.Errorf("error to get shell: %w", err)
	}

	historyPath := filepath.Join(homeDir, s.HistoryFile)

	file, err := os.Open(historyPath)
	if err != nil {
		return fmt.Errorf("error to open history file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if s.Shell == Zsh {
			// Zsh history lines can have timestamps like:
			// : 1687154450:0;git status
			// Let's strip the timestamp if present
			if idx := strings.Index(line, ";"); idx != -1 {
				line = line[idx+1:]
			}

		}

		s.Commands = append(s.Commands, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading history file: %w", err)
	}

	return nil
}

// RunCommand executes a command with the given name and arguments.
func (s *System) RunCommand(args string) {
	name := s.Shell
	cmd := exec.Command(name, "-c", args)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command")
	}
}

// func IsDangerousCommand(cmd string) bool {
// 	dangerousPatterns := []string{
// 		"rm -rf /",     // distrugge tutto il filesystem
// 		"rm -rf *",     // elimina tutto nella dir corrente
// 		"sudo",         // potenzialmente pericoloso
// 		"mkfs",         // formatta filesystem
// 		":(){ :|:& };", // fork bomb
// 		"shutdown",     // spegne il sistema
// 		"reboot",       // riavvia
// 	}

// 	for _, pattern := range dangerousPatterns {
// 		if strings.Contains(cmd, pattern) {
// 			return true
// 		}
// 	}

// 	// eventualmente, regex o controlli più intelligenti
// 	return false
// }
