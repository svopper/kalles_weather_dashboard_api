# Kalle's Weather Dashboard API

> API that consumes the DMI API and presents weather data through different endpoints.

## Hosting

App is hosted on [api.weather.kols.dk](https://api.weather.kols.dk/).

## Setup

This app uses env variables to get API keys. The following keys are needed:

| **Variable**          | **Description**                                              | **Status** |
| --------------------- | ------------------------------------------------------------ | ---------- |
| DMI_MET_OBS_API_KEY   | Key to access to weather observation API                     | Required   |
| DMI_OCEAN_OBS_API_KEY | Key to access to ocean observartion API                      | Required   |
| PORT                  | The port that the app is exposed through on the host machine | Optional   |

During development, API keys should be stored as env variables on the OS. When deploying to production, proper API keys are injected in [pkg/common/envs/config.yaml](https://github.com/svopper/kalles_weather_dashboard_v2/blob/main/pkg/common/envs/config.yaml) from a GitHub Action. API keys are found on the [DMI Open Data platform](https://dmiapi.govcloud.dk/).

## Deployment

The application is deployed to Google Cloud Platform on every push to `main`. It is not possible to deploy through the `gcloud` CLI, due to API keu injection in deployment process.
