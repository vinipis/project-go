package manipulador

import (
	"fmt"
	"net/http"
	"project-go/avancado/conexao_bd_sql/model"
	"project-go/avancado/conexao_bd_sql/repo"
	"strconv"
	"time"
)

//Local é o manipilador de rota para a página
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}

	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um numero valido", http.StatusBadRequest)
		fmt.Println("[Local] Erro ao converter o numero enviado: ", err.Error())
		return
	}

	sql := "select pais, cidade, telcode from testego where telcode = ?"
	rows, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse numero", http.StatusInternalServerError)
		fmt.Println("[Local] não foi possivel executar a query: ", sql, "erro: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possivel pesquisar esse numero", http.StatusInternalServerError)
			fmt.Println("[Local] não foi possivel fazer o binding dos dados do banco na struct: ", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Local] Erro na execucao do modelo: ", err.Error())
	}

	sql = "insert into logquery (daterequest) values (?)"
	res, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[Local] Erro na inclusão do log: ", sql, " - ", err.Error())
	}

	rowsAfetadas, err := res.RowsAffected()
	if err != nil {
		fmt.Println("[Local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}
	fmt.Println("Sucesso! ", rowsAfetadas, " Linha(s) afetada(s)")
}
