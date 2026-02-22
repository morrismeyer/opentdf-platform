package dto

// PolicyBinding ensures cryptographic integrity between policy and wrapped key.
// Prevents policy tampering by binding the policy hash to the encrypted key.
type PolicyBinding struct {
	// Algorithm is the cryptographic hashing algorithm used for policy binding.
	// Value: Always "HS256" (HMAC-SHA256).
	Algorithm string `fory:"algorithm"`

	// Hash is the HMAC-SHA256 hash of the base64-encoded policy using the DEK as the secret key.
	Hash string `fory:"hash"`
}
