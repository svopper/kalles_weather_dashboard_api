package metObs

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/models"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/util"
)

func getAverageMaxTemp(observations []float64) float64 {
	var sum float64
	iterations := 0
	outliersRemoved := removeMinAndMaxValue(observations)
	fmt.Println(observations)
	fmt.Println(outliersRemoved)
	for _, observation := range outliersRemoved {
		if observation == math.Inf(1) {
			continue
		}
		sum += observation
		iterations++
	}
	return util.RoundToTwoDecimal(sum / float64(iterations))
}

func getAverageMinTemp(observations []float64) float64 {
	var sum float64
	iterations := 0
	outliersRemoved := removeMinAndMaxValue(observations)
	for _, observation := range outliersRemoved {
		if observation == math.Inf(-1) {
			continue
		}
		sum += observation
		iterations++

	}
	return util.RoundToTwoDecimal(sum / float64(iterations))
}

func removeMinAndMaxValue(observations []float64) []float64 {
	sorted := sort.Float64Slice(observations)
	sort.Sort(sorted)
	return sorted[1 : len(sorted)-1]
}

func getMaxTempValues(observations []models.TemperatureObservation) []float64 {
	result := []float64{}
	for _, observation := range observations {
		result = append(result, observation.Max)
	}
	return result
}

func getMinTempValues(observations []models.TemperatureObservation) []float64 {
	result := []float64{}
	for _, observation := range observations {
		result = append(result, observation.Min)
	}
	return result
}

func getWatherObservations(from, to time.Time) models.DMIObservation {
	uri := generateTemperatureUri(from, to)
	fmt.Println(uri)
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

func getTemperatureObservation(h handler, keyMin, keyMax string, fromDate, toDate time.Time) models.TemperatureObservation {
	var to models.TemperatureObservation
	value_min, err_min := h.DB.Get(keyMin).Result()
	value_max, err_max := h.DB.Get(keyMax).Result()
	fmt.Println(value_min)
	fmt.Println(value_max)
	if err_min != redis.Nil && err_max != redis.Nil {
		v_min, _ := strconv.ParseFloat(value_min, 64)
		v_max, _ := strconv.ParseFloat(value_max, 64)
		to.Min = math.Round(v_min)
		to.Max = math.Round(v_max)
	} else {
		w := getWatherObservations(fromDate, toDate)
		min, max := getMinAndMax(w.Features)
		h.DB.Set(keyMin, min, 0)
		h.DB.Set(keyMax, max, 0)
		to.Min = math.Round(min)
		to.Max = math.Round(max)
	}
	to.Year = fromDate.Year()
	return to

}

func (h handler) GetAverage(c *gin.Context) {

}

// GetMetObs godoc
// @Description  Get metrological observations for a given date. If no date is given, the current date is used.
// @Summary  	 Get metrological observations
// @Tags         metObs
// @Produce      json
// @Param        date	query	string  false  "Date (default: today)"
// @Param        numberOfYears	query	int  false  "Number of years back to get observations for (default: 10)"
// @Success      200  {object}  models.MetObservationResponse
// @Router       /metObs/ [get]
func (h handler) GetMetObs(c *gin.Context) {
	requestedDate := util.GetQueryParameter(c, "date", time.Now().String())

	var parsedDate time.Time
	if tmpDate, err := time.Parse("2006-01-02", requestedDate); err == nil {
		parsedDate = tmpDate
	} else {
		parsedDate = time.Now()
	}

	numberOfYears, err := strconv.Atoi(util.GetQueryParameter(c, "numberOfYears", "10"))
	if err != nil {
		numberOfYears = 10
	}

	viewModel := models.MetObservationResponse{
		Date:                    parsedDate.Format("January 02"),
		TemperatureObservations: []models.TemperatureObservation{},
	}
	for i := 1; i <= numberOfYears; i++ {
		year := parsedDate.Year() - i
		month := parsedDate.Month()
		day := parsedDate.Day()
		if !util.IsLeapYear(year) && month == time.February && day == 29 {
			viewModel.TemperatureObservations = append(viewModel.TemperatureObservations, models.TemperatureObservation{Year: year, Min: math.Inf(-1), Max: math.Inf(1)})
			continue
		}
		fromDate := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
		toDate := time.Date(year, month, day, 23, 59, 0, 0, time.Now().Location())

		key_min := fmt.Sprintf("%d-%d-%d-min", year, month, day)
		key_max := fmt.Sprintf("%d-%d-%d-max", year, month, day)

		tempObs := getTemperatureObservation(h, key_min, key_max, fromDate, toDate)

		viewModel.TemperatureObservations = append(viewModel.TemperatureObservations, tempObs)

	}
	viewModel.MaxAverage = getAverageMaxTemp(getMaxTempValues(viewModel.TemperatureObservations))
	viewModel.MinAverage = getAverageMinTemp(getMinTempValues(viewModel.TemperatureObservations))

	c.JSON(http.StatusOK, viewModel)
}
