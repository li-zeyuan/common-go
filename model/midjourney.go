package model

type MjBaseResponse struct {
	Code        int       `json:"code"`
	Description string    `json:"description"`
	Properties  DiscordPP `json:"properties"`
	Result      string    `json:"result"`
}

type DiscordPP struct {
	DiscordChannelId  string `json:"discordChannelId"`
	DiscordInstanceId string `json:"discordInstanceId"`
}

type ImagineReq struct {
	Base64Array []string `json:"base64Array"`
	Modes       []string `json:"modes"`
	Prompt      string   `json:"prompt"`
	Remix       bool     `json:"remix"`
}

type ActionReq struct {
	TaskId     string `json:"taskId"`
	CustomId   string `json:"customId"`
	NotifyHook string `json:"notifyHook"`
	State      string `json:"state"`
}

type FetchTasksReq struct {
	Ids []string `json:"ids"`
}

type FetchTasksResp struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	FailReason  string    `json:"failReason"`
	ImageUrl    string    `json:"imageUrl"`
	Progress    string    `json:"progress"`
	Status      string    `json:"status"`
	Action      string    `json:"action"`
	Buttons     []*Button `json:"buttons"`
}

type Button struct {
	CustomId string `json:"customId"`
	Emoji    string `json:"emoji"`
	Label    string `json:"label"`
	Style    int    `json:"style"`
	Type     int    `json:"type"`
}
