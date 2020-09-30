package structure

type Conversations struct {
	Ok       bool `json:"ok"`
	Messages []struct {
		Type string `json:"type"`
		User string `json:"user"`
		Text string `json:"text"`
		Ts   string `json:"ts"`
	} `json:"messages"`
	HasMore          bool `json:"has_more"`
	PinCount         int  `json:"pin_count"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}
