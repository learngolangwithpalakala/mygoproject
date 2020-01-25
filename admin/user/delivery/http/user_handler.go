package http

import (
	"context"
	"encoding/json"
	"github.com/bxcodec/go-clean-arch/admin/user"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var store = sessions.NewCookieStore([]byte("secret"))

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Status      string      `json:"status"`
	ErrorCode   string      `json:"errorCode"`
	Description string      `json:"description"`
	SessionKey  string      `json:"sessionKey"`
	Role        models.Role `json:"role"`
	Email       string      `json:"email"`
	EmpId       string      `json:"empId"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Position    string      `json:"position"`
	Timestamp   string      `json:"timestamp"`
}

// UserHandler  represent the httphandler for article
type UserHandler struct {
	Userusecase user.Usecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, us user.Usecase) {
	handler := &UserHandler{
		Userusecase: us,
	}

	e.GET("/user", handler.FetchUsers)
	e.POST("/user", handler.Store)
	e.PUT("/user", handler.Update)
	e.GET("/user/:id", handler.GetByEmpNumber)
	e.DELETE("/user/:id", handler.Delete)
	e.POST("/login", handler.DoLogin)

}

// Store will store the user by given request body
func (a *UserHandler) Store(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.Userusecase.Store(ctx, &user)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

//Update
func (a *UserHandler) Update(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.Userusecase.Update(ctx, &user)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

// Delete will delete user by given param
func (a *UserHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	id := strconv.Itoa(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.Userusecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// GetByEmpNumber will get user by given emp number
func (a *UserHandler) GetByEmpNumber(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := strconv.Itoa(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := a.Userusecase.GetByEmpNumber(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func isRequestValid(m *models.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// FetchUsers will fetch the users based on given params
func (a *UserHandler) FetchUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listUsrs, err := a.Userusecase.Fetch(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listUsrs)
}

func (a *UserHandler) DoLogin(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err = a.Userusecase.Login(ctx, user)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	session, err := store.Get(c.Request(), "session-name")
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return err
	}
	// Set some session values.
	session.Values["foo"] = "bar"
	// Save it before we write to the response/return from the handler.
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return err
	}

	response := &LoginResponse{Status: http.StatusText(200), ErrorCode: "", Description: "", SessionKey: session.ID, Role: models.Role{ID: 3, Role: "ADMIN"}, Email: user.WorkEmail, EmpId: user.EmpNumber,
		FirstName: user.FirstName, LastName: user.LastName, Position: user.Position}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(response)
}
