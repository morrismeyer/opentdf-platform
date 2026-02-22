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
	ID string `fury:"id"`

	// Name is used to partition Attribute Definitions and enable federation.
	Name string `fury:"name"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fury:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fury:"active"`

	// KasKeys are the keys for the namespace.
	KasKeys []*SimpleKasKey `fury:"kasKeys"`
}

// Attribute is the attribute definition within a namespace.
type Attribute struct {
	// ID is the generated uuid in database.
	ID string `fury:"id"`

	// Namespace is the namespace of the attribute.
	Namespace *Namespace `fury:"namespace"`

	// Name is the attribute name.
	Name string `fury:"name"`

	// Rule is the attribute rule type.
	Rule RuleType `fury:"rule"`

	// Values are the attribute values.
	Values []*Value `fury:"values"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fury:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fury:"active"`

	// KasKeys are the keys associated with the attribute.
	KasKeys []*SimpleKasKey `fury:"kasKeys"`

	// AllowTraversal indicates whether to use the attribute definition during encryption
	// if the attribute value is missing.
	AllowTraversal *bool `fury:"allowTraversal"`
}

// Value is the attribute value within an attribute definition.
type Value struct {
	// ID is the generated uuid in database.
	ID string `fury:"id"`

	// Attribute is the parent attribute.
	Attribute *Attribute `fury:"attribute"`

	// Value is the value string.
	Value string `fury:"value"`

	// Fqn is the Fully Qualified Name.
	Fqn string `fury:"fqn"`

	// Active is true by default until explicitly deactivated.
	Active *bool `fury:"active"`

	// KasKeys are the keys associated with the value.
	KasKeys []*SimpleKasKey `fury:"kasKeys"`
}

// SimpleKasKey contains simple KAS key information.
type SimpleKasKey struct {
	// KasURI is the URL of the Key Access Server.
	KasURI string `fury:"kasUri"`

	// PublicKey is the public key that belongs to the KAS.
	PublicKey *SimpleKasPublicKey `fury:"publicKey"`

	// KasID is the ID of the Key Access Server.
	KasID string `fury:"kasId"`
}

// SimpleKasPublicKey contains simple KAS public key information.
type SimpleKasPublicKey struct {
	// Algorithm is the algorithm of the key.
	Algorithm Algorithm `fury:"algorithm"`

	// Kid is the key identifier.
	Kid string `fury:"kid"`

	// Pem is the PEM-encoded public key.
	Pem string `fury:"pem"`
}
