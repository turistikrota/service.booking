package invite

import (
	"fmt"

	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/config"
	"github.com/turistikrota/service.shared/helper"
)

type Events interface {
	Invite(event InviteEvent)
}

type InviteEvent struct {
	Locale     string
	Email      string
	InviteUUID string
	UserName   string
}

type inviteEvents struct {
	publisher events.Publisher
	topics    config.Topics
	i18n      *i18np.I18n
}

type EventConfig struct {
	Publisher events.Publisher
	Topics    config.Topics
	I18n      *i18np.I18n
}

func NewEvents(cnf EventConfig) Events {
	return &inviteEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
		i18n:      cnf.I18n,
	}
}

func (e inviteEvents) Invite(event InviteEvent) {
	subject := e.i18n.Translate(I18nMessages.InviteSubject, event.Locale)
	template := fmt.Sprintf("booking/invite.%s", event.Locale)
	_ = e.publisher.Publish(e.topics.Notify.SendMail, helper.Notify.BuildEmail(event.Email, subject, i18np.P{
		"UserName":   event.UserName,
		"InviteUUID": event.InviteUUID,
	}, event.Email, template))
}
