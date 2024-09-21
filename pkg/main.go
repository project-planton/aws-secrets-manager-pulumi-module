package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretsmanager"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	PlaceholderSecretValue = "placeholder"
)

func Resources(ctx *pulumi.Context, stackInput *awssecretsmanager.AwsSecretsManagerStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.AwsCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"aws-classic",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	// For each secret in the input spec, create a secret in AWS Secrets Manager
	for _, secretName := range locals.AwsSecretsManager.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		// Construct the secret ID to make it unique within the AWS account
		secretId := fmt.Sprintf("%s-%s", locals.AwsSecretsManager.Metadata.Id, secretName)

		// Create the secret resource
		createdSecret, err := secretsmanager.NewSecret(ctx,
			secretName,
			&secretsmanager.SecretArgs{
				Name: pulumi.String(secretId),
				Tags: pulumi.ToStringMap(locals.AwsTags),
			}, pulumi.Provider(awsProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create secret")
		}

		// Create a secret version with a placeholder value
		_, err = secretsmanager.NewSecretVersion(ctx, secretId, &secretsmanager.SecretVersionArgs{
			SecretId:     createdSecret.ID(),
			SecretString: pulumi.String(PlaceholderSecretValue),
		}, pulumi.Parent(createdSecret), pulumi.IgnoreChanges([]string{"secretString"})) // Ignore secret value changes to avoid diffs
		if err != nil {
			return errors.Wrap(err, "failed to create placeholder secret version")
		}

		// Export the secret ID
		ctx.Export(fmt.Sprintf("%s-arn", secretName), createdSecret.Arn)
	}

	return nil
}
