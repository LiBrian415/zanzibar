// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/foo"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Bar_TooManyArgs_Args represents the arguments for the Bar.tooManyArgs function.
//
// The arguments for tooManyArgs are sent and received over the wire as this struct.
type Bar_TooManyArgs_Args struct {
	Request *BarRequest    `json:"request,required"`
	Foo     *foo.FooStruct `json:"foo,omitempty"`
}

// ToWire translates a Bar_TooManyArgs_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_TooManyArgs_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request == nil {
		return w, errors.New("field Request of Bar_TooManyArgs_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Foo != nil {
		w, err = v.Foo.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _FooStruct_Read(w wire.Value) (*foo.FooStruct, error) {
	var v foo.FooStruct
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_TooManyArgs_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_TooManyArgs_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_TooManyArgs_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_TooManyArgs_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _BarRequest_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Foo, err = _FooStruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_TooManyArgs_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_TooManyArgs_Args
// struct.
func (v *Bar_TooManyArgs_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++
	if v.Foo != nil {
		fields[i] = fmt.Sprintf("Foo: %v", v.Foo)
		i++
	}

	return fmt.Sprintf("Bar_TooManyArgs_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_TooManyArgs_Args match the
// provided Bar_TooManyArgs_Args.
//
// This function performs a deep comparison.
func (v *Bar_TooManyArgs_Args) Equals(rhs *Bar_TooManyArgs_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Request.Equals(rhs.Request) {
		return false
	}
	if !((v.Foo == nil && rhs.Foo == nil) || (v.Foo != nil && rhs.Foo != nil && v.Foo.Equals(rhs.Foo))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_TooManyArgs_Args.
func (v *Bar_TooManyArgs_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("request", v.Request))
	if v.Foo != nil {
		err = multierr.Append(err, enc.AddObject("foo", v.Foo))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_TooManyArgs_Args) GetRequest() (o *BarRequest) { return v.Request }

// IsSetRequest returns true if Request is not nil.
func (v *Bar_TooManyArgs_Args) IsSetRequest() bool {
	return v.Request != nil
}

// GetFoo returns the value of Foo if it is set or its
// zero value if it is unset.
func (v *Bar_TooManyArgs_Args) GetFoo() (o *foo.FooStruct) {
	if v.Foo != nil {
		return v.Foo
	}

	return
}

// IsSetFoo returns true if Foo is not nil.
func (v *Bar_TooManyArgs_Args) IsSetFoo() bool {
	return v.Foo != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "tooManyArgs" for this struct.
func (v *Bar_TooManyArgs_Args) MethodName() string {
	return "tooManyArgs"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_TooManyArgs_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_TooManyArgs_Helper provides functions that aid in handling the
// parameters and return values of the Bar.tooManyArgs
// function.
var Bar_TooManyArgs_Helper = struct {
	// Args accepts the parameters of tooManyArgs in-order and returns
	// the arguments struct for the function.
	Args func(
		request *BarRequest,
		foo *foo.FooStruct,
	) *Bar_TooManyArgs_Args

	// IsException returns true if the given error can be thrown
	// by tooManyArgs.
	//
	// An error can be thrown by tooManyArgs only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for tooManyArgs
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// tooManyArgs into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by tooManyArgs
	//
	//   value, err := tooManyArgs(args)
	//   result, err := Bar_TooManyArgs_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from tooManyArgs: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponse, error) (*Bar_TooManyArgs_Result, error)

	// UnwrapResponse takes the result struct for tooManyArgs
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if tooManyArgs threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_TooManyArgs_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_TooManyArgs_Result) (*BarResponse, error)
}{}

func init() {
	Bar_TooManyArgs_Helper.Args = func(
		request *BarRequest,
		foo *foo.FooStruct,
	) *Bar_TooManyArgs_Args {
		return &Bar_TooManyArgs_Args{
			Request: request,
			Foo:     foo,
		}
	}

	Bar_TooManyArgs_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		case *foo.FooException:
			return true
		default:
			return false
		}
	}

	Bar_TooManyArgs_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_TooManyArgs_Result, error) {
		if err == nil {
			return &Bar_TooManyArgs_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_TooManyArgs_Result.BarException")
			}
			return &Bar_TooManyArgs_Result{BarException: e}, nil
		case *foo.FooException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_TooManyArgs_Result.FooException")
			}
			return &Bar_TooManyArgs_Result{FooException: e}, nil
		}

		return nil, err
	}
	Bar_TooManyArgs_Helper.UnwrapResponse = func(result *Bar_TooManyArgs_Result) (success *BarResponse, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}
		if result.FooException != nil {
			err = result.FooException
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_TooManyArgs_Result represents the result of a Bar.tooManyArgs function call.
//
// The result of a tooManyArgs execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_TooManyArgs_Result struct {
	// Value returned by tooManyArgs after a successful execution.
	Success      *BarResponse      `json:"success,omitempty"`
	BarException *BarException     `json:"barException,omitempty"`
	FooException *foo.FooException `json:"fooException,omitempty"`
}

// ToWire translates a Bar_TooManyArgs_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_TooManyArgs_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.FooException != nil {
		w, err = v.FooException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_TooManyArgs_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _FooException_Read(w wire.Value) (*foo.FooException, error) {
	var v foo.FooException
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_TooManyArgs_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_TooManyArgs_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_TooManyArgs_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_TooManyArgs_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.FooException, err = _FooException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if v.FooException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_TooManyArgs_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_TooManyArgs_Result
// struct.
func (v *Bar_TooManyArgs_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}
	if v.FooException != nil {
		fields[i] = fmt.Sprintf("FooException: %v", v.FooException)
		i++
	}

	return fmt.Sprintf("Bar_TooManyArgs_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_TooManyArgs_Result match the
// provided Bar_TooManyArgs_Result.
//
// This function performs a deep comparison.
func (v *Bar_TooManyArgs_Result) Equals(rhs *Bar_TooManyArgs_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}
	if !((v.FooException == nil && rhs.FooException == nil) || (v.FooException != nil && rhs.FooException != nil && v.FooException.Equals(rhs.FooException))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_TooManyArgs_Result.
func (v *Bar_TooManyArgs_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BarException != nil {
		err = multierr.Append(err, enc.AddObject("barException", v.BarException))
	}
	if v.FooException != nil {
		err = multierr.Append(err, enc.AddObject("fooException", v.FooException))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_TooManyArgs_Result) GetSuccess() (o *BarResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_TooManyArgs_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// GetBarException returns the value of BarException if it is set or its
// zero value if it is unset.
func (v *Bar_TooManyArgs_Result) GetBarException() (o *BarException) {
	if v.BarException != nil {
		return v.BarException
	}

	return
}

// IsSetBarException returns true if BarException is not nil.
func (v *Bar_TooManyArgs_Result) IsSetBarException() bool {
	return v.BarException != nil
}

// GetFooException returns the value of FooException if it is set or its
// zero value if it is unset.
func (v *Bar_TooManyArgs_Result) GetFooException() (o *foo.FooException) {
	if v.FooException != nil {
		return v.FooException
	}

	return
}

// IsSetFooException returns true if FooException is not nil.
func (v *Bar_TooManyArgs_Result) IsSetFooException() bool {
	return v.FooException != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "tooManyArgs" for this struct.
func (v *Bar_TooManyArgs_Result) MethodName() string {
	return "tooManyArgs"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_TooManyArgs_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
