package service

import (
	"TeamTickBackend/dal/models"
	"TeamTickBackend/pkg"
	apperrors "TeamTickBackend/pkg/errors"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// --- Mock 实现 ---

// Mock UserDAO
type mockUserDAO struct {
	mock.Mock
}

type mockEmailRedisDAO struct {
	mock.Mock
}

type mockEmailService struct {
	mock.Mock
}

func (m *mockEmailService) GenerateVerificationCode(length int) (string, error) {
	args := m.Called(length)
	return args.String(0), args.Error(1)
}

func (m *mockEmailService) SendVerificationEmail(ctx context.Context, email, code string) error {
	args := m.Called(ctx, email, code)
	return args.Error(0)
}

func (m *mockEmailRedisDAO) GetVerificationCodeByEmail(ctx context.Context, email string, tx ...*redis.Client) (string, error) {
	args := m.Called(ctx, email, tx)
	return args.String(0), args.Error(1)
}

func (m *mockEmailRedisDAO) SetVerificationCodeByEmail(ctx context.Context, email string, code string) error {
	args := m.Called(ctx, email, code)
	return args.Error(0)
}

func (m *mockEmailRedisDAO) DeleteCacheByEmail(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}

func (m *mockUserDAO) GetByUsername(ctx context.Context, username string, tx ...*gorm.DB) (*models.User, error) {
	args := m.Called(ctx, username, tx)
	// 处理预期返回nil用户的情况
	userArg := args.Get(0)
	if userArg == nil {
		return nil, args.Error(1)
	}
	return userArg.(*models.User), args.Error(1)
}

func (m *mockUserDAO) Create(ctx context.Context, user *models.User, tx ...*gorm.DB) error {
	args := m.Called(ctx, user, tx)
	// 模拟创建时分配ID和时间
	if args.Error(0) == nil {
		user.UserID = 1
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
	}
	return args.Error(0)
}

// GetByID方法实现
func (m *mockUserDAO) GetByID(ctx context.Context, id int, tx ...*gorm.DB) (*models.User, error) {
	args := m.Called(ctx, id, tx)
	userArg := args.Get(0)
	if userArg == nil {
		return nil, args.Error(1)
	}
	return userArg.(*models.User), args.Error(1)
}

func (m *mockUserDAO) UpdatePassword(ctx context.Context, userID int, password string, tx ...*gorm.DB) error {
	args := m.Called(ctx, userID, password, tx)
	return args.Error(0)
}

func (m *mockUserDAO) GetByEmail(ctx context.Context, email string, tx ...*gorm.DB) (*models.User, error) {
	args := m.Called(ctx, email, tx)
	return args.Get(0).(*models.User), args.Error(1)
}

// Mock JwtHandler
type mockJwtHandler struct {
	mock.Mock
}

func (m *mockJwtHandler) GenerateJWTToken(username string, userID int) (string, error) {
	args := m.Called(username, userID)
	return args.String(0), args.Error(1)
}

// ParseJWTToken方法实现
func (m *mockJwtHandler) ParseJWTToken(tokenString string) (pkg.JwtPayload, error) {
	args := m.Called(tokenString)
	payloadArg := args.Get(0)
	if payloadArg == nil {
		return pkg.JwtPayload{}, args.Error(1)
	}
	return payloadArg.(pkg.JwtPayload), args.Error(1)
}

// --- 测试准备 ---

func setupAuthServiceTest() (*AuthService, *mockUserDAO, *mockTransactionManager, *mockJwtHandler, *mockEmailRedisDAO) {
	mockUserDao := new(mockUserDAO)
	mockTxManager := new(mockTransactionManager)
	mockJwt := new(mockJwtHandler)
	mockEmailRedisDAO := new(mockEmailRedisDAO)
	mockEmailService := new(mockEmailService)
	authService := NewAuthService(mockUserDao, mockTxManager, mockJwt, mockEmailRedisDAO, mockEmailService)
	return authService, mockUserDao, mockTxManager, mockJwt, mockEmailRedisDAO
}

// --- AuthRegister 测试 ---

func TestAuthRegister_Success(t *testing.T) {
	authService, mockUserDao, mockTxManager, _, _ := setupAuthServiceTest()
	ctx := context.Background()
	username := "newuser"
	password := "password123"

	// Mock期望
	mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(*gorm.DB) error")).Return(nil)
	mockUserDao.On("GetByUsername", ctx, username, mock.AnythingOfType("[]*gorm.DB")).Return(nil, gorm.ErrRecordNotFound)
	mockUserDao.On("Create", ctx, mock.AnythingOfType("*models.User"), mock.AnythingOfType("[]*gorm.DB")).Return(nil).Run(func(args mock.Arguments) {
		// 验证传给Create的用户对象
		userArg := args.Get(1).(*models.User)
		assert.Equal(t, username, userArg.Username)
		assert.NotEmpty(t, userArg.Password, "密码应被哈希，不为空")
	})

	// 调用函数
	createdUser, err := authService.AuthRegister(ctx, username, password, "", "")

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, username, createdUser.Username)
	assert.NotEqual(t, password, createdUser.Password) // 确保密码非明文存储
	assert.NotZero(t, createdUser.UserID)              // 检查ID是否被分配

	// 验证mock调用
	mockTxManager.AssertExpectations(t)
	mockUserDao.AssertExpectations(t)
}

func TestAuthRegister_UserAlreadyExists(t *testing.T) {
	authService, mockUserDao, mockTxManager, _, _ := setupAuthServiceTest()
	ctx := context.Background()
	username := "existinguser"
	password := "password123"
	existingUser := &models.User{UserID: 1, Username: username}

	// Mock期望
	mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(*gorm.DB) error")).Return(nil)
	mockUserDao.On("GetByUsername", ctx, username, mock.AnythingOfType("[]*gorm.DB")).Return(existingUser, nil) // 用户已存在

	// 调用函数
	createdUser, err := authService.AuthRegister(ctx, username, password, "", "")

	// 断言
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrUserAlreadyExists))
	assert.Nil(t, createdUser)

	// 验证mock调用
	mockTxManager.AssertExpectations(t)
	mockUserDao.AssertExpectations(t)
	mockUserDao.AssertNotCalled(t, "Create", mock.Anything, mock.Anything, mock.Anything) // Create不应被调用
}

// --- AuthLogin 测试 ---

// 密码哈希辅助函数
func getTestPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func TestAuthLogin_Success(t *testing.T) {
	authService, mockUserDao, mockTxManager, mockJwt, _ := setupAuthServiceTest()
	ctx := context.Background()
	username := "testuser"
	password := "password123"
	userID := 5
	expectedToken := "valid.jwt.token"
	hash, err := getTestPasswordHash(password)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}
	foundUser := &models.User{
		UserID:   userID,
		Username: username,
		Password: hash,
	}

	// Mock期望
	mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(*gorm.DB) error")).Return(nil)
	mockUserDao.On("GetByUsername", ctx, username, mock.AnythingOfType("[]*gorm.DB")).Return(foundUser, nil)
	mockJwt.On("GenerateJWTToken", username, userID).Return(expectedToken, nil)

	// 调用函数
	loggedInUser, token, err := authService.AuthLogin(ctx, username, password)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, loggedInUser)
	assert.Equal(t, foundUser, loggedInUser)
	assert.Equal(t, expectedToken, token)

	// 验证mock
	mockTxManager.AssertExpectations(t)
	mockUserDao.AssertExpectations(t)
	mockJwt.AssertExpectations(t)
}

func TestAuthLogin_UserNotFound(t *testing.T) {
	authService, mockUserDao, mockTxManager, mockJwt, _ := setupAuthServiceTest()
	ctx := context.Background()
	username := "nonexistentuser"
	password := "password123"

	// Mock期望
	mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(*gorm.DB) error")).Return(nil)
	mockUserDao.On("GetByUsername", ctx, username, mock.AnythingOfType("[]*gorm.DB")).Return(nil, gorm.ErrRecordNotFound)

	// 调用函数
	loggedInUser, token, err := authService.AuthLogin(ctx, username, password)

	// 断言
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrUserNotFound))
	assert.Nil(t, loggedInUser)
	assert.Empty(t, token)

	// 验证mock
	mockTxManager.AssertExpectations(t)
	mockUserDao.AssertExpectations(t)
	mockJwt.AssertNotCalled(t, "GenerateJWTToken", mock.Anything, mock.Anything)
}

func TestAuthLogin_InvalidPassword(t *testing.T) {
	authService, mockUserDao, mockTxManager, mockJwt, _ := setupAuthServiceTest()
	ctx := context.Background()
	username := "testuser"
	wrongPassword := "wrongpassword"
	correctPassword := "password123"
	userID := 5
	hash, err := getTestPasswordHash(correctPassword)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}
	foundUser := &models.User{
		UserID:   userID,
		Username: username,
		Password: hash,
	}

	// Mock期望
	mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(*gorm.DB) error")).Return(nil)
	mockUserDao.On("GetByUsername", ctx, username, mock.AnythingOfType("[]*gorm.DB")).Return(foundUser, nil)

	// 调用函数
	loggedInUser, token, err := authService.AuthLogin(ctx, username, wrongPassword)

	// 断言
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrInvalidPassword))
	assert.Nil(t, loggedInUser)
	assert.Empty(t, token)

	// 验证mock
	mockTxManager.AssertExpectations(t)
	mockUserDao.AssertExpectations(t)
	mockJwt.AssertNotCalled(t, "GenerateJWTToken", mock.Anything, mock.Anything)
}

// --- Add more tests as needed ---
// e.g., tests for database errors during Create in Register,
// database errors (non-NotFound) during GetByUsername in Login,
// errors returned directly by WithTransaction itself.
