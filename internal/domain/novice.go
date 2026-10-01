package domain

const (
	NoviceLimit = 6
)

// Интерфейс, описывающий, что предмет имеет уровень простоты для новичков
type NoviceRated interface {
	Simplicity() int
}

// Можно ли новичку
func IsNoviceFriendly(n NoviceRated) bool {
	if n.Simplicity() >= NoviceLimit {
		return true
	}
	return false
}
