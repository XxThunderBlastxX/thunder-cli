package cmd

import (
	"fmt"
	"time"

	"github.com/99designs/keyring"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/internal/config"
	"github.com/XxThunderBlastxX/thunder-cli/internal/service"
)

func LoginAction(appConfig *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		client := resty.New()
		var token interface{}
		k, err := service.NewKeyRingService()
		if err != nil {
			return err
		}

		// Start the Device Authorization Flow
		res, err := client.R().
			SetFormData(map[string]string{
				"client_id": appConfig.AuthClientId,
				"scope":     appConfig.AuthScope,
				"audience":  appConfig.AuthAudience,
			}).
			Post("https://" + appConfig.AuthDomain + "/oauth/device/code")

		if err != nil {
			fmt.Println("Error requesting device code:", err)
			return err
		}

		if res.StatusCode() != 200 {
			fmt.Println("Error requesting device code")
			return nil
		}

		// Extract needed details from response
		var result struct {
			DeviceCode              string `json:"device_code"`
			UserCode                string `json:"user_code"`
			VerificationUriComplete string `json:"verification_uri_complete"`
			Interval                int    `json:"interval"`
		}

		err = client.JSONUnmarshal(res.Body(), &result)
		if err != nil {
			fmt.Println("Error unmarshalling response:", err)
			return nil
		}

		// Show URL to user or open in browser
		fmt.Println("Please open the following URL in your browser :", result.VerificationUriComplete)
		fmt.Println("and verify the user code:", result.UserCode)
		if err := browser.OpenURL(result.VerificationUriComplete); err != nil {
			fmt.Println("Error opening browser:", err)
			return err
		}

		// Poll for token
		for {
			authResponse, err := client.R().
				SetFormData(map[string]string{
					"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
					"device_code": result.DeviceCode,
					"client_id":   appConfig.AuthClientId,
				}).
				Post("https://" + appConfig.AuthDomain + "/oauth/token")

			if err == nil && authResponse.StatusCode() == 200 {
				fmt.Println("✅ Login successful.")

				if err := json.Unmarshal(authResponse.Body(), &token); err != nil {
					return err
				}

				// Saving token to local keyring
				if err := k.Set(keyring.Item{
					Key:  "AUTH_ACCESS_TOKEN",
					Data: []byte(token.(map[string]interface{})["access_token"].(string)),
				}); err != nil {
					return err
				}
				// TODO: Save refresh token

				break
			}

			time.Sleep(time.Duration(result.Interval) * time.Second)
		}

		return nil
	}
}
