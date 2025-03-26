package authors

type Authors struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

type RequestAuthorsDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *RequestAuthorsDTO) ToAuthors() Authors {
	return Authors{
		ID:   r.ID,
		Name: r.Name,
	}
}

type ResponseAuthorsDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Authors) ToResponseAuthorsDTO() ResponseAuthorsDTO {
	return ResponseAuthorsDTO{
		ID:   r.ID,
		Name: r.Name,
	}
}
