package job

type Student struct {
	IdEstudiante    int
	PrimerNombre    string
	SegundoNombre   string
	PrimerApellido  string
	SegundoApellido string
	Correo          string
}

type Job struct {
	NameFileEvaluation string
	NameFileAnswer     string
	NameFileParametes  string
	NameBucket         string
	IDEValuation       int
	Student            Student
}
