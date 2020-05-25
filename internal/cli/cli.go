package cli

import (
	"flag"
	"fmt"
	"os"
)

// DefaultPort is a default port for serve command
const DefaultPort string = "9005"

// Executor is an interface for commands
type Executor interface {
	Execute([]string) error
}

var commands = map[string]Executor{
	"serve": Serve{},
}

// Run function reads flags and arguments, and then decides what to execute
func Run(version string, buildDate string) {
	helpPtr := flag.Bool("help", false, "")
	versionPtr := flag.Bool("version", false, "")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr,
			`Usage: pfms COMMAND [arg...]
   or: pfms -help
   or: pfms -version
Options:
      -help           Print usage
      -version        Print version and build date
Commands:
       serve [port]   Start a server (default port `+DefaultPort+`)`, "\n")
	}

	flag.Parse()

	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if *versionPtr {
		fmt.Fprintf(os.Stderr, "Personal Finance Management System\nVersion:    %s\nBuild date: %s\n", version, buildDate)
		os.Exit(0)
	}

	if len(os.Args) > 1 {
		c, ok := commands[os.Args[1]]
		if ok {
			err := c.Execute(os.Args[2:])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "command provided but not defined: %s\n", os.Args[1])
	}

	flag.Usage()

	// If a command is not found we exit with a status 2 to match the behavior
	// of flag.Parse() with flag.ExitOnError when parsing an invalid flag.
	os.Exit(2)
}
