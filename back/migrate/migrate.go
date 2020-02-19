package main

import (
	"fmt"
	"io/ioutil"
	"log"

	bancoDeDados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	modelmenu "github.com/eajardini/SigaGoPedidos/back/controler/menu/model"
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

// LeArquivoMenuSQL :zz
func LeArquivoMenuSQL() []byte {
	sqlArquivo, err := ioutil.ReadFile("scriptCriaMenuSIGA.sql")
	if err != nil {
		log.Fatal(err)
	}

	return sqlArquivo
}

// MostraItensDoMenu : zz
func MostraItensDoMenu() error {

	var (
		sql               string
		itensDaTabelaMenu []modelmenu.ItensDaTabelaMenu
	)

	sql = `
								WITH RECURSIVE submenus AS (
									SELECT  menuID, descricao, nivel, link,icone,codigoMenuSuperiorID, CAST(descricao As varchar(1000)) As Itens_Menu
									FROM    menu  
									where codigoMenuSuperiorID = 0
									UNION
									SELECT  m.menuID, m.descricao, m.nivel, m.link, m.icone,m.codigoMenuSuperiorID, CAST(s.Itens_Menu || '->' || m.descricao As varchar(1000)) As Itens_Menu
									FROM menu m
									INNER JOIN submenus s ON s.menuID = m.codigoMenuSuperiorID
							) 
							SELECT  menuID, codigoMenuSuperiorID, descricao, link , icone  ,nivel
								FROM submenus
							ORDER BY Itens_Menu
				`

	err := bd.BD.Select(&itensDaTabelaMenu, sql)

	if err != nil {
		log.Println("[menu | SelectItensDoMenu] ", "Erro ao selecionar os itens de menu "+err.Error())
	}

	for _, itensMenu := range itensDaTabelaMenu {
		fmt.Println("menuID, codigoMenuSuperiorID, descricao, nivel",
			itensMenu.MenuID, " ", itensMenu.CodigoMenuSuperiorID, " ",
			itensMenu.Descricao, " ", itensMenu.Nivel)
	}

	return err
}

func main() {
	bd.ConfiguraStringDeConexao("../config/ConfigBancoDados.toml")
	bd.IniciaConexao()

	schemaSQL := LeArquivoScriptSQL()
	bd.AbreConexao()
	bd.ExecutaMigrate(schemaSQL)
	bd.FechaConexao()

	//Migra as tabelas do Menu
	schemaSQL = LeArquivoMenuSQL()
	bd.AbreConexao()
	bd.ExecutaMigrate(schemaSQL)
	MostraItensDoMenu() //Usado para verificar como ficou a estrutura de Menu

	bd.FechaConexao()

}
