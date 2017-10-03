// Code generated by zanzibar
// @generated
// Checksum : wybJ0UraYROv+n1RAzXyPQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(in *jlexer.Lexer, out *SecondService_EchoEnum_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(Fruit)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Success).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(out *jwriter.Writer, in SecondService_EchoEnum_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Success).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoEnum_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoEnum_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoEnum_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoEnum_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum(l, v)
}
func easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(in *jlexer.Lexer, out *SecondService_EchoEnum_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "arg":
			if in.IsNull() {
				in.Skip()
				out.Arg = nil
			} else {
				if out.Arg == nil {
					out.Arg = new(Fruit)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Arg).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(out *jwriter.Writer, in SecondService_EchoEnum_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Arg != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"arg\":")
		if in.Arg == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Arg).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoEnum_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoEnum_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a651e90EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoEnum_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoEnum_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a651e90DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoEnum1(l, v)
}
