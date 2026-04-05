package hooks

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/cloudinary/account-provisioning-go/cldprovisioning/internal/config"
)

// CloudinaryAccountHook populates account_id from CLOUDINARY_ACCOUNT_URL on init
// and attaches Basic auth (provisioning_api_key:provisioning_api_secret) to every request.
//
// Credential resolution order:
//  1. Explicit security passed to the SDK constructor
//  2. CLOUDINARY_ACCOUNT_URL / individual CLOUDINARY_ACCOUNT_* env vars
type CloudinaryAccountHook struct {
	config *AccountConfig
}

func NewCloudinaryAccountHook(cfg *AccountConfig) *CloudinaryAccountHook {
	return &CloudinaryAccountHook{config: cfg}
}

func (h *CloudinaryAccountHook) SDKInit(sdkConfig config.SDKConfiguration) config.SDKConfiguration {
	if sdkConfig.Globals.AccountID == nil && h.config.AccountID != "" {
		id := h.config.AccountID
		sdkConfig.Globals.AccountID = &id
	}
	return sdkConfig
}

func (h *CloudinaryAccountHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	key, secret, err := h.resolveCredentials(hookCtx)
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", key, secret)))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))

	return req, nil
}

func (h *CloudinaryAccountHook) resolveCredentials(hookCtx BeforeRequestContext) (string, string, error) {
	key := h.config.ProvisioningAPIKey
	secret := h.config.ProvisioningAPISecret

	if hookCtx.SecuritySource != nil {
		sec, err := hookCtx.SecuritySource(hookCtx.Context)
		if err == nil && sec != nil {
			if s, ok := sec.(interface{ GetProvisioningAPIKey() *string }); ok {
				if v := s.GetProvisioningAPIKey(); v != nil && *v != "" {
					key = *v
				}
			}
			if s, ok := sec.(interface{ GetProvisioningAPISecret() *string }); ok {
				if v := s.GetProvisioningAPISecret(); v != nil && *v != "" {
					secret = *v
				}
			}
		}
	}

	if key == "" || secret == "" {
		return "", "", fmt.Errorf("provisioning API key and secret are required")
	}

	return key, secret, nil
}
