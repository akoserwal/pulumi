// Copyright 2016 Marapongo, Inc. All rights reserved.

// Because of the complex structure of the MuPack and MuIL metadata formats, we cannot rely on the standard JSON
// marshaling and unmarshaling routines.  Instead, we will need to do it mostly "by hand".  This package does that.
package encoding

import (
	"reflect"

	"github.com/marapongo/mu/pkg/encoding"
	"github.com/marapongo/mu/pkg/pack"
	"github.com/marapongo/mu/pkg/pack/ast"
	"github.com/marapongo/mu/pkg/util/mapper"
)

// Decode unmarshals the entire contents of the given byte array into a Package object.
func Decode(m encoding.Marshaler, b []byte) (*pack.Package, error) {
	// First convert the whole contents of the metadata into a map.  Although it would be more efficient to walk the
	// token stream, token by token, this allows us to reuse existing YAML packages in addition to JSON ones.
	var tree mapper.Object
	if err := m.Unmarshal(b, &tree); err != nil {
		return nil, err
	}

	// Now decode the top-level Package metadata; this will automatically recurse throughout the whole structure.
	md := mapper.New(customDecoders())
	var pack pack.Package
	if err := md.Decode(tree, &pack); err != nil {
		return nil, err
	}
	return &pack, nil
}

// customDecoders makes a complete map of all known custom AST decoders.  In general, any polymorphic node kind that
// appears as a field in another concrete marshalable structure  must have an associated custom decoder.  If not, the
// Mapper will error out.  This is typically an interface type and that method typically switches on the kind property.
// Note that interfaces that are used as "markers" and don't show up in fields are okay and don't require a decoder.
func customDecoders() mapper.Decoders {
	return mapper.Decoders{
		reflect.TypeOf((*ast.ModuleMember)(nil)).Elem(): moduleMemberDecoder,
		reflect.TypeOf((*ast.ClassMember)(nil)).Elem():  classMemberDecoder,
		reflect.TypeOf((*ast.Statement)(nil)).Elem():    statementDecoder,
		reflect.TypeOf((*ast.Expression)(nil)).Elem():   expressionDecoder,
	}
}

// Each of the custom decoders is a varaible that points to a decoder function; this is done so that the decode*
// functions can remain strongly typed, as the mapper's decoder signature requires a weakly-typed interface{} return.

var moduleMemberDecoder = func(m mapper.Mapper, tree mapper.Object) (interface{}, error) {
	return decodeModuleMember(m, tree)
}
var classMemberDecoder = func(m mapper.Mapper, tree mapper.Object) (interface{}, error) {
	return decodeClassMember(m, tree)
}
var statementDecoder = func(m mapper.Mapper, tree mapper.Object) (interface{}, error) {
	return decodeStatement(m, tree)
}
var expressionDecoder = func(m mapper.Mapper, tree mapper.Object) (interface{}, error) {
	return decodeExpression(m, tree)
}