package user

type User struct {
	ID         string     `json:"id"`
	AccountID  string     `json:"account_id"`
	Name       string     `json:"name"`
	Age        uint       `json:"age"`
	Gender     string     `json:"gender"`
	Photos     []Photo    `json:"photos"`
	Interests  []Interest `json:"interests"`
	Longitude  float64    `json:"longitude"`
	Latitude   float64    `json:"latitude"`
	Preference Preference `json:"preference"`
}

type Photo struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Interest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Preference struct {
	AgeMin   uint   `json:"age_min"`
	AgeMax   uint   `json:"age_max"`
	Gender   string `json:"gender"`
	Distance uint   `json:"distance"`
}

type UpdateCurrentUserRequest struct {
	Name         string   `json:"name"`
	Age          uint     `json:"age"`
	Gender       string   `json:"gender"`
	Photos       []string `json:"photos"`
	Interests    []int    `json:"interests"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	AgeMin       uint     `json:"age_min"`
	AgeMax       uint     `json:"age_max"`
	GenderFilter string   `json:"gender_filter"`
	Distance     uint     `json:"distance"`
}
