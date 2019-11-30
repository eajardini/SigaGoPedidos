package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeArquivoScriptSQL(t *testing.T) {
	assert := assert.New(t)

	// Execução normal de acordo com a configuração no arquivo ConfigBancoDados.toml
	err := bd.ConfiguraStringDeConexao("../config/ConfigBancoDados.toml")
	assert.Equal(err, nil, "[Erro ao CONFIGURA a string de conexão do Banco de Dados]")

	err = bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")

	schemaSQL := LeArquivoScriptSQL()
	assert.NotEqual(schemaSQL, nil, "Não foi possível ler o arquivo de script de SQL!")

	err = bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")

	err = bd.ExecutaMigrate(schemaSQL)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para migrar o Banco de Dados]")

	schemaMENU := LeArquivoMenuSQL()
	err = bd.ExecutaMigrate(schemaMENU)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para CRIAR o MENU no Banco de Dados]")

	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")

	//***
	// Execução forçada no modo Teste
	//***
	log.Println("Realizando testes forçados no banco de dados no modo Teste")
	err = bd.AbreArquivoDeConfiguracaoDoBancoDeDados("../config/ConfigBancoDados.toml")
	assert.Equal(err, nil, "[Erro ao ABRIR o arquivo de Configuração do Banco de Dados]")

	bd.SetaStringDeConexao("Teste")

	err = bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")

	schemaSQL = LeArquivoScriptSQL()
	assert.NotEqual(schemaSQL, nil, "Não foi possível ler o arquivo de script de SQL!")

	err = bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")

	err = bd.ExecutaMigrate(schemaSQL)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para migrar o Banco de Dados]")

	schemaMENU = LeArquivoMenuSQL()
	err = bd.ExecutaMigrate(schemaMENU)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para CRIAR o MENU no Banco de Dados]")

	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")

	//***
	// Execução forçada no modo Prod
	//***
	log.Println("Realizando testes forçados no banco de dados no modo Teste")
	err = bd.AbreArquivoDeConfiguracaoDoBancoDeDados("../config/ConfigBancoDados.toml")
	assert.Equal(err, nil, "[Erro ao ABRIR o arquivo de Configuração do Banco de Dados]")

	bd.SetaStringDeConexao("Prod")

	err = bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")

	schemaSQL = LeArquivoScriptSQL()
	assert.NotEqual(schemaSQL, nil, "Não foi possível ler o arquivo de script de SQL!")

	err = bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")

	err = bd.ExecutaMigrate(schemaSQL)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para migrar o Banco de Dados]")

	schemaMENU = LeArquivoMenuSQL()
	err = bd.ExecutaMigrate(schemaMENU)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para CRIAR o MENU no Banco de Dados]")

	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")

	//***
	// Execução forçada no modo Desenvolvimento
	//***
	log.Println("Realizando testes forçados no banco de dados no modo Teste")
	err = bd.AbreArquivoDeConfiguracaoDoBancoDeDados("../config/ConfigBancoDados.toml")
	assert.Equal(err, nil, "[Erro ao ABRIR o arquivo de Configuração do Banco de Dados]")

	bd.SetaStringDeConexao("Desenvolvimento")

	err = bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")

	schemaSQL = LeArquivoScriptSQL()
	assert.NotEqual(schemaSQL, nil, "Não foi possível ler o arquivo de script de SQL!")

	err = bd.AbreConexao()
	assert.Equal(err, nil, "[Erro ao ABRIR o Banco de Dados]")

	err = bd.ExecutaMigrate(schemaSQL)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para migrar o Banco de Dados]")

	schemaMENU = LeArquivoMenuSQL()
	err = bd.ExecutaMigrate(schemaMENU)
	assert.Equal(err, nil, "[Erro ao EXECUTAR o arquivo de SQL para CRIAR o MENU no Banco de Dados]")

	err = bd.FechaConexao()
	assert.Equal(err, nil, "[Erro ao FECHAR o Banco de Dados]")
}
