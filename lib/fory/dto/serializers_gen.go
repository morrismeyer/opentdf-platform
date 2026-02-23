// Code generated manually based on forygen patterns. DO NOT EDIT.
// This file contains optimized serializers for DecisionResponse and Action
// that use deferred error checking to eliminate per-field HasError() overhead (~9.4%)

package dto

import (
	"reflect"

	"github.com/apache/fory/go/fory"
)

func init() {
	fory.RegisterSerializerFactory((*DecisionResponse)(nil), NewSerializerFor_DecisionResponse)
	fory.RegisterSerializerFactory((*Action)(nil), NewSerializerFor_Action)
}

// =============================================================================
// DecisionResponse Serializer
// =============================================================================

type DecisionResponse_ForyGenSerializer struct {
	structHash       int32
	actionSerializer *Action_ForyGenSerializer
}

func NewSerializerFor_DecisionResponse() fory.Serializer {
	return &DecisionResponse_ForyGenSerializer{
		actionSerializer: &Action_ForyGenSerializer{},
	}
}

func (g *DecisionResponse_ForyGenSerializer) initHash(resolver *fory.TypeResolver) {
	if g.structHash == 0 {
		g.structHash = fory.GetStructHash(reflect.TypeOf(DecisionResponse{}), resolver)
	}
	g.actionSerializer.initHash(resolver)
}

// Write is the entry point for serialization with ref/type handling
func (g *DecisionResponse_ForyGenSerializer) Write(ctx *fory.WriteContext, refMode fory.RefMode, writeType bool, hasGenerics bool, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	_ = hasGenerics // not used for struct serializers
	switch refMode {
	case fory.RefModeTracking:
		if !value.IsValid() || (value.Kind() == reflect.Ptr && value.IsNil()) {
			ctx.Buffer().WriteInt8(-3) // NullFlag
			return
		}
		refWritten, err := ctx.RefResolver().WriteRefOrNull(ctx.Buffer(), value)
		if err != nil {
			ctx.SetError(fory.FromError(err))
			return
		}
		if refWritten {
			return
		}
	case fory.RefModeNullOnly:
		if !value.IsValid() || (value.Kind() == reflect.Ptr && value.IsNil()) {
			ctx.Buffer().WriteInt8(-3) // NullFlag
			return
		}
		ctx.Buffer().WriteInt8(-1) // NotNullValueFlag
	}
	if writeType {
		ctx.Buffer().WriteVarUint32(uint32(fory.NAMED_STRUCT))
	}
	g.WriteData(ctx, value)
}

// WriteTyped provides strongly-typed serialization with no reflection overhead
func (g *DecisionResponse_ForyGenSerializer) WriteTyped(ctx *fory.WriteContext, v *DecisionResponse) error {
	buf := ctx.Buffer()
	// Write struct hash for compatibility checking
	buf.WriteInt32(g.structHash)

	// Write fields in sorted order (matching Fory's field ordering by name)
	// Field: Action (*Action)
	if v.Action == nil {
		buf.WriteInt8(-3) // NullFlag
	} else {
		buf.WriteInt8(-1) // NotNullValueFlag
		_ = g.actionSerializer.WriteTyped(ctx, v.Action)
	}
	// Field: Decision (enum as int)
	buf.WriteVarint32(int32(v.Decision))
	// Field: EntityChainID (string)
	ctx.WriteString(v.EntityChainID)
	// Field: Obligations ([]string)
	{
		sliceLen := len(v.Obligations)
		buf.WriteVarUint32(uint32(sliceLen))
		if sliceLen > 0 {
			buf.WriteInt8(0) // Collection flags - no ref tracking for string slices
			for _, s := range v.Obligations {
				ctx.WriteString(s)
			}
		}
	}
	// Field: ResourceAttributesID (string)
	ctx.WriteString(v.ResourceAttributesID)
	return nil
}

// WriteData provides reflect.Value interface compatibility (implements fory.Serializer)
func (g *DecisionResponse_ForyGenSerializer) WriteData(ctx *fory.WriteContext, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	var v *DecisionResponse
	if value.Kind() == reflect.Ptr {
		v = value.Interface().(*DecisionResponse)
	} else {
		temp := value.Interface().(DecisionResponse)
		v = &temp
	}
	if err := g.WriteTyped(ctx, v); err != nil {
		ctx.SetError(fory.FromError(err))
	}
}

// Read is the entry point for deserialization with ref/type handling
func (g *DecisionResponse_ForyGenSerializer) Read(ctx *fory.ReadContext, refMode fory.RefMode, readType bool, hasGenerics bool, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	_ = hasGenerics  // not used for struct serializers
	err := ctx.Err() // Get error pointer for deferred error checking
	switch refMode {
	case fory.RefModeTracking:
		refID, refErr := ctx.RefResolver().TryPreserveRefId(ctx.Buffer())
		if refErr != nil {
			ctx.SetError(fory.FromError(refErr))
			return
		}
		if int8(refID) < -1 { // NotNullValueFlag
			obj := ctx.RefResolver().GetReadObject(refID)
			if obj.IsValid() {
				value.Set(obj)
			}
			return
		}
	case fory.RefModeNullOnly:
		flag := ctx.Buffer().ReadInt8(err)
		if flag == -3 { // NullFlag
			return
		}
	}
	if readType {
		ctx.TypeResolver().ReadTypeInfo(ctx.Buffer(), err)
	}
	g.ReadData(ctx, value)
}

// ReadTyped provides strongly-typed deserialization with no reflection overhead
// Uses deferred error checking pattern: only checks ctx.HasError() once at the end
func (g *DecisionResponse_ForyGenSerializer) ReadTyped(ctx *fory.ReadContext, v *DecisionResponse) error {
	buf := ctx.Buffer()
	err := ctx.Err() // Get error pointer for deferred error checking

	// Read and verify struct hash
	if got := buf.ReadInt32(err); got != g.structHash {
		if ctx.HasError() {
			return ctx.TakeError()
		}
		return fory.HashMismatchError(got, g.structHash, "DecisionResponse")
	}

	// Read fields in same order as write (sorted by field name)
	// Field: Action (*Action) - read null flag then nested struct
	actionFlag := buf.ReadInt8(err)
	if actionFlag == -3 { // NullFlag
		v.Action = nil
	} else {
		if v.Action == nil {
			v.Action = &Action{}
		}
		_ = g.actionSerializer.ReadTyped(ctx, v.Action)
	}
	// Field: Decision (enum as int)
	v.Decision = Decision(buf.ReadVarint32(err))
	// Field: EntityChainID (string)
	v.EntityChainID = ctx.ReadString()
	// Field: Obligations ([]string)
	{
		sliceLen := int(buf.ReadVarUint32(err))
		if sliceLen == 0 {
			v.Obligations = v.Obligations[:0] // Reuse slice, clear contents
		} else {
			_ = buf.ReadInt8(err) // Collection flags
			if cap(v.Obligations) >= sliceLen {
				v.Obligations = v.Obligations[:sliceLen]
			} else {
				v.Obligations = make([]string, sliceLen)
			}
			for i := 0; i < sliceLen; i++ {
				v.Obligations[i] = ctx.ReadString()
			}
		}
	}
	// Field: ResourceAttributesID (string)
	v.ResourceAttributesID = ctx.ReadString()

	// Final deferred error check - only ONE HasError() call for entire struct
	if ctx.HasError() {
		return ctx.TakeError()
	}
	return nil
}

// ReadData provides reflect.Value interface compatibility (implements fory.Serializer)
func (g *DecisionResponse_ForyGenSerializer) ReadData(ctx *fory.ReadContext, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	var v *DecisionResponse
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		v = value.Interface().(*DecisionResponse)
	} else {
		v = value.Addr().Interface().(*DecisionResponse)
	}
	if err := g.ReadTyped(ctx, v); err != nil {
		ctx.SetError(fory.FromError(err))
	}
}

// ReadWithTypeInfo deserializes with pre-read type information
func (g *DecisionResponse_ForyGenSerializer) ReadWithTypeInfo(ctx *fory.ReadContext, refMode fory.RefMode, typeInfo *fory.TypeInfo, value reflect.Value) {
	g.Read(ctx, refMode, false, false, value)
}

// =============================================================================
// Action Serializer
// =============================================================================

type Action_ForyGenSerializer struct {
	structHash int32
}

func NewSerializerFor_Action() fory.Serializer {
	return &Action_ForyGenSerializer{}
}

func (g *Action_ForyGenSerializer) initHash(resolver *fory.TypeResolver) {
	if g.structHash == 0 {
		g.structHash = fory.GetStructHash(reflect.TypeOf(Action{}), resolver)
	}
}

// Write is the entry point for serialization with ref/type handling
func (g *Action_ForyGenSerializer) Write(ctx *fory.WriteContext, refMode fory.RefMode, writeType bool, hasGenerics bool, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	_ = hasGenerics
	switch refMode {
	case fory.RefModeTracking:
		if !value.IsValid() || (value.Kind() == reflect.Ptr && value.IsNil()) {
			ctx.Buffer().WriteInt8(-3)
			return
		}
		refWritten, err := ctx.RefResolver().WriteRefOrNull(ctx.Buffer(), value)
		if err != nil {
			ctx.SetError(fory.FromError(err))
			return
		}
		if refWritten {
			return
		}
	case fory.RefModeNullOnly:
		if !value.IsValid() || (value.Kind() == reflect.Ptr && value.IsNil()) {
			ctx.Buffer().WriteInt8(-3)
			return
		}
		ctx.Buffer().WriteInt8(-1)
	}
	if writeType {
		ctx.Buffer().WriteVarUint32(uint32(fory.NAMED_STRUCT))
	}
	g.WriteData(ctx, value)
}

// WriteTyped provides strongly-typed serialization
func (g *Action_ForyGenSerializer) WriteTyped(ctx *fory.WriteContext, v *Action) error {
	buf := ctx.Buffer()
	buf.WriteInt32(g.structHash)

	// Fields sorted by name: CustomAction, ID, Name, StandardAction
	ctx.WriteString(v.CustomAction)
	ctx.WriteString(v.ID)
	ctx.WriteString(v.Name)
	buf.WriteVarint32(int32(v.StandardAction))
	return nil
}

func (g *Action_ForyGenSerializer) WriteData(ctx *fory.WriteContext, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	var v *Action
	if value.Kind() == reflect.Ptr {
		v = value.Interface().(*Action)
	} else {
		temp := value.Interface().(Action)
		v = &temp
	}
	if err := g.WriteTyped(ctx, v); err != nil {
		ctx.SetError(fory.FromError(err))
	}
}

func (g *Action_ForyGenSerializer) Read(ctx *fory.ReadContext, refMode fory.RefMode, readType bool, hasGenerics bool, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	_ = hasGenerics
	err := ctx.Err()
	switch refMode {
	case fory.RefModeTracking:
		refID, refErr := ctx.RefResolver().TryPreserveRefId(ctx.Buffer())
		if refErr != nil {
			ctx.SetError(fory.FromError(refErr))
			return
		}
		if int8(refID) < -1 {
			obj := ctx.RefResolver().GetReadObject(refID)
			if obj.IsValid() {
				value.Set(obj)
			}
			return
		}
	case fory.RefModeNullOnly:
		flag := ctx.Buffer().ReadInt8(err)
		if flag == -3 {
			return
		}
	}
	if readType {
		ctx.TypeResolver().ReadTypeInfo(ctx.Buffer(), err)
	}
	g.ReadData(ctx, value)
}

// ReadTyped uses deferred error checking - only one HasError() call at the end
func (g *Action_ForyGenSerializer) ReadTyped(ctx *fory.ReadContext, v *Action) error {
	buf := ctx.Buffer()
	err := ctx.Err()

	if got := buf.ReadInt32(err); got != g.structHash {
		if ctx.HasError() {
			return ctx.TakeError()
		}
		return fory.HashMismatchError(got, g.structHash, "Action")
	}

	// Fields sorted by name: CustomAction, ID, Name, StandardAction
	v.CustomAction = ctx.ReadString()
	v.ID = ctx.ReadString()
	v.Name = ctx.ReadString()
	v.StandardAction = StandardAction(buf.ReadVarint32(err))

	if ctx.HasError() {
		return ctx.TakeError()
	}
	return nil
}

func (g *Action_ForyGenSerializer) ReadData(ctx *fory.ReadContext, value reflect.Value) {
	g.initHash(ctx.TypeResolver())
	var v *Action
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		v = value.Interface().(*Action)
	} else {
		v = value.Addr().Interface().(*Action)
	}
	if err := g.ReadTyped(ctx, v); err != nil {
		ctx.SetError(fory.FromError(err))
	}
}

func (g *Action_ForyGenSerializer) ReadWithTypeInfo(ctx *fory.ReadContext, refMode fory.RefMode, typeInfo *fory.TypeInfo, value reflect.Value) {
	g.Read(ctx, refMode, false, false, value)
}
