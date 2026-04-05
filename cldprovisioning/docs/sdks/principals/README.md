# Principals

## Overview

View roles associated with a specified principal, and add or remove a role associated with a principal.


### Available Operations

* [ListRoles](#listroles) - Get a principal's roles
* [UpdateRoles](#updateroles) - Assign roles to a principal
* [Inspect](#inspect) - Inspect
* [InspectMultiple](#inspectmultiple) - Inspect multiple

## ListRoles

Retrieve all roles associated with a specific principal.

### Example Usage: GetPrincipalRolesExample

<!-- UsageSnippet language="go" operationID="getPrincipalRoles" method="get" path="/v2/accounts/{account_id}/permissions/principal_roles" example="GetPrincipalRolesExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.ListRoles(ctx, operations.GetPrincipalRolesRequest{
        PrincipalType: components.PrincipalTypeEnumUser,
        PrincipalID: "<id>",
        PermissionType: components.PermissionTypeEnumGlobal,
        ManagementType: components.ManagementTypeEnumSystem.ToPointer(),
        ScopeType: components.ScopeTypeEnumProdenv.ToPointer(),
        ParamKey: []string{
            "folder_id",
        },
        ParamValue: []string{
            "asdfjkl12347890",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: GetPrincipalRolesWithParamKeyValueExample

<!-- UsageSnippet language="go" operationID="getPrincipalRoles" method="get" path="/v2/accounts/{account_id}/permissions/principal_roles" example="GetPrincipalRolesWithParamKeyValueExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.ListRoles(ctx, operations.GetPrincipalRolesRequest{
        PrincipalType: components.PrincipalTypeEnumUser,
        PrincipalID: "<id>",
        PermissionType: components.PermissionTypeEnumGlobal,
        ManagementType: components.ManagementTypeEnumSystem.ToPointer(),
        ScopeType: components.ScopeTypeEnumProdenv.ToPointer(),
        ParamKey: []string{
            "folder_id",
        },
        ParamValue: []string{
            "asdfjkl12347890",
        },
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetPrincipalRolesRequest](../../models/operations/getprincipalrolesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.PrincipalRolesResponse](../../models/components/principalrolesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 404                      | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## UpdateRoles

Add or remove roles associated with a principal.

### Example Usage: InvalidOperationFieldErrorExample

<!-- UsageSnippet language="go" operationID="updatePrincipalRoles" method="put" path="/v2/accounts/{account_id}/permissions/principal_roles" example="InvalidOperationFieldErrorExample" -->
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
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    err := s.Principals.UpdateRoles(ctx, components.UpdatePrincipalRolesRequest{
        Operation: components.OperationEnumAdd.ToPointer(),
        Principal: &components.Principal{
            Type: components.PrincipalTypeEnumUser,
            ID: "1234abc",
        },
        Roles: []components.RoleToManage{
            components.RoleToManage{
                ID: "marketing_content_contributor_1357fhe",
                ScopeID: cldprovisioning.Pointer("975l29lz02jt0836fhwi"),
                PolicyParameters: &components.RoleToManagePolicyParameters{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```
### Example Usage: missingPrincipalFieldsErrorExample

<!-- UsageSnippet language="go" operationID="updatePrincipalRoles" method="put" path="/v2/accounts/{account_id}/permissions/principal_roles" example="missingPrincipalFieldsErrorExample" -->
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
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    err := s.Principals.UpdateRoles(ctx, components.UpdatePrincipalRolesRequest{
        Operation: components.OperationEnumAdd.ToPointer(),
        Principal: &components.Principal{
            Type: components.PrincipalTypeEnumUser,
            ID: "1234abc",
        },
        Roles: []components.RoleToManage{
            components.RoleToManage{
                ID: "marketing_content_contributor_1357fhe",
                ScopeID: cldprovisioning.Pointer("975l29lz02jt0836fhwi"),
                PolicyParameters: &components.RoleToManagePolicyParameters{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.UpdatePrincipalRolesRequest](../../models/components/updateprincipalrolesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 404                 | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Inspect

Retrieves principals and their assigned roles within a specified scope, based on query filters.

Use this endpoint to determine which users, groups, API keys, or account API keys have roles applied at:
* A specific `scope_type` ("account" or "prodenv")
* A specific `scope_id` (for product environments)
* Optional content filters (such as `folder_id`, `collection_id`, or `param_key` / `param_value`)
* Optional `principal_type` (and `principal_id`)

The response includes all principals that match the specified filters, along with the roles applied to them within that context.

This endpoint is useful when you want to:
* Audit who has access within a given scope
* Identify all principals assigned to roles for a specific folder, collection, asset, or product environment
* Retrieve role assignments dynamically based on filtering criteria


### Example Usage: InspectExampleAllFolders

<!-- UsageSnippet language="go" operationID="inspect" method="get" path="/v2/accounts/{account_id}/permissions/principal_roles/inspect" example="InspectExampleAllFolders" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.Inspect(ctx, operations.InspectRequest{
        ScopeType: components.ScopeTypeEnumProdenv,
        PrincipalType: components.PrincipalTypeEnumUser.ToPointer(),
        ParamKey: []string{
            "folder_id",
        },
        ParamValue: []string{
            "asdfjkl12347890",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: InspectExampleFolder

<!-- UsageSnippet language="go" operationID="inspect" method="get" path="/v2/accounts/{account_id}/permissions/principal_roles/inspect" example="InspectExampleFolder" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.Inspect(ctx, operations.InspectRequest{
        ScopeType: components.ScopeTypeEnumProdenv,
        PrincipalType: components.PrincipalTypeEnumUser.ToPointer(),
        ParamKey: []string{
            "folder_id",
        },
        ParamValue: []string{
            "asdfjkl12347890",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: InspectExampleProdEnv

<!-- UsageSnippet language="go" operationID="inspect" method="get" path="/v2/accounts/{account_id}/permissions/principal_roles/inspect" example="InspectExampleProdEnv" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cldprovisioning.New(
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.Inspect(ctx, operations.InspectRequest{
        ScopeType: components.ScopeTypeEnumProdenv,
        PrincipalType: components.PrincipalTypeEnumUser.ToPointer(),
        ParamKey: []string{
            "folder_id",
        },
        ParamValue: []string{
            "asdfjkl12347890",
        },
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.InspectRequest](../../models/operations/inspectrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*components.PrincipalRolesInspectResponse](../../models/components/principalrolesinspectresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 404                      | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## InspectMultiple

Retrieves role assignments for a specified list of principals within a shared scope and optional policy parameters.

Use this endpoint when you already know which principals you want to evaluate and want to check their role assignments in a single request.

In this request:
* You explicitly provide the `principals` array in the request body.
* All principals are evaluated against the same `scope_type`, optional `scope_id`, and optional `policy_parameters`.
* Unlike `Inspect`, this endpoint doesn't search for matching principals. It evaluates only the principals you provide.

This endpoint is useful for:
* Verify role assignments for multiple specified principals at once
* Comparing role assignments across specific users, groups, or keys
* Reducing multiple per-principal `Inspect` calls into a single request


### Example Usage: InspectMultipleRequestExample

<!-- UsageSnippet language="go" operationID="inspectMultiple" method="post" path="/v2/accounts/{account_id}/permissions/principal_roles/inspect" example="InspectMultipleRequestExample" -->
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
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.InspectMultiple(ctx, components.InspectRequest{
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: cldprovisioning.Pointer("975l29lz02jt0836fhwi"),
        Principals: []components.Principal{
            components.Principal{
                Type: components.PrincipalTypeEnumUser,
                ID: "1234abc",
            },
            components.Principal{
                Type: components.PrincipalTypeEnumUser,
                ID: "4567xyz",
            },
            components.Principal{
                Type: components.PrincipalTypeEnumAPIKey,
                ID: "a382ltieo893jhioqpg8urp",
            },
        },
        PolicyParameters: &components.InspectRequestPolicyParameters{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: InspectMultipleResponseExample

<!-- UsageSnippet language="go" operationID="inspectMultiple" method="post" path="/v2/accounts/{account_id}/permissions/principal_roles/inspect" example="InspectMultipleResponseExample" -->
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
        cldprovisioning.WithAccountID("<id>"),
        cldprovisioning.WithSecurity(components.Security{
            ProvisioningAPIKey: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
            ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
        }),
    )

    res, err := s.Principals.InspectMultiple(ctx, components.InspectRequest{
        ScopeType: components.ScopeTypeEnumProdenv,
        Principals: []components.Principal{
            components.Principal{
                Type: components.PrincipalTypeEnumUser,
                ID: "1234abc",
            },
        },
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.InspectRequest](../../models/components/inspectrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*components.PrincipalRolesInspectResponse](../../models/components/principalrolesinspectresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 404                      | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |