package security

import (
	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
	"aahframework.org/security.v0/authz"
)

// FormAuthorizationProvider struct implements `authz.Authorizer` interface.
type FormAuthorizationProvider struct {
	// for demo purpose in-memory authorization info's
	// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
	// Information.
	users map[string]authz.AuthorizationInfo
}

// Init method initializes the FormAuthoriationProvider, this method gets called
// during server start up.
func (fa *FormAuthorizationProvider) Init(cfg *config.Config) error {

	// NOTE: for demo purpose I'm creating set users in the map.
	// Typically you will be using Database, API calls, LDAP, etc to get the Authorization
	// Information.

	fa.users = make(map[string]authz.AuthorizationInfo)

	// Subject 1
	authzInfo1 := authz.NewAuthorizationInfo()

	// Adding roles
	authzInfo1.AddRole("user", "manager")

	// Adding permissions
	// composed of domain is users, resource is manage, action is view
	// Learn about permission: http://docs.aahframework.org/security-permissions.html
	authzInfo1.AddPermissionString("users:manage:view")
	fa.users["user1@example.com"] = *authzInfo1

	// Subject 2
	authzInfo2 := authz.NewAuthorizationInfo()
	authzInfo2.AddRole("user", "administrator")
	// Learn about permission: http://docs.aahframework.org/security-permissions.html
	authzInfo2.AddPermissionString("users:*")
	fa.users["admin@example.com"] = *authzInfo2

	// Subject 3
	// We are not adding AuthorizationInfo since that user is locked.
	// Not required for demo.

	// Subject 4
	authzInfo4 := authz.NewAuthorizationInfo()
	authzInfo4.AddRole("user")
	fa.users["user3@example.com"] = *authzInfo4

	return nil
}

// GetAuthorizationInfo method is `authz.Authorizer` interface.
//
// GetAuthorizationInfo method gets called after authentication is successful
// to get Subject's (aka User) access control information such as roles and permissions.
func (fa *FormAuthorizationProvider) GetAuthorizationInfo(authcInfo *authc.AuthenticationInfo) *authz.AuthorizationInfo {

	if ai, found := fa.users[authcInfo.PrimaryPrincipal().Value]; found {
		return &ai
	}

	// No AuthorizationInfo for the subject, return nil
	return nil
}
