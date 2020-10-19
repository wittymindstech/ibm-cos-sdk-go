// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/awserr"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To abort a multipart upload
//
// The following example aborts a multipart upload.
func ExampleS3_AbortMultipartUpload_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.AbortMultipartUploadInput{
		Bucket:   aws.String("examplebucket"),
		Key:      aws.String("bigobject"),
		UploadId: aws.String("xadcOB_7YPBOJuoFiQ9cz4P3Pe6FIZwO4f7wN93uHsNBEw97pl5eNwzExg0LAT2dUN91cOmrEQHDsP3WA60CEg--"),
	}

	result, err := svc.AbortMultipartUpload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchUpload:
				fmt.Println(s3.ErrCodeNoSuchUpload, aerr.Error())
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

	fmt.Println(result)
}

// To complete multipart upload
//
// The following example completes a multipart upload.
func ExampleS3_CompleteMultipartUpload_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.CompleteMultipartUploadInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("bigobject"),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: []*s3.CompletedPart{
				{
					ETag:       aws.String("\"d8c2eafd90c266e19ab9dcacc479f8af\""),
					PartNumber: aws.Int64(1),
				},
				{
					ETag:       aws.String("\"d8c2eafd90c266e19ab9dcacc479f8af\""),
					PartNumber: aws.Int64(2),
				},
			},
		},
		UploadId: aws.String("7YPBOJuoFiQ9cz4P3Pe6FIZwO4f7wN93uHsNBEw97pl5eNwzExg0LAT2dUN91cOmrEQHDsP3WA60CEg--"),
	}

	result, err := svc.CompleteMultipartUpload(input)
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

	fmt.Println(result)
}

// To copy an object
//
// The following example copies an object from one bucket to another.
func ExampleS3_CopyObject_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.CopyObjectInput{
		Bucket:     aws.String("destinationbucket"),
		CopySource: aws.String("/sourcebucket/HappyFacejpg"),
		Key:        aws.String("HappyFaceCopyjpg"),
	}

	result, err := svc.CopyObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeObjectNotInActiveTierError:
				fmt.Println(s3.ErrCodeObjectNotInActiveTierError, aerr.Error())
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

	fmt.Println(result)
}

// To create a bucket in a specific region
//
// The following example creates a bucket. The request specifies an AWS region where
// to create the bucket.
func ExampleS3_CreateBucket_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.CreateBucketInput{
		Bucket: aws.String("examplebucket"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String("eu-west-1"),
		},
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
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

	fmt.Println(result)
}

// To create a bucket
//
// The following example creates a bucket.
func ExampleS3_CreateBucket_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.CreateBucketInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
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

	fmt.Println(result)
}

// To initiate a multipart upload
//
// The following example initiates a multipart upload.
func ExampleS3_CreateMultipartUpload_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("largeobject"),
	}

	result, err := svc.CreateMultipartUpload(input)
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

	fmt.Println(result)
}

// To delete a bucket
//
// The following example deletes the specified bucket.
func ExampleS3_DeleteBucket_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteBucketInput{
		Bucket: aws.String("forrandall2"),
	}

	result, err := svc.DeleteBucket(input)
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

	fmt.Println(result)
}

// To delete cors configuration on a bucket.
//
// The following example deletes CORS configuration on a bucket.
func ExampleS3_DeleteBucketCors_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteBucketCorsInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.DeleteBucketCors(input)
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

	fmt.Println(result)
}

// To delete bucket website configuration
//
// The following example deletes bucket website configuration.
func ExampleS3_DeleteBucketWebsite_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteBucketWebsiteInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.DeleteBucketWebsite(input)
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

	fmt.Println(result)
}

// To delete an object (from a non-versioned bucket)
//
// The following example deletes an object from a non-versioned bucket.
func ExampleS3_DeleteObject_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("ExampleBucket"),
		Key:    aws.String("HappyFace.jpg"),
	}

	result, err := svc.DeleteObject(input)
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

	fmt.Println(result)
}

// To delete multiple objects from a versioned bucket
//
// The following example deletes objects from a bucket. The bucket is versioned, and
// the request does not specify the object version to delete. In this case, all versions
// remain in the bucket and S3 adds a delete marker.
func ExampleS3_DeleteObjects_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String("examplebucket"),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{
				{
					Key: aws.String("objectkey1"),
				},
				{
					Key: aws.String("objectkey2"),
				},
			},
			Quiet: aws.Bool(false),
		},
	}

	result, err := svc.DeleteObjects(input)
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

	fmt.Println(result)
}

// To delete multiple object versions from a versioned bucket
//
// The following example deletes objects from a bucket. The request specifies object
// versions. S3 deletes specific object versions and returns the key and versions of
// deleted objects in the response.
func ExampleS3_DeleteObjects_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String("examplebucket"),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{
				{
					Key:       aws.String("HappyFace.jpg"),
					VersionId: aws.String("2LWg7lQLnY41.maGB5Z6SWW.dcq0vx7b"),
				},
				{
					Key:       aws.String("HappyFace.jpg"),
					VersionId: aws.String("yoz3HB.ZhCS_tKVEmIOr7qYyyAaZSKVd"),
				},
			},
			Quiet: aws.Bool(false),
		},
	}

	result, err := svc.DeleteObjects(input)
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

	fmt.Println(result)
}

// To get cors configuration set on a bucket
//
// The following example returns cross-origin resource sharing (CORS) configuration
// set on a bucket.
func ExampleS3_GetBucketCors_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetBucketCorsInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.GetBucketCors(input)
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

	fmt.Println(result)
}

// To get bucket location
//
// The following example returns bucket location.
func ExampleS3_GetBucketLocation_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetBucketLocationInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.GetBucketLocation(input)
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

	fmt.Println(result)
}

// To get bucket website configuration
//
// The following example retrieves website configuration of a bucket.
func ExampleS3_GetBucketWebsite_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetBucketWebsiteInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.GetBucketWebsite(input)
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

	fmt.Println(result)
}

// To retrieve an object
//
// The following example retrieves an object for an S3 bucket.
func ExampleS3_GetObject_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetObjectInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("HappyFace.jpg"),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
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

	fmt.Println(result)
}

// To retrieve a byte range of an object
//
// The following example retrieves an object for an S3 bucket. The request specifies
// the range header to retrieve a specific byte range.
func ExampleS3_GetObject_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetObjectInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("SampleFile.txt"),
		Range:  aws.String("bytes=0-9"),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
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

	fmt.Println(result)
}

// To retrieve object ACL
//
// The following example retrieves access control list (ACL) of an object.
func ExampleS3_GetObjectAcl_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.GetObjectAclInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("HappyFace.jpg"),
	}

	result, err := svc.GetObjectAcl(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
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

	fmt.Println(result)
}

// To determine if bucket exists
//
// This operation checks to see if a bucket exists.
func ExampleS3_HeadBucket_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.HeadBucketInput{
		Bucket: aws.String("acl1"),
	}

	result, err := svc.HeadBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
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

	fmt.Println(result)
}

// To retrieve metadata of an object without returning the object itself
//
// The following example retrieves an object metadata.
func ExampleS3_HeadObject_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.HeadObjectInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("HappyFace.jpg"),
	}

	result, err := svc.HeadObject(input)
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

	fmt.Println(result)
}

// To list object versions
//
// The following example return versions of an object with specific key name prefix.
// The request limits the number of items returned to two. If there are are more than
// two object version, S3 returns NextToken in the response. You can specify this token
// value in your next request to fetch next set of object versions.
func ExampleS3_ListBuckets_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
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

	fmt.Println(result)
}

// To list in-progress multipart uploads on a bucket
//
// The following example lists in-progress multipart uploads on a specific bucket.
func ExampleS3_ListMultipartUploads_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListMultipartUploadsInput{
		Bucket: aws.String("examplebucket"),
	}

	result, err := svc.ListMultipartUploads(input)
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

	fmt.Println(result)
}

// List next set of multipart uploads when previous result is truncated
//
// The following example specifies the upload-id-marker and key-marker from previous
// truncated response to retrieve next setup of multipart uploads.
func ExampleS3_ListMultipartUploads_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListMultipartUploadsInput{
		Bucket:         aws.String("examplebucket"),
		KeyMarker:      aws.String("nextkeyfrompreviousresponse"),
		MaxUploads:     aws.Int64(2),
		UploadIdMarker: aws.String("valuefrompreviousresponse"),
	}

	result, err := svc.ListMultipartUploads(input)
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

	fmt.Println(result)
}

// To list objects in a bucket
//
// The following example list two objects in a bucket.
func ExampleS3_ListObjects_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListObjectsInput{
		Bucket:  aws.String("examplebucket"),
		MaxKeys: aws.Int64(2),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
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

	fmt.Println(result)
}

// To get object list
//
// The following example retrieves object list. The request specifies max keys to limit
// response to include only 2 object keys.
func ExampleS3_ListObjectsV2_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String("examplebucket"),
		MaxKeys: aws.Int64(2),
	}

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
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

	fmt.Println(result)
}

// To list parts of a multipart upload.
//
// The following example lists parts uploaded for a specific multipart upload.
func ExampleS3_ListParts_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.ListPartsInput{
		Bucket:   aws.String("examplebucket"),
		Key:      aws.String("bigobject"),
		UploadId: aws.String("example7YPBOJuoFiQ9cz4P3Pe6FIZwO4f7wN93uHsNBEw97pl5eNwzExg0LAT2dUN91cOmrEQHDsP3WA60CEg--"),
	}

	result, err := svc.ListParts(input)
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

	fmt.Println(result)
}

// Put bucket acl
//
// The following example replaces existing ACL on a bucket. The ACL grants the bucket
// owner (specified using the owner ID) and write permission to the LogDelivery group.
// Because this is a replace operation, you must specify all the grants in your request.
// To incrementally add or remove ACL grants, you might use the console.
func ExampleS3_PutBucketAcl_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutBucketAclInput{
		Bucket:           aws.String("examplebucket"),
		GrantFullControl: aws.String("id=examplee7a2f25102679df27bb0ae12b3f85be6f290b936c4393484"),
		GrantWrite:       aws.String("uri=http://acs.amazonaws.com/groups/s3/LogDelivery"),
	}

	result, err := svc.PutBucketAcl(input)
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

	fmt.Println(result)
}

// To set cors configuration on a bucket.
//
// The following example enables PUT, POST, and DELETE requests from www.example.com,
// and enables GET requests from any domain.
func ExampleS3_PutBucketCors_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutBucketCorsInput{
		Bucket: aws.String(""),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{
				{
					AllowedHeaders: []*string{
						aws.String("*"),
					},
					AllowedMethods: []*string{
						aws.String("PUT"),
						aws.String("POST"),
						aws.String("DELETE"),
					},
					AllowedOrigins: []*string{
						aws.String("http://www.example.com"),
					},
					ExposeHeaders: []*string{
						aws.String("x-amz-server-side-encryption"),
					},
					MaxAgeSeconds: aws.Int64(3000),
				},
				{
					AllowedHeaders: []*string{
						aws.String("Authorization"),
					},
					AllowedMethods: []*string{
						aws.String("GET"),
					},
					AllowedOrigins: []*string{
						aws.String("*"),
					},
					MaxAgeSeconds: aws.Int64(3000),
				},
			},
		},
	}

	result, err := svc.PutBucketCors(input)
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

	fmt.Println(result)
}

// To upload an object and specify canned ACL.
//
// The following example uploads and object. The request specifies optional canned ACL
// (access control list) to all READ access to authenticated users. If the bucket is
// versioning enabled, S3 returns version ID in response.
func ExampleS3_PutObject_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		ACL:    aws.String("authenticated-read"),
		Body:   aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("exampleobject"),
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To create an object.
//
// The following example creates an object. If the bucket is versioning enabled, S3
// returns version ID in response.
func ExampleS3_PutObject_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("objectkey"),
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To upload an object (specify optional headers)
//
// The following example uploads an object. The request specifies optional request headers
// to directs S3 to use specific storage class and use server-side encryption.
func ExampleS3_PutObject_shared02() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("HappyFace.jpg")),
		Bucket:               aws.String("examplebucket"),
		Key:                  aws.String("HappyFace.jpg"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("STANDARD_IA"),
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To upload an object
//
// The following example uploads an object to a versioning-enabled bucket. The source
// file is specified using Windows file syntax. S3 returns VersionId of the newly created
// object.
func ExampleS3_PutObject_shared03() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("HappyFace.jpg")),
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("HappyFace.jpg"),
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To upload object and specify user-defined metadata
//
// The following example creates an object. The request also specifies optional metadata.
// If the bucket is versioning enabled, S3 returns version ID in response.
func ExampleS3_PutObject_shared04() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("exampleobject"),
		Metadata: map[string]*string{
			"metadata1": aws.String("value1"),
			"metadata2": aws.String("value2"),
		},
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To upload an object and specify server-side encryption and object tags
//
// The following example uploads and object. The request specifies the optional server-side
// encryption option. The request also specifies optional object tags. If the bucket
// is versioning enabled, S3 returns version ID in response.
func ExampleS3_PutObject_shared05() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectInput{
		Body:                 aws.ReadSeekCloser(strings.NewReader("filetoupload")),
		Bucket:               aws.String("examplebucket"),
		Key:                  aws.String("exampleobject"),
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              aws.String("key1=value1&key2=value2"),
	}

	result, err := svc.PutObject(input)
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

	fmt.Println(result)
}

// To grant permissions using object ACL
//
// The following example adds grants to an object ACL. The first permission grants user1
// and user2 FULL_CONTROL and the AllUsers group READ permission.
func ExampleS3_PutObjectAcl_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.PutObjectAclInput{
		AccessControlPolicy: &s3.AccessControlPolicy{},
		Bucket:              aws.String("examplebucket"),
		GrantFullControl:    aws.String("emailaddress=user1@example.com,emailaddress=user2@example.com"),
		GrantRead:           aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"),
		Key:                 aws.String("HappyFace.jpg"),
	}

	result, err := svc.PutObjectAcl(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
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

	fmt.Println(result)
}

// To upload a part
//
// The following example uploads part 1 of a multipart upload. The example specifies
// a file name for the part data. The Upload ID is same that is returned by the initiate
// multipart upload.
func ExampleS3_UploadPart_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.UploadPartInput{
		Body:       aws.ReadSeekCloser(strings.NewReader("fileToUpload")),
		Bucket:     aws.String("examplebucket"),
		Key:        aws.String("examplelargeobject"),
		PartNumber: aws.Int64(1),
		UploadId:   aws.String("xadcOB_7YPBOJuoFiQ9cz4P3Pe6FIZwO4f7wN93uHsNBEw97pl5eNwzExg0LAT2dUN91cOmrEQHDsP3WA60CEg--"),
	}

	result, err := svc.UploadPart(input)
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

	fmt.Println(result)
}

// To upload a part by copying byte range from an existing object as data source
//
// The following example uploads a part of a multipart upload by copying a specified
// byte range from an existing object as data source.
func ExampleS3_UploadPartCopy_shared00() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.UploadPartCopyInput{
		Bucket:          aws.String("examplebucket"),
		CopySource:      aws.String("/bucketname/sourceobjectkey"),
		CopySourceRange: aws.String("bytes=1-100000"),
		Key:             aws.String("examplelargeobject"),
		PartNumber:      aws.Int64(2),
		UploadId:        aws.String("exampleuoh_10OhKhT7YukE9bjzTPRiuaCotmZM_pFngJFir9OZNrSr5cWa3cq3LZSUsfjI4FI7PkP91We7Nrw--"),
	}

	result, err := svc.UploadPartCopy(input)
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

	fmt.Println(result)
}

// To upload a part by copying data from an existing object as data source
//
// The following example uploads a part of a multipart upload by copying data from an
// existing object as data source.
func ExampleS3_UploadPartCopy_shared01() {
	svc := s3.New(session.Must(session.NewSession()))
	input := &s3.UploadPartCopyInput{
		Bucket:     aws.String("examplebucket"),
		CopySource: aws.String("/bucketname/sourceobjectkey"),
		Key:        aws.String("examplelargeobject"),
		PartNumber: aws.Int64(1),
		UploadId:   aws.String("exampleuoh_10OhKhT7YukE9bjzTPRiuaCotmZM_pFngJFir9OZNrSr5cWa3cq3LZSUsfjI4FI7PkP91We7Nrw--"),
	}

	result, err := svc.UploadPartCopy(input)
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

	fmt.Println(result)
}