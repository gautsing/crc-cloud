// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a global forwarding rule within GCE from its name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.LookupGlobalForwardingRule(ctx, &compute.LookupGlobalForwardingRuleArgs{
//				Name: "forwarding-rule-global",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGlobalForwardingRule(ctx *pulumi.Context, args *LookupGlobalForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupGlobalForwardingRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalForwardingRuleResult
	err := ctx.Invoke("gcp:compute/getGlobalForwardingRule:getGlobalForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGlobalForwardingRule.
type LookupGlobalForwardingRuleArgs struct {
	// The name of the global forwarding rule.
	//
	// ***
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getGlobalForwardingRule.
type LookupGlobalForwardingRuleResult struct {
	AllowPscGlobalAccess bool   `pulumi:"allowPscGlobalAccess"`
	BaseForwardingRule   string `pulumi:"baseForwardingRule"`
	Description          string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                                  `pulumi:"id"`
	IpAddress           string                                  `pulumi:"ipAddress"`
	IpProtocol          string                                  `pulumi:"ipProtocol"`
	IpVersion           string                                  `pulumi:"ipVersion"`
	LabelFingerprint    string                                  `pulumi:"labelFingerprint"`
	Labels              map[string]string                       `pulumi:"labels"`
	LoadBalancingScheme string                                  `pulumi:"loadBalancingScheme"`
	MetadataFilters     []GetGlobalForwardingRuleMetadataFilter `pulumi:"metadataFilters"`
	Name                string                                  `pulumi:"name"`
	Network             string                                  `pulumi:"network"`
	NoAutomateDnsZone   bool                                    `pulumi:"noAutomateDnsZone"`
	PortRange           string                                  `pulumi:"portRange"`
	Project             *string                                 `pulumi:"project"`
	PscConnectionId     string                                  `pulumi:"pscConnectionId"`
	PscConnectionStatus string                                  `pulumi:"pscConnectionStatus"`
	SelfLink            string                                  `pulumi:"selfLink"`
	SourceIpRanges      []string                                `pulumi:"sourceIpRanges"`
	Subnetwork          string                                  `pulumi:"subnetwork"`
	Target              string                                  `pulumi:"target"`
}

func LookupGlobalForwardingRuleOutput(ctx *pulumi.Context, args LookupGlobalForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalForwardingRuleResult, error) {
			args := v.(LookupGlobalForwardingRuleArgs)
			r, err := LookupGlobalForwardingRule(ctx, &args, opts...)
			var s LookupGlobalForwardingRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalForwardingRuleResultOutput)
}

// A collection of arguments for invoking getGlobalForwardingRule.
type LookupGlobalForwardingRuleOutputArgs struct {
	// The name of the global forwarding rule.
	//
	// ***
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGlobalForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalForwardingRuleArgs)(nil)).Elem()
}

// A collection of values returned by getGlobalForwardingRule.
type LookupGlobalForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalForwardingRuleResult)(nil)).Elem()
}

func (o LookupGlobalForwardingRuleResultOutput) ToLookupGlobalForwardingRuleResultOutput() LookupGlobalForwardingRuleResultOutput {
	return o
}

func (o LookupGlobalForwardingRuleResultOutput) ToLookupGlobalForwardingRuleResultOutputWithContext(ctx context.Context) LookupGlobalForwardingRuleResultOutput {
	return o
}

func (o LookupGlobalForwardingRuleResultOutput) AllowPscGlobalAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) bool { return v.AllowPscGlobalAccess }).(pulumi.BoolOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) BaseForwardingRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.BaseForwardingRule }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGlobalForwardingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpProtocol }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) LoadBalancingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.LoadBalancingScheme }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) MetadataFilters() GetGlobalForwardingRuleMetadataFilterArrayOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) []GetGlobalForwardingRuleMetadataFilter {
		return v.MetadataFilters
	}).(GetGlobalForwardingRuleMetadataFilterArrayOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) NoAutomateDnsZone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) bool { return v.NoAutomateDnsZone }).(pulumi.BoolOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) PortRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PortRange }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) PscConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PscConnectionId }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) PscConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PscConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) SourceIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) []string { return v.SourceIpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Target }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalForwardingRuleResultOutput{})
}