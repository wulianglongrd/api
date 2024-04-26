// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: security/v1beta1/jwt.proto

// $schema: istio.security.v1beta1.JWTRule
// $title: JWTRule
// $description: Configuration to validate JWT.
// $location: https://istio.io/docs/reference/config/security/jwt.html
// $aliases: [/docs/reference/config/security/v1beta1/jwt]

package v1beta1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// <!-- istio code generation tags
// +istio.io/sync-start
// -->
// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Examples:
//
// Spec for a JWT that is issued by `https://example.com`, with the audience claims must be either
// `bookstore_android.apps.example.com` or `bookstore_web.apps.example.com`.
// The token should be presented at the `Authorization` header (default). The JSON Web Key Set (JWKS)
// will be discovered following OpenID Connect protocol.
//
// ```yaml
// issuer: https://example.com
// audiences:
//   - bookstore_android.apps.example.com
//     bookstore_web.apps.example.com
//
// ```
//
// This example specifies a token in a non-default location (`x-goog-iap-jwt-assertion` header). It also
// defines the URI to fetch JWKS explicitly.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.secret/jwks.json
// fromHeaders:
// - "x-goog-iap-jwt-assertion"
// ```
type JWTRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// A JWT with different `iss` claim will be rejected.
	//
	// Example: `https://foobar.auth0.com`
	// Example: `1234567-compute@developer.gserviceaccount.com`
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3)
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	// ```yaml
	// audiences:
	//   - bookstore_android.apps.example.com
	//     bookstore_web.apps.example.com
	//
	// ```
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: `https://www.googleapis.com/oauth2/v1/certs`
	//
	// Note: Only one of `jwksUri`, `jwks` and `jwksFile` should be used.
	// +kubebuilder:altName=jwks_uri
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See https://auth0.com/docs/jwks.
	//
	// Note: Only one of `jwksUri`, `jwks` and `jwksFile` should be used.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// Local file of the provider's public key set to validate signature of the JWT.
	//
	// Note: Only one of `jwksUri`, `jwks` and `jwksFile` should be used.
	JwksFile string `protobuf:"bytes,14,opt,name=jwks_file,json=jwksFile,proto3" json:"jwks_file,omitempty"`
	// List of header locations from which JWT is expected. For example, below is the location spec
	// if JWT is expected to be found in `x-jwt-assertion` header, and have `Bearer` prefix:
	//
	// ```yaml
	//
	//	fromHeaders:
	//	- name: x-jwt-assertion
	//	  prefix: "Bearer "
	//
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromHeaders []*JWTHeader `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	// List of query parameters from which JWT is expected. For example, if JWT is provided via query
	// parameter `my_token` (e.g `/path?my_token=<JWT>`), the config is:
	//
	// ```yaml
	//
	//	fromParams:
	//	- "my_token"
	//
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromParams []string `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	// This field specifies the header name to output a successfully verified JWT payload to the
	// backend. The forwarded data is `base64_encoded(jwt_payload_in_JSON)`. If it is not specified,
	// the payload will not be emitted.
	OutputPayloadToHeader string `protobuf:"bytes,8,opt,name=output_payload_to_header,json=outputPayloadToHeader,proto3" json:"output_payload_to_header,omitempty"`
	// List of cookie names from which JWT is expected.	//
	// For example, if config is:
	//
	// ``` yaml
	//
	//	from_cookies:
	//	- auth-token
	//
	// ```
	// Then JWT will be extracted from “auth-token“ cookie in the request.
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromCookies []string `protobuf:"bytes,12,rep,name=from_cookies,json=fromCookies,proto3" json:"from_cookies,omitempty"`
	// If set to true, the original token will be kept for the upstream request. Default is false.
	ForwardOriginalToken bool `protobuf:"varint,9,opt,name=forward_original_token,json=forwardOriginalToken,proto3" json:"forward_original_token,omitempty"`
	// This field specifies a list of operations to copy the claim to HTTP headers on a successfully verified token.
	// This differs from the `output_payload_to_header` by allowing outputting individual claims instead of the whole payload.
	// The header specified in each operation in the list must be unique. Nested claims of type string/int/bool is supported as well.
	// ```
	//
	//	outputClaimToHeaders:
	//	- header: x-my-company-jwt-group
	//	  claim: my-group
	//	- header: x-test-environment-flag
	//	  claim: test-flag
	//	- header: x-jwt-claim-group
	//	  claim: nested.key.group
	//
	// ```
	// [Experimental] This feature is a experimental feature.
	OutputClaimToHeaders []*ClaimToHeader `protobuf:"bytes,11,rep,name=output_claim_to_headers,json=outputClaimToHeaders,proto3" json:"output_claim_to_headers,omitempty"` // [TODO:Update the status whenever this feature is promoted.]
	// The maximum amount of time that the resolver, determined by the PILOT_JWT_ENABLE_REMOTE_JWKS environment variable,
	// will spend waiting for the JWKS to be fetched. Default is 5s.
	Timeout *duration.Duration `protobuf:"bytes,13,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *JWTRule) Reset() {
	*x = JWTRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTRule) ProtoMessage() {}

func (x *JWTRule) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTRule.ProtoReflect.Descriptor instead.
func (*JWTRule) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *JWTRule) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *JWTRule) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *JWTRule) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

func (x *JWTRule) GetJwks() string {
	if x != nil {
		return x.Jwks
	}
	return ""
}

func (x *JWTRule) GetJwksFile() string {
	if x != nil {
		return x.JwksFile
	}
	return ""
}

func (x *JWTRule) GetFromHeaders() []*JWTHeader {
	if x != nil {
		return x.FromHeaders
	}
	return nil
}

func (x *JWTRule) GetFromParams() []string {
	if x != nil {
		return x.FromParams
	}
	return nil
}

func (x *JWTRule) GetOutputPayloadToHeader() string {
	if x != nil {
		return x.OutputPayloadToHeader
	}
	return ""
}

func (x *JWTRule) GetFromCookies() []string {
	if x != nil {
		return x.FromCookies
	}
	return nil
}

func (x *JWTRule) GetForwardOriginalToken() bool {
	if x != nil {
		return x.ForwardOriginalToken
	}
	return false
}

func (x *JWTRule) GetOutputClaimToHeaders() []*ClaimToHeader {
	if x != nil {
		return x.OutputClaimToHeaders
	}
	return nil
}

func (x *JWTRule) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// This message specifies a header location to extract JWT token.
type JWTHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP header name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The prefix that should be stripped before decoding the token.
	// For example, for `Authorization: Bearer <token>`, prefix=`Bearer` with a space at the end.
	// If the header doesn't have this exact prefix, it is considered invalid.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *JWTHeader) Reset() {
	*x = JWTHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTHeader) ProtoMessage() {}

func (x *JWTHeader) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTHeader.ProtoReflect.Descriptor instead.
func (*JWTHeader) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *JWTHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JWTHeader) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// This message specifies the detail for copying claim to header.
type ClaimToHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the header to be created. The header will be overridden if it already exists in the request.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The name of the claim to be copied from. Only claim of type string/int/bool is supported.
	// The header will not be there if the claim does not exist or the type of the claim is not supported.
	Claim string `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *ClaimToHeader) Reset() {
	*x = ClaimToHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimToHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimToHeader) ProtoMessage() {}

func (x *ClaimToHeader) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimToHeader.ProtoReflect.Descriptor instead.
func (*ClaimToHeader) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *ClaimToHeader) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *ClaimToHeader) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

var File_security_v1beta1_jwt_proto protoreflect.FileDescriptor

var file_security_v1beta1_jwt_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x07, 0x4a, 0x57, 0x54, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x77, 0x6b, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x77, 0x6b, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4a, 0x57, 0x54, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x37, 0x0a, 0x18, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x5c, 0x0a, 0x17, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x14, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x22, 0x3c, 0x0a, 0x09, 0x4a, 0x57, 0x54, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x22, 0x3d, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_v1beta1_jwt_proto_rawDescOnce sync.Once
	file_security_v1beta1_jwt_proto_rawDescData = file_security_v1beta1_jwt_proto_rawDesc
)

func file_security_v1beta1_jwt_proto_rawDescGZIP() []byte {
	file_security_v1beta1_jwt_proto_rawDescOnce.Do(func() {
		file_security_v1beta1_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_v1beta1_jwt_proto_rawDescData)
	})
	return file_security_v1beta1_jwt_proto_rawDescData
}

var file_security_v1beta1_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_security_v1beta1_jwt_proto_goTypes = []interface{}{
	(*JWTRule)(nil),           // 0: istio.security.v1beta1.JWTRule
	(*JWTHeader)(nil),         // 1: istio.security.v1beta1.JWTHeader
	(*ClaimToHeader)(nil),     // 2: istio.security.v1beta1.ClaimToHeader
	(*duration.Duration)(nil), // 3: google.protobuf.Duration
}
var file_security_v1beta1_jwt_proto_depIdxs = []int32{
	1, // 0: istio.security.v1beta1.JWTRule.from_headers:type_name -> istio.security.v1beta1.JWTHeader
	2, // 1: istio.security.v1beta1.JWTRule.output_claim_to_headers:type_name -> istio.security.v1beta1.ClaimToHeader
	3, // 2: istio.security.v1beta1.JWTRule.timeout:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_security_v1beta1_jwt_proto_init() }
func file_security_v1beta1_jwt_proto_init() {
	if File_security_v1beta1_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_v1beta1_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_security_v1beta1_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_security_v1beta1_jwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimToHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_v1beta1_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_v1beta1_jwt_proto_goTypes,
		DependencyIndexes: file_security_v1beta1_jwt_proto_depIdxs,
		MessageInfos:      file_security_v1beta1_jwt_proto_msgTypes,
	}.Build()
	File_security_v1beta1_jwt_proto = out.File
	file_security_v1beta1_jwt_proto_rawDesc = nil
	file_security_v1beta1_jwt_proto_goTypes = nil
	file_security_v1beta1_jwt_proto_depIdxs = nil
}
