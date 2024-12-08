package logic

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"testing"
)

func TestUpload(T *testing.T) {
	// 加载 .env 文件
	err := godotenv.Load("../../../../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	// 从环境变量获取配置
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET")
	location := os.Getenv("MINIO_REGION")

	// 初始化 MinIO 客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // 如果需要使用 HTTPS，请设置为 true
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 创建存储桶
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 检查桶是否已存在
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// 上传文件
	objectName := "test_avatar.jpg"
	filePath := os.Getenv("TEST_FILE_PATH")
	contentType := "image/jpeg"

	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	log.Printf("See http://%s to browse to the object\n", endpoint+"/"+bucketName+"/"+objectName)
}
