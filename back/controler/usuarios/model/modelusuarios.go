package modelusuarios

//STUsuarios : zz
type STUsuarios struct {
	Userid        int    `json:"userid"`
	Funcnome      string `json:"funcnome"`
	Login         string `json:"login"`
	Datacriacao   string `json:"datacriacao"`
	Datavalidade  string `json:"datavalidade"`
	Userbloqueado string `json:"userbloqueado"`
}
