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
		Port: 6969,
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Bind(sock, sockAddr)
	check(err)

	err = syscall.Listen(sock, 1)
	check(err)

	nfd, addr, err := syscall.Accept(sock)
	check(err)

	buf := make([]byte, 1024)
	nfd, addr, err = syscall.Recvfrom(nfd, buf, 0)
	fmt.Println(buf)
	check(err)

	err = syscall.Sendto(sock, buf, 0, sockAddr)
	check(err)

	fmt.Println(buf)
	fmt.Println(addr)
}
