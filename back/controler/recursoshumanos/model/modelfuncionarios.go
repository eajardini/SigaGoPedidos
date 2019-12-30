package modelfuncionarios

//STFuncionarios : zz
type STFuncionarios struct {
	Funcid              int    `json:"funcid"`
	Cpf                 string `json:"cpf"`
	Rg                  string `json:"rg"`
	Funcnome            string `json:"funcnome"`
	Datanasc            string `json:"datanasc"`
	Funcdatacontratacao string `json:"funcdatacontratacao"`
	Funcdatadispensa    string `json:"funcdatadispensa"`
}
