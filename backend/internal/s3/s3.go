package bucket

import (
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var client *s3.S3

func Setup() {
	sess, err := session.NewSession(&aws.Config{

		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT_URL")),
	})

	if err != nil {
		log.Fatal(err)
	}

	client = s3.New(sess)

}

func Upload(file multipart.File, fileheader *multipart.FileHeader, image_name string) (string, error) {

	_, err := client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("/media"),
		Key:         aws.String(image_name + "/" + fileheader.Filename),
		Body:        file,
		ContentType: aws.String("image"),
	})

	if err != nil {
		return "", err
	}

	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("/media"),
		Key:    aws.String(image_name + "/" + fileheader.Filename),
	})
	object_url, err := req.Presign(time.Hour * 24 * 365)

	if err != nil {
		return "", err
	}

	return object_url, nil
}
