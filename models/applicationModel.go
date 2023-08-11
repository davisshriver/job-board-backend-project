package models

import "time"

type Application struct {
	ApplicationID int              `gorm:"primaryKey;autoIncrement"`
	UserID        int              `json:"user_id"`
    PostID        int              `json:"post_id"`
	FirstName     string           `json:"first_name" validate:"required,min=2,max=100"`
	LastName      string           `json:"first_name" validate:"requiredmin=2,max=100"`
	Email         string           `json:"email" validate:"email,required"`
	Phone         string           `json:"phone" validate:"required,min=1,max=10"`
	Address       string           `json:"address" validate:"required"`
	City          string           `json:"city" validate:"required"`
	State         string           `json:"state" validate:"required"`
	PostalCode    string           `json:"postal_code" validate:"required"`
	CoverLetter   string           `json:"cover_letter"`
	ResumeURL     string           `json:"resume_url" validate:"required"`
	LinkedInURL   string           `json:"linkedin_url"`
	PortfolioURL  string           `json:"portfolio_url"`
	References    []string         `json:"references"`  
	DesiredSalary float64          `json:"desired_salary" validate:"required"`  
	Availability  string           `json:"availability" validate:"required"`
	Education     []EducationInfo  `json:"education"`
	WorkHistory   []WorkExperience `json:"work_history"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"expires_at"`
}

type EducationInfo struct {
	Degree   string `json:"degree"`
	School   string `json:"school"`
	Location string `json:"location"`
	GradYear int    `json:"grad_year"`
}

type WorkExperience struct {
	Position         string   `json:"position"`
	Company          string   `json:"company"`
	Location         string   `json:"location"`
	StartYear        int      `json:"start_year"`
	EndYear          int      `json:"end_year"`
	Responsibilities []string `json:"responsibilities"`
}
