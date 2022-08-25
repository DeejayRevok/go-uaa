package getAuthenticatedUser

import (
	"errors"
	"go-uaa/mocks"
	"go-uaa/src/domain/auth/accessToken"
	"go-uaa/src/domain/user"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type testCase struct {
	UserRepo *mocks.UserRepository
	UseCase  *GetAuthenticatedUserUseCase
}

func setUp(t *testing.T) testCase {
	logger, _ := zap.NewDevelopment()
	userRepositoryMock := mocks.NewUserRepository(t)
	return testCase{
		UserRepo: userRepositoryMock,
		UseCase:  NewGetAuthenticatedUserUseCase(userRepositoryMock, logger),
	}
}

func TestExecuteWrongRequest(t *testing.T) {
	testCase := setUp(t)
	request := "wrongRequest"

	response := testCase.UseCase.Execute(request)

	if response.Err == nil {
		t.Fatal("Expected use case to return error")
	}
	testCase.UserRepo.AssertNotCalled(t, "FindByUsername")
}

func TestExecuteFindError(t *testing.T) {
	testCase := setUp(t)
	findError := errors.New("Test find user error")
	testCase.UserRepo.On("FindByUsername", mock.Anything).Return(nil, findError)
	testUsername := "TestUser"
	testToken := accessToken.AccessToken{
		Sub: testUsername,
	}
	request := GetAuthenticatedUserRequest{
		Token: testToken,
	}

	response := testCase.UseCase.Execute(&request)

	if response.Err == nil {
		t.Fatal("Expected use case to return error")
	}
	if response.Err != findError {
		t.Fatal("Expected use case to return user repository find error")
	}
	testCase.UserRepo.AssertCalled(t, "FindByUsername", testUsername)
}

func TestExecuteSuccess(t *testing.T) {
	testCase := setUp(t)
	testUsername := "TestUser"
	testUser := user.User{
		Username: testUsername,
	}
	testCase.UserRepo.On("FindByUsername", mock.Anything).Return(&testUser, nil)
	testToken := accessToken.AccessToken{
		Sub: testUsername,
	}
	request := GetAuthenticatedUserRequest{
		Token: testToken,
	}

	response := testCase.UseCase.Execute(&request)

	if response.Err != nil {
		t.Fatal("Expected use case not to return error")
	}
	responseUser := *response.Content.(*user.User)
	if !reflect.DeepEqual(responseUser, testUser) {
		t.Fatal("Expected use case ro return same user as the repository")
	}
	testCase.UserRepo.AssertCalled(t, "FindByUsername", testUsername)
}
