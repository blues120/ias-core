package pagination

// Pagination 分页
type Pagination struct {
	PageNum  int
	PageSize int
}

// IsValid 判断分页参数是否合法
func (p *Pagination) IsValid() bool {
	if p.PageSize == 0 {
		return false
	}
	return true
}

// Offset 计算偏移量
func (p *Pagination) Offset() int {
	if p.PageNum > 1 {
		return (p.PageNum - 1) * p.PageSize
	}
	return 0
}
