package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/chyroc/lark"
)

func main() {
	appID := ""
	appSecret := ""
	accessToken := ""
	refreshToken := ""
	file := ""
	flag.StringVar(&appID, "app_id", "", "")
	flag.StringVar(&appSecret, "app_secret", "", "")
	flag.StringVar(&accessToken, "access_token", "", "")
	flag.StringVar(&refreshToken, "refresh_token", "", "")
	flag.StringVar(&file, "file", "", "")
	flag.Parse()

	if appID == "" {
		appID = os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID")
	}
	if appSecret == "" {
		appSecret = os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET")
	}
	if file == "" {
		file = "./token.json"
	}
	if accessToken == "" {
		accessToken = os.Getenv("LARK_ACCESS_TOKEN_ALL_PERMISSION_APP")
	}
	if refreshToken == "" {
		refreshToken = os.Getenv("LARK_REFRESH_TOKEN_ALL_PERMISSION_APP")
	}
	refreshTimestamp := os.Getenv("INTERNAL_REFRESH_TIMESTAMP")

	fmt.Println("start refresh token")
	defer func() {
		fmt.Println("end refresh token")
	}()

	timestamp := int64(0)
	if refreshTimestamp == "" {
		timestamp = time.Now().Unix()
	} else {
		t, err := strconv.ParseInt(refreshTimestamp, 10, 64)
		if err != nil {
			panic(err)
		}
		timestamp = t
	}
	if timestamp == 0 {
		panic(fmt.Sprintf("get timestamp fail"))
	}

	// 60*60<=3600
	if time.Now().Unix()-timestamp < 3600 {
		bs, _ := json.Marshal(map[string]interface{}{
			"quick": "true",
		})
		if err := ioutil.WriteFile(file, bs, 0o666); err != nil {
			panic(err)
		}
		return
	}

	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential(appID, appSecret))

	resp, _, err := cli.Auth.RefreshAccessToken(ctx, &lark.RefreshAccessTokenReq{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	})
	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(map[string]interface{}{
		"refresh_token": resp.RefreshToken,
		"access_token":  resp.AccessToken,
		"timestamp":     strconv.FormatInt(timestamp, 10),
	})
	if err = ioutil.WriteFile(file, bs, 0o666); err != nil {
		panic(err)
	}
}
