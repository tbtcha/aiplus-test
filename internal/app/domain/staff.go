package domain

type StaffRequest struct {
	Lastname    string `json:"lastname" example:"Azamatov"`
	Firstname   string `json:"firstname" example:"Azamat"`
	Middlename  string `json:"middlename" example:"Azamatuly"`
	PhoneNumber string `json:"phoneNumber" example:"8(777)777-77-77"`
	CityName    string `json:"cityName" example:"Almaty"`
}

type StaffResponse struct {
	ID          string `json:"id" example:"bbb7b76f-7bb5-11ee-9ef7-9828a62c67e0"`
	Lastname    string `json:"lastname" example:"Azamatov"`
	Firstname   string `json:"firstname" example:"Azamat"`
	Middlename  string `json:"middlename" example:"Azamatuly"`
	PhoneNumber string `json:"phoneNumber" example:"8(777)777-77-77"`
	CityName    string `json:"cityName" example:"Almaty"`
}
