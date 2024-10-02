package pkg

import (
	awssecretsmanagerv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/aws/awssecretsmanager/v1"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	AwsSecretsManager *awssecretsmanagerv1.AwsSecretsManager
	AwsTags           map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awssecretsmanagerv1.AwsSecretsManagerStackInput) *Locals {
	locals := &Locals{}
	locals.AwsSecretsManager = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsSecretsManager.Spec.EnvironmentInfo.OrgId,
		awstagkeys.Environment:  locals.AwsSecretsManager.Spec.EnvironmentInfo.EnvId,
		awstagkeys.ResourceKind: "aws-secrets-manager",
		awstagkeys.ResourceId:   locals.AwsSecretsManager.Metadata.Id,
	}
	return locals
}
