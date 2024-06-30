package http

import (
	"encoding/json"
	"fmt"
)

func (r *Response) StatusCode() int {
	return r.fhttpResponse.StatusCode
}

func (r *Response) Status() string {
	return r.fhttpResponse.Status
}

func (r *Response) BodyDecode(out any) error {
	return json.Unmarshal(r.Body, out)
}

func (r *Response) BodyString() string {
	return string(r.Body)
}

func (r *Response) BodyStringJsonIndented(out any) string {
	err := r.BodyDecode(out)

	if err != nil {
		return ""
	}

	indented, err := json.MarshalIndent(out, "", "	")

	if err != nil {
		return ""
	}

	return string(indented)
}

func (r *Response) BodyBytes() []byte {
	return r.Body
}

// just for logging purposes
func (r *Response) StatusError() error {
	return fmt.Errorf("server responded with unexpected status code: %d", r.fhttpResponse.StatusCode)
}
