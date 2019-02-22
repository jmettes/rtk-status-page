# build images (for now)
`docker-compose build`

# run the logger image
```
docker run -e "STATION=ALIC7" -e "DBHOST=my-url-to-db" -e "DBPT=5432" -e "DBU=postgres" -e "DBPW=my-db-password" -e "DBNAME=rtkstatuspage" --mount src=$(pwd)/logs,target=/logs,type=bind -t rtk-status-page_logger
```

# run rtklib image
```
docker run -e "STATION=ALIC7" -e "DBHOST=my-url-to-db" -e "DBPT=5432" -e "DBU=postgres" -e "DBPW=my-db-password" -e "DBNAME=rtkstatuspage" -it --mount src=$(pwd)/logs,target=/logs,type=bind rtk-status-page_rtkrcv sh -c "sed -i 's/STATION/ALIC7/' single.conf && sed -i 's/AUSCORS/MYUSERNAME:MYPASSWORD@auscors.ga.gov.au:2101/' single.conf && ./rtkrcv -o single.conf -s -d /dev/tty"
```
