package cmd

import (
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/s3"
"github.com/spf13/cobra"
	"os"
)

func s3Buckets() {
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
}

var listS3Buckets = &cobra.Command{
	Use:   "analyze IMAGE",
	Short: "Analyze Docker image",
	Long:  `Analyze a Docker image with Clair, against Ubuntu, Red hat and Debian vulnerabilities databases`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Printf("clairctl: \"analyze\" requires a minimum of 1 argument")
			os.Exit(1)
		}

		config.ImageName = args[0]
		image, manifest, err := docker.RetrieveManifest(config.ImageName, true)
		if err != nil {
			fmt.Println(errInternalError)
			log.Fatalf("retrieving manifest for %q: %v", config.ImageName, err)
		}

		startLocalServer()
		if err := clair.Push(image, manifest); err != nil {
			if err != nil {
				fmt.Println(errInternalError)
				log.Fatalf("pushing image %q: %v", image.String(), err)
			}
		}

		analysis := clair.Analyze(image, manifest)
		err = template.Must(template.New("analysis").Parse(analyzeTplt)).Execute(os.Stdout, analysis)
		if err != nil {
			fmt.Println(errInternalError)
			log.Fatalf("rendering analysis: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(listS3Buckets)
	listS3Buckets.Flags().BoolVarP(false, "local", "l", false, "Use local images")
}