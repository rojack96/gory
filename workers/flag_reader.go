package workers

import "flag"

type FlagReaderStruct struct {
	Number int
	Search string
}

// FlagReader reads command line flags.
func FlagReader() FlagReaderStruct {
	var (
		number *int    = flag.Int("number", 10, "Number of commands to show")
		n      *int    = flag.Int("n", 10, "Number of commands to show")
		search *string = flag.String("search", "", "Search for a specific command in the history")
		s      *string = flag.String("s", "", "Search for a specific command in the history")
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
		Number: finalNumber,
		Search: finalSearch,
	}

}
