package fory

import (
	"testing"

	"github.com/opentdf/platform/lib/fory/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPolicyBindingRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.PolicyBinding{
		Algorithm: "HS256",
		Hash:      "abc123hash",
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.PolicyBinding
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.Algorithm, deserialized.Algorithm)
	assert.Equal(t, original.Hash, deserialized.Hash)
}

func TestKeyAccessRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.KeyAccess{
		KeyType:    "wrapped",
		KasURL:     "https://kas.example.com",
		Kid:        "key-123",
		Protocol:   "kas",
		WrappedKey: []byte{1, 2, 3, 4, 5},
		PolicyBinding: &dto.PolicyBinding{
			Algorithm: "HS256",
			Hash:      "hashvalue",
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.KeyAccess
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.KeyType, deserialized.KeyType)
	assert.Equal(t, original.KasURL, deserialized.KasURL)
	assert.Equal(t, original.Kid, deserialized.Kid)
	assert.Equal(t, original.WrappedKey, deserialized.WrappedKey)
}

func TestEntityRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.Entity{
		EphemeralID: "entity-1",
		EntityType:  dto.EntityTypeEmailAddress,
		EntityValue: "user@example.com",
		Category:    dto.CategorySubject,
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.Entity
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.EphemeralID, deserialized.EphemeralID)
	assert.Equal(t, original.EntityType, deserialized.EntityType)
	assert.Equal(t, original.EntityValue, deserialized.EntityValue)
	assert.Equal(t, original.Category, deserialized.Category)
}

func TestEntityChainRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.EntityChain{
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
				EntityValue: "client-123",
				Category:    dto.CategoryEnvironment,
			},
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.EntityChain
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.EphemeralID, deserialized.EphemeralID)
	assert.Len(t, deserialized.Entities, 2)
}

func TestRewrapRequestRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.RewrapRequest{
		SignedRequestToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.RewrapRequest
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.SignedRequestToken, deserialized.SignedRequestToken)
}

func TestDecisionResponseRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.DecisionResponse{
		EntityChainID:        "ec1",
		ResourceAttributesID: "attr-set-1",
		Decision:             dto.DecisionPermit,
		Obligations:          []string{"http://example.org/obligation/watermark"},
		Action: &dto.Action{
			StandardAction: dto.StandardActionTransmit,
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.DecisionResponse
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.EntityChainID, deserialized.EntityChainID)
	assert.Equal(t, original.ResourceAttributesID, deserialized.ResourceAttributesID)
	assert.Equal(t, original.Decision, deserialized.Decision)
	assert.True(t, deserialized.IsPermit())
	assert.Equal(t, original.Obligations, deserialized.Obligations)
}

func TestResourceRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.Resource{
		ResourceAttributesID: "resource-1",
		AttributeValueFqns: []string{
			"https://example.com/attr/classification/value/secret",
			"https://example.com/attr/region/value/us",
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.Resource
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.ResourceAttributesID, deserialized.ResourceAttributesID)
	assert.Equal(t, original.AttributeValueFqns, deserialized.AttributeValueFqns)
}

func TestNamespaceRoundTrip(t *testing.T) {
	codec := NewForyCodec()
	active := true

	original := &dto.Namespace{
		ID:     "ns-1",
		Name:   "example.com",
		Fqn:    "https://example.com",
		Active: &active,
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.Namespace
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.ID, deserialized.ID)
	assert.Equal(t, original.Name, deserialized.Name)
	assert.Equal(t, original.Fqn, deserialized.Fqn)
}

func TestAttributeRoundTrip(t *testing.T) {
	codec := NewForyCodec()
	active := true

	original := &dto.Attribute{
		ID:     "attr-1",
		Name:   "classification",
		Fqn:    "https://example.com/attr/classification",
		Rule:   dto.RuleTypeHierarchy,
		Active: &active,
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.Attribute
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.ID, deserialized.ID)
	assert.Equal(t, original.Name, deserialized.Name)
	assert.Equal(t, original.Fqn, deserialized.Fqn)
	assert.Equal(t, original.Rule, deserialized.Rule)
}

func TestRewrapResponseRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.RewrapResponse{
		SessionPublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgk...\n-----END PUBLIC KEY-----",
		Responses: []*dto.PolicyRewrapResult{
			{
				PolicyID: "policy-1",
				Results: []*dto.KeyAccessRewrapResult{
					{
						KeyAccessObjectID: "kao-1",
						Status:            "permit",
						KasWrappedKey:     []byte{10, 20, 30, 40},
					},
				},
			},
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.RewrapResponse
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.SessionPublicKey, deserialized.SessionPublicKey)
	assert.Len(t, deserialized.Responses, 1)
	assert.True(t, deserialized.Responses[0].Results[0].IsSuccess())
}

func TestSimpleKasKeyRoundTrip(t *testing.T) {
	codec := NewForyCodec()

	original := &dto.SimpleKasKey{
		KasURI: "https://kas.example.com",
		KasID:  "kas-1",
		PublicKey: &dto.SimpleKasPublicKey{
			Algorithm: dto.AlgorithmRSA2048,
			Kid:       "key-1",
			Pem:       "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgk...\n-----END PUBLIC KEY-----",
		},
	}

	data, err := codec.Serialize(original)
	require.NoError(t, err)

	var deserialized dto.SimpleKasKey
	err = codec.Deserialize(data, &deserialized)
	require.NoError(t, err)

	assert.Equal(t, original.KasURI, deserialized.KasURI)
	assert.Equal(t, original.KasID, deserialized.KasID)
	assert.Equal(t, original.PublicKey.Algorithm, deserialized.PublicKey.Algorithm)
	assert.Equal(t, original.PublicKey.Kid, deserialized.PublicKey.Kid)
}
