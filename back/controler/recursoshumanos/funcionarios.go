package recursoshumanos

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"io"
	"log"
	_ "log"
	"net/http"
	"strconv"
	"strings"

	"time"

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

	// sql := `
	// SELECT Funcid,  Cpf, Rg, Funcnome, to_char(Datanasc, 'DD/MM/YYYY') Datanasc,
	// 		 to_char(Funcdatacontratacao, 'DD/MM/YYYY') Funcdatacontratacao,
	//      CASE WHEN  Funcdatadispensa isnull THEN '-'
	//           ELSE to_char(Funcdatadispensa, 'DD/MM/YYYY')
	//       END AS Funcdatadispensa
	// from rh_funcionarios
	// `

	sql := `
	SELECT Funcid,  Cpf, Rg, Funcnome, 
	CASE WHEN  Datanasc = '01/01/0001' THEN null ELSE to_char(Datanasc, 'DD/MM/YYYY') end Datanasc ,
	to_char(Funcdatacontratacao, 'DD/MM/YYYY')  Funcdatacontratacao,
	to_char(Funcdatadispensa, 'DD/MM/YYYY')   Funcdatadispensa
  from rh_funcionarios
`
	err := bd.BD.Select(&mFuncionarios, sql)

	if err != nil {
		log.Println("[*** funcionarios | selecionaTodosUsuarios] ", "Erro ao listar todos os Funcionarios: "+err.Error())
	}
	//log.Println("[funcionarios | selecionaTodosUsuarios] ", "Valor da Data Nasci: "+mFuncionarios[0].Datanasc.Time.String())

	return mFuncionarios
}

//ListaTodosFuncionarios : zz
func ListaTodosFuncionarios(c *gin.Context) {

	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  http://localhost:8081/

	//TodosFuncionarios := selecionaTodosFuncionarios()

	type retornoFunc struct {
		Funcid              sql.NullString `json:"funcid"`
		Cpf                 sql.NullString `json:"cpf"`
		Rg                  sql.NullString `json:"rg"`
		Funcnome            sql.NullString `json:"funcnome"`
		Datanasc            sql.NullString `json:"datanasc"`
		Funcdatacontratacao sql.NullString `json:"funcdatacontratacao"`
		Funcdatadispensa    sql.NullString `json:"funcdatadispensa"`
		Foto                []byte         `json:"foto"`
	}

	// var mFuncionarios []modelfuncionarios.STFuncionarios
	var retornoFuncionarios []retornoFunc

	bd.AbreConexao()
	defer bd.FechaConexao()

	sql := `
	SELECT Funcid,  Cpf, Rg, Funcnome, 
	CASE WHEN  Datanasc = '01/01/0001' THEN null ELSE  to_char(Datanasc, 'DD/MM/YYYY')    end  Datanasc,
	CASE WHEN  Funcdatacontratacao = '01/01/0001' THEN null ELSE  to_char(Funcdatacontratacao, 'DD/MM/YYYY')    end  Funcdatacontratacao,
	CASE WHEN  Funcdatadispensa = '01/01/0001' THEN null ELSE  to_char(Funcdatadispensa, 'DD/MM/YYYY')    end  Funcdatadispensa,
	Foto
  from rh_funcionarios
 `
	err := bd.BD.Select(&retornoFuncionarios, sql)

	if err != nil {
		log.Println("[funcionarios | ListaTodosFuncionarios] ", "Erro ao listar todos os Funcionarios: "+err.Error())
	}

	for i := range retornoFuncionarios {
		if retornoFuncionarios[i].Rg.String == "null" {
			retornoFuncionarios[i].Rg.String = ""
		}
		// retornoFuncionarios[i].Datanasc.Time.= retornoFuncionarios[i].Datanasc.Time.Format("02/01/2006")
	}

	c.JSON(200, gin.H{
		"resposta": retornoFuncionarios,
	})

}

// RetornaFotoFuncionario : zz
func RetornaFotoFuncionario(c *gin.Context) {
	type retornoFoto struct {
		Foto []byte `json:"foto"`
	}

	// idFuncionario := c.Request.FormValue("idfuncionario")
	// idFuncionario64, _ := strconv.Atoi(idFuncionario)
	// idFuncionario64 := 104
	idFuncionario := c.Param("idFuncionario")

	idFuncionario64, _ := strconv.Atoi(idFuncionario)
	log.Println("[funcionarios | idFuncionario64] ", "Valor de idFuncionario64: "+strconv.Itoa(idFuncionario64))

	var retornoFuncionarios retornoFoto

	bd.AbreConexao()
	defer bd.FechaConexao()

	sql := `
		SELECT foto
		from rh_funcionarios
		where funcid = ` + idFuncionario

	err := bd.BD.Get(&retornoFuncionarios, sql)

	if err != nil {
		log.Println("[funcionarios | RetornaFotoFuncionario] ", "Erro ao encotrar a foto do Funcionario: "+err.Error())
		c.String(http.StatusInternalServerError, "Erro ao encotrar a foto do Funcionario: "+err.Error())
		return
	}

	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, retornoFuncionarios.Foto); err != nil {
	// 	log.Println("[funcionarios | UPLoadFotoFuncionario] ", "Erro ao converter a Foto do funcionário: "+err.Error())
	// }

	c.Data(http.StatusOK, "image/jpeg", retornoFuncionarios.Foto) // buf.Bytes())

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

//CastNil : zz
func CastNil(dataConvertida sql.NullTime) driver.Value {
	log.Println("[funcionarios | CastNil**] ", "Valor de Data valid: "+strconv.FormatBool(dataConvertida.Valid))
	log.Println("[funcionarios | CastNil**] ", "Valor de Data Format: "+dataConvertida.Time.Format("02/01/2006"))
	if dataConvertida.Time.Format("02/01/2006") == "01/01/0001" {
		return nil
	}
	return dataConvertida.Value
}

//CadastroFuncionario : zz
func CadastroFuncionario(c *gin.Context) {

	var (
		fotoByte *bytes.Buffer
		// Mensagemretorno string 						 `json:"mensagemretorno"`
		MensagemRetorno string
		err             error
	)

	// curl -X POST http://localhost:8081/rh/upLoadFotoFuncionario \
	// -F "foto=@/home/eajardini/pessoal/Fotos/FotoEvandro.png" \
	// -H "Content-Type: multipart/form-data"
	nomeFunc := c.Request.FormValue("nomeFunc")
	CPFFunc := c.Request.FormValue("CPFFunc")
	RGFunc := c.Request.FormValue("RGFunc")
	CEPFunc := c.Request.FormValue("CEPFunc")
	EnderFunc := c.Request.FormValue("EnderFunc")
	CidadeFunc := c.Request.FormValue("CidadeFunc")
	UFFunc := c.Request.FormValue("UFFunc")
	EstadoFunc := c.Request.FormValue("EstadoFunc")
	DataNascFunc := c.Request.FormValue("DataNascFunc")
	DataContratacaoFunc := c.Request.FormValue("DataContratacaoFunc")
	DataDispensaFunc := c.Request.FormValue("DataDispensaFunc")
	SalarioFunc := c.Request.FormValue("SalarioFunc")
	fotoFuncionario, _, errFoto := c.Request.FormFile("foto")

	CPFFunc = strings.ReplaceAll(CPFFunc, ".", "")
	CPFFunc = strings.ReplaceAll(CPFFunc, "-", "")
	CEPFunc = strings.ReplaceAll(CEPFunc, "-", "")

	DataNascFuncConvertida, _ := time.Parse("02/01/2006", DataNascFunc) //DD-MM-YYYY

	if err != nil {
		log.Println("[funcionarios | CadastraFuncionario**] ", "Erro r de Data Nascimento Convertida: "+err.Error())
	}
	dataContratConvertida, _ := time.Parse("02/01/2006", DataContratacaoFunc)   //DD-MM-YYYY
	dataDispensaFuncConvertida, _ := time.Parse("02/01/2006", DataDispensaFunc) //DD-MM-YYYY

	salarioFuncionarioConvertido, _ := strconv.ParseFloat(SalarioFunc, 32)

	if errFoto != nil {
		fotoByte = bytes.NewBuffer(nil)
	} else {
		fotoByte = bytes.NewBuffer(nil)
		_, err := io.Copy(fotoByte, fotoFuncionario)

		if err != nil {
			log.Println("[**[ERRO]funcionarios | CadastraFuncionario**] ", "Valor de converter a foto para Byte: "+err.Error())
		}
	}

	bd.AbreConexao()
	defer bd.FechaConexao()

	tx := bd.BD.MustBegin()

	sql := `
					INSERT INTO rh_funcionarios (funcid, cpf, rg, funcnome, datanasc, funcdatacontratacao, 
																			 funcdatadispensa, foto, funcsalario, cep,
																			 endereco, cidade, uf, estado) 
					VALUES (nextval('seq_rhfuncionarios'), $1, $2, $3, to_timestamp($4,'DD/MM/YYYY'), $5, $6, $7, $8, $9,
																								$10, $11, $12, $13)
					`
	_, err = tx.Exec(sql, CPFFunc, RGFunc, nomeFunc, DataNascFuncConvertida.Format("02/01/2006"), dataContratConvertida,
		dataDispensaFuncConvertida, fotoByte.Bytes(), salarioFuncionarioConvertido,
		CEPFunc, EnderFunc, CidadeFunc, UFFunc, EstadoFunc)

	if err != nil {
		log.Println("[**[ERRO]funcionarios | CadastraFuncionario**] ", "Erros ao inserir: "+err.Error())
		MensagemRetorno = "Erro ao cadastrar funcionário:" + err.Error()
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Println("[**[ERRO]funcionarios | CadastraFuncionario**] ", "Erros ao fechar transação: "+err.Error())
		MensagemRetorno = "Erro ao fechar transação no cadastro funcionário:" + err.Error()
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

	// Upload the file to specific dst.
	//c.SaveUploadedFile(foto, "./testes/"+foto.Filename+"_Foto")

	//criando arquivo Byte em memoria para retornar a imagem
	MensagemRetorno = "Funcionário cadastrado com sucesso!"
	c.String(http.StatusOK, MensagemRetorno)
}
