package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jhonM8a/worker-evaluacion/internal/minio"
	"github.com/jhonM8a/worker-evaluacion/internal/rabittmq"
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

	// Guardar el contenido de la respuesta en un archivo Python temporal
	tmpFileName := "/tmp/tmp_answer.py"
	err = writeToFile(tmpFileName, contentFileAnswer)
	if err != nil {
		fmt.Printf("Error al escribir archivo temporal: %v\n", err)
		isValid = false
	}

	// Ejecutar el archivo Python y capturar su salida
	resultAnswer, err := executePythonScript(tmpFileName)
	if err != nil {
		fmt.Printf("Error al ejecutar el archivo Python: %v\n", err)
		isValid = false
	}

	// Comparamos los contenidos línea por línea
	readerAnswer := bufio.NewScanner(strings.NewReader(resultAnswer))
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

		message := rabittmq.Message{
			IdEvaluation: idEValuation,
			IsValid:      isValid,
		}

		rabittmq.Enqueue(message)
	}

}

// Función para escribir contenido a un archivo temporal
func writeToFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Función para ejecutar un script Python y capturar su salida
func executePythonScript(filePath string) (string, error) {
	cmd := exec.Command("python3", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
