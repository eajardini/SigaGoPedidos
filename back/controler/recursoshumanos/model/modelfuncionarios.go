package modelfuncionarios

//STFuncionarios : zz
type STFuncionarios struct {
	funcid              int    `json:"funcid"`
	cpf                 string `json:"cpf"`
	rg                  string `json:"rg"`
	funcnome            string `json:"funcnome"`
	datanasc            string `json:"datanasc"`
	funcdatacontratacao string `json:"funcdatacontratacao"`
	funcdatadispensa    string `json:"funcdatadispensa"`
}
