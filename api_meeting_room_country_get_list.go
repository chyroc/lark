package lark

import (
	"context"
)

// GetCountryList 新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN
func (r *MeetingRoomAPI) GetCountryList(ctx context.Context, request *GetCountryListReq) (*GetCountryListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/country/list",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCountryListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "GetCountryList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCountryListReq struct{}

type getCountryListResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetCountryListResp `json:"data,omitempty"` // 返回业务信息
}

type GetCountryListResp struct {
	Countries *GetCountryListRespCountries `json:"countries,omitempty"` // 国家地区列表
}

type GetCountryListRespCountries struct {
	CountryID string `json:"country_id,omitempty"` // 国家地区ID
	Name      string `json:"name,omitempty"`       // 国家地区名称
}
