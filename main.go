package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// Load AWS configuration with provided credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Set AWS credentials
	cfg.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("<your S3 key>", "<your S3 secret>", ""))

	// Specify the AWS region
	cfg.Region = " <your s3 region>" // Mumbai region

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)
	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))
			return
		}
		// Upload the file to S3 directly without saving it locally
		fileName := header.Filename
		fileKey := fmt.Sprintf("uploads/%s", fileName)

		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("<your bucket name>"), // Replace with your bucket name
			Key:    aws.String(fileKey),
			Body:   file,
		})
		if err != nil {
			log.Fatalf("failed to upload file to S3: %v", err)
		}

		// Generate a public URL for the uploaded file
		publicURL := fmt.Sprintf("https://<your bucket name>.s3.<your s3 region>.amazonaws.com/%s", fileKey)
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully. Public URL: %s", fileName, publicURL))
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server")
	}
}

echo "# S3-file-upload" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Ankit9114/S3-file-upload.git
git push -u origin main