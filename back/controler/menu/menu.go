package menu

import (
	"github.com/gin-gonic/gin"
)

// {
// 	"label": "Financeiro",
// 	"items":[
// 			{ "label": "Clientes",
// 				"items" :[
// 										 {"label": "Cadastro de Clientes"}
// 						 ]
// 								}
// 							]
// }

// type MenuItem struct {
// 	Label string `json:"label"`
// 	Items struct {
// 		Label string `json:"label"`
// 		Items struct {
// 			Label string `json:"label"`
// 		} `json:"items"`
// 	} `json:"items"`
// }

// type MenuItem struct {
// 	Menu ItemsNivel1 `json:",omitempty"`
// }

type ItemsNivel1 struct {
	Label string        `json:"label"`
	Items []ItemsNivel2 `json:"items"`
}

type ItemsNivel2 struct {
	Label string        `json:"label"`
	Items []ItemsNivel3 `json:"items"`
}
type ItemsNivel3 struct {
	Label string `json:"label"`
}

var menu []ItemsNivel1

// Principal : zz
func Principal(c *gin.Context) {

	// https://stackoverflow.com/questions/42802136/golang-error-missing-type-in-composite-literal

	// https://stackoverflow.com/questions/38965693/missing-type-in-composite-literal-error
	// menu = []MenuItem{
	// 	MenuItem{
	// 		"label": "Financeiro",
	// 		Items: []ItemsNivel2{
	// 			{
	// 				"label": "Contas a Pagar",
	// 				Items: []ItemsNivel3{
	// 					{"label": "Cadastro"},
	// 					{"label": "Relatório"},
	// 					{"label": "Impressão"},
	// 				},
	// 			},
	// 			{
	// 				"label": "Contas a Receber",
	// 				Items: []ItemsNivel3{
	// 					{"label": "Cadastro"},
	// 					{"label": "Relatório"},
	// 					{"label": "Configuração"},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	MenuItem{
	// 		"label": "CRM",
	// 		Items: []ItemsNivel2{
	// 			{
	// 				"label": "Clientes",
	// 				Items: []ItemsNivel3{
	// 					{"label": "Manutenção"},
	// 					{"label": "Relatório"},
	// 					{"label": "Impressão"},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	menu = []ItemsNivel1{
		{
			Label: "Financeiro",
			Items: []ItemsNivel2{
				{
					Label: "Contas a Pagar",
					Items: []ItemsNivel3{
						{Label: "Cadastro"},
					},
				},
			},
		},
		{
			Label: "CRM",
			Items: []ItemsNivel2{
				{
					Label: "Clientes",
					Items: []ItemsNivel3{
						{Label: "Cadastro"},
						{Label: "Relatórios"},
					},
				},
			},
		},
	}
	// MenuItem{
	// 	"label": "CRM",
	// },

	//	menu[0].item[0].Label = "Vendas"

	c.JSON(200, gin.H{
		"resposta": menu,
	})

}
