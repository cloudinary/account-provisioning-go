# Users

## Overview

Manage the users for your account.


Learn more about users management.
<https://cloudinary.com/documentation/provisioning_api#users>

### Available Operations

* [List](#list) - Get users
* [Create](#create) - Create user
* [Get](#get) - Get user
* [Update](#update) - Update user
* [Delete](#delete) - Delete user
* [GetGroups](#getgroups) - Get user groups
* [ListSubAccounts](#listsubaccounts) - Get user sub-accounts

## List

Returns an array of all users in the account, or if conditions are specified, returns the relevant users.


Learn more about retrieving users.
<https://cloudinary.com/documentation/provisioning_api#get_users>

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsers" method="get" path="/v1_1/provisioning/accounts/{account_id}/users" example="Users" -->
```go
package main

import(
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/types"
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

    res, err := s.Users.List(ctx, &operations.GetUsersRequest{
        Pending: cldprovisioning.Pointer(false),
        Prefix: cldprovisioning.Pointer("john"),
        From: types.MustNewDateFromString("2023-01-01T00:00:00Z"),
        To: types.MustNewDateFromString("2024-12-31T00:00:00Z"),
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
| `request`                                                                | [operations.GetUsersRequest](../../models/operations/getusersrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*components.UsersResponse](../../models/components/usersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Create

Create a new user.

Learn more about creating users.
<https://cloudinary.com/documentation/provisioning_api#create_user>

### Example Usage: CreateUserExample

<!-- UsageSnippet language="go" operationID="createUser" method="post" path="/v1_1/provisioning/accounts/{account_id}/users" example="CreateUserExample" -->
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

    res, err := s.Users.Create(ctx, operations.CreateUserRequestBody{
        Name: "jenny_bar",
        Email: "jenny_bar@example.com",
        Role: operations.RoleMasterAdmin.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: UserResponseExample

<!-- UsageSnippet language="go" operationID="createUser" method="post" path="/v1_1/provisioning/accounts/{account_id}/users" example="UserResponseExample" -->
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

    res, err := s.Users.Create(ctx, operations.CreateUserRequestBody{
        Name: "John",
        Email: "john@example.com",
        Role: operations.RoleTechnicalAdmin.ToPointer(),
        Enabled: cldprovisioning.Pointer(true),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateUserRequestBody](../../models/operations/createuserrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*components.User](../../models/components/user.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Get

Retrieve a specific user.

Learn more about retrieving a specific user.
<https://cloudinary.com/documentation/provisioning_api#get_user>

### Example Usage

<!-- UsageSnippet language="go" operationID="getUser" method="get" path="/v1_1/provisioning/accounts/{account_id}/users/{user_id}" example="UserResponseExample" -->
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

    res, err := s.Users.Get(ctx, "0abed8dfcc039ea05e2a1d494fd442")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the user.                                      | 0abed8dfcc039ea05e2a1d494fd442                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.User](../../models/components/user.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update the details of a user.

Learn more about updating a specific user.
<https://cloudinary.com/documentation/provisioning_api#update_user>

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUser" method="put" path="/v1_1/provisioning/accounts/{account_id}/users/{user_id}" example="UserJohn" -->
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

    res, err := s.Users.Update(ctx, operations.UpdateUserRequest{
        UserID: "0abed8dfcc039ea05e2a1d494fd442",
        UserRequest: components.UserRequest{
            Name: cldprovisioning.Pointer("John"),
            Email: cldprovisioning.Pointer("john@example.com"),
            Role: components.UserRequestRoleTechnicalAdmin.ToPointer(),
            Enabled: cldprovisioning.Pointer(true),
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
| `request`                                                                    | [operations.UpdateUserRequest](../../models/operations/updateuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.User](../../models/components/user.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete a specific user.

Learn more about deleting a specific user.
<https://cloudinary.com/documentation/provisioning_api#delete_user>

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUser" method="delete" path="/v1_1/provisioning/accounts/{account_id}/users/{user_id}" -->
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

    res, err := s.Users.Delete(ctx, "0abed8dfcc039ea05e2a1d494fd442")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the user.                                      | 0abed8dfcc039ea05e2a1d494fd442                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.SuccessResponse](../../models/components/successresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 420, 429 | application/json             |
| sdkerrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## GetGroups

Retrieve the groups that a specific user belongs to.

Learn more about retrieving user groups.
<https://cloudinary.com/documentation/provisioning_api#get_users_groups>

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsForUser" method="get" path="/v1_1/provisioning/accounts/{account_id}/users/{user_id}/groups" example="UserGroupsList" -->
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

    res, err := s.Users.GetGroups(ctx, "0abed8dfcc039ea05e2a1d494fd442")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the user.                                      | 0abed8dfcc039ea05e2a1d494fd442                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.UserGroupsListResponse](../../models/components/usergroupslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListSubAccounts

Retrieve the product environments (sub-accounts) that a specific user has access to.

Learn more about retrieving user sub-accounts.
<https://cloudinary.com/documentation/provisioning_api#get_user_sub_accounts>

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubAccountsForUser" method="get" path="/v1_1/provisioning/accounts/{account_id}/users/{user_id}/sub_accounts" example="UserSubAccounts" -->
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

    res, err := s.Users.ListSubAccounts(ctx, "0abed8dfcc039ea05e2a1d494fd442")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the user.                                      | 0abed8dfcc039ea05e2a1d494fd442                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.UserSubAccountsResponse](../../models/components/usersubaccountsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |