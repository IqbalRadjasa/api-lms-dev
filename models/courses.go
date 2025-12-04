package models

type Courses struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	CategoryId   int    `json:"category_id" validate:"required"`
	InstructorId int    `json:"instructor_id" validate:"required"`
	Slug         string `json:"slug" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Thumbnail    string `json:"thumbnail" validate:"required"`
}
