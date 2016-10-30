package domain

type PersistenceApiClient interface {
	GetWins() ([]*Entry, error)
	GetFails() ([]*Entry, error)
	AddWin() error
	AddFail() error
}

