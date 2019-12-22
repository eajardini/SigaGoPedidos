package usuarios

import (
	"log"

	bancodedados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	modelusuarios "github.com/eajardini/SigaGoPedidos/back/controler/usuarios/model"
	"github.com/gin-gonic/gin"
)

var (
	bd bancodedados.BDCon
)

func listaTodosUsuarios() []modelusuarios.STUsuarios {

	var mUsuarios []modelusuarios.STUsuarios

	bd.AbreConexao()
	defer bd.FechaConexao()

	sql := `
	SELECT userid,  funcnome, login, to_char(datacriacao, 'DD/MM/YYYY') datacriacao,
       CASE WHEN  datavalidade isnull THEN '-'         
            ELSE to_char(datavalidade, 'DD/MM/YYYY')
        END AS datavalidade , 
				 userbloqueado
	from acl_user u, rh_funcionarios func
	where u.funcid = func.funcid
	`
	err := bd.BD.Select(&mUsuarios, sql)

	if err != nil {
		log.Println("[usuarios | listaTodosUsuarios] ", "Erro ao listar todos os Usuários "+err.Error())
	}

	return mUsuarios
}

//ListaTodosUsuarios : zz
func ListaTodosUsuarios(c *gin.Context) {

	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  http://localhost:8081/

	TodosUsuarios := listaTodosUsuarios()

	c.JSON(200, gin.H{
		"resposta": TodosUsuarios,
	})

}
