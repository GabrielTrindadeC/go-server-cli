package cmd

import (
	"fmt"
	"go-cli/utils"
	"log"
	"net"
	"strings"

	"github.com/urfave/cli"
)

func SearchNs(c *cli.Context) {
	host := c.String("host")
	utils.CleanHost(&host)
	host = strings.TrimPrefix(host, "www.")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}