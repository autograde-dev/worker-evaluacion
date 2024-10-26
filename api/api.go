package api

import (
	job "github.com/jhonM8a/worker-evaluacion/internal/job"
)

func RequestHandler(jobQueue chan job.Job, nameFileEvaluation string, nameFileAnswer string, nameBucket string, idEValuation int, idSudent int, firstName string, secondName string, firstLastName string, secondLastName string, email string, nameFileParameters string) {

	job := job.Job{
		NameFileEvaluation: nameFileEvaluation,
		NameFileAnswer:     nameFileAnswer,
		NameBucket:         nameBucket,
		IDEValuation:       idEValuation,
		NameFileParametes:  nameFileParameters,
		Student: job.Student{
			IdEstudiante:    idSudent,
			PrimerNombre:    firstName,
			SegundoNombre:   secondName,
			PrimerApellido:  firstLastName,
			SegundoApellido: secondLastName,
			Correo:          email,
		},
	}
	jobQueue <- job
}
