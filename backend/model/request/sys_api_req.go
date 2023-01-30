package request

import "job-maching/model/system"

type SearchApiReq struct {
	system.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
