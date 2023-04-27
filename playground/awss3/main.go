package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Person 객체 생성
	person := Person{
		Name:  "s3-kukim-tmp",
		Email: "kukim@example.com-tmp",
	}

	// Person 객체를 JSON으로 변환
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Failed to marshal person to JSON", err)
		return
	}

	// S3 버킷 이름과 객체 키
	bucket := "pg-go-s3"
	key := "path/s3/person-tmp.json"

	// 액세스 키 ID와 시크릿 액세스 키
	accessKey := ""
	secretKey := ""

	// AWS 세션 생성
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"), // 지역 설정
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		fmt.Println("Failed to create session", err)
		return
	}

	// S3 서비스 생성
	svc := s3.New(sess)

	// 업로드할 파일 정보 설정
	params := &s3.PutObjectInput{
		Bucket:  aws.String(bucket),
		Key:     aws.String(key),
		Body:    bytes.NewReader(jsonBytes),
		Expires: aws.Time(time.Now().Add(24 * time.Hour)), // 만료 시간 설정
	}

	// 파일 업로드
	_, err = svc.PutObject(params)
	if err != nil {
		fmt.Println("Failed to upload file", err)
		return
	}

	fmt.Println("File uploaded successfully")
}
