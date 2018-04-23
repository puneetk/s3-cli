package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/urfave/cli"
)

//
// Move object
// s3cmd mv s3://BUCKET1/OBJECT1 s3://BUCKET2[/OBJECT2]
//
func CmdMove(config *Config, c *cli.Context) error {
	args := c.Args()

	if len(args) < 2 {
		return fmt.Errorf("Not enought arguments to the move command")
	}
	dst := args[1]
	src := args[0]

	dst_u, err := FileURINew(dst)
	if err != nil {
		return fmt.Errorf("Invalid destination argument")
	}
	src_u, err := FileURINew(src)
	if err != nil {
		return fmt.Errorf("Invalid source argument")
	}
	if err := moveFile(config, src_u, dst_u); err != nil {
		return err
	}
	return nil
}
func moveFile(config *Config, src *FileURI, dst *FileURI) error {

	if src.Scheme != "s3" && dst.Scheme != "s3" {
		return fmt.Errorf("cp only supports s3 URLs")
	}
	copyOnS3(config, src, dst)

	svc, err := SessionForBucket(SessionNew(config), src.Bucket)
	if err != nil {
		return err
	}
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(src.Bucket),
		Key:    aws.String(src.Path),
	}
	result, err := svc.DeleteObject(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	return nil
}
