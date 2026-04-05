# UserGroups

## Overview

### Available Operations

* [List](#list) - Get User Groups
* [Create](#create) - Create User Group
* [Get](#get) - Get User Group
* [Update](#update) - Update User Group
* [Delete](#delete) - Delete User Group
* [ListUsers](#listusers) - Get Users in User Group
* [AddUser](#adduser) - Add User to User Group
* [RemoveUser](#removeuser) - Remove User from User Group

## List

Retrieve an array of all user groups in the account.

Learn more about retrieving user groups.
<https://cloudinary.com/documentation/provisioning_api#get_user_groups>

### Example Usage: UserGroups

<!-- UsageSnippet language="go" operationID="getUserGroups" method="get" path="/v1_1/provisioning/accounts/{account_id}/user_groups" example="UserGroups" -->
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

    res, err := s.UserGroups.List(ctx, cldprovisioning.Pointer(true))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: UserGroupsExtendedDetails

<!-- UsageSnippet language="go" operationID="getUserGroups" method="get" path="/v1_1/provisioning/accounts/{account_id}/user_groups" example="UserGroupsExtendedDetails" -->
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

    res, err := s.UserGroups.List(ctx, cldprovisioning.Pointer(true))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `extendedDetails`                                                                                         | `*bool`                                                                                                   | :heavy_minus_sign:                                                                                        | Whether to only return extended (true) or basic information about the group (false).<br/>**Default**: false.<br/> | true                                                                                                      |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*components.UserGroupsResponse](../../models/components/usergroupsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Create

Create a new user group for the account.

Learn more about creating user groups.
<https://cloudinary.com/documentation/provisioning_api#create_user_group>

### Example Usage

<!-- UsageSnippet language="go" operationID="createUserGroup" method="post" path="/v1_1/provisioning/accounts/{account_id}/user_groups" example="UserGroupDesigners" -->
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

    res, err := s.UserGroups.Create(ctx, components.UserGroupRequest{
        Name: cldprovisioning.Pointer("Designers"),
        ExtendedDetails: cldprovisioning.Pointer(false),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.UserGroupRequest](../../models/components/usergrouprequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*components.UserGroup](../../models/components/usergroup.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Get

Retrieve a specific user group.

Learn more about retrieving a specific user group.
<https://cloudinary.com/documentation/provisioning_api#get_user_group>

### Example Usage: UserGroupDesigners

<!-- UsageSnippet language="go" operationID="getUserGroup" method="get" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}" example="UserGroupDesigners" -->
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

    res, err := s.UserGroups.Get(ctx, operations.GetUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        ExtendedDetails: cldprovisioning.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: UserGroupDesignersExtendedDetails

<!-- UsageSnippet language="go" operationID="getUserGroup" method="get" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}" example="UserGroupDesignersExtendedDetails" -->
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

    res, err := s.UserGroups.Get(ctx, operations.GetUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        ExtendedDetails: cldprovisioning.Pointer(true),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetUserGroupRequest](../../models/operations/getusergrouprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*components.UserGroup](../../models/components/usergroup.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update the name of a specified user group.

Learn more about updating a specific user group.
<https://cloudinary.com/documentation/provisioning_api#update_user_group>

### Example Usage: UserGroupDesignersRequest

<!-- UsageSnippet language="go" operationID="updateUserGroup" method="put" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}" example="UserGroupDesignersRequest" -->
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

    res, err := s.UserGroups.Update(ctx, operations.UpdateUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        UserGroupRequest: components.UserGroupRequest{
            Name: cldprovisioning.Pointer("Designers and Creatives"),
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
### Example Usage: UserGroupUpdate

<!-- UsageSnippet language="go" operationID="updateUserGroup" method="put" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}" example="UserGroupUpdate" -->
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

    res, err := s.UserGroups.Update(ctx, operations.UpdateUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        UserGroupRequest: components.UserGroupRequest{
            Name: cldprovisioning.Pointer("Designers"),
            ExtendedDetails: cldprovisioning.Pointer(false),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateUserGroupRequest](../../models/operations/updateusergrouprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.UserGroup](../../models/components/usergroup.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete a user group with the specified ID.

Learn more about deleting a specific user group.
<https://cloudinary.com/documentation/provisioning_api#delete_user_group>

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUserGroup" method="delete" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}" -->
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

    res, err := s.UserGroups.Delete(ctx, "7f08f1f1fc910bf1f25274aef11d27")
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
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the user group.                                | 7f08f1f1fc910bf1f25274aef11d27                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.SuccessResponse](../../models/components/successresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 420, 429 | application/json             |
| sdkerrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## ListUsers

Retrieve the users in the group with the specified ID.

Learn more about retrieving users in a specific user group.
<https://cloudinary.com/documentation/provisioning_api#get_user_group_users>

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsersInUserGroup" method="get" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}/users" example="UserGroupUsers" -->
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

    res, err := s.UserGroups.ListUsers(ctx, "7f08f1f1fc910bf1f25274aef11d27")
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
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the user group.                                | 7f08f1f1fc910bf1f25274aef11d27                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.UserGroupUsersResponse](../../models/components/usergroupusersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 401, 404, 420, 429      | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AddUser

Add a user to a group with the specified ID.

Learn more about adding users to a specific user group.
<https://cloudinary.com/documentation/provisioning_api#add_user_to_group>

### Example Usage

<!-- UsageSnippet language="go" operationID="addUserToUserGroup" method="post" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}/users/{user_id}" example="UserGroupUserJohn" -->
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

    res, err := s.UserGroups.AddUser(ctx, operations.AddUserToUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        UserID: "0abed8dfcc039ea05e2a1d494fd442",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AddUserToUserGroupRequest](../../models/operations/addusertousergrouprequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.UserGroupUser](../../models/components/usergroupuser.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## RemoveUser

Remove a user from a group with the specified ID.

Learn more about removing users from a specific user group.
<https://cloudinary.com/documentation/provisioning_api#remove_user_from_group>

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUserFromUserGroup" method="delete" path="/v1_1/provisioning/accounts/{account_id}/user_groups/{group_id}/users/{user_id}" example="UserGroupUsers" -->
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

    res, err := s.UserGroups.RemoveUser(ctx, operations.RemoveUserFromUserGroupRequest{
        GroupID: "7f08f1f1fc910bf1f25274aef11d27",
        UserID: "0abed8dfcc039ea05e2a1d494fd442",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveUserFromUserGroupRequest](../../models/operations/removeuserfromusergrouprequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*components.UserGroupUsersResponse](../../models/components/usergroupusersresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 420, 429 | application/json             |
| sdkerrors.APIError           | 4XX, 5XX                     | \*/\*                        |