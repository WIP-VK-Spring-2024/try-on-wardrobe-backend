// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package feed

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

func easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed(in *jlexer.Lexer, out *rateRequest) {
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
		case "rating":
			out.Rating = int(in.Int())
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
func easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed(out *jwriter.Writer, in rateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Rating != 0 {
		const prefix string = ",\"rating\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Rating))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v rateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v rateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *rateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *rateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed(l, v)
}
func easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed1(in *jlexer.Lexer, out *createCommentResponse) {
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
		case "uuid":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Uuid).UnmarshalText(data))
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
func easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed1(out *jwriter.Writer, in createCommentResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if (in.Uuid).IsDefined() {
		const prefix string = ",\"uuid\":"
		first = false
		out.RawString(prefix[1:])
		out.RawText((in.Uuid).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createCommentResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createCommentResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD77e0694EncodeTryOnInternalPkgDeliveryFeed1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createCommentResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createCommentResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD77e0694DecodeTryOnInternalPkgDeliveryFeed1(l, v)
}
