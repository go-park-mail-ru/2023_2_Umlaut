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

func easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto(in *jlexer.Lexer, out *FilterParams) {
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
		case "user_id":
			out.UserId = int(in.Int())
		case "min_age":
			out.MinAge = int(in.Int())
		case "max_age":
			out.MaxAge = int(in.Int())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto(out *jwriter.Writer, in FilterParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.UserId))
	}
	{
		const prefix string = ",\"min_age\":"
		out.RawString(prefix)
		out.Int(int(in.MinAge))
	}
	{
		const prefix string = ",\"max_age\":"
		out.RawString(prefix)
		out.Int(int(in.MaxAge))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tags {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilterParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilterParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilterParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilterParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto(l, v)
}
func easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto1(in *jlexer.Lexer, out *FeedData) {
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
		case "user":
			(out.User).UnmarshalEasyJSON(in)
		case "like_counter":
			out.LikeCounter = int(in.Int())
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
func easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto1(out *jwriter.Writer, in FeedData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		(in.User).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"like_counter\":"
		out.RawString(prefix)
		out.Int(int(in.LikeCounter))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FeedData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FeedData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD77e0694EncodeGithubComGoParkMailRu20232UmlautInternalModelDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FeedData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FeedData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD77e0694DecodeGithubComGoParkMailRu20232UmlautInternalModelDto1(l, v)
}
