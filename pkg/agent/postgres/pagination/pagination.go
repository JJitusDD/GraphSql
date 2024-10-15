package pagination

import (
	"gorm.io/gorm"
)

// IPaginator ...
type IPaginatorPostgre interface {
	Search(out interface{}) error
	Sort(sortField string, sortValue string) IPaginatorPostgre
	Limit(limit int64) IPaginatorPostgre
	SetModel(model interface{}) IPaginatorPostgre
	SetDB(db *gorm.DB) IPaginatorPostgre
	SetTable(table string) IPaginatorPostgre
	Filter(filter map[string]interface{}) IPaginatorPostgre
	Page(cursor Cursor) IPaginatorPostgre
	Data() *PaginationData
	SetPreLoads(preloads []string) IPaginatorPostgre
	SetDisabledCount(disable bool) IPaginatorPostgre
	SetJoins(join []string) IPaginatorPostgre
	SetSelectCol(col string) IPaginatorPostgre
	SetGroup(group string) IPaginatorPostgre
}
