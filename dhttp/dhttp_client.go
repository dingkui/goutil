package dhttp

import (
	"gitee.com/dk83/goutils/djson"
	"io"
)

type Client struct {
	method  string
	handler *HttpHandler
}
type Clients struct {
	Get     *Client
	Post    *Client
	Put     *Client
	Head    *Client
	Delete  *Client
	handler *HttpHandler
}

func NewClients(url iUrlHandler, req iReqHandler, res iResHandler) *Clients {
	c := &HttpHandler{}
	if url != nil {
		c.url = url
	}
	if req != nil {
		c.req = req
	}
	if res != nil {
		c.res = res
	}
	return &Clients{
		Get:     &Client{"GET", c},
		Post:    &Client{"POST", c},
		Put:     &Client{"PUT", c},
		Head:    &Client{"HEAD", c},
		Delete:  &Client{"DELETE", c},
		handler: c,
	}
}
func NewDefaultClient(base string) *Clients {
	return NewClients(&DefaultUrlHandler{base}, &DefaultReqHandler{}, &DefaultResHandler{})
}

func mkOptionWithHandler(url string, option []*Options, handler *HttpHandler) *Options {
	options := &Options{}
	if len(option) > 0 {
		options = option[0]
	}
	options.Path(url)
	if options.GetArg(httpHandler) == nil {
		options.HttpHandler(handler)
	}
	return options
}
func (cli *Client) Send(url string, option ...*Options) (*Response, error) {
	return HTTPMethod(cli.method).Do(mkOptionWithHandler(url, option, cli.handler))
}
func (cli *Client) SendData(url string, data io.Reader, option ...*Options) (*Response, error) {
	return HTTPMethod(cli.method).Do(mkOptionWithHandler(url, option, cli.handler).DataStream(data))
}
func (cli *Client) SendForm(url string, data *djson.JsonGo, option ...*Options) (*Response, error) {
	return HTTPMethod(cli.method).Do(mkOptionWithHandler(url, option, cli.handler).DataFrom(data))
}
func (cli *Client) SendJson(url string, data *djson.JsonGo, option ...*Options) (*Response, error) {
	return HTTPMethod(cli.method).Do(mkOptionWithHandler(url, option, cli.handler).DataJson(data))
}

func (cli *Clients) Send(method string, url string, option ...*Options) (*Response, error) {
	return (&Client{method, cli.handler}).Send(url, option...)
}
func (cli *Clients) SendData(method string, url string, data io.Reader, option ...*Options) (*Response, error) {
	return (&Client{method, cli.handler}).SendData(url, data, option...)
}
func (cli *Clients) SendForm(method string, url string, data *djson.JsonGo, option ...*Options) (*Response, error) {
	return (&Client{method, cli.handler}).SendForm(url, data, option...)
}
func (cli *Clients) SendJson(method string, url string, data *djson.JsonGo, option ...*Options) (*Response, error) {
	return (&Client{method, cli.handler}).SendJson(url, data, option...)
}
