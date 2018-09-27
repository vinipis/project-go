package structs

import (
	"fmt"

	"github.com/caarlos0/env"
)

type parameters struct {
	Host     string `env:"h" envDefault:"127.0.0.12"`
	Port     string `env:"p" envDefault:"3306"`
	User     string `env:"u" envDefault:"${HOME}" envExpand:"true"`
	Password string `env:"o"`
}

//ParametersTerminal é a declaração dos parametros aceitaveis na linha de comando
func ParametersTerminal() (valueParameters [4]string, err error) {

	cfg := parameters{}
	fmt.Println(env.Parse(cfg))
	err = env.Parse(&cfg)
	fmt.Println(cfg)
	fmt.Println(env.Parse(cfg))
	switch {
	case cfg.Host != "":
		valueParameters[0] = cfg.Host
		fallthrough
	case cfg.Port != "":
		valueParameters[1] = string(cfg.Port)
		fallthrough
	case cfg.User != "":
		valueParameters[2] = cfg.User
		fallthrough
	case cfg.Password != "":
		valueParameters[3] = cfg.Password
	}
	return valueParameters, err
}
