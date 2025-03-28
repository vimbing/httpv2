package http

import (
	"io"
	"time"

	fhttp "github.com/vimbing/fhttp"
	"github.com/vimbing/fhttp/cookiejar"
	"github.com/vimbing/fhttp/http2"
	tls "github.com/vimbing/vutls"
)

type OptionStringJa string
type OptionTimeout time.Duration
type OptionProxy string
type OptionDisallowRedirect bool
type OptionForcedProxyRotation bool
type OptionUtlsJa3HelloId tls.ClientHelloID
type OptionUtlsJa3HelloSpec tls.ClientHelloSpec
type OptionTlsProfile TlsProfile
type OptionInsecureSkipVerify bool
type OptionCookieJar *cookiejar.Jar
type OptionRequestMiddleware []RequestMiddlewareFunc
type OptionResponseMiddleware []ResponseMiddlewareFunc
type OptionResponseErrorMiddleware []ResponseErrorMiddlewareFunc
type OptionHttpSettings Http2Settings
type OptionRetry *Retry
type OptionStatusValidationFunc StatusValidationFunc

type Client struct {
	fhttpClient *fhttp.Client
	cfg         *Config
}

type RequestMiddlewareFunc func(*Request) error
type ResponseMiddlewareFunc func(*Response) error
type ResponseErrorMiddlewareFunc func(*Request, error)

type Http2Settings struct {
	Order       []http2.SettingID
	Settings    map[http2.SettingID]uint32
	DisablePush bool
}

type Config struct {
	insecureSkipVerify      bool
	requestMiddleware       []RequestMiddlewareFunc
	responseMiddleware      []ResponseMiddlewareFunc
	responseErrorMiddleware []ResponseErrorMiddlewareFunc
	proxies                 []string
	forceRotation           bool
	allowRedirect           bool
	timeout                 time.Duration
	ja3                     tls.ClientHelloID
	tlsProfile              *TlsProfile
	jar                     *cookiejar.Jar
	httpSettings            Http2Settings
	retry                   *Retry
	statusValidationFunc    StatusValidationFunc
}

type RequestJsonBody any
type QueryParams map[string]string
type FormUrlEncoded map[string]string

type Request struct {
	Method string
	Body   io.Reader
	Header fhttp.Header
	Url    string

	protoMinor int
	protoMajor int
	proto      string

	host         *string
	tlsProfile   *TlsProfile
	fhttpRequest *fhttp.Request
}

type Response struct {
	Body          []byte
	fhttpResponse *fhttp.Response
}

type TlsProfile struct {
	Http2Settings Http2Settings
	Ja3           tls.ClientHelloID
}

type requestExecutionResult struct {
	res   *Response
	error error
}

type Retry struct {
	Max           int
	Delay         time.Duration
	IgnoredErrors []error
	EndingErrors  []error
	OnError       func(error)
}

type doFunc func(*Request) (*Response, error)

type StatusValidationFunc func(status int, client *Client) error
