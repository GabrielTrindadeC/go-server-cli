package cmd

import (
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/urfave/cli"
)
func SearchIp (c *cli.Context ) {
	host := c.String("host")
	fmt.Println(host)
	if u, err := url.Parse(host); err == nil && u.Host != "" {
		hostname, _, err := net.SplitHostPort(u.Host)
		if err == nil {
			fmt.Println(hostname)
			host = hostname
	} else {
			host = u.Host
	}
	}
	ips, err := net.LookupIP(host)
    if err != nil {
        log.Fatal(err)
    }
    for _, ip := range ips {
        fmt.Println(ip)
    }
}