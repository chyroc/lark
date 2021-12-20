package lark

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/chyroc/lark/internal"
)

// GenerateJssdkSignature ...
func (r *JssdkService) GenerateJssdkSignature(ctx context.Context, request *GenerateJssdkSignatureReq) (string, error) {
	resp, _, err := r.GetJssdkTicket(ctx, &GetJssdkTicketReq{})
	if err != nil {
		return "", err
	}
	noncestr := internal.RandString(16)
	timestamp := time.Now().UnixNano() / 1000 / 1000
	s := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", resp.Ticket, noncestr, timestamp, request.URL)

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(s))

	return fmt.Sprintf("%x", sha1Hash.Sum(nil)), nil
}

// GenerateJssdkSignatureReq ...
type GenerateJssdkSignatureReq struct {
	URL string `json:"url"` // 调用飞书组件的页面的url，不要包括#、?后面的参数。 比如原url是https://m.mm.cn/ttc/3541093/3131_1.html#subtitle?foo=bar，则实际应该传入https://m.mm.cn/ttc/3541093/3131_1.html
}
