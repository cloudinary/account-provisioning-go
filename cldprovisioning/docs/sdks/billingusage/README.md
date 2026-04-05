# BillingUsage

## Overview

### Available Operations

* [Get](#get) - Get billing usage information

## Get

Retrieve billing usage metrics for your account and product environments,
including storage, bandwidth, impressions, and transformations.

If `year` is provided:
  - **Annual subscription**: Returns the single annual usage cycle that began in the specified year (if any).
  - **Monthly subscription**: Returns all monthly usage cycles that began in the specified year.

If `year` is not provided:
  - Returns usage for the year in which the current active subscription period began.
  - If no active subscription exists, an empty list is returned.

**Note**: Usage metrics are updated daily, so calling this endpoint once per day gives you the latest data.


### Example Usage: AdvancedExtraPAYG

<!-- UsageSnippet language="go" operationID="getBillingUsage" method="get" path="/v1_1/provisioning/accounts/{account_id}/billing_usage" example="AdvancedExtraPAYG" -->
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

    res, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{
        Year: cldprovisioning.Pointer[int](2024),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: AdvancedPAYG

<!-- UsageSnippet language="go" operationID="getBillingUsage" method="get" path="/v1_1/provisioning/accounts/{account_id}/billing_usage" example="AdvancedPAYG" -->
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

    res, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{
        Year: cldprovisioning.Pointer[int](2024),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: MultiPlan

<!-- UsageSnippet language="go" operationID="getBillingUsage" method="get" path="/v1_1/provisioning/accounts/{account_id}/billing_usage" example="MultiPlan" -->
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

    res, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{
        Year: cldprovisioning.Pointer[int](2024),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: ProPAYG

<!-- UsageSnippet language="go" operationID="getBillingUsage" method="get" path="/v1_1/provisioning/accounts/{account_id}/billing_usage" example="ProPAYG" -->
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

    res, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{
        Year: cldprovisioning.Pointer[int](2024),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: ProPAYGPlus

<!-- UsageSnippet language="go" operationID="getBillingUsage" method="get" path="/v1_1/provisioning/accounts/{account_id}/billing_usage" example="ProPAYGPlus" -->
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

    res, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{
        Year: cldprovisioning.Pointer[int](2024),
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
| `request`                                                                              | [operations.GetBillingUsageRequest](../../models/operations/getbillingusagerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[[]components.BillingUsageResponse](../../.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400, 401, 403, 404, 429 | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |