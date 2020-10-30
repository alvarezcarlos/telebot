package notigram

type MessageInfo struct {
	AppName string `json:"app_name"`
	Message string `json:"message"`
	ChatId  string `json:"chat_id"`
}
