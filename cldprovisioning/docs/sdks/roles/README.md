# Roles

## Overview

All roles include one or more predefined **system policies**, and can be applied to principals to enforce these policies.

You can **manage custom roles**, where you select the system policies.

You can **view system roles**, which contain a fixed set of system policies and are provided by Cloudinary.


### Available Operations

* [List](#list) - Get roles
* [Create](#create) - Create custom role
* [Get](#get) - Get role
* [Update](#update) - Update custom role
* [Delete](#delete) - Delete custom role
* [ListPrincipals](#listprincipals) - Get a role's principals
* [UpdatePrincipals](#updateprincipals) - Assign principals to a role

## List

Retrieve all system and custom roles of a specific `permission_type` (global or content-specific). Optionally filter by `scope_type`, `management_type`, and `policy_parameters`.


### Example Usage

<!-- UsageSnippet language="go" operationID="getRoles" method="get" path="/v2/accounts/{account_id}/permissions/roles" example="GetRolesExample" -->
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

    res, err := s.Roles.List(ctx, operations.GetRolesRequest{
        ScopeType: components.ScopeTypeEnumProdenv.ToPointer(),
        PermissionType: components.PermissionTypeEnumGlobal,
        ManagementType: components.ManagementTypeEnumSystem.ToPointer(),
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetRolesRequest](../../models/operations/getrolesrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*components.RolesResponse](../../models/components/rolesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403                           | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a new custom role.

### Example Usage: createAcctRoleExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createAcctRoleExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("sensitive_account_details_manager_88825hl"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumAccount,
        Name: "Sensitive account details manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing users and account security.")),
        SystemPolicyIds: []string{
            "cld::global::users_and_groups::manage",
            "cld::policy::global::account_security::manage",
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
### Example Usage: createAcctRoleResponseExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createAcctRoleResponseExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: createContentRoleExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createContentRoleExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("marketing_content_contributor_1357fhe"),
        PermissionType: "content",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Marketing content contributor",
        Description: optionalnullable.From(cldprovisioning.Pointer("View all assets in selected folders and contribute to those folders, without permission to download.")),
        SystemPolicyIds: []string{
            "cld::policy::content::folder::view_download",
            "cld::policy::content::folder::add_assets",
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
### Example Usage: createContentRoleResponseExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createContentRoleResponseExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: createGlobalRoleExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createGlobalRoleExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing uploads and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: createGlobalRoleResponseExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="createGlobalRoleResponseExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: invalidFieldErrorExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="invalidFieldErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: missingFieldsErrorExample

<!-- UsageSnippet language="go" operationID="createRole" method="post" path="/v2/accounts/{account_id}/permissions/roles" example="missingFieldsErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Create(ctx, components.CreateRole{
        ID: cldprovisioning.Pointer("upload_manager_12334565"),
        PermissionType: "global",
        ScopeType: components.ScopeTypeEnumProdenv,
        Name: "Upload manager",
        Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
        SystemPolicyIds: []string{
            "cld::policy::global::upload_presets::manage",
            "cld::policy::global::folder_and_asset_management::create_asset",
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.CreateRole](../../models/components/createrole.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*components.RoleResponse](../../models/components/roleresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 409                 | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Get

Get a specific system or custom role.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRole" method="get" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" example="GetRoleExample" -->
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

    res, err := s.Roles.Get(ctx, operations.GetRoleRequest{
        RoleID: "<id>",
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
| `request`                                                              | [operations.GetRoleRequest](../../models/operations/getrolerequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*components.RoleResponse](../../models/components/roleresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403, 404                      | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Update

Update a specific custom role by providing all relevant details, including those you want to keep as-is.

*Note:* Updating a role with a name that already exists will trigger a 409 error.


### Example Usage: invalidPolicyIDFieldErrorExample

<!-- UsageSnippet language="go" operationID="updateRole" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" example="invalidPolicyIDFieldErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Update(ctx, operations.UpdateRoleRequest{
        RoleID: "<id>",
        Role: components.Role{
            Name: "Upload manager",
            Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
            SystemPolicyIds: []string{
                "cld::policy::global::upload_presets::manage",
                "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: missingFieldsErrorExample

<!-- UsageSnippet language="go" operationID="updateRole" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" example="missingFieldsErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Update(ctx, operations.UpdateRoleRequest{
        RoleID: "<id>",
        Role: components.Role{
            Name: "Upload manager",
            Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
            SystemPolicyIds: []string{
                "cld::policy::global::upload_presets::manage",
                "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: updatePolicyResponse

<!-- UsageSnippet language="go" operationID="updateRole" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" example="updatePolicyResponse" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Update(ctx, operations.UpdateRoleRequest{
        RoleID: "<id>",
        Role: components.Role{
            Name: "Upload manager",
            Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing upload presets and uploading new assets.")),
            SystemPolicyIds: []string{
                "cld::policy::global::upload_presets::manage",
                "cld::policy::global::folder_and_asset_management::create_asset",
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
### Example Usage: updateRoleExample

<!-- UsageSnippet language="go" operationID="updateRole" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" example="updateRoleExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    res, err := s.Roles.Update(ctx, operations.UpdateRoleRequest{
        RoleID: "<id>",
        Role: components.Role{
            Name: "Upload manager",
            Description: optionalnullable.From(cldprovisioning.Pointer("Responsible for managing uploads and uploading new assets.")),
            SystemPolicyIds: []string{
                "cld::policy::global::upload_presets::manage",
                "cld::policy::global::folder_and_asset_management::create_asset",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateRoleRequest](../../models/operations/updaterolerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.RoleResponse](../../models/components/roleresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 404, 409            | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Delete

Delete a specific custom role.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteRole" method="delete" path="/v2/accounts/{account_id}/permissions/roles/{role_id}" -->
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

    err := s.Roles.Delete(ctx, operations.DeleteRoleRequest{
        RoleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteRoleRequest](../../models/operations/deleterolerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403, 404                      | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## ListPrincipals

Retrieve all principals associated with a specific role.

### Example Usage: GetRolePrincipalsExample

<!-- UsageSnippet language="go" operationID="getRolePrincipals" method="get" path="/v2/accounts/{account_id}/permissions/roles/{role_id}/principals" example="GetRolePrincipalsExample" -->
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

    res, err := s.Roles.ListPrincipals(ctx, operations.GetRolePrincipalsRequest{
        RoleID: "<id>",
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
### Example Usage: GetRolePrincipalsWithParamKeyValueExample

<!-- UsageSnippet language="go" operationID="getRolePrincipals" method="get" path="/v2/accounts/{account_id}/permissions/roles/{role_id}/principals" example="GetRolePrincipalsWithParamKeyValueExample" -->
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

    res, err := s.Roles.ListPrincipals(ctx, operations.GetRolePrincipalsRequest{
        RoleID: "<id>",
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
| `request`                                                                                  | [operations.GetRolePrincipalsRequest](../../models/operations/getroleprincipalsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.PrincipalsResponse](../../models/components/principalsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403, 404                      | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## UpdatePrincipals

Add or remove principals associated with a role.

### Example Usage: invalidOperationFieldErrorExample

<!-- UsageSnippet language="go" operationID="updateRolePrincipals" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}/principals" example="invalidOperationFieldErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    err := s.Roles.UpdatePrincipals(ctx, operations.UpdateRolePrincipalsRequest{
        RoleID: "<id>",
        UpdateRolePrincipalsRequest: components.UpdateRolePrincipalsRequest{
            Operation: components.OperationEnumAdd.ToPointer(),
            Principals: []components.UpdateRolePrincipal{
                components.UpdateRolePrincipal{
                    PrincipalType: components.PrincipalTypeEnumUser,
                    PrincipalID: "1234abc",
                    ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
                    PolicyParameters: &components.UpdateRolePrincipalPolicyParameters{},
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```
### Example Usage: missingFieldsErrorExample

<!-- UsageSnippet language="go" operationID="updateRolePrincipals" method="put" path="/v2/accounts/{account_id}/permissions/roles/{role_id}/principals" example="missingFieldsErrorExample" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/optionalnullable"
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

    err := s.Roles.UpdatePrincipals(ctx, operations.UpdateRolePrincipalsRequest{
        RoleID: "<id>",
        UpdateRolePrincipalsRequest: components.UpdateRolePrincipalsRequest{
            Operation: components.OperationEnumAdd.ToPointer(),
            Principals: []components.UpdateRolePrincipal{
                components.UpdateRolePrincipal{
                    PrincipalType: components.PrincipalTypeEnumUser,
                    PrincipalID: "1234abc",
                    ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
                    PolicyParameters: &components.UpdateRolePrincipalPolicyParameters{},
                },
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
| `request`                                                                                        | [operations.UpdateRolePrincipalsRequest](../../models/operations/updateroleprincipalsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 404                 | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |