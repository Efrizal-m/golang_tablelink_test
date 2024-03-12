package user

import (
	"time"
)

// User ...
type User struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RoleID     string    `json:"role_id"`
	RoleName   string    `json:"role_name"`
	Name       int       `json:"name"`
	Email      int       `json:"email"`
	Password   int       `json:"password"`
	LastAccess time.Time `json:"last_access"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// FilterSpec ...
type FilterSpec struct {
	Name string
	ID   int
}
