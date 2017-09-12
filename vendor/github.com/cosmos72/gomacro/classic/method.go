/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * method.go
 *
 *  Created on: Mar 30, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// registerMethod registers a function value for the given receiver type recvType and method name.
// The type typ of the function must not include a receiver,
// while the function value val must include a receiver as first argument
// in particular, the equality recvType == val.Type().In(0) must be true
func (ir *ThreadGlobals) registerMethod(recvType r.Type, name string, typ r.Type, val r.Value) {
	if _, ok := ir.AllMethods[recvType][name]; ok {
		ir.Warnf("redefined method %s for <%v>", name, recvType)
	} else {
		// do not allow duplicate methods, one with pointer receiver and one with value receiver
		var altType r.Type
		if recvType.Kind() == r.Ptr {
			altType = recvType.Elem()
		} else {
			altType = r.PtrTo(recvType)
		}
		if _, ok := ir.AllMethods[altType][name]; ok {
			ir.Warnf("redefined method %s for <%v>", name, recvType)
			delete(ir.AllMethods[altType], name)
		}
	}

	methods, ok := ir.AllMethods[recvType]
	if !ok {
		methods = make(map[string]TypedValue)
		ir.AllMethods[recvType] = methods
	}
	methods[name] = TypedValue{typ: typ, val: val}
}

// ObjMethodByName returns a function value corresponding to the method
// of obj with the given name.
// The arguments to a Call on the returned function should not include
// a receiver; the returned function will always use obj as the receiver.
// It returns the zero Value if no method was found.
func (ir *ThreadGlobals) ObjMethodByName(obj r.Value, name string) r.Value {
	// search for methods known to the compiler
	val := obj.MethodByName(name)
	if val == Nil {
		// search for methods known to the intepreter
		t := obj.Type()
		if method, ok := ir.AllMethods[t][name]; ok {
			// cumbersome... we must create a closure on-the-fly
			val = r.MakeFunc(method.typ, func(args []r.Value) []r.Value {
				return method.val.Call(append([]r.Value{obj}, args...))
			})
		}
	}
	return val
}