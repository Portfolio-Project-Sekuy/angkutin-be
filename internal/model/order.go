package model

// Order :nodoc:
type Order struct {
	ID                  int64  `json:"id"`
	SenderName          string `json:"senderName"`
	SenderPhoneNumber   string `json:"senderPhoneNumber"`
	ReceiverName        string `json:"receiverName"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber"`
}

// IsEntity implement gqlgen fedruntime Entity
func (Order) IsEntity() {}
