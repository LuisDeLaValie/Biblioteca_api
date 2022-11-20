package model

type ErrorRes struct {
	Error   string `json:"error,omitempty"`
	Mensaje string `json:"mensaje,omitempty"`
	Cuerpo  error  `json:"cuerpo,omitempty"`
}
