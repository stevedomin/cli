package main

import (
	"fmt"
	"github.com/stevedomin/cli"
)

var (
	force  bool
	port   int
	host   string
	domain string
)

func main() {
	rootCmd := cli.NewCommand("myapp")
	rootCmd.Flags.BoolVar(&force, "force", false, "Force")
	rootCmd.HandlerFunc = func(args []string) {
		fmt.Printf("Handler for myapp, args: %v, force: %t\n", args, force)
	}

	runCmd := cli.NewCommand("run")
	runCmd.Flags.IntVar(&port, "p", 80, "Port")
	runCmd.HandlerFunc = func(args []string) {
		fmt.Printf("Handler for run, args: %v, port: %d\n", args, port)
	}

	serverCmd := cli.NewCommand("server")
	serverCmd.Flags.StringVar(&host, "h", "localhost", "Host")
	serverCmd.HandlerFunc = func(args []string) {
		fmt.Printf("Handler for run server, args: %v, port: %d, host: %s\n", args, port, host)
	}

	clientCmd := cli.NewCommand("client")
	clientCmd.Flags.StringVar(&domain, "d", "myapp.com", "Domain")
	clientCmd.HandlerFunc = func(args []string) {
		fmt.Printf("Handler for run client, args: %v, port: %d, domain: %s\n", args, port, domain)
	}

	rootCmd.AddCommands(runCmd)
	runCmd.AddCommands(serverCmd, clientCmd)
	rootCmd.Execute(nil)
}
