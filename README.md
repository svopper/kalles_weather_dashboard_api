# Kalle's Weather Dashboard API

> API that consumes the DMI API and presents weather data through different endpoints.

🏠 [api.weather.kols.dk](https://api.weather.kols.dk/)

## Hosting

App is hosted on [api.weather.kols.dk](https://api.weather.kols.dk/).

SSL are managed through [GCP](https://console.cloud.google.com/net-services/loadbalancing/advanced/sslCertificates/list?project=kalles-weather-dashboard-api).

## Setup

This app uses env variables to get API keys. The following keys are needed:

| **Variable**          | **Description**                                              | **Status** |
| --------------------- | ------------------------------------------------------------ | ---------- |
| DMI_MET_OBS_API_KEY   | Key to access to weather observation API                     | Required   |
| DMI_OCEAN_OBS_API_KEY | Key to access to ocean observartion API                      | Required   |
| PORT                  | The port that the app is exposed through on the host machine | Optional   |

During development, API keys should be stored as env variables on the OS. When deploying to production, proper API keys are injected in [pkg/common/envs/config.yaml](https://github.com/svopper/kalles_weather_dashboard_v2/blob/main/pkg/common/envs/config.yaml) from a GitHub Action. API keys are found on the [DMI Open Data platform](https://dmiapi.govcloud.dk/).

## Deployment

The project is deployed to [jakkes-weather-dashboard-api](https://console.cloud.google.com/welcome?project=kalles-weather-dashboard-api) on Google Cloud Platform.

The application is deployed to GCP on every push to `main`. It is not possible to deploy through the `gcloud` CLI, due to API key injection in deployment process.

## Documentation

Using Swagger ([gin-swagger](https://github.com/swaggo/gin-swagger)) to give endpoint overview. Navigate to `/swagger/index.html` to browse.

Swagger doc strings has to be created on each endpoint. See https://github.com/swaggo/swag/blob/master/README.md#api-operation for more info.

## Dependecies

| **Package**                 | **Description**                | **Link**                            |
| --------------------------- | ------------------------------ | ----------------------------------- |
| github.com/gin-gonic/gin    | HTTP web framework             | https://github.com/gin-gonic/gin    |
| github.com/gin-contrib/cors | Manage CORS                    | https://github.com/gin-contrib/cors |
| github.com/spf13/viper      | Managing of configuration      | https://github.com/spf13/viper      |
| github.com/swaggo/swag      | API documentation with Swagger | https://github.com/swaggo/swag      |
