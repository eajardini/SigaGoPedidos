package main

import (
	bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	"testing"
)

var (
	bd bancoDeDados.BDCon
)

func setaStringDeConexao(t *testing.T) {
	bd.IniciaConexao()
}
