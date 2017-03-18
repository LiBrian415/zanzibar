// Code generated by thriftrw v1.0.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/clients/foo/foo"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_TooManyArgs_Args struct {
	Request *BarRequest    `json:"request"`
	Foo     *foo.FooStruct `json:"foo,omitempty"`
}

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

func (v *Bar_TooManyArgs_Args) String() string {
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

func (v *Bar_TooManyArgs_Args) MethodName() string {
	return "tooManyArgs"
}

func (v *Bar_TooManyArgs_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_TooManyArgs_Helper = struct {
	Args           func(request *BarRequest, foo *foo.FooStruct) *Bar_TooManyArgs_Args
	IsException    func(error) bool
	WrapResponse   func(*BarResponse, error) (*Bar_TooManyArgs_Result, error)
	UnwrapResponse func(*Bar_TooManyArgs_Result) (*BarResponse, error)
}{}

func init() {
	Bar_TooManyArgs_Helper.Args = func(request *BarRequest, foo *foo.FooStruct) *Bar_TooManyArgs_Args {
		return &Bar_TooManyArgs_Args{Request: request, Foo: foo}
	}
	Bar_TooManyArgs_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
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
		}
		return nil, err
	}
	Bar_TooManyArgs_Helper.UnwrapResponse = func(result *Bar_TooManyArgs_Result) (success *BarResponse, err error) {
		if result.BarException != nil {
			err = result.BarException
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

type Bar_TooManyArgs_Result struct {
	Success      *BarResponse  `json:"success,omitempty"`
	BarException *BarException `json:"barException,omitempty"`
}

func (v *Bar_TooManyArgs_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
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
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_TooManyArgs_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

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
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_TooManyArgs_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_TooManyArgs_Result) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}
	return fmt.Sprintf("Bar_TooManyArgs_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_TooManyArgs_Result) MethodName() string {
	return "tooManyArgs"
}

func (v *Bar_TooManyArgs_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
