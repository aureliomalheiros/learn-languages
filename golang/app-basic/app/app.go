package app

import (
	"github.com/urfave/cli"
	"log"
	"net"
	"fmt"
)
//Generation 
//Return app
func Generation() *cli.App {
	app := cli.NewApp()
	app.Name = "App line command"
	app.Usage = "research IPs and names server in the internet"

	flags := []cli.Flag {
		cli.StringFlag {
			Name: "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands =  []cli.Command{
		{
			Name: "ip", 
			Usage: "Search IPs",
			Flags: flags,
			Action: searchIPs,
		},
		{
			Name: "server",
			Usage: "Search severs in the internet",
			Flags: flags,
			Action: searchServers,

		},
	}
	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) 

	if erro != nil {
		log.Fatal(erro)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
