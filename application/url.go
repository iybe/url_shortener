package application

type urlApp struct{}

func NewUrlApp() *urlApp {
	return &urlApp{}
}

type UrlAppInterface interface {
	ShortenUrl() (string, error)
}

var _ UrlAppInterface = &urlApp{}

func (u *urlApp) ShortenUrl() (string, error) {
	return "", nil
}
