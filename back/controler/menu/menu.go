package menu

import (
	"log"

	bancodedados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	modelmenu "github.com/eajardini/SigaGoPedidos/back/controler/menu/model"
	"github.com/gin-gonic/gin"
)

var (
	mMenu modelmenu.ModelMenuRaiz
	bd    bancodedados.BDCon
)

//SelectItensDoMenu : zz
func SelectItensDoMenu() ([]modelmenu.ItensDaTabelaMenu, error) {

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
	bd.AbreConexao()
	defer bd.FechaConexao()

	err := bd.BD.Select(&itensDaTabelaMenu, sql)

	if err != nil {
		log.Println("[menu | SelectItensDoMenu] ", "Erro ao selecionar os itens de menu "+err.Error())
	}

	return itensDaTabelaMenu, err
}

// MontaEstruturaDoMenu : a partir dos itens inseridos na estrutura ItensDaTabelaMenu
//	eu monto um
func MontaEstruturaDoMenu() (menuMontado []modelmenu.ItemsNivel1) {

	var (
		itensDaTabelaMenu []modelmenu.ItensDaTabelaMenu
		menuLocal         []modelmenu.ItemsNivel1
		menuItemNivel1    modelmenu.ItemsNivel1
		menuItemNivel2    modelmenu.ItemsNivel2
		menuItemNivel3    modelmenu.ItemsNivel3
		posN1, posN2      int
	)

	//Executa a consulta que retorna o menu do sistema
	itensDaTabelaMenu, _ = SelectItensDoMenu()

	posN1 = -1
	posN2 = -1

	for _, itensMenu := range itensDaTabelaMenu {
		// fmt.Println("Nivel", itensMenu.Nivel)
		if itensMenu.Nivel == 1 {
			posN1++

			menuItemNivel1 = modelmenu.ItemsNivel1{
				Label: itensMenu.Descricao,
			}

			menuLocal = append(menuLocal, menuItemNivel1)
			posN2 = -1

		} else if itensMenu.Nivel == 2 {
			posN2++
			menuItemNivel2 = modelmenu.ItemsNivel2{
				Label: itensMenu.Descricao,
			}
			menuLocal[posN1].Items = append(menuLocal[posN1].Items, menuItemNivel2)

		} else if itensMenu.Nivel == 3 {

			menuItemNivel3 = modelmenu.ItemsNivel3{
				Label: itensMenu.Descricao,
				Link:  itensMenu.Link,
			}
			menuLocal[posN1].Items[posN2].Items = append(menuLocal[posN1].Items[posN2].Items, menuItemNivel3)

		}
	}

	menuMontado = menuLocal

	return menuMontado

}

//RetornaEstruturaDoMenu : zz
func RetornaEstruturaDoMenu() (menulocal []modelmenu.ItemsNivel1) {

	menulocal = MontaEstruturaDoMenu()

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
