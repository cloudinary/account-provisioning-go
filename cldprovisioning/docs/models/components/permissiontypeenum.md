# PermissionTypeEnum

Specifies whether permissions apply globally or to specific content.

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
)

value := components.PermissionTypeEnumGlobal

// Open enum: custom values can be created with a direct type cast
custom := components.PermissionTypeEnum("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PermissionTypeEnumGlobal`  | global                      |
| `PermissionTypeEnumContent` | content                     |