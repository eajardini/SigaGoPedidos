package recursoshumanos

import (
	"bytes"
	"database/sql"
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
	"github.com/go-playground/validator/v10"
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

//CadastroFuncionario : zz
func CadastroFuncionario(c *gin.Context) {

	var (
		fotoByte *bytes.Buffer
		// Mensagemretorno string 						 `json:"mensagemretorno"`
		MensagemRetorno string
		err             error
		validate        = validator.New()
	)

	// curl -X POST http://localhost:8081/rh/upLoadFotoFuncionario \
	// -F "foto=@/home/eajardini/pessoal/Fotos/FotoEvandro.png" \
	// -H "Content-Type: multipart/form-data"
	nomeFunc := c.Request.FormValue("nomeFunc")
	err = validate.Var(nomeFunc, "required")
	if err != nil {
		MensagemRetorno = "Erro de Validação do Nome do Funcionário: " + err.Error()
		log.Println("[funcionarios | CadastraFuncionario**] ", MensagemRetorno) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		c.JSON(500, gin.H{
			"resposta": MensagemRetorno,
		})
		return
	}

	CPFFunc := c.Request.FormValue("CPFFunc")
	err = validate.Var(CPFFunc, "required")
	if err != nil {
		MensagemRetorno = "Erro de Validação do CPF: " + err.Error()
		log.Println("[funcionarios | CadastraFuncionario**] ", MensagemRetorno) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		// c.String(http.StatusInternalServerError, MensagemRetorno)
		c.JSON(400, gin.H{
			"resposta": MensagemRetorno,
		})
		return
	}

	RGFunc := c.Request.FormValue("RGFunc")
	CEPFunc := c.Request.FormValue("CEPFunc")
	EnderFunc := c.Request.FormValue("EnderFunc")
	CidadeFunc := c.Request.FormValue("CidadeFunc")
	UFFunc := c.Request.FormValue("UFFunc")
	EstadoFunc := c.Request.FormValue("EstadoFunc")
	DataNascFunc := c.Request.FormValue("DataNascFunc")
	DataContratacaoFunc := c.Request.FormValue("DataContratacaoFunc")
	err = validate.Var(DataContratacaoFunc, "required")

	if err != nil {
		MensagemRetorno = "Erro de Validação da Data de Contratação: " + err.Error()
		log.Println("[funcionarios | CadastraFuncionario**] ", MensagemRetorno) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

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

	log.Println("[funcionarios | CadastraFuncionario**] ", "Salario Incial: "+SalarioFunc)
	// SalarioFunc = strings.ReplaceAll(SalarioFunc, ".", "")
	// log.Println("[funcionarios | CadastraFuncionario**] ", "Salario sem ponto: "+SalarioFunc)
	SalarioFunc = strings.ReplaceAll(SalarioFunc, ",", ".")
	log.Println("[funcionarios | CadastraFuncionario**] ", "Salario sem virgula: "+SalarioFunc)
	salarioFuncionarioConvertido, _ := strconv.ParseFloat(SalarioFunc, 32)
	log.Println("[funcionarios | CadastraFuncionario**] ", "Salario ParseFloat: "+strconv.FormatFloat(salarioFuncionarioConvertido, 'f', 2, 64))

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

//BuscaDadosFuncionariosParaAtualizar : zz
func BuscaDadosFuncionariosParaAtualizar(c *gin.Context) {

	var (
		MensagemRetorno     string
		RetornoFuncionarios modelfuncionarios.STFuncionariosParaRetorno
	)

	idFuncionario := c.Request.FormValue("funcid")

	log.Println("[funcionarios | BuscaDadosFuncionariosParaAtualizar] idFuncionario: " + idFuncionario)

	bd.AbreConexao()
	defer bd.FechaConexao()

	sql := `
			select * from rh_funcionarios
			where funcid = ` + idFuncionario

	err := bd.BD.Get(&RetornoFuncionarios, sql)

	if err != nil {
		MensagemRetorno = "Erro ao buscar os dados do funcionário: " + err.Error()
		log.Println("[funcionarios | BuscaDadosFuncionariosParaAtualizar] ", MensagemRetorno) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

	if RetornoFuncionarios.Endereco.String == "null" {
		RetornoFuncionarios.Endereco.String = ""
	}
	if RetornoFuncionarios.Cidade.String == "null" {
		RetornoFuncionarios.Cidade.String = ""
	}
	if RetornoFuncionarios.Estado.String == "null" {
		RetornoFuncionarios.Estado.String = ""
	}

	if RetornoFuncionarios.Uf.String == "null" {
		RetornoFuncionarios.Uf.String = ""
	}

	if RetornoFuncionarios.Rg.String == "null" {
		RetornoFuncionarios.Rg.String = ""
	}

	log.Println("[funcionarios | BuscaDadosFuncionariosParaAtualizar] Salario: " + strconv.FormatFloat(RetornoFuncionarios.Funcsalario.Float64, 'f', 2, 64))

	//Usado para converter o formato dos campos datas
	if strings.HasPrefix(RetornoFuncionarios.Datanasc.String, "0001") {
		RetornoFuncionarios.Datanasc.String = ""
	} else {
		dataNasc, _ := time.Parse("2006-01-02T15:04:05Z", RetornoFuncionarios.Datanasc.String)
		RetornoFuncionarios.Datanasc.String = dataNasc.Format("02/01/2006")
	}

	if strings.HasPrefix(RetornoFuncionarios.Funcdatacontratacao.String, "0001") {
		RetornoFuncionarios.Funcdatacontratacao.String = ""
	} else {
		datacontratacao, _ := time.Parse("2006-01-02T15:04:05Z", RetornoFuncionarios.Funcdatacontratacao.String)
		RetornoFuncionarios.Funcdatacontratacao.String = datacontratacao.Format("02/01/2006")
	}

	if strings.HasPrefix(RetornoFuncionarios.Funcdatadispensa.String, "0001") {
		RetornoFuncionarios.Funcdatadispensa.String = ""
	} else {
		datadispensa, _ := time.Parse("2006-01-02T15:04:05Z", RetornoFuncionarios.Funcdatadispensa.String)
		RetornoFuncionarios.Funcdatadispensa.String = datadispensa.Format("02/01/2006")
	}

	c.JSON(200, RetornoFuncionarios)

}

//AtualizaFuncionarios : zz
func AtualizaFuncionarios(c *gin.Context) {

	var (
		Funcionarios modelfuncionarios.STFuncionarios
		fotoByte     *bytes.Buffer
		// Mensagemretorno string 						 `json:"mensagemretorno"`
		MensagemRetorno string
		err             error
		validate        = validator.New()
	)

	// curl -X POST http://localhost:8081/rh/upLoadFotoFuncionario \
	// -F "foto=@/home/eajardini/pessoal/Fotos/FotoEvandro.png" \
	// -H "Content-Type: multipart/form-data"
	Funcionarios.Funcid, _ = strconv.Atoi(c.Request.FormValue("funcid"))
	Funcionarios.Funcnome.String = c.Request.FormValue("nomeFunc")
	Funcionarios.Cpf.String = c.Request.FormValue("CPFFunc")
	Funcionarios.Rg.String = c.Request.FormValue("RGFunc")
	Funcionarios.Cep.String = c.Request.FormValue("CEPFunc")
	Funcionarios.Endereco.String = c.Request.FormValue("EnderFunc")
	Funcionarios.Cidade.String = c.Request.FormValue("CidadeFunc")
	Funcionarios.Uf.String = c.Request.FormValue("UFFunc")
	Funcionarios.Estado.String = c.Request.FormValue("EstadoFunc")
	Funcionarios.Datanasc.Time, _ = time.Parse("02/01/2006", c.Request.FormValue("DataNascFunc"))
	Funcionarios.Funcdatacontratacao.Time, err = time.Parse("02/01/2006", c.Request.FormValue("DataContratacaoFunc"))
	if err != nil {
		MensagemRetorno = "[**[ERRO]funcionarios | AtualizaFuncionario**] Data de contratação inválida: " + err.Error()
		log.Println(MensagemRetorno)
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return

	}
	Funcionarios.Funcdatadispensa.Time, _ = time.Parse("02/01/2006", c.Request.FormValue("DataDispensaFunc"))
	Funcionarios.Funcsalario.Float64, _ = strconv.ParseFloat(c.Request.FormValue("SalarioFunc"), 32)
	log.Printf("[funcionarios | AtualizaFuncionario**] Salario ParseFloat: %v\n", Funcionarios.Funcsalario.Float64)
	fotoFuncionario, _, errFoto := c.Request.FormFile("foto")
	if errFoto != nil {
		fotoByte = bytes.NewBuffer(nil)
	} else {
		fotoByte = bytes.NewBuffer(nil)
		_, err := io.Copy(fotoByte, fotoFuncionario)
		if err != nil {
			MensagemRetorno = "[**[ERRO]funcionarios | AtualizaFuncionario**] Erro ao converter a foto para Byte: " + err.Error()
			log.Println(MensagemRetorno)
			c.String(http.StatusInternalServerError, MensagemRetorno)
			return
		}
	}
	Funcionarios.Foto = fotoByte.Bytes()

	Funcionarios.Cpf.String = strings.ReplaceAll(Funcionarios.Cpf.String, ".", "")
	Funcionarios.Cpf.String = strings.ReplaceAll(Funcionarios.Cpf.String, "-", "")
	Funcionarios.Cep.String = strings.ReplaceAll(Funcionarios.Cep.String, "-", "")

	//Validação de todos os campos
	validate = validator.New()
	err = validate.Struct(Funcionarios)
	if err != nil {
		MensagemRetorno = "[**[ERRO]funcionarios | AtualizaFuncionario**] Erro ao validar todos os dados: " + err.Error()
		log.Println(MensagemRetorno)
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

	bd.AbreConexao()
	defer bd.FechaConexao()

	tx := bd.BD.MustBegin()

	sql := `
				update rh_funcionarios
				set cpf = $2,
						rg =  $3,
						funcnome = $4,
						datanasc = $5,
						funcdatacontratacao = $6,
						funcdatadispensa = $7,
						foto = $8,
						funcsalario = $9,
						cep = $10,
						endereco = $11,
						cidade = $12,
						uf = $13,
						estado = $14
				where funcid = $1	
			`

	_, err = tx.Exec(sql, Funcionarios.Funcid, Funcionarios.Cpf.String, Funcionarios.Rg.String,
		Funcionarios.Funcnome.String, Funcionarios.Datanasc.Time.Format("02/01/2006"),
		Funcionarios.Funcdatacontratacao.Time.Format("02/01/2006"), Funcionarios.Funcdatadispensa.Time.Format("02/01/2006"),
		Funcionarios.Foto, Funcionarios.Funcsalario.Float64, Funcionarios.Cep.String, Funcionarios.Endereco.String,
		Funcionarios.Cidade.String, Funcionarios.Uf.String, Funcionarios.Estado.String)

	if err != nil {
		MensagemRetorno = "Erro ao atualizar os dados dos funcionário no BD:" + err.Error()
		log.Println(MensagemRetorno)
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}

	err = tx.Commit()

	if err != nil {
		MensagemRetorno = "Erro ao fechar transação na atualização funcionário:" + err.Error()
		log.Println(MensagemRetorno)
		c.String(http.StatusInternalServerError, MensagemRetorno)
		return
	}
	MensagemRetorno = "Funcionário atualizado com sucesso!"
	log.Println(MensagemRetorno)
	c.String(http.StatusOK, MensagemRetorno)
}
