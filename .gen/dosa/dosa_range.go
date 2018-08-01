// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Dosa_Range_Args represents the arguments for the Dosa.range function.
//
// The arguments for range are sent and received over the wire as this struct.
type Dosa_Range_Args struct {
	Request *RangeRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_Range_Args struct into a Thrift-level intermediate
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
func (v *Dosa_Range_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RangeRequest_Read(w wire.Value) (*RangeRequest, error) {
	var v RangeRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_Range_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Range_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Range_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Range_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _RangeRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_Range_Args
// struct.
func (v *Dosa_Range_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_Range_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Range_Args match the
// provided Dosa_Range_Args.
//
// This function performs a deep comparison.
func (v *Dosa_Range_Args) Equals(rhs *Dosa_Range_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Dosa_Range_Args) GetRequest() (o *RangeRequest) {
	if v.Request != nil {
		return v.Request
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "range" for this struct.
func (v *Dosa_Range_Args) MethodName() string {
	return "range"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_Range_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_Range_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.range
// function.
var Dosa_Range_Helper = struct {
	// Args accepts the parameters of range in-order and returns
	// the arguments struct for the function.
	Args func(
		request *RangeRequest,
	) *Dosa_Range_Args

	// IsException returns true if the given error can be thrown
	// by range.
	//
	// An error can be thrown by range only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for range
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// range into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by range
	//
	//   value, err := range(args)
	//   result, err := Dosa_Range_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from range: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*RangeResponse, error) (*Dosa_Range_Result, error)

	// UnwrapResponse takes the result struct for range
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if range threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Dosa_Range_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_Range_Result) (*RangeResponse, error)
}{}

func init() {
	Dosa_Range_Helper.Args = func(
		request *RangeRequest,
	) *Dosa_Range_Args {
		return &Dosa_Range_Args{
			Request: request,
		}
	}

	Dosa_Range_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_Range_Helper.WrapResponse = func(success *RangeResponse, err error) (*Dosa_Range_Result, error) {
		if err == nil {
			return &Dosa_Range_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Range_Result.ClientError")
			}
			return &Dosa_Range_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Range_Result.ServerError")
			}
			return &Dosa_Range_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_Range_Helper.UnwrapResponse = func(result *Dosa_Range_Result) (success *RangeResponse, err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
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

// Dosa_Range_Result represents the result of a Dosa.range function call.
//
// The result of a range execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Dosa_Range_Result struct {
	// Value returned by range after a successful execution.
	Success     *RangeResponse       `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_Range_Result struct into a Thrift-level intermediate
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
func (v *Dosa_Range_Result) ToWire() (wire.Value, error) {
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
	if v.ClientError != nil {
		w, err = v.ClientError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerError != nil {
		w, err = v.ServerError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Dosa_Range_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RangeResponse_Read(w wire.Value) (*RangeResponse, error) {
	var v RangeResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_Range_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Range_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Range_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Range_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _RangeResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ClientError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerError, err = _InternalServerError_Read(field.Value)
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
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Dosa_Range_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_Range_Result
// struct.
func (v *Dosa_Range_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}

	return fmt.Sprintf("Dosa_Range_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Range_Result match the
// provided Dosa_Range_Result.
//
// This function performs a deep comparison.
func (v *Dosa_Range_Result) Equals(rhs *Dosa_Range_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Dosa_Range_Result) GetSuccess() (o *RangeResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetClientError returns the value of ClientError if it is set or its
// zero value if it is unset.
func (v *Dosa_Range_Result) GetClientError() (o *BadRequestError) {
	if v.ClientError != nil {
		return v.ClientError
	}

	return
}

// GetServerError returns the value of ServerError if it is set or its
// zero value if it is unset.
func (v *Dosa_Range_Result) GetServerError() (o *InternalServerError) {
	if v.ServerError != nil {
		return v.ServerError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "range" for this struct.
func (v *Dosa_Range_Result) MethodName() string {
	return "range"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_Range_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
