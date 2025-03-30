package sso

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/adwinugroho/wedding-management-system/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     config.SSOConfig.GoogleClientID,
	ClientSecret: config.SSOConfig.GoogleClientSecret,
	RedirectURL:  config.SSOConfig.GoogleRedirectURL,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func GetGoogleOAuthConfig() *oauth2.Config {
	return googleOAuthConfig
}

func GetGoogleOAuthURL(state string) string {
	return googleOAuthConfig.AuthCodeURL(state)
}

func GetGoogleOAuthToken(code string, ctx context.Context) (*oauth2.Token, error) {
	return googleOAuthConfig.Exchange(ctx, code)
}

func GetGoogleOAuthUserInfo(token *oauth2.Token, ctx context.Context) (*GoogleUserInfo, error) {
	resp, err := googleOAuthConfig.Client(ctx, token).Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println("Error get google user info, cause:", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read google user info, cause:", err)
		return nil, err
	}

	var userInfo GoogleUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Println("Error unmarshal google user info, cause:", err)
		return nil, err
	}

	return &userInfo, nil
}
