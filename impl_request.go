/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lark

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/chyroc/lark/internal"
)

// Response contains the general return value of the http request, such as request-id, etc.
type Response struct {
	Method        string      // request method
	URL           string      // request url
	LogID         string      // log id, if you got some error and oncall lark/feishu team, please with this log id
	StatusCode    int         // http response status code
	Header        http.Header // http response header
	ContentLength int64       // http response content length

	// deprecated
	RequestID string
}

// RawRequest Send a http request of lark
//
// for example, you can see BotService.GetBotInfo
func (r *Lark) RawRequest(ctx context.Context, req *RawRequestReq, resp interface{}) (response *Response, err error) {
	if r.mock.mockRawRequest != nil {
		return r.mock.mockRawRequest(ctx, req, resp)
	}

	return r.wrapDoRequest(ctx, req, resp)
}

// MockRawRequest mock request
func (r *Mock) MockRawRequest(f func(ctx context.Context, req *RawRequestReq, resp interface{}) (response *Response, err error)) {
	r.mockRawRequest = f
}

// UnMockRawRequest un mock request
func (r *Mock) UnMockRawRequest() {
	r.mockRawRequest = nil
}

func (r *Lark) rawRequest(ctx context.Context, req *RawRequestReq, resp interface{}) (response *Response, err error) {
	// 1. parse request
	rawHttpReq, err := r.parseRawHttpRequest(ctx, req)
	if err != nil {
		// 这里日志不需要区分 level, 输出 [error] 日志
		r.Log(ctx, LogLevelError, "[lark] %s#%s %s %s parse_req failed, err=%s", req.Scope, req.API, req.Method, req.URL, err)
		return response, err
	}

	// 2. request log
	logLevel := r.getLogLevel(req.Scope, req.API)
	switch logLevel {
	case LogLevelDebug:
		r.Log(ctx, LogLevelDebug, "[lark] %s#%s %s %s start, body=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, string(rawHttpReq.RawBody))
	case LogLevelInfo:
		r.Log(ctx, LogLevelInfo, "[lark] %s#%s %s %s start", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL)
	default:
		// error 不需要 req 日志, 合并到 resp 一起
	}

	// 3. do request
	var respContent string
	response, respContent, err = r.doRequest(ctx, rawHttpReq, resp)
	if err != nil {
		switch logLevel {
		case LogLevelDebug:
			// [debug]: 详细 error 日志
			r.Log(ctx, LogLevelError, "[lark] %s#%s %s %s failed, log_id=%s, status=%d, body=%s, err=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, response.LogID, response.StatusCode, respContent, err)
		default:
			// [其他]: 简单 error 日志
			r.Log(ctx, LogLevelError, "[lark] %s#%s %s %s failed, log_id=%s, status=%d, err=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, response.LogID, response.StatusCode, err)
		}
		return response, err
	}
	code, msg, detailErr := getCodeMsg(resp)

	// 4. response log
	if response.StatusCode >= http.StatusBadRequest || code != 0 {
		r.Log(ctx, LogLevelError, "[lark] %s#%s %s %s failed, log_id=%s, status=%d, code=%d, msg=%s, detail=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, response.LogID, response.StatusCode, code, msg, detailErr)
	} else {
		switch logLevel {
		case LogLevelDebug:
			r.Log(ctx, LogLevelDebug, "[lark] %s#%s %s %s success, log_id=%s, body=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, response.LogID, respContent)
		case LogLevelInfo:
			r.Log(ctx, LogLevelInfo, "[lark] %s#%s %s %s success, log_id=%s", rawHttpReq.Scope, rawHttpReq.API, rawHttpReq.Method, rawHttpReq.URL, response.LogID)
		default:
			// error 不需要 resp 日志
		}
	}

	// 5. response
	if code != 0 {
		e := NewError(req.Scope, req.API, code, msg)
		e.(*Error).ErrorDetail = detailErr
		return response, e
	}

	return response, nil
}

// RawRequestReq is the parameter that composes the lark sdk request
type RawRequestReq struct {
	Scope                 string        // api domain, such as Auth, Chat, Mail
	API                   string        // api name, as in CreateMailGroup
	Method                string        // http request method, such as GET, POST
	URL                   string        // http request url
	Body                  interface{}   // http request body, query, path and other parameter information
	IsFile                bool          // send body data as a file
	NeedTenantAccessToken bool          // do you need TenantAccessToken
	NeedAppAccessToken    bool          // do you need AppAccessToken
	NeedUserAccessToken   bool          // do you need UserAccessToken
	NeedHelpdeskAuth      bool          // do you need HelpdeskAccessToken
	MethodOption          *MethodOption // method option, such as userAccessToken
}

// 把可读的 RawRequestReq ，解析为 http 请求的参数 rawHttpRequestParam
func (r *Lark) parseRawHttpRequest(ctx context.Context, req *RawRequestReq) (*rawHttpRequest, error) {
	if req.MethodOption == nil {
		req.MethodOption = new(MethodOption)
	}

	// 0 init
	rawHttpReq := &rawHttpRequest{
		Scope:   req.Scope,
		API:     req.API,
		Method:  strings.ToUpper(req.Method),
		Headers: map[string]string{},
		URL:     req.URL,
		Timeout: r.timeout,
	}
	if req.MethodOption.timeout > 0 {
		rawHttpReq.Timeout = req.MethodOption.timeout
	}
	if len(req.MethodOption.headers) > 0 {
		for k, v := range req.MethodOption.headers {
			rawHttpReq.Headers[k] = v
		}
	}

	// 1 headers
	if err := rawHttpReq.parseHeader(ctx, r, req); err != nil {
		return nil, err
	}

	// 2 body
	if err := rawHttpReq.parseRawRequestReqBody(req.Body, req.IsFile); err != nil {
		return nil, err
	}

	// 3 return
	return rawHttpReq, nil
}

func (r *Lark) doRequest(ctx context.Context, rawHttpReq *rawHttpRequest, realResponse interface{}) (*Response, string, error) {
	response := new(Response)
	response.Method = rawHttpReq.Method
	response.URL = rawHttpReq.URL
	response.Header = map[string][]string{}

	if rawHttpReq.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, rawHttpReq.Timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, rawHttpReq.Method, rawHttpReq.URL, rawHttpReq.Body)
	if err != nil {
		return response, "", err
	}
	for k, v := range rawHttpReq.Headers {
		req.Header.Set(k, v)
	}

	resp, err := r.httpClient.Do(ctx, req)
	if err != nil {
		return response, "", err
	}

	_, media, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	respFilename := ""
	if media != nil {
		respFilename = media["filename"]
	}

	response.StatusCode = resp.StatusCode
	response.RequestID = resp.Header.Get("X-Request-Id")
	response.LogID = resp.Header.Get("X-Tt-Logid")
	response.Header = resp.Header
	response.ContentLength = resp.ContentLength

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, "", err
	}

	var respContent string
	if respFilename == "" {
		respContent = string(bs)
	} else {
		respContent = fmt.Sprintf("<FILE: %d>", len(bs))
	}

	if realResponse != nil {
		if resp != nil && resp.StatusCode == http.StatusOK {
			isSpecResp := false
			if setter, ok := realResponse.(readerSetter); ok {
				isSpecResp = true
				setter.SetReader(bytes.NewReader(bs))
			}
			if setter, ok := realResponse.(filenameSetter); ok {
				isSpecResp = true
				setter.SetFilename(respFilename)
			}
			if isSpecResp {
				return response, respContent, nil
			}
		}

		if len(bs) == 0 && resp.StatusCode >= http.StatusBadRequest {
			return response, respContent, fmt.Errorf("request fail: %s", resp.Status)
		}

		if err = json.Unmarshal(bs, realResponse); err != nil {
			return response, respContent, fmt.Errorf("invalid json: %s, err: %s", bs, err)
		}
	}

	return response, respContent, nil
}

func (r *rawHttpRequest) parseHeader(ctx context.Context, ins *Lark, req *RawRequestReq) error {
	if ins.isEnableLogID {
		logID, ok := getStringFromContext(ctx, rpcLogIDKey)
		if ok {
			r.Headers[httpHeaderLogIDKey] = logID
		}
	}
	r.Headers["User-Agent"] = fmt.Sprintf("chyroc-go-lark/%s (https://github.com/chyroc/lark)", version)

	if req.NeedUserAccessToken && req.MethodOption.userAccessToken != "" {
		r.Headers[httpRequestHeaderAuthorization] = "Bearer " + req.MethodOption.userAccessToken
	} else if req.NeedTenantAccessToken {
		token, _, err := ins.Auth.GetTenantAccessToken(ctx)
		if err != nil {
			return err
		}
		r.Headers[httpRequestHeaderAuthorization] = "Bearer " + token.Token
	} else if req.NeedAppAccessToken {
		token, _, err := ins.Auth.GetAppAccessToken(ctx)
		if err != nil {
			return err
		}
		r.Headers[httpRequestHeaderAuthorization] = "Bearer " + token.Token
	}

	if req.NeedHelpdeskAuth {
		r.Headers[httpRequestHeaderHelpdeskAuth] = base64.StdEncoding.EncodeToString([]byte(ins.helpdeskID + ":" + ins.helpdeskToken))
	}

	return nil
}

func (r *rawHttpRequest) parseRawRequestReqBody(body interface{}, isFile bool) error {
	var reader io.Reader
	fileKey := ""
	query := url.Values{}
	isNeedBody := false
	fileData := map[string]string{}

	if err := rangeStruct(body, func(fieldVV reflect.Value, fieldVT reflect.StructField) error {
		if path := fieldVT.Tag.Get("path"); path != "" {
			if strings.Contains(r.URL, ":"+path) {
				r.URL = strings.ReplaceAll(r.URL, ":"+path, internal.ReflectToString(fieldVV))
			} else {
				r.URL = strings.ReplaceAll(r.URL, "{"+path+"}", internal.ReflectToString(fieldVV))
			}
		} else if queryKey := fieldVT.Tag.Get("query"); queryKey != "" {
			value := internal.ReflectToQueryString(fieldVV)
			sep := fieldVT.Tag.Get("join_sep")
			if sep != "" {
				query.Add(queryKey, strings.Join(value, sep))
			} else {
				for _, v := range value {
					query.Add(queryKey, v)
				}
			}
		} else if header := fieldVT.Tag.Get("header"); header != "" {
			switch header {
			case "range":
				if fieldVV.Kind() != reflect.Array || fieldVV.Len() != 2 || fieldVV.Index(0).Kind() != reflect.Int64 {
					return fmt.Errorf("with range header, value must be [2]int64")
				}
				from := fieldVV.Index(0).Int()
				to := fieldVV.Index(1).Int()
				if from != 0 || to != 0 {
					r.Headers["Range"] = "bytes=" + strconv.FormatInt(from, 10) + "-" + strconv.FormatInt(to, 10)
				}
			}
		} else if j := fieldVT.Tag.Get("json"); j != "" {
			if strings.HasSuffix(j, ",omitempty") {
				j = j[:len(j)-10]
			}
			if isFile {
				fileKey = j
				if r, ok := fieldVV.Interface().(io.Reader); ok {
					reader = r
				} else {
					fileData[j] = internal.ReflectToString(fieldVV)
				}
			} else {
				isNeedBody = true
			}
		}
		return nil
	}); err != nil {
		return err
	}

	if isNeedBody {
		bs, err := json.Marshal(body)
		if err != nil {
			return err
		}
		r.Body = bytes.NewBuffer(bs)
		r.RawBody = bs
		r.Headers["Content-Type"] = "application/json; charset=utf-8"
	}

	if isFile {
		contentType, bod, err := newFileUploadRequest(fileData, fileKey, reader)
		if err != nil {
			return err
		}
		r.Headers["Content-Type"] = contentType
		r.Body = bod
		r.RawBody = []byte("<FILE>")
	}

	if len(query) > 0 {
		r.URL = r.URL + "?" + query.Encode()
	}

	return nil
}

type rawHttpRequest struct {
	Scope   string
	API     string
	Method  string
	URL     string
	Body    io.Reader
	RawBody []byte
	Headers map[string]string
	Timeout time.Duration
}

func newFileUploadRequest(params map[string]string, filekey string, reader io.Reader) (string, io.Reader, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filekey, "unknown-file")
	if err != nil {
		return "", nil, err
	}
	if reader != nil {
		if _, err = io.Copy(part, reader); err != nil {
			return "", nil, err
		}
	}
	for key, val := range params {
		if err = writer.WriteField(key, val); err != nil {
			return "", nil, err
		}
	}
	if err = writer.Close(); err != nil {
		return "", nil, err
	}

	return writer.FormDataContentType(), body, nil
}

type readerSetter interface {
	SetReader(file io.Reader)
}

type filenameSetter interface {
	SetFilename(filename string)
}

func getCodeMsg(v interface{}) (code int64, msg string, detail *ErrorDetail) {
	if v == nil {
		return 0, "", nil
	}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	if vv.Kind() != reflect.Struct {
		return 0, "", nil
	}
	codeField := vv.FieldByName("Code")
	if codeField.IsValid() {
		if internal.IsInReflectKind(codeField.Kind(), []reflect.Kind{
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
		}) {
			code = int64(codeField.Int())
		} else if internal.IsInReflectKind(codeField.Kind(), []reflect.Kind{
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
		}) {
			code = int64(codeField.Uint())
		}
	}

	msgField := vv.FieldByName("Msg")
	if msgField.IsValid() {
		if msgField.Kind() == reflect.String {
			msg = msgField.String()
		}
	}

	errorField := vv.FieldByName("Error")
	if errorField.IsValid() {
		detail, _ = errorField.Interface().(*ErrorDetail)
	}
	return
}

func getResponseLogID(response *Response) (logID string, statusCode int) {
	if response == nil {
		return
	}
	logID = response.LogID
	statusCode = response.StatusCode
	return
}

func getStringFromContext(ctx context.Context, key string) (string, bool) {
	if ctx == nil {
		return "", false
	}

	v := ctx.Value(key)
	if v == nil {
		return "", false
	}

	switch v := v.(type) {
	case string:
		return v, true
	case *string:
		if v == nil {
			return "", false
		}
		return *v, true
	}
	return "", false
}

func rangeStruct(v interface{}, f func(fieldVV reflect.Value, fieldVT reflect.StructField) error) error {
	vv := reflect.ValueOf(v)
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}

	for i := 0; i < vt.NumField(); i++ {
		fieldVV := vv.Field(i)
		fieldVT := vt.Field(i)

		if fieldVV.Kind() == reflect.Ptr && fieldVV.IsNil() {
			continue
		}
		if fieldVV.Kind() == reflect.Slice && fieldVV.Len() == 0 {
			continue
		}

		err := f(fieldVV, fieldVT)
		if err != nil {
			return err
		}
	}

	return nil
}

type defaultHttpClient struct {
	ins *http.Client
}

func newDefaultHttpClient(cli *http.Client) HttpClient {
	return &defaultHttpClient{
		ins: cli,
	}
}

// Do ...
func (r *defaultHttpClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return r.ins.Do(req)
}

func FormatRequestBody(body interface{}) interface{} {
	if body == nil {
		return nil
	}
	vt := reflect.TypeOf(body)
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() != reflect.Struct {
		return body
	}

	queries := url.Values{}
	res := map[string]interface{}{}

	_ = rangeStruct(body, func(fieldVV reflect.Value, fieldVT reflect.StructField) error {
		key, val := formatRequestBodyKey(fieldVV, fieldVT, &queries)
		if key != "" {
			res[key] = val
		}
		return nil
	})

	for k, v := range queries {
		if len(v) == 1 {
			res[k] = v[0]
		} else if len(v) > 1 {
			res[k] = v
		}
	}
	return res
}

func formatRequestBodyKey(fieldVV reflect.Value, fieldVT reflect.StructField, queries *url.Values) (string, interface{}) {
	if path := fieldVT.Tag.Get("path"); path != "" {
		return path, internal.ReflectToString(fieldVV)
	} else if queryKey := fieldVT.Tag.Get("query"); queryKey != "" {
		value := internal.ReflectToQueryString(fieldVV)
		sep := fieldVT.Tag.Get("join_sep")
		if sep != "" {
			queries.Add(queryKey, strings.Join(value, sep))
		} else {
			for _, v := range value {
				queries.Add(queryKey, v)
			}
		}
		return "", nil
	} else if header := fieldVT.Tag.Get("header"); header != "" {
		switch header {
		case "range":
			if fieldVV.Kind() != reflect.Array || fieldVV.Len() != 2 || fieldVV.Index(0).Kind() != reflect.Int64 {
				return "", nil
			}
			from := fieldVV.Index(0).Int()
			to := fieldVV.Index(1).Int()
			if from != 0 || to != 0 {
				return "range", fmt.Sprintf("bytes=%d-%d", from, to)
			}
		}
	} else if j := fieldVT.Tag.Get("json"); j != "" {
		if strings.HasSuffix(j, ",omitempty") {
			j = j[:len(j)-10]
		}
		if _, ok := fieldVV.Interface().(io.Reader); ok {
			return j, "<FILE>"
		}
		for fieldVV.Kind() == reflect.Ptr {
			fieldVV = fieldVV.Elem()
		}
		return j, FormatRequestBody(fieldVV.Interface())
	}
	return "", nil
}
