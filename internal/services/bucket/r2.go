package bucket

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Config struct {
	Bucket    string
	AccountID string
	Key       string
	Secret    string
}

type R2Bucket struct {
	Client *s3.Client
	Config R2Config
}

func InitR2Bucket() Bucket {
	r2 := R2Config{
		Bucket:    "beholder",
		AccountID: os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		Key:       os.Getenv("CLOUDFLARE_API_KEY"),
		Secret:    os.Getenv("CLOUDFLARE_API_SECRET"),
	}
	r2revolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: os.Getenv("CLOUDFLARE_R2_URL"),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2revolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.Key, r2.Secret, "")),
		config.WithRegion("auto"),
	)

	if err != nil {
		fmt.Printf("[R2 Bucket] Fail to init %s", err.Error())
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	return R2Bucket{
		Client: client,
		Config: r2,
	}
}

func (b R2Bucket) DownloadFile(objectKey string, fileName string) error {
	output, err := b.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(b.Config.Bucket),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		return err
	}

	defer output.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Couldn't create file ", objectKey, err)
		return err
	}
	defer file.Close()
	file.ReadFrom(output.Body)
	return nil
}

func (b R2Bucket) UploadFile(objectKey string, filename string, filetype string) error {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Couldn't open file: ", err)
		return err
	}

	defer file.Close()
	_, err = b.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(b.Config.Bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(filetype),
	})

	if err != nil {
		log.Println("Couldn't upload file to S3: ", err)
		return err
	}

	return nil
}
