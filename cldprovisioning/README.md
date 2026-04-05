# Cloudinary Account Provisioning Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/cloudinary/account-provisioning-go/cldprovisioning.svg)](https://pkg.go.dev/github.com/cloudinary/account-provisioning-go/cldprovisioning)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Developer-friendly & type-safe Go SDK specifically catered to leverage the *Cloudinary Account Provisioning* API.

<!-- Start Summary [summary] -->
## Summary

Cloudinary Account Provisioning API: Accounts with provisioning API access can create and manage their **product environments**, **users** and **user groups** using the RESTful Provisioning API.

Provisioning API access is available [upon request](https://cloudinary.com/contact?plan=enterprise) for accounts on an [Enterprise plan](https://cloudinary.com/pricing#pricing-enterprise).

The API uses **Basic Authentication** over HTTPS. Your **Account API Key** and **Account API Secret** (previously referred to as **Provisioning API keys**) are used for the authentication. These credentials (as well as your ACCOUNT_ID) are located in the [Cloudinary Console](https://console.cloudinary.com/pm) under **Settings > Account API Keys**.

The Provisioning API has dedicated SDKs for the following languages:

* [JavaScript](https://github.com/cloudinary/account-provisioning-js)
* [PHP](https://github.com/cloudinary/account-provisioning-php)
* [Java](https://github.com/cloudinary/account-provisioning-java)

Useful links:
* [Provisioning API reference (Classic)](https://cloudinary.com/documentation/provisioning_api_1) (includes SDKs for additional languages)

Accounts with Permissions API access can assign roles, made up of system policies, to control what principals (users, groups, and API keys) can do across the Cloudinary account and product environments. For more information about Cloudinary roles and permissions, see the [Role-based permissions](permissions_overview) guide.

Permissions API access is available [upon request](https://cloudinary.com/contact?plan=enterprise) for accounts on an [Enterprise plan](https://cloudinary.com/pricing#pricing-enterprise).

The API uses **Basic Authentication** over HTTPS. Your **Account API Key** and **Account API Secret** (previously referred to as **Provisioning API keys**) are used for the authentication. These credentials (as well as your ACCOUNT_ID) are located in the [Cloudinary Console](https://console.cloudinary.com/app/settings/account-api-keys) under **Settings > Account API Keys**.

_**Important:**_

_Cloudinary's **Roles and Permissions Management** is now available as a **Beta**. This is an early stage release, and while it's functional and ready for real-world testing, it's subject to change as we continue refining the experience based on what we learn, including your feedback.
During the Beta period, core functionality is considered stable, though some APIs, scopes, or response formats may evolve._

_**How you can help:**_

* _Use Roles and Permissions Management in real projects, prototypes, or tests._
* _Share feedback, issues, or ideas with our support team._

_Thank you for exploring this early release and helping us shape these tools to best meet your needs._
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Cloudinary Account Provisioning Go SDK](#cloudinary-account-provisioning-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/cloudinary/account-provisioning-go/cldprovisioning
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                                             | Type | Scheme      | Environment Variable                                                       |
| ------------------------------------------------ | ---- | ----------- | -------------------------------------------------------------------------- |
| `ProvisioningAPIKey`<br/>`ProvisioningAPISecret` | http | Custom HTTP | `CLOUDINARY_PROVISIONING_API_KEY`<br/>`CLOUDINARY_PROVISIONING_API_SECRET` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
		cldprovisioning.WithAccountID("<id>"),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccessKeys](docs/sdks/accesskeys/README.md)

* [List](docs/sdks/accesskeys/README.md#list) - Get access keys
* [Generate](docs/sdks/accesskeys/README.md#generate) - Generate an access key
* [DeleteByName](docs/sdks/accesskeys/README.md#deletebyname) - Delete access key by name
* [Update](docs/sdks/accesskeys/README.md#update) - Update an access key
* [Delete](docs/sdks/accesskeys/README.md#delete) - Delete access key

### [BillingUsage](docs/sdks/billingusage/README.md)

* [Get](docs/sdks/billingusage/README.md#get) - Get billing usage information

### [CustomPolicies](docs/sdks/custompolicies/README.md)

* [List](docs/sdks/custompolicies/README.md#list) - Get custom policies
* [Create](docs/sdks/custompolicies/README.md#create) - Create custom policy
* [Get](docs/sdks/custompolicies/README.md#get) - Get custom policy
* [Update](docs/sdks/custompolicies/README.md#update) - Update custom policy
* [Delete](docs/sdks/custompolicies/README.md#delete) - Delete custom policy

### [EffectivePolicies](docs/sdks/effectivepolicies/README.md)

* [List](docs/sdks/effectivepolicies/README.md#list) - Get effective policies

### [Principals](docs/sdks/principals/README.md)

* [ListRoles](docs/sdks/principals/README.md#listroles) - Get a principal's roles
* [UpdateRoles](docs/sdks/principals/README.md#updateroles) - Assign roles to a principal
* [Inspect](docs/sdks/principals/README.md#inspect) - Inspect
* [InspectMultiple](docs/sdks/principals/README.md#inspectmultiple) - Inspect multiple

### [ProductEnvironments](docs/sdks/productenvironments/README.md)

* [List](docs/sdks/productenvironments/README.md#list) - Get product environments
* [Create](docs/sdks/productenvironments/README.md#create) - Create product environment
* [Get](docs/sdks/productenvironments/README.md#get) - Get product environment
* [Update](docs/sdks/productenvironments/README.md#update) - Update product environment
* [Delete](docs/sdks/productenvironments/README.md#delete) - Delete product environment

### [Public](docs/sdks/public/README.md)

* [GetCatalog](docs/sdks/public/README.md#getcatalog) - Get system roles and policies catalog
* [ValidatePolicy](docs/sdks/public/README.md#validatepolicy) - Validate a Cedar policy
* [GetSchema](docs/sdks/public/README.md#getschema) - Get Cedar schema

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - Get roles
* [Create](docs/sdks/roles/README.md#create) - Create custom role
* [Get](docs/sdks/roles/README.md#get) - Get role
* [Update](docs/sdks/roles/README.md#update) - Update custom role
* [Delete](docs/sdks/roles/README.md#delete) - Delete custom role
* [ListPrincipals](docs/sdks/roles/README.md#listprincipals) - Get a role's principals
* [UpdatePrincipals](docs/sdks/roles/README.md#updateprincipals) - Assign principals to a role

### [SystemPolicies](docs/sdks/systempolicies/README.md)

* [List](docs/sdks/systempolicies/README.md#list) - Get system policies

### [UserGroups](docs/sdks/usergroups/README.md)

* [List](docs/sdks/usergroups/README.md#list) - Get User Groups
* [Create](docs/sdks/usergroups/README.md#create) - Create User Group
* [Get](docs/sdks/usergroups/README.md#get) - Get User Group
* [Update](docs/sdks/usergroups/README.md#update) - Update User Group
* [Delete](docs/sdks/usergroups/README.md#delete) - Delete User Group
* [ListUsers](docs/sdks/usergroups/README.md#listusers) - Get Users in User Group
* [AddUser](docs/sdks/usergroups/README.md#adduser) - Add User to User Group
* [RemoveUser](docs/sdks/usergroups/README.md#removeuser) - Remove User from User Group

### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - Get users
* [Create](docs/sdks/users/README.md#create) - Create user
* [Get](docs/sdks/users/README.md#get) - Get user
* [Update](docs/sdks/users/README.md#update) - Update user
* [Delete](docs/sdks/users/README.md#delete) - Delete user
* [GetGroups](docs/sdks/users/README.md#getgroups) - Get user groups
* [ListSubAccounts](docs/sdks/users/README.md#listsubaccounts) - Get user sub-accounts

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

A parameter is configured globally. This parameter may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `account_id` to `"<id>"` at SDK initialization and then you do not have to pass the same value on calls to operations like `List`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameter is available.
Global parameters can also be set via environment variable.

| Name      | Type   | Description | Environment           |
| --------- | ------ | ----------- | --------------------- |
| AccountID | string | Account ID  | CLOUDINARY_ACCOUNT_ID |

### Example

```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Global Parameters [global-parameters] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type              | Status Code                       | Content Type     |
| ----------------------- | --------------------------------- | ---------------- |
| sdkerrors.ErrorResponse | 400, 401, 403, 404, 409, 420, 429 | application/json |
| sdkerrors.APIError      | 4XX, 5XX                          | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                            | Variables | Description                                     |
| --- | --------------------------------- | --------- | ----------------------------------------------- |
| 0   | `https://{region}.cloudinary.com` | `region`  | Regional API endpoints for optimal performance. |
| 1   | `https://{host}`                  | `host`    | Custom domains for enterprise deployments.      |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable | Option                            | Supported Values                            | Default                | Description                 |
| -------- | --------------------------------- | ------------------------------------------- | ---------------------- | --------------------------- |
| `region` | `WithRegion(region ServerRegion)` | - `"api"`<br/>- `"api-eu"`<br/>- `"api-ap"` | `"api"`                | Regional endpoint selection |
| `host`   | `WithHost(host string)`           | string                                      | `"api.cloudinary.com"` | API host domain.            |

#### Example

```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithServerIndex(0),
		cldprovisioning.WithRegion("api-ap"),
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cldprovisioning.New(
		cldprovisioning.WithServerURL("https://api.cloudinary.com"),
		cldprovisioning.WithAccountID("<id>"),
		cldprovisioning.WithSecurity(components.Security{
			ProvisioningAPIKey:    cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_KEY"),
			ProvisioningAPISecret: cldprovisioning.Pointer("CLOUDINARY_PROVISIONING_API_SECRET"),
		}),
	)

	res, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{
		Enabled: cldprovisioning.Pointer(true),
		Prefix:  cldprovisioning.Pointer("product"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = cldprovisioning.New(cldprovisioning.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 


