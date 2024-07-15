package models

import "time"

type UsuarioEditadoReponse struct {
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NovoUsuarioResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	CreatedAt time.Time `json:"createdAt"`
}

type UsuariosResponse struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []usuarios `json:"data"`
}

type usuarios struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}
