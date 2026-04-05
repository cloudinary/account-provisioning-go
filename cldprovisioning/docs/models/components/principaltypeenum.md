# PrincipalTypeEnum

The type of principal.

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
)

value := components.PrincipalTypeEnumUser

// Open enum: custom values can be created with a direct type cast
custom := components.PrincipalTypeEnum("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PrincipalTypeEnumUser`            | user                               |
| `PrincipalTypeEnumGroup`           | group                              |
| `PrincipalTypeEnumAPIKey`          | apiKey                             |
| `PrincipalTypeEnumProvisioningKey` | provisioningKey                    |