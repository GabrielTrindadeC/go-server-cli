package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/urfave/cli"
)

func Ping(c *cli.Context) {
	ip := c.String("ip")
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", "3", ip)
	case "linux", "darwin": 
		cmd = exec.Command("ping", "-c", "3", ip) 
	default:
		return 
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println(string(output))
	
}

