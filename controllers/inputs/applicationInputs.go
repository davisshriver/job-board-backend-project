package inputs

type ApplicationInput struct {
	FirstName     string           `json:"first_name" validate:"required,min=2,max=100"`
	LastName      string           `json:"last_name" validate:"required,min=2,max=100"`
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
	DesiredSalary float64          `json:"desired_salary" validate:"required"`
	Availability  string           `json:"availability" validate:"required"`
	Education     []EducationInfo  `json:"education"`
	Referrals     []Referral       `json:"referrals" validate:"required"`
	WorkHistory   []WorkExperience `json:"work_history"`
}

type ApplicationUpdateInput struct {
	FirstName     *string          `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName      *string          `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email         *string          `json:"email" validate:"omitempty,email"`
	Phone         *string          `json:"phone" validate:"omitempty"`
	Address       *string          `json:"address" validate:"omitempty"`
	City          *string          `json:"city" validate:"omitempty"`
	State         *string          `json:"state" validate:"omitempty"`
	PostalCode    *string          `json:"postal_code" validate:"omitempty"`
	CoverLetter   *string          `json:"cover_letter" validate:"omitempty"`
	ResumeURL     *string          `json:"resume_url" validate:"omitempty"`
	LinkedInURL   *string          `json:"linkedin_url" validate:"omitempty"`
	PortfolioURL  *string          `json:"portfolio_url" validate:"omitempty"`
	DesiredSalary *float64         `json:"desired_salary" validate:"omitempty"`
	Availability  *string          `json:"availability" validate:"omitempty"`
	Education     []EducationInfo  `json:"education" validate:"omitempty"`
	Referrals     []Referral       `json:"referrals" validate:"omitempty"`
	WorkHistory   []WorkExperience `json:"work_history" validate:"omitempty"`
}

type Referral struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Relation string `json:"relation"`
	Title    string `json:"title"`
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
