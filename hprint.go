package main

import "fmt"

func hprint(pinfo packetInfo, short bool) {

	if short {
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("URL:\t\t%s\n", pinfo.Hreq.URL)
		if pinfo.Hreq.BasicAuth != "" {
			fmt.Printf("Basic-Auth:\t%s\n", pinfo.Hreq.BasicAuth)
		}
		fmt.Printf("Source:\t\t%s:%s (%s)\n", pinfo.SrcIP, pinfo.SrcPort, pinfo.SrcMAC)
		fmt.Printf("Destination:\t%s:%s (%s)\n", pinfo.DstIP, pinfo.DstPort, pinfo.DstMAC)
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("Date:\t\t%s\n", pinfo.Date)
		fmt.Printf("Method:\t\t%s\n", pinfo.Hreq.Method)
		fmt.Printf("URL:\t\t%s\n", pinfo.Hreq.URL)
		if pinfo.Hreq.BasicAuth != "" {
			fmt.Printf("Basic-Auth:\t%s\n", pinfo.Hreq.BasicAuth)
		}
		if pinfo.Hreq.UseProxy {
			fmt.Printf("UseProxy:\t%t\n", pinfo.Hreq.UseProxy)
			fmt.Printf("ProxyAuth:\t%s\n", pinfo.Hreq.ProxyAuth)
		}
		if pinfo.Hreq.Referer != "" {
			fmt.Printf("Referer:\t%s\n", pinfo.Hreq.Referer)
		}
		if pinfo.Hreq.UserAgent != "" {
			fmt.Printf("UserAgent:\t%s\n", pinfo.Hreq.UserAgent)
		}
		fmt.Printf("Source:\t\t%s:%s (%s)\n", pinfo.SrcIP, pinfo.SrcPort, pinfo.SrcMAC)
		fmt.Printf("Destination:\t%s:%s (%s)\n", pinfo.DstIP, pinfo.DstPort, pinfo.DstMAC)
	}
}
