package cmd

import (
	"fmt"
	"go-cli/utils"
	"log"
	"net"

	"github.com/urfave/cli"
)

func SearchIp(c *cli.Context) {
	host := c.String("host")
	utils.CleanHost(&host)
	
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}