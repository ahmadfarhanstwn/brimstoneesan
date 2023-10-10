package brimstoneesan

import "net/http"

func (b *Brimstoneesan) SessionLoad(next http.Handler) http.Handler {
	return b.Session.LoadAndSave(next)
}
