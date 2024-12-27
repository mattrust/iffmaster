// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

/*
Iffmaster is a tool to inspect IFF files as defined under EA-85.

Usage: [options] filename

	iffmaster

		-version: Show the application's version.

		filename: The IFF file to inspect (optional).
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattrust/iffmaster/internal/gui"
)

func main() {
	var filename string

	showVersion := flag.Bool("version", false, "Display the version of iffmaster")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(),
			"\n  filename: The IFF file to inspect (optional).\n")
	}
	flag.Parse()

	if *showVersion {
		fmt.Println("iffmaster version", Version)
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		filename = flag.Arg(0)
	}

	gui.OpenGUI(filename, Version)
}
