package core

type Dialog struct {
	Id                  int       `json:"id" db:"id"`
	User1Id             int       `json:"user1_id" db:"user1_id"`
	User2Id             int       `json:"user2_id" db:"user2_id"`
	Banned              bool      `json:"banned" db:"banned"`
	Companion           string    `json:"companion"`
	CompanionImagePaths *[]string `json:"сompanion_image_paths"`
	LastMessage         *Message  `json:"last_message"`
}

func (d *Dialog) Sanitize() {
	d.Companion = policy.Sanitize(d.Companion)
}
