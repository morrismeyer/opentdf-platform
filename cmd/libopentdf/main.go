// Package main provides a C-shared library exposing the OpenTDF platform services.
// This library is designed to be loaded via FFM (Foreign Function & Memory API) from Java
// to enable native image support with GraalVM.
//
// Build with: go build -buildmode=c-shared -o libopentdf.so .
package main

/*
#include <stdint.h>
#include <stdlib.h>

// Error codes
#define OPENTDF_OK 0
#define OPENTDF_ERR_INVALID_HANDLE 1
#define OPENTDF_ERR_INVALID_INPUT 2
#define OPENTDF_ERR_SERIALIZATION 3
#define OPENTDF_ERR_INTERNAL 4
#define OPENTDF_ERR_NOT_FOUND 5
#define OPENTDF_ERR_UNAUTHORIZED 6
#define OPENTDF_ERR_BUFFER_TOO_SMALL 7
*/
import "C"

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/opentdf/platform/lib/fory"
	"github.com/opentdf/platform/lib/fory/dto"
)

// PlatformConfig holds configuration for initializing the platform.
type PlatformConfig struct {
	PlatformEndpoint string `json:"platformEndpoint"`
	ClientID         string `json:"clientId"`
	ClientSecret     string `json:"clientSecret"`
	TokenEndpoint    string `json:"tokenEndpoint"`
}

// TDFService represents an initialized OpenTDF platform connection.
type TDFService struct {
	config *PlatformConfig
	codec  *fory.ForyCodec
}

var (
	// Global service registry
	services     = make(map[uint64]*TDFService)
	servicesMu   sync.RWMutex
	handleCounter uint64
)

// allocateHandle returns a new unique handle for a service instance.
func allocateHandle() uint64 {
	return atomic.AddUint64(&handleCounter, 1)
}

// getService retrieves a service instance by handle.
func getService(handle uint64) (*TDFService, bool) {
	servicesMu.RLock()
	defer servicesMu.RUnlock()
	svc, ok := services[handle]
	return svc, ok
}

// ============== Platform Lifecycle ==============

//export PlatformNew
func PlatformNew(configPtr *C.char) C.uint64_t {
	if configPtr == nil {
		return 0
	}

	configJSON := C.GoString(configPtr)
	var config PlatformConfig
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		return 0
	}

	svc := &TDFService{
		config: &config,
		codec:  fory.NewForyCodec(),
	}

	handle := allocateHandle()
	servicesMu.Lock()
	services[handle] = svc
	servicesMu.Unlock()

	return C.uint64_t(handle)
}

//export PlatformFree
func PlatformFree(handle C.uint64_t) {
	servicesMu.Lock()
	delete(services, uint64(handle))
	servicesMu.Unlock()
}

//export PlatformGetVersion
func PlatformGetVersion(versionPtr *C.char, versionLen *C.int) C.int {
	version := "0.12.0"
	if int(*versionLen) < len(version)+1 {
		*versionLen = C.int(len(version) + 1)
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	cVersion := C.CString(version)
	defer C.free(unsafe.Pointer(cVersion))

	// Copy to output buffer
	copy(unsafe.Slice((*byte)(unsafe.Pointer(versionPtr)), len(version)+1),
		unsafe.Slice((*byte)(unsafe.Pointer(cVersion)), len(version)+1))
	*versionLen = C.int(len(version))

	return C.OPENTDF_OK
}

// ============== KAS (Key Access Service) ==============

//export KASGetPublicKey
func KASGetPublicKey(handle C.uint64_t, algorithm C.int, keyPtr *C.uint8_t, keyLen *C.int) C.int {
	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	// Create a sample public key response (in real implementation, this would call the KAS)
	response := &dto.PublicKeyResponse{
		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
		Kid:       "key-1",
	}

	data, err := svc.codec.Serialize(response)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*keyLen) < len(data) {
		*keyLen = C.int(len(data))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	// Copy data to output buffer
	copy(unsafe.Slice((*byte)(unsafe.Pointer(keyPtr)), len(data)), data)
	*keyLen = C.int(len(data))

	return C.OPENTDF_OK
}

//export KASRewrap
func KASRewrap(handle C.uint64_t,
	requestPtr *C.uint8_t, requestLen C.int,
	responsePtr *C.uint8_t, responseLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	if requestPtr == nil || requestLen <= 0 {
		return C.OPENTDF_ERR_INVALID_INPUT
	}

	// Deserialize the request
	requestData := C.GoBytes(unsafe.Pointer(requestPtr), requestLen)
	var request dto.RewrapRequest
	if err := svc.codec.Deserialize(requestData, &request); err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	// Create a sample response (in real implementation, this would call the KAS service)
	response := &dto.RewrapResponse{
		SessionPublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A...\n-----END PUBLIC KEY-----",
		Responses: []*dto.PolicyRewrapResult{
			{
				PolicyID: "policy-1",
				Results: []*dto.KeyAccessRewrapResult{
					{
						KeyAccessObjectID: "kao-1",
						Status:            "permit",
						KasWrappedKey:     make([]byte, 256),
					},
				},
			},
		},
	}

	// Serialize the response
	responseData, err := svc.codec.Serialize(response)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*responseLen) < len(responseData) {
		*responseLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	// Copy response data
	copy(unsafe.Slice((*byte)(unsafe.Pointer(responsePtr)), len(responseData)), responseData)
	*responseLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

// ============== Authorization Service ==============

//export AuthGetDecision
func AuthGetDecision(handle C.uint64_t,
	entityPtr *C.uint8_t, entityLen C.int,
	resourcePtr *C.uint8_t, resourceLen C.int,
	decisionPtr *C.uint8_t, decisionLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	if entityPtr == nil || resourcePtr == nil {
		return C.OPENTDF_ERR_INVALID_INPUT
	}

	// Deserialize entity chain
	entityData := C.GoBytes(unsafe.Pointer(entityPtr), entityLen)
	var entityChain dto.EntityChain
	if err := svc.codec.Deserialize(entityData, &entityChain); err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	// Deserialize resource
	resourceData := C.GoBytes(unsafe.Pointer(resourcePtr), resourceLen)
	var resource dto.Resource
	if err := svc.codec.Deserialize(resourceData, &resource); err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	// Create decision response (in real implementation, this would call the authorization service)
	response := &dto.DecisionResponse{
		EntityChainID:        entityChain.EphemeralID,
		ResourceAttributesID: resource.ResourceAttributesID,
		Decision:             dto.DecisionPermit,
		Action: &dto.Action{
			StandardAction: dto.StandardActionTransmit,
		},
		Obligations: []string{},
	}

	// Serialize response
	responseData, err := svc.codec.Serialize(response)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*decisionLen) < len(responseData) {
		*decisionLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(decisionPtr)), len(responseData)), responseData)
	*decisionLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

//export AuthGetEntitlements
func AuthGetEntitlements(handle C.uint64_t,
	entityPtr *C.uint8_t, entityLen C.int,
	entitlementsPtr *C.uint8_t, entitlementsLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	if entityPtr == nil {
		return C.OPENTDF_ERR_INVALID_INPUT
	}

	// Deserialize entity
	entityData := C.GoBytes(unsafe.Pointer(entityPtr), entityLen)
	var entity dto.Entity
	if err := svc.codec.Deserialize(entityData, &entity); err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	// Create entitlements response
	response := &dto.EntitlementsResponse{
		Entitlements: []*dto.EntityEntitlements{
			{
				EntityID: entity.EphemeralID,
				AttributeValueFqns: []string{
					"https://example.com/attr/classification/value/secret",
					"https://example.com/attr/region/value/us",
				},
			},
		},
	}

	// Serialize response
	responseData, err := svc.codec.Serialize(response)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*entitlementsLen) < len(responseData) {
		*entitlementsLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(entitlementsPtr)), len(responseData)), responseData)
	*entitlementsLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

// ============== Policy Service ==============

//export PolicyListAttributes
func PolicyListAttributes(handle C.uint64_t,
	namespacePtr *C.char,
	attributesPtr *C.uint8_t, attributesLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	namespace := ""
	if namespacePtr != nil {
		namespace = C.GoString(namespacePtr)
	}

	// Create sample attributes response
	active := true
	attributes := []*dto.Attribute{
		{
			ID:   "attr-1",
			Name: "classification",
			Fqn:  namespace + "/attr/classification",
			Rule: dto.RuleTypeHierarchy,
			Namespace: &dto.Namespace{
				ID:     "ns-1",
				Name:   namespace,
				Fqn:    namespace,
				Active: &active,
			},
			Active: &active,
		},
	}

	// Serialize response
	responseData, err := svc.codec.Serialize(attributes)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*attributesLen) < len(responseData) {
		*attributesLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(attributesPtr)), len(responseData)), responseData)
	*attributesLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

//export PolicyGetAttributeValues
func PolicyGetAttributeValues(handle C.uint64_t,
	fqnsPtr *C.uint8_t, fqnsLen C.int,
	valuesPtr *C.uint8_t, valuesLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	if fqnsPtr == nil {
		return C.OPENTDF_ERR_INVALID_INPUT
	}

	// Deserialize FQNs
	fqnsData := C.GoBytes(unsafe.Pointer(fqnsPtr), fqnsLen)
	var fqns []string
	if err := svc.codec.Deserialize(fqnsData, &fqns); err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	// Create sample values response
	active := true
	values := make([]*dto.Value, 0, len(fqns))
	for i, fqn := range fqns {
		values = append(values, &dto.Value{
			ID:     "val-" + string(rune('0'+i)),
			Value:  "value-" + string(rune('0'+i)),
			Fqn:    fqn,
			Active: &active,
		})
	}

	// Serialize response
	responseData, err := svc.codec.Serialize(values)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*valuesLen) < len(responseData) {
		*valuesLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(valuesPtr)), len(responseData)), responseData)
	*valuesLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

// ============== Entity Resolution Service ==============

//export EntityResolve
func EntityResolve(handle C.uint64_t,
	tokenPtr *C.char, tokenLen C.int,
	entityChainPtr *C.uint8_t, entityChainLen *C.int) C.int {

	svc, ok := getService(uint64(handle))
	if !ok {
		return C.OPENTDF_ERR_INVALID_HANDLE
	}

	if tokenPtr == nil || tokenLen <= 0 {
		return C.OPENTDF_ERR_INVALID_INPUT
	}

	// Get the token (in real implementation, this would be validated and resolved)
	// token := C.GoStringN(tokenPtr, tokenLen)

	// Create sample entity chain response
	entityChain := &dto.EntityChain{
		EphemeralID: "resolved-chain-1",
		Entities: []*dto.Entity{
			{
				EphemeralID: "entity-1",
				EntityType:  dto.EntityTypeEmailAddress,
				EntityValue: "user@example.com",
				Category:    dto.CategorySubject,
			},
		},
	}

	// Serialize response
	responseData, err := svc.codec.Serialize(entityChain)
	if err != nil {
		return C.OPENTDF_ERR_SERIALIZATION
	}

	if int(*entityChainLen) < len(responseData) {
		*entityChainLen = C.int(len(responseData))
		return C.OPENTDF_ERR_BUFFER_TOO_SMALL
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(entityChainPtr)), len(responseData)), responseData)
	*entityChainLen = C.int(len(responseData))

	return C.OPENTDF_OK
}

func main() {
	// Required for c-shared build mode
}
