package domain

// PersistenceAPIClient if the interface definition to any persistence system that want to adapt MyWins API.
type PersistenceAPIClient interface {
	GetWins() ([]*Entry, error)
	GetFails() ([]*Entry, error)
	AddWin() error
	AddFail() error
}
