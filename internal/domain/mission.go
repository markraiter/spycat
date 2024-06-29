package domain

type Mission struct {
	ID        int    `json:"id"`
	CatID     int    `json:"cat_id" validate:"required" example:"1"`
	TargetIDs []int  `json:"targets" validate:"dive,required" example:"[1,2,3]"`
	Notes     string `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool   `json:"completed" validate:"required" example:"false"`
}
