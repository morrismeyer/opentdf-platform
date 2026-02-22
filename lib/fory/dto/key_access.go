package dto

// KeyAccess contains cryptographic material and metadata for TDF decryption.
type KeyAccess struct {
	// EncryptedMetadata is base64-encoded encrypted metadata containing additional key information.
	EncryptedMetadata string `fury:"encryptedMetadata"`

	// PolicyBinding ensures cryptographic integrity between policy and wrapped key.
	PolicyBinding *PolicyBinding `fury:"policyBinding"`

	// Protocol is the identifier for the key access mechanism. Typically 'kas'.
	Protocol string `fury:"protocol"`

	// KeyType is the type of key wrapping used for the data encryption key.
	// Values: 'wrapped' (RSA-wrapped), 'ec-wrapped' (ECDH-wrapped).
	KeyType string `fury:"keyType"`

	// KasURL is the URL of the Key Access Server that can unwrap this key.
	KasURL string `fury:"kasUrl"`

	// Kid is the key identifier for the KAS public key used for wrapping.
	Kid string `fury:"kid"`

	// SplitID is the split identifier for key splitting scenarios.
	SplitID string `fury:"splitId"`

	// WrappedKey is the client-generated data encryption key wrapped by KAS.
	WrappedKey []byte `fury:"wrappedKey"`

	// Header contains all metadata and policy information (for formats that embed it).
	Header []byte `fury:"header"`

	// EphemeralPublicKey is the ephemeral public key for ECDH key derivation (ec-wrapped type only).
	EphemeralPublicKey string `fury:"ephemeralPublicKey"`
}
