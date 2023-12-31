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

func easyjson4086215fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(in *jlexer.Lexer, out *Message) {
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
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(int)
				}
				*out.Id = int(in.Int())
			}
		case "sender_id":
			if in.IsNull() {
				in.Skip()
				out.SenderId = nil
			} else {
				if out.SenderId == nil {
					out.SenderId = new(int)
				}
				*out.SenderId = int(in.Int())
			}
		case "recipient_id":
			if in.IsNull() {
				in.Skip()
				out.RecipientId = nil
			} else {
				if out.RecipientId == nil {
					out.RecipientId = new(int)
				}
				*out.RecipientId = int(in.Int())
			}
		case "dialog_id":
			if in.IsNull() {
				in.Skip()
				out.DialogId = nil
			} else {
				if out.DialogId == nil {
					out.DialogId = new(int)
				}
				*out.DialogId = int(in.Int())
			}
		case "message_text":
			if in.IsNull() {
				in.Skip()
				out.Text = nil
			} else {
				if out.Text == nil {
					out.Text = new(string)
				}
				*out.Text = string(in.String())
			}
		case "is_read":
			if in.IsNull() {
				in.Skip()
				out.IsRead = nil
			} else {
				if out.IsRead == nil {
					out.IsRead = new(bool)
				}
				*out.IsRead = bool(in.Bool())
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
func easyjson4086215fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		if in.Id == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Id))
		}
	}
	{
		const prefix string = ",\"sender_id\":"
		out.RawString(prefix)
		if in.SenderId == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.SenderId))
		}
	}
	{
		const prefix string = ",\"recipient_id\":"
		out.RawString(prefix)
		if in.RecipientId == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.RecipientId))
		}
	}
	{
		const prefix string = ",\"dialog_id\":"
		out.RawString(prefix)
		if in.DialogId == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.DialogId))
		}
	}
	{
		const prefix string = ",\"message_text\":"
		out.RawString(prefix)
		if in.Text == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Text))
		}
	}
	{
		const prefix string = ",\"is_read\":"
		out.RawString(prefix)
		if in.IsRead == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.IsRead))
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
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(l, v)
}
