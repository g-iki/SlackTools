package structure

type Users struct {
	Ok      bool `json:"ok"`
	Members []struct {
		ID       string `json:"id"`
		TeamID   string `json:"team_id"`
		Name     string `json:"name"`
		Deleted  bool   `json:"deleted"`
		Color    string `json:"color"`
		RealName string `json:"real_name"`
		Tz       string `json:"tz"`
		TzLabel  string `json:"tz_label"`
		TzOffset int    `json:"tz_offset"`
		Profile  struct {
			AvatarHash            string `json:"avatar_hash"`
			StatusText            string `json:"status_text"`
			StatusEmoji           string `json:"status_emoji"`
			RealName              string `json:"real_name"`
			DisplayName           string `json:"display_name"`
			RealNameNormalized    string `json:"real_name_normalized"`
			DisplayNameNormalized string `json:"display_name_normalized"`
			Email                 string `json:"email"`
			Image24               string `json:"image_24"`
			Image32               string `json:"image_32"`
			Image48               string `json:"image_48"`
			Image72               string `json:"image_72"`
			Image192              string `json:"image_192"`
			Image512              string `json:"image_512"`
			Image1024             string `json:"image_1024"`
			Team                  string `json:"team"`
			ImageOriginal         string `json:"image_original"`
			FirstName             string `json:"first_name"`
			LastName              string `json:"last_name"`
			Title                 string `json:"title"`
			Phone                 string `json:"phone"`
			Skype                 string `json:"skype"`
		} `json:"profile,omitempty"`
		IsAdmin           bool `json:"is_admin"`
		IsOwner           bool `json:"is_owner"`
		IsPrimaryOwner    bool `json:"is_primary_owner"`
		IsRestricted      bool `json:"is_restricted"`
		IsUltraRestricted bool `json:"is_ultra_restricted"`
		IsBot             bool `json:"is_bot"`
		Updated           int  `json:"updated"`
		IsAppUser         bool `json:"is_app_user,omitempty"`
		Has2Fa            bool `json:"has_2fa"`
	} `json:"members"`
	CacheTs          int `json:"cache_ts"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}
