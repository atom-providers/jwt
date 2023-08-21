package jwt

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	UserID   int64  `json:"user_id,omitempty"`
	TenantID int64  `json:"tenant_id,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Claims struct {
	BaseClaims
	jwt.RegisteredClaims
}

func (claim *Claims) IsAdmin() bool {
	return claim.Role == RoleSysAdmin.String() || claim.Role == RoleSuperAdmin.String()
}

func (claim *Claims) IsTenantAdmin() bool {
	return claim.Role == RoleTenantAdmin.String()
}
