package workers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ReadBashHistory reads the bash history file from the user's home directory and returns a slice of commands.
func ReadBashHistory() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("errore nel recuperare la home directory: %w", err)
	}

	historyPath := filepath.Join(homeDir, ".bash_history")

	file, err := os.Open(historyPath)
	if err != nil {
		return nil, fmt.Errorf("errore nell'apertura di %s: %w", historyPath, err)
	}
	defer file.Close()

	var commands []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("errore durante la lettura del file: %w", err)
	}

	return commands, nil
}
