package lark

import (
	"context"
)

// UpdateBuilding 该接口用于编辑建筑信息，添加楼层，删除楼层，编辑楼层信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uETNwYjLxUDM24SM1AjN
func (r *MeetingRoomAPI) UpdateBuilding(ctx context.Context, request *UpdateBuildingReq) (*UpdateBuildingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/update",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateBuildingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "UpdateBuilding", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateBuildingReq struct {
	BuildingID       string   `json:"building_id,omitempty"`        // 要更新的建筑ID
	Name             *string  `json:"name,omitempty"`               // 建筑名称
	Floors           []string `json:"floors,omitempty"`             // 楼层列表
	CountryID        *string  `json:"country_id,omitempty"`         // 国家/地区ID
	DistrictID       *string  `json:"district_id,omitempty"`        // 城市ID
	CustomBuildingID *string  `json:"custom_building_id,omitempty"` // 租户自定义建筑ID
}

type updateBuildingResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *UpdateBuildingResp `json:"data,omitempty"`
}

type UpdateBuildingResp struct{}
