package service

import (
	"backend/internal/config"
	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/forms/cms_sys_form"
	"backend/internal/repositories/sys_repositories"
	"backend/pkg/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IUserService interface {
	Get(uid int64) (*entity.SysUser, error)
	FindByAccount(account string) (*entity.SysUser, error)
	FindByPhone(phone string) (*entity.SysUser, error)
	FindAllPager(searchKey string, page, size int) ([]entity.SysUser, int64, error)
	Insert(data *entity.SysUser) error
	Update(data *entity.SysUser, musColumns ...string) error
	Save(data *entity.SysUser, musColumns ...string) error
	CreateUser(data *cms_sys_form.CreateAdminUserForm) (*entity.SysUser, error)
	GenerateToken(id int64, username string) (string, error)
	AdminLogin(param cms_sys_form.AdminLoginParam) (*dto.PwdLoginDTO, error)
}

type UserService struct {
	repo sys_repositories.IUserRepository
}

func NewUserService() IUserService {
	return &UserService{
		repo: sys_repositories.NewUserRepository(),
	}
}

func (s *UserService) Get(uid int64) (*entity.SysUser, error) {
	return s.repo.Get(uid)
}

func (s *UserService) FindByAccount(account string) (*entity.SysUser, error) {
	return s.repo.FindByAccount(account)
}

func (s *UserService) FindByPhone(phone string) (*entity.SysUser, error) {
	return s.repo.FindByPhone(phone)
}

func (s *UserService) FindAllPager(searchKey string, page, size int) ([]entity.SysUser, int64, error) {
	return s.repo.FindAllPager(searchKey, page, size)
}

func (s *UserService) Insert(data *entity.SysUser) error {
	return s.repo.Insert(data)
}
func (s *UserService) Update(data *entity.SysUser, musColumns ...string) error {
	return s.repo.Update(data, musColumns...)
}
func (s *UserService) Save(data *entity.SysUser, musColumns ...string) error {
	return s.repo.Save(data, musColumns...)
}

func (s *UserService) CreateUser(data *cms_sys_form.CreateAdminUserForm) (*entity.SysUser, error) {
	account, err := s.repo.FindByAccount(data.UserName)
	if err != nil {
		return nil, err
	}
	if account != nil {
		return nil, errors.New("账号已存在，无法重复创建")
	}
	pwdStr, err := utils.GeneratePwd(data.Password)
	if err != nil {
		return nil, errors.New("生成密码失败")
	}
	user := entity.SysUser{
		Username:     data.UserName,
		Password:     pwdStr,
		Email:        "",
		Mobile:       data.Mobile,
		Status:       0,
		CreateUserID: 0,
		ShopID:       0,
	}

	err = s.repo.Insert(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *UserService) GenerateToken(id int64, username string) (string, error) {
	jwtCfg := config.ConfigInstance.Jwt
	j := utils.NewJwt()
	claims := utils.CustomClaims{
		ID:          uint(id),
		Email:       "",
		UserName:    username,
		AuthorityId: 0,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                             // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 3600*jwtCfg.JwtExpireHour, // 签名的过期时间时间
			Issuer:    "",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return token, nil
}

func (s *UserService) AdminLogin(param cms_sys_form.AdminLoginParam) (*dto.PwdLoginDTO, error) {
	account, err := s.FindByAccount(param.UserName)
	if err != nil {
		return nil, err
	}
	// 校验密码
	pwd := utils.CheckPwd(account.Password, account.Password)
	if !pwd {
		return nil, errors.New("密码不正确")
	}
	// 生成Token
	token, err := s.GenerateToken(account.UserID, account.Username)
	if err != nil {
		return nil, err
	}
	res := dto.PwdLoginDTO{AccessToken: token}
	return &res, nil
}
