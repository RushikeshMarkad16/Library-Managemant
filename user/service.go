package user

import (
	"context"
	"time"

	"github.com/RushikeshMarkad16/Library-Managemant/db"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	Create(ctx context.Context, req User) (err error)
	FindByID(ctx context.Context, id string) (response FindByIDResponse, err error)
	FindByData(ctx context.Context, filterData string) (respose FindByDataResponse, err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req User) (err error)
	UpdatePassword(ctx context.Context, req ChangePassword) (err error)
	GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("jsd549$^&")

func (cs *userService) GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error) {

	// var cs *userService
	user, err := cs.store.FindUserByEmail(ctx, Email)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return "", errWrongEmail
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "email", Email)
		return
	}

	if !CheckPasswordHash(Password, user.Password) {
		return "", errWrongPassword
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (cs *userService) List(ctx context.Context) (response ListResponse, err error) {
	dbUsers, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	for _, dbUser := range dbUsers {
		var userData User
		userData.ID = dbUser.ID
		userData.FirstName = dbUser.First_name
		userData.Last_name = dbUser.Last_name
		userData.Gender = dbUser.Gender
		userData.Address = dbUser.Address
		userData.Email = dbUser.Email
		userData.Password = dbUser.Password
		userData.Mob_no = dbUser.Mob_no
		userData.Role = dbUser.Role

		response.Users = append(response.Users, userData)
	}

	return
}

func (cs *userService) Create(ctx context.Context, c User) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for user create", "msg", err.Error(), "user", c)
		return
	}

	uuidgen := uuid.New()
	c.ID = uuidgen.String()

	err = cs.store.CreateUser(ctx, &db.User{
		ID:         c.ID,
		First_name: c.FirstName,
		Last_name:  c.Last_name,
		Gender:     c.Gender,
		Address:    c.Address,
		Email:      c.Email,
		Password:   c.Password,
		Mob_no:     c.Mob_no,
		Role:       c.Role,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}

	return
}

func (cs *userService) Update(ctx context.Context, c User) (err error) {
	err = c.ValidateUpdate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		ID:         c.ID,
		First_name: c.FirstName,
		Last_name:  c.Last_name,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *userService) UpdatePassword(ctx context.Context, c ChangePassword) (err error) {

	err = cs.store.UpdatePassword(ctx, &db.User{
		ID:       c.ID,
		Password: c.NewPassword,
	})
	if err != nil {
		cs.logger.Error("Error updating Password", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *userService) FindByID(ctx context.Context, id string) (response FindByIDResponse, err error) {
	user, err := cs.store.FindUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
		return
	}

	response.User.ID = user.ID
	response.User.FirstName = user.First_name
	response.User.Last_name = user.Last_name
	response.User.Gender = user.Gender
	response.User.Address = user.Address
	response.User.Email = user.Email
	response.User.Password = user.Password
	response.User.Mob_no = user.Mob_no
	response.User.Role = user.Role

	return
}

func (cs *userService) FindByData(ctx context.Context, filterData string) (response FindByDataResponse, err error) {
	user, err := cs.store.FindUserByData(ctx, filterData)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "filterData", filterData)
		return
	}

	response.User.ID = user.ID
	response.User.FirstName = user.First_name
	response.User.Last_name = user.Last_name
	response.User.Gender = user.Gender
	response.User.Address = user.Address
	response.User.Email = user.Email
	response.User.Password = user.Password
	response.User.Mob_no = user.Mob_no
	response.User.Role = user.Role

	return
}

func (cs *userService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("User Not present", "err", err.Error(), "id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
