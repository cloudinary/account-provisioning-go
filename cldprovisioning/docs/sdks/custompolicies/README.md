# CustomPolicies

## Overview

### Available Operations

* [List](#list) - Get custom policies
* [Create](#create) - Create custom policy
* [Get](#get) - Get custom policy
* [Update](#update) - Update custom policy
* [Delete](#delete) - Delete custom policy

## List

Retrieve all custom permission policies defined for a specific scope. The scope can be at the account level or within a specific product environment.

**Notes:**
  * If `scope_type` isn't specified as `prodenv`, the account-level policies will be returned.

  * If `scope_type` is specified as `prodenv`, `scope_id` is required.


### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomPolicies" method="get" path="/v2/accounts/{account_id}/permissions/policies/custom" example="CustomPoliciesExample" -->
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

    res, err := s.CustomPolicies.List(ctx, &operations.GetCustomPoliciesRequest{
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
| `request`                                                                                  | [operations.GetCustomPoliciesRequest](../../models/operations/getcustompoliciesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.CustomPoliciesResponse](../../models/components/custompoliciesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403                           | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a new custom policy with a statement defined in Cedar. This policy specifies permissions for a principal, allowing a specific action on a designated resource, within a particular scope.

Learn more about creating a new custom policy
<https://cloudinary.com/documentation/permissions_api_guide#create_a_policy_statement>

### Example Usage: cedarErrorExample

<!-- UsageSnippet language="go" operationID="createCustomPolicy" method="post" path="/v2/accounts/{account_id}/permissions/policies/custom" example="cedarErrorExample" -->
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

    res, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
        PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
        Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
        Name: "Read access to Clothing folder",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: createCustomPolicyExample

<!-- UsageSnippet language="go" operationID="createCustomPolicy" method="post" path="/v2/accounts/{account_id}/permissions/policies/custom" example="createCustomPolicyExample" -->
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

    res, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
        PolicyStatement: "permit(principal == Cloudinary::APIKey::\"1234\",action==Cloudinary::Action::\"read\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\"asdfjkl12347890\")} ;",
        Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
        Name: "Read access to Clothing folder",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: createPolicyResponse

<!-- UsageSnippet language="go" operationID="createCustomPolicy" method="post" path="/v2/accounts/{account_id}/permissions/policies/custom" example="createPolicyResponse" -->
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

    res, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
        PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
        Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
        Name: "Read access to Clothing folder",
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

<!-- UsageSnippet language="go" operationID="createCustomPolicy" method="post" path="/v2/accounts/{account_id}/permissions/policies/custom" example="invalidFieldErrorExample" -->
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

    res, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
        PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
        Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
        Name: "Read access to Clothing folder",
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

<!-- UsageSnippet language="go" operationID="createCustomPolicy" method="post" path="/v2/accounts/{account_id}/permissions/policies/custom" example="missingFieldsErrorExample" -->
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

    res, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
        PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
        Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
        ScopeType: components.ScopeTypeEnumProdenv,
        ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
        Name: "Read access to Clothing folder",
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.CustomPolicy](../../models/components/custompolicy.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*components.CustomPolicyResponse](../../models/components/custompolicyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 409                 | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Get

Get a specific custom policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomPolicy" method="get" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="CustomPolicyExample" -->
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

    res, err := s.CustomPolicies.Get(ctx, "<id>")
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
| `policyID`                                               | `string`                                                 | :heavy_check_mark:                                       | A unique identifier for the custom permission policy.    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.CustomPolicyResponse](../../models/components/custompolicyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403, 404                      | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Update

Update a specific custom policy by providing the entire policy entity in the request body. Ensure the new `policy_statement` includes modifications. Existing permissions will be replaced with the new data.

*Note:* Updating a policy with an identical `policy_statement` will trigger a 409 error.


### Example Usage: cedarErrorExample

<!-- UsageSnippet language="go" operationID="updateCustomPolicy" method="put" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="cedarErrorExample" -->
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

    res, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
        PolicyID: "<id>",
        UpdateCustomPolicy: components.UpdateCustomPolicy{
            PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
            Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
            ScopeType: components.ScopeTypeEnumProdenv,
            ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
            Name: "Read access to Clothing folder",
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

<!-- UsageSnippet language="go" operationID="updateCustomPolicy" method="put" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="invalidFieldErrorExample" -->
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

    res, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
        PolicyID: "<id>",
        UpdateCustomPolicy: components.UpdateCustomPolicy{
            PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
            Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
            ScopeType: components.ScopeTypeEnumProdenv,
            ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
            Name: "Read access to Clothing folder",
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
### Example Usage: missingFieldsUpdateErrorExample

<!-- UsageSnippet language="go" operationID="updateCustomPolicy" method="put" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="missingFieldsUpdateErrorExample" -->
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

    res, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
        PolicyID: "<id>",
        UpdateCustomPolicy: components.UpdateCustomPolicy{
            PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
            Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
            ScopeType: components.ScopeTypeEnumProdenv,
            ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
            Name: "Read access to Clothing folder",
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
### Example Usage: updateCustomPolicyExample

<!-- UsageSnippet language="go" operationID="updateCustomPolicy" method="put" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="updateCustomPolicyExample" -->
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

    res, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
        PolicyID: "<id>",
        UpdateCustomPolicy: components.UpdateCustomPolicy{
            PolicyStatement: "permit(principal == Cloudinary::APIKey::\"1234\",action==Cloudinary::Action::\"read\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\"asdfjkl12347890\")} ;",
            Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Accessories folder with external ID asdfjkl12347890.")),
            ScopeType: components.ScopeTypeEnumProdenv,
            ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
            Name: "Read access to Accessories folder",
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

<!-- UsageSnippet language="go" operationID="updateCustomPolicy" method="put" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" example="updatePolicyResponse" -->
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

    res, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
        PolicyID: "<id>",
        UpdateCustomPolicy: components.UpdateCustomPolicy{
            PolicyStatement: "permit(principal == Cloudinary::APIKey::\\\"1234\\\",action==Cloudinary::Action::\\\"read\\\",resource is Cloudinary::Folder ) when {resource.ancestor_ids.contains(\\\"asdfjkl12347890\\\")} ;\n",
            Description: optionalnullable.From(cldprovisioning.Pointer("Permit read access to the Clothing folder with external ID asdfjkl12347890.")),
            ScopeType: components.ScopeTypeEnumProdenv,
            ScopeID: optionalnullable.From(cldprovisioning.Pointer("975l29lz02jt0836fhwi")),
            Name: "Read access to Clothing folder",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCustomPolicyRequest](../../models/operations/updatecustompolicyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.CustomPolicyResponse](../../models/components/custompolicyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 400, 401, 403, 404, 409            | application/json                   |
| sdkerrors.PermissionsErrorResponse | 500                                | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Delete

Delete a specific custom policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCustomPolicy" method="delete" path="/v2/accounts/{account_id}/permissions/policies/custom/{policy_id}" -->
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

    err := s.CustomPolicies.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policyID`                                               | `string`                                                 | :heavy_check_mark:                                       | A unique identifier for the custom permission policy.    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.PermissionsErrorResponse | 401, 403, 404                      | application/json                   |
| sdkerrors.APIError                 | 4XX, 5XX                           | \*/\*                              |