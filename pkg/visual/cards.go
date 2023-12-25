package visual

import (
	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
)

type Cards map[entity.CardId]string

func (c Cards) Ids() []entity.CardId {
	r := make([]entity.CardId, 0, len(c))
	for id := range c {
		r = append(r, id)
	}
	return r
}
