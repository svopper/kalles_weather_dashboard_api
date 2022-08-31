# Kalle's Weather Dashboard 2

> App that consumes the DMI API and presents weather data through different endpoints.

## Setup

This app uses env variables to get API keys. The following keys are needed:

| **Variable**          | **Description**                                              | **Status** |
| --------------------- | ------------------------------------------------------------ | ---------- |
| DMI_MET_OBS_API_KEY   | Key to access to weather observation API                     | Required   |
| DMI_OCEAN_OBS_API_KEY | Key to access to ocean observartion API                      | Required   |
| PORT                  | The port that the app is exposed through on the host machine | Optional   |
