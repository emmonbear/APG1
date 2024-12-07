package xargs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type XArgs struct {
	InputArgs []string
}

func New() *XArgs {
	return &XArgs{}
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
