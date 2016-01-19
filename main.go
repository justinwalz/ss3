package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {
	var err error
	var match = flag.String("match", "", "pattern to match when listing keys")

	flag.Parse()
	args := flag.Args()

	auth, err := aws.SharedAuth()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to authenticate: %s\n", err.Error())
		os.Exit(1)
	}

	client := s3.New(auth, aws.USEast)

	if len(args) == 0 {
		err = listBuckets(client, *match)
	} else {
		err = listObjectsInBuckets(client, args, *match)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

}

func listBuckets(client *s3.S3, match string) error {
	resp, err := client.ListBuckets()
	if err != nil {
		return err
	}

	for _, bucket := range resp.Buckets {
		if len(match) > 0 {
			if strings.HasPrefix(bucket.Name, match) {
				fmt.Printf("%s\n", bucket.Name)
			}
		} else {
			fmt.Printf("%s\n", bucket.Name)
		}
	}

	return nil
}

func listObjectsInBuckets(client *s3.S3, buckets []string, match string) error {
	for _, b := range buckets {
		bucket := client.Bucket(b)

		contents, err := bucket.GetBucketContents()
		if err != nil {
			return err
		}

		matches := []string{}
		for key, value := range *contents {
			if len(match) > 0 {
				if strings.HasPrefix(key, match) {
					matches = append(matches, key)
				}
			} else {
				fmt.Printf("%s (%d)\n", key, value.Size)
			}
		}

		if len(match) > 0 {
			sort.Strings(matches)
			fmt.Printf("%s\n", matches[len(matches)-1])
		}

		return nil
	}

	return nil
}
