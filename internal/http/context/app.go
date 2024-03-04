package context

import (
	context2 "context"
	"net/http"

	"github.com/gorilla/sessions"
)

type App struct {
	SessionsStore sessions.Store
	Request       *http.Request
	Response      http.ResponseWriter
}

func (context *App) getSession() (*sessions.Session, error) {
	return context.SessionsStore.Get(context.Request, "user")
}

func (context *App) IsAuth() bool {
	if context.SessionsStore != nil {
		session, err := context.getSession()
		if err != nil {
			return false
		}

		_, ok := session.Values["userId"]
		return ok
	}

	return false
}

func (context *App) GetUserId() string {
	if context.SessionsStore != nil {
		session, err := context.getSession()
		if err != nil {
			panic(err)
		}

		userId, ok := session.Values["userId"]
		if !ok {
			panic("No userid")
		}
		return userId.(string)
	}

	panic("No session store")
}

func (context *App) SaveRedirectUrl(url string) {
	session, err := context.getSession()
	if err != nil {
		panic(err)
	}

	session.Values["redirectUrl"] = url
	err = session.Save(context.Request, context.Response)
	if err != nil {
		panic(err)
	}
}

func (context *App) Login(userId string) {

	session, err := context.getSession()
	if err != nil {
		panic(err)
	}

	session.Values["userId"] = userId

	err = session.Save(context.Request, context.Response)
	if err != nil {
		panic(err)
	}
}

func (context *App) GetRedirectUrl() string {

	session, err := context.getSession()
	if err != nil {
		panic(err)
	}

	url, ok := session.Values["redirectUrl"]

	if ok {
		delete(session.Values, "redirectUrl")
	} else {
		url = "/"
	}

	err = session.Save(context.Request, context.Response)
	if err != nil {
		panic(err)
	}

	return url.(string)
}

func (context *App) Logout() {

	session, err := context.getSession()
	if err != nil {
		return
	}

	delete(session.Values, "userId")

	err = session.Save(context.Request, context.Response)
	if err != nil {
		panic(err)
	}
}

func (context *App) IsHxRequest() bool {
	return context.Request.Header.Get("HX-Request") == "true"
}

func (context *App) Url() string {
	return context.Request.URL.Path
}

func GetApp(ctx context2.Context) *App {
	return (ctx.Value("app")).(*App)
}

func SetApp(sessionStore sessions.Store, r *http.Request, w http.ResponseWriter) *http.Request {
	appContext := &App{SessionsStore: sessionStore, Request: r, Response: w}
	return r.WithContext(context2.WithValue(r.Context(), "app", appContext))
}
