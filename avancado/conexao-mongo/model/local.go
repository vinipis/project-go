package model

//Local armazena os dados da localidade pelo seu codigo telefonico
type Local struct {
	Pais             string `json:"Pais" db:"pais" bson:"pais"`
	Cidade           string `json:"Cidade" db:"cidade" bson:"cidade"`
	CodigoTelefonico int    `json:"Cod_Telefone" db:"telcode" bson:"telcode"`
}
