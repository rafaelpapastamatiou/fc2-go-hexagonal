package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rafaelpapastamatiou/fc2-go-hexagonal/adapters/cli"
	mock_application "github.com/rafaelpapastamatiou/fc2-go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	name := "test product"
	price := 25.99
	status := "enabled"
	id := "abc"

	mock := mock_application.NewMockProductInterface(ctrl)

	mock.EXPECT().GetID().Return(id).AnyTimes()
	mock.EXPECT().GetName().Return(name).AnyTimes()
	mock.EXPECT().GetPrice().Return(price).AnyTimes()
	mock.EXPECT().GetStatus().Return(status).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)

	service.EXPECT().Create(name, price).Return(mock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(mock, nil).AnyTimes()
	service.EXPECT().Enable(mock).Return(mock, nil).AnyTimes()
	service.EXPECT().Disable(mock).Return(mock, nil).AnyTimes()

	expectedResult := fmt.Sprintf(
		cli.RESULT_CREATE_MESSAGE,
		id,
		name,
		price,
		status,
	)
	result, err := cli.Run(service, "create", "", name, price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		cli.RESULT_ENABLE_MESSAGE,
		name,
	)
	result, err = cli.Run(service, "enable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		cli.RESULT_DISABLE_MESSAGE,
		name,
	)
	result, err = cli.Run(service, "disable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		cli.RESULT_GET_MESSAGE,
		id,
		name,
		price,
		status,
	)
	result, err = cli.Run(service, "get", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
