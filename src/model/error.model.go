package model

type ErrorRes struct {
	Titulo  string `json:"error,omitempty"`
	Mensaje string `json:"mensaje,omitempty"`
	Cuerpo  error  `json:"cuerpo,omitempty"`
}

func (t ErrorRes) Error() string {

	return t.Titulo + ": " + t.Mensaje + "\n" + t.Cuerpo.Error()
}
