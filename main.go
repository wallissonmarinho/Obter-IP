package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wallissonmarinho/Obter-IPs/app"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
