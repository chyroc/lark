package lark

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/chyroc/lark/internal"
)

type Response struct {
	// HTTPResponse *http.Response

	Method     string
	URL        string
	RequestID  string
	StatusCode int
}

func (r *Lark) RawRequest(ctx context.Context, req *RawRequestReq, resp interface{}) (*Response, error) {
	r.log(ctx, LogLevelInfo, "[lark] %s#%s call api", req.Scope, req.API)
	r.log(ctx, LogLevelDebug, "[lark] %s#%s request: %s", req.Scope, req.API, jsonString(req))

	headers, err := r.prepareHeaders(ctx, req)
	if err != nil {
		return nil, err
	}
	response, err := r.doRequest(ctx, r.httpClient, req, headers, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.log(ctx, LogLevelError, "[lark] %s#%s %s %s failed, request_id: %s, status_code: %d, error: %s", req.Scope, req.API, req.Method, req.URL, requestID, statusCode, err)
		return response, err
	}
	code, msg := getCodeMsg(resp)
	if code != 0 {
		r.log(ctx, LogLevelError, "[lark] %s#%s %s %s failed, request_id: %s, status_code: %d, code: %d, msg: %s", req.Scope, req.API, req.Method, req.URL, requestID, statusCode, code, msg)
		return response, NewError(req.Scope, req.API, code, msg)
	}

	r.log(ctx, LogLevelDebug, "[lark] %s#%s success, request_id: %s, status_code: %d, response: %s", req.Scope, req.API, requestID, statusCode, "TODO")

	return response, nil
}

type RawRequestReq struct {
	Scope                 string
	API                   string
	Method                string
	URL                   string
	Body                  interface{}
	IsFile                bool
	NeedTenantAccessToken bool
	NeedAppAccessToken    bool
	NeedUserAccessToken   bool
	NeedHelpdeskAuth      bool
	MethodOption          *MethodOption
}

func (r *Lark) prepareHeaders(ctx context.Context, req *RawRequestReq) (map[string]string, error) {
	headers := map[string]string{}
	if req.Method != http.MethodGet {
		headers["Content-Type"] = "application/json; charset=utf-8"
	}
	if req.NeedUserAccessToken && req.MethodOption.userAccessToken != "" {
		headers["Authorization"] = "Bearer " + req.MethodOption.userAccessToken
	} else if req.NeedTenantAccessToken {
		token, _, err := r.Auth.GetTenantAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		headers["Authorization"] = "Bearer " + token.Token
	} else if req.NeedAppAccessToken {
		token, _, err := r.Auth.GetAppAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		headers["Authorization"] = "Bearer " + token.Token
	}

	if req.NeedHelpdeskAuth {
		headers["X-Lark-Helpdesk-Authorization"] = base64.StdEncoding.EncodeToString([]byte(r.helpdeskID + ":" + r.helpdeskToken))
	}

	return headers, nil
}

func parseRequestParam(req *RawRequestReq) (*realRequestParam, error) {
	uri := req.URL
	var body io.Reader
	var reader io.Reader
	headers := map[string]string{}
	fileKey := ""

	vv := reflect.ValueOf(req.Body)
	vt := reflect.TypeOf(req.Body)

	if vt.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}

	q := url.Values{}
	isNeedQuery := false
	isNeedBody := false
	filedata := map[string]string{}

	for i := 0; i < vt.NumField(); i++ {
		fieldVV := vv.Field(i)
		fieldVT := vt.Field(i)

		if fieldVV.Kind() == reflect.Ptr && fieldVV.IsNil() {
			continue
		}
		if fieldVV.Kind() == reflect.Slice && fieldVV.Len() == 0 {
			continue
		}

		if path := fieldVT.Tag.Get("path"); path != "" {
			if strings.Contains(uri, ":"+path) {
				uri = strings.ReplaceAll(uri, ":"+path, internal.ReflectToString(fieldVV))
			} else {
				uri = strings.ReplaceAll(uri, "{"+path+"}", internal.ReflectToString(fieldVV))
			}
			continue
		} else if query := fieldVT.Tag.Get("query"); query != "" {
			isNeedQuery = true
			for _, v := range internal.ReflectToQueryString(fieldVV) {
				q.Add(query, v)
			}
			continue
		} else if j := fieldVT.Tag.Get("json"); j != "" {
			if strings.HasSuffix(j, ",omitempty") {
				j = j[:len(j)-10]
			}
			if req.IsFile {
				fileKey = j
				if r, ok := fieldVV.Interface().(io.Reader); ok {
					reader = r
				} else {
					filedata[j] = internal.ReflectToString(fieldVV)
				}
			} else {
				isNeedBody = true
			}
			continue
		}
	}

	if isNeedBody {
		bs, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(bs)
	}

	if req.IsFile {
		contentType, bod, err := newFileUploadRequest(filedata, fileKey, reader)
		if err != nil {
			return nil, err
		}
		headers["Content-Type"] = contentType
		body = bod
	}

	if isNeedQuery {
		uri = uri + "?" + q.Encode()
	}

	return &realRequestParam{
		Method:  strings.ToUpper(req.Method),
		URL:     uri,
		Body:    body,
		Headers: headers,
	}, nil
}

type realRequestParam struct {
	Method  string
	URL     string
	Body    io.Reader
	Headers map[string]string
}

func (r *Lark) doRequest(ctx context.Context, cli *http.Client, requestParam *RawRequestReq, headers map[string]string, realResponse interface{}) (*Response, error) {
	response := new(Response)
	realReq, err := parseRequestParam(requestParam)
	if err != nil {
		return response, err
	}

	response.Method = realReq.Method
	response.URL = realReq.URL
	for k, v := range realReq.Headers {
		headers[k] = v
	}

	req, err := http.NewRequest(realReq.Method, realReq.URL, realReq.Body)
	if err != nil {
		return response, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := cli.Do(req)
	if err != nil {
		return response, err
	}
	// response.HTTPResponse = resp
	response.StatusCode = resp.StatusCode
	response.RequestID = resp.Header.Get("x-doRequest-id")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	r.log(ctx, LogLevelDebug, "%s %s, got: %s", realReq.Method, realReq.URL, bs)

	if realResponse != nil {
		if resp != nil && resp.StatusCode == http.StatusOK {
			respFileSetter, ok := realResponse.(readerSetter)
			if ok {
				respFileSetter.SetReader(bytes.NewReader(bs))
				return response, nil
			}
		}

		if err = json.Unmarshal(bs, realResponse); err != nil {
			return response, fmt.Errorf("invalid json: %s", bs)
		}
	}

	return response, nil
}

func newFileUploadRequest(params map[string]string, filekey string, reader io.Reader) (string, io.Reader, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filekey, "file.file")
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

func getCodeMsg(v interface{}) (code int64, msg string) {
	if v == nil {
		return 0, ""
	}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	if vv.Kind() != reflect.Struct {
		return 0, ""
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

	codeMsg := vv.FieldByName("Msg")
	if codeField.IsValid() {
		if codeMsg.Kind() == reflect.String {
			msg = codeMsg.String()
		}
	}
	return
}

func getResponseRequestID(response *Response) (requestID string, statusCode int) {
	if response == nil {
		return
	}
	requestID = response.RequestID
	statusCode = response.StatusCode
	return
}
