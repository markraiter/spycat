package domain

type Mission struct {
	ID        int      `json:"id"`
	CatID     int      `json:"cat_id" validate:"required" example:"1"`
	Targets   []Target `json:"targets" validate:"dive,required" example:"[{\"id\":1,\"mission_id\":1,\"name\":\"John Doe\",\"country\":\"USA\",\"notes\":\"Lorem ipsum\",\"completed\":false}]"`
	Notes     string   `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool     `json:"completed" validate:"required" example:"false"`
}

type Target struct {
	ID        int    `json:"id"`
	MissionID int    `json:"mission_id" validate:"required" example:"1"`
	Name      string `json:"name" validate:"required" example:"John Doe"`
	Country   string `json:"country" validate:"required" example:"USA"`
	Notes     string `json:"notes" validate:"omitempty" example:"Lorem ipsum"`
	Completed bool   `json:"completed" validate:"required" example:"false"`
}
