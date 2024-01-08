// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified firewall rule.
func (c *Client) UpdateFirewallRule(ctx context.Context, params *UpdateFirewallRuleInput, optFns ...func(*Options)) (*UpdateFirewallRuleOutput, error) {
	if params == nil {
		params = &UpdateFirewallRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallRule", params, optFns, c.addOperationUpdateFirewallRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallRuleInput struct {

	// The ID of the domain list to use in the rule.
	//
	// This member is required.
	FirewallDomainListId *string

	// The unique identifier of the firewall rule group for the rule.
	//
	// This member is required.
	FirewallRuleGroupId *string

	// The action that DNS Firewall should take on a DNS query when it matches one of
	// the domains in the rule's domain list:
	//   - ALLOW - Permit the request to go through.
	//   - ALERT - Permit the request to go through but send an alert to the logs.
	//   - BLOCK - Disallow the request. This option requires additional details in the
	//   rule's BlockResponse .
	Action types.Action

	// The DNS record's type. This determines the format of the record value that you
	// provided in BlockOverrideDomain . Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE .
	BlockOverrideDnsType types.BlockOverrideDnsType

	// The custom DNS record to send back in response to the query. Used for the rule
	// action BLOCK with a BlockResponse setting of OVERRIDE .
	BlockOverrideDomain *string

	// The recommended amount of time, in seconds, for the DNS resolver or web browser
	// to cache the provided override record. Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE .
	BlockOverrideTtl *int32

	// The way that you want DNS Firewall to block the request. Used for the rule
	// action setting BLOCK .
	//   - NODATA - Respond indicating that the query was successful, but no response
	//   is available for it.
	//   - NXDOMAIN - Respond indicating that the domain name that's in the query
	//   doesn't exist.
	//   - OVERRIDE - Provide a custom override in the response. This option requires
	//   custom handling details in the rule's BlockOverride* settings.
	BlockResponse types.BlockResponse

	// The name of the rule.
	Name *string

	// The setting that determines the processing order of the rule in the rule group.
	// DNS Firewall processes the rules in a rule group by order of priority, starting
	// from the lowest setting. You must specify a unique priority for each rule in a
	// rule group. To make it easier to insert rules later, leave space between the
	// numbers, for example, use 100, 200, and so on. You can change the priority
	// setting for the rules in a rule group at any time.
	Priority *int32

	// The DNS query type you want the rule to evaluate. Allowed values are;
	//   - A: Returns an IPv4 address.
	//   - AAAA: Returns an Ipv6 address.
	//   - CAA: Restricts CAs that can create SSL/TLS certifications for the domain.
	//   - CNAME: Returns another domain name.
	//   - DS: Record that identifies the DNSSEC signing key of a delegated zone.
	//   - MX: Specifies mail servers.
	//   - NAPTR: Regular-expression-based rewriting of domain names.
	//   - NS: Authoritative name servers.
	//   - PTR: Maps an IP address to a domain name.
	//   - SOA: Start of authority record for the zone.
	//   - SPF: Lists the servers authorized to send emails from a domain.
	//   - SRV: Application specific values that identify servers.
	//   - TXT: Verifies email senders and application-specific values.
	Qtype *string

	noSmithyDocumentSerde
}

type UpdateFirewallRuleOutput struct {

	// The firewall rule that you just updated.
	FirewallRule *types.FirewallRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFirewallRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFirewallRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFirewallRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFirewallRule"); err != nil {
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
	if err = addOpUpdateFirewallRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFirewallRule",
	}
}
