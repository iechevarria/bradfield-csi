package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

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

			// try to fork + exec
			args := strings.Fields(line)
			cmd := exec.Command(args[0], args[1:]...)

			// hook up command stdin/stdout to the shell's and run
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			err := cmd.Run()
			if err != nil {
				fmt.Println("🐒💩 " + args[0] + ": command not found")
			}

			// fmt.Println(line)

			// reset terminal
			fmt.Print("🐵 ")
			sb.Reset()

		} else {
			sb.WriteRune(r)
		}	
	}
}
