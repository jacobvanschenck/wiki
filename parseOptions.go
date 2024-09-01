package main

type Flags struct {
	intro bool
}

func parseFlagsString(flags []string) Flags {
	var options Flags
	for _, flag := range flags {
		switch flag {
		case "-i":
			options.intro = true
		}
	}

	return options
}
