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

func (r *Sheet) setSheetTitle(ctx context.Context, sheetID, name string) error {
	// TODO lark sdk 还不支持
	return nil
}

func (r *Sheet) insertDimension(ctx context.Context, majorDimension, sheetID string, startIndex int, count int) error {
	_, _, err := r.larkClient.Drive.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.InsertSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: ptr.String(majorDimension),
			StartIndex:     int64(startIndex),
			EndIndex:       int64(count),
		},
		InheritStyle: nil,
	})
	return err
}

// startIndex 从 0 开始计算
// startIndex startIndex 是: 左闭区间 右闭区间，即 [start, end]
// 当上/左移的时候，destIndex 是上半边的 index，当下/右移的时候，destIndex 是下半边的 index
// dest 是: 左闭区间，右开区间，即 [dest, dest)
// -1 start end 是 左闭区间 右闭区间，dest 是做 左闭区间，右开区间
// 上移看：左区间，下移看：右区间
func (r *Sheet) moveDimension(ctx context.Context, majorDimension, sheetID string, startIndex, count, diff int) error {
	if diff == 0 {
		return nil
	}

	startIndex -= 1                    // 因为从 0 开始算
	endIndex := startIndex + count - 1 // 结束位置，
	destIndex := 0
	if diff < 0 {
		destIndex = startIndex - (-diff) // 左闭区间，不需要再+1
	} else {
		destIndex = endIndex + diff + 1 // 右开区间，所以需要 +1
	}

	_, _, err := r.larkClient.Drive.MoveSheetDimension(ctx, &lark.MoveSheetDimensionReq{
		SpreadSheetToken: r.sheetToken,
		SheetID:          sheetID,
		Source: &lark.MoveSheetDimensionReqSource{
			MajorDimension: ptr.String(majorDimension),
			StartIndex:     ptr.Int64(int64(startIndex)),
			EndIndex:       ptr.Int64(int64(endIndex)),
		},
		DestinationIndex: ptr.Int64(int64(destIndex)),
	})
	return err
}

func (r *Sheet) appendDimension(ctx context.Context, cellRange string, values [][]lark.SheetContent, option *string) error {
	_, _, err := r.larkClient.Drive.AppendSheetValue(ctx, &lark.AppendSheetValueReq{
		InsertDataOption: option,
		SpreadSheetToken: r.sheetToken,
		ValueRange: &lark.AppendSheetValueReqValueRange{
			Range:  cellRange,
			Values: values,
		},
	})
	return err
}

func (r *Sheet) addDimension(ctx context.Context, dimension string, sheetID string, count int) error {
	_, _, err := r.larkClient.Drive.AddSheetDimensionRange(ctx, &lark.AddSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.AddSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: &dimension,
			Length:         int64(count),
		},
	})
	return err
}

func (r *Sheet) updateDimension(ctx context.Context, dimension, sheetID string, startIndex, count int, visible *bool, fixedSize *int64) error {
	_, _, err := r.larkClient.Drive.UpdateSheetDimensionRange(ctx, &lark.UpdateSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.UpdateSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: &dimension,
			StartIndex:     int64(startIndex),
			EndIndex:       int64(startIndex + count - 1),
		},
		DimensionProperties: &lark.UpdateSheetDimensionRangeReqDimensionProperties{
			Visible:   visible,
			FixedSize: fixedSize,
		},
	})
	return err
}

func (r *Sheet) deleteDimension(ctx context.Context, dimension, sheetID string, startIndex, count int) error {
	_, _, err := r.larkClient.Drive.DeleteSheetDimensionRange(ctx, &lark.DeleteSheetDimensionRangeReq{
		SpreadSheetToken: r.sheetToken,
		Dimension: &lark.DeleteSheetDimensionRangeReqDimension{
			SheetID:        sheetID,
			MajorDimension: &dimension,
			StartIndex:     int64(startIndex),
			EndIndex:       int64(startIndex + count - 1),
		},
	})
	return err
}

// TODO sheet 内容解析需要完善
func (r *Sheet) getValue(ctx context.Context, cellRange string, option *lark.GetSheetValueReq) error {
	if option == nil {
		option = &lark.GetSheetValueReq{}
	}
	res, _, err := r.larkClient.Drive.GetSheetValue(ctx, &lark.GetSheetValueReq{
		ValueRenderOption:    option.ValueRenderOption,
		DateTimeRenderOption: option.DateTimeRenderOption,
		UserIDType:           option.UserIDType,
		SpreadSheetToken:     r.sheetToken,
		Range:                cellRange,
	})
	_ = res
	return err
}

func (r *Sheet) setCellStyle(ctx context.Context, cellRange string, style *lark.SetSheetStyleReqAppendStyleStyle) error {
	_, _, err := r.larkClient.Drive.SetSheetStyle(ctx, &lark.SetSheetStyleReq{
		SpreadSheetToken: r.sheetToken,
		AppendStyle: &lark.SetSheetStyleReqAppendStyle{
			Range: cellRange,
			Style: style,
		},
	})
	return err
}

func (r *Sheet) batchSetCellStyle(ctx context.Context, styles []*lark.BatchSetSheetStyleReqData) error {
	_, _, err := r.larkClient.Drive.BatchSetSheetStyle(ctx, &lark.BatchSetSheetStyleReq{
		SpreadSheetToken: r.sheetToken,
		Data:             styles,
	})
	return err
}

func (r *Sheet) mergeCell(ctx context.Context, cellRange, mergeType string) error {
	if mergeType == "" {
		mergeType = "MERGE_ALL"
	}
	_, _, err := r.larkClient.Drive.MergeSheetCell(ctx, &lark.MergeSheetCellReq{
		SpreadSheetToken: r.sheetToken,
		Range:            cellRange,
		MergeType:        mergeType,
	})
	return err
}

func (r *Sheet) unmergeCell(ctx context.Context, cellRange string) error {
	_, _, err := r.larkClient.Drive.UnmergeSheetCell(ctx, &lark.UnmergeSheetCellReq{
		SpreadSheetToken: r.sheetToken,
		Range:            cellRange,
	})
	return err
}

func (r *Sheet) setSheetValue(ctx context.Context, cellRange string, contents [][]lark.SheetContent) error {
	_, _, err := r.larkClient.Drive.SetSheetValue(ctx, &lark.SetSheetValueReq{
		SpreadSheetToken: r.sheetToken,
		ValueRange: &lark.SetSheetValueReqValueRange{
			Range:  cellRange,
			Values: contents,
		},
	})
	return err
}

func (r *Sheet) batchSetSheetValue(ctx context.Context, values []*lark.BatchSetSheetValueReqValueRange) error {
	_, _, err := r.larkClient.Drive.BatchSetSheetValue(ctx, &lark.BatchSetSheetValueReq{
		SpreadSheetToken: r.sheetToken,
		ValueRanges:      values,
	})
	return err
}

func (r *Sheet) setSheetValueImage(ctx context.Context, cellRange string, image []byte) error {
	_, _, err := r.larkClient.Drive.SetSheetValueImage(ctx, &lark.SetSheetValueImageReq{
		SpreadSheetToken: r.sheetToken,
		Range:            cellRange,
		Image:            image,
		Name:             "a.png",
	})
	return err
}

func (r *Sheet) search(ctx context.Context, sheetID, value string, condition *lark.FindSheetReqFindCondition) (*lark.FindSheetRespFindResult, error) {
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

func (r *Sheet) replace(ctx context.Context, sheetID, old, new string, condition *lark.ReplaceSheetReqFindCondition) (*lark.ReplaceSheetRespReplaceResult, error) {
	if condition == nil {
		condition = &lark.ReplaceSheetReqFindCondition{}
	}
	if condition.Range == "" {
		condition.Range = sheetID
	}
	resp, _, err := r.larkClient.Drive.ReplaceSheet(ctx, &lark.ReplaceSheetReq{
		SpreadSheetToken: r.sheetToken,
		SheetID:          sheetID,
		FindCondition:    condition,
		Find:             old,
		Replacement:      new,
	})
	if err != nil {
		return nil, err
	}
	return resp.ReplaceResult, nil
}
