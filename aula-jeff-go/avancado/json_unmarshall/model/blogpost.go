package model

//BlgPost armazena dados de um post no blog
type BlogPost struct {
	UsuarioID int    `json:"userId"`
	ID        int    `json:"id"`
	Titulo    string `json:"title"`
	Texto     string `json:"body"`
}
