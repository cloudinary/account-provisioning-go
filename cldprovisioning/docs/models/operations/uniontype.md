# UnionType

Whether to return users who last logged in within the specified date range (include) or those who didn't
last log in within the range (exclude). **Possible values**: `include`, `exclude`. **Default**: `include`.


## Example Usage

```go
import (
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
)

value := operations.UnionTypeInclude
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `UnionTypeInclude` | include            |
| `UnionTypeExclude` | exclude            |