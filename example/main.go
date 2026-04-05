package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/account-provisioning-go/cldprovisioning"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/components"
	"github.com/cloudinary/account-provisioning-go/cldprovisioning/models/operations"
)

type results struct {
	passed int
	failed int
	errors []string
}

func (r *results) pass(name string) {
	r.passed++
	fmt.Printf("  PASS: %s\n", name)
}

func (r *results) fail(name string, err error) {
	r.failed++
	msg := fmt.Sprintf("FAIL: %s — %v", name, err)
	r.errors = append(r.errors, msg)
	fmt.Printf("  %s\n", msg)
}

func main() {
	ctx := context.Background()
	s := cldprovisioning.New()
	r := &results{}

	fmt.Println("============================================================")
	fmt.Println("  Cloudinary Account Provisioning Go SDK — Full Endpoint Test")
	fmt.Println("============================================================")

	var firstProdEnvID, firstUserID, firstGroupID, firstRoleID string

	// =====================================================================
	// PUBLIC
	// =====================================================================
	section("Public")

	catalogRes, err := s.Public.GetCatalog(ctx)
	if err != nil {
		r.fail("Public.GetCatalog", err)
	} else {
		count := 0
		for _, v := range catalogRes.Catalog {
			if items, ok := v.([]any); ok {
				count += len(items)
			}
		}
		fmt.Printf("    catalog keys=%v, total items=%d\n", mapKeys(catalogRes.Catalog), count)
		r.pass("Public.GetCatalog")
	}

	schemaRes, err := s.Public.GetSchema(ctx)
	if err != nil {
		r.fail("Public.GetSchema", err)
	} else {
		data, _ := json.Marshal(schemaRes.Schema)
		fmt.Printf("    schema json_len=%d\n", len(data))
		r.pass("Public.GetSchema")
	}

	validateRes, err := s.Public.ValidatePolicy(ctx, components.ValidateCedarPolicyRequest{
		Policy: `permit(principal, action, resource);`,
	})
	if err != nil {
		r.fail("Public.ValidatePolicy", err)
	} else {
		fmt.Printf("    valid=%v, errors=%d\n", validateRes.Valid, len(validateRes.Errors))
		r.pass("Public.ValidatePolicy")
	}

	// =====================================================================
	// PRODUCT ENVIRONMENTS
	// =====================================================================
	section("ProductEnvironments")

	peList, err := s.ProductEnvironments.List(ctx, &operations.GetProductEnvironmentsRequest{})
	if err != nil {
		r.fail("ProductEnvironments.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(peList.SubAccounts))
		if len(peList.SubAccounts) > 0 {
			firstProdEnvID = deref(peList.SubAccounts[0].ID)
		}
		r.pass("ProductEnvironments.List")
	}

	if firstProdEnvID != "" {
		peGet, err := s.ProductEnvironments.Get(ctx, firstProdEnvID)
		if err != nil {
			r.fail("ProductEnvironments.Get", err)
		} else {
			fmt.Printf("    id=%s name=%s cloud=%s\n", deref(peGet.ID), deref(peGet.Name), deref(peGet.CloudName))
			r.pass("ProductEnvironments.Get")
		}
	}

	testPEName := fmt.Sprintf("go-sdk-test-%d", time.Now().Unix())
	peCreated, err := s.ProductEnvironments.Create(ctx, components.ProductEnvironmentRequest{Name: testPEName})
	if err != nil {
		r.fail("ProductEnvironments.Create", err)
	} else {
		createdID := deref(peCreated.ID)
		fmt.Printf("    created id=%s name=%s\n", createdID, deref(peCreated.Name))
		r.pass("ProductEnvironments.Create")

		updatedName := testPEName + "-updated"
		peUpdated, err := s.ProductEnvironments.Update(ctx, operations.UpdateProductEnvironmentRequest{
			SubAccountID:                       createdID,
			ProductEnvironmentUpdateRequest: components.ProductEnvironmentUpdateRequest{Name: &updatedName},
		})
		if err != nil {
			r.fail("ProductEnvironments.Update", err)
		} else {
			fmt.Printf("    updated name=%s\n", deref(peUpdated.Name))
			r.pass("ProductEnvironments.Update")
		}

		_, err = s.ProductEnvironments.Delete(ctx, createdID)
		if err != nil {
			r.fail("ProductEnvironments.Delete", err)
		} else {
			fmt.Printf("    deleted id=%s\n", createdID)
			r.pass("ProductEnvironments.Delete")
		}
	}

	// =====================================================================
	// USERS
	// =====================================================================
	section("Users")

	userList, err := s.Users.List(ctx, &operations.GetUsersRequest{})
	if err != nil {
		r.fail("Users.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(userList.Users))
		if len(userList.Users) > 0 {
			firstUserID = deref(userList.Users[0].ID)
		}
		r.pass("Users.List")
	}

	if firstUserID != "" {
		userGet, err := s.Users.Get(ctx, firstUserID)
		if err != nil {
			r.fail("Users.Get", err)
		} else {
			fmt.Printf("    id=%s name=%s email=%s\n", deref(userGet.ID), deref(userGet.Name), deref(userGet.Email))
			r.pass("Users.Get")
		}
	}

	if firstUserID != "" {
		userGroups, err := s.Users.GetGroups(ctx, firstUserID)
		if err != nil {
			r.fail("Users.GetGroups", err)
		} else {
			fmt.Printf("    groups=%d\n", len(userGroups.Groups))
			r.pass("Users.GetGroups")
		}
	}

	if firstUserID != "" {
		userSubs, err := s.Users.ListSubAccounts(ctx, firstUserID)
		if err != nil {
			r.fail("Users.ListSubAccounts", err)
		} else {
			fmt.Printf("    sub_accounts=%d\n", len(userSubs.SubAccounts))
			r.pass("Users.ListSubAccounts")
		}
	}

	testEmail := fmt.Sprintf("go-sdk-test-%d@cloudinary.com", time.Now().Unix())
	testRole := operations.RoleAdmin
	userCreated, err := s.Users.Create(ctx, operations.CreateUserRequestBody{
		Name:  "Go SDK Test User",
		Email: testEmail,
		Role:  &testRole,
	})
	if err != nil {
		r.fail("Users.Create", err)
	} else {
		createdUserID := deref(userCreated.ID)
		fmt.Printf("    created id=%s email=%s\n", createdUserID, deref(userCreated.Email))
		r.pass("Users.Create")

		updatedName := "Go SDK Test User Updated"
		userUpdated, err := s.Users.Update(ctx, operations.UpdateUserRequest{
			UserID:      createdUserID,
			UserRequest: components.UserRequest{Name: &updatedName},
		})
		if err != nil {
			r.fail("Users.Update", err)
		} else {
			fmt.Printf("    updated name=%s\n", deref(userUpdated.Name))
			r.pass("Users.Update")
		}

		_, err = s.Users.Delete(ctx, createdUserID)
		if err != nil {
			r.fail("Users.Delete", err)
		} else {
			fmt.Printf("    deleted id=%s\n", createdUserID)
			r.pass("Users.Delete")
		}
	}

	// =====================================================================
	// USER GROUPS
	// =====================================================================
	section("UserGroups")

	ugList, err := s.UserGroups.List(ctx, nil)
	if err != nil {
		r.fail("UserGroups.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(ugList.UserGroups))
		if len(ugList.UserGroups) > 0 {
			firstGroupID = deref(ugList.UserGroups[0].ID)
		}
		r.pass("UserGroups.List")
	}

	if firstGroupID != "" {
		ugGet, err := s.UserGroups.Get(ctx, operations.GetUserGroupRequest{GroupID: firstGroupID})
		if err != nil {
			r.fail("UserGroups.Get", err)
		} else {
			fmt.Printf("    id=%s name=%s\n", deref(ugGet.ID), deref(ugGet.Name))
			r.pass("UserGroups.Get")
		}
	}

	if firstGroupID != "" {
		ugUsers, err := s.UserGroups.ListUsers(ctx, firstGroupID)
		if err != nil {
			r.fail("UserGroups.ListUsers", err)
		} else {
			fmt.Printf("    users=%d\n", len(ugUsers.Users))
			r.pass("UserGroups.ListUsers")
		}
	}

	testGroupName := fmt.Sprintf("go-sdk-test-group-%d", time.Now().Unix())
	ugCreated, err := s.UserGroups.Create(ctx, components.UserGroupRequest{Name: &testGroupName})
	if err != nil {
		r.fail("UserGroups.Create", err)
	} else {
		createdGroupID := deref(ugCreated.ID)
		fmt.Printf("    created id=%s name=%s\n", createdGroupID, deref(ugCreated.Name))
		r.pass("UserGroups.Create")

		if firstUserID != "" {
			addedUser, err := s.UserGroups.AddUser(ctx, operations.AddUserToUserGroupRequest{
				GroupID: createdGroupID,
				UserID:  firstUserID,
			})
			if err != nil {
				r.fail("UserGroups.AddUser", err)
			} else {
				fmt.Printf("    added user=%s\n", deref(addedUser.ID))
				r.pass("UserGroups.AddUser")
			}

			_, err = s.UserGroups.RemoveUser(ctx, operations.RemoveUserFromUserGroupRequest{
				GroupID: createdGroupID,
				UserID:  firstUserID,
			})
			if err != nil {
				r.fail("UserGroups.RemoveUser", err)
			} else {
				fmt.Printf("    removed user=%s\n", firstUserID)
				r.pass("UserGroups.RemoveUser")
			}
		}

		updatedGroupName := testGroupName + "-updated"
		ugUpdated, err := s.UserGroups.Update(ctx, operations.UpdateUserGroupRequest{
			GroupID:          createdGroupID,
			UserGroupRequest: components.UserGroupRequest{Name: &updatedGroupName},
		})
		if err != nil {
			r.fail("UserGroups.Update", err)
		} else {
			fmt.Printf("    updated name=%s\n", deref(ugUpdated.Name))
			r.pass("UserGroups.Update")
		}

		_, err = s.UserGroups.Delete(ctx, createdGroupID)
		if err != nil {
			r.fail("UserGroups.Delete", err)
		} else {
			fmt.Printf("    deleted id=%s\n", createdGroupID)
			r.pass("UserGroups.Delete")
		}
	}

	// =====================================================================
	// ACCESS KEYS
	// =====================================================================
	section("AccessKeys")

	if firstProdEnvID != "" {
		akList, err := s.AccessKeys.List(ctx, operations.GetAccessKeysRequest{SubAccountID: firstProdEnvID})
		if err != nil {
			r.fail("AccessKeys.List", err)
		} else {
			fmt.Printf("    count=%d\n", len(akList.AccessKeys))
			r.pass("AccessKeys.List")
		}

		akCreated, err := s.AccessKeys.Generate(ctx, operations.GenerateAccessKeyRequest{
			SubAccountID:   firstProdEnvID,
			AccessKeyRequest: components.AccessKeyRequest{
				Name: cldprovisioning.String(fmt.Sprintf("go-sdk-test-key-%d", time.Now().Unix())),
			},
		})
		if err != nil {
			r.fail("AccessKeys.Generate", err)
		} else {
			akKey := deref(akCreated.APIKey)
			fmt.Printf("    generated key=%s name=%s\n", akKey, deref(akCreated.Name))
			r.pass("AccessKeys.Generate")

			updatedKeyName := "go-sdk-test-key-updated"
			akUpdated, err := s.AccessKeys.Update(ctx, operations.UpdateAccessKeyRequest{
				SubAccountID:         firstProdEnvID,
				Key:                  akKey,
				AccessKeyUpdateRequest: components.AccessKeyUpdateRequest{Name: &updatedKeyName},
			})
			if err != nil {
				r.fail("AccessKeys.Update", err)
			} else {
				fmt.Printf("    updated name=%s\n", deref(akUpdated.Name))
				r.pass("AccessKeys.Update")
			}

			_, err = s.AccessKeys.Delete(ctx, operations.DeleteAccessKeyRequest{
				SubAccountID: firstProdEnvID,
				Key:          akKey,
			})
			if err != nil {
				r.fail("AccessKeys.Delete", err)
			} else {
				fmt.Printf("    deleted key=%s\n", akKey)
				r.pass("AccessKeys.Delete")
			}
		}

		deleteName := fmt.Sprintf("go-sdk-delbyname-%d", time.Now().Unix())
		akForDel, err := s.AccessKeys.Generate(ctx, operations.GenerateAccessKeyRequest{
			SubAccountID:   firstProdEnvID,
			AccessKeyRequest: components.AccessKeyRequest{Name: &deleteName},
		})
		if err != nil {
			r.fail("AccessKeys.Generate (for DeleteByName)", err)
		} else {
			r.pass("AccessKeys.Generate (for DeleteByName)")
			_, err = s.AccessKeys.DeleteByName(ctx, operations.DeleteAccessKeyByNameRequest{
				SubAccountID: firstProdEnvID,
				Name:         deleteName,
			})
			if err != nil {
				r.fail("AccessKeys.DeleteByName", err)
				s.AccessKeys.Delete(ctx, operations.DeleteAccessKeyRequest{
					SubAccountID: firstProdEnvID,
					Key:          deref(akForDel.APIKey),
				})
			} else {
				fmt.Printf("    deleted by name=%s\n", deleteName)
				r.pass("AccessKeys.DeleteByName")
			}
		}
	} else {
		fmt.Println("  SKIP: no product environment ID available")
	}

	// =====================================================================
	// BILLING USAGE
	// =====================================================================
	section("BillingUsage")

	billingRes, err := s.BillingUsage.Get(ctx, &operations.GetBillingUsageRequest{})
	if err != nil {
		r.fail("BillingUsage.Get", err)
	} else {
		fmt.Printf("    plans=%d\n", len(billingRes))
		r.pass("BillingUsage.Get")
	}

	// =====================================================================
	// SYSTEM POLICIES
	// =====================================================================
	section("SystemPolicies")

	spList, err := s.SystemPolicies.List(ctx, &operations.GetSystemPoliciesRequest{})
	if err != nil {
		r.fail("SystemPolicies.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(spList.Data))
		r.pass("SystemPolicies.List")
	}

	// =====================================================================
	// CUSTOM POLICIES
	// =====================================================================
	section("CustomPolicies")

	cpList, err := s.CustomPolicies.List(ctx, &operations.GetCustomPoliciesRequest{})
	if err != nil {
		r.fail("CustomPolicies.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(cpList.Data))
		r.pass("CustomPolicies.List")
	}

	cpCreated, err := s.CustomPolicies.Create(ctx, components.CustomPolicy{
		Name:            fmt.Sprintf("go-sdk-test-policy-%d", time.Now().Unix()),
		PolicyStatement: `permit(principal, action, resource == Cloudinary::Feature::"cld::global::ml::access");`,
		ScopeType:       components.ScopeTypeEnumProdenv,
	})
	if err != nil {
		r.fail("CustomPolicies.Create", err)
	} else {
		cpID := cpCreated.Data.ID
		fmt.Printf("    created id=%s name=%s\n", cpID, cpCreated.Data.Name)
		r.pass("CustomPolicies.Create")

		cpGet, err := s.CustomPolicies.Get(ctx, cpID)
		if err != nil {
			r.fail("CustomPolicies.Get", err)
		} else {
			fmt.Printf("    got id=%s name=%s\n", cpGet.Data.ID, cpGet.Data.Name)
			r.pass("CustomPolicies.Get")
		}

		updatedPolicyName := fmt.Sprintf("go-sdk-test-policy-%d-updated", time.Now().Unix())
		cpUpdated, err := s.CustomPolicies.Update(ctx, operations.UpdateCustomPolicyRequest{
			PolicyID: cpID,
			UpdateCustomPolicy: components.UpdateCustomPolicy{
				Name:            updatedPolicyName,
				PolicyStatement: `permit(principal, action, resource == Cloudinary::Feature::"cld::global::ml::access");`,
				ScopeType:       components.ScopeTypeEnumProdenv,
			},
		})
		if err != nil {
			r.fail("CustomPolicies.Update", err)
		} else {
			fmt.Printf("    updated name=%s\n", cpUpdated.Data.Name)
			r.pass("CustomPolicies.Update")
		}

		err = s.CustomPolicies.Delete(ctx, cpID)
		if err != nil {
			r.fail("CustomPolicies.Delete", err)
		} else {
			fmt.Printf("    deleted id=%s\n", cpID)
			r.pass("CustomPolicies.Delete")
		}
	}

	// =====================================================================
	// EFFECTIVE POLICIES
	// =====================================================================
	section("EffectivePolicies")

	epScopeType := components.ScopeTypeEnumAccount
	epList, err := s.EffectivePolicies.List(ctx, &operations.GetEffectivePoliciesRequest{
		ScopeType: &epScopeType,
	})
	if err != nil {
		r.fail("EffectivePolicies.List", err)
	} else {
		hasData := epList.Data != nil
		fmt.Printf("    has_data=%v\n", hasData)
		r.pass("EffectivePolicies.List")
	}

	// =====================================================================
	// ROLES
	// =====================================================================
	section("Roles")

	rolesList, err := s.Roles.List(ctx, operations.GetRolesRequest{
		PermissionType: components.PermissionTypeEnumGlobal,
	})
	if err != nil {
		r.fail("Roles.List", err)
	} else {
		fmt.Printf("    count=%d\n", len(rolesList.Data))
		if len(rolesList.Data) > 0 {
			firstRoleID = rolesList.Data[0].ID
		}
		r.pass("Roles.List")
	}

	if firstRoleID != "" {
		roleGet, err := s.Roles.Get(ctx, operations.GetRoleRequest{RoleID: firstRoleID})
		if err != nil {
			r.fail("Roles.Get", err)
		} else {
			fmt.Printf("    id=%s name=%s\n", roleGet.Data.ID, roleGet.Data.Name)
			r.pass("Roles.Get")
		}
	}

	if firstRoleID != "" {
		rolePrincipals, err := s.Roles.ListPrincipals(ctx, operations.GetRolePrincipalsRequest{RoleID: firstRoleID})
		if err != nil {
			r.fail("Roles.ListPrincipals", err)
		} else {
			fmt.Printf("    principals=%d\n", len(rolePrincipals.Data))
			r.pass("Roles.ListPrincipals")
		}
	}

	roleCreated, err := s.Roles.Create(ctx, components.CreateRole{
		Name:            fmt.Sprintf("go-sdk-test-role-%d", time.Now().Unix()),
		PermissionType:  "global",
		ScopeType:       components.ScopeTypeEnumProdenv,
		SystemPolicyIds: []string{},
	})
	if err != nil {
		r.fail("Roles.Create", err)
	} else {
		roleID := roleCreated.Data.ID
		fmt.Printf("    created id=%s name=%s\n", roleID, roleCreated.Data.Name)
		r.pass("Roles.Create")

		updatedRoleName := fmt.Sprintf("go-sdk-test-role-%d-updated", time.Now().Unix())
		roleUpdated, err := s.Roles.Update(ctx, operations.UpdateRoleRequest{
			RoleID: roleID,
			Role:   components.Role{Name: updatedRoleName, SystemPolicyIds: []string{}},
		})
		if err != nil {
			r.fail("Roles.Update", err)
		} else {
			fmt.Printf("    updated name=%s\n", roleUpdated.Data.Name)
			r.pass("Roles.Update")
		}

		if firstUserID != "" {
			err = s.Roles.UpdatePrincipals(ctx, operations.UpdateRolePrincipalsRequest{
				RoleID: roleID,
				UpdateRolePrincipalsRequest: components.UpdateRolePrincipalsRequest{
					Operation: components.OperationEnumAdd.ToPointer(),
					Principals: []components.UpdateRolePrincipal{{
						PrincipalType: components.PrincipalTypeEnumUser,
						PrincipalID:   firstUserID,
					}},
				},
			})
			if err != nil {
				r.fail("Roles.UpdatePrincipals (add)", err)
			} else {
				fmt.Printf("    added principal user=%s\n", firstUserID)
				r.pass("Roles.UpdatePrincipals (add)")
			}

			err = s.Roles.UpdatePrincipals(ctx, operations.UpdateRolePrincipalsRequest{
				RoleID: roleID,
				UpdateRolePrincipalsRequest: components.UpdateRolePrincipalsRequest{
					Operation: components.OperationEnumRemove.ToPointer(),
					Principals: []components.UpdateRolePrincipal{{
						PrincipalType: components.PrincipalTypeEnumUser,
						PrincipalID:   firstUserID,
					}},
				},
			})
			if err != nil {
				r.fail("Roles.UpdatePrincipals (remove)", err)
			} else {
				fmt.Printf("    removed principal user=%s\n", firstUserID)
				r.pass("Roles.UpdatePrincipals (remove)")
			}
		}

		err = s.Roles.Delete(ctx, operations.DeleteRoleRequest{RoleID: roleID})
		if err != nil {
			r.fail("Roles.Delete", err)
		} else {
			fmt.Printf("    deleted id=%s\n", roleID)
			r.pass("Roles.Delete")
		}
	}

	// =====================================================================
	// PRINCIPALS
	// =====================================================================
	section("Principals")

	if firstUserID != "" {
		principalRoles, err := s.Principals.ListRoles(ctx, operations.GetPrincipalRolesRequest{
			PrincipalType:  components.PrincipalTypeEnumUser,
			PrincipalID:    firstUserID,
			PermissionType: components.PermissionTypeEnumGlobal,
		})
		if err != nil {
			r.fail("Principals.ListRoles", err)
		} else {
			fmt.Printf("    roles=%d\n", len(principalRoles.Data))
			r.pass("Principals.ListRoles")
		}
	}

	if firstUserID != "" {
		inspectRes, err := s.Principals.Inspect(ctx, operations.InspectRequest{
			ScopeType:     components.ScopeTypeEnumAccount,
			PrincipalType: components.PrincipalTypeEnumUser.ToPointer(),
			PrincipalID:   &firstUserID,
		})
		if err != nil {
			r.fail("Principals.Inspect", err)
		} else {
			fmt.Printf("    inspect result=%v\n", inspectRes != nil)
			r.pass("Principals.Inspect")
		}
	}

	if firstUserID != "" {
		inspectMultiRes, err := s.Principals.InspectMultiple(ctx, components.InspectRequest{
			ScopeType: components.ScopeTypeEnumAccount,
			Principals: []components.Principal{{
				Type: components.PrincipalTypeEnumUser,
				ID:   firstUserID,
			}},
		})
		if err != nil {
			r.fail("Principals.InspectMultiple", err)
		} else {
			fmt.Printf("    inspect_multiple result=%v\n", inspectMultiRes != nil)
			r.pass("Principals.InspectMultiple")
		}
	}

	fmt.Println("  SKIP: Principals.UpdateRoles (tested indirectly via Roles.UpdatePrincipals)")

	// =====================================================================
	// SUMMARY
	// =====================================================================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Printf("  RESULTS: %d passed, %d failed\n", r.passed, r.failed)
	fmt.Println("============================================================")
	if r.failed > 0 {
		fmt.Println("\n  FAILURES:")
		for _, e := range r.errors {
			fmt.Printf("    %s\n", e)
		}
		os.Exit(1)
	}
}

func section(name string) {
	fmt.Printf("\n--- %s ---\n", name)
}

func deref(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func mapKeys(m map[string]any) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
