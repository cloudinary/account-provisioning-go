## Go SDK Changes:
* `CldProvisioning.Roles.Create()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Users.Update()`: `request.Request` **Changed** (Breaking ⚠️)
    - `RequestBody` **Added** (Breaking ⚠️)
    - `UserRequest` **Removed** (Breaking ⚠️)
* `CldProvisioning.SystemPolicies.List()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.CustomPolicies.List()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.CustomPolicies.Create()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.CustomPolicies.Get()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.CustomPolicies.Update()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.CustomPolicies.Delete()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.List()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.EffectivePolicies.List()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.Update()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Role` **Removed** (Breaking ⚠️)
    - `UpdateRole` **Added** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.Get()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Principals.InspectMultiple()`: 
  *  `request.Request.PolicyParameters` **Changed** (Breaking ⚠️)
  *  `response.Data.Principals[].Roles[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.Delete()`:  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.ListPrincipals()`: 
  *  `response.Data[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Roles.UpdatePrincipals()`: 
  *  `request.Request.UpdateRolePrincipalsRequest.Principals[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Principals.ListRoles()`: 
  *  `response.Data[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Principals.UpdateRoles()`: 
  *  `request.Request.Roles[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Principals.Inspect()`: 
  *  `response.Data.Principals[].Roles[].PolicyParameters` **Changed** (Breaking ⚠️)
  *  `error.Error.Details` **Changed** (Breaking ⚠️)
* `CldProvisioning.Users.Create()`: `request.Request` **Changed**
    - `Role` **Deprecated**
    - `Roles` **Added**
