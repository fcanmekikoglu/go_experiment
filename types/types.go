package types

type User struct {
	ID        string `json:"_id"`
	Version   string `json:"_v"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
	Name      struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Google struct {
		ID           string `json:"id"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"google"`
}

type Fact struct {
	ID        string `json:"_id"`
	Version   int    `json:"_v"`
	User      string `json:"user"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Deleted   bool   `json:"deleted"`
	Source    string `json:"source"`
	Type      string `json:"type"`
	Text      string `json:"text"`
	Status    struct {
		Verified  bool   `json:"verified"`
		SentCount int    `json:"sentCount"`
		Feedback  string `json:"feedback"`
	} `json:"status"`
}
