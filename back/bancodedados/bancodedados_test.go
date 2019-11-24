package bancodedados

import (
	"testing"

	// bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	"github.com/stretchr/testify/assert"
)

var (
	bd BDCon
)

// https://github.com/stretchr/testify
func TestSetaStringDeConexao(t *testing.T) {
	assert := assert.New(t)

	err := bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")
	err = bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")
	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")

}

// func main() {
// 	fmt.Println("Testando...")
// }
