// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package outfits

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

func easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits(in *jlexer.Lexer, out *getOutfitsParams) {
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
		case "limit":
			out.Limit = int(in.Int())
		case "since":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Since).UnmarshalJSON(data))
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
func easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits(out *jwriter.Writer, in getOutfitsParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Limit != 0 {
		const prefix string = ",\"limit\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Limit))
	}
	if true {
		const prefix string = ",\"since\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Since).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getOutfitsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getOutfitsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getOutfitsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getOutfitsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits(l, v)
}
func easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits1(in *jlexer.Lexer, out *createdResponse) {
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
		case "image":
			out.Image = string(in.String())
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
func easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits1(out *jwriter.Writer, in createdResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if (in.Uuid).IsDefined() {
		const prefix string = ",\"uuid\":"
		first = false
		out.RawString(prefix[1:])
		out.RawText((in.Uuid).MarshalText())
	}
	if in.Image != "" {
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createdResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createdResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE39f981aEncodeTryOnInternalPkgDeliveryOutfits1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createdResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createdResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE39f981aDecodeTryOnInternalPkgDeliveryOutfits1(l, v)
}
