package lark

import (
	"context"
)

// QueryRoomConfig
//
// > 查询一个范围内的会议室配置
// 根据查询范围传入对应的参数
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/room_config/query
func (r *VCAPI) QueryRoomConfig(ctx context.Context, request *QueryRoomConfigReq) (*QueryRoomConfigResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/room_configs/query",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(queryRoomConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "QueryRoomConfig", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type QueryRoomConfigReq struct {
	Scope      int    `query:"scope" json:"-"`       // 查询节点范围，必选字段，可用值：【1（租户），2（国家/地区），3（城市），4（建筑），5（楼层），6（会议室）】，示例值：6
	CountryID  string `query:"country_id" json:"-"`  // 国家/地区id scope为2/3时此参数必填，示例值："666"
	DistrictID string `query:"district_id" json:"-"` // 城市id scope为3时此参数必填，示例值："555"
	BuildingID string `query:"building_id" json:"-"` // 建筑id scope为4/5时此参数必填，示例值："123"
	FloorName  string `query:"floor_name" json:"-"`  // 楼层 scope为5时此参数必填，示例值："8"
	RoomID     string `query:"room_id" json:"-"`     // 会议室id scope为5时需要此参数，示例值："6383786266263"
}

type queryRoomConfigResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *QueryRoomConfigResp `json:"data,omitempty"` // -
}

type QueryRoomConfigResp struct {
	RoomBackground    string                             `json:"room_background,omitempty"`    // 飞书会议室背景图
	DisplayBackground string                             `json:"display_background,omitempty"` // 飞书签到板背景图
	DigitalSignage    *QueryRoomConfigRespDigitalSignage `json:"digital_signage,omitempty"`    // 飞书会议室数字标牌
}

type QueryRoomConfigRespDigitalSignage struct {
	Enable       bool                                         `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                         `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int                                          `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int                                          `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*QueryRoomConfigRespDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

type QueryRoomConfigRespDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材id
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int    `json:"material_type,omitempty"` // 素材类型，可用值：【1（图片），2（视频），3（GIF）】
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int    `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
}
