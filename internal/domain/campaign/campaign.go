package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignID string `gorm:"size:50"`
}

const (
	Pending = "Pending"
	Started = "Started"
	Done    = "Done"
)

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].ID = xid.New().String()
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}

	return nil, err
}
