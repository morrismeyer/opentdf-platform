package dto

// EntityType enumeration for entity identification methods.
type EntityType int

const (
	EntityTypeUnspecified EntityType = iota
	EntityTypeEmailAddress
	EntityTypeUserName
	EntityTypeClaims
	EntityTypeClientID
)

// Category enumeration for entity categories.
type Category int

const (
	CategoryUnspecified Category = iota
	CategorySubject
	CategoryEnvironment
)

// Entity represents a Person Entity (PE) or Non-Person Entity (NPE).
type Entity struct {
	// EphemeralID is for tracking between request and response.
	EphemeralID string `fory:"ephemeralId"`

	// EntityType is the type of entity.
	EntityType EntityType `fory:"entityType"`

	// EntityValue is the entity identifier value (email, username, client_id, etc.).
	EntityValue string `fory:"entityValue"`

	// Claims is the claims data as JSON string (when EntityType is Claims).
	Claims string `fory:"claims"`

	// Category is the entity category.
	Category Category `fory:"category"`
}

// EntityChain is a set of related Person Entities (PE) and Non-Person Entities (NPE).
type EntityChain struct {
	// EphemeralID is for tracking between request and response.
	EphemeralID string `fory:"ephemeralId"`

	// Entities is the list of entities in this chain.
	Entities []*Entity `fory:"entities"`
}

// Token represents an authentication/authorization token.
type Token struct {
	// EphemeralID is for tracking between request and response.
	EphemeralID string `fory:"ephemeralId"`

	// JWT is the JWT token.
	JWT string `fory:"jwt"`
}
