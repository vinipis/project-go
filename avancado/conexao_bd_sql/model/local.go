package model

import (
	"database/sql"
)

//Local armazena os dados da localidade pelo seu codio telefonico
type Local struct {
	Pais             string         `json:"País" db:"pais"`
	Cidade           sql.NullString `json:"Cidade" db:"cidade"`
	CodigoTelefonico int            `json:"Cod_Telefone" db:"telcode"`
}
