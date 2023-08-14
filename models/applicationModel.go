package models

import "time"

type Application struct {
	ApplicationID int       `gorm:"primaryKey;autoIncrement"`
	UserID        int       `gorm:"column:user_id"`
	PostID        int       `gorm:"column:post_id"`
	FirstName     string    `gorm:"column:first_name"`
	LastName      string    `gorm:"column:last_name"`
	Email         string    `gorm:"column:email;type:varchar(255);not null"`
	Phone         string    `gorm:"column:phone;type:varchar(10);not null"`
	Address       string    `gorm:"column:address;type:varchar(255);not null"`
	City          string    `gorm:"column:city;type:varchar(255);not null"`
	State         string    `gorm:"column:state;type:varchar(255);not null"`
	PostalCode    string    `gorm:"column:postal_code;type:varchar(10);not null"`
	CoverLetter   string    `gorm:"column:cover_letter"`
	ResumeURL     string    `gorm:"column:resume_url;type:varchar(255);not null"`
	LinkedInURL   string    `gorm:"column:linkedin_url"`
	PortfolioURL  string    `gorm:"column:portfolio_url"`
	DesiredSalary float64   `gorm:"column:desired_salary;not null"`
	Availability  string    `gorm:"column:availability;not null"`
	Education     []byte    `gorm:"column:education;type:jsonb"`
	Referrals     []byte    `gorm:"column:referrals;type:jsonb"`
	WorkHistory   []byte    `gorm:"column:work_history;type:jsonb"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:expires_at"`
}
