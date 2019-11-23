package testes

import (
	"testing"

	bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
)

var (
	bd bancoDeDados.BDCon
)

// https://github.com/stretchr/testify
func TestSetaStringDeConexao(t *testing.T) {
	//assert.Equal(t, 1, 2, "Isto é verdadeiro!")
	bd.IniciaConexao()

}
