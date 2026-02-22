package dto

// StandardAction enumeration for predefined actions.
type StandardAction int

const (
	StandardActionUnspecified StandardAction = iota
	StandardActionDecrypt
	StandardActionTransmit
)

// Action represents an action an entity can take.
type Action struct {
	// ID is the generated uuid in database.
	ID string `fory:"id"`

	// StandardAction is the standard action type.
	StandardAction StandardAction `fory:"standardAction"`

	// CustomAction is the custom action name.
	CustomAction string `fory:"customAction"`

	// Name is the action name.
	Name string `fory:"name"`
}

// Decision enumeration for authorization decisions.
type Decision int

const (
	DecisionUnspecified Decision = iota
	DecisionDeny
	DecisionPermit
)

// Resource is a logical bucket of attributes belonging to a "Resource".
type Resource struct {
	// ResourceAttributesID is the unique identifier for this resource's attributes.
	ResourceAttributesID string `fory:"resourceAttributesId"`

	// AttributeValueFqns is the list of attribute value FQNs associated with this resource.
	AttributeValueFqns []string `fory:"attributeValueFqns"`
}

// DecisionRequest is the request for authorization decisions.
type DecisionRequest struct {
	// Actions is the list of actions to evaluate.
	Actions []*Action `fory:"actions"`

	// EntityChains is the list of entity chains to evaluate.
	EntityChains []*EntityChain `fory:"entityChains"`

	// ResourceAttributes is the list of resource attributes to evaluate against.
	ResourceAttributes []*Resource `fory:"resourceAttributes"`
}

// DecisionResponse contains authorization decision result.
type DecisionResponse struct {
	// EntityChainID is the ephemeral entity chain id from the request.
	EntityChainID string `fory:"entityChainId"`

	// ResourceAttributesID is the ephemeral resource attributes id from the request.
	ResourceAttributesID string `fory:"resourceAttributesId"`

	// Action is the action of the decision response.
	Action *Action `fory:"action"`

	// Decision is the authorization decision.
	Decision Decision `fory:"decision"`

	// Obligations is the optional list of obligations represented in URI format.
	Obligations []string `fory:"obligations"`
}

// IsPermit returns true if the decision is PERMIT.
func (d *DecisionResponse) IsPermit() bool {
	return d.Decision == DecisionPermit
}

// IsDeny returns true if the decision is DENY.
func (d *DecisionResponse) IsDeny() bool {
	return d.Decision == DecisionDeny
}

// EntitlementsResponse contains entity entitlements.
type EntitlementsResponse struct {
	// Entitlements is the list of entity entitlements.
	Entitlements []*EntityEntitlements `fory:"entitlements"`
}

// EntityEntitlements contains entitlements for a single entity.
type EntityEntitlements struct {
	// EntityID is the entity identifier.
	EntityID string `fory:"entityId"`

	// AttributeValueFqns is the list of attribute value FQNs the entity is entitled to.
	AttributeValueFqns []string `fory:"attributeValueFqns"`
}
