package xargs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type XArgs struct {
	Command   string
	InputArgs []string
	BaseArgs  []string
}

func New(command string, baseArgs []string) *XArgs {
	return &XArgs{
		Command:  command,
		BaseArgs: baseArgs,
	}
}

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
