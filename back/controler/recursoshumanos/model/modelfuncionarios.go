package modelfuncionarios

import (
	"database/sql"
)

//STFuncionarios : zz
type STFuncionarios struct {
	Funcid              int             `json:"funcid"`
	Cpf                 sql.NullString  `json:"CPFFunc" validate:"required"`
	Rg                  sql.NullString  `json:"RGFunc"`
	Funcnome            sql.NullString  `json:"nomeFunc" validate:"required"`
	Datanasc            sql.NullTime    `json:"DataNascFunc"`
	Funcdatacontratacao sql.NullTime    `json:"DataContratacaoFunc" validate:"required"`
	Funcdatadispensa    sql.NullTime    `json:"DataDispensaFunc"`
	Foto                []byte          `json:"fotoFuncionario"`
	Funcsalario         sql.NullFloat64 `json:"money"`
	Cep                 sql.NullString  `json:"CEPFunc"`
	Endereco            sql.NullString  `json:"EnderFunc"`
	Cidade              sql.NullString  `json:"CidadeFunc"`
	Uf                  sql.NullString  `json:"UFFunc"`
	Estado              sql.NullString  `json:"EstadoFunc"`
}

//STFuncionariosParaRetorno : zz
type STFuncionariosParaRetorno struct {
	Funcid              int             `json:"funcid"`
	Cpf                 sql.NullString  `json:"CPFFunc" validate:"required"`
	Rg                  sql.NullString  `json:"RGFunc"`
	Funcnome            sql.NullString  `json:"nomeFunc"`
	Datanasc            sql.NullString  `json:"DataNascFunc"`
	Funcdatacontratacao sql.NullString  `json:"DataContratacaoFunc"`
	Funcdatadispensa    sql.NullString  `json:"DataDispensaFunc"`
	Foto                []byte          `json:"fotoFuncionario"`
	Funcsalario         sql.NullFloat64 `json:"money"`
	Cep                 sql.NullString  `json:"CEPFunc"`
	Endereco            sql.NullString  `json:"EnderFunc"`
	Cidade              sql.NullString  `json:"CidadeFunc"`
	Uf                  sql.NullString  `json:"UFFunc"`
	Estado              sql.NullString  `json:"EstadoFunc"`
}
