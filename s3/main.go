package main

import (
	"s3/libs"
)

func main() {
	libs.S3.Region = "sa-east-1"
	libs.S3.NewSession(libs.S3.Region)
	libs.S3.Ls()
	libs.S3.Upload("UploadTeste.txt", "vms-triple", "UploadTesteS3.txt")
	libs.S3.GenerateUrl("vms-triple", "UploadTesteS3.txt")
	libs.S3.DeletaObjeto("vms-triple", "UploadTesteS3.txt")
}
