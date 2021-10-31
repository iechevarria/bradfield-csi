package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func run (args []string) error {
	if len(args) == 0 {
		return nil
	}

	fn := args[0];

	switch fn {
	case "exit":
		// pretty weird to use io.EOF as the err here
		return io.EOF
	case "cd":
		if len(args) < 2 {
			return nil
		}
		return syscall.Chdir(args[1])
	default:
		// errors are ugly here. ideally process them
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return nil
}

func main () {
	var sb strings.Builder
	in := bufio.NewReader(os.Stdin)
	fmt.Print("🐵 ")

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		} else if r == '\n' {
			line := sb.String()
			args := strings.Fields(line)
			err := run(args)
			if err == io.EOF {
				break	
			}

			if err != nil {
				fmt.Println("🐒💩 " + args[0] + ": " + err.Error())
			}

			// reset terminal
			fmt.Print("🐵 ")
			sb.Reset()
		} else {
			sb.WriteRune(r)
		}	
	}

	fmt.Println("\n🙉🙈🙊")
}
