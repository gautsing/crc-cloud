// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "openstack:images/image:Image":
		r = &Image{}
	case "openstack:images/imageAccess:ImageAccess":
		r = &ImageAccess{}
	case "openstack:images/imageAccessAccept:ImageAccessAccept":
		r = &ImageAccessAccept{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"openstack",
		"images/image",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"images/imageAccess",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"images/imageAccessAccept",
		&module{version},
	)
}
