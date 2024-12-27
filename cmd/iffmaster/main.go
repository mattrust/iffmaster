// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

/*
Iffmaster is a tool to inspect IFF files as defined under EA-85.

Usage: [options]

	iffmaster

		-version: Show the application's version.
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
