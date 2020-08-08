package op

type Item struct {
	Uuid         string       `json:"uuid,omitempty"`
	TemplateUuid string       `json:"templateUuid,omitempty"`
	Trashed      string       `json:"trashed,omitempty"`
	CreatedAt    string       `json:"createdAt,omitempty"`
	UpdatedAt    string       `json:"updatedAt,omitempty"`
	ChangerUuid  string       `json:"changerUuid,omitempty"`
	ItemVersion  int          `json:"itemVersion,omitempty"`
	VaultUuid    string       `json:"vaultUuid,omitempty"`
	Overview     itemOverview `json:"overview,omitempty"`
}

type itemOverview struct {
	URLs []struct {
		L string `json:"l,omitempty"`
		U string `json:"u,omitempty"`
	} `json:"URLs"`
	Ainfo  string `json:"ainfo,omitempty"`
	Appids []struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"appids"`
	Pbe   float32  `json:"pbe,omitempty"`
	Pgrng bool     `json:"pgrng,omitempty"`
	Ps    float32  `json:"ps,omitempty"`
	Tags  []string `json:"tags"`
	Title string   `json:"title,omitempty"`
	Url   string   `json:"url,omitempty"`
}

type ItemWithDetails struct {
	Item
	Details itemDetails
}

type itemDetails struct {
	Fields []struct {
		Designation string `json:"designation,omitempty"`
		Name        string `json:"name,omitempty"`
		Type        string `json:"type,omitempty"`
		Value       string `json:"value,omitempty"`
	} `json:"fields"`
	NotesPlain      string `json:"notesPlain"`
	PasswordHistory []struct {
		Time  int    `json:"time"`
		Value string `json:"value"`
	}
	// Sections []struct{} `json:"sections"` // TODO parsing of sections
}
