package main

import (
	"fmt"
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	check(err)
	defer syscall.Close(sock)

	sockAddr := &syscall.SockaddrInet4{
		Port: 1991,
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Bind(sock, sockAddr)
	check(err)

	err = syscall.Listen(sock, 128)
	check(err)

	for {
		client_fd, addr, err := syscall.Accept(sock)
		check(err)

		fmt.Println(client_fd)
		fmt.Println(addr)
		fmt.Println(sock)

		for {
			buf := make([]byte, 1024)
			_, _, err = syscall.Recvfrom(client_fd, buf, 0)
			check(err)

			err = syscall.Sendto(client_fd, buf, 0, sockAddr)
			check(err)
		}
	}
}
