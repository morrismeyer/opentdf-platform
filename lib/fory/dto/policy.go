package dto

// Algorithm enumeration for supported key algorithms.
type Algorithm int

const (
	AlgorithmUnspecified Algorithm = iota
	AlgorithmRSA2048
	AlgorithmRSA4096
	AlgorithmECP256
	AlgorithmECP384
	AlgorithmECP521
)

// RuleType enumeration for attribute rule types.
type RuleType int

const (
	RuleTypeUnspecified RuleType = iota
	RuleTypeAllOf
	RuleTypeAnyOf
	RuleTypeHierarchy
)

// Namespace is used for partitioning Attribute Definitions.
type Namespace struct {
	// ID is the generated uuid in database.
	ID string `fory:"id"`

	// Name is used to partition Attribute Definitions and enable federation.
	Name string `fory:"name"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fory:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fory:"active"`

	// KasKeys are the keys for the namespace.
	KasKeys []*SimpleKasKey `fory:"kasKeys"`
}

// Attribute is the attribute definition within a namespace.
type Attribute struct {
	// ID is the generated uuid in database.
	ID string `fory:"id"`

	// Namespace is the namespace of the attribute.
	Namespace *Namespace `fory:"namespace"`

	// Name is the attribute name.
	Name string `fory:"name"`

	// Rule is the attribute rule type.
	Rule RuleType `fory:"rule"`

	// Values are the attribute values.
	Values []*Value `fory:"values"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fory:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fory:"active"`

	// KasKeys are the keys associated with the attribute.
	KasKeys []*SimpleKasKey `fory:"kasKeys"`

	// AllowTraversal indicates whether to use the attribute definition during encryption
	// if the attribute value is missing.
	AllowTraversal *bool `fory:"allowTraversal"`
}

// Value is the attribute value within an attribute definition.
type Value struct {
	// ID is the generated uuid in database.
	ID string `fory:"id"`

	// Attribute is the parent attribute.
	Attribute *Attribute `fory:"attribute"`

	// Value is the value string.
	Value string `fory:"value"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fory:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fory:"active"`

	// KasKeys are the keys associated with the value.
	KasKeys []*SimpleKasKey `fory:"kasKeys"`
}

// SimpleKasKey contains simple KAS key information.
type SimpleKasKey struct {
	// KasURI is the URL of the Key Access Server.
	KasURI string `fory:"kasUri"`

	// PublicKey is the public key that belongs to the KAS.
	PublicKey *SimpleKasPublicKey `fory:"publicKey"`

	// KasID is the ID of the Key Access Server.
	KasID string `fory:"kasId"`
}

// SimpleKasPublicKey contains simple KAS public key information.
type SimpleKasPublicKey struct {
	// Algorithm is the algorithm of the key.
	Algorithm Algorithm `fory:"algorithm"`

	// Kid is the key identifier.
	Kid string `fory:"kid"`

	// Pem is the PEM-encoded public key.
	Pem string `fory:"pem"`
}
