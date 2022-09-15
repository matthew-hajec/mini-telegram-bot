package tgbot

type TelegramMessageHandler func(update *TelegramUpdateResult)

type TelegramUser struct {
	Id                    int64  `json:"id"`
	IsBot                 bool   `json:"is_bot"`
	Firstname             string `json:"first_name"`
	Lastname              string `json:"last_name"`
	Username              string `json:"username"`
	LanguageCode          string `json:"language_code"`
	IsPremimum            bool   `json:"is_premium"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu"`
}

type TelegramChat struct {
	Id        int64  `json:"id"`
	ChatType  string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

type TelegramMessageEntity struct {
	EntityType string        `json:"type"`
	Offset     int           `json:"offset"`
	Length     int           `json:"length"`
	URL        string        `json:"url"`
	User       *TelegramUser `json:"user"`
}

type TelegramMessage struct {
	Id                   int                      `json:"message_id"`
	From                 *TelegramUser            `json:"from"`
	Date                 int                      `json:"date"`
	Chat                 *TelegramChat            `json:"chat"`
	ForwardFromMessageId int                      `json:"forward_from_message_id"`
	ViaBot               bool                     `json:"via_bot"`
	EditDate             int                      `json:"edit_date"`
	Text                 string                   `json:"text"`
	Entities             []*TelegramMessageEntity `json:"entities"`
}

type TelegramUpdate struct {
	ID            int              `json:"update_id"`
	Message       *TelegramMessage `json:"message"`
	EditedMessage *TelegramMessage `json:"edited_message"`
}

type TelegramUpdateResult struct {
	Ok      bool              `json:"ok"`
	Results []*TelegramUpdate `json:"result"`
}

type TelegramOutgoingMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramOKResponse struct {
	Ok bool `json:"ok"`
}
