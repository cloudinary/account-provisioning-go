# Cloudinary Account Provisioning Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/cloudinary/account-provisioning-go/cldprovisioning.svg)](https://pkg.go.dev/github.com/cloudinary/account-provisioning-go/cldprovisioning)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go SDK for the [Cloudinary Account Provisioning](https://cloudinary.com/documentation/provisioning_api_1) and Permissions APIs.

## Quick Start

```bash
go get github.com/cloudinary/account-provisioning-go/cldprovisioning
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
)

func main() {
	// Credentials loaded automatically from environment variables:
	// CLOUDINARY_ACCOUNT_URL or CLOUDINARY_ACCOUNT_ID + CLOUDINARY_PROVISIONING_API_KEY + CLOUDINARY_PROVISIONING_API_SECRET
	s := cldprovisioning.New()

	res, err := s.ProductEnvironments.List(context.Background(),
		&operations.GetProductEnvironmentsRequest{
			Enabled: cldprovisioning.Bool(true),
			Prefix:  cldprovisioning.String("my"),
		})
	if err != nil {
		log.Fatal(err)
	}
	for _, pe := range res.SubAccounts {
		fmt.Printf("%s  %s\n", *pe.ID, *pe.Name)
	}
}
```

## Documentation

See the full SDK documentation in the [`cldprovisioning`](cldprovisioning/README.md) module.
