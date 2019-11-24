package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeArquivoScriptSQL(t *testing.T) {
	assert := assert.New(t)
	bd.IniciaConexao()

	sqlArquivo := LeArquivoScriptSQL()
	assert.NotEqual(sqlArquivo, nil, "Não foi possível ler o arquivo de script de SQL!")
	err := bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")
	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")
}
