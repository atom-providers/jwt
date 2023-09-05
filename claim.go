package jwt

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	UserID   uint64 `json:"user_id,omitempty"`
	TenantID uint64 `json:"tenant_id,omitempty"`
	RoleID   uint64 `json:"role_id,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Claims struct {
	BaseClaims
	jwt.RegisteredClaims
}

func (claim *Claims) IsAdmin() bool {
	return claim.Role == RoleSystemAdmin.String() || claim.Role == RoleSuperAdmin.String()
}

func (claim *Claims) IsTenantAdmin() bool {
	return claim.Role == RoleTenantAdmin.String()
}
