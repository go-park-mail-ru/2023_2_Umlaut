// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

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

func easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelDto(in *jlexer.Lexer, out *PremiumLike) {
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
func easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelDto(out *jwriter.Writer, in PremiumLike) {
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
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PremiumLike) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52421b6dEncodeGithubComGoParkMailRu20232UmlautInternalModelDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PremiumLike) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PremiumLike) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52421b6dDecodeGithubComGoParkMailRu20232UmlautInternalModelDto(l, v)
}
