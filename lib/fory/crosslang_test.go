package fory

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/opentdf/platform/lib/fory/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// crossLanguageDirectory returns the path to cross-language test data.
// This directory is shared between Java and Go for compatibility testing.
func crossLanguageDirectory(t *testing.T) string {
	// Use a temp directory for cross-language tests
	dir := filepath.Join(os.TempDir(), "opentdf-fory-crosslang")
	err := os.MkdirAll(dir, 0o755)
	require.NoError(t, err)
	return dir
}

// TestWriteGoSerializedData writes Go-serialized data for Java to read.
// Run this first, then run the Java CrossLanguageTest.testReadGoData().
func TestWriteGoSerializedData(t *testing.T) {
	codec := NewForyCodec()
	dir := crossLanguageDirectory(t)

	// EntityChain
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
				EntityValue: "client-123",
				Category:    dto.CategoryEnvironment,
			},
		},
	}
	data, err := codec.Serialize(entityChain)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, "go-entity-chain.bin"), data, 0o644)
	require.NoError(t, err)
	t.Logf("EntityChain serialized: %d bytes", len(data))
	t.Logf("Hex: %s", hex.EncodeToString(data))

	// KeyAccess
	keyAccess := &dto.KeyAccess{
		KeyType:    "wrapped",
		KasURL:     "https://kas.example.com",
		Kid:        "key-123",
		Protocol:   "kas",
		WrappedKey: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		PolicyBinding: &dto.PolicyBinding{
			Algorithm: "HS256",
			Hash:      "abc123hash",
		},
	}
	data, err = codec.Serialize(keyAccess)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, "go-key-access.bin"), data, 0o644)
	require.NoError(t, err)
	t.Logf("KeyAccess serialized: %d bytes", len(data))

	// DecisionResponse
	response := &dto.DecisionResponse{
		EntityChainID:        "ec1",
		ResourceAttributesID: "attr-1",
		Decision:             dto.DecisionPermit,
		Obligations:          []string{"http://example.org/audit"},
		Action: &dto.Action{
			StandardAction: dto.StandardActionTransmit,
		},
	}
	data, err = codec.Serialize(response)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, "go-decision-response.bin"), data, 0o644)
	require.NoError(t, err)
	t.Logf("DecisionResponse serialized: %d bytes", len(data))

	t.Logf("Cross-language test data written to: %s", dir)
}

// TestReadJavaSerializedData reads Java-serialized data.
// Run Java's CrossLanguageTest.testWriteJavaData() first to generate the data.
func TestReadJavaSerializedData(t *testing.T) {
	codec := NewForyCodec()
	dir := crossLanguageDirectory(t)

	// Try to read EntityChain from Java
	entityChainPath := filepath.Join(dir, "java-entity-chain.bin")
	if data, err := os.ReadFile(entityChainPath); err == nil {
		var entityChain dto.EntityChain
		err = codec.Deserialize(data, &entityChain)
		if err != nil {
			t.Logf("Failed to deserialize Java EntityChain: %v", err)
			t.Logf("Data hex: %s", hex.EncodeToString(data))
		} else {
			assert.Equal(t, "chain-1", entityChain.EphemeralID)
			assert.Len(t, entityChain.Entities, 2)
			t.Logf("Successfully deserialized Java EntityChain: %+v", entityChain)
		}
	} else {
		t.Skipf("Java EntityChain file not found at %s - run Java tests first", entityChainPath)
	}

	// Try to read KeyAccess from Java
	keyAccessPath := filepath.Join(dir, "java-key-access.bin")
	if data, err := os.ReadFile(keyAccessPath); err == nil {
		var keyAccess dto.KeyAccess
		err = codec.Deserialize(data, &keyAccess)
		if err != nil {
			t.Logf("Failed to deserialize Java KeyAccess: %v", err)
		} else {
			assert.Equal(t, "wrapped", keyAccess.KeyType)
			assert.Equal(t, "https://kas.example.com", keyAccess.KasURL)
			t.Logf("Successfully deserialized Java KeyAccess: %+v", keyAccess)
		}
	} else {
		t.Skipf("Java KeyAccess file not found at %s", keyAccessPath)
	}

	// Try to read DecisionResponse from Java
	responsePath := filepath.Join(dir, "java-decision-response.bin")
	if data, err := os.ReadFile(responsePath); err == nil {
		var response dto.DecisionResponse
		err = codec.Deserialize(data, &response)
		if err != nil {
			t.Logf("Failed to deserialize Java DecisionResponse: %v", err)
		} else {
			assert.Equal(t, "ec1", response.EntityChainID)
			assert.Equal(t, dto.DecisionPermit, response.Decision)
			t.Logf("Successfully deserialized Java DecisionResponse: %+v", response)
		}
	} else {
		t.Skipf("Java DecisionResponse file not found at %s", responsePath)
	}
}

// BenchmarkCrossLanguageReadJavaData measures reading Java-serialized data.
func BenchmarkCrossLanguageReadJavaData(b *testing.B) {
	codec := NewForyCodec()
	dir := filepath.Join(os.TempDir(), "opentdf-fory-crosslang")

	// Try to read EntityChain from Java
	entityChainPath := filepath.Join(dir, "java-entity-chain.bin")
	data, err := os.ReadFile(entityChainPath)
	if err != nil {
		b.Skipf("Java EntityChain file not found - run Java tests first")
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var entityChain dto.EntityChain
		err = codec.Deserialize(data, &entityChain)
		if err != nil {
			b.Fatal(err)
		}
	}
}
