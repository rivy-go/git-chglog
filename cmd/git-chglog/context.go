package main

import (
	"io"
)

// CLIContext ...
type CLIContext struct {
	WorkingDir       string
	Stdout           io.Writer
	Stderr           io.Writer
	ConfigPath       string
	OutputPath       string
	Silent           bool
	NoColor          bool
	NoEmoji          bool
	NoCaseSensitive  bool
	Query            string
	NextTag          string
	TagFilterPattern string
	Unreleased       bool
}

// InitContext ...
type InitContext struct {
	WorkingDir string
	Stdout     io.Writer
	Stderr     io.Writer
}
