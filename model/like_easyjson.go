// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel(in *jlexer.Lexer, out *PremiumLike) {
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
		case "image_paths":
			if in.IsNull() {
				in.Skip()
				out.ImagePaths = nil
			} else {
				if out.ImagePaths == nil {
					out.ImagePaths = new([]string)
				}
				if in.IsNull() {
					in.Skip()
					*out.ImagePaths = nil
				} else {
					in.Delim('[')
					if *out.ImagePaths == nil {
						if !in.IsDelim(']') {
							*out.ImagePaths = make([]string, 0, 4)
						} else {
							*out.ImagePaths = []string{}
						}
					} else {
						*out.ImagePaths = (*out.ImagePaths)[:0]
					}
					for !in.IsDelim(']') {
						var v1 string
						v1 = string(in.String())
						*out.ImagePaths = append(*out.ImagePaths, v1)
						in.WantComma()
					}
					in.Delim(']')
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
func easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel(out *jwriter.Writer, in PremiumLike) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"liked_by_user_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.LikedByUserId))
	}
	{
		const prefix string = ",\"image_paths\":"
		out.RawString(prefix)
		if in.ImagePaths == nil {
			out.RawString("null")
		} else {
			if *in.ImagePaths == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v2, v3 := range *in.ImagePaths {
					if v2 > 0 {
						out.RawByte(',')
					}
					out.String(string(v3))
				}
				out.RawByte(']')
			}
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PremiumLike) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PremiumLike) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PremiumLike) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PremiumLike) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel(l, v)
}
func easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel1(in *jlexer.Lexer, out *Like) {
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
func easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel1(out *jwriter.Writer, in Like) {
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
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Like) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Like) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Like) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautModel1(l, v)
}
