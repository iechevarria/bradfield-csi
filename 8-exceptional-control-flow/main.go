package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func run (args []string) {
	if len(args) == 0 {
		return
	}

	fn := args[0];

	switch fn {
	case "exit":
		// TODO: raise some kind of exception here
		fmt.Println("You suck!")
	default:
		// try to fork + exec
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr		
		err := cmd.Run()
		if err != nil {
			fmt.Println("🐒💩 " + args[0] + ": command not found")
		}
	}
}

func main () {
	var sb strings.Builder
	in := bufio.NewReader(os.Stdin)
	fmt.Print("🐵 ")

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("\n🙉🙈🙊")
			break
		} else if r == '\n' {
			line := sb.String()
			args := strings.Fields(line)
			run(args)

			// reset terminal
			fmt.Print("🐵 ")
			sb.Reset()
		} else {
			sb.WriteRune(r)
		}	
	}
}
