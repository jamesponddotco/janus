// Package app holds the main entry point for the application.
package app

import (
	"flag"
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/janus/internal/build"
	"git.sr.ht/~jamesponddotco/janus/internal/diff"
)

func Run() int {
	var (
		noColor = flag.Bool("no-color", false, "do not colorize the output")
		version = flag.Bool("version", false, "display version information")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <oldFile> <newFile>\nOptions:\n", build.Name)

		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Fprintln(os.Stdout, "Version: ", build.Version)

		return 0
	}

	if flag.NArg() != 2 {
		flag.Usage()

		return 1
	}

	var (
		oldFile = flag.Arg(0)
		newFile = flag.Arg(1)
	)

	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		*noColor = true
	}

	opts := &diff.Options{
		Colorize: !*noColor,
	}

	result, err := diff.Compare(oldFile, newFile, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)

		return 1
	}

	fmt.Fprint(os.Stdout, result)

	return 0
}
