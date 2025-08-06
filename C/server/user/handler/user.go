package handler

import (
	"context"
	"errors"
	"math/rand"
	"regexp"
	"server/models"
	"server/pkg"
	"server/user/basic/inits"
	__ "server/user/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

// 示例
func (s *Server) Stream(_ context.Context, in *__.StreamReq) (*__.StreamResp, error) {
	return &__.StreamResp{}, nil
}

func (s *Server) SendSms(_ context.Context, in *__.SendSmsRequest) (*__.SendSmsResponse, error) {
	code := rand.Intn(9000) + 1000
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Mobile) {
		return nil, errors.New("手机号格式不正确")
	}
	get, _ := inits.RedisGet("SendSmsIncr" + in.Mobile).Int()
	if get == 1 {
		return nil, errors.New("已发送短信，60秒后可重试")
	}

	inits.RedisSet("SendSms"+in.Source+in.Mobile, code)
	incr := inits.RedisIncr("SendSmsIncr" + in.Mobile)
	if incr == 1 {
		inits.RedisExpire("SendSmsIncr" + in.Mobile)
	}

	return &__.SendSmsResponse{Greet: "验证码发送成功"}, nil
}

func (s *Server) UserLogin(_ context.Context, in *__.UserLoginRequest) (*__.UserLoginResponse, error) {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Mobile) {
		return nil, errors.New("手机号格式不正确")
	}
	get := inits.RedisGet("SendSms" + "login" + in.Mobile).Val()
	if get != in.Code {
		return nil, errors.New("短信验证码错误")
	}
	if get == "" {
		return nil, errors.New("短信验证码已过期")
	}
	defer inits.RedisDel("SendSms" + "login" + in.Mobile)
	u := &models.User{}
	user, err := u.FindUserByMobile(in.Mobile)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if user.Id == 0 {
		users, err := u.CreateUser(in.Mobile)
		if err != nil {
			return nil, errors.New("注册失败")
		}
		user = users
	}
	return &__.UserLoginResponse{Greet: int64(user.Id)}, nil
}

func (s *Server) UserPassword(_ context.Context, in *__.UserPasswordRequest) (*__.UserPasswordResponse, error) {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Mobile) {
		return nil, errors.New("手机号格式不正确")
	}
	u := &models.User{}
	user, err := u.FindUserByMobile(in.Mobile)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if user.Id != 0 {
		return nil, errors.New("还未注册，请前往注册")
	}
	if pkg.PwdMd5(in.Password) != user.Password {
		return nil, errors.New("密码错误,请重新输入")
	}
	return &__.UserPasswordResponse{Greet: int64(user.Id)}, nil
}
