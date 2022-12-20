package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NoteType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNoteType(noteType string) {
	nb.NoteType = noteType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("icon kullanacaksanız geçerli bir subTitle girmelisin")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("Öncelik 0 - 5 arasında olmalı")
	}

	return &Notification{
		Title:    nb.Title,
		SubTitle: nb.SubTitle,
		Message:  nb.Message,
		Image:    nb.Image,
		Icon:     nb.Icon,
		Priority: nb.Priority,
		NoteType: nb.NoteType,
	}, nil
}
