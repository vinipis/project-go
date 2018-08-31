package manipulador

import (
	"html/template"
	"os"
)

//ModeloOla armazena os modelos de pagina que serão executados pelos manipuladores
var ModeloOla = template.Must(template.ParseFiles(os.Getenv("HOME") + "/go/src/project-go/avancado/conexao_bd_sql/html/ola.html"))

//ModeloLocal armazena os modelos de pagina Local
var ModeloLocal = template.Must(template.ParseFiles(os.Getenv("HOME") + "/go/src/project-go/avancado/conexao_bd_sql/html/local.html"))
