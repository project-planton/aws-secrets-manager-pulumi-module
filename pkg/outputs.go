package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretsmanagersecretset/model"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *model.AwsSecretsManagerSecretSetStackInput) *model.AwsSecretsManagerSecretSetStackOutputs {
	return &model.AwsSecretsManagerSecretSetStackOutputs{}
}
