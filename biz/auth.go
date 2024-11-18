package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/crypto"
)

type UserRepo interface {
	// Create 创建用户
	Create(ctx context.Context, user *User) error

	// FindByName 根据用户名查询
	FindByName(ctx context.Context, name string) (*User, error)
}

type CaptchaRepo interface {
	// Generate 生成验证码
	Generate(ctx context.Context) (*Captcha, error)

	// Verify 验证验证码是否正确
	Verify(ctx context.Context, id string, answer string) (bool, error)
}

type JwtRepo interface {
	// NewToken 生成 accessToken 和 refreshToken
	NewToken(ctx context.Context, claims CustomClaims, accessExpiration, refreshExpiration time.Duration) (*Token, error)

	// GetUserIdFromRefreshToken 根据 refreshToken 获取 userId
	GetUserIdFromRefreshToken(ctx context.Context, refreshToken string) (uint64, error)

	// DeleteRefreshToken 删除 RefreshToken
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
}

type AuthUsecase struct {
	conf        *conf.Auth
	userRepo    UserRepo
	captchaRepo CaptchaRepo
	jwtRepo     JwtRepo
	log         *log.Helper
}

func NewAuthUsecase(conf *conf.Auth, userRepo UserRepo, captchaRepo CaptchaRepo, tokenRepo JwtRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{
		conf:        conf,
		userRepo:    userRepo,
		captchaRepo: captchaRepo,
		jwtRepo:     tokenRepo,
		log:         log.NewHelper(logger),
	}
}

type User struct {
	Id       uint64
	Name     string
	Password string
}

// Register 用户注册
func (uc *AuthUsecase) Register(ctx context.Context, user *User) error {
	_, err := uc.userRepo.FindByName(ctx, user.Name)
	if err == nil {
		return errors.ErrorUserNameExist("用户名已存在")
	}
	if !errors.IsUserNotFound(err) {
		return err
	}

	// 密码加密
	user.Password, err = crypto.BcryptMake([]byte(user.Password))
	if err != nil {
		return err
	}

	return uc.userRepo.Create(ctx, user)
}

type Token struct {
	AccessToken            string `json:"accessToken"`            // 访问token
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`  // 访问token的过期时间
	RefreshToken           string `json:"refreshToken"`           // 刷新token
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"` // 刷新token的过期时间
}

// Captcha 验证码
type Captcha struct {
	Id        string
	Base64Img string
	Answer    string
}

// GetCaptcha 获取验证码
func (uc *AuthUsecase) GetCaptcha(ctx context.Context) (*Captcha, error) {
	return uc.captchaRepo.Generate(ctx)
}

type Login struct {
	User User
	Cap  Captcha
}

// Login 用户登录
func (uc *AuthUsecase) Login(ctx context.Context, login Login) (*Token, error) {
	ok, err := uc.captchaRepo.Verify(ctx, login.Cap.Id, login.Cap.Answer)
	if err != nil {
		return nil, errors.ErrorUserCaptchaVerifyError("验证码错误")
	}
	if !ok {
		return nil, errors.ErrorUserCaptchaVerifyError("验证码错误")
	}
	u, err := uc.userRepo.FindByName(ctx, login.User.Name)
	if err != nil {
		return nil, errors.ErrorUserNameOrPasswordError("用户名密码错误")
	}
	if ok := crypto.BcryptMakeCheck([]byte(login.User.Password), u.Password); !ok {
		return nil, errors.ErrorUserNameOrPasswordError("用户名密码错误")
	}
	return uc.newToken(ctx, u.Id)
}

func (uc *AuthUsecase) newToken(ctx context.Context, userId uint64) (*Token, error) {
	claims := CustomClaims{UserId: userId}
	return uc.jwtRepo.NewToken(ctx, claims, uc.conf.AccessExpiration.AsDuration(), uc.conf.RefreshExpiration.AsDuration())
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	UserId uint64
	jwt.RegisteredClaims
}

// Logout 用户登出
func (uc *AuthUsecase) Logout(ctx context.Context, refreshToken string) error {
	return uc.jwtRepo.DeleteRefreshToken(ctx, refreshToken)
}

// RefreshToken 刷新令牌
func (uc *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	userId, err := uc.jwtRepo.GetUserIdFromRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if err := uc.jwtRepo.DeleteRefreshToken(ctx, refreshToken); err != nil {
		return nil, err
	}
	return uc.newToken(ctx, userId)
}
