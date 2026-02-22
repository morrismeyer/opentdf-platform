package fory

import (
	"testing"

	"github.com/opentdf/platform/lib/fory/dto"
)

// BenchmarkEntityChainSerialize measures EntityChain serialization throughput.
func BenchmarkEntityChainSerialize(b *testing.B) {
	codec := NewForyCodec()

	entityChain := &dto.EntityChain{
		EphemeralID: "chain-1",
		Entities: []*dto.Entity{
			{
				EphemeralID: "e1",
				EntityType:  dto.EntityTypeEmailAddress,
				EntityValue: "bob@example.com",
				Category:    dto.CategorySubject,
			},
			{
				EphemeralID: "e2",
				EntityType:  dto.EntityTypeClientID,
				EntityValue: "client-application-123",
				Category:    dto.CategoryEnvironment,
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(entityChain)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEntityChainDeserialize measures EntityChain deserialization throughput.
func BenchmarkEntityChainDeserialize(b *testing.B) {
	codec := NewForyCodec()

	entityChain := &dto.EntityChain{
		EphemeralID: "chain-1",
		Entities: []*dto.Entity{
			{
				EphemeralID: "e1",
				EntityType:  dto.EntityTypeEmailAddress,
				EntityValue: "bob@example.com",
				Category:    dto.CategorySubject,
			},
			{
				EphemeralID: "e2",
				EntityType:  dto.EntityTypeClientID,
				EntityValue: "client-application-123",
				Category:    dto.CategoryEnvironment,
			},
		},
	}

	data, err := codec.Serialize(entityChain)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result dto.EntityChain
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkKeyAccessSerialize measures KeyAccess serialization throughput.
func BenchmarkKeyAccessSerialize(b *testing.B) {
	codec := NewForyCodec()

	keyAccess := &dto.KeyAccess{
		KeyType:    "wrapped",
		KasURL:     "https://kas.example.com/api/v1",
		Kid:        "key-identifier-12345",
		Protocol:   "kas",
		WrappedKey: make([]byte, 256),
		PolicyBinding: &dto.PolicyBinding{
			Algorithm: "HS256",
			Hash:      "abc123hashvalue456def",
		},
		EncryptedMetadata: "base64encodedmetadata==",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(keyAccess)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkKeyAccessDeserialize measures KeyAccess deserialization throughput.
func BenchmarkKeyAccessDeserialize(b *testing.B) {
	codec := NewForyCodec()

	keyAccess := &dto.KeyAccess{
		KeyType:    "wrapped",
		KasURL:     "https://kas.example.com/api/v1",
		Kid:        "key-identifier-12345",
		Protocol:   "kas",
		WrappedKey: make([]byte, 256),
		PolicyBinding: &dto.PolicyBinding{
			Algorithm: "HS256",
			Hash:      "abc123hashvalue456def",
		},
		EncryptedMetadata: "base64encodedmetadata==",
	}

	data, err := codec.Serialize(keyAccess)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result dto.KeyAccess
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDecisionResponseSerialize measures DecisionResponse serialization throughput.
func BenchmarkDecisionResponseSerialize(b *testing.B) {
	codec := NewForyCodec()

	response := &dto.DecisionResponse{
		EntityChainID:        "ec1",
		ResourceAttributesID: "attr-set-1",
		Decision:             dto.DecisionPermit,
		Obligations: []string{
			"http://example.org/obligation/watermark",
			"http://example.org/obligation/audit",
		},
		Action: &dto.Action{
			StandardAction: dto.StandardActionTransmit,
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDecisionResponseDeserialize measures DecisionResponse deserialization throughput.
func BenchmarkDecisionResponseDeserialize(b *testing.B) {
	codec := NewForyCodec()

	response := &dto.DecisionResponse{
		EntityChainID:        "ec1",
		ResourceAttributesID: "attr-set-1",
		Decision:             dto.DecisionPermit,
		Obligations: []string{
			"http://example.org/obligation/watermark",
			"http://example.org/obligation/audit",
		},
		Action: &dto.Action{
			StandardAction: dto.StandardActionTransmit,
		},
	}

	data, err := codec.Serialize(response)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result dto.DecisionResponse
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkPolicyBindingSerialize measures PolicyBinding serialization throughput.
func BenchmarkPolicyBindingSerialize(b *testing.B) {
	codec := NewForyCodec()

	binding := &dto.PolicyBinding{
		Algorithm: "HS256",
		Hash:      "abc123hashvalue456def",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(binding)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkPolicyBindingDeserialize measures PolicyBinding deserialization throughput.
func BenchmarkPolicyBindingDeserialize(b *testing.B) {
	codec := NewForyCodec()

	binding := &dto.PolicyBinding{
		Algorithm: "HS256",
		Hash:      "abc123hashvalue456def",
	}

	data, err := codec.Serialize(binding)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result dto.PolicyBinding
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRewrapResponseSerialize measures RewrapResponse serialization throughput.
func BenchmarkRewrapResponseSerialize(b *testing.B) {
	codec := NewForyCodec()

	response := &dto.RewrapResponse{
		SessionPublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
		Responses: []*dto.PolicyRewrapResult{
			{
				PolicyID: "policy-1",
				Results: []*dto.KeyAccessRewrapResult{
					{
						KeyAccessObjectID: "kao-1",
						Status:            "permit",
						KasWrappedKey:     make([]byte, 256),
					},
				},
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRewrapResponseDeserialize measures RewrapResponse deserialization throughput.
func BenchmarkRewrapResponseDeserialize(b *testing.B) {
	codec := NewForyCodec()

	response := &dto.RewrapResponse{
		SessionPublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
		Responses: []*dto.PolicyRewrapResult{
			{
				PolicyID: "policy-1",
				Results: []*dto.KeyAccessRewrapResult{
					{
						KeyAccessObjectID: "kao-1",
						Status:            "permit",
						KasWrappedKey:     make([]byte, 256),
					},
				},
			},
		},
	}

	data, err := codec.Serialize(response)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result dto.RewrapResponse
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEntityChainRoundTrip measures full serialize+deserialize cycle.
func BenchmarkEntityChainRoundTrip(b *testing.B) {
	codec := NewForyCodec()

	entityChain := &dto.EntityChain{
		EphemeralID: "chain-1",
		Entities: []*dto.Entity{
			{
				EphemeralID: "e1",
				EntityType:  dto.EntityTypeEmailAddress,
				EntityValue: "bob@example.com",
				Category:    dto.CategorySubject,
			},
			{
				EphemeralID: "e2",
				EntityType:  dto.EntityTypeClientID,
				EntityValue: "client-application-123",
				Category:    dto.CategoryEnvironment,
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		data, err := codec.Serialize(entityChain)
		if err != nil {
			b.Fatal(err)
		}

		var result dto.EntityChain
		err = codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkKeyAccessRoundTrip measures full serialize+deserialize cycle.
func BenchmarkKeyAccessRoundTrip(b *testing.B) {
	codec := NewForyCodec()

	keyAccess := &dto.KeyAccess{
		KeyType:    "wrapped",
		KasURL:     "https://kas.example.com/api/v1",
		Kid:        "key-identifier-12345",
		Protocol:   "kas",
		WrappedKey: make([]byte, 256),
		PolicyBinding: &dto.PolicyBinding{
			Algorithm: "HS256",
			Hash:      "abc123hashvalue456def",
		},
		EncryptedMetadata: "base64encodedmetadata==",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		data, err := codec.Serialize(keyAccess)
		if err != nil {
			b.Fatal(err)
		}

		var result dto.KeyAccess
		err = codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}
