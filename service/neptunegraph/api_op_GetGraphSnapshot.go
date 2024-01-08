// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Retrieves a specified graph snapshot.
func (c *Client) GetGraphSnapshot(ctx context.Context, params *GetGraphSnapshotInput, optFns ...func(*Options)) (*GetGraphSnapshotOutput, error) {
	if params == nil {
		params = &GetGraphSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGraphSnapshot", params, optFns, c.addOperationGetGraphSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGraphSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGraphSnapshotInput struct {

	// The ID of the snapshot to retrieve.
	//
	// This member is required.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

func (in *GetGraphSnapshotInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type GetGraphSnapshotOutput struct {

	// The ARN of the graph snapshot.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph snapshot.
	//
	// This member is required.
	Id *string

	// The snapshot name. For example: my-snapshot-1 . The name must contain from 1 to
	// 63 letters, numbers, or hyphens, and its first character must be a letter. It
	// cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	Name *string

	// The ID of the KMS key used to encrypt and decrypt the snapshot.
	KmsKeyIdentifier *string

	// The time when the snapshot was created.
	SnapshotCreateTime *time.Time

	// The graph identifier for the graph for which a snapshot is to be created.
	SourceGraphId *string

	// The status of the graph snapshot.
	Status types.SnapshotStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGraphSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGraphSnapshot"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetGraphSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGraphSnapshot(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// GetGraphSnapshotAPIClient is a client that implements the GetGraphSnapshot
// operation.
type GetGraphSnapshotAPIClient interface {
	GetGraphSnapshot(context.Context, *GetGraphSnapshotInput, ...func(*Options)) (*GetGraphSnapshotOutput, error)
}

var _ GetGraphSnapshotAPIClient = (*Client)(nil)

// GraphSnapshotAvailableWaiterOptions are waiter options for
// GraphSnapshotAvailableWaiter
type GraphSnapshotAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// GraphSnapshotAvailableWaiter will use default minimum delay of 60 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, GraphSnapshotAvailableWaiter will use default max delay of 7200
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetGraphSnapshotInput, *GetGraphSnapshotOutput, error) (bool, error)
}

// GraphSnapshotAvailableWaiter defines the waiters for GraphSnapshotAvailable
type GraphSnapshotAvailableWaiter struct {
	client GetGraphSnapshotAPIClient

	options GraphSnapshotAvailableWaiterOptions
}

// NewGraphSnapshotAvailableWaiter constructs a GraphSnapshotAvailableWaiter.
func NewGraphSnapshotAvailableWaiter(client GetGraphSnapshotAPIClient, optFns ...func(*GraphSnapshotAvailableWaiterOptions)) *GraphSnapshotAvailableWaiter {
	options := GraphSnapshotAvailableWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 7200 * time.Second
	options.Retryable = graphSnapshotAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &GraphSnapshotAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for GraphSnapshotAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *GraphSnapshotAvailableWaiter) Wait(ctx context.Context, params *GetGraphSnapshotInput, maxWaitDur time.Duration, optFns ...func(*GraphSnapshotAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for GraphSnapshotAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *GraphSnapshotAvailableWaiter) WaitForOutput(ctx context.Context, params *GetGraphSnapshotInput, maxWaitDur time.Duration, optFns ...func(*GraphSnapshotAvailableWaiterOptions)) (*GetGraphSnapshotOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 7200 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetGraphSnapshot(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for GraphSnapshotAvailable waiter")
}

func graphSnapshotAvailableStateRetryable(ctx context.Context, input *GetGraphSnapshotInput, output *GetGraphSnapshotOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(types.SnapshotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SnapshotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.SnapshotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SnapshotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "AVAILABLE"
		value, ok := pathValue.(types.SnapshotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SnapshotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

// GraphSnapshotDeletedWaiterOptions are waiter options for
// GraphSnapshotDeletedWaiter
type GraphSnapshotDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// GraphSnapshotDeletedWaiter will use default minimum delay of 60 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, GraphSnapshotDeletedWaiter will use default max delay of 3600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetGraphSnapshotInput, *GetGraphSnapshotOutput, error) (bool, error)
}

// GraphSnapshotDeletedWaiter defines the waiters for GraphSnapshotDeleted
type GraphSnapshotDeletedWaiter struct {
	client GetGraphSnapshotAPIClient

	options GraphSnapshotDeletedWaiterOptions
}

// NewGraphSnapshotDeletedWaiter constructs a GraphSnapshotDeletedWaiter.
func NewGraphSnapshotDeletedWaiter(client GetGraphSnapshotAPIClient, optFns ...func(*GraphSnapshotDeletedWaiterOptions)) *GraphSnapshotDeletedWaiter {
	options := GraphSnapshotDeletedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 3600 * time.Second
	options.Retryable = graphSnapshotDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &GraphSnapshotDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for GraphSnapshotDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *GraphSnapshotDeletedWaiter) Wait(ctx context.Context, params *GetGraphSnapshotInput, maxWaitDur time.Duration, optFns ...func(*GraphSnapshotDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for GraphSnapshotDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *GraphSnapshotDeletedWaiter) WaitForOutput(ctx context.Context, params *GetGraphSnapshotInput, maxWaitDur time.Duration, optFns ...func(*GraphSnapshotDeletedWaiterOptions)) (*GetGraphSnapshotOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 3600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetGraphSnapshot(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for GraphSnapshotDeleted waiter")
}

func graphSnapshotDeletedStateRetryable(ctx context.Context, input *GetGraphSnapshotInput, output *GetGraphSnapshotOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status != 'DELETING'", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetGraphSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGraphSnapshot",
	}
}
