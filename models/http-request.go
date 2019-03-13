package models

import (
	"io"
)

type HttpRequest struct {
	Headers     map[string]string
	QueryParams map[string]string
	Path        string
	Body        io.Reader
}
