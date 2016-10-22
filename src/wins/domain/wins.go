package domain

type win struct {
	Success []string `json:"wins"`
	Fails   []string `json:"defeats"`
}

func (w *win) getSuccess() []string {
	return w.Success
}

func (w *win) getFails() []string {
	return w.Fails
}
