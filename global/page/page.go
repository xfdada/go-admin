package page

type Page struct {
	Total    int64   `json:"total"`     //总记录条数
	Pages    float64 `json:"pages"`     //页码总数
	PageSize int64   `json:"page_size"` //每页数量
	NowPage  int     `json:"now_page"`  //当前页
}
