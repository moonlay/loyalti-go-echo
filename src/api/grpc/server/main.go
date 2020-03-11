package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"github.com/radyatamaa/loyalti-go-echo/src/api/grpc/pb"
)

const (
	port = "localhost:50051"
)

type server struct {
}

//func (s *server) Upload(stream pb.GuploadService_UploadServer) (err error) {
//	// open output file
//	fo, err := os.Create("./output")
//	if err != nil {
//		return errors.New("failed to create file")
//	}
//	// close fo on exit and check for its returned error
//	defer func() {
//		if err := fo.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	var res *pb.Chunk
//	for {
//		res, err = stream.Recv()
//
//		if err == io.EOF {
//			err = stream.SendAndClose(&pb.UploadStatus{
//				Message: "Upload received with success",
//				Code:    pb.UploadStatusCode_Ok,
//			})
//			if err != nil {
//				err =  errors.New("failed to send status code")
//				return err
//			}
//			return nil
//		}
//		//fmt.Println(res.Received)
//		// write a chunk
//		if _, err := fo.Write(res.Content); err != nil {
//			err =  errors.New(err.Error())
//			return err
//		}
//	}
//}
func (s *server) Upload(stream pb.GuploadService_UploadServer) (err error) {

	//var filename = "C:/Users/radyatama.suryana/Go/loyalti-go-echo/src/70815263.jpg"
	var myBucket = "loyalti-storage"
	var myKey = "merchant/merchant_profile/test21.jpg"
	var accessKey = "AKIAJCWHV6G7IRA2G2WA"
	var accessSecret = "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"
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" {
		//load default credentials
		awsConfig = &aws.Config{
			Region: aws.String("ap-southeast-1"),
		}
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	// Create an uploader with the session and default options
	//uploader := s3manager.NewUploader(sess)

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 2            // default is 5
	})

	var image = "test.jpg"
	// open output file
	f, err := os.Create("." + "/" + image)
	if err != nil {
		return errors.New("failed to create file")
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	var res *pb.Chunk
	for {
		res, err = stream.Recv()

		if err == io.EOF {
			err = stream.SendAndClose(&pb.UploadStatus{
				Message: "Upload received with success",
				Code:    pb.UploadStatusCode_Ok,
			})
			if err != nil {
				err =  errors.New("failed to send status code")
				return err
			}
			return nil
		}
		//checkTypeData
		//fmt.Println(reflect.TypeOf(test).String())
		//fmt.Print(&test)

		//// write a chunk
		if err != nil {
			fmt.Printf("failed to upload file, %v", err)
			return err
		}


		if _, err := f.Write(res.Content); err != nil {
			err =  errors.New(err.Error())
			return err
		}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)

		var tempImage = dir + "\\" + image
		tempImage = strings.Replace(tempImage, "\\", "/", -1)
		fmt.Println(tempImage)
		//var imagePath = "queen.jpg"

		//fmt.Print(res.String())
		fo, err := os.Open(tempImage)
		if err != nil {
			fmt.Printf("failed to open file %q, %v", fo, err)
			return err
		}
		fo.Write(res.Content)

		//fmt.Print(fo)
		//
		//defer f.Close()

		// Upload the file to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(myKey),
			ACL:    aws.String("public-read"),
			Body:   fo,
		})

		fmt.Printf("file uploaded to, %s\n", result.Location)

	}
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterGuploadServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("server Start")
}