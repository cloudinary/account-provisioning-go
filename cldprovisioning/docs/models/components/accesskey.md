# AccessKey


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Name`                                         | `*string`                                      | :heavy_minus_sign:                             | The name of the access key.                    |
| `APIKey`                                       | `*string`                                      | :heavy_minus_sign:                             | The API key.                                   |
| `APISecret`                                    | `*string`                                      | :heavy_minus_sign:                             | The API secret.                                |
| `CreatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The date when the access key was created.      |
| `UpdatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The date when the access key was last updated. |
| `Enabled`                                      | `*bool`                                        | :heavy_minus_sign:                             | Whether the access key is enabled or disabled. |