package pkg

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/jhonM8a/worker-evaluacion/internal/minio"
)

func Evaluate(idEValuation int, nameFileAnswer string, nameFileEvaluation string, nameBucket string) {
	fmt.Printf("idEValuation %d Iniciado\n", idEValuation)
	isValid := true

	contentFileAnswer, err := minio.GetFileFromMinio(nameBucket, nameFileAnswer)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de %s: %v", nameFileAnswer, err)
		isValid = false
	}

	if contentFileAnswer == "" {
		fmt.Println("El contenido del archivo de respuesta está vacío")
		isValid = false
	}

	contentFileEvaluation, err := minio.GetFileFromMinio(nameBucket, nameFileEvaluation)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de %s: %v", nameFileEvaluation, err)
		isValid = false
	}

	if contentFileEvaluation == "" {
		fmt.Println("El contenido del archivo de evaluación está vacío")
		isValid = false
	}

	// Comparamos los contenidos línea por línea
	readerAnswer := bufio.NewScanner(strings.NewReader(contentFileAnswer))
	readerEvaluation := bufio.NewScanner(strings.NewReader(contentFileEvaluation))

	lineNumber := 1

	if isValid {
		for readerAnswer.Scan() && readerEvaluation.Scan() {
			lineAnswer := readerAnswer.Text()
			lineEvaluation := readerEvaluation.Text()

			// Si las líneas son diferentes, marcamos la bandera como falsa
			if lineAnswer != lineEvaluation {
				fmt.Printf("Error: la línea %d es diferente entre los archivos\n", lineNumber)
				isValid = false
				break
			}
			lineNumber++
		}

		// Verificamos si los archivos tienen diferente número de líneas
		if readerAnswer.Scan() || readerEvaluation.Scan() {
			fmt.Println("Error: los archivos tienen diferente número de líneas")
			isValid = false
		}

		if isValid {
			fmt.Println("Los archivos son iguales.")
		} else {
			fmt.Println("Los archivos no son válidos.")
		}
	}

}
