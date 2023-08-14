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
	Referrals     []Referral       `json:"referrals" validate:"required"`
	DesiredSalary float64          `json:"desired_salary" validate:"required"`
	Availability  string           `json:"availability" validate:"required"`
	Education     []EducationInfo  `json:"education"`
	WorkHistory   []WorkExperience `json:"work_history"`
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
