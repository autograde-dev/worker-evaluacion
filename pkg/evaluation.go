package pkg

import (
	"fmt"

	"github.com/jhonM8a/worker-evaluacion/internal/minio"
)

func Evaluate(idEValuation int, nameFileAnswer string, nameFileEvaluation string, nameBucket string) {
	fmt.Printf("idEValuation %d Iniciado\n", idEValuation)

	contentFileAnswer, err := minio.GetFileFromMinio(nameBucket, nameFileAnswer)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de  %s: %v", nameFileAnswer, err)
	}

	if contentFileAnswer != "" {
		fmt.Printf(contentFileAnswer)
	}

	contentFileEvaluation, err := minio.GetFileFromMinio(nameBucket, nameFileEvaluation)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de  %s: %v", nameFileEvaluation, err)
	}

	if contentFileEvaluation != "" {
		fmt.Printf(contentFileEvaluation)
	}

}
