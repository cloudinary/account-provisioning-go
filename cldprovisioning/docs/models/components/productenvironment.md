# ProductEnvironment

Product Environment


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CloudName`                                                          | `*string`                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Name`                                                               | `*string`                                                            | :heavy_minus_sign:                                                   | The cloud name of the product environment.                           |
| `Enabled`                                                            | `*bool`                                                              | :heavy_minus_sign:                                                   | Whether the product environment is enabled or disabled.              |
| `ID`                                                                 | `*string`                                                            | :heavy_minus_sign:                                                   | The ID of the product environment.                                   |
| `APIAccessKeys`                                                      | [][components.APIAccessKey](../../models/components/apiaccesskey.md) | :heavy_minus_sign:                                                   | The list of the product environment's access keys.                   |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | The date when the product environment was created.                   |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | The date when the product environment was last updated.              |