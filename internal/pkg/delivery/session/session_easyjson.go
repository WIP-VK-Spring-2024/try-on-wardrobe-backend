// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package session

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

func easyjsonA818f49aDecodeTryOnInternalPkgDeliverySession(in *jlexer.Lexer, out *tokenResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		case "user_id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.UserID).UnmarshalText(data))
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
func easyjsonA818f49aEncodeTryOnInternalPkgDeliverySession(out *jwriter.Writer, in tokenResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Token != "" {
		const prefix string = ",\"token\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	if (in.UserID).IsDefined() {
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.UserID).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v tokenResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA818f49aEncodeTryOnInternalPkgDeliverySession(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v tokenResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA818f49aEncodeTryOnInternalPkgDeliverySession(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *tokenResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA818f49aDecodeTryOnInternalPkgDeliverySession(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *tokenResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA818f49aDecodeTryOnInternalPkgDeliverySession(l, v)
}
