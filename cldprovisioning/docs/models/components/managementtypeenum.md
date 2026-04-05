# ManagementTypeEnum

Specifies whether the role's policies are managed by Cloudinary (`system`) or by the user (`custom`).

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
)

value := components.ManagementTypeEnumSystem

// Open enum: custom values can be created with a direct type cast
custom := components.ManagementTypeEnum("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ManagementTypeEnumSystem` | system                     |
| `ManagementTypeEnumCustom` | custom                     |