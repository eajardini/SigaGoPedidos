package menu

import (
	"encoding/json"
	"fmt"

	bancodedados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	modelmenu "github.com/eajardini/SigaGoPedidos/back/controler/menu/model"
	"github.com/gin-gonic/gin"
)

var (
	mMenu modelmenu.ModelMenu
	bd    bancodedados.BDCon
)

//SelectItensDoMenu : zz
func SelectItensDoMenu() []modelmenu.ItensDaTabelaMenu {

	var (
		sql               string
		itensDaTabelaMenu []modelmenu.ItensDaTabelaMenu
	)

	sql = `
					WITH RECURSIVE submenus AS (
						SELECT  menuID, descricao, nivel, codigoMenuSuperiorID, CAST(descricao As varchar(1000)) As Itens_Menu
						FROM    menu  
						where codigoMenuSuperiorID is null
						UNION
						SELECT  m.menuID, m.descricao, m.nivel, m.codigoMenuSuperiorID, CAST(s.Itens_Menu || '->' || m.descricao As varchar(1000)) As Itens_Menu
						FROM menu m
						INNER JOIN submenus s ON s.menuID = m.codigoMenuSuperiorID
					) 
					SELECT  descricao,  nivel
						FROM submenus
					ORDER BY Itens_Menu		
				`
	bd.AbreConexao()
	defer bd.FechaConexao()

	bd.BD.Select(&itensDaTabelaMenu, sql)

	return itensDaTabelaMenu
}

// MontaEstruturaDoMenu : a partir dos itens inseridos na estrutura ItensDaTabelaMenu
//	eu monto um
func MontaEstruturaDoMenu() string {

	var (
		itensDaTabelaMenu []modelmenu.ItensDaTabelaMenu
	)

	itensDaTabelaMenu = SelectItensDoMenu()

	for i, itensMenu := range itensDaTabelaMenu {
		fmt.Println("i, itensMenu, nível", i, " ", itensMenu.Descricao, " ", itensMenu.Nivel)

	}

	return "Borodin"

}

//RetornaEstruturaDoMenu : zz
func RetornaEstruturaDoMenu() (menulocal []modelmenu.ItemsNivel1) {

	MontaEstruturaDoMenu()

	itemsDoMenu := `[{"label":"Finanças","items":[{"label":"Contas a Pagar","items":[{"label":"Manutenção"}]}]},{"label":"CRM","items":null}]`
	json.Unmarshal([]byte(itemsDoMenu), &menulocal)

	return menulocal
}

// MenuPrincipal : zz
func MenuPrincipal(c *gin.Context) {

	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  http://localhost:8081/

	menu := RetornaEstruturaDoMenu()

	c.JSON(200, gin.H{
		"resposta": menu,
	})

}
