package oceanObs

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/models"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/util"
)

func generateOceanUri(stationId int) string {
	uri := fmt.Sprintf(
		"https://dmigw.govcloud.dk/v2/oceanObs/collections/observation/items?period=latest-day&stationId=%d&parameterId=tw&sortorder=observed,DESC&bbox-crs=https://www.opengis.net/def/crs/OGC/1.3/CRS84&api-key=%s",
		stationId,
		util.GetEnvVariable("DMI_OCEAN_OBS_API_KEY"),
	)
	return uri
}

func getOceanObservations(stationId int) models.DMIObservation {
	uri := generateOceanUri(stationId)
	request := util.BuildRequest(uri)
	response := util.DoRequest(request)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	weatherObs, err := util.UnmarshalDMIObservation(body)

	if err != nil {
		panic(err)
	}
	return weatherObs
}

func getMax(features []models.Feature) float64 {
	max := math.Inf(-1)
	for _, feature := range features {
		if feature.Properties.Value > max {
			max = feature.Properties.Value
		}
	}
	return max
}

// GetOceanObs	 godoc
// @Description  Get the current ocean temperature and highest temperature in the past 24h from a fixed set of locations
// @Summary  	 Get current ocean temperature
// @Tags         oceanObs
// @Produce      json
// @Success      200  {object}  models.OceanObservationResponse
// @Router       /oceanObs/ [get]
func GetOceanObs(c *gin.Context) {
	oceanObsModel := models.OceanObservationResponse{
		Date: time.Now().Format("January 02"),
	}

	for stationId, stationName := range util.OCEAN_STATION_MAP {
		obs := getOceanObservations(stationId)
		observation := models.OceanObservation{
			StationId:         stationId,
			StationName:       stationName,
			MaxTemp24H:        getMax(obs.Features),
			LatestTemperature: obs.Features[0].Properties.Value,
		}
		oceanObsModel.Observations = append(oceanObsModel.Observations, observation)
	}

	c.JSON(http.StatusOK, oceanObsModel)
}
