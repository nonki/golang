package foo_test

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/nonki/golang/mocks"
	"testing"
)

func Test_MyIssuer_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	container := mocks.NewMockContainer(ctrl)
	container.EXPECT().Get().Return("hi")
	issuer := mocks.NewMockIssuer(ctrl)
	issuer.EXPECT().Get().Return(*container)

	issuer.Get().Get()
}
