package outputs

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretsmanager"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *awssecretsmanager.AwsSecretsManagerStackInput) *awssecretsmanager.AwsSecretsManagerStackOutputs {
	stackOutputs := &awssecretsmanager.AwsSecretsManagerStackOutputs{
		SecretArnMap: make(map[string]string),
	}
	for _, secretName := range input.ApiResource.Spec.SecretNames {
		stackOutputs.SecretArnMap[secretName] = autoapistackoutput.GetVal(pulumiOutputs,
			fmt.Sprintf("%s-arn", secretName))
	}
	return stackOutputs
}
