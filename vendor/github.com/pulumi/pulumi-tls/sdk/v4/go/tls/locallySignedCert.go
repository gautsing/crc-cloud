// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type LocallySignedCert struct {
	pulumi.CustomResourceState

	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayOutput `pulumi:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringOutput `pulumi:"caCertPem"`
	// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
	CaKeyAlgorithm pulumi.StringOutput `pulumi:"caKeyAlgorithm"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem pulumi.StringOutput `pulumi:"caPrivateKeyPem"`
	// Certificate data in PEM (RFC 1421).
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringOutput `pulumi:"certRequestPem"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours pulumi.IntPtrOutput `pulumi:"earlyRenewalHours"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrOutput `pulumi:"isCaCertificate"`
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal pulumi.BoolOutput `pulumi:"readyForRenewal"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrOutput `pulumi:"setSubjectKeyId"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringOutput `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringOutput `pulumi:"validityStartTime"`
}

// NewLocallySignedCert registers a new resource with the given unique name, arguments, and options.
func NewLocallySignedCert(ctx *pulumi.Context,
	name string, args *LocallySignedCertArgs, opts ...pulumi.ResourceOption) (*LocallySignedCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedUses == nil {
		return nil, errors.New("invalid value for required argument 'AllowedUses'")
	}
	if args.CaCertPem == nil {
		return nil, errors.New("invalid value for required argument 'CaCertPem'")
	}
	if args.CaPrivateKeyPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPrivateKeyPem'")
	}
	if args.CertRequestPem == nil {
		return nil, errors.New("invalid value for required argument 'CertRequestPem'")
	}
	if args.ValidityPeriodHours == nil {
		return nil, errors.New("invalid value for required argument 'ValidityPeriodHours'")
	}
	if args.CaPrivateKeyPem != nil {
		args.CaPrivateKeyPem = pulumi.ToSecret(args.CaPrivateKeyPem).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"caPrivateKeyPem",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocallySignedCert
	err := ctx.RegisterResource("tls:index/locallySignedCert:LocallySignedCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocallySignedCert gets an existing LocallySignedCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocallySignedCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocallySignedCertState, opts ...pulumi.ResourceOption) (*LocallySignedCert, error) {
	var resource LocallySignedCert
	err := ctx.ReadResource("tls:index/locallySignedCert:LocallySignedCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocallySignedCert resources.
type locallySignedCertState struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses []string `pulumi:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem *string `pulumi:"caCertPem"`
	// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
	CaKeyAlgorithm *string `pulumi:"caKeyAlgorithm"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem *string `pulumi:"caPrivateKeyPem"`
	// Certificate data in PEM (RFC 1421).
	CertPem *string `pulumi:"certPem"`
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem *string `pulumi:"certRequestPem"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal *bool `pulumi:"readyForRenewal"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime *string `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours *int `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime *string `pulumi:"validityStartTime"`
}

type LocallySignedCertState struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayInput
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringPtrInput
	// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
	CaKeyAlgorithm pulumi.StringPtrInput
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem pulumi.StringPtrInput
	// Certificate data in PEM (RFC 1421).
	CertPem pulumi.StringPtrInput
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringPtrInput
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours pulumi.IntPtrInput
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrInput
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal pulumi.BoolPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringPtrInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntPtrInput
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringPtrInput
}

func (LocallySignedCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*locallySignedCertState)(nil)).Elem()
}

type locallySignedCertArgs struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses []string `pulumi:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem string `pulumi:"caCertPem"`
	// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
	CaKeyAlgorithm *string `pulumi:"caKeyAlgorithm"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem string `pulumi:"caPrivateKeyPem"`
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem string `pulumi:"certRequestPem"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a LocallySignedCert resource.
type LocallySignedCertArgs struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayInput
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringInput
	// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
	CaKeyAlgorithm pulumi.StringPtrInput
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem pulumi.StringInput
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringInput
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours pulumi.IntPtrInput
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntInput
}

func (LocallySignedCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locallySignedCertArgs)(nil)).Elem()
}

type LocallySignedCertInput interface {
	pulumi.Input

	ToLocallySignedCertOutput() LocallySignedCertOutput
	ToLocallySignedCertOutputWithContext(ctx context.Context) LocallySignedCertOutput
}

func (*LocallySignedCert) ElementType() reflect.Type {
	return reflect.TypeOf((**LocallySignedCert)(nil)).Elem()
}

func (i *LocallySignedCert) ToLocallySignedCertOutput() LocallySignedCertOutput {
	return i.ToLocallySignedCertOutputWithContext(context.Background())
}

func (i *LocallySignedCert) ToLocallySignedCertOutputWithContext(ctx context.Context) LocallySignedCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocallySignedCertOutput)
}

func (i *LocallySignedCert) ToOutput(ctx context.Context) pulumix.Output[*LocallySignedCert] {
	return pulumix.Output[*LocallySignedCert]{
		OutputState: i.ToLocallySignedCertOutputWithContext(ctx).OutputState,
	}
}

// LocallySignedCertArrayInput is an input type that accepts LocallySignedCertArray and LocallySignedCertArrayOutput values.
// You can construct a concrete instance of `LocallySignedCertArrayInput` via:
//
//	LocallySignedCertArray{ LocallySignedCertArgs{...} }
type LocallySignedCertArrayInput interface {
	pulumi.Input

	ToLocallySignedCertArrayOutput() LocallySignedCertArrayOutput
	ToLocallySignedCertArrayOutputWithContext(context.Context) LocallySignedCertArrayOutput
}

type LocallySignedCertArray []LocallySignedCertInput

func (LocallySignedCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocallySignedCert)(nil)).Elem()
}

func (i LocallySignedCertArray) ToLocallySignedCertArrayOutput() LocallySignedCertArrayOutput {
	return i.ToLocallySignedCertArrayOutputWithContext(context.Background())
}

func (i LocallySignedCertArray) ToLocallySignedCertArrayOutputWithContext(ctx context.Context) LocallySignedCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocallySignedCertArrayOutput)
}

func (i LocallySignedCertArray) ToOutput(ctx context.Context) pulumix.Output[[]*LocallySignedCert] {
	return pulumix.Output[[]*LocallySignedCert]{
		OutputState: i.ToLocallySignedCertArrayOutputWithContext(ctx).OutputState,
	}
}

// LocallySignedCertMapInput is an input type that accepts LocallySignedCertMap and LocallySignedCertMapOutput values.
// You can construct a concrete instance of `LocallySignedCertMapInput` via:
//
//	LocallySignedCertMap{ "key": LocallySignedCertArgs{...} }
type LocallySignedCertMapInput interface {
	pulumi.Input

	ToLocallySignedCertMapOutput() LocallySignedCertMapOutput
	ToLocallySignedCertMapOutputWithContext(context.Context) LocallySignedCertMapOutput
}

type LocallySignedCertMap map[string]LocallySignedCertInput

func (LocallySignedCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocallySignedCert)(nil)).Elem()
}

func (i LocallySignedCertMap) ToLocallySignedCertMapOutput() LocallySignedCertMapOutput {
	return i.ToLocallySignedCertMapOutputWithContext(context.Background())
}

func (i LocallySignedCertMap) ToLocallySignedCertMapOutputWithContext(ctx context.Context) LocallySignedCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocallySignedCertMapOutput)
}

func (i LocallySignedCertMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocallySignedCert] {
	return pulumix.Output[map[string]*LocallySignedCert]{
		OutputState: i.ToLocallySignedCertMapOutputWithContext(ctx).OutputState,
	}
}

type LocallySignedCertOutput struct{ *pulumi.OutputState }

func (LocallySignedCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocallySignedCert)(nil)).Elem()
}

func (o LocallySignedCertOutput) ToLocallySignedCertOutput() LocallySignedCertOutput {
	return o
}

func (o LocallySignedCertOutput) ToLocallySignedCertOutputWithContext(ctx context.Context) LocallySignedCertOutput {
	return o
}

func (o LocallySignedCertOutput) ToOutput(ctx context.Context) pulumix.Output[*LocallySignedCert] {
	return pulumix.Output[*LocallySignedCert]{
		OutputState: o.OutputState,
	}
}

// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
func (o LocallySignedCertOutput) AllowedUses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringArrayOutput { return v.AllowedUses }).(pulumi.StringArrayOutput)
}

// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o LocallySignedCertOutput) CaCertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.CaCertPem }).(pulumi.StringOutput)
}

// Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
//
// Deprecated: This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
func (o LocallySignedCertOutput) CaKeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.CaKeyAlgorithm }).(pulumi.StringOutput)
}

// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o LocallySignedCertOutput) CaPrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.CaPrivateKeyPem }).(pulumi.StringOutput)
}

// Certificate data in PEM (RFC 1421).
func (o LocallySignedCertOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o LocallySignedCertOutput) CertRequestPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.CertRequestPem }).(pulumi.StringOutput)
}

// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
// early renewal period. (default: `0`)
func (o LocallySignedCertOutput) EarlyRenewalHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.IntPtrOutput { return v.EarlyRenewalHours }).(pulumi.IntPtrOutput)
}

// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
func (o LocallySignedCertOutput) IsCaCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.BoolPtrOutput { return v.IsCaCertificate }).(pulumi.BoolPtrOutput)
}

// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
func (o LocallySignedCertOutput) ReadyForRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.BoolOutput { return v.ReadyForRenewal }).(pulumi.BoolOutput)
}

// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o LocallySignedCertOutput) SetSubjectKeyId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.BoolPtrOutput { return v.SetSubjectKeyId }).(pulumi.BoolPtrOutput)
}

// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o LocallySignedCertOutput) ValidityEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.ValidityEndTime }).(pulumi.StringOutput)
}

// Number of hours, after initial issuing, that the certificate will remain valid for.
func (o LocallySignedCertOutput) ValidityPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.IntOutput { return v.ValidityPeriodHours }).(pulumi.IntOutput)
}

// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o LocallySignedCertOutput) ValidityStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LocallySignedCert) pulumi.StringOutput { return v.ValidityStartTime }).(pulumi.StringOutput)
}

type LocallySignedCertArrayOutput struct{ *pulumi.OutputState }

func (LocallySignedCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocallySignedCert)(nil)).Elem()
}

func (o LocallySignedCertArrayOutput) ToLocallySignedCertArrayOutput() LocallySignedCertArrayOutput {
	return o
}

func (o LocallySignedCertArrayOutput) ToLocallySignedCertArrayOutputWithContext(ctx context.Context) LocallySignedCertArrayOutput {
	return o
}

func (o LocallySignedCertArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LocallySignedCert] {
	return pulumix.Output[[]*LocallySignedCert]{
		OutputState: o.OutputState,
	}
}

func (o LocallySignedCertArrayOutput) Index(i pulumi.IntInput) LocallySignedCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocallySignedCert {
		return vs[0].([]*LocallySignedCert)[vs[1].(int)]
	}).(LocallySignedCertOutput)
}

type LocallySignedCertMapOutput struct{ *pulumi.OutputState }

func (LocallySignedCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocallySignedCert)(nil)).Elem()
}

func (o LocallySignedCertMapOutput) ToLocallySignedCertMapOutput() LocallySignedCertMapOutput {
	return o
}

func (o LocallySignedCertMapOutput) ToLocallySignedCertMapOutputWithContext(ctx context.Context) LocallySignedCertMapOutput {
	return o
}

func (o LocallySignedCertMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocallySignedCert] {
	return pulumix.Output[map[string]*LocallySignedCert]{
		OutputState: o.OutputState,
	}
}

func (o LocallySignedCertMapOutput) MapIndex(k pulumi.StringInput) LocallySignedCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocallySignedCert {
		return vs[0].(map[string]*LocallySignedCert)[vs[1].(string)]
	}).(LocallySignedCertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocallySignedCertInput)(nil)).Elem(), &LocallySignedCert{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocallySignedCertArrayInput)(nil)).Elem(), LocallySignedCertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocallySignedCertMapInput)(nil)).Elem(), LocallySignedCertMap{})
	pulumi.RegisterOutputType(LocallySignedCertOutput{})
	pulumi.RegisterOutputType(LocallySignedCertArrayOutput{})
	pulumi.RegisterOutputType(LocallySignedCertMapOutput{})
}
