package service

import (
	"ccloud/web/constant"
	"ccloud/web/dao"
	"ccloud/web/entity"
	"ccloud/web/entity/cmd"
	"ccloud/web/entity/dto"
	"ccloud/web/log"
	"ccloud/web/security"
	"ccloud/web/token"
	"errors"
	"strconv"
	"time"
)

type LoginService interface {
	Login(cmd *cmd.LoginCmd) entity.Response
}

type loginserviceimpl struct {
	accountdao      dao.AccountDao
	passwordencoder security.PasswordEncoder
	tokenservice    token.AuthTokenService
}

func NewLoginService() (LoginService, error) {
	accountdao, err := dao.NewAccountDao()
	if err != nil {
		return nil, err
	}

	return &loginserviceimpl{
		accountdao:      accountdao,
		passwordencoder: security.NewBcryptPasswordEncoder(),
		tokenservice:    token.NewAuthTokenService(),
	}, nil
}

func (impl loginserviceimpl) Login(cmd *cmd.LoginCmd) entity.Response {
	// 参数校验
	if cmd.User == "" || cmd.Pass == "" {
		return entity.Fail(constant.ParamErr)
	}

	// 根据用户名从数据库查询用户
	user, err := impl.accountdao.SelectByUsername(cmd.User)
	if err != nil {
		log.Logger.Warnf("username not found, username: %v", cmd.User)
		return entity.Fail(constant.UsernameNotFound)
	}

	// 密码校验
	matches := impl.passwordencoder.Matches(cmd.User, user.Password)
	if matches {
		token, err := impl.createtoken(user.Id)
		if err != nil {
			return entity.Fail(constant.TokenGenerateFailed)
		}

		logindto := dto.LoginDto{
			Id:    user.Id,
			User:  cmd.User,
			Token: token,
		}

		return entity.OKWithData(logindto)
	}

	return entity.Fail(constant.PasswordMismatch)
}

func (impl loginserviceimpl) createtoken(userId int64) (string, error) {
	var payload map[string]string = map[string]string{}
	payload["userId"] = strconv.FormatInt(userId, 10)
	// 颁发Token
	tokenpalyload := token.TokenPayload{
		Expire:  time.Now().Add(7 * 24 * time.Hour),
		Payload: payload,
	}
	token := impl.tokenservice.CreateToken(tokenpalyload)
	if token != "" {
		return token, nil
	}

	return "", errors.New("create access token error")
}
