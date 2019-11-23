package bancodedados

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
)

//BDCon : zz
type BDCon struct {
	BD *sqlx.DB
}

//SIGAGoPedidosConfig :zz
type SIGAGoPedidosConfig struct {
	Principal    principalConfig
	BancoDeDados map[string]bancoDeDadosConfig
}

//SIGAGoPedidosPrincipal : zz
type principalConfig struct {
	Modo string
}

//BancoDeDados : zz
type bancoDeDadosConfig struct {
	SGBD     string
	Host     string
	User     string
	Port     string
	Ssl      string
	Database string
	Password string
}

//SIGAConfig : zz
var (
	SIGAConfig    SIGAGoPedidosConfig
	modo          string
	sgbd          string
	stringConexao string
)

func setaStringDeConexao() {
	if _, err := toml.DecodeFile("../config/ConfigBancoDados.toml", &SIGAConfig); err != nil {
		fmt.Println("Não deu certo")
		log.Fatal(err)
	}
	modo = SIGAConfig.Principal.Modo
	sgbd = SIGAConfig.BancoDeDados[modo].SGBD
	stringConexao = fmt.Sprintf("user=%s dbname=%s host=%s sslmode=%s",
		SIGAConfig.BancoDeDados[modo].User, SIGAConfig.BancoDeDados[modo].Database,
		SIGAConfig.BancoDeDados[modo].Host, SIGAConfig.BancoDeDados[modo].Ssl)
	fmt.Println("[string de Coneção:]", stringConexao)
	// fmt.Println("Host:", SIGAConfig.BancoDeDados[SIGAConfig.Principal.Modo].Database)
}

// GetStringConexao : zz
func GetStringConexao() string {
	return stringConexao
}

//IniciaConexao : zz
func (bdcon *BDCon) IniciaConexao() {
	setaStringDeConexao()
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=crudbd host=172.17.0.1  sslmode=disable")
	db, err := sqlx.Connect(sgbd, stringConexao)
	if err != nil {
		log.Fatal("[bancodados, Inicia] Erro ao Conectar ao Banco de Dados!!")
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("[bancodados, Inicia] Erro ao Pingar ao Banco de Dados!!")
	} else {
		log.Println("[bancodados, Inicia] Banco de Dados conectado corretamente!!")
	}

	db.Close()
}

//AbreConexao :zz
func (bdcon *BDCon) AbreConexao() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=crudbd host=172.17.0.1  sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	//seta a variavel BD para fazer as consultas SQL
	bdcon.BD = db

}

//FechaConexao :zz
func (bdcon *BDCon) FechaConexao() {
	bdcon.BD.Close()
}

//Insert : faz o insert no Banco de Dados
func (bdcon *BDCon) Insert(sgbd string, sqlinsert interface{}) {
	fmt.Println(sqlinsert)
}
