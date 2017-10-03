// Code generated by zanzibar
// @generated
// Checksum : B2+P1BPi6lc35v8+HqAuqA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(in *jlexer.Lexer, out *Bar_NormalRecur_Result) {
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
					out.Success = new(BarResponseRecur)
				}
				easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*out.Success)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in, &*out.BarException)
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
func easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(out *jwriter.Writer, in Bar_NormalRecur_Result) {
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
			easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.Success)
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out, *in.BarException)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_NormalRecur_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_NormalRecur_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_NormalRecur_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_NormalRecur_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur(l, v)
}
func easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	out.RawByte('}')
}
func easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarResponseRecur) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var NodesSet bool
	var HeightSet bool
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
		case "nodes":
			if in.IsNull() {
				in.Skip()
				out.Nodes = nil
			} else {
				in.Delim('[')
				if out.Nodes == nil {
					if !in.IsDelim(']') {
						out.Nodes = make([]string, 0, 4)
					} else {
						out.Nodes = []string{}
					}
				} else {
					out.Nodes = (out.Nodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Nodes = append(out.Nodes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
			NodesSet = true
		case "height":
			out.Height = int32(in.Int32())
			HeightSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !NodesSet {
		in.AddError(fmt.Errorf("key 'nodes' is required"))
	}
	if !HeightSet {
		in.AddError(fmt.Errorf("key 'height' is required"))
	}
}
func easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarResponseRecur) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodes\":")
	if in.Nodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Nodes {
			if v2 > 0 {
				out.RawByte(',')
			}
			out.String(string(v3))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"height\":")
	out.Int32(int32(in.Height))
	out.RawByte('}')
}
func easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(in *jlexer.Lexer, out *Bar_NormalRecur_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequestRecur)
				}
				easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(in, &*out.Request)
			}
			RequestSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(out *jwriter.Writer, in Bar_NormalRecur_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(out, *in.Request)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_NormalRecur_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_NormalRecur_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_NormalRecur_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_NormalRecur_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNormalRecur1(l, v)
}
func easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(in *jlexer.Lexer, out *BarRequestRecur) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var NameSet bool
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
		case "name":
			out.Name = string(in.String())
			NameSet = true
		case "recur":
			if in.IsNull() {
				in.Skip()
				out.Recur = nil
			} else {
				if out.Recur == nil {
					out.Recur = new(BarRequestRecur)
				}
				easyjson47f5be94DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(in, &*out.Recur)
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
	if !NameSet {
		in.AddError(fmt.Errorf("key 'name' is required"))
	}
}
func easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(out *jwriter.Writer, in BarRequestRecur) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if in.Recur != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"recur\":")
		if in.Recur == nil {
			out.RawString("null")
		} else {
			easyjson47f5be94EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(out, *in.Recur)
		}
	}
	out.RawByte('}')
}
