package service

import (
	"backend/internal/config"
	"backend/internal/dto"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/repositories/users_repositories"
	"backend/pkg/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IUserService interface {
	Get(uid int64) (*models.SysUserModel, error)
	FindByAccount(account string) (*models.SysUserModel, error)
	FindByPhone(phone string) (*models.SysUserModel, error)
	FindAllPager(searchKey string, page, size int) ([]models.SysUserModel, int64, error)
	Insert(data *models.SysUserModel) error
	Update(data *models.SysUserModel, musColumns ...string) error
	Save(data *models.SysUserModel, musColumns ...string) error
	CreateUser(data *forms.CreateAdminUserForm) (*models.SysUserModel, error)
	GenerateToken(id int64, username string) (string, error)
	AdminLogin(param forms.AdminLoginParam) (*dto.PwdLoginDTO, error)
}

type UserService struct {
	repo users_repositories.IUserRepository
}

func NewUserService(r users_repositories.IUserRepository) IUserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Get(uid int64) (*models.SysUserModel, error) {
	return s.repo.Get(uid)
}

func (s *UserService) FindByAccount(account string) (*models.SysUserModel, error) {
	return s.repo.FindByAccount(account)
}

func (s *UserService) FindByPhone(phone string) (*models.SysUserModel, error) {
	return s.repo.FindByPhone(phone)
}

func (s *UserService) FindAllPager(searchKey string, page, size int) ([]models.SysUserModel, int64, error) {
	return s.repo.FindAllPager(searchKey, page, size)
}

func (s *UserService) Insert(data *models.SysUserModel) error {
	return s.repo.Insert(data)
}
func (s *UserService) Update(data *models.SysUserModel, musColumns ...string) error {
	return s.repo.Update(data, musColumns...)
}
func (s *UserService) Save(data *models.SysUserModel, musColumns ...string) error {
	return s.repo.Save(data, musColumns...)
}

func (s *UserService) CreateUser(data *forms.CreateAdminUserForm) (*models.SysUserModel, error) {
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
	user := models.SysUserModel{
		Username:    data.UserName,
		Password:    pwdStr,
		Email:       "",
		Mobile:      data.Mobile,
		Status:      0,
		CreatUserId: 0,
		ShopId:      0,
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

func (s *UserService) AdminLogin(param forms.AdminLoginParam) (*dto.PwdLoginDTO, error) {
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
	token, err := s.GenerateToken(account.ID, account.Username)
	if err != nil {
		return nil, err
	}
	res := dto.PwdLoginDTO{AccessToken: token}
	return &res, nil
}
