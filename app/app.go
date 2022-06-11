package app

import (
	"github.com/urfave/cli"
)

//Funcao que vai gerar alguma coisa
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplication command in line"
	app.Usage = "Search IPs and names of sites"

	return app
}
