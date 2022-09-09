package stations

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/util"
)

func generateTemperatureUri() string {
	uri := fmt.Sprintf(
		"https://dmigw.govcloud.dk/v2/metObs/collections/station/items?datetime=%s/&status=Active&bbox-crs=https://www.opengis.net/def/crs/OGC/1.3/CRS84&api-key=%s",
		util.FormatDate(time.Now()),
		util.GetEnvVariable("DMI_MET_OBS_API_KEY"),
	)
	return uri
}

func (h handler) GetMetObsStations(c *gin.Context) {
	uri := generateTemperatureUri()
	request := util.BuildRequest(uri)
	response := util.DoRequest(request)
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var r DMIStationResponse
	err = json.Unmarshal(body, &r)

	if err != nil {
		panic(err)
	}

	c.JSON(200, r)
}
