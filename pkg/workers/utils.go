package workers

import "flag"

type FlagReaderStruct struct {
	Number int
	Search string
	//Unsafe bool
	ReadOnly bool
}

// FlagReader reads command line flags.
func FlagReader() FlagReaderStruct {
	var (
		number   *int    = flag.Int("number", 10, "Number of commands to show")
		n        *int    = flag.Int("n", 10, "Number of commands to show")
		search   *string = flag.String("search", "", "Search for a specific command in the history")
		s        *string = flag.String("s", "", "Search for a specific command in the history")
		readOnly *bool   = flag.Bool("read-only", false, "Not execute commands, just show them")
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

	return FlagReaderStruct{
		Number:   finalNumber,
		Search:   finalSearch,
		ReadOnly: *readOnly,
		//Unsafe: *unsafe,
	}

}

func LastNCommands(slice []string, n int) []string {
	if len(slice) < n {
		return slice
	}
	return slice[len(slice)-n:]
}
