# ProductEnvironments

## Overview

### Available Operations

* [List](#list) - Get product environments
* [Create](#create) - Create product environment
* [Get](#get) - Get product environment
* [Update](#update) - Update product environment
* [Delete](#delete) - Delete product environment

## List

Return an array of all product environments, or if conditions are specified, 
return the relevant product environments.


Learn more about getting product environments.
<https://cloudinary.com/documentation/provisioning_api#get_product_environments>

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductEnvironments" method="get" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts" example="ProductEnvironments" -->
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

    res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
        Enabled: cldprovisioning.Pointer(true),
        Prefix: cldprovisioning.Pointer("product"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetProductEnvironmentsRequest](../../models/operations/getproductenvironmentsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*components.ProductEnvironmentsResponse](../../models/components/productenvironmentsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Create

Create a new product environment. Any users that have access to all product environments will also automatically
have access to the new product environment.


Learn more about creating product environments.
<https://cloudinary.com/documentation/provisioning_api#create_product_environment>

### Example Usage

<!-- UsageSnippet language="go" operationID="createProductEnvironment" method="post" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts" example="ProductEnvironment3" -->
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

    res, err := s.ProductEnvironments.Create(ctx, components.ProductEnvironmentRequest{
        Name: "Product3 Application",
        CloudName: cldprovisioning.Pointer("product3"),
        BaseSubAccountID: cldprovisioning.Pointer("0aaaaa1bbbbb2ccccc3ddddd4eeeee5f"),
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
| `request`                                                                                    | [components.ProductEnvironmentRequest](../../models/components/productenvironmentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.ProductEnvironment](../../models/components/productenvironment.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 409, 420, 429 | application/json             |
| sdkerrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Get

Retrieve a specific product environment.

Learn more about retrieving specific product environments.
<https://cloudinary.com/documentation/provisioning_api#get_product_environment>

### Example Usage: ProductEnvironment1

<!-- UsageSnippet language="go" operationID="getProductEnvironment" method="get" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts/{sub_account_id}" example="ProductEnvironment1" -->
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

    res, err := s.ProductEnvironments.Get(ctx, "abcde1fghij2klmno3pqrst4uvwxy5z")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: ProductEnvironment2

<!-- UsageSnippet language="go" operationID="getProductEnvironment" method="get" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts/{sub_account_id}" example="ProductEnvironment2" -->
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

    res, err := s.ProductEnvironments.Get(ctx, "abcde1fghij2klmno3pqrst4uvwxy5z")
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
| `subAccountID`                                           | `string`                                                 | :heavy_check_mark:                                       | The ID of the product environment.                       | abcde1fghij2klmno3pqrst4uvwxy5z                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.ProductEnvironment](../../models/components/productenvironment.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400, 401, 404, 420, 429 | application/json        |
| sdkerrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update the details of a product environment.

Learn more about updating product environments.
<https://cloudinary.com/documentation/provisioning_api#update_product_environment>

### Example Usage

<!-- UsageSnippet language="go" operationID="updateProductEnvironment" method="put" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts/{sub_account_id}" example="ProductEnvironment1" -->
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

    res, err := s.ProductEnvironments.Update(ctx, operations.UpdateProductEnvironmentRequest{
        SubAccountID: "abcde1fghij2klmno3pqrst4uvwxy5z",
        ProductEnvironmentUpdateRequest: components.ProductEnvironmentUpdateRequest{
            Name: cldprovisioning.Pointer("Product3 Application"),
            CloudName: cldprovisioning.Pointer("product3"),
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateProductEnvironmentRequest](../../models/operations/updateproductenvironmentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*components.ProductEnvironment](../../models/components/productenvironment.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorResponse           | 400, 401, 403, 404, 409, 420, 429 | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete a specific product environment.

Learn more about deleting product environments.
<https://cloudinary.com/documentation/provisioning_api#delete_product_environment>

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteProductEnvironment" method="delete" path="/v1_1/provisioning/accounts/{account_id}/sub_accounts/{sub_account_id}" -->
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

    res, err := s.ProductEnvironments.Delete(ctx, "abcde1fghij2klmno3pqrst4uvwxy5z")
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
| `subAccountID`                                           | `string`                                                 | :heavy_check_mark:                                       | The ID of the product environment.                       | abcde1fghij2klmno3pqrst4uvwxy5z                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.SuccessResponse](../../models/components/successresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 420, 429 | application/json             |
| sdkerrors.APIError           | 4XX, 5XX                     | \*/\*                        |