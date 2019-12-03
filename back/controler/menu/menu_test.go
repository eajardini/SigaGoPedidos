package menu

import (
	"reflect"
	"testing"

	// modelmenu "github.com/eajardini/SigaGoPedidos/back/controler/menu/model"
	modelmenu "github.com/eajardini/SigaGoPedidos/back/controler/menu/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSelectItensDoMenu(t *testing.T) {
	assert := assert.New(t)
	err := bd.ConfiguraStringDeConexao("../../config/ConfigBancoDados.toml")
	assert.Equal(err, nil, "[menu_test|SelectItensDoMenu]: Erro ao CONFIGURA a string de conexão do Banco de Dados")
	err = bd.IniciaConexao()
	assert.Equal(err, nil, "[Erro ao INICIAR o Banco de Dados]")

	itensDaTabelaMenu, err := SelectItensDoMenu()
	assert.Equal(err, nil, "[menu_test|SelectItensDoMenu]: Problema ao executar a função SelectItensDoMenu()")

	assert.Greater(len(itensDaTabelaMenu), 0, "[menu_test|SelectItensDoMenu]: A função SelectItensDoMenu() retornou 0 registros")

}

func TestMontaEstruturaDoMenu(t *testing.T) {
	tests := []struct {
		name            string
		wantMenuMontado []modelmenu.ItemsNivel1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMenuMontado := MontaEstruturaDoMenu(); !reflect.DeepEqual(gotMenuMontado, tt.wantMenuMontado) {
				t.Errorf("MontaEstruturaDoMenu() = %v, want %v", gotMenuMontado, tt.wantMenuMontado)
			}
		})
	}
}

func TestRetornaEstruturaDoMenu(t *testing.T) {
	tests := []struct {
		name          string
		wantMenulocal []modelmenu.ItemsNivel1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMenulocal := RetornaEstruturaDoMenu(); !reflect.DeepEqual(gotMenulocal, tt.wantMenulocal) {
				t.Errorf("RetornaEstruturaDoMenu() = %v, want %v", gotMenulocal, tt.wantMenulocal)
			}
		})
	}
}

func TestMenuPrincipal(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MenuPrincipal(tt.args.c)
		})
	}
}
