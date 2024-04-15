// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package weather

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

func easyjson4f98044aDecodeTryOnInternalPkgRepositoryWeather(in *jlexer.Lexer, out *weatherApiResponse) {
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
		case "location":
			(out.Location).UnmarshalEasyJSON(in)
		case "current":
			(out.Current).UnmarshalEasyJSON(in)
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
func easyjson4f98044aEncodeTryOnInternalPkgRepositoryWeather(out *jwriter.Writer, in weatherApiResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"location\":"
		first = false
		out.RawString(prefix[1:])
		(in.Location).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"current\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Current).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v weatherApiResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f98044aEncodeTryOnInternalPkgRepositoryWeather(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v weatherApiResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f98044aEncodeTryOnInternalPkgRepositoryWeather(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *weatherApiResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f98044aDecodeTryOnInternalPkgRepositoryWeather(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *weatherApiResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f98044aDecodeTryOnInternalPkgRepositoryWeather(l, v)
}