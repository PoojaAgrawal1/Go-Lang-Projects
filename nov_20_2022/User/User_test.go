package User_test

import (
	"sample/test/User"
	"sample/test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t) //attaching testing function to mock controller
	defer mockCtrl.Finish()             //dispose the object

	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &User.User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "sample test").Return(nil).Times(1)
	testUser.Use()
}
