package validators

import (
	"errors"
	"github.com/urcop/go-fiber-template/internal/model"
	"strings"
)

func LinkValidator(link *model.Link) error {
	if !(strings.HasPrefix(link.Link, "https://") || strings.HasPrefix(link.Link, "http://")) {
		return errors.New("link must begin with https:// or http://")
	}

	return nil
}
