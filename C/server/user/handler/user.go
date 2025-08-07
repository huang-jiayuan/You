package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
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
	if !phoneRegex.MatchString(in.Username) {
		return nil, errors.New("手机号格式不正确")
	}
	get, _ := inits.RedisGet("SendSmsIncr" + in.Username).Int()
	if get == 1 {
		return nil, errors.New("已发送短信，60秒后可重试")
	}

	inits.RedisSet("SendSms"+in.Source+in.Username, code)
	incr := inits.RedisIncr("SendSmsIncr" + in.Username)
	if incr == 1 {
		inits.RedisExpire("SendSmsIncr" + in.Username)
	}

	return &__.SendSmsResponse{Greet: "验证码发送成功"}, nil
}

func (s *Server) UserLogin(_ context.Context, in *__.UserLoginRequest) (*__.UserLoginResponse, error) {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Username) {
		return nil, errors.New("手机号格式不正确")
	}
	get := inits.RedisGet("SendSms" + "login" + in.Username).Val()
	if get != in.Code {
		return nil, errors.New("短信验证码错误")
	}
	if get == "" {
		return nil, errors.New("短信验证码已过期")
	}
	defer inits.RedisDel("SendSms" + "login" + in.Username)
	u := &models.User{}
	user, err := u.FindUserByMobile(in.Username)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if user.Id == 0 {
		users, err := u.CreateUser(in.Username, in.LastLoginIp)
		fmt.Println(err)
		if err != nil {
			return nil, errors.New("注册失败")
		}
		user = users
	}
	return &__.UserLoginResponse{Greet: int64(user.Id)}, nil
}

// 生成随机令牌
func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

const (
	tokenPrefix = "remember_token:" // 令牌键前缀，避免键冲突
)

func (s *Server) UserPassword(_ context.Context, in *__.UserPasswordRequest) (*__.UserPasswordResponse, error) {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Username) {
		return nil, errors.New("手机号格式不正确")
	}
	u := &models.User{}

	user, err := u.FindUserByMobile(in.Username)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if user.Id == 0 {
		return nil, errors.New("还未注册，请前往注册")
	}
	if pkg.PwdMd5(in.Password) != user.Password {
		return nil, errors.New("密码错误,请重新输入")
	}
	var token string
	if in.RememberMe {
		token = generateToken()
		// 存储令牌到Redis，设置过期时间
		key := tokenPrefix + token
		err = inits.Set(
			key,
			in.Username, // 存储用户名
		)
		if err != nil {
			return nil, errors.New("存储令牌失败")
		}
	}
	return &__.UserPasswordResponse{Greet: int64(user.Id), Token: token}, nil
}

func (s *Server) VerifyToken(_ context.Context, req *__.VerifyTokenRequest) (*__.VerifyTokenResponse, error) {
	if req.Token == "" {
		return &__.VerifyTokenResponse{Valid: false}, nil
	}
	// 从Redis查询令牌
	key := tokenPrefix + req.Token
	username, err := inits.RedisGet(key).Result()
	// 处理Redis查询错误
	if errors.Is(err, redis.Nil) {
		// 令牌不存在或已过期
		return &__.VerifyTokenResponse{Valid: false}, nil
	}
	if err != nil {
		// Redis操作出错
		return nil, fmt.Errorf("验证令牌失败: %v", err)
	}
	// 令牌有效，返回用户名
	return &__.VerifyTokenResponse{
		Valid:    true,
		Username: username,
	}, nil
}

func (s *Server) Logout(_ context.Context, req *__.LogoutRequest) (*__.LogoutResponse, error) {
	if req.Token == "" {
		return &__.LogoutResponse{
			Greet: "已登出",
		}, nil
	}
	// 从Redis删除令牌
	key := tokenPrefix + req.Token
	inits.RedisDel(key)
	return &__.LogoutResponse{
		Greet: "登出成功",
	}, nil
}

func (s *Server) UpdatePassword(_ context.Context, in *__.UpdatePasswordRequest) (*__.UpdatePasswordResponse, error) {
	u := &models.User{}
	id, err := u.FindUserById(in.UserId)
	fmt.Println(in.UserId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if pkg.PwdMd5(in.Password) == id.Password {
		return nil, errors.New("新密码与旧密码不能一致")
	}
	err = u.UpdatePassword(in.UserId, in.Password)
	if err != nil {
		return nil, errors.New("密码修改失败")
	}
	return &__.UpdatePasswordResponse{Greet: "密码修改成功"}, nil
}

func (s *Server) ImproveUserMessage(_ context.Context, in *__.ImproveUserMessageRequest) (*__.ImproveUserMessageResponse, error) {
	u := &models.User{}
	err := u.ImproveUserMessage(in)
	if err != nil {
		return nil, err
	}
	return &__.ImproveUserMessageResponse{Greet: "完善成功"}, nil
}
