package helper

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	id := uuid.New()
	uuid := strings.Replace(id.String(), "-", "", -1)
	return uuid
}
