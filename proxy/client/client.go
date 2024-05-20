package client

import (
	"errors"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/no-src/nscache/proxy"
	"github.com/no-src/nscache/proxy/request"
	"github.com/no-src/nscache/proxy/response"
	"github.com/no-src/nsgo/httputil"
)

const (
	// DriverName the unique name of the proxy driver for register
	DriverName = "proxy"
)

type proxyCache struct {
	nscache.NSCacheExt

	baseUrl    string
	serializer encoding.Serializer
	client     httputil.HttpClient
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	insecureSkipVerify, certFile, enableHTTP3, https, err := parseConnection(conn)
	if err != nil {
		return nil, err
	}
	client, err := httputil.NewHttpClient(insecureSkipVerify, certFile, enableHTTP3)
	if err != nil {
		return nil, err
	}
	var baseUrl string
	if https {
		baseUrl = "https://" + conn.Host
	} else {
		baseUrl = "http://" + conn.Host
	}
	c := &proxyCache{
		baseUrl:    baseUrl,
		serializer: encoding.DefaultSerializer,
		client:     client,
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *proxyCache) Get(k string, v any) error {
	resp, err := c.client.HttpGet(c.url(k))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var getResp response.Response
	err = c.serializer.Deserialize(data, &getResp)
	if err != nil {
		return err
	}
	if getResp.Code == proxy.StatusNilError {
		return nscache.ErrNil
	} else if getResp.Code != proxy.StatusSuccess {
		return errors.New(getResp.Message)
	}
	return c.serializer.Deserialize([]byte(getResp.Data), &v)
}

func (c *proxyCache) Set(k string, v any, expiration time.Duration) error {
	req := request.SetRequest{
		Value:      v,
		Expiration: expiration,
	}
	data, err := c.serializer.Serialize(req)
	if err != nil {
		return err
	}
	resp, err := c.client.HttpPut(c.url(k), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var setResp response.Response
	err = c.serializer.Deserialize(data, &setResp)
	if err != nil {
		return err
	}
	if setResp.Code != proxy.StatusSuccess {
		return errors.New(setResp.Message)
	}
	return nil
}

func (c *proxyCache) Remove(k string) error {
	resp, err := c.client.HttpDelete(c.url(k), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var setResp response.Response
	err = c.serializer.Deserialize(data, &setResp)
	if err != nil {
		return err
	}
	if setResp.Code != proxy.StatusSuccess {
		return errors.New(setResp.Message)
	}
	return nil
}

func (c *proxyCache) Close() error {
	return nil
}

func (c *proxyCache) url(key string) string {
	return c.baseUrl + "/" + key
}

func parseConnection(u *url.URL) (insecureSkipVerify bool, certFile string, enableHTTP3 bool, https bool, err error) {
	insecureSkipVerify = true
	if u == nil {
		err = errors.New("invalid proxy connection string")
		return
	}
	insecureSkipVerifyStr := u.Query().Get("tls_insecure_skip_verify")
	if len(insecureSkipVerifyStr) > 0 {
		insecureSkipVerify, err = strconv.ParseBool(insecureSkipVerifyStr)
		if err != nil {
			err = errors.New("invalid tls_insecure_skip_verify parameter in the proxy connection string")
			return
		}
	}

	certFile = u.Query().Get("tls_cert_file")

	enableHTTP3Str := u.Query().Get("http3")
	if len(enableHTTP3Str) > 0 {
		enableHTTP3, err = strconv.ParseBool(enableHTTP3Str)
		if err != nil {
			err = errors.New("invalid http3 parameter in the proxy connection string")
			return
		}
	}

	httpsStr := u.Query().Get("https")
	if len(httpsStr) > 0 {
		https, err = strconv.ParseBool(httpsStr)
		if err != nil {
			err = errors.New("invalid https parameter in the proxy connection string")
			return
		}
	}

	if enableHTTP3 {
		https = true
	}
	return
}

func init() {
	nscache.Register(DriverName, newCache)
}
