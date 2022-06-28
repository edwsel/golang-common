package rule

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
)

// UniqueRule to check data is unique or not in db
type UniqueRule struct {
	db *sql.DB
}

func NewUniqueRule(db *sql.DB, ruleName string) *UniqueRule {
	return &UniqueRule{db}
}

// Rule is func to register as custom rule
func (r *UniqueRule) Rule(filed validator.FieldLevel) error {
	filed.
}
