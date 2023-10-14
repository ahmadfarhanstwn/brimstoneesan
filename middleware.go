package brimstoneesan

import "net/http"

func (b *Brimstoneesan) SessionLoad(next http.Handler) http.Handler {
	b.InfoLog.Println("Session Load!")
	return b.Session.LoadAndSave(next)
}
