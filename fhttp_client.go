package http

import (
	http "github.com/vimbing/fhttp"

	"golang.org/x/net/proxy"
)

func newFhttpClient(cfg *Config) (*http.Client, error) {
	client := &http.Client{
		Timeout: cfg.timeout,
	}

	if !cfg.allowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	var dialer proxy.ContextDialer
	var err error

	if len(cfg.proxies) > 0 {
		dialer, err = newConnectDialer(cfg.proxies[RandomInt(0, len(cfg.proxies))])
	} else {
		dialer = proxy.Direct
	}

	if err != nil {
		return &http.Client{}, err
	}

	client.Transport = newRoundTripper(roundTripperSettings{
		clientHello:        cfg.ja3,
		insecureSkipVerify: cfg.insecureSkipVerify,
		dialer:             dialer,
		http2Settings:      cfg.httpSettings.Settings,
		http2SettingsOrder: cfg.httpSettings.Order,
		disablePush:        cfg.httpSettings.DisablePush,
	})

	client.Timeout = cfg.timeout

	return client, nil
}
