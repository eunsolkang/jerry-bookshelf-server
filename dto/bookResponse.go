package dto

type BookResponse struct {
	Id        string  `json:"id`
	Name      string  `json:"name"`
	Read_date string  `json:"read_date"`
	Report    string  `json:"report"`
	Rating    float64 `json:"rating"`
	Img_url   string  `json:"img_url"`
}

