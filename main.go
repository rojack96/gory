package main

import (
	"fmt"
	"strings"

	"github.com/rojack96/gory/workers"

	"github.com/charmbracelet/huh"
	"github.com/rojack96/treje"
	"github.com/rojack96/treje/set/types"
)

// Esempio d'uso
func main() {
	var (
		result  []huh.Option[string]
		command string
		run     bool
	)

	fr := workers.FlagReader()
	listOfUniqueCommand, _ := treje.NewSet().String()

	commands, err := workers.ReadBashHistory()
	if err != nil {
		fmt.Printf("Error reading bash history: %v", err)
		return
	}

	for i, _ := range commands {
		listOfUniqueCommand.Add(types.Str(commands[i]))
	}

	listOfCommand, _ := listOfUniqueCommand.ToSlice()
	listOfCommandToView := listOfCommand

	if fr.Search == "" {
		listOfCommandToView = listOfCommand[len(listOfCommand)-fr.Number:]
	}

	for _, cmd := range listOfCommandToView {
		if strings.Contains(cmd, fr.Search) {
			result = append(result, huh.NewOption(cmd, cmd))
		}
	}

	form := huh.NewSelect[string]().
		Title("Choose command to run:").
		Options(result...).
		Value(&command)

	if err = form.Run(); err != nil {
		fmt.Println("Error running select:", err)
		return
	}

	test := huh.NewConfirm().
		Title(fmt.Sprintf(`Are you sure to run "%s"?`, command)).
		Value(&run)

	if err = test.Run(); err != nil {
		fmt.Println("Error running confirmation:", err)
		return
	}

	if run {
		fmt.Println(command)
		workers.RunCommand("bash", "-c", command)
	}

	return
}
