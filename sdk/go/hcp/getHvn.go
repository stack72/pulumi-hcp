// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The HVN data source provides information about an existing HashiCorp Virtual Network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/grapl-security/pulumi-hcp/sdk/go/hcp"
// 	"github.com/pulumi/pulumi-hcp/sdk/go/hcp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hcp.LookupHvn(ctx, &GetHvnArgs{
// 			HvnId: _var.Hvn_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupHvn(ctx *pulumi.Context, args *LookupHvnArgs, opts ...pulumi.InvokeOption) (*LookupHvnResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupHvnResult
	err := ctx.Invoke("hcp:index/getHvn:getHvn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHvn.
type LookupHvnArgs struct {
	// The ID of the HashiCorp Virtual Network (HVN).
	HvnId string `pulumi:"hvnId"`
}

// A collection of values returned by getHvn.
type LookupHvnResult struct {
	// The CIDR range of the HVN.
	CidrBlock string `pulumi:"cidrBlock"`
	// The provider where the HVN is located.
	CloudProvider string `pulumi:"cloudProvider"`
	// The time that the HVN was created.
	CreatedAt string `pulumi:"createdAt"`
	// The ID of the HashiCorp Virtual Network (HVN).
	HvnId string `pulumi:"hvnId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the HCP organization where the HVN is located.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the HCP project where the HVN is located.
	ProjectId string `pulumi:"projectId"`
	// The provider account ID where the HVN is located.
	ProviderAccountId string `pulumi:"providerAccountId"`
	// The region where the HVN is located.
	Region string `pulumi:"region"`
	// A unique URL identifying the HVN.
	SelfLink string `pulumi:"selfLink"`
}

func LookupHvnOutput(ctx *pulumi.Context, args LookupHvnOutputArgs, opts ...pulumi.InvokeOption) LookupHvnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHvnResult, error) {
			args := v.(LookupHvnArgs)
			r, err := LookupHvn(ctx, &args, opts...)
			return *r, err
		}).(LookupHvnResultOutput)
}

// A collection of arguments for invoking getHvn.
type LookupHvnOutputArgs struct {
	// The ID of the HashiCorp Virtual Network (HVN).
	HvnId pulumi.StringInput `pulumi:"hvnId"`
}

func (LookupHvnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHvnArgs)(nil)).Elem()
}

// A collection of values returned by getHvn.
type LookupHvnResultOutput struct{ *pulumi.OutputState }

func (LookupHvnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHvnResult)(nil)).Elem()
}

func (o LookupHvnResultOutput) ToLookupHvnResultOutput() LookupHvnResultOutput {
	return o
}

func (o LookupHvnResultOutput) ToLookupHvnResultOutputWithContext(ctx context.Context) LookupHvnResultOutput {
	return o
}

// The CIDR range of the HVN.
func (o LookupHvnResultOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// The provider where the HVN is located.
func (o LookupHvnResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

// The time that the HVN was created.
func (o LookupHvnResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the HashiCorp Virtual Network (HVN).
func (o LookupHvnResultOutput) HvnId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.HvnId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupHvnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the HCP organization where the HVN is located.
func (o LookupHvnResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the HCP project where the HVN is located.
func (o LookupHvnResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The provider account ID where the HVN is located.
func (o LookupHvnResultOutput) ProviderAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.ProviderAccountId }).(pulumi.StringOutput)
}

// The region where the HVN is located.
func (o LookupHvnResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.Region }).(pulumi.StringOutput)
}

// A unique URL identifying the HVN.
func (o LookupHvnResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHvnResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHvnResultOutput{})
}
