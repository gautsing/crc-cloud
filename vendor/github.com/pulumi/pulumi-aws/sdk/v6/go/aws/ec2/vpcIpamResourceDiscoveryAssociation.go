// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an association between an Amazon IP Address Manager (IPAM) and a IPAM Resource Discovery. IPAM Resource Discoveries are resources meant for multi-organization customers. If you wish to use a single IPAM across multiple orgs, a resource discovery can be created and shared from a subordinate organization to the management organizations IPAM delegated admin account.
//
// Once an association is created between two organizations via IPAM & a IPAM Resource Discovery, IPAM Pools can be shared via Resource Access Manager (RAM) to accounts in the subordinate organization; these RAM shares must be accepted by the end user account. Pools can then also discover and monitor IPAM resources in the subordinate organization.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpcIpamResourceDiscoveryAssociation(ctx, "test", &ec2.VpcIpamResourceDiscoveryAssociationArgs{
//				IpamId:                  pulumi.Any(aws_vpc_ipam.Test.Id),
//				IpamResourceDiscoveryId: pulumi.Any(aws_vpc_ipam_resource_discovery.Test.Id),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import IPAMs using the IPAM resource discovery association `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpamResourceDiscoveryAssociation:VpcIpamResourceDiscoveryAssociation example ipam-res-disco-assoc-0178368ad2146a492
//
// ```
type VpcIpamResourceDiscoveryAssociation struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of IPAM Resource Discovery Association.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the IPAM.
	IpamArn pulumi.StringOutput `pulumi:"ipamArn"`
	// The ID of the IPAM to associate.
	IpamId pulumi.StringOutput `pulumi:"ipamId"`
	// The home region of the IPAM.
	IpamRegion pulumi.StringOutput `pulumi:"ipamRegion"`
	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId pulumi.StringOutput `pulumi:"ipamResourceDiscoveryId"`
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The account ID for the account that manages the Resource Discovery
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The lifecycle state of the association when you associate or disassociate a resource discovery.
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags to add to the IPAM resource discovery association resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcIpamResourceDiscoveryAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamResourceDiscoveryAssociation(ctx *pulumi.Context,
	name string, args *VpcIpamResourceDiscoveryAssociationArgs, opts ...pulumi.ResourceOption) (*VpcIpamResourceDiscoveryAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpamId == nil {
		return nil, errors.New("invalid value for required argument 'IpamId'")
	}
	if args.IpamResourceDiscoveryId == nil {
		return nil, errors.New("invalid value for required argument 'IpamResourceDiscoveryId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpamResourceDiscoveryAssociation
	err := ctx.RegisterResource("aws:ec2/vpcIpamResourceDiscoveryAssociation:VpcIpamResourceDiscoveryAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamResourceDiscoveryAssociation gets an existing VpcIpamResourceDiscoveryAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamResourceDiscoveryAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamResourceDiscoveryAssociationState, opts ...pulumi.ResourceOption) (*VpcIpamResourceDiscoveryAssociation, error) {
	var resource VpcIpamResourceDiscoveryAssociation
	err := ctx.ReadResource("aws:ec2/vpcIpamResourceDiscoveryAssociation:VpcIpamResourceDiscoveryAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamResourceDiscoveryAssociation resources.
type vpcIpamResourceDiscoveryAssociationState struct {
	// The Amazon Resource Name (ARN) of IPAM Resource Discovery Association.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the IPAM.
	IpamArn *string `pulumi:"ipamArn"`
	// The ID of the IPAM to associate.
	IpamId *string `pulumi:"ipamId"`
	// The home region of the IPAM.
	IpamRegion *string `pulumi:"ipamRegion"`
	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId *string `pulumi:"ipamResourceDiscoveryId"`
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery.
	IsDefault *bool `pulumi:"isDefault"`
	// The account ID for the account that manages the Resource Discovery
	OwnerId *string `pulumi:"ownerId"`
	// The lifecycle state of the association when you associate or disassociate a resource discovery.
	State *string `pulumi:"state"`
	// A map of tags to add to the IPAM resource discovery association resource.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcIpamResourceDiscoveryAssociationState struct {
	// The Amazon Resource Name (ARN) of IPAM Resource Discovery Association.
	Arn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IPAM.
	IpamArn pulumi.StringPtrInput
	// The ID of the IPAM to associate.
	IpamId pulumi.StringPtrInput
	// The home region of the IPAM.
	IpamRegion pulumi.StringPtrInput
	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId pulumi.StringPtrInput
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery.
	IsDefault pulumi.BoolPtrInput
	// The account ID for the account that manages the Resource Discovery
	OwnerId pulumi.StringPtrInput
	// The lifecycle state of the association when you associate or disassociate a resource discovery.
	State pulumi.StringPtrInput
	// A map of tags to add to the IPAM resource discovery association resource.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (VpcIpamResourceDiscoveryAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamResourceDiscoveryAssociationState)(nil)).Elem()
}

type vpcIpamResourceDiscoveryAssociationArgs struct {
	// The ID of the IPAM to associate.
	IpamId string `pulumi:"ipamId"`
	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId string `pulumi:"ipamResourceDiscoveryId"`
	// A map of tags to add to the IPAM resource discovery association resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcIpamResourceDiscoveryAssociation resource.
type VpcIpamResourceDiscoveryAssociationArgs struct {
	// The ID of the IPAM to associate.
	IpamId pulumi.StringInput
	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId pulumi.StringInput
	// A map of tags to add to the IPAM resource discovery association resource.
	Tags pulumi.StringMapInput
}

func (VpcIpamResourceDiscoveryAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamResourceDiscoveryAssociationArgs)(nil)).Elem()
}

type VpcIpamResourceDiscoveryAssociationInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryAssociationOutput() VpcIpamResourceDiscoveryAssociationOutput
	ToVpcIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationOutput
}

func (*VpcIpamResourceDiscoveryAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (i *VpcIpamResourceDiscoveryAssociation) ToVpcIpamResourceDiscoveryAssociationOutput() VpcIpamResourceDiscoveryAssociationOutput {
	return i.ToVpcIpamResourceDiscoveryAssociationOutputWithContext(context.Background())
}

func (i *VpcIpamResourceDiscoveryAssociation) ToVpcIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryAssociationOutput)
}

func (i *VpcIpamResourceDiscoveryAssociation) ToOutput(ctx context.Context) pulumix.Output[*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[*VpcIpamResourceDiscoveryAssociation]{
		OutputState: i.ToVpcIpamResourceDiscoveryAssociationOutputWithContext(ctx).OutputState,
	}
}

// VpcIpamResourceDiscoveryAssociationArrayInput is an input type that accepts VpcIpamResourceDiscoveryAssociationArray and VpcIpamResourceDiscoveryAssociationArrayOutput values.
// You can construct a concrete instance of `VpcIpamResourceDiscoveryAssociationArrayInput` via:
//
//	VpcIpamResourceDiscoveryAssociationArray{ VpcIpamResourceDiscoveryAssociationArgs{...} }
type VpcIpamResourceDiscoveryAssociationArrayInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryAssociationArrayOutput() VpcIpamResourceDiscoveryAssociationArrayOutput
	ToVpcIpamResourceDiscoveryAssociationArrayOutputWithContext(context.Context) VpcIpamResourceDiscoveryAssociationArrayOutput
}

type VpcIpamResourceDiscoveryAssociationArray []VpcIpamResourceDiscoveryAssociationInput

func (VpcIpamResourceDiscoveryAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (i VpcIpamResourceDiscoveryAssociationArray) ToVpcIpamResourceDiscoveryAssociationArrayOutput() VpcIpamResourceDiscoveryAssociationArrayOutput {
	return i.ToVpcIpamResourceDiscoveryAssociationArrayOutputWithContext(context.Background())
}

func (i VpcIpamResourceDiscoveryAssociationArray) ToVpcIpamResourceDiscoveryAssociationArrayOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryAssociationArrayOutput)
}

func (i VpcIpamResourceDiscoveryAssociationArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[[]*VpcIpamResourceDiscoveryAssociation]{
		OutputState: i.ToVpcIpamResourceDiscoveryAssociationArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcIpamResourceDiscoveryAssociationMapInput is an input type that accepts VpcIpamResourceDiscoveryAssociationMap and VpcIpamResourceDiscoveryAssociationMapOutput values.
// You can construct a concrete instance of `VpcIpamResourceDiscoveryAssociationMapInput` via:
//
//	VpcIpamResourceDiscoveryAssociationMap{ "key": VpcIpamResourceDiscoveryAssociationArgs{...} }
type VpcIpamResourceDiscoveryAssociationMapInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryAssociationMapOutput() VpcIpamResourceDiscoveryAssociationMapOutput
	ToVpcIpamResourceDiscoveryAssociationMapOutputWithContext(context.Context) VpcIpamResourceDiscoveryAssociationMapOutput
}

type VpcIpamResourceDiscoveryAssociationMap map[string]VpcIpamResourceDiscoveryAssociationInput

func (VpcIpamResourceDiscoveryAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (i VpcIpamResourceDiscoveryAssociationMap) ToVpcIpamResourceDiscoveryAssociationMapOutput() VpcIpamResourceDiscoveryAssociationMapOutput {
	return i.ToVpcIpamResourceDiscoveryAssociationMapOutputWithContext(context.Background())
}

func (i VpcIpamResourceDiscoveryAssociationMap) ToVpcIpamResourceDiscoveryAssociationMapOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryAssociationMapOutput)
}

func (i VpcIpamResourceDiscoveryAssociationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[map[string]*VpcIpamResourceDiscoveryAssociation]{
		OutputState: i.ToVpcIpamResourceDiscoveryAssociationMapOutputWithContext(ctx).OutputState,
	}
}

type VpcIpamResourceDiscoveryAssociationOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryAssociationOutput) ToVpcIpamResourceDiscoveryAssociationOutput() VpcIpamResourceDiscoveryAssociationOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationOutput) ToVpcIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[*VpcIpamResourceDiscoveryAssociation]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of IPAM Resource Discovery Association.
func (o VpcIpamResourceDiscoveryAssociationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the IPAM.
func (o VpcIpamResourceDiscoveryAssociationOutput) IpamArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamArn }).(pulumi.StringOutput)
}

// The ID of the IPAM to associate.
func (o VpcIpamResourceDiscoveryAssociationOutput) IpamId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamId }).(pulumi.StringOutput)
}

// The home region of the IPAM.
func (o VpcIpamResourceDiscoveryAssociationOutput) IpamRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamRegion }).(pulumi.StringOutput)
}

// The ID of the Resource Discovery to associate.
func (o VpcIpamResourceDiscoveryAssociationOutput) IpamResourceDiscoveryId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamResourceDiscoveryId }).(pulumi.StringOutput)
}

// A boolean to identify if the Resource Discovery is the accounts default resource discovery.
func (o VpcIpamResourceDiscoveryAssociationOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The account ID for the account that manages the Resource Discovery
func (o VpcIpamResourceDiscoveryAssociationOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The lifecycle state of the association when you associate or disassociate a resource discovery.
func (o VpcIpamResourceDiscoveryAssociationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A map of tags to add to the IPAM resource discovery association resource.
func (o VpcIpamResourceDiscoveryAssociationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VpcIpamResourceDiscoveryAssociationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscoveryAssociation) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcIpamResourceDiscoveryAssociationArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryAssociationArrayOutput) ToVpcIpamResourceDiscoveryAssociationArrayOutput() VpcIpamResourceDiscoveryAssociationArrayOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationArrayOutput) ToVpcIpamResourceDiscoveryAssociationArrayOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationArrayOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[[]*VpcIpamResourceDiscoveryAssociation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpamResourceDiscoveryAssociationArrayOutput) Index(i pulumi.IntInput) VpcIpamResourceDiscoveryAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamResourceDiscoveryAssociation {
		return vs[0].([]*VpcIpamResourceDiscoveryAssociation)[vs[1].(int)]
	}).(VpcIpamResourceDiscoveryAssociationOutput)
}

type VpcIpamResourceDiscoveryAssociationMapOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryAssociationMapOutput) ToVpcIpamResourceDiscoveryAssociationMapOutput() VpcIpamResourceDiscoveryAssociationMapOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationMapOutput) ToVpcIpamResourceDiscoveryAssociationMapOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryAssociationMapOutput {
	return o
}

func (o VpcIpamResourceDiscoveryAssociationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpamResourceDiscoveryAssociation] {
	return pulumix.Output[map[string]*VpcIpamResourceDiscoveryAssociation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpamResourceDiscoveryAssociationMapOutput) MapIndex(k pulumi.StringInput) VpcIpamResourceDiscoveryAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamResourceDiscoveryAssociation {
		return vs[0].(map[string]*VpcIpamResourceDiscoveryAssociation)[vs[1].(string)]
	}).(VpcIpamResourceDiscoveryAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryAssociationInput)(nil)).Elem(), &VpcIpamResourceDiscoveryAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryAssociationArrayInput)(nil)).Elem(), VpcIpamResourceDiscoveryAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryAssociationMapInput)(nil)).Elem(), VpcIpamResourceDiscoveryAssociationMap{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryAssociationOutput{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryAssociationArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryAssociationMapOutput{})
}
