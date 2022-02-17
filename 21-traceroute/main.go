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


/*
0                   1                   2                   3
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|     Type      |      Code     |          Checksum             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           Identifier          |        Sequence Number        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

The checksum is the 16-bit ones's complement of the one's
complement sum of the ICMP message starting with the ICMP Type.
For computing the checksum , the checksum field should be zero.
This checksum may be replaced in the future.
*/
type Icmp struct {
	Type           uint8	// 8 for ping request
	Code           uint8	// 0
	Checksum       uint16
	Identifier     uint16
	SequenceNumber uint16
}

/*
data is this? wtf?
0000   08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17
0010   18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27
0020   28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37
*/

func main() {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 1)
	check(err)
	defer syscall.Close(sock)

	fmt.Println(sock)

	// sockAddr := &syscall.SockaddrInet4{
	// 	Port: 1991,
	// 	Addr: [4]byte{127, 0, 0, 1},
	// }
	// err = syscall.Bind(sock, sockAddr)
	// check(err)

	// err = syscall.Listen(sock, 128)
	// check(err)

	// for {
	// 	client_fd, addr, err := syscall.Accept(sock)
	// 	check(err)

	// 	fmt.Println(client_fd)
	// 	fmt.Println(addr)
	// 	fmt.Println(sock)

	// 	for {
	// 		buf := make([]byte, 1024)
	// 		_, _, err = syscall.Recvfrom(client_fd, buf, 0)
	// 		check(err)

	// 		err = syscall.Sendto(client_fd, buf, 0, sockAddr)
	// 		check(err)
	// 	}
	// }
}
