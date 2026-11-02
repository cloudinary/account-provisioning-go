# ~~Role~~

The legacy role to assign. Deprecated: use `roles` (create) or `role_management` (update) instead. Can't be combined with `roles` or `role_management`.

> :warning: **DEPRECATED**: Use `roles` (create) or `role_management` (update) instead..

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
)

value := operations.RoleMasterAdmin
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `RoleMasterAdmin`       | master_admin            |
| `RoleAdmin`             | admin                   |
| `RoleBilling`           | billing                 |
| `RoleTechnicalAdmin`    | technical_admin         |
| `RoleReports`           | reports                 |
| `RoleMediaLibraryAdmin` | media_library_admin     |
| `RoleMediaLibraryUser`  | media_library_user      |