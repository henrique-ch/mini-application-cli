package main

import (
	"command-in-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplication := app.Gerar()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
