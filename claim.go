package jwt

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/samber/lo"
)

type BaseClaims struct {
	UserID   int64    `json:"user_id,omitempty"`
	TenantID int64    `json:"tenant_id,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

type Claims struct {
	BaseClaims
	jwt.RegisteredClaims
}

func (claim *Claims) IsAdmin() bool {
	return lo.Contains(claim.Roles, RoleSysAdmin.String()) || lo.Contains(claim.Roles, RoleSuperAdmin.String())
}

func (claim *Claims) IsTenantAdmin() bool {
	return lo.Contains(claim.Roles, RoleTenantAdmin.String())
}
