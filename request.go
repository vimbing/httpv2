package http

import (
	"context"
	"time"

	fhttp "github.com/vimbing/fhttp"
)

func (r *Request) useTlsProfile() {
	if r.tlsProfile == nil || r.tlsProfile.Headers == nil {
		return
	}

	for k, v := range r.tlsProfile.Headers {
		r.Header[k] = v
	}

	if len(r.tlsProfile.HeaderOrder) > 0 {
		r.Header[fhttp.HeaderOrderKey] = r.tlsProfile.HeaderOrder
	}

	if len(r.tlsProfile.PseudoHeaderOrder) > 0 {
		r.Header[fhttp.PHeaderOrderKey] = r.tlsProfile.PseudoHeaderOrder
	}
}

func (r *Request) Build(timeout time.Duration) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)

	r.useTlsProfile()

	req, err := fhttp.NewRequestWithContext(ctx, r.Method, r.Url, r.Body)

	if err != nil {
		return ctx, cancel, err
	}

	req.Header = r.Header
	r.fhttpRequest = req

	if r.host != nil {
		r.fhttpRequest.Host = *r.host
	}

	if len(r.proto) > 0 {
		req.Proto = r.proto
		req.ProtoMajor = r.protoMajor
		req.ProtoMinor = r.protoMinor
	}

	return ctx, cancel, nil
}

func (r *Request) SetHost(host string) {
	r.host = &host
}

func (r *Request) SetProto(proto string, major int, minor int) {
	r.proto = proto
	r.protoMajor = major
	r.protoMinor = minor
}
