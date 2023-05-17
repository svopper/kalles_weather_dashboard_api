#! /bin/bash

# Write me a script that calls an api with curl in a foorloop with a formatted date
# The api is https://api.weather.kols.dk/metObs?date=2023-03-15
# The date should be the current date + 1 day

date=$(date +%F)

echo $date

for i in {1..20}
do
    curl https://api.weather.kols.dk/metObs?date=$date
    date=$(gdate -d "$date + 1 day" +%F)
    echo " "
    echo " "

done