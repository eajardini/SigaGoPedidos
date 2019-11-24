package testes

import (
	"testing"

	bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	"github.com/stretchr/testify/assert"
)

var (
	bd bancoDeDados.BDCon
)

// https://github.com/stretchr/testify
func TestSetaStringDeConexao(t *testing.T) {
	assert.Equal(t, 1, 1, "Isto é verdadeiro!")
	bd.IniciaConexao()

}
