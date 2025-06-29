package workers

import "flag"

type FlagReaderStruct struct {
	Modify   bool
	Number   int
	ReadOnly bool
	Search   string
	//Unsafe bool
}

// FlagReader reads command line flags.
func FlagReader() FlagReaderStruct {
	var (
		number   *int    = flag.Int("number", 10, "Number of commands to show")
		n        *int    = flag.Int("n", 10, "Number of commands to show")
		readOnly *bool   = flag.Bool("read-only", false, "Not execute commands, just show them")
		modify   *bool   = flag.Bool("modify", false, "Allow the user to modify the selected command")
		m        *bool   = flag.Bool("m", false, "Allow the user to modify the selected command")
		search   *string = flag.String("search", "", "Search for a specific command in the history")
		s        *string = flag.String("s", "", "Search for a specific command in the history")
		//unsafe *bool   = flag.Bool("unsafe", false, "Run unsafe commands")
	)

	flag.Parse()

	finalNumber := *number
	if *n != 10 {
		finalNumber = *n
	}

	finalSearch := *search
	if *s != "" {
		finalSearch = *s
	}

	finalModify := *modify
	if *m {
		finalModify = *m
	}

	return FlagReaderStruct{
		Modify:   finalModify,
		Number:   finalNumber,
		ReadOnly: *readOnly,
		Search:   finalSearch,
		//Unsafe: *unsafe,
	}

}

func LastNCommands(slice []string, n int) []string {
	if len(slice) < n {
		return slice
	}
	return slice[len(slice)-n:]
}
