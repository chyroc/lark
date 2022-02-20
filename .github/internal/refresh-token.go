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
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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

	fmt.Println("start refresh token")
	defer func() {
		fmt.Println("end refresh token")
	}()

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
	})
	if err = ioutil.WriteFile(file, bs, 0o666); err != nil {
		panic(err)
	}
}
