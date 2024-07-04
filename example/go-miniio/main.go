package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"os"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

	ctx := context.Background()
	endpoint := "localhost:10002"
	accessKeyID := "kQk1geF89bkWzKeuy8Ns"
	secretAccessKey := "AqZglpzEMnokSxUcxF4iC7rop6DkkWwWd31eVa14"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	var (
		backetName  = "test-miniio"
		fileName    = "test"
		filePath    = "./test"
		contentType = "application/octet-stream"
	)

	exists, err := minioClient.BucketExists(ctx, backetName)
	if err != nil {
		panic(err)
	}
	if exists {
		// Upload file
		info, err := minioClient.FPutObject(ctx, backetName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			panic(err)
		}
		log.Printf("Successfully uploaded：%v of size %d\n", fileName, info.Size)
	}

	os.Mkdir("./download", os.ModePerm)
	// Download file
	err = minioClient.FGetObject(ctx, backetName, fileName, "./download", minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
}
