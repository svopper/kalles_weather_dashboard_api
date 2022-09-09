package util

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/models"
)

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func RoundToTwoDecimal(num float64) float64 {
	rounded := math.Round(num*10) / 10
	return rounded
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02T15:04:05Z")
}

func BuildRequest(uri string) *http.Request {
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("accept", "application/geo+json")
	return req
}

func DoRequest(request *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(request)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)
		panic(fmt.Sprintf("Request failed with status code %d. Error: %s", resp.StatusCode, buf.String()))
	}
	if err != nil {
		panic(err)
	}

	return resp
}

func UnmarshalDMIObservation(data []byte) (models.DMIObservation, error) {
	var r models.DMIObservation
	err := json.Unmarshal(data, &r)
	return r, err
}

func GetEnvVariable(key string) string {
	value := viper.Get(key)
	if value == "xxx" {
		return os.Getenv(key)
	}
	return value.(string)
}

func GetQueryParameter(c *gin.Context, key, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}
