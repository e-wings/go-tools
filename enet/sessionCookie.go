package enet

import (
	"github.com/revel/revel"
	"net/http"
	"net/url"
	"time"
)

/*SetCookie用于设置cookie值
 */
func SetCookie(name, value string, dest *revel.Response, days int) {
	ExpireTime := time.Now().AddDate(0, 0, days)
	var sessionValue string
	sessionValue = value
	sessionData := url.QueryEscape(sessionValue)
	cookie := &http.Cookie{
		Name:     name,
		Value:    sessionData,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  ExpireTime.UTC(),
	}
	http.SetCookie(dest.Out, cookie)
}
