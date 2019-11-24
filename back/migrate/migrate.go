package main

import (
	"io/ioutil"
	"log"

	bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
)

var (
	bd bancoDeDados.BDCon
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

type bancoDeDadosConfig struct {
	SGBD     string
	Host     string
	User     string
	Port     string
	Ssl      string
	Database string
	Password string
}

// LeArquivoScriptSQL :zz
func LeArquivoScriptSQL() []byte {
	sqlArquivo, err := ioutil.ReadFile("scriptBaseDadosSigaGoPedidos.sql")
	if err != nil {
		log.Fatal(err)
	}

	return sqlArquivo
}

func main() {
	bd.IniciaConexao()

	schemaSQL := LeArquivoScriptSQL()

	bd.AbreConexao()
	bd.ExecutaMigrate(schemaSQL)
	bd.FechaConexao()

}

// porta 1,20 x 2,085
// bandeira 1,20 x 0,765
// Guarnição 0,12 cm
