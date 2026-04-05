<!-- Start SDK Example Usage [usage] -->
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