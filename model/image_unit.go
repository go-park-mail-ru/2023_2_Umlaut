package model

import "io"

type ImageUnit struct {
	Name        string
	FileLoad    io.Reader
	Size        int64
	ContentType string
}
