// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Packer Image data source iteration gets the most recent iteration (or build) of an image, given an iteration id.
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
// 		hardened_source, err := hcp.GetPackerIteration(ctx, &GetPackerIterationArgs{
// 			BucketName: "hardened-ubuntu-16-04",
// 			Channel:    "production",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := hcp.GetPackerImage(ctx, &GetPackerImageArgs{
// 			BucketName:    "hardened-ubuntu-16-04",
// 			CloudProvider: "aws",
// 			IterationId:   hardened_source.Ulid,
// 			Region:        "us-east-1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("packer-registry-ubuntu", foo.CloudImageId)
// 		return nil
// 	})
// }
// ```
func GetPackerImage(ctx *pulumi.Context, args *GetPackerImageArgs, opts ...pulumi.InvokeOption) (*GetPackerImageResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPackerImageResult
	err := ctx.Invoke("hcp:index/getPackerImage:getPackerImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPackerImage.
type GetPackerImageArgs struct {
	// The slug of the HCP Packer Registry image bucket to pull from.
	BucketName string `pulumi:"bucketName"`
	// Name of the cloud provider this image is stored-in.
	CloudProvider string `pulumi:"cloudProvider"`
	// HCP ID of this image.
	IterationId string `pulumi:"iterationId"`
	// Region this image is stored in, if any.
	Region string `pulumi:"region"`
}

// A collection of values returned by getPackerImage.
type GetPackerImageResult struct {
	// The slug of the HCP Packer Registry image bucket to pull from.
	BucketName string `pulumi:"bucketName"`
	// HCP ID of this build.
	BuildId string `pulumi:"buildId"`
	// Cloud Image ID or URL string identifying this image for the builder that built it.
	CloudImageId string `pulumi:"cloudImageId"`
	// Name of the cloud provider this image is stored-in.
	CloudProvider string `pulumi:"cloudProvider"`
	// Name of the builder that built this. Ex: 'amazon-ebs.example'
	ComponentType string `pulumi:"componentType"`
	// Creation time of this build.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// HCP ID of this image.
	IterationId string `pulumi:"iterationId"`
	// Labels associated with this build.
	Labels map[string]interface{} `pulumi:"labels"`
	// The ID of the organization this HCP Packer registry is located in.
	OrganizationId string `pulumi:"organizationId"`
	// UUID of this build.
	PackerRunUuid string `pulumi:"packerRunUuid"`
	// The ID of the project this HCP Packer registry is located in.
	ProjectId string `pulumi:"projectId"`
	// Region this image is stored in, if any.
	Region string `pulumi:"region"`
}

func GetPackerImageOutput(ctx *pulumi.Context, args GetPackerImageOutputArgs, opts ...pulumi.InvokeOption) GetPackerImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPackerImageResult, error) {
			args := v.(GetPackerImageArgs)
			r, err := GetPackerImage(ctx, &args, opts...)
			return *r, err
		}).(GetPackerImageResultOutput)
}

// A collection of arguments for invoking getPackerImage.
type GetPackerImageOutputArgs struct {
	// The slug of the HCP Packer Registry image bucket to pull from.
	BucketName pulumi.StringInput `pulumi:"bucketName"`
	// Name of the cloud provider this image is stored-in.
	CloudProvider pulumi.StringInput `pulumi:"cloudProvider"`
	// HCP ID of this image.
	IterationId pulumi.StringInput `pulumi:"iterationId"`
	// Region this image is stored in, if any.
	Region pulumi.StringInput `pulumi:"region"`
}

func (GetPackerImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPackerImageArgs)(nil)).Elem()
}

// A collection of values returned by getPackerImage.
type GetPackerImageResultOutput struct{ *pulumi.OutputState }

func (GetPackerImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPackerImageResult)(nil)).Elem()
}

func (o GetPackerImageResultOutput) ToGetPackerImageResultOutput() GetPackerImageResultOutput {
	return o
}

func (o GetPackerImageResultOutput) ToGetPackerImageResultOutputWithContext(ctx context.Context) GetPackerImageResultOutput {
	return o
}

// The slug of the HCP Packer Registry image bucket to pull from.
func (o GetPackerImageResultOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.BucketName }).(pulumi.StringOutput)
}

// HCP ID of this build.
func (o GetPackerImageResultOutput) BuildId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.BuildId }).(pulumi.StringOutput)
}

// Cloud Image ID or URL string identifying this image for the builder that built it.
func (o GetPackerImageResultOutput) CloudImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.CloudImageId }).(pulumi.StringOutput)
}

// Name of the cloud provider this image is stored-in.
func (o GetPackerImageResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

// Name of the builder that built this. Ex: 'amazon-ebs.example'
func (o GetPackerImageResultOutput) ComponentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.ComponentType }).(pulumi.StringOutput)
}

// Creation time of this build.
func (o GetPackerImageResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPackerImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// HCP ID of this image.
func (o GetPackerImageResultOutput) IterationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.IterationId }).(pulumi.StringOutput)
}

// Labels associated with this build.
func (o GetPackerImageResultOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v GetPackerImageResult) map[string]interface{} { return v.Labels }).(pulumi.MapOutput)
}

// The ID of the organization this HCP Packer registry is located in.
func (o GetPackerImageResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// UUID of this build.
func (o GetPackerImageResultOutput) PackerRunUuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.PackerRunUuid }).(pulumi.StringOutput)
}

// The ID of the project this HCP Packer registry is located in.
func (o GetPackerImageResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Region this image is stored in, if any.
func (o GetPackerImageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackerImageResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPackerImageResultOutput{})
}
