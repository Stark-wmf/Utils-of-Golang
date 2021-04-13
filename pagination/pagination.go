package pagination

import (
	"golangutils/pconst"
	"math"
)
type Paginatemodel struct {
	Total       uint32 `json:"total"`        //总条数
	CurrentPage uint32 `json:"current_page"` //当前页
	PerPage     uint32 `json:"per_page"`     //每页数
	Start       uint32 `json:"start"`        //当页开始数
	End         uint32 `json:"end"`          //当页结束数
	Offset      uint32 `json:"offset"`       //偏移量
}

func Paginate(currentPage int, perPage int, total int) (int, int, bool, bool) {
	var hasPrev, hasNext bool
	var offset int
	if currentPage == 0 {
		currentPage = 1
	}
	offset = (currentPage - 1) * perPage
	if currentPage > 1 {
		hasPrev = true
	}
	if total > (currentPage * perPage) {
		hasNext = true
	}
	return offset, currentPage, hasPrev, hasNext
}

type Pagination struct {
	Total       uint32 `json:"total"`        //总条数
	CurrentPage uint32 `json:"current_page"` //当前页
	PerPage     uint32 `json:"per_page"`     //每页数
	Start       uint32 `json:"start"`        //当页开始数
	End         uint32 `json:"end"`          //当页结束数
	Offset      uint32 `json:"offset"`       //偏移量
}

func GetPaginate(total, currentPage, perPage uint32) *Pagination {
	pt := new(Pagination)
	pt.Total = total
	pt.CurrentPage = GetCurrentPage(total, currentPage, perPage)
	pt.PerPage = GetPerPage(perPage)
	pt.Start = pt.GetStartNum()
	pt.End = pt.GetEndNum()
	pt.Offset = GetOffset(pt.CurrentPage, pt.PerPage)
	return pt
}

func GetListPaginate(total, currentPage, perPage uint32) (paginate Paginatemodel) {
	pt := GetPaginate(total, currentPage, perPage)
	paginate.Total = pt.Total
	paginate.CurrentPage = pt.CurrentPage
	paginate.PerPage = pt.PerPage
	paginate.Start = pt.Start
	paginate.End = pt.End
	return paginate
}

func GetPerPage(perPage uint32) uint32 {
	if perPage == 0 {
		perPage = pconst.COMMON_PAGE_LIMIT_NUM_20
	}
	return perPage
}

func GetPage(page uint32) uint32 {
	if page == 0 {
		page = 1
	}
	return page
}

func GetCurrentPage(total, currentPage, perPage uint32) uint32 {
	maxPage := uint32(math.Ceil(float64(total) / float64(perPage)))
	if currentPage == 0 {
		currentPage = 1
	}
	if currentPage > maxPage {
		currentPage = maxPage
	}
	return currentPage
}

func GetOffset(currentPage, perPage uint32) uint32 {
	perPage = GetPerPage(perPage)
	currentPage = GetPage(currentPage)
	offset := (currentPage - 1) * perPage
	if offset < 0 {
		offset = 0
	}
	return offset
}

func (pt *Pagination) GetCurrentPage() uint32 {
	maxPage := uint32(math.Floor(float64(pt.Total) / float64(pt.PerPage)))
	if pt.CurrentPage == 0 {
		pt.CurrentPage = 1
	}
	if pt.CurrentPage > maxPage {
		pt.CurrentPage = maxPage
	}
	return pt.CurrentPage
}

func (pt *Pagination) GetStartNum() uint32 {
	pt.Start = pt.PerPage*(pt.CurrentPage-1) + 1
	if pt.Start < 0 {
		pt.Start = 0
	}
	return pt.Start
}

func (pt *Pagination) GetEndNum() uint32 {
	pt.End = pt.PerPage * pt.CurrentPage
	if pt.End > pt.Total {
		pt.End = pt.Total
	}
	return pt.End
}
