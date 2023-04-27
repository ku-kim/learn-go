package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/s3"
	"reflect"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Person 객체 생성
	person := Person{
		Name:  "s3-kukim-tmp-5-Cache-Control-333",
		Email: "kukim@example.com-tmp-5-Cache-Control-333",
	}

	//// Person 객체를 JSON으로 변환
	//jsonBytes, err := json.Marshal(person)
	//if err != nil {
	//	fmt.Println("Failed to marshal person to JSON", err)
	//	return
	//}

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

	fmt.Println(isPersonChanged(sess, bucket, key, person))

	uploadPersonToS3(sess, bucket, key, person)
}

func uploadPersonToS3(sess *session.Session, bucket string, key string, person Person) error {
	// 이전 Person 데이터와 현재 Person 데이터 비교
	if !isPersonChanged(sess, bucket, key, person) {
		fmt.Println("No changes in Person data")
		return nil
	}

	// Person 데이터를 JSON 형식으로 변환
	data, err := json.Marshal(person)
	if err != nil {
		return err
	}

	// S3에 파일 업로드
	//cacheControl := "no-cache, no-store, must-revalidate"
	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		//Expires: aws.Time(time.Now().Add(24 * time.Hour)), // 만료 시간 설정
		//CacheControl: &cacheControl,
	})
	if err != nil {
		return err
	}

	// CloudFront 캐시 무효화
	_, err = cloudfront.New(sess).CreateInvalidation(&cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(""),
		InvalidationBatch: &cloudfront.InvalidationBatch{ // 캐시 무효화 수행할 파일 목록, caller reference 값들
			CallerReference: aws.String(fmt.Sprintf("invalidation-%d", time.Now().Unix())), // 캐시 무효화 구분하기 위한 유니크 값
			Paths: &cloudfront.Paths{
				Items:    []*string{aws.String(fmt.Sprintf("/%s", key))}, // 캐시 무효화 수행할 파일 경로 목록
				Quantity: aws.Int64(1),                                   // 캐시 무효화할 파일 경로 개수와 동일해야함
			},
		},
	})
	if err != nil {
		return err
	}

	fmt.Println("Person data uploaded successfully and cache invalidated")
	return nil
}

func isPersonChanged(sess *session.Session, bucket string, key string, curr Person) bool {
	svc := s3.New(sess)
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return true
	}
	defer resp.Body.Close()

	var prev Person
	if err := json.NewDecoder(resp.Body).Decode(&prev); err != nil {
		return true
	}
	fmt.Println("prev: ", prev)
	fmt.Println("curr: ", curr)
	return !reflect.DeepEqual(prev, curr)
}
