package entity

type StaffEntity struct {
	ID          string `gorm:"primarykey"`
	Lastname    string
	Firstname   string
	Middlename  string
	PhoneNumber string
	CityName    string
}
