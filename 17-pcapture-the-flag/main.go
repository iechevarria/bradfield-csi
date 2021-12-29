package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

/*
pcap file header:
+------------------------------+
|        Magic number          | 32
+--------------+---------------+
|Major version | Minor version | 16 | 16
+--------------+---------------+
|      Time zone offset        | 32
+------------------------------+
|     Time stamp accuracy      | 32
+------------------------------+
|       Snapshot length        | 32
+------------------------------+
|   Link-layer header type     | 32
+------------------------------+
*/
type PcapFileHeader struct {
	MagicNumber         uint32 // 0xa1b2c3d4
	MajorVersion        uint16 // 2
	MinorVersion        uint16 // 4
	TzOffset            uint32 // 0
	TsAccuracy          uint32 // 0
	SnapshotLength      uint32 // 1514
	LinkLayerHeaderType uint32 // 1
}

func check_pcap_file_header(pcap PcapFileHeader) {
	if pcap.MagicNumber != 0xa1b2c3d4 {
		panic(fmt.Sprintf("Wrong magic number: %x", pcap.MagicNumber))
	}
	if pcap.MajorVersion != 2 {
		panic(fmt.Sprintln("Wrong major version:", pcap.MajorVersion))
	}
	if pcap.MinorVersion != 4 {
		panic(fmt.Sprintln("Wrong minor version:", pcap.MinorVersion))
	}
	if pcap.TzOffset != 0 {
		panic(fmt.Sprintln("Wrong tz offset:", pcap.TzOffset))
	}
	if pcap.TsAccuracy != 0 {
		panic(fmt.Sprintln("Wrong ts accuracy:", pcap.TsAccuracy))
	}
	// link layer header 1 is Ethernet
	if pcap.LinkLayerHeaderType != 1 {
		panic(fmt.Sprintln("Wrong link layer header type:", pcap.LinkLayerHeaderType))
	}
}

/*
pcap packet header:
+----------------------------------------------+
|          Time stamp, seconds value           |
+----------------------------------------------+
|Time stamp, microseconds or nanoseconds value |
+----------------------------------------------+
|       Length of captured packet data         |
+----------------------------------------------+
|   Un-truncated length of the packet data     |
+----------------------------------------------+
*/
type PcapPacketHeader struct {
	TsSeconds         uint32
	TsMicroseconds    uint32
	Length            uint32
	UntruncatedLength uint32
}

func check_pcap_packet_header(pcap PcapPacketHeader) {
	if pcap.Length != pcap.UntruncatedLength {
		panic(fmt.Sprintln("Length", pcap.Length, "!= UntruncatedLength", pcap.UntruncatedLength))
	}
}

type EthernetHeader struct {
	MACDestination [6]byte
	MACSource      [6]byte
	EtherType      uint16 // 0x800 - IPv4
}

func check_ethernet_header(ethernet EthernetHeader) {
    if ethernet.EtherType != 0x800 {
        panic(fmt.Sprintln("Wrong EtherType:", ethernet.EtherType))
    }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("net.cap")
	check(err)

	pcap_file := PcapFileHeader{}
	// little endian byte ordering indicates this is from machine, not network
	err = binary.Read(f, binary.LittleEndian, &pcap_file)
	check(err)
	check_pcap_file_header(pcap_file)

	pcap_packet := PcapPacketHeader{}
	ethernet := EthernetHeader{}
	var n_packets int
	for {
		_, err = f.Seek(int64(pcap_packet.Length), 1)
		check(err)

        // pcap packet
		err = binary.Read(f, binary.LittleEndian, &pcap_packet)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		check_pcap_packet_header(pcap_packet)

        // ethernet
		err = binary.Read(f, binary.BigEndian, &ethernet)
		check_ethernet_header(ethernet)
        fmt.Println(ethernet)
        
        n_packets += 1
		fmt.Println(n_packets, pcap_packet)

        // go back to beginning of pcap packet to make math at the top work. hacky.
        _, err = f.Seek(-14, 1)
        check(err)
	}

	if n_packets != 99 {
		panic(fmt.Sprintln("Wrong number of packets:", n_packets))
	}

	f.Close()
}
