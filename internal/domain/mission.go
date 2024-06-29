package domain

type Mission struct {
	ID        int      `json:"id"`
	CatID     int      `json:"cat_id" validate:"required" example:"1"`
	Targets   []Target `json:"targets" validate:"dive,required"`
	Notes     string   `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool     `json:"completed" validate:"omitempty" example:"false"`
}

type MissionRequest struct {
	CatID     int      `json:"cat_id" validate:"required" example:"1"`
	Targets   []Target `json:"targets" validate:"dive,required"`
	Notes     string   `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool     `json:"completed" validate:"omitempty" example:"false"`
}
