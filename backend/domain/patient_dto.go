package domain

import "github.com/KScaesar/jubo-homework/backend/util"

func TransformPatientModel(patient *Patient) DtoPatientResponse {
	return DtoPatientResponse{
		Id:   patient.Id,
		Name: patient.Name,
	}
}

// read dto

type DtoPatientResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DtoQryPatientParam struct {
	DtoFilterPatientParam
	DtoSortPatientParam
	util.DtoPageParam
}

func (d *DtoQryPatientParam) FilterParam() any {
	return &d.DtoFilterPatientParam
}

func (d *DtoQryPatientParam) SortParam() any {
	return &d.DtoSortPatientParam
}

func (d *DtoQryPatientParam) PageParam() util.DtoPageParam {
	return d.DtoPageParam
}

type DtoFilterPatientParam struct {
	Name *string `form:"name" rdb:"name = ?"`
}

type DtoSortPatientParam struct {
	SortCreatedAt util.SortKind `form:"sort_created_at" rdb:"created_at" validate:"sort"`
}

func (dto *DtoSortPatientParam) SetDefault() {
	if dto.SortCreatedAt == "" {
		dto.SortCreatedAt = util.SortDesc
	}
}
