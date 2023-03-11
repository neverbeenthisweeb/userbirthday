package inmemory

import (
	"context"
	"errors"
	"sync"
)

var (
	errTemplateNotFound = errors.New("template not found")
)

var storage map[string]string = map[string]string{
	"email.birthday": "",
	"wa.birthday":    "",
}

type NotificationTemplateRepository struct {
	mu sync.Mutex
}

// FIXME: Test later after notification impl
func (nt *NotificationTemplateRepository) GetNotificationTemplate(ctx context.CancelFunc, ID string) (string, error) {
	nt.mu.Lock()
	defer nt.mu.Unlock()
	v, ok := storage[ID]
	if !ok {
		return "", errTemplateNotFound
	}
	return v, nil
}
