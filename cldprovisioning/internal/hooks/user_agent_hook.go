package hooks

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cloudinary/account-provisioning-go/cldprovisioning/internal/config"
)

// UserAgentHook replaces the default Speakeasy user-agent with a Cloudinary-branded one.
//
// Format: Cloudinary/AccountProvisioning Go/{sdkVersion} Gen/{genVersion} Schema/{schemaVersion} (Go {goVersion}; {os} {arch})
type UserAgentHook struct{}

func (h *UserAgentHook) SDKInit(sdkConfig config.SDKConfiguration) config.SDKConfiguration {
	original := sdkConfig.UserAgent
	if original == "" || !strings.HasPrefix(original, "speakeasy-sdk/") {
		return sdkConfig
	}

	parts := strings.SplitN(original, " ", 5)
	if len(parts) < 5 {
		return sdkConfig
	}

	sdkVersion := parts[1]
	genVersion := parts[2]
	schemaVersion := parts[3]

	goVersion := strings.TrimPrefix(runtime.Version(), "go")
	systemInfo := fmt.Sprintf("Go %s; %s %s", goVersion, runtime.GOOS, runtime.GOARCH)

	sdkConfig.UserAgent = fmt.Sprintf(
		"Cloudinary/AccountProvisioning Go/%s Gen/%s Schema/%s (%s)",
		sdkVersion, genVersion, schemaVersion, systemInfo,
	)

	return sdkConfig
}
