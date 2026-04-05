# EffectivePolicies

## Overview

### Available Operations

* [List](#list) - Get effective policies

## List

Retrieve all roles (along with their system permission policies) associated with principals, and all custom permission policies directly added to principals, all within a specified scope.

Optionally specify a principal to only include roles and policies associated with that principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="getEffectivePolicies" method="get" path="/v2/accounts/{account_id}/permissions/policies/effective" example="EffectivePoliciesExample" -->
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

    res, err := s.EffectivePolicies.List(ctx, &operations.GetEffectivePoliciesRequest{
        ScopeType: components.ScopeTypeEnumProdenv.ToPointer(),
        PrincipalType: components.PrincipalTypeEnumUser.ToPointer(),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetEffectivePoliciesRequest](../../models/operations/geteffectivepoliciesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*components.EffectivePoliciesResponse](../../models/components/effectivepoliciesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 404                 | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |