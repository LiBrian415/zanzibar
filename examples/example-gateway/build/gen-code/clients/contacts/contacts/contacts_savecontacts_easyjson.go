// Code generated by zanzibar
// @generated
// Checksum : SimV3XUDOpm8ah4fj34vKQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package contacts

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

func easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(in *jlexer.Lexer, out *Contacts_SaveContacts_Result) {
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
					out.Success = new(SaveContactsResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
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
func easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(out *jwriter.Writer, in Contacts_SaveContacts_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Success).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contacts_SaveContacts_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contacts_SaveContacts_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contacts_SaveContacts_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contacts_SaveContacts_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts(l, v)
}
func easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(in *jlexer.Lexer, out *Contacts_SaveContacts_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var SaveContactsRequestSet bool
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
		case "saveContactsRequest":
			if in.IsNull() {
				in.Skip()
				out.SaveContactsRequest = nil
			} else {
				if out.SaveContactsRequest == nil {
					out.SaveContactsRequest = new(SaveContactsRequest)
				}
				(*out.SaveContactsRequest).UnmarshalEasyJSON(in)
			}
			SaveContactsRequestSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !SaveContactsRequestSet {
		in.AddError(fmt.Errorf("key 'saveContactsRequest' is required"))
	}
}
func easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(out *jwriter.Writer, in Contacts_SaveContacts_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"saveContactsRequest\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SaveContactsRequest == nil {
			out.RawString("null")
		} else {
			(*in.SaveContactsRequest).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contacts_SaveContacts_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contacts_SaveContacts_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3457510EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contacts_SaveContacts_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contacts_SaveContacts_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3457510DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsContactsContactsContactsSaveContacts1(l, v)
}
