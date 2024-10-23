package minio

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/jhonM8a/worker-evaluacion/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
	once        sync.Once
)

func getMinioClient() *minio.Client {

	once.Do(func() {
		conf := config.LoadConfMinio()

		endpoint := conf.Endpoint
		accessKeyID := conf.AccessKeyID
		secretAccessKey := conf.SecretAccessKey
		useSSL := conf.UseSSL

		var err error
		minioClient, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln("Error al crear el cliente de MinIO:", err)
		}
		fmt.Println("Cliente de MinIO inicializado.")
	})
	return minioClient
}

func GetFileFromMinio(bucketName, objectName string) (string, error) {
	client := getMinioClient()

	// Obtener el archivo (objeto) desde MinIO
	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("error al obtener el objeto %s: %v", objectName, err)
	}
	defer object.Close()

	// Leer el contenido del archivo
	var content string
	buffer := make([]byte, 1024)
	for {
		n, err := object.Read(buffer)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("error al leer el archivo: %v", err)
		}
		if n == 0 {
			break
		}
		content += string(buffer[:n])
	}

	return content, nil
}
