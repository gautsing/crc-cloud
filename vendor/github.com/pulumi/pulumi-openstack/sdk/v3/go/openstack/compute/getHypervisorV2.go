// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about hypervisors
// by hostname.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.GetHypervisorV2(ctx, &compute.GetHypervisorV2Args{
//				Hostname: "host01",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetHypervisorV2(ctx *pulumi.Context, args *GetHypervisorV2Args, opts ...pulumi.InvokeOption) (*GetHypervisorV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHypervisorV2Result
	err := ctx.Invoke("openstack:compute/getHypervisorV2:getHypervisorV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHypervisorV2.
type GetHypervisorV2Args struct {
	// The hostname of the hypervisor
	Hostname string `pulumi:"hostname"`
}

// A collection of values returned by getHypervisorV2.
type GetHypervisorV2Result struct {
	// The amount in GigaBytes of local storage the hypervisor can provide
	Disk int `pulumi:"disk"`
	// The IP address of the Hypervisor
	HostIp string `pulumi:"hostIp"`
	// See Argument Reference above.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The number in MegaBytes of memory the hypervisor can provide
	Memory int `pulumi:"memory"`
	// The state of the hypervisor (`up` or `down`)
	State string `pulumi:"state"`
	// The status of the hypervisor (`enabled` or `disabled`)
	Status string `pulumi:"status"`
	// The type of the hypervisor (example: `QEMU`)
	Type string `pulumi:"type"`
	// The number of virtual CPUs the hypervisor can provide
	Vcpus int `pulumi:"vcpus"`
}

func GetHypervisorV2Output(ctx *pulumi.Context, args GetHypervisorV2OutputArgs, opts ...pulumi.InvokeOption) GetHypervisorV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHypervisorV2Result, error) {
			args := v.(GetHypervisorV2Args)
			r, err := GetHypervisorV2(ctx, &args, opts...)
			var s GetHypervisorV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHypervisorV2ResultOutput)
}

// A collection of arguments for invoking getHypervisorV2.
type GetHypervisorV2OutputArgs struct {
	// The hostname of the hypervisor
	Hostname pulumi.StringInput `pulumi:"hostname"`
}

func (GetHypervisorV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHypervisorV2Args)(nil)).Elem()
}

// A collection of values returned by getHypervisorV2.
type GetHypervisorV2ResultOutput struct{ *pulumi.OutputState }

func (GetHypervisorV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHypervisorV2Result)(nil)).Elem()
}

func (o GetHypervisorV2ResultOutput) ToGetHypervisorV2ResultOutput() GetHypervisorV2ResultOutput {
	return o
}

func (o GetHypervisorV2ResultOutput) ToGetHypervisorV2ResultOutputWithContext(ctx context.Context) GetHypervisorV2ResultOutput {
	return o
}

// The amount in GigaBytes of local storage the hypervisor can provide
func (o GetHypervisorV2ResultOutput) Disk() pulumi.IntOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) int { return v.Disk }).(pulumi.IntOutput)
}

// The IP address of the Hypervisor
func (o GetHypervisorV2ResultOutput) HostIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.HostIp }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetHypervisorV2ResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHypervisorV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// The number in MegaBytes of memory the hypervisor can provide
func (o GetHypervisorV2ResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) int { return v.Memory }).(pulumi.IntOutput)
}

// The state of the hypervisor (`up` or `down`)
func (o GetHypervisorV2ResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.State }).(pulumi.StringOutput)
}

// The status of the hypervisor (`enabled` or `disabled`)
func (o GetHypervisorV2ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.Status }).(pulumi.StringOutput)
}

// The type of the hypervisor (example: `QEMU`)
func (o GetHypervisorV2ResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) string { return v.Type }).(pulumi.StringOutput)
}

// The number of virtual CPUs the hypervisor can provide
func (o GetHypervisorV2ResultOutput) Vcpus() pulumi.IntOutput {
	return o.ApplyT(func(v GetHypervisorV2Result) int { return v.Vcpus }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHypervisorV2ResultOutput{})
}