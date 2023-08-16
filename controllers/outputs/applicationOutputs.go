package outputs

import (
	"time"

	inputs "github.com/davisshriver/job-board-backend-project/controllers/inputs"
)

type ApplicationOutput struct {
	ApplicationID int                     `json:"application_id"`
	UserID        int                     `json:"user_id"`
	PostID        int                     `json:"post_id"`
	FirstName     string                  `json:"first_name"`
	LastName      string                  `json:"last_name"`
	Email         string                  `json:"email"`
	Phone         string                  `json:"phone"`
	Address       string                  `json:"address"`
	City          string                  `json:"city"`
	State         string                  `json:"state"`
	PostalCode    string                  `json:"postal_code"`
	CoverLetter   string                  `json:"cover_letter"`
	ResumeURL     string                  `json:"resume_url"`
	LinkedInURL   string                  `json:"linkedin_url"`
	PortfolioURL  string                  `json:"portfolio_url"`
	Referrals     []inputs.Referral       `json:"referrals"`
	DesiredSalary float64                 `json:"desired_salary"`
	Availability  string                  `json:"availability"`
	Education     []inputs.EducationInfo  `json:"education"`
	WorkHistory   []inputs.WorkExperience `json:"work_history"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"expires_at"`
}
