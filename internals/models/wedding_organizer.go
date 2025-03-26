package models

type (
	WeddingOrganizer struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		ContactWO *ContactWO `json:"contact_wo"`
	}

	ContactWO struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
		LinkGMaps string `json:"link_gmaps"`
	}
)
