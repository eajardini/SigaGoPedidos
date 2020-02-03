package modelfuncionarios

import (
	"database/sql"
	"time"
)

//STFuncionarios : zz
type STFuncionarios struct {
	Funcid              int             `json:"funcid"`
	Cpf                 sql.NullString  `json:"CPFFunc" validate:"required"`
	Rg                  sql.NullString  `json:"RGFunc"`
	Funcnome            sql.NullString  `json:"nomeFunc" validate:"required"`
	Datanasc            time.Time       `json:"DataNascFunc"`
	Funcdatacontratacao time.Time       `json:"DataContratacaoFunc" validate:"required"`
	Funcdatadispensa    time.Time       `json:"DataDispensaFunc"`
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
