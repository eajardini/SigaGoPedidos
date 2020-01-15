package modelfuncionarios

import "database/sql"

//STFuncionarios : zz
type STFuncionarios struct {
	Funcid              int             `json:"funcid"`
	Cpf                 sql.NullString  `json:"cpf"`
	Rg                  sql.NullString  `json:"rg"`
	Funcnome            sql.NullString  `json:"funcnome"`
	Datanasc            sql.NullTime    `json:"datanasc"`
	Funcdatacontratacao sql.NullTime    `json:"funcdatacontratacao"`
	Funcdatadispensa    sql.NullTime    `json:"funcdatadispensa"`
	Cep                 sql.NullString  `json:"cep"`
	Endereco            sql.NullString  `json:"endereco"`
	Cidade              sql.NullString  `json:"cidade"`
	Uf                  sql.NullString  `json:"uf"`
	Estado              sql.NullString  `json:"estado"`
	Funcsalario         sql.NullFloat64 `json:"funcsalario"`
	foto                sql.RawBytes
}
