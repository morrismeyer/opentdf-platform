// Package fory provides Apache Fory serialization support for OpenTDF.
// It enables high-performance binary serialization for cross-language communication
// between Go and Java components.
package fory

import (
	"github.com/apache/fory/go/fory"
	"github.com/opentdf/platform/lib/fory/dto"
)

// ForyCodec wraps the Fory serializer with OpenTDF-specific configuration.
type ForyCodec struct {
	fory *fory.Fory
}

// NewForyCodec creates a new ForyCodec with all OpenTDF types registered.
func NewForyCodec() *ForyCodec {
	f := fory.NewFory(
		fory.WithRefTracking(true),  // Enable reference tracking for schema compatibility
		fory.WithXlang(true),        // Enable cross-language support
		fory.WithCompatible(true),   // Enable schema compatibility mode
	)
	registerTypes(f)
	return &ForyCodec{fory: f}
}

// Serialize converts an object to a byte array.
func (c *ForyCodec) Serialize(obj interface{}) ([]byte, error) {
	return c.fory.Marshal(obj)
}

// Deserialize converts a byte array back to an object.
func (c *ForyCodec) Deserialize(data []byte, obj interface{}) error {
	return c.fory.Unmarshal(data, obj)
}

// registerTypes registers all OpenTDF DTO types with Fory.
func registerTypes(f *fory.Fory) {
	// Enum types must be registered first
	_ = f.RegisterNamedEnum(dto.EntityTypeUnspecified, "io.opentdf.fory.dto.EntityDto$EntityType")
	_ = f.RegisterNamedEnum(dto.CategoryUnspecified, "io.opentdf.fory.dto.EntityDto$Category")
	_ = f.RegisterNamedEnum(dto.StandardActionUnspecified, "io.opentdf.fory.dto.ActionDto$StandardAction")
	_ = f.RegisterNamedEnum(dto.DecisionUnspecified, "io.opentdf.fory.dto.DecisionResponseDto$Decision")
	_ = f.RegisterNamedEnum(dto.RuleTypeUnspecified, "io.opentdf.fory.dto.AttributeDto$RuleType")
	_ = f.RegisterNamedEnum(dto.AlgorithmUnspecified, "io.opentdf.fory.dto.SimpleKasKeyDto$Algorithm")

	// Core DTOs
	_ = f.RegisterNamedStruct(dto.PolicyBinding{}, "io.opentdf.fory.dto.PolicyBindingDto")
	_ = f.RegisterNamedStruct(dto.KeyAccess{}, "io.opentdf.fory.dto.KeyAccessDto")
	_ = f.RegisterNamedStruct(dto.Entity{}, "io.opentdf.fory.dto.EntityDto")
	_ = f.RegisterNamedStruct(dto.EntityChain{}, "io.opentdf.fory.dto.EntityChainDto")
	_ = f.RegisterNamedStruct(dto.Token{}, "io.opentdf.fory.dto.TokenDto")

	// KAS DTOs
	_ = f.RegisterNamedStruct(dto.RewrapRequest{}, "io.opentdf.fory.dto.RewrapRequestDto")
	_ = f.RegisterNamedStruct(dto.UnsignedRewrapRequest{}, "io.opentdf.fory.dto.RewrapRequestDto$UnsignedRewrapRequestDto")
	_ = f.RegisterNamedStruct(dto.WithPolicy{}, "io.opentdf.fory.dto.RewrapRequestDto$WithPolicyDto")
	_ = f.RegisterNamedStruct(dto.WithKeyAccessObject{}, "io.opentdf.fory.dto.RewrapRequestDto$WithKeyAccessObjectDto")
	_ = f.RegisterNamedStruct(dto.WithPolicyRequest{}, "io.opentdf.fory.dto.RewrapRequestDto$WithPolicyRequestDto")
	_ = f.RegisterNamedStruct(dto.RewrapResponse{}, "io.opentdf.fory.dto.RewrapResponseDto")
	_ = f.RegisterNamedStruct(dto.PolicyRewrapResult{}, "io.opentdf.fory.dto.RewrapResponseDto$PolicyRewrapResultDto")
	_ = f.RegisterNamedStruct(dto.KeyAccessRewrapResult{}, "io.opentdf.fory.dto.RewrapResponseDto$KeyAccessRewrapResultDto")
	_ = f.RegisterNamedStruct(dto.PublicKeyResponse{}, "io.opentdf.fory.dto.PublicKeyResponseDto")

	// Authorization DTOs
	_ = f.RegisterNamedStruct(dto.Action{}, "io.opentdf.fory.dto.ActionDto")
	_ = f.RegisterNamedStruct(dto.Resource{}, "io.opentdf.fory.dto.ResourceDto")
	_ = f.RegisterNamedStruct(dto.DecisionRequest{}, "io.opentdf.fory.dto.DecisionRequestDto")
	_ = f.RegisterNamedStruct(dto.DecisionResponse{}, "io.opentdf.fory.dto.DecisionResponseDto")
	_ = f.RegisterNamedStruct(dto.EntitlementsResponse{}, "io.opentdf.fory.dto.EntitlementsResponseDto")
	_ = f.RegisterNamedStruct(dto.EntityEntitlements{}, "io.opentdf.fory.dto.EntitlementsResponseDto$EntityEntitlementsDto")

	// Policy DTOs
	_ = f.RegisterNamedStruct(dto.Namespace{}, "io.opentdf.fory.dto.NamespaceDto")
	_ = f.RegisterNamedStruct(dto.Attribute{}, "io.opentdf.fory.dto.AttributeDto")
	_ = f.RegisterNamedStruct(dto.Value{}, "io.opentdf.fory.dto.ValueDto")
	_ = f.RegisterNamedStruct(dto.SimpleKasKey{}, "io.opentdf.fory.dto.SimpleKasKeyDto")
	_ = f.RegisterNamedStruct(dto.SimpleKasPublicKey{}, "io.opentdf.fory.dto.SimpleKasKeyDto$SimpleKasPublicKeyDto")
}
