package modelfuncionarios

//STFuncionarios : zz
type STFuncionarios struct {
	Funcid              int     `json:"funcid"`
	Cpf                 string  `json:"cpf"`
	Rg                  string  `json:"rg"`
	Funcnome            string  `json:"funcnome"`
	Datanasc            string  `json:"datanasc"`
	Funcdatacontratacao string  `json:"funcdatacontratacao"`
	Funcdatadispensa    string  `json:"funcdatadispensa"`
	Cep                 string  `json:"cep"`
	Endereco            string  `json:"endereco"`
	Cidade              string  `json:"cidade"`
	Uf                  string  `json:"uf"`
	Estado              string  `json:"estado"`
	Funcsalario         float32 `json:"funcsalario"`
}
