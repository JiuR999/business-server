package common

import "BusinessServer/env"

type PageModel struct {
	PageNum  int `json:"pageNum"`  //页码
	PageSize int `json:"pageSize"` //分页大小
}

func (p *PageModel) IfAbsent() {
	if p.PageSize <= 0 {
		p.PageSize = env.GetConfig().PageSize
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
}
