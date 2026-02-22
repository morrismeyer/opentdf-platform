package dto

// RewrapRequest is the request to rewrap (decrypt and re-encrypt) TDF keys for client access.
type RewrapRequest struct {
	// SignedRequestToken is a JWT signed by the DPoP private key.
	SignedRequestToken string `fury:"signedRequestToken"`
}

// UnsignedRewrapRequest is the unsigned rewrap request payload embedded in a JWT.
type UnsignedRewrapRequest struct {
	// ClientPublicKey is the client's public key in PEM format for establishing a session key.
	ClientPublicKey string `fury:"clientPublicKey"`

	// Requests is the list of policy requests to be processed.
	Requests []*WithPolicyRequest `fury:"requests"`
}

// WithPolicy contains policy metadata and content for a group of KeyAccessObjects.
type WithPolicy struct {
	// ID is an identifier unique within the scope of the rewrap request.
	ID string `fury:"id"`

	// Body is the policy content - Base64-encoded JSON policy object.
	Body string `fury:"body"`
}

// WithKeyAccessObject is the Key Access Object wrapper with identifier.
type WithKeyAccessObject struct {
	// KeyAccessObjectID is the ephemeral, unique identifier for this KAO within the request.
	KeyAccessObjectID string `fury:"keyAccessObjectId"`

	// KeyAccessObject is the actual Key Access Object containing cryptographic material.
	KeyAccessObject *KeyAccess `fury:"keyAccessObject"`
}

// WithPolicyRequest groups policy with associated key access objects.
type WithPolicyRequest struct {
	// KeyAccessObjects is the list of Key Access Objects associated with this policy.
	KeyAccessObjects []*WithKeyAccessObject `fury:"keyAccessObjects"`

	// Policy is the policy information for this group of KAOs.
	Policy *WithPolicy `fury:"policy"`

	// Algorithm is the cryptographic algorithm identifier for the TDF type.
	Algorithm string `fury:"algorithm"`
}

// RewrapResponse contains rewrapped keys and session information.
type RewrapResponse struct {
	// SessionPublicKey is the KAS's ephemeral session public key in PEM format.
	SessionPublicKey string `fury:"sessionPublicKey"`

	// Responses is the policy-grouped rewrap results for the bulk API.
	Responses []*PolicyRewrapResult `fury:"responses"`
}

// PolicyRewrapResult contains results for all KAOs associated with a single policy.
type PolicyRewrapResult struct {
	// PolicyID matches the policy.id from the request.
	PolicyID string `fury:"policyId"`

	// Results contains results for each KAO under this policy.
	Results []*KeyAccessRewrapResult `fury:"results"`
}

// KeyAccessRewrapResult is the result of a key access object rewrap operation.
type KeyAccessRewrapResult struct {
	// Metadata associated with this KAO result (e.g., required obligations).
	Metadata map[string]interface{} `fury:"metadata"`

	// KeyAccessObjectID matches the key_access_object_id from the request.
	KeyAccessObjectID string `fury:"keyAccessObjectId"`

	// Status of the rewrap operation for this KAO.
	// Values: "permit" (success), "fail" (failure).
	Status string `fury:"status"`

	// KasWrappedKey is the successfully rewrapped key encrypted with the session key.
	// Present when Status="permit".
	KasWrappedKey []byte `fury:"kasWrappedKey"`

	// Error is the error message when rewrap failed.
	// Present when Status="fail".
	Error string `fury:"error"`
}

// IsSuccess returns true if the rewrap operation succeeded.
func (r *KeyAccessRewrapResult) IsSuccess() bool {
	return r.Status == "permit"
}

// PublicKeyResponse contains a KAS public key.
type PublicKeyResponse struct {
	// PublicKey is the public key in PEM format.
	PublicKey string `fury:"publicKey"`

	// Kid is the key identifier.
	Kid string `fury:"kid"`
}
