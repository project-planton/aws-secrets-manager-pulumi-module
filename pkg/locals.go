package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretsmanager"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	AwsSecretsManager *awssecretsmanager.AwsSecretsManager
	AwsTags           map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awssecretsmanager.AwsSecretsManagerStackInput) *Locals {
	locals := &Locals{}
	locals.AwsSecretsManager = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsSecretsManager.Spec.EnvironmentInfo.OrgId,
		awstagkeys.Environment:  locals.AwsSecretsManager.Spec.EnvironmentInfo.EnvId,
		awstagkeys.ResourceKind: apiresourcekind.ApiResourceKind_aws_secrets_manager.String(),
		awstagkeys.ResourceId:   locals.AwsSecretsManager.Metadata.Id,
	}
	return locals
}
