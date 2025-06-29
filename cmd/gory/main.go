package main

import (
	"fmt"
	"strings"

	"github.com/rojack96/gory/pkg/workers"

	"github.com/charmbracelet/huh"
	"github.com/rojack96/treje"
	"github.com/rojack96/treje/set/types"
)

// returns a list of unique commands from the system's command history
func getCommands(system workers.System, fr workers.FlagReaderStruct) []huh.Option[string] {
	var (
		result           []huh.Option[string]
		filteredCommands []string
	)

	listOfUniqueCommand, _ := treje.NewSet().String()

	for i := range system.Commands {
		listOfUniqueCommand.Add(types.Str(system.Commands[i]))
	}

	listOfCommand, _ := listOfUniqueCommand.ToSlice()
	listOfCommandToView := listOfCommand
	// filter command if search is provided
	if fr.Search != "" {
		for _, cmd := range listOfCommand {
			if strings.Contains(cmd, fr.Search) {
				filteredCommands = append(filteredCommands, cmd)
			}
		}
		listOfCommandToView = filteredCommands
	}

	listOfCommandToView = workers.LastNCommands(listOfCommandToView, fr.Number)

	for _, cmd := range listOfCommandToView {
		if strings.Contains(cmd, fr.Search) {
			result = append(result, huh.NewOption(cmd, cmd))
		}
	}

	return result
}

func form(options []huh.Option[string], fr workers.FlagReaderStruct) (string, bool) {
	var (
		command string
		run     bool
		err     error
	)

	if len(options) == 0 {
		fmt.Println("No commands found in history.")
		return "", false
	}

	form := huh.NewSelect[string]().
		Title("Choose command to run:").
		Options(options...).
		Value(&command)

	if err = form.Run(); err != nil {
		fmt.Println("Error running select:", err)
		return "", false
	}

	if fr.ReadOnly {
		return command, true
	}

	if fr.Modify {
		huh.NewInput().Value(&command).Run()
	}

	confirm := huh.NewConfirm().
		Title(`Are you sure to run "` + command + `"?`).
		Value(&run)

	if err = confirm.Run(); err != nil {
		fmt.Println("Error running confirmation:", err)
		return "", false
	}

	return command, run
}

func main() {
	var (
		commands []huh.Option[string]
		command  string
		run      bool
		err      error
	)

	system := workers.System{}

	if err = system.ReadHistory(); err != nil {
		fmt.Println("error reading history file 📂❌")
		return
	}

	fr := workers.FlagReader()

	if commands = getCommands(system, fr); len(commands) == 0 {
		fmt.Println("No commands found in history.")
		return
	}

	if command, run = form(commands, fr); !run {
		fmt.Println("Bye bye 👋")
		return
	}

	// if !fr.Unsafe {
	// 	if workers.IsDangerousCommand(command) {
	// 		fmt.Println("⚠️ This command is considered unsafe and will not be executed.")
	// 		return
	// 	}
	// }

	fmt.Println(command)
	if !fr.ReadOnly {
		system.RunCommand(command)
	}

}
