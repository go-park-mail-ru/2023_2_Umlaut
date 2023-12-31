// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package core

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

func easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(in *jlexer.Lexer, out *Like) {
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
		case "liked_by_user_id":
			out.LikedByUserId = int(in.Int())
		case "liked_to_user_id":
			out.LikedToUserId = int(in.Int())
		case "is_like":
			out.IsLike = bool(in.Bool())
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
func easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(out *jwriter.Writer, in Like) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"liked_by_user_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.LikedByUserId))
	}
	{
		const prefix string = ",\"liked_to_user_id\":"
		out.RawString(prefix)
		out.Int(int(in.LikedToUserId))
	}
	{
		const prefix string = ",\"is_like\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsLike))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Like) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Like) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Like) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Like) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(l, v)
}
