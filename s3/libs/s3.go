package libs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var S3 *S3Client

func init() {
	S3 = new(S3Client)
}

type S3Client struct {
	Region string
	Sess   *session.Session
	Svc    *s3.S3
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
func (t *S3Client) NewSession(region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		exitErrorf("Problema com a sessão do S3, %v", err)
	}
	t.Sess = sess
	t.Svc = s3.New(t.Sess)
}

/*
lista todos Buckets da conta aws
*/
func (t *S3Client) Ls() {
	result, err := t.Svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Não foi possivel listar os Buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s creado em %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

/*
sobe um arquivo para o S3

filename string arquivo locaFailed to sign requestl para upload
myBucket string nome do bucket no s3
keyName  string nome do objetos final criado no S3 com o caminho completo (sem o nome do bucket)
*/

func (t *S3Client) Upload(filename string, myBucket string, keyName string) {
	uploader := s3manager.NewUploader(t.Sess)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file %q, %v", filename, err))
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(keyName),
		Body:   f,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to upload file, %v", err))
	}
	fmt.Println(result)
}

/*
	Gera URL publica do objeto no S3

	myBucket string nome do bucket no s3

	keyName string nome do objeto para download com o caminho completo
*/
func (t *S3Client) GenerateUrl(myBucket string, keyName string) string {
	req, _ := t.Svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(keyName),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Falha ao gerar a requisição", err)
	}
	fmt.Println(urlStr)
	return urlStr

}

func (t *S3Client) DeletaObjeto(myBucket string, keyName string) {
	input := &s3.DeleteObjectInput{
		Bucket: &myBucket,
		Key:    &keyName,
	}

	_, err := t.Svc.DeleteObject(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return

	}

	fmt.Println(keyName + " deletado com sucesso!")
}
