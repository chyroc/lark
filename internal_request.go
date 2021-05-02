package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/chyroc/lark/internal"
)

type Response struct {
	HTTPResponse *http.Response

	RequestID string
}

func (r *apiImpl) request(ctx context.Context, req *requestParam, resp interface{}) (*Response, error) {
	headers := map[string]string{}
	if req.NeedTenantAccessToken {
		token, _, err := r.Token().GetTenantAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		headers["Authorization"] = "Bearer " + token.Token
	}
	if req.NeedAppAccessToken {
		token, _, err := r.Token().GetAppAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		headers["Authorization"] = "Bearer " + token.Token
	}
	return request(ctx, r.httpClient, req, headers, resp)
}

type requestParam struct {
	Method                string
	URL                   string
	Body                  interface{}
	NeedTenantAccessToken bool
	NeedAppAccessToken    bool
}

func parseRequestParam(req *requestParam) (*realRequestParam, error) {
	vv := reflect.ValueOf(req.Body)
	vt := reflect.TypeOf(req.Body)

	if vt.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}

	q := url.Values{}
	uri := req.URL
	isNeedQuery := false
	isNeedBody := false

	for i := 0; i < vt.NumField(); i++ {
		fieldVV := vv.Field(i)
		fieldVT := vt.Field(i)
		if path := fieldVT.Tag.Get("path"); path != "" {
			uri = strings.ReplaceAll(uri, ":"+path, internal.ReflectToString(fieldVV))
			continue
		} else if query := fieldVT.Tag.Get("query"); query != "" {
			isNeedQuery = true
			q.Set(query, internal.ReflectToString(fieldVV))
			continue
		} else if j := fieldVT.Tag.Get("json"); j != "" {
			isNeedBody = true
			continue
		}
		spew.Dump(req)
		os.Exit(0)
	}

	var body io.Reader
	if isNeedBody {
		bs, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(bs)
	}

	if isNeedQuery {
		uri = uri + "?" + q.Encode()
	}

	return &realRequestParam{
		Method: strings.ToUpper(req.Method),
		URL:    uri,
		Body:   body,
	}, nil
}

type realRequestParam struct {
	Method string
	URL    string
	Body   io.Reader
}

func request(ctx context.Context, cli *http.Client, requestParam *requestParam, headers map[string]string, realResponse interface{}) (*Response, error) {
	response := new(Response)
	realReq, err := parseRequestParam(requestParam)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(realReq.Method, realReq.URL, realReq.Body)
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := cli.Do(req)
	if err != nil {
		return response, err
	}
	response.HTTPResponse = resp
	response.RequestID = resp.Header.Get("x-request-id")

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(bs, realResponse)
	if err != nil {
		return response, fmt.Errorf("invalid json: %s", bs)
	}
	return response, nil
}
