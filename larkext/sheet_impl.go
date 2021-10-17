package larkext

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

func (r *Sheet) meta(ctx context.Context) (*lark.GetSheetMetaResp, error) {
	res, _, err := r.larkClient.Drive.GetSheetMeta(ctx, &lark.GetSheetMetaReq{
		ExtFields:        nil,
		UserIDType:       lark.IDTypePtr(lark.IDTypeOpenID),
		SpreadSheetToken: r.sheetToken,
	})
	return res, err
}

func (r *Sheet) delete(ctx context.Context) error {
	_, _, err := r.larkClient.Drive.DeleteDriveSheetFile(ctx, &lark.DeleteDriveSheetFileReq{
		SpreadSheetToken: r.sheetToken,
	})
	return err
}

func (r *Sheet) setTitle(ctx context.Context, title string) error {
	_, _, err := r.larkClient.Drive.UpdateSheetProperty(ctx, &lark.UpdateSheetPropertyReq{
		SpreadSheetToken: r.sheetToken,
		Properties: &lark.UpdateSheetPropertyReqProperties{
			Title: title,
		},
	})
	return err
}

func (r *Sheet) newSheet(ctx context.Context, title string) (string, error) {
	res, _, err := r.larkClient.Drive.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{
		SpreadSheetToken: r.sheetToken,
		Requests: []*lark.BatchUpdateSheetReqRequest{
			{
				AddSheet: &lark.BatchUpdateSheetReqRequestAddSheet{
					Properties: &lark.BatchUpdateSheetReqRequestAddSheetProperties{
						Title: title,
						Index: nil,
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}
	if len(res.Replies) > 0 && res.Replies[0].AddSheet != nil && res.Replies[0].AddSheet.Properties != nil {
		return res.Replies[0].AddSheet.Properties.SheetID, nil
	}
	return "", fmt.Errorf("new sheet empty response")
}

func (r *Sheet) deleteSheet(ctx context.Context, sheetID string) error {
	_, _, err := r.larkClient.Drive.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{
		SpreadSheetToken: r.sheetToken,
		Requests: []*lark.BatchUpdateSheetReqRequest{
			{
				DeleteSheet: &lark.BatchUpdateSheetReqRequestDeleteSheet{
					SheetID: sheetID,
				},
			},
		},
	})
	return err
}

func (r *Sheet) copySheet(ctx context.Context, sheetID string, title *string) (string, error) {
	resp, _, err := r.larkClient.Drive.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{
		SpreadSheetToken: r.sheetToken,
		Requests: []*lark.BatchUpdateSheetReqRequest{
			{
				CopySheet: &lark.BatchUpdateSheetReqRequestCopySheet{
					Source: &lark.BatchUpdateSheetReqRequestCopySheetSource{
						SheetID: sheetID,
					},
					Destination: &lark.BatchUpdateSheetReqRequestCopySheetDestination{
						Title: title,
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}
	if len(resp.Replies) > 0 && resp.Replies[0].CopySheet != nil && resp.Replies[0].CopySheet.Properties != nil {
		return resp.Replies[0].CopySheet.Properties.SheetID, nil
	}
	return "", fmt.Errorf("copy sheet empty response")
}

func (r *Sheet) setSheetName(ctx context.Context, sheetID, name string) error {
	// TODO lark sdk 还不支持
	return nil
}

func (r *Sheet) insertCol(ctx context.Context, sheetID string, startIndex int, count int) error {
	_, _, err := r.larkClient.Drive.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.InsertSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: ptr.String("COLUMNS"),
			StartIndex:     int64(startIndex),
			EndIndex:       int64(startIndex + count),
		},
		InheritStyle: nil,
	})
	return err
}

func (r *Sheet) insertRow(ctx context.Context, sheetID string, startIndex int, count int) error {
	_, _, err := r.larkClient.Drive.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.InsertSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: ptr.String("ROWS"),
			StartIndex:     int64(startIndex),
			EndIndex:       int64(startIndex + count),
		},
		InheritStyle: nil,
	})
	return err
}

func (r *Sheet) searchSheet(ctx context.Context, sheetID, value string, condition *lark.FindSheetReqFindCondition) (*lark.FindSheetRespFindResult, error) {
	if condition == nil {
		condition = new(lark.FindSheetReqFindCondition)
	}
	if condition.Range == "" {
		condition.Range = sheetID
	}
	resp, _, err := r.larkClient.Drive.FindSheet(ctx, &lark.FindSheetReq{
		SpreadSheetToken: r.sheetToken,
		SheetID:          sheetID,
		FindCondition:    condition,
		Find:             value,
	})
	if err != nil {
		return nil, err
	}
	return resp.FindResult, nil
}
