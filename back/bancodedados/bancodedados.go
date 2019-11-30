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

// AbreArquivoDeConfiguracaoDoBancoDeDados : abre o arquivo ConfigBancoDados.toml e seta a
// estrutura SIGAGoPedidosConfig.
func (bdcon *BDCon) AbreArquivoDeConfiguracaoDoBancoDeDados(caminhoDoArquivoToml string) error {

	var err error

	if _, err = toml.DecodeFile(caminhoDoArquivoToml, &SIGAConfig); err != nil {
		fmt.Println("[bancodedados:setaStringDeConexao] Erro ao abrio o arquivo TOML de configuração")
		log.Fatal(err)
	}
	return err
}

// SetaStringDeConexao : zz
func (bdcon *BDCon) SetaStringDeConexao(modo string) {

	log.Println("[bancodados, SetaStringDeConexao] Modo de Abertura: ", modo)

	sgbd = SIGAConfig.BancoDeDados[modo].SGBD
	stringConexao = fmt.Sprintf("user=%s dbname=%s host=%s sslmode=%s",
		SIGAConfig.BancoDeDados[modo].User, SIGAConfig.BancoDeDados[modo].Database,
		SIGAConfig.BancoDeDados[modo].Host, SIGAConfig.BancoDeDados[modo].Ssl)
	fmt.Println("[string de Coneção:]", stringConexao)
	// fmt.Println("Host:", SIGAConfig.BancoDeDados[SIGAConfig.Principal.Modo].Database)
}

//ConfiguraStringDeConexao : este método abre o arquivo de configuração do banco de dados
//	e depois seta a string de conexão
func (bdcon *BDCon) ConfiguraStringDeConexao(caminhoDoArquivoToml string) error {
	err := bdcon.AbreArquivoDeConfiguracaoDoBancoDeDados(caminhoDoArquivoToml)
	bdcon.SetaStringDeConexao(SIGAConfig.Principal.Modo)

	return err
}

//IniciaConexao : Faz a conexão inicial. A string de conexão já deve estar
//								configurada previamente pelo método ConfiguraStringDeConexao()
func (bdcon *BDCon) IniciaConexao() error {
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
	return err
}

//AbreConexao :zz
func (bdcon *BDCon) AbreConexao() error {
	db, err := sqlx.Connect(sgbd, stringConexao)
	if err != nil {
		log.Fatalln(err)
	}
	//seta a variavel BD para fazer as consultas SQL
	bdcon.BD = db

	return err

}

//FechaConexao :zz
func (bdcon *BDCon) FechaConexao() error {
	err := bdcon.BD.Close()
	return err
}

//Insert : faz o insert no Banco de Dados
func (bdcon *BDCon) Insert(sgbd string, sqlinsert interface{}) {
	fmt.Println(sqlinsert)
}

// ExecutaMigrate : zz
func (bdcon *BDCon) ExecutaMigrate(schemaSQL []byte) error {
	_, err := bdcon.BD.Exec(string(schemaSQL))
	if err != nil {
		log.Println("[bancodedados:ExecutaMigrate] Erro ao realizar o Migrate")
		log.Fatalln(err)
	}

	return err
}
