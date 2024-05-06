package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatRequestBody(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, map[string]interface{}{
			"app_id":       "app",
			"lang":         "lang",
			"user_id_type": "app_id",
		}, FormatRequestBody(GetApplicationReq{
			AppID:      "app",
			Lang:       "lang",
			UserIDType: IDTypePtr(IDTypeAppID),
		}))
	})
	t.Run("", func(t *testing.T) {
		assert.Equal(t, map[string]interface{}{
			"user_ids":     []string{"1"},
			"user_id_type": "app_id",
			"app_feed_card": map[string]interface{}{
				"biz_id": "biz",
				"link": map[string]interface{}{
					"link": "link",
				},
			},
		}, FormatRequestBody(CreateAppFeedCardReq{
			UserIDType: IDTypePtr(IDTypeAppID),
			AppFeedCard: &CreateAppFeedCardReqAppFeedCard{
				BizID: &[]string{"biz"}[0],
				Link: &CreateAppFeedCardReqAppFeedCardLink{
					Link: &[]string{"link"}[0],
				},
			},
			UserIDs: []string{"1"},
		}))
	})
}
