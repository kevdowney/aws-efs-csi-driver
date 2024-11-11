// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to manage storage for your file system. A LifecycleConfiguration
// consists of one or more LifecyclePolicy objects that define the following:
//
//   - TransitionToIA – When to move files in the file system from primary storage
//     (Standard storage class) into the Infrequent Access (IA) storage.
//
//   - TransitionToArchive – When to move files in the file system from their
//     current storage class (either IA or Standard storage) into the Archive storage.
//
// File systems cannot transition into Archive storage before transitioning into
//
//	IA storage. Therefore, TransitionToArchive must either not be set or must be
//	later than TransitionToIA.
//
// The Archive storage class is available only for file systems that use the
//
//	Elastic Throughput mode and the General Purpose Performance mode.
//
//	- TransitionToPrimaryStorageClass – Whether to move files in the file system
//	back to primary storage (Standard storage class) after they are accessed in IA
//	or Archive storage.
//
// For more information, see [Managing file system storage].
//
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration . In the request, specify the following:
//
//   - The ID for the file system for which you are enabling, disabling, or
//     modifying Lifecycle management.
//
//   - A LifecyclePolicies array of LifecyclePolicy objects that define when to
//     move files to IA storage, to Archive storage, and back to primary storage.
//
// Amazon EFS requires that each LifecyclePolicy object have only have a single
//
//	transition, so the LifecyclePolicies array needs to be structured with
//	separate LifecyclePolicy objects. See the example requests in the following
//	section for more information.
//
// This operation requires permissions for the
// elasticfilesystem:PutLifecycleConfiguration operation.
//
// To apply a LifecycleConfiguration object to an encrypted file system, you need
// the same Key Management Service permissions as when you created the encrypted
// file system.
//
// [Managing file system storage]: https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html
func (c *Client) PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleConfiguration", params, optFns, c.addOperationPutLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleConfigurationInput struct {

	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	//
	// This member is required.
	FileSystemId *string

	// An array of LifecyclePolicy objects that define the file system's
	// LifecycleConfiguration object. A LifecycleConfiguration object informs EFS
	// Lifecycle management of the following:
	//
	//   - TransitionToIA – When to move files in the file system from primary storage
	//   (Standard storage class) into the Infrequent Access (IA) storage.
	//
	//   - TransitionToArchive – When to move files in the file system from their
	//   current storage class (either IA or Standard storage) into the Archive storage.
	//
	// File systems cannot transition into Archive storage before transitioning into
	//   IA storage. Therefore, TransitionToArchive must either not be set or must be
	//   later than TransitionToIA.
	//
	// The Archive storage class is available only for file systems that use the
	//   Elastic Throughput mode and the General Purpose Performance mode.
	//
	//   - TransitionToPrimaryStorageClass – Whether to move files in the file system
	//   back to primary storage (Standard storage class) after they are accessed in IA
	//   or Archive storage.
	//
	// When using the put-lifecycle-configuration CLI command or the
	// PutLifecycleConfiguration API action, Amazon EFS requires that each
	// LifecyclePolicy object have only a single transition. This means that in a
	// request body, LifecyclePolicies must be structured as an array of
	// LifecyclePolicy objects, one object for each storage transition. See the example
	// requests in the following section for more information.
	//
	// This member is required.
	LifecyclePolicies []types.LifecyclePolicy

	noSmithyDocumentSerde
}

type PutLifecycleConfigurationOutput struct {

	// An array of lifecycle management policies. EFS supports a maximum of one policy
	// per file system.
	LifecyclePolicies []types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutLifecycleConfiguration"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpPutLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutLifecycleConfiguration",
	}
}