package cmd

import (
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/s3"
"github.com/spf13/cobra"
)


var listS3Buckets = &cobra.Command{
	Use:   "s3-buckets",
	Short: "List S3 buckets",
	Long:  `List S3 buckets`,
	Run: func(cmd *cobra.Command, args []string) {

		s3svc := s3.New(session.New())
		result, err := s3svc.ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Failed to list buckets", err)
			return
		}

		fmt.Println("Buckets:")
		for _, bucket := range result.Buckets {
			fmt.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
		}
	},
}

func init() {
	RootCmd.AddCommand(listS3Buckets)
	//listS3Buckets.Flags().Bool(false, "local", "l", false, "Use local images")
}