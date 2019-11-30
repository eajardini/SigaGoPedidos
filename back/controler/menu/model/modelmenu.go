package modelmenu

var ()

// ModelMenu : zz
type ModelMenu struct {
	MMenu []ItemsNivel1
}

//ItemsNivel1 : zz
type ItemsNivel1 struct {
	Label string        `json:"label"`
	Items []ItemsNivel2 `json:"items"`
}

//ItemsNivel2 : zz
type ItemsNivel2 struct {
	Label string        `json:"label"`
	Items []ItemsNivel3 `json:"items"`
}

//ItemsNivel3 : zz
type ItemsNivel3 struct {
	Label string `json:"label"`
}

// ItensDaTabelaMenu : armazena os itens de menu que estão salvos
// na tabela menu. Vou usar essa estrutura para criar o JSON
//contendo os menu que irão aparece no menu principal da aplicação
type ItensDaTabelaMenu struct {
	Descricao string
	Nivel     int
}

// menulocal = []ItemsNivel1{
// 	{
// 		Label: "Financeiro",
// 		Items: []ItemsNivel2{
// 			{
// 				Label: "Contas a Pagar",
// 				Items: []ItemsNivel3{
// 					{Label: "Cadastro"},
// 				},
// 			},
// 		},
// 	},
// 	{
// 		Label: "CRM",
// 		Items: []ItemsNivel2{
// 			{
// 				Label: "Clientes",
// 				Items: []ItemsNivel3{
// 					{Label: "Cadastro"},
// 					{Label: "Relatórios"},
// 				},
// 			},
// 		},
// 	},
// }
