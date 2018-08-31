package manipulador

import "html/template"

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("/home/carlos/go/src/avancado/servidor_web/html/ola.html"))
