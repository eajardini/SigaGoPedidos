package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

//SIGAGoPedidosConfig :zz
type SIGAGoPedidosConfig struct {
	Principal    principalConfig
	BancoDeDados map[string]bancoDeDadosConfig
}

//SIGAGoPedidosPrincipal : zz
type principalConfig struct {
	Modo string
}

//BancoDeDados : zz
type bancoDeDadosConfig struct {
	Host     string
	User     string
	Port     string
	Ssl      string
	Database string
	Password string
}

func main() {
	var SIGAConfig SIGAGoPedidosConfig

	if _, err := toml.DecodeFile("../config/ConfigBancoDados.toml", &SIGAConfig); err != nil {
		fmt.Println("Não deu certo")
		log.Fatal(err)
	}

	fmt.Println("Modo:", SIGAConfig.Principal.Modo)
	for serverName, server := range SIGAConfig.BancoDeDados {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.Host, server.Database)
	}

	fmt.Println("Host:", SIGAConfig.BancoDeDados[SIGAConfig.Principal.Modo].Host)
	fmt.Println("Host:", SIGAConfig.BancoDeDados[SIGAConfig.Principal.Modo].Database)

	fmt.Println("Host:", SIGAConfig.BancoDeDados["Teste"].Host)
	fmt.Println("Host:", SIGAConfig.BancoDeDados["Teste"].Database)

}

// porta 1,20 x 2,085
// bandeira 1,20 x 0,765
// Guarnição 0,12 cm
