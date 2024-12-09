// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package xargs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// XArgs represents the structure for managing arguments passed to a command.
// It stores the command to be executed, input arguments from stdin, and base arguments.
type XArgs struct {
	Command   string   // The command to be executed (e.g., "ls", "grep")
	InputArgs []string // Arguments read from stdin
	BaseArgs  []string // Base arguments passed during initialization
}

// New initializes a new XArgs instance with a specified command and base arguments.
// The command is the executable to be run, and baseArgs are arguments that are passed with the command.
func New(command string, baseArgs []string) *XArgs {
	return &XArgs{
		Command:  command,
		BaseArgs: baseArgs,
	}
}

// ParseCommandLine reads from the provided input reader (usually stdin) and splits the input into arguments.
// It appends each line's arguments to the `InputArgs` field of the XArgs struct.
func (x *XArgs) ParseCommandLine(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)

		x.InputArgs = append(x.InputArgs, args...)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input %v", err)
	}

	return nil
}

// Execute runs the command with the base arguments and the arguments accumulated from stdin.
// The results are written to the provided output writer. Any errors during execution are returned.
func (x *XArgs) Execute(out io.Writer) error {
	allArgs := append(x.BaseArgs, x.InputArgs...)
	cmd := exec.Command(x.Command, allArgs...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command execution error '%s': %v", x.Command, err)
	}

	return nil
}
