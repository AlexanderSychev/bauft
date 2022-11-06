package data

import (
	"fmt"
	"reflect"
	"sync"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "valueFactory" corresponding errors
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// TypeNotRegisteredError returns by valueFactory.CreateValueByTypeId method when factory has no type
// with received identifier
type TypeNotRegisteredError struct {
	// Type unique identifier
	typeId uint64
}

func (err TypeNotRegisteredError) Error() string {
	return fmt.Sprintf("type with id %d not registered", err.typeId)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "valueFactory" type and methods
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// valueFactory contains types which can contain Row instance. This types can be serialized
type valueFactory struct {
	// Mutual exclusion lock for thread-safe work with factory
	mutex sync.RWMutex
	// Registered types' dictionary where key is type unique identifier and value is type
	types map[uint64]reflect.Type
}

func (vf *valueFactory) registerTypeByRef(ref Value) {
	typeId := ref.TypeId()
	tp := reflect.TypeOf(ref)
	vf.types[typeId] = tp
}

func (vf *valueFactory) RegisterTypeByRef(ref Value) {
	vf.mutex.Lock()
	defer vf.mutex.Unlock()

	vf.registerTypeByRef(ref)
}

func (vf *valueFactory) hasTypeWithId(typeId uint64) bool {
	_, result := vf.types[typeId]
	return result
}

func (vf *valueFactory) HasTypeWithId(typeId uint64) bool {
	vf.mutex.RLock()
	result := vf.hasTypeWithId(typeId)
	vf.mutex.RUnlock()

	return result
}

func (vf *valueFactory) createValueByTypeId(typeId uint64) (Value, error) {
	// Check whether factory has type with received identifier
	tp, hasType := vf.types[typeId]
	if !hasType {
		return nil, TypeNotRegisteredError{
			typeId: typeId,
		}
	}

	// Created value instance and return it
	result := reflect.New(tp).Interface().(Value)

	return result, nil
}

func (vf *valueFactory) CreateValueByTypeId(typeId uint64) (Value, error) {
	vf.mutex.RLock()
	result, err := vf.createValueByTypeId(typeId)
	vf.mutex.RUnlock()

	return result, err
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "valueFactory" instance initialize
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

var ValueFactory *valueFactory = nil
var valueFactoryInit = &sync.Once{}

func init() {
	valueFactoryInit.Do(func() {
		ValueFactory = &valueFactory{
			mutex: sync.RWMutex{},
			types: make(map[uint64]reflect.Type),
		}

		numberRef := Number(0)
		ValueFactory.registerTypeByRef(&numberRef)
	})
}
