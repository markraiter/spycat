package domain

type Target struct {
	ID        int    `json:"id"`
	MissionID int    `json:"mission_id" validate:"required" example:"1"`
	Name      string `json:"name" validate:"required" example:"John Doe"`
	Country   string `json:"country" validate:"required" example:"USA"`
	Notes     string `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool   `json:"completed" validate:"omitepmty" example:"false"`
}
