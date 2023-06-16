package utils

import (
	"fmt"
	"strings"
)

func CreateListFilter(tag, author, favorited, limit, offset string) string {
	filterString := ""
	var filters []string
	if tag != "" {
		filters = append(filters, fmt.Sprintf("'%s' = ANY(array_agg(DISTINCT t.title))", tag))
	}
	if author != "" {
		filters = append(filters, fmt.Sprintf("'%s' = u.username", author))
	}
	if favorited != "" {
		filters = append(filters, fmt.Sprintf("'%s' = ANY(array_agg(DISTINCT u_fav.username))", favorited))
	}

	if len(filters) != 0 {
		filterString = " HAVING " + strings.Join(filters, " AND ")
	}

	if limit != "" {
		filterString += fmt.Sprintf(" LIMIT %s", limit)
	}
	if offset != "" {
		filterString += fmt.Sprintf(" OFFSET %s", offset)
	}
	return filterString
}
