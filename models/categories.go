package models

type Categories struct {
	ID           int         `json:"id" gorm:"primaryKey"`
	DepartmentId int         `json:"department_id" validate:"required"`
	Name         string      `json:"name" validate:"max=100"`
	Slug         string      `json:"slug" validate:"max=100"`
	Department   Departments `json:"departments" gorm:"foreignKey=DepartmentId"`
}
