package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Funcao que vai gerar alguma coisa
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplication command in line"
	app.Usage = "Search IPs and names of sites"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidor",
			Usage:  "Busca por nome dos servidores",
			Flags:  flags,
			Action: buscaServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
