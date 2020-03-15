package api_types

type Menus struct {
	HomeInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
		Image string `json:"image"`
	} `json:"logoInfo"`
	MenuInfo []struct {
		Title  string `json:"title"`
		Href   string `json:"href"`
		Icon   string `json:"icon"`
		Target string `json:"target"`
		Child  []struct {
			Title  string `json:"title"`
			Href   string `json:"href"`
			Icon   string `json:"icon"`
			Target string `json:"target"`
			Child  []struct {
				Title  string `json:"title"`
				Href   string `json:"href"`
				Icon   string `json:"icon"`
				Target string `json:"target"`
			} `json:"child"`
		} `json:"child"`
	} `json:"menuInfo"`
}
