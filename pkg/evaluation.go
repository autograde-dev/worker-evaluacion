package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jhonM8a/worker-evaluacion/internal/job"
	"github.com/jhonM8a/worker-evaluacion/internal/minio"
	"github.com/jhonM8a/worker-evaluacion/internal/rabittmq"
)

func Evaluate(idEValuation int, nameFileAnswer string, nameFileEvaluation string, nameBucket string, studentJob job.Student, nameFileparametes string) {
	fmt.Printf("idEValuation %d Iniciado\n", idEValuation)
	fmt.Printf("nameFileAnswer %s Iniciado\n", nameFileAnswer)
	fmt.Printf("nameFileEvaluation %s Iniciado\n", nameFileEvaluation)
	fmt.Printf("nameBucket %s Iniciado\n", nameBucket)
	fmt.Printf("nameFileparametes %s Iniciado\n", nameFileparametes)
	isValid := true

	contentFileParametersToCode, err := minio.GetFileFromMinio(nameBucket, nameFileparametes)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de %s: %v", nameFileparametes, err)
		isValid = false
	}

	if contentFileParametersToCode == "" {
		fmt.Println("El contenido del archivo de parametros está vacío")
		isValid = false
	}

	contentFileAnswerCode, err := minio.GetFileFromMinio(nameBucket, nameFileAnswer)
	if err != nil {
		fmt.Errorf("error al obtener el contenido de %s: %v", nameFileAnswer, err)
		isValid = false
	}

	if contentFileAnswerCode == "" {
		fmt.Println("El contenido del archivo de respuestas está vacío")
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
	err = writeToFile(tmpFileName, contentFileAnswerCode)
	if err != nil {
		fmt.Printf("Error al escribir archivo temporal: %v\n", err)
		isValid = false
	}

	// Ejecutar el archivo Python y capturar su salida

	resultAnswer, err := executePythonScript(tmpFileName, contentFileParametersToCode)
	if err != nil {
		fmt.Printf("Error al ejecutar el archivo Python: %v\n", err)
		isValid = false
	}

	fmt.Println("resultAnswer--->" + resultAnswer)
	// Comparamos los contenidos línea por línea
	readerAnswer := bufio.NewScanner(strings.NewReader(resultAnswer))
	readerEvaluation := bufio.NewScanner(strings.NewReader(contentFileEvaluation))

	lineNumber := 1

	if isValid {
		for readerAnswer.Scan() && readerEvaluation.Scan() {
			lineAnswer := readerAnswer.Text()
			lineEvaluation := readerEvaluation.Text()
			fmt.Println(lineAnswer + ":" + lineEvaluation)

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

	}

	if isValid {
		fmt.Println("Los archivos son iguales.")
	} else {
		fmt.Println("Los archivos no son válidos.")
	}

	message := rabittmq.Message{
		IdEvaluation: idEValuation,
		IsValid:      isValid,
		Student: rabittmq.Student{
			IdEstudiante:    studentJob.IdEstudiante,
			PrimerNombre:    studentJob.PrimerNombre,
			SegundoNombre:   studentJob.SegundoNombre,
			PrimerApellido:  studentJob.PrimerApellido,
			SegundoApellido: studentJob.SegundoApellido,
			Correo:          studentJob.Correo,
		},
	}

	rabittmq.Enqueue(message)

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
func executePythonScript(filePath string, paremetersCode string) (string, error) {
	var resultBuilder strings.Builder

	// Divide `paremetersCode` en líneas
	parameters := strings.Split(paremetersCode, "\n")
	for _, param := range parameters {
		param = strings.TrimSpace(param)
		if param == "" {
			continue // Saltar líneas vacías
		}

		fmt.Println("--->" + param)

		// Ejecuta el comando con el parámetro
		cmd := exec.Command("python", filePath, param)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error compilando codigo python")
			return "", err
		}

		fmt.Println("out.String()--->" + out.String())
		output := out.String()
		if strings.HasSuffix(output, "\n") {
			resultBuilder.WriteString(output) // Ya incluye un salto de línea
		} else {
			resultBuilder.WriteString(output + "\n") // Agregar salto de línea si falta
		}
	}

	return resultBuilder.String(), nil
}
