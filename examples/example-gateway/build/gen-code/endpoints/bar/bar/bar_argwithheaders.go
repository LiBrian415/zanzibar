// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Bar_ArgWithHeaders_Args represents the arguments for the Bar.argWithHeaders function.
//
// The arguments for argWithHeaders are sent and received over the wire as this struct.
type Bar_ArgWithHeaders_Args struct {
	Name     string  `json:"name,required"`
	UserUUID *string `json:"-"`
}

// ToWire translates a Bar_ArgWithHeaders_Args struct into a Thrift-level intermediate
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
func (v *Bar_ArgWithHeaders_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.UserUUID != nil {
		w, err = wire.NewValueString(*(v.UserUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ArgWithHeaders_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithHeaders_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithHeaders_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithHeaders_Args) FromWire(w wire.Value) error {
	var err error

	nameIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UserUUID = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !nameIsSet {
		return errors.New("field Name of Bar_ArgWithHeaders_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithHeaders_Args
// struct.
func (v *Bar_ArgWithHeaders_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.UserUUID != nil {
		fields[i] = fmt.Sprintf("UserUUID: %v", *(v.UserUUID))
		i++
	}

	return fmt.Sprintf("Bar_ArgWithHeaders_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgWithHeaders_Args match the
// provided Bar_ArgWithHeaders_Args.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithHeaders_Args) Equals(rhs *Bar_ArgWithHeaders_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Name == rhs.Name) {
		return false
	}
	if !_String_EqualsPtr(v.UserUUID, rhs.UserUUID) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithHeaders_Args.
func (v *Bar_ArgWithHeaders_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("name", v.Name)
	if v.UserUUID != nil {
		enc.AddString("userUUID", *v.UserUUID)
	}
	return err
}

// GetName returns the value of Name if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithHeaders_Args) GetName() (o string) { return v.Name }

// GetUserUUID returns the value of UserUUID if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithHeaders_Args) GetUserUUID() (o string) {
	if v.UserUUID != nil {
		return *v.UserUUID
	}

	return
}

// IsSetUserUUID returns true if UserUUID is not nil.
func (v *Bar_ArgWithHeaders_Args) IsSetUserUUID() bool {
	return v.UserUUID != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "argWithHeaders" for this struct.
func (v *Bar_ArgWithHeaders_Args) MethodName() string {
	return "argWithHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_ArgWithHeaders_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_ArgWithHeaders_Helper provides functions that aid in handling the
// parameters and return values of the Bar.argWithHeaders
// function.
var Bar_ArgWithHeaders_Helper = struct {
	// Args accepts the parameters of argWithHeaders in-order and returns
	// the arguments struct for the function.
	Args func(
		name string,
		userUUID *string,
	) *Bar_ArgWithHeaders_Args

	// IsException returns true if the given error can be thrown
	// by argWithHeaders.
	//
	// An error can be thrown by argWithHeaders only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for argWithHeaders
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// argWithHeaders into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by argWithHeaders
	//
	//   value, err := argWithHeaders(args)
	//   result, err := Bar_ArgWithHeaders_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from argWithHeaders: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponse, error) (*Bar_ArgWithHeaders_Result, error)

	// UnwrapResponse takes the result struct for argWithHeaders
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if argWithHeaders threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_ArgWithHeaders_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_ArgWithHeaders_Result) (*BarResponse, error)
}{}

func init() {
	Bar_ArgWithHeaders_Helper.Args = func(
		name string,
		userUUID *string,
	) *Bar_ArgWithHeaders_Args {
		return &Bar_ArgWithHeaders_Args{
			Name:     name,
			UserUUID: userUUID,
		}
	}

	Bar_ArgWithHeaders_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Bar_ArgWithHeaders_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_ArgWithHeaders_Result, error) {
		if err == nil {
			return &Bar_ArgWithHeaders_Result{Success: success}, nil
		}

		return nil, err
	}
	Bar_ArgWithHeaders_Helper.UnwrapResponse = func(result *Bar_ArgWithHeaders_Result) (success *BarResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_ArgWithHeaders_Result represents the result of a Bar.argWithHeaders function call.
//
// The result of a argWithHeaders execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_ArgWithHeaders_Result struct {
	// Value returned by argWithHeaders after a successful execution.
	Success *BarResponse `json:"success,omitempty"`
}

// ToWire translates a Bar_ArgWithHeaders_Result struct into a Thrift-level intermediate
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
func (v *Bar_ArgWithHeaders_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
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

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_ArgWithHeaders_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ArgWithHeaders_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithHeaders_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithHeaders_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithHeaders_Result) FromWire(w wire.Value) error {
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
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_ArgWithHeaders_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithHeaders_Result
// struct.
func (v *Bar_ArgWithHeaders_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Bar_ArgWithHeaders_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgWithHeaders_Result match the
// provided Bar_ArgWithHeaders_Result.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithHeaders_Result) Equals(rhs *Bar_ArgWithHeaders_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithHeaders_Result.
func (v *Bar_ArgWithHeaders_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithHeaders_Result) GetSuccess() (o *BarResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_ArgWithHeaders_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "argWithHeaders" for this struct.
func (v *Bar_ArgWithHeaders_Result) MethodName() string {
	return "argWithHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_ArgWithHeaders_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
