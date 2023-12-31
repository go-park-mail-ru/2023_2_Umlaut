// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package core

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore(in *jlexer.Lexer, out *NeedFixObject) {
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
		case "count":
			out.Count = int32(in.Int32())
		case "comment_fix":
			if in.IsNull() {
				in.Skip()
				out.CommentFix = nil
			} else {
				in.Delim('[')
				if out.CommentFix == nil {
					if !in.IsDelim(']') {
						out.CommentFix = make([]string, 0, 4)
					} else {
						out.CommentFix = []string{}
					}
				} else {
					out.CommentFix = (out.CommentFix)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.CommentFix = append(out.CommentFix, v1)
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
func easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore(out *jwriter.Writer, in NeedFixObject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.Count))
	}
	{
		const prefix string = ",\"comment_fix\":"
		out.RawString(prefix)
		if in.CommentFix == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.CommentFix {
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
func (v NeedFixObject) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NeedFixObject) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NeedFixObject) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NeedFixObject) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore(l, v)
}
func easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore1(in *jlexer.Lexer, out *FeedbackStatistic) {
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
		case "avg-rating":
			out.AvgRating = float32(in.Float32())
		case "rating-count":
			if in.IsNull() {
				in.Skip()
				out.RatingCount = nil
			} else {
				in.Delim('[')
				if out.RatingCount == nil {
					if !in.IsDelim(']') {
						out.RatingCount = make([]int32, 0, 16)
					} else {
						out.RatingCount = []int32{}
					}
				} else {
					out.RatingCount = (out.RatingCount)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.RatingCount = append(out.RatingCount, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "liked-map":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.LikedMap = make(map[string]int32)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 int32
					v5 = int32(in.Int32())
					(out.LikedMap)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "need-fix-map":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.NeedFixMap = make(map[string]NeedFixObject)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 NeedFixObject
					(v6).UnmarshalEasyJSON(in)
					(out.NeedFixMap)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]string, 0, 4)
					} else {
						out.Comments = []string{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Comments = append(out.Comments, v7)
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
func easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore1(out *jwriter.Writer, in FeedbackStatistic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"avg-rating\":"
		out.RawString(prefix[1:])
		out.Float32(float32(in.AvgRating))
	}
	{
		const prefix string = ",\"rating-count\":"
		out.RawString(prefix)
		if in.RatingCount == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.RatingCount {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"liked-map\":"
		out.RawString(prefix)
		if in.LikedMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.LikedMap {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.Int32(int32(v10Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"need-fix-map\":"
		out.RawString(prefix)
		if in.NeedFixMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.NeedFixMap {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				(v11Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		if in.Comments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Comments {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FeedbackStatistic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FeedbackStatistic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FeedbackStatistic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FeedbackStatistic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore1(l, v)
}
func easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore2(in *jlexer.Lexer, out *Feedback) {
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
		case "id":
			out.Id = int(in.Int())
		case "user_id":
			out.UserId = int(in.Int())
		case "rating":
			if in.IsNull() {
				in.Skip()
				out.Rating = nil
			} else {
				if out.Rating == nil {
					out.Rating = new(int)
				}
				*out.Rating = int(in.Int())
			}
		case "liked":
			if in.IsNull() {
				in.Skip()
				out.Liked = nil
			} else {
				if out.Liked == nil {
					out.Liked = new(string)
				}
				*out.Liked = string(in.String())
			}
		case "need_fix":
			if in.IsNull() {
				in.Skip()
				out.NeedFix = nil
			} else {
				if out.NeedFix == nil {
					out.NeedFix = new(string)
				}
				*out.NeedFix = string(in.String())
			}
		case "comment":
			if in.IsNull() {
				in.Skip()
				out.Comment = nil
			} else {
				if out.Comment == nil {
					out.Comment = new(string)
				}
				*out.Comment = string(in.String())
			}
		case "created_at":
			if in.IsNull() {
				in.Skip()
				out.CreatedAt = nil
			} else {
				if out.CreatedAt == nil {
					out.CreatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreatedAt).UnmarshalJSON(data))
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
func easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore2(out *jwriter.Writer, in Feedback) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Int(int(in.UserId))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		if in.Rating == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Rating))
		}
	}
	{
		const prefix string = ",\"liked\":"
		out.RawString(prefix)
		if in.Liked == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Liked))
		}
	}
	{
		const prefix string = ",\"need_fix\":"
		out.RawString(prefix)
		if in.NeedFix == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NeedFix))
		}
	}
	{
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		if in.Comment == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Comment))
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		if in.CreatedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.CreatedAt).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Feedback) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Feedback) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson345d1f75EncodeGithubComGoParkMailRu20232UmlautInternalModelCore2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Feedback) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Feedback) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson345d1f75DecodeGithubComGoParkMailRu20232UmlautInternalModelCore2(l, v)
}
