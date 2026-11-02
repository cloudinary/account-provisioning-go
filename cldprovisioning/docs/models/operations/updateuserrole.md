# ~~UpdateUserRole~~

The legacy role to assign. Deprecated: use `roles` (create) or `role_management` (update) instead. Can't be combined with `roles` or `role_management`.

> :warning: **DEPRECATED**: Use `roles` (create) or `role_management` (update) instead..

## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
)

value := operations.UpdateUserRoleMasterAdmin
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `UpdateUserRoleMasterAdmin`       | master_admin                      |
| `UpdateUserRoleAdmin`             | admin                             |
| `UpdateUserRoleBilling`           | billing                           |
| `UpdateUserRoleTechnicalAdmin`    | technical_admin                   |
| `UpdateUserRoleReports`           | reports                           |
| `UpdateUserRoleMediaLibraryAdmin` | media_library_admin               |
| `UpdateUserRoleMediaLibraryUser`  | media_library_user                |