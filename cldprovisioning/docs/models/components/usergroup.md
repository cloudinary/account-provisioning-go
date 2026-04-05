# UserGroup

User group.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ID`                                                                               | `*string`                                                                          | :heavy_minus_sign:                                                                 | The user group's ID.                                                               |
| `Name`                                                                             | `*string`                                                                          | :heavy_minus_sign:                                                                 | The user group's name.                                                             |
| `Users`                                                                            | [][components.UserGroupUser](../../models/components/usergroupuser.md)             | :heavy_minus_sign:                                                                 | A list of user group users.<br/>Only returned in extended_details.<br/>            |
| `CreatedAt`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_minus_sign:                                                                 | The date when the user group was created.<br/>Only returned in extended_details.<br/> |
| `UpdatedAt`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_minus_sign:                                                                 | The date when the user group was last updated.<br/>Only returned in extended_details.<br/> |