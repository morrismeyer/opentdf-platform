package fory

import (
	"testing"

	"github.com/opentdf/platform/lib/fory/dto"
	"github.com/opentdf/platform/protocol/go/authorization"
	"github.com/opentdf/platform/protocol/go/entity"
	"github.com/opentdf/platform/protocol/go/policy"
	"google.golang.org/protobuf/proto"
)

// ============================================================================
// Fory vs Protobuf Benchmark Comparison
// ============================================================================
//
// Run with: go test -bench=. -benchmem
//
// This file provides head-to-head comparisons between Apache Fory and
// Google Protobuf for identical message structures.

// ============== EntityChain Benchmarks ==============

func BenchmarkForyEntityChainSerialize(b *testing.B) {
	codec := NewForyCodec()
	chain := &dto.EntityChain{
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
		_, err := codec.Serialize(chain)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufEntityChainSerialize(b *testing.B) {
	chain := &entity.EntityChain{
		EphemeralId: "chain-1",
		Entities: []*entity.Entity{
			{
				EphemeralId: "e1",
				EntityType:  &entity.Entity_EmailAddress{EmailAddress: "bob@example.com"},
				Category:    entity.Entity_CATEGORY_SUBJECT,
			},
			{
				EphemeralId: "e2",
				EntityType:  &entity.Entity_ClientId{ClientId: "client-application-123"},
				Category:    entity.Entity_CATEGORY_ENVIRONMENT,
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(chain)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkForyEntityChainDeserialize(b *testing.B) {
	codec := NewForyCodec()
	chain := &dto.EntityChain{
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
	data, _ := codec.Serialize(chain)

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

func BenchmarkProtobufEntityChainDeserialize(b *testing.B) {
	chain := &entity.EntityChain{
		EphemeralId: "chain-1",
		Entities: []*entity.Entity{
			{
				EphemeralId: "e1",
				EntityType:  &entity.Entity_EmailAddress{EmailAddress: "bob@example.com"},
				Category:    entity.Entity_CATEGORY_SUBJECT,
			},
			{
				EphemeralId: "e2",
				EntityType:  &entity.Entity_ClientId{ClientId: "client-application-123"},
				Category:    entity.Entity_CATEGORY_ENVIRONMENT,
			},
		},
	}
	data, _ := proto.Marshal(chain)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result entity.EntityChain
		err := proto.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ============== DecisionResponse Benchmarks ==============

func BenchmarkForyDecisionResponseSerialize(b *testing.B) {
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

func BenchmarkProtobufDecisionResponseSerialize(b *testing.B) {
	response := &authorization.DecisionResponse{
		EntityChainId:        "ec1",
		ResourceAttributesId: "attr-set-1",
		Decision:             authorization.DecisionResponse_DECISION_PERMIT,
		Obligations: []string{
			"http://example.org/obligation/watermark",
			"http://example.org/obligation/audit",
		},
		Action: &policy.Action{
			Value: &policy.Action_Standard{Standard: policy.Action_STANDARD_ACTION_TRANSMIT},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkForyDecisionResponseDeserialize(b *testing.B) {
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
	data, _ := codec.Serialize(response)

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

func BenchmarkProtobufDecisionResponseDeserialize(b *testing.B) {
	response := &authorization.DecisionResponse{
		EntityChainId:        "ec1",
		ResourceAttributesId: "attr-set-1",
		Decision:             authorization.DecisionResponse_DECISION_PERMIT,
		Obligations: []string{
			"http://example.org/obligation/watermark",
			"http://example.org/obligation/audit",
		},
		Action: &policy.Action{
			Value: &policy.Action_Standard{Standard: policy.Action_STANDARD_ACTION_TRANSMIT},
		},
	}
	data, _ := proto.Marshal(response)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result authorization.DecisionResponse
		err := proto.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ============== Token (Simple Object) Benchmarks ==============

func BenchmarkForyTokenSerialize(b *testing.B) {
	codec := NewForyCodec()
	token := &dto.Token{
		EphemeralID: "tok-1",
		JWT:         "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA",
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := codec.Serialize(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufTokenSerialize(b *testing.B) {
	token := &entity.Token{
		EphemeralId: "tok-1",
		Jwt:         "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA",
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkForyTokenDeserialize(b *testing.B) {
	codec := NewForyCodec()
	token := &dto.Token{
		EphemeralID: "tok-1",
		JWT:         "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA",
	}
	data, _ := codec.Serialize(token)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result dto.Token
		err := codec.Deserialize(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufTokenDeserialize(b *testing.B) {
	token := &entity.Token{
		EphemeralId: "tok-1",
		Jwt:         "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA",
	}
	data, _ := proto.Marshal(token)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result entity.Token
		err := proto.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ============== Round-Trip Benchmarks ==============

func BenchmarkForyEntityChainRoundTrip(b *testing.B) {
	codec := NewForyCodec()
	chain := &dto.EntityChain{
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
		data, err := codec.Serialize(chain)
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

func BenchmarkProtobufEntityChainRoundTrip(b *testing.B) {
	chain := &entity.EntityChain{
		EphemeralId: "chain-1",
		Entities: []*entity.Entity{
			{
				EphemeralId: "e1",
				EntityType:  &entity.Entity_EmailAddress{EmailAddress: "bob@example.com"},
				Category:    entity.Entity_CATEGORY_SUBJECT,
			},
			{
				EphemeralId: "e2",
				EntityType:  &entity.Entity_ClientId{ClientId: "client-application-123"},
				Category:    entity.Entity_CATEGORY_ENVIRONMENT,
			},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, err := proto.Marshal(chain)
		if err != nil {
			b.Fatal(err)
		}
		var result entity.EntityChain
		err = proto.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ============== Serialized Size Comparison ==============

func BenchmarkForySerializedSize(b *testing.B) {
	codec := NewForyCodec()
	chain := &dto.EntityChain{
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
	data, _ := codec.Serialize(chain)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}

func BenchmarkProtobufSerializedSize(b *testing.B) {
	chain := &entity.EntityChain{
		EphemeralId: "chain-1",
		Entities: []*entity.Entity{
			{
				EphemeralId: "e1",
				EntityType:  &entity.Entity_EmailAddress{EmailAddress: "bob@example.com"},
				Category:    entity.Entity_CATEGORY_SUBJECT,
			},
			{
				EphemeralId: "e2",
				EntityType:  &entity.Entity_ClientId{ClientId: "client-application-123"},
				Category:    entity.Entity_CATEGORY_ENVIRONMENT,
			},
		},
	}
	data, _ := proto.Marshal(chain)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}
