package structs

//Parameters é a declaração dos parametros aceitaveis na linha de comando
type Parameters struct {
	host     string `env:"-h" envDefault:"127.0.0.1"`
	port     string `env:"-p" envDefault:"3306"`
	user     string `env:"-u" envDefault:"${HOME}"`
	password string `env:"-o, required"`
}
