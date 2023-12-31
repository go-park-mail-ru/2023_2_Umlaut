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

func easyjson9e1087fdDecodeGithubComGoParkMailRu20232UmlautPkgModelCore(in *jlexer.Lexer, out *User) {
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
		case "name":
			out.Name = string(in.String())
		case "mail":
			out.Mail = string(in.String())
		case "password":
			out.PasswordHash = string(in.String())
		case "user_gender":
			if in.IsNull() {
				in.Skip()
				out.UserGender = nil
			} else {
				if out.UserGender == nil {
					out.UserGender = new(int)
				}
				*out.UserGender = int(in.Int())
			}
		case "prefer_gender":
			if in.IsNull() {
				in.Skip()
				out.PreferGender = nil
			} else {
				if out.PreferGender == nil {
					out.PreferGender = new(int)
				}
				*out.PreferGender = int(in.Int())
			}
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
			}
		case "age":
			if in.IsNull() {
				in.Skip()
				out.Age = nil
			} else {
				if out.Age == nil {
					out.Age = new(int)
				}
				*out.Age = int(in.Int())
			}
		case "looking":
			if in.IsNull() {
				in.Skip()
				out.Looking = nil
			} else {
				if out.Looking == nil {
					out.Looking = new(string)
				}
				*out.Looking = string(in.String())
			}
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
		case "education":
			if in.IsNull() {
				in.Skip()
				out.Education = nil
			} else {
				if out.Education == nil {
					out.Education = new(string)
				}
				*out.Education = string(in.String())
			}
		case "hobbies":
			if in.IsNull() {
				in.Skip()
				out.Hobbies = nil
			} else {
				if out.Hobbies == nil {
					out.Hobbies = new(string)
				}
				*out.Hobbies = string(in.String())
			}
		case "birthday":
			if in.IsNull() {
				in.Skip()
				out.Birthday = nil
			} else {
				if out.Birthday == nil {
					out.Birthday = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Birthday).UnmarshalJSON(data))
				}
			}
		case "online":
			out.Online = bool(in.Bool())
		case "role":
			out.Role = int(in.Int())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				if out.Tags == nil {
					out.Tags = new([]string)
				}
				if in.IsNull() {
					in.Skip()
					*out.Tags = nil
				} else {
					in.Delim('[')
					if *out.Tags == nil {
						if !in.IsDelim(']') {
							*out.Tags = make([]string, 0, 4)
						} else {
							*out.Tags = []string{}
						}
					} else {
						*out.Tags = (*out.Tags)[:0]
					}
					for !in.IsDelim(']') {
						var v2 string
						v2 = string(in.String())
						*out.Tags = append(*out.Tags, v2)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
		case "oauthId":
			if in.IsNull() {
				in.Skip()
				out.OauthId = nil
			} else {
				if out.OauthId == nil {
					out.OauthId = new(int)
				}
				*out.OauthId = int(in.Int())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20232UmlautPkgModelCore(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"mail\":"
		out.RawString(prefix)
		out.String(string(in.Mail))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.PasswordHash))
	}
	{
		const prefix string = ",\"user_gender\":"
		out.RawString(prefix)
		if in.UserGender == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.UserGender))
		}
	}
	{
		const prefix string = ",\"prefer_gender\":"
		out.RawString(prefix)
		if in.PreferGender == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.PreferGender))
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		if in.Age == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Age))
		}
	}
	{
		const prefix string = ",\"looking\":"
		out.RawString(prefix)
		if in.Looking == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Looking))
		}
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
				for v3, v4 := range *in.ImagePaths {
					if v3 > 0 {
						out.RawByte(',')
					}
					out.String(string(v4))
				}
				out.RawByte(']')
			}
		}
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		if in.Education == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Education))
		}
	}
	{
		const prefix string = ",\"hobbies\":"
		out.RawString(prefix)
		if in.Hobbies == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Hobbies))
		}
	}
	{
		const prefix string = ",\"birthday\":"
		out.RawString(prefix)
		if in.Birthday == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Birthday).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"online\":"
		out.RawString(prefix)
		out.Bool(bool(in.Online))
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.Int(int(in.Role))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil {
			out.RawString("null")
		} else {
			if *in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v5, v6 := range *in.Tags {
					if v5 > 0 {
						out.RawByte(',')
					}
					out.String(string(v6))
				}
				out.RawByte(']')
			}
		}
	}
	{
		const prefix string = ",\"oauthId\":"
		out.RawString(prefix)
		if in.OauthId == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.OauthId))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20232UmlautPkgModelCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20232UmlautPkgModelCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20232UmlautPkgModelCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20232UmlautPkgModelCore(l, v)
}
