package cmd

import (
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

// initialize Command
var addCmd = &cobra.Command{
	Use:     "ec2 [command name]",
	Aliases: []string{"command"},
	Short:   "List EC2 instances",
	Long: `EC2 (cobra ec2)
Example: cobra add server  -> resulting in a new cmd/server.go
  `,

	Run: func(cmd *cobra.Command, args []string) {
		ec2svc := ec2.New(session.New())
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("tag:Environment"),
					Values: []*string{aws.String("prod")},
				},
				{
					Name:   aws.String("instance-state-name"),
					Values: []*string{aws.String("running"), aws.String("pending")},
				},
			},
		}
		resp, err := ec2svc.DescribeInstances(params)
		if err != nil {
			fmt.Println("there was an error listing instances in", err.Error())
			log.Fatal(err.Error())
		}

		for idx, res := range resp.Reservations {
			fmt.Println("  > Reservation Id", *res.ReservationId, " Num Instances: ", len(res.Instances))
			for _, inst := range resp.Reservations[idx].Instances {
				fmt.Println("    - Instance ID: ", *inst.InstanceId)
			}
		}
	},
}

