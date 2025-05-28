package helper

import (
	"strings"

	"gorm.io/gorm"
)

func BuildQuery(filters *map[string]string, tx *gorm.DB) {
	idLike, ok := (*filters)["id.in"]
	if ok {
		values := strings.Split(idLike, ",")
		conditions := make([]string, len(values))

		for i := range values {
			conditions[i] = "panen_payments.id LIKE ?"
		}

		tx.Where(strings.Join(conditions, " OR "), addWildcards(values)...)
		delete(*filters, "id.in")
	}
}

func addWildcards(values []string) []interface{} {
	wildcards := make([]interface{}, len(values))
	for i, v := range values {
		wildcards[i] = "%" + v + "%"
	}
	return wildcards
}
