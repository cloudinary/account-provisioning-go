# ScopeTypeEnum

Specifies where the role/policy is applied, either at the account level or within product environments.

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
)

value := components.ScopeTypeEnumAccount

// Open enum: custom values can be created with a direct type cast
custom := components.ScopeTypeEnum("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ScopeTypeEnumAccount` | account                |
| `ScopeTypeEnumProdenv` | prodenv                |