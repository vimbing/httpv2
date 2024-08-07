package http

import (
	"io"
	"time"

	fhttp "github.com/vimbing/fhttp"
	"github.com/vimbing/fhttp/cookiejar"
	"github.com/vimbing/retry"
	tls "github.com/vimbing/vutls"
)

type OptionStringJa string
type OptionTimeout time.Duration
type OptionProxy string
type OptionDisallowRedirect bool
type OptionUtlsJa3HelloId tls.ClientHelloID
type OptionUtlsJa3HelloSpec tls.ClientHelloSpec
type OptionTlsProfile TlsProfile
type OptionInsecureSkipVerify bool
type OptionCookieJar *cookiejar.Jar
type OptionRequestMiddleware []RequestMiddlewareFunc
type OptionResponseMiddleware []ResponseMiddlewareFunc
type OptionResponseErrorMiddleware []ResponseErrorMiddlewareFunc

type Client struct {
	fhttpClient *fhttp.Client
	cfg         *Config
}

type RequestMiddlewareFunc func(*Request) error
type ResponseMiddlewareFunc func(*Response) error
type ResponseErrorMiddlewareFunc func(*Request, error)

type Config struct {
	insecureSkipVerify      bool
	requestMiddleware       []RequestMiddlewareFunc
	responseMiddleware      []ResponseMiddlewareFunc
	responseErrorMiddleware []ResponseErrorMiddlewareFunc
	proxies                 []string
	allowRedirect           bool
	timeout                 time.Duration
	ja3                     tls.ClientHelloID
	tlsProfile              *TlsProfile
	jar                     *cookiejar.Jar
}

type RequestJsonBody any
type QueryParams map[string]string
type FormUrlEncoded map[string]string

type Request struct {
	Method string
	Body   io.Reader
	Header fhttp.Header
	Url    string

	host         *string
	retrier      *retry.Retrier
	tlsProfile   *TlsProfile
	fhttpRequest *fhttp.Request
}

type Response struct {
	Body          []byte
	fhttpResponse *fhttp.Response
}

type TlsProfile struct {
	SecChUa          string
	SecChUaMobile    string
	SecChaUaPlatform string
	UserAgent        string
}
