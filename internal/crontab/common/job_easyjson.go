// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package common

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

func easyjson8a33d6c7DecodeGoCrawlerDistributedInternalCrontabCommon(in *jlexer.Lexer, out *Job) {
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
		case "name":
			out.Name = string(in.String())
		case "command":
			out.Command = string(in.String())
		case "cron_expr":
			out.CronExpr = string(in.String())
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
func easyjson8a33d6c7EncodeGoCrawlerDistributedInternalCrontabCommon(out *jwriter.Writer, in Job) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"command\":"
		out.RawString(prefix)
		out.String(string(in.Command))
	}
	{
		const prefix string = ",\"cron_expr\":"
		out.RawString(prefix)
		out.String(string(in.CronExpr))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Job) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a33d6c7EncodeGoCrawlerDistributedInternalCrontabCommon(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Job) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a33d6c7EncodeGoCrawlerDistributedInternalCrontabCommon(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Job) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a33d6c7DecodeGoCrawlerDistributedInternalCrontabCommon(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Job) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a33d6c7DecodeGoCrawlerDistributedInternalCrontabCommon(l, v)
}
