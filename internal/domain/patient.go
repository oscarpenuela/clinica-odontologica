package domain

type Patient struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Domicilio   string `json:"domicilio"`
	DNI         string `json:"dni"`
	FechaDeAlta string `json:"fecha_de_alta"`
}