package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
       	                      1  1  1  1  1  1
0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                      ID                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   | QR = 0, Opcode = 0, AA = 0, TC = 0, RD = 1, RA = 0, Z = 0, RCODE = 0
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    QDCOUNT                    | number of entries in the question section = 1
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    ANCOUNT                    | number of resource records in the answer section = 0
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    NSCOUNT                    | number of name server resource records in the authority records section = 0
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    ARCOUNT                    | number of resource records in the additional records section = 1
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
*/
type DnsHeader struct {
	Id      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

func main() {
	//    4    d    o    c    s    6    g    o    o    g   l     e    3    c   o    m   end
	// 0x04 0x64 0x6f 0x63 0x73 0x06 0x67 0x6f 0x6f 0x67 0x6c 0x65 0x03 0x63 0x6f 0x6d 0x00
	request_qname := [...]byte{0x04, 0x64, 0x6f, 0x63, 0x73, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00}

	dns_query := struct {
		Header DnsHeader
		QName  [len(request_qname)]byte
		QType  uint16
		QClass uint16
	}{
		DnsHeader{
			12345,  // id
			0x0100, // flags
			1,      // qd count
			0,      // an count
			0,      // ns count
			0,      // ar count
		},
		request_qname, // qname
		1,             // qtype
		1,             // qclass
	}

	// send request
	conn, err := net.Dial("udp", "8.8.8.8:53")
	check(err)
	binary.Write(conn, binary.BigEndian, dns_query)

	// get response
	response := make([]byte, 2048)
	_, err = bufio.NewReader(conn).Read(response)
	check(err)

	response_reader := bytes.NewReader(response)
	response_header := DnsHeader{}
	response_question := struct {
		QName  [len(request_qname)]byte
		QType  uint16
		QClass uint16
	}{}
	binary.Read(response_reader, binary.BigEndian, &response_header)
	binary.Read(response_reader, binary.BigEndian, &response_question)

	if response_header.Id != 12345 {
		panic("Wrong response id")
	}
	if response_header.QdCount != 1 {
		panic("Wrong response number of queries")
	}
	if response_question.QName != request_qname {
		panic("Wrong response qname")
	}
	if response_question.QType != 1 {
		panic("Wrong response qtype")
	}
	if response_question.QClass != 1 {
		panic("Wrong response qclass")
	}

	fmt.Println(response_question)
	conn.Close()
}
