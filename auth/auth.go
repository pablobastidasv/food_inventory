package auth

import (
	"context"
	"encoding/gob"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
)

const sessionName = "_bastriguez_session"
const userCookieValueName = "user"

type AuthHandler struct {
	clientId     string
	clientSecret string
	callbackUrl  string
	returnUrl    string
	domain       string
	loginUri     string
}

type User struct {
	UserId    string
	Name      string
	FirstName string
	LastName  string
	AvatarURL string
}

// New creates a instance with all Handlers for the auth, it receives the URL where the login page is invoked
func New(loginUrl string) *AuthHandler {
	gob.Register(User{})

	ah := &AuthHandler{
		clientId:     os.Getenv("AUTH_CLIENT_ID"),
		clientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		callbackUrl:  os.Getenv("AUTH_CALLBACK_URL"),
		returnUrl:    os.Getenv("AUTH_REDIRECT_URL"), // Redirect after logout
		domain:       os.Getenv("AUTH_DOMAIN"),
		loginUri:     loginUrl,
	}

	auth0Provider := auth0.New(ah.clientId, ah.clientSecret, ah.callbackUrl, ah.domain)
	goth.UseProviders(auth0Provider)

	return ah
}

func (ah *AuthHandler) GetLogin(c echo.Context) error {
	q := c.Request().URL.Query()
	q.Add("provider", "auth0")
	c.Request().URL.RawQuery = q.Encode()

	req := c.Request()
	res := c.Response().Writer
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		return c.JSON(http.StatusOK, gothUser)
	}
	gothic.BeginAuthHandler(res, req)
	return nil
}

func (ah *AuthHandler) GetCallback(c echo.Context) error {
	req := c.Request()
	res := c.Response().Writer

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := createUserSession(c, user); err != nil {
		return err
	}

	// Redirect
	return c.Redirect(http.StatusSeeOther, "/")
}

func (ah *AuthHandler) GetLogout(c echo.Context) error {
	q := c.Request().URL.Query()
	q.Add("provider", "auth0")
	c.Request().URL.RawQuery = q.Encode()

	req := c.Request()
	res := c.Response().Writer

	gothic.Logout(res, req)

	// remove custom cookie
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	sess.Values = make(map[interface{}]interface{})
	err = sess.Save(req, res)
	if err != nil {
		return err
	}

	// Close auth0 session
	logoutUrl, err := url.Parse("https://" + ah.domain + "/v2/logout")
	if err != nil {
		return err
	}

	parameters := url.Values{}
	parameters.Add("returnTo", ah.returnUrl)
	parameters.Add("client_id", ah.clientId)
	logoutUrl.RawQuery = parameters.Encode()

	return c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

func (ah *AuthHandler) PageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Debug("Page authentication middleware")

		sess, err := session.Get(sessionName, c)
		if err != nil {
			return err
		}

		user := sess.Values[userCookieValueName]
		if user == nil {
			return c.Redirect(http.StatusTemporaryRedirect, ah.loginUri)
		}

		// Add user to the context
		addUserToContext(c, user.(User))

		return next(c)
	}
}

// addUserToContext adds the given user to che given context
func addUserToContext(c echo.Context, user User){
	ctx := context.WithValue(c.Request().Context(), "user", user)
	r := c.Request().WithContext(ctx)
	c.SetRequest(r)
}

func (ah *AuthHandler) FragmentMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Debug("Fragment authentication middleware")

		sess, err := session.Get(sessionName, c)
		if err != nil {
			return err
		}

		user := sess.Values[userCookieValueName]
		if user == nil {
			c.Response().Header().Add("HX-Redirect", ah.loginUri)
			return nil
		}

		// Add user to the context
		addUserToContext(c, user.(User))

		return next(c)
	}
}

func GetUser(c echo.Context) (User, error) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return User{}, err
	}
	u := sess.Values[userCookieValueName]

	return u.(User), nil
}

func createUserSession(c echo.Context, user goth.User) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   c.IsTLS(),
	}

	u := User{
		UserId:    user.UserID,
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		AvatarURL: user.AvatarURL,
	}
	sess.Values[userCookieValueName] = u
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}
