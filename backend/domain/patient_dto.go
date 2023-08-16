package domain

import "github.com/KScaesar/jubo-homework/backend/util"

func TransformPatientModel(patient *Patient) PatientResponse {
	return PatientResponse{
		Id:   patient.Id,
		Name: patient.Name,
	}
}

// read dto

type PatientResponse struct {
	Id   string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type QryPatientParam struct {
	FilterPatientParam
	util.SortParam
	util.PageParam
}

func (d *QryPatientParam) PreProcess(isPagination bool) {
	d.SortParam.SetDefaultIfInvalid("created_at", util.SortDesc)
	if isPagination {
		d.PageParam.SetDefaultIfInvalid()
		return
	}
	d.PageParam.SetWithoutPagination()
}

type FilterPatientParam struct {
	Name *string `form:"name" rdb:"name = ?"`
}
