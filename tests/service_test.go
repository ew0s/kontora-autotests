package tests

import (
	"context"
	"net/http"
	"sort"
	"testing"

	"github.com/ew0s/kontora-autotests/internal/service"
	"github.com/ew0s/kontora-autotests/swagger"
	"github.com/stretchr/testify/require"
)

func Test_Service_GetAllUsingGET1(t *testing.T) {
	ctx := context.Background()

	serviceControllerClient := service.GetKontoraClient().ServiceControllerApi

	expectedDTOs := []swagger.ServiceDto{
		{
			Id:   1,
			Name: "Стрижка налысо",
		},
		{
			Id:   2,
			Name: "Барбер стрижка",
		},
		{
			Id:   3,
			Name: "Маникюр",
		},
		{
			Id:   4,
			Name: "Стрижка + маникюр",
		},
	}

	actualDTOs, resp, err := serviceControllerClient.GetAllUsingGET1(ctx)
	require.NoError(t, err)

	sort.Slice(actualDTOs, func(i, j int) bool {
		return actualDTOs[i].Id < actualDTOs[j].Id
	})

	require.Equal(t, expectedDTOs, actualDTOs)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
