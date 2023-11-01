package link

type Link struct {
	ID        string `json:"ID"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
	Clicked   uint32 `json:"clicked"`
	Random    bool   `json:"random"`
}
