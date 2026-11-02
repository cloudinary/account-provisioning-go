# RoleManagementRequest

A role management operation and the roles it applies to.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Operation`                                                              | [components.Operation](../../models/components/operation.md)             | :heavy_check_mark:                                                       | Whether to add or remove the listed roles.                               | add                                                                      |
| `Roles`                                                                  | [][components.RoleAssignment](../../models/components/roleassignment.md) | :heavy_check_mark:                                                       | The roles to add or remove.                                              |                                                                          |