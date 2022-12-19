package tests

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"testing"

	repo "github.com/ew0s/kontora-autotests/internal/domain/repository"
	"github.com/ew0s/kontora-autotests/internal/mapper"
	"github.com/ew0s/kontora-autotests/internal/repository"
	"github.com/ew0s/kontora-autotests/internal/service"
	"github.com/ew0s/kontora-autotests/swagger"
	"github.com/ew0s/kontora-autotests/tests/testdata"
	"github.com/stretchr/testify/require"
)

func Test_Hairdresser_GetAllUsingGET_Success(t *testing.T) {
	ctx := context.Background()

	hairdresserClient := service.GetKontoraClient().HairdresserControllerApi

	expectedDTOs := testdata.BaseHairdressers
	actualDTOs, resp, err := hairdresserClient.GetAllUsingGET(ctx)
	require.NoError(t, err)

	sort.Slice(actualDTOs, func(i, j int) bool {
		return actualDTOs[i].Id < actualDTOs[j].Id
	})

	require.Equal(t, expectedDTOs, actualDTOs)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_Hairdresser_GetAllUsingGET_Empty(t *testing.T) {
	ctx := context.Background()

	hairdresserClient := service.GetKontoraClient().HairdresserControllerApi
	hairdresserRepo := repository.GetHairdresserRepository()
	serviceInfoRepo := repository.GetServiceInfoRepository()

	err := hairdresserRepo.Truncate(ctx)
	require.NoError(t, err)

	expectedDTOs := make([]swagger.HairdresserDto, 0)
	actualDTOs, resp, err := hairdresserClient.GetAllUsingGET(ctx)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, expectedDTOs, actualDTOs)

	if err = hairdressersSetupAsBefore(ctx, hairdresserRepo); err != nil {
		t.Fatalf("setting up hairdressers as before: %s", err)
	}
	if err = serviceInfoSetupAsBefore(ctx, serviceInfoRepo); err != nil {
		t.Fatalf("setting up service_info as before: %s", err)
	}
}

func hairdressersSetupAsBefore(ctx context.Context, repo repo.Hairdresser) error {
	err := repo.InsertHairdressers(ctx, mapper.MakeHairdresserEntities(testdata.BaseHairdressers))
	if err != nil {
		return fmt.Errorf("inseting hairdressers: %w", err)
	}

	return nil
}
