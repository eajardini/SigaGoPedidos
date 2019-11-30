package route

import (
	bd "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	menu "github.com/eajardini/SigaGoPedidos/back/controler/menu"

	//	cliente "eajardini/gin/gocrud/controler/cliente"

	"github.com/gin-gonic/gin"
)

var bancodados bd.BDCon

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// fmt.Println(c.Request.Method)
		c.Next()
	}
}

//IniciaServidor : sd
func IniciaServidor() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	bancodados.ConfiguraStringDeConexao("./config/ConfigBancoDados.toml")
	bancodados.IniciaConexao()

	r.GET("/", menu.MenuPrincipal)

	// cli := r.Group("/cliente")
	// {
	// 	cli.GET("", cliente.OlaCliente)
	// 	cli.GET("/selecionatodos", cliente.SelecionaTodosOsCliente)
	// 	cli.GET("/selecionaclientepornome", cliente.SelecionaClientePorNome)
	// 	cli.POST("/insere", cliente.InsereCliente)
	// 	cli.PUT("/atualizacliente", cliente.AtualizaCliente)
	// 	cli.DELETE("/apagacliente", cliente.ApagaCliente)
	// }

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
