package store

// Recipient struct to store mail id and name
type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// MessageInfo to store email message info
type MessageInfo struct {
	To       []Recipient `json:"to"`
	CC       []Recipient `json:"cc"`
	BCC      []Recipient `json:"bcc"`
	Subject  string      `json:"subject"`
	TextPart string      `json:"textPart"`
	HTMLPart string      `json:"htmlPart"`
}
