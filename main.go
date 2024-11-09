package main

import (
	"github.com/pkg/errors"
	"github.com/project-planton/aws-secrets-manager-pulumi-module/pkg"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/aws/awssecretsmanager/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &awssecretsmanagerv1.AwsSecretsManagerStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
