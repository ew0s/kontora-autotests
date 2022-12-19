package tests

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"testing"

	repo "github.com/ew0s/kontora-autotests/internal/domain/repository"
	"github.com/ew0s/kontora-autotests/internal/mapper"
	"github.com/ew0s/kontora-autotests/internal/service"
	"github.com/ew0s/kontora-autotests/tests/testdata"
	"github.com/stretchr/testify/require"
)

func Test_ServiceInfo_GetAllUsingGET2(t *testing.T) {
	ctx := context.Background()

	serviceControllerClient := service.GetKontoraClient().ServiceInfoControllerApi

	expectedDTOs := testdata.BaseServiceInfos

	actualDTOs, resp, err := serviceControllerClient.GetAllUsingGET2(ctx)
	require.NoError(t, err)

	sort.Slice(actualDTOs, func(i, j int) bool {
		return actualDTOs[i].Id < actualDTOs[j].Id
	})

	require.Equal(t, expectedDTOs, actualDTOs)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func serviceInfoSetupAsBefore(ctx context.Context, repo repo.ServiceInfo) error {
	err := repo.InsertServiceInfos(ctx, mapper.MakeServiceInfoEntities(testdata.BaseServiceInfos))
	if err != nil {
		return fmt.Errorf("inseting service info: %w", err)
	}

	return nil
}
