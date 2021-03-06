package route

import (
	bd "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	menu "github.com/eajardini/SigaGoPedidos/back/controler/menu"
	funcionarios "github.com/eajardini/SigaGoPedidos/back/controler/recursoshumanos"
	usuarios "github.com/eajardini/SigaGoPedidos/back/controler/usuarios"

	//	cliente "eajardini/gin/gocrud/controler/cliente"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	r.Static("/fotos", "./fotos")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sigagopedidos", store))

	r.Use(CORSMiddleware())

	bancodados.ConfiguraStringDeConexao("./config/ConfigBancoDados.toml")
	bancodados.IniciaConexao()

	r.GET("/", menu.MenuPrincipal)

	// Módulo de Administração
	r.GET("/listaTodosUsuarios", usuarios.ListaTodosUsuarios)

	// Módulo de Recursos Humanos
	rh := r.Group("/rh")
	{
		rh.GET("/listaTodosFuncionarios", funcionarios.ListaTodosFuncionarios)
		rh.GET("/retornafotofuncionario/:idFuncionario", funcionarios.RetornaFotoFuncionario)
		rh.POST("/atualizaFuncionarios", funcionarios.AtualizaFuncionarios)
		rh.POST("/buscaDadosFuncionariosParaAtualizar", funcionarios.BuscaDadosFuncionariosParaAtualizar)
		rh.POST("/upLoadFotoFuncionario", funcionarios.UPLoadFotoFuncionario)
		rh.POST("/cadastroFuncionario", funcionarios.CadastroFuncionario)

	}

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
