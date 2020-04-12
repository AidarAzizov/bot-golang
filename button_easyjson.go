// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package botgolang

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

func easyjsonF248ab8DecodeBotGolang(in *jlexer.Lexer, out *ButtonResponse) {
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
		case "queryId":
			out.QueryID = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "showAlert":
			out.ShowAlert = bool(in.Bool())
		case "url":
			out.URL = string(in.String())
		case "callbackData":
			out.CallbackData = string(in.String())
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
func easyjsonF248ab8EncodeBotGolang(out *jwriter.Writer, in ButtonResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"queryId\":"
		out.RawString(prefix[1:])
		out.String(string(in.QueryID))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"showAlert\":"
		out.RawString(prefix)
		out.Bool(bool(in.ShowAlert))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"callbackData\":"
		out.RawString(prefix)
		out.String(string(in.CallbackData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ButtonResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF248ab8EncodeBotGolang(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ButtonResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF248ab8EncodeBotGolang(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ButtonResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF248ab8DecodeBotGolang(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ButtonResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF248ab8DecodeBotGolang(l, v)
}
func easyjsonF248ab8DecodeBotGolang1(in *jlexer.Lexer, out *Button) {
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
		case "text":
			out.Text = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "callbackData":
			out.CallbackData = string(in.String())
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
func easyjsonF248ab8EncodeBotGolang1(out *jwriter.Writer, in Button) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	if in.CallbackData != "" {
		const prefix string = ",\"callbackData\":"
		out.RawString(prefix)
		out.String(string(in.CallbackData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Button) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF248ab8EncodeBotGolang1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Button) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF248ab8EncodeBotGolang1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Button) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF248ab8DecodeBotGolang1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Button) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF248ab8DecodeBotGolang1(l, v)
}
