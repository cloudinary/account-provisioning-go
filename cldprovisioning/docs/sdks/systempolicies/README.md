# SystemPolicies

## Overview

### Available Operations

* [List](#list) - Get system policies

## List

Retrieve all system-defined permission policies. Optionally filter the results by `scope_type` (account or product environment) and/or `permission_type` (global or content-specific).

**Note:** If `scope_type` isn't specified as `prodenv`, the account-level policies will be returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemPolicies" method="get" path="/v2/accounts/{account_id}/permissions/policies/system" example="SystemPoliciesExample" -->
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

    res, err := s.SystemPolicies.List(ctx, &operations.GetSystemPoliciesRequest{
        PermissionType: components.PermissionTypeEnumGlobal.ToPointer(),
        ScopeType: components.ScopeTypeEnumProdenv.ToPointer(),
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
| `request`                                                                                  | [operations.GetSystemPoliciesRequest](../../models/operations/getsystempoliciesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SystemPoliciesResponse](../../models/components/systempoliciesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403                           | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |