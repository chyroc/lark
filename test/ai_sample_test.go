// Code generated by lark_sdk_gen. DO NOT EDIT.
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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_AI_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ParseAIResume(ctx, &lark.ParseAIResumeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIParseAIResume(func(ctx context.Context, request *lark.ParseAIResumeReq, options ...lark.MethodOptionFunc) (*lark.ParseAIResumeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIParseAIResume()

			_, _, err := moduleCli.ParseAIResume(ctx, &lark.ParseAIResumeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIVehicleInvoice(func(ctx context.Context, request *lark.RecognizeAIVehicleInvoiceReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIVehicleInvoiceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIVehicleInvoice()

			_, _, err := moduleCli.RecognizeAIVehicleInvoice(ctx, &lark.RecognizeAIVehicleInvoiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIHealthCertificate(func(ctx context.Context, request *lark.RecognizeAIHealthCertificateReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIHealthCertificateResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIHealthCertificate()

			_, _, err := moduleCli.RecognizeAIHealthCertificate(ctx, &lark.RecognizeAIHealthCertificateReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIHkmMainlandTravelPermit(func(ctx context.Context, request *lark.RecognizeAIHkmMainlandTravelPermitReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIHkmMainlandTravelPermitResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIHkmMainlandTravelPermit()

			_, _, err := moduleCli.RecognizeAIHkmMainlandTravelPermit(ctx, &lark.RecognizeAIHkmMainlandTravelPermitReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAITwMainlandTravelPermit(func(ctx context.Context, request *lark.RecognizeAITwMainlandTravelPermitReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAITwMainlandTravelPermitResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAITwMainlandTravelPermit()

			_, _, err := moduleCli.RecognizeAITwMainlandTravelPermit(ctx, &lark.RecognizeAITwMainlandTravelPermitReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIChinesePassport(func(ctx context.Context, request *lark.RecognizeAIChinesePassportReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIChinesePassportResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIChinesePassport()

			_, _, err := moduleCli.RecognizeAIChinesePassport(ctx, &lark.RecognizeAIChinesePassportReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIBankCard(func(ctx context.Context, request *lark.RecognizeAIBankCardReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIBankCardResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIBankCard()

			_, _, err := moduleCli.RecognizeAIBankCard(ctx, &lark.RecognizeAIBankCardReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIVehicleLicense(func(ctx context.Context, request *lark.RecognizeAIVehicleLicenseReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIVehicleLicenseResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIVehicleLicense()

			_, _, err := moduleCli.RecognizeAIVehicleLicense(ctx, &lark.RecognizeAIVehicleLicenseReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAITrainInvoice(func(ctx context.Context, request *lark.RecognizeAITrainInvoiceReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAITrainInvoiceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAITrainInvoice()

			_, _, err := moduleCli.RecognizeAITrainInvoice(ctx, &lark.RecognizeAITrainInvoiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAITaxiInvoice(func(ctx context.Context, request *lark.RecognizeAITaxiInvoiceReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAITaxiInvoiceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAITaxiInvoice()

			_, _, err := moduleCli.RecognizeAITaxiInvoice(ctx, &lark.RecognizeAITaxiInvoiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAiidCard(func(ctx context.Context, request *lark.RecognizeAiidCardReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAiidCardResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAiidCard()

			_, _, err := moduleCli.RecognizeAiidCard(ctx, &lark.RecognizeAiidCardReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIFoodProduceLicense(func(ctx context.Context, request *lark.RecognizeAIFoodProduceLicenseReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIFoodProduceLicenseResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIFoodProduceLicense()

			_, _, err := moduleCli.RecognizeAIFoodProduceLicense(ctx, &lark.RecognizeAIFoodProduceLicenseReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIFoodManageLicense(func(ctx context.Context, request *lark.RecognizeAIFoodManageLicenseReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIFoodManageLicenseResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIFoodManageLicense()

			_, _, err := moduleCli.RecognizeAIFoodManageLicense(ctx, &lark.RecognizeAIFoodManageLicenseReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIDrivingLicense(func(ctx context.Context, request *lark.RecognizeAIDrivingLicenseReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIDrivingLicenseResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIDrivingLicense()

			_, _, err := moduleCli.RecognizeAIDrivingLicense(ctx, &lark.RecognizeAIDrivingLicenseReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIVatInvoice(func(ctx context.Context, request *lark.RecognizeAIVatInvoiceReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIVatInvoiceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIVatInvoice()

			_, _, err := moduleCli.RecognizeAIVatInvoice(ctx, &lark.RecognizeAIVatInvoiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIBusinessLicense(func(ctx context.Context, request *lark.RecognizeAIBusinessLicenseReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIBusinessLicenseResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIBusinessLicense()

			_, _, err := moduleCli.RecognizeAIBusinessLicense(ctx, &lark.RecognizeAIBusinessLicenseReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIExtractAIContractField(func(ctx context.Context, request *lark.ExtractAIContractFieldReq, options ...lark.MethodOptionFunc) (*lark.ExtractAIContractFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIExtractAIContractField()

			_, _, err := moduleCli.ExtractAIContractField(ctx, &lark.ExtractAIContractFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeAIBusinessCard(func(ctx context.Context, request *lark.RecognizeAIBusinessCardReq, options ...lark.MethodOptionFunc) (*lark.RecognizeAIBusinessCardResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeAIBusinessCard()

			_, _, err := moduleCli.RecognizeAIBusinessCard(ctx, &lark.RecognizeAIBusinessCardReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeBasicImage(func(ctx context.Context, request *lark.RecognizeBasicImageReq, options ...lark.MethodOptionFunc) (*lark.RecognizeBasicImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeBasicImage()

			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeSpeechStream(func(ctx context.Context, request *lark.RecognizeSpeechStreamReq, options ...lark.MethodOptionFunc) (*lark.RecognizeSpeechStreamResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeSpeechStream()

			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIRecognizeSpeechFile(func(ctx context.Context, request *lark.RecognizeSpeechFileReq, options ...lark.MethodOptionFunc) (*lark.RecognizeSpeechFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeSpeechFile()

			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAITranslateText(func(ctx context.Context, request *lark.TranslateTextReq, options ...lark.MethodOptionFunc) (*lark.TranslateTextResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAITranslateText()

			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIDetectTextLanguage(func(ctx context.Context, request *lark.DetectTextLanguageReq, options ...lark.MethodOptionFunc) (*lark.DetectTextLanguageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIDetectTextLanguage()

			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAIDetectFaceAttributes(func(ctx context.Context, request *lark.DetectFaceAttributesReq, options ...lark.MethodOptionFunc) (*lark.DetectFaceAttributesResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIDetectFaceAttributes()

			_, _, err := moduleCli.DetectFaceAttributes(ctx, &lark.DetectFaceAttributesReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ParseAIResume(ctx, &lark.ParseAIResumeReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVehicleInvoice(ctx, &lark.RecognizeAIVehicleInvoiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIHealthCertificate(ctx, &lark.RecognizeAIHealthCertificateReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIHkmMainlandTravelPermit(ctx, &lark.RecognizeAIHkmMainlandTravelPermitReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITwMainlandTravelPermit(ctx, &lark.RecognizeAITwMainlandTravelPermitReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIChinesePassport(ctx, &lark.RecognizeAIChinesePassportReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBankCard(ctx, &lark.RecognizeAIBankCardReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVehicleLicense(ctx, &lark.RecognizeAIVehicleLicenseReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITrainInvoice(ctx, &lark.RecognizeAITrainInvoiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITaxiInvoice(ctx, &lark.RecognizeAITaxiInvoiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAiidCard(ctx, &lark.RecognizeAiidCardReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIFoodProduceLicense(ctx, &lark.RecognizeAIFoodProduceLicenseReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIFoodManageLicense(ctx, &lark.RecognizeAIFoodManageLicenseReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIDrivingLicense(ctx, &lark.RecognizeAIDrivingLicenseReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVatInvoice(ctx, &lark.RecognizeAIVatInvoiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBusinessLicense(ctx, &lark.RecognizeAIBusinessLicenseReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ExtractAIContractField(ctx, &lark.ExtractAIContractFieldReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBusinessCard(ctx, &lark.RecognizeAIBusinessCardReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DetectFaceAttributes(ctx, &lark.DetectFaceAttributesReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.AI
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ParseAIResume(ctx, &lark.ParseAIResumeReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVehicleInvoice(ctx, &lark.RecognizeAIVehicleInvoiceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIHealthCertificate(ctx, &lark.RecognizeAIHealthCertificateReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIHkmMainlandTravelPermit(ctx, &lark.RecognizeAIHkmMainlandTravelPermitReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITwMainlandTravelPermit(ctx, &lark.RecognizeAITwMainlandTravelPermitReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIChinesePassport(ctx, &lark.RecognizeAIChinesePassportReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBankCard(ctx, &lark.RecognizeAIBankCardReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVehicleLicense(ctx, &lark.RecognizeAIVehicleLicenseReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITrainInvoice(ctx, &lark.RecognizeAITrainInvoiceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAITaxiInvoice(ctx, &lark.RecognizeAITaxiInvoiceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAiidCard(ctx, &lark.RecognizeAiidCardReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIFoodProduceLicense(ctx, &lark.RecognizeAIFoodProduceLicenseReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIFoodManageLicense(ctx, &lark.RecognizeAIFoodManageLicenseReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIDrivingLicense(ctx, &lark.RecognizeAIDrivingLicenseReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIVatInvoice(ctx, &lark.RecognizeAIVatInvoiceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBusinessLicense(ctx, &lark.RecognizeAIBusinessLicenseReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ExtractAIContractField(ctx, &lark.ExtractAIContractFieldReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeAIBusinessCard(ctx, &lark.RecognizeAIBusinessCardReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DetectFaceAttributes(ctx, &lark.DetectFaceAttributesReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
