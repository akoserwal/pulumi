// Copyright 2016 Marapongo, Inc. All rights reserved.

package eval

import (
	"github.com/marapongo/mu/pkg/compiler/ast"
	"github.com/marapongo/mu/pkg/compiler/symbols"
	"github.com/marapongo/mu/pkg/diag"
	"github.com/marapongo/mu/pkg/eval/rt"
	"github.com/marapongo/mu/pkg/tokens"
	"github.com/marapongo/mu/pkg/util/contract"
)

// Invoker implements an intrinsic function's functionality.
type Invoker func(intrin *Intrinsic, e *evaluator, this *rt.Object, args []*rt.Object) *rt.Unwind

// Intrinsics contains the set of runtime functions that are callable by name through the Mu standard library
// package.  Their functionality is implemented in the runtime because MuIL cannot express the concepts they require to
// get their job done.  This includes things like dynamic introspection, invocation, and more.
var Intrinsics map[tokens.Token]Invoker

func init() {
	Intrinsics = map[tokens.Token]Invoker{
		"mu:runtime:isFunction":    isFunction,
		"mu:runtime:dynamicInvoke": dynamicInvoke,
	}
}

// Intrinsic is a special intrinsic function whose behavior is implemented by the runtime.
type Intrinsic struct {
	Node    ast.Node     // the contextual node representing the place where this intrinsic got created.
	Func    ast.Function // the underlying function's node (before mapping to an intrinsic).
	Nm      tokens.Name
	Tok     tokens.Token
	Sig     *symbols.FunctionType
	Invoker Invoker
}

var _ symbols.Function = (*Intrinsic)(nil)

func (node *Intrinsic) Symbol()                          {}
func (node *Intrinsic) Name() tokens.Name                { return node.Nm }
func (node *Intrinsic) Token() tokens.Token              { return node.Tok }
func (node *Intrinsic) Special() bool                    { return false }
func (node *Intrinsic) Tree() diag.Diagable              { return node.Func }
func (node *Intrinsic) Function() ast.Function           { return node.Func }
func (node *Intrinsic) Signature() *symbols.FunctionType { return node.Sig }
func (node *Intrinsic) String() string                   { return string(node.Name()) }

// Invoke calls the intrinsic routine, passing additional context from the intrinsic struct itself.
func (node *Intrinsic) Invoke(e *evaluator, this *rt.Object, args []*rt.Object) *rt.Unwind {
	return node.Invoker(node, e, this, args)

}

// newIntrinsic returns a new intrinsic function symbol with the given information.
func NewIntrinsic(node ast.Node, fnc ast.Function, tok tokens.Token, nm tokens.Name,
	sig *symbols.FunctionType, invoker Invoker) *Intrinsic {
	return &Intrinsic{
		Node:    node,
		Func:    fnc,
		Nm:      nm,
		Tok:     tok,
		Sig:     sig,
		Invoker: invoker,
	}
}

// MaybeIntrinsic checks whether the given symbol is an intrinsic and, if so, swaps it out with the actual runtime
// implementation of that intrinsic.  If the symbol is not an intrinsic, the original symbol is simply returned.
func MaybeIntrinsic(node ast.Node, sym symbols.Symbol) symbols.Symbol {
	switch s := sym.(type) {
	case *Intrinsic:
		// Already an intrinsic; do not swap it out.
	case symbols.Function:
		// If this is a function whose token we recognize, create a new intrinsic symbol.  Note that we do not currently
		// cache these symbols because of the need to associate the AST node with the resulting symbol.
		tok := s.Token()
		if invoker, isintrinsic := Intrinsics[tok]; isintrinsic {
			sym = NewIntrinsic(node, s.Function(), tok, tok.Name(), s.Signature(), invoker)
		}
	}
	return sym
}

func isFunction(intrin *Intrinsic, e *evaluator, this *rt.Object, args []*rt.Object) *rt.Unwind {
	contract.Assert(this == nil)    // module function
	contract.Assert(len(args) == 1) // just one arg: the object to inquire about functionness
	_, isfunc := args[0].Type().(*symbols.FunctionType)
	ret := e.alloc.NewBool(isfunc)
	return rt.NewReturnUnwind(ret)
}

func dynamicInvoke(intrin *Intrinsic, e *evaluator, this *rt.Object, args []*rt.Object) *rt.Unwind {
	contract.Assert(this == nil)    // module function
	contract.Assert(len(args) == 3) // three args: obj, thisArg, and args

	// First ensure the target is a function.
	t := args[0].Type()
	if _, isfunc := t.(*symbols.FunctionType); !isfunc {
		return e.NewIllegalInvokeTargetException(intrin.Node, t)
	}

	// Fetch the function stub information (ignoring `this`, since we will pass our own).
	stub := args[0].FunctionValue()

	// Next, simply call through to the evalCall function, which will do all additional verification.
	obj, uw := e.evalCall(intrin.Node, stub.Func, this, args...)
	if uw != nil {
		return uw
	}
	return rt.NewReturnUnwind(obj)
}