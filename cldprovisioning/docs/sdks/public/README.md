# Public

## Overview

View shared Permissions system utilities, including the system roles and policies catalog, Cedar policy validation, and the Cedar schema. These endpoints are publicly accessible and don't require authentication.


### Available Operations

* [GetCatalog](#getcatalog) - Get system roles and policies catalog
* [ValidatePolicy](#validatepolicy) - Validate a Cedar policy
* [GetSchema](#getschema) - Get Cedar schema

## GetCatalog

Retrieves the catalog of predefined system roles and their associated policies.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemRolesAndPoliciesCatalog" method="get" path="/v2/accounts/permissions/public/catalog" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Public.GetCatalog(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SystemRolesAndPoliciesCatalogResponse](../../models/components/systemrolesandpoliciescatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## ValidatePolicy

Validate a Cedar policy

### Example Usage

<!-- UsageSnippet language="go" operationID="validateCedarPolicy" method="post" path="/v2/accounts/permissions/public/validate" example="ValidateCedarPolicyResponseExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Public.ValidatePolicy(ctx, components.ValidateCedarPolicyRequest{
        Policy: "permit(principal == Cloudinary::APIKy::\"1234\", action == Cloudinary::Action::\"read\", resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\"asdfjkl12347890\")} ;",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.ValidateCedarPolicyRequest](../../models/components/validatecedarpolicyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.ValidateCedarPolicyResponse](../../models/components/validatecedarpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSchema

Get Cedar schema

### Example Usage

<!-- UsageSnippet language="go" operationID="getCedarSchema" method="get" path="/v2/accounts/permissions/public/schema" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Public.GetSchema(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.CedarSchemaResponse](../../models/components/cedarschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |