package metObs

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

func getAverageMaxTemp(observations []models.TemperatureObservation) float64 {
	var sum float64
	iterations := 0
	for _, observation := range observations {
		if observation.Max == math.Inf(1) {
			continue
		}
		sum += observation.Max
		iterations++
	}
	return util.RoundToTwoDecimal(sum / float64(iterations))
}

func getAverageMinTemp(observations []models.TemperatureObservation) float64 {
	var sum float64
	iterations := 0
	for _, observation := range observations {
		if observation.Min == math.Inf(-1) {
			continue
		}
		sum += observation.Min
		iterations++

	}
	return util.RoundToTwoDecimal(sum / float64(iterations))
}

func getWatherObservations(from, to time.Time) models.DMIObservation {
	uri := generateTemperatureUri(from, to)
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

func generateTemperatureUri(fromDate, toDate time.Time) string {
	uri := fmt.Sprintf(
		"https://dmigw.govcloud.dk/v2/metObs/collections/observation/items?datetime=%s/%s&stationId=06180&parameterId=temp_dry&bbox-crs=https://www.opengis.net/def/crs/OGC/1.3/CRS84&api-key=%s",
		util.FormatDate(fromDate),
		util.FormatDate(toDate),
		util.GetEnvVariable("DMI_MET_OBS_API_KEY"),
	)
	return uri
}

func getMinAndMax(features []models.Feature) (float64, float64) {
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, feature := range features {
		if feature.Properties.Value < min {
			min = feature.Properties.Value
		}
		if feature.Properties.Value > max {
			max = feature.Properties.Value
		}
	}
	return min, max
}

// GetMetObs godoc
// @Description  Get metrological observations for a given date. If no date is given, the current date is used.
// @Summary  	 Get metrological observations
// @Tags         metObs
// @Produce      json
// @Param        date	query	string  false  "Date"
// @Success      200  {object}  models.MetObservationResponse
// @Router       /metObs/ [get]
func GetMetObs(c *gin.Context) {
	requestedDate := c.Query("date")
	var parsedDate time.Time

	if tmpDate, err := time.Parse("2006-01-02", requestedDate); err == nil {
		parsedDate = tmpDate
	} else {
		parsedDate = time.Now()
	}

	viewModel := models.MetObservationResponse{
		Date:                    parsedDate.Format("January 02"),
		TemperatureObservations: []models.TemperatureObservation{},
	}
	for i := 1; i <= 10; i++ {
		year := parsedDate.Year() - i
		month := parsedDate.Month()
		day := parsedDate.Day()
		if !util.IsLeapYear(year) && month == time.February && day == 29 {
			viewModel.TemperatureObservations = append(viewModel.TemperatureObservations, models.TemperatureObservation{Year: year, Min: math.Inf(-1), Max: math.Inf(1)})
			continue
		}
		fromDate := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
		toDate := time.Date(year, month, day, 23, 59, 0, 0, time.Now().Location())
		w := getWatherObservations(fromDate, toDate)
		min, max := getMinAndMax(w.Features)
		obs := models.TemperatureObservation{Year: year, Min: min, Max: max}
		viewModel.TemperatureObservations = append(viewModel.TemperatureObservations, obs)
	}
	viewModel.MaxAverage = getAverageMaxTemp(viewModel.TemperatureObservations)
	viewModel.MinAverage = getAverageMinTemp(viewModel.TemperatureObservations)

	c.JSON(http.StatusOK, viewModel)
}
