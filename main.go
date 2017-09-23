package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type packetInfo struct {
	Date    string      `json:"Date"`
	Hreq    httpRequest `json:"HTTP"`
	SrcPort string      `json:"SrcPort"`
	DstPort string      `json:"DstPort"`
	SrcIP   string      `json:"SrcIP"`
	DstIP   string      `json:"DstIP"`
	SrcMAC  string      `json:"SrcMAC"`
	DstMAC  string      `json:"DstMAC"`
}

type httpRequest struct {
	URL       string `json:"URL"`
	Method    string `json:"Method"`
	Version   string `json:"Version"`
	BasicAuth string `json:"BasicAuth"`
	UseProxy  bool   `json:"UseProxy"`
	ProxyAuth string `json:"ProxyAuth"`
	Referer   string `json:"Referer"`
	UserAgent string `json:"UserAgent"`
}

var (
	iface    = flag.String("i", "en0", "Interface to get packets from")
	pcapFile = flag.String("p", "", "Filename to read from, overrides -i")
	filter   = flag.String("f", "tcp", "BPF filter for pcap")
	logFile  = flag.String("l", "", "Log file")
	authOnly = flag.Bool("a", false, "Display only information including Basic-Auth")
	short    = flag.Bool("s", false, "Display information in short format")
	snaplen  = flag.Int("S", 1600, "SnapLen for pcap packet capture")
)

func main() {

	var handle *pcap.Handle
	var pinfo packetInfo
	var err error

	flag.Parse()

	file, err := os.OpenFile(*logFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		if *logFile != "" {
			log.Fatal(err)
		}
	} else {
		defer file.Close()
	}

	// Set up pcap packet capture
	if *pcapFile != "" {
		log.Printf("Reading from pcap dump %q", *pcapFile)
		handle, err = pcap.OpenOffline(*pcapFile)
	} else {
		log.Printf("Starting capture on interface %q", *iface)
		handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err)
	}

	log.Println("reading in packets")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	for packet := range packets {
		// A nil packet indicates the end of a pcap file.
		if packet == nil {
			return
		}
		if application := packet.ApplicationLayer(); application != nil {
			payload := string(application.Payload())
			hreq, err := hparser(payload)
			if err != nil {
				continue
			}
			pinfo.Hreq = hreq

			pinfo.Date = time.Now().Format("2006/01/02 15:04:05")

			tcp := packet.TransportLayer().TransportFlow()
			pinfo.SrcPort = tcp.Src().String()
			pinfo.DstPort = tcp.Dst().String()

			net := packet.NetworkLayer().NetworkFlow()
			pinfo.SrcIP = net.Src().String()
			pinfo.DstIP = net.Dst().String()

			link := packet.LinkLayer().LinkFlow()
			pinfo.SrcMAC = link.Src().String()
			pinfo.DstMAC = link.Dst().String()

			b, _ := json.Marshal(&pinfo)
			fmt.Fprintln(file, string(b))

			if *authOnly {
				if pinfo.Hreq.BasicAuth != "" {
					hprint(pinfo, *short)
				}
			} else {
				hprint(pinfo, *short)
			}
		}
	}
}
