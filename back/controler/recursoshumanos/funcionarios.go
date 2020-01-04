package recursoshumanos

import (
	"bytes"
	"io"
	"log"
	_ "log"
	"net/http"

	bancodedados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	modelfuncionarios "github.com/eajardini/SigaGoPedidos/back/controler/recursoshumanos/model"
	"github.com/gin-gonic/gin"
)

var (
	bd bancodedados.BDCon
)

func selecionaTodosFuncionarios() []modelfuncionarios.STFuncionarios {

	var mFuncionarios []modelfuncionarios.STFuncionarios

	bd.AbreConexao()
	defer bd.FechaConexao()

	sql := `
	SELECT Funcid,  Cpf, Rg, Funcnome, to_char(Datanasc, 'DD/MM/YYYY') Datanasc,
			 to_char(Funcdatacontratacao, 'DD/MM/YYYY') Funcdatacontratacao,
       CASE WHEN  Funcdatadispensa isnull THEN '-'         
            ELSE to_char(Funcdatadispensa, 'DD/MM/YYYY')
        END AS Funcdatadispensa 
	from rh_funcionarios
	`
	err := bd.BD.Select(&mFuncionarios, sql)

	if err != nil {
		log.Println("[funcionarios | selecionaTodosUsuarios] ", "Erro ao listar todos os Funcionarios: "+err.Error())
	}

	return mFuncionarios
}

//ListaTodosFuncionarios : zz
func ListaTodosFuncionarios(c *gin.Context) {

	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  http://localhost:8081/

	TodosFuncionarios := selecionaTodosFuncionarios()

	c.JSON(200, gin.H{
		"resposta": TodosFuncionarios,
	})

}

//UPLoadFotoFuncionario : zz
func UPLoadFotoFuncionario(c *gin.Context) {

	// curl -X POST http://localhost:8081/rh/upLoadFotoFuncionario \
	// -F "foto=@/home/eajardini/pessoal/Fotos/FotoEvandro.png" \
	// -H "Content-Type: multipart/form-data"
	fotoretorno, _, _ := c.Request.FormFile("foto")

	// Upload the file to specific dst.
	//c.SaveUploadedFile(foto, "./testes/"+foto.Filename+"_Foto")

	//criando arquivo Byte em memoria para retornar a imagem
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, fotoretorno); err != nil {
		log.Println("[funcionarios | UPLoadFotoFuncionario] ", "Erro ao converter a Foto do funcionário: "+err.Error())
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())

}

//CadastroFuncionario : zz
func CadastroFuncionario(c *gin.Context) {

	// curl -X POST http://localhost:8081/rh/upLoadFotoFuncionario \
	// -F "foto=@/home/eajardini/pessoal/Fotos/FotoEvandro.png" \
	// -H "Content-Type: multipart/form-data"
	nomeFunc := c.Request.FormValue("nomeFunc")
	// CPFFunc := c.Request.FormValue("CPFFunc")
	// RGFunc := c.Request.FormValue("RGFunc")
	// CEPFunc := c.Request.FormValue("CEPFunc")
	// EnderFunc := c.Request.FormValue("EnderFunc")
	// CidadeFunc := c.Request.FormValue("CidadeFunc")
	// UFFunc := c.Request.FormValue("UFFunc")
	// EstadoFunc := c.Request.FormValue("EstadoFunc")
	// fotoFuncionario := c.Request.FormValue("fotoFuncionario")
	// SalarioFunc := c.Request.FormValue("SalarioFunc")
	// fotofunc, _, _ := c.Request.FormFile("foto")

	log.Println("[**funcionarios | CadastraFuncionario**] ", "Valor dos Dados do funcionário: "+nomeFunc)

	// Upload the file to specific dst.
	//c.SaveUploadedFile(foto, "./testes/"+foto.Filename+"_Foto")

	//criando arquivo Byte em memoria para retornar a imagem

	c.String(http.StatusOK, "Valor Retornado pelo Gin:"+nomeFunc)

}
