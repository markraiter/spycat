package domain

type Cat struct {
	ID                int    `json:"id"`
	Name              string `json:"name" validate:"required" example:"Tom"`
	YearsOfExperience int    `json:"years_of_experience" validate:"omitempty" example:"5"`
	Breed             string `json:"breed" validate:"required" example:"Siamese"`
	Salary            int    `json:"salary" validate:"omitempty" example:"1000"`
}
