package pagination

import (
	"gorm.io/gorm"
	"math"
	"strings"
)

type Cursor int64

type Paginator struct {
	page        Cursor
	next        Cursor
	prev        Cursor
	limit       int64
	offset      int64
	totalRecord int64
	totalPage   int64

	filterQuery map[string]interface{}

	db    *gorm.DB
	model interface{}
	table string

	sortKeys string

	aggCount     int64
	sf           map[string]string
	preload      []string
	disableCount bool
	joins        []string
	selectCol    string
	group        string
}

// PaginationData struct for returning pagination stat
type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

func (p *Paginator) Search(out interface{}) error {
	p.initOptions()
	ctx := getCtx()
	tx := p.db.WithContext(ctx).Table(p.table)
	if len(p.filterQuery) > 0 {
		for k, v := range p.filterQuery {
			if strings.Contains(k, "NULL") {
				tx = tx.Where(k)
			} else {
				tx = tx.Where(k, v)
			}
		}
	}
	skip := getSkip(int64(p.page), p.limit)
	if p.sortKeys != "" {
		tx = tx.Order(p.sortKeys)
	}

	if len(p.joins) > 0 {
		for _, join := range p.joins {
			tx = tx.Joins(join)
		}
	}
	if p.selectCol != "" {
		tx = tx.Select(p.selectCol)
	}
	if len(p.preload) > 0 {
		for _, table := range p.preload {
			tx = tx.Preload(table)
		}
	}
	if p.group != "" {
		tx = tx.Group(p.group)
	}

	tx = tx.Offset(int(skip)).Limit(int(p.limit)).Find(out)
	if tx.Error != nil {
		return tx.Error
	}
	if !p.disableCount {
		p.PagingV2()
	}
	return nil
}

func (p *Paginator) Sort(sortField string, sortValue string) IPaginatorPostgre {
	if p.sortKeys == "" {
		p.sortKeys = sortField + " " + sortValue
	}
	return p
}

func (p *Paginator) Limit(limit int64) IPaginatorPostgre {
	p.limit = limit
	return p
}

func (p *Paginator) SetPreLoads(preloads []string) IPaginatorPostgre {
	p.preload = preloads
	return p
}

func (p *Paginator) SetGroup(group string) IPaginatorPostgre {
	p.group = group
	return p
}

func (p *Paginator) SetModel(model interface{}) IPaginatorPostgre {
	p.model = model
	return p
}

func (p *Paginator) SetDisabledCount(disable bool) IPaginatorPostgre {
	p.disableCount = disable
	return p
}

func (p *Paginator) SetDB(db *gorm.DB) IPaginatorPostgre {
	p.db = db
	return p
}

func (p *Paginator) SetTable(table string) IPaginatorPostgre {
	p.table = table
	return p
}

func (p *Paginator) Filter(filter map[string]interface{}) IPaginatorPostgre {
	p.filterQuery = filter
	return p
}

func (p *Paginator) SetJoins(join []string) IPaginatorPostgre {
	p.joins = join
	return p
}

func (p *Paginator) SetSelectCol(col string) IPaginatorPostgre {
	p.selectCol = col
	return p
}

func (p *Paginator) Page(cursor Cursor) IPaginatorPostgre {
	if cursor < 1 {
		p.page = 1
	} else {
		p.page = cursor
	}
	return p

}

func (p Paginator) Data() *PaginationData {
	return &PaginationData{
		Total:     p.totalRecord,
		Page:      int64(p.page),
		PerPage:   p.limit,
		Prev:      int64(p.prev),
		Next:      int64(p.next),
		TotalPage: p.totalPage,
	}
}

func (p *Paginator) initOptions() {
	p.limit = getPaginationLimit(p.limit)
}

func (p *Paginator) PagingV2() {
	var offset int64
	var count int64
	ctx := getCtx()
	tx := p.db.WithContext(ctx).Table(p.table)
	if len(p.filterQuery) > 0 {
		for k, v := range p.filterQuery {
			if strings.Contains(k, "NULL") {
				tx = tx.Where(k)
			} else {
				tx = tx.Where(k, v)
			}
		}
	}
	if len(p.joins) > 0 {
		for _, join := range p.joins {
			tx = tx.Joins(join)
		}
	}

	tx = tx.Count(&count)

	if p.page == 1 {
		offset = 0
	} else {
		offset = int64(p.page) - 1*p.limit
	}

	p.offset = offset
	p.totalRecord = count
	p.totalPage = int64(math.Ceil(float64(count) / float64(p.limit)))

	if p.page > 1 {
		p.prev = p.page - 1
	} else {
		p.prev = p.page
	}
	if int64(p.page) >= p.totalPage {
		p.next = p.page
	} else {
		p.next = p.page + 1
	}
}

func New() IPaginatorPostgre {
	return &Paginator{}
}
