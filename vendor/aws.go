package vendor

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	s3region = "us-west-1"
)

// UploadToS3 func, uploads a recording to s3
// once ge a response from the server then
func UploadToS3(filename string, key string) (string, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(s3region)})
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("dsound-boy/" + filename),
		Key:    aws.String(key + ".caf"),
		Body:   f,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return aws.StringValue(&result.Location), nil
}

// CleanupAfterUpload func, will clean up
// the temp dirs and files created during
// the multipart form and upload
func CleanupAfterUpload(temp string) {
	err := os.RemoveAll(temp)
	if err != nil {
		fmt.Println("error deleting temp folder ", err)
	}

}
