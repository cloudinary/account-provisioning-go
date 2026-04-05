package hooks

import (
	"fmt"
	"net/url"
	"os"
)

// AccountConfig holds Cloudinary account provisioning credentials parsed from
// CLOUDINARY_ACCOUNT_URL or individual CLOUDINARY_ACCOUNT_* environment variables.
//
// URL format: account://<provisioning_api_key>:<provisioning_api_secret>@<account_id>
type AccountConfig struct {
	AccountID            string
	ProvisioningAPIKey   string
	ProvisioningAPISecret string
}

// NewAccountConfig creates a new AccountConfig populated from environment variables.
func NewAccountConfig() *AccountConfig {
	cfg := &AccountConfig{}

	if accountURL := os.Getenv("CLOUDINARY_ACCOUNT_URL"); accountURL != "" {
		cfg.loadFromURL(accountURL)
	}

	if v := os.Getenv("CLOUDINARY_ACCOUNT_ID"); v != "" {
		cfg.AccountID = v
	}
	if v := os.Getenv("CLOUDINARY_PROVISIONING_API_KEY"); v != "" {
		cfg.ProvisioningAPIKey = v
	}
	if v := os.Getenv("CLOUDINARY_PROVISIONING_API_SECRET"); v != "" {
		cfg.ProvisioningAPISecret = v
	}

	return cfg
}

func (c *AccountConfig) loadFromURL(rawURL string) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(fmt.Sprintf("invalid CLOUDINARY_ACCOUNT_URL: %v", err))
	}
	if parsed.Scheme != "account" {
		panic("invalid CLOUDINARY_ACCOUNT_URL scheme, expecting 'account://'")
	}

	c.AccountID = parsed.Hostname()
	c.ProvisioningAPIKey = parsed.User.Username()
	if secret, ok := parsed.User.Password(); ok {
		c.ProvisioningAPISecret = secret
	}
}
