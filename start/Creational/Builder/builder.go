package main


type NotificationBuilder struct {
	Title string
	Subtitle string
	Message string
	Image string
	Icon string
	Priority int
	NotType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType= notType
}

// The build method returns fully finished notification object

func (nb *NotificationBuilder) Build() (*Notification, error) {
	// TODO: Error checking can be done in the build stage

	// TODO: Returns a newly created notification object

	return &Notification{
		title: nb.Title,
		subtitle: nb.Subtitle,
		message: nb.Message,
		image: nb.Image,
		icon: nb.Icon,
		priority: nb.Priority,
		notType: nb.NotType,
	}, nil
}

