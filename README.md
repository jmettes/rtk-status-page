## build images
`docker-compose build`

## deploy images
`docker tag rtk-status-page_rtkrcv asia.gcr.io/<PROJECT-NAME>/rtk-status-page_rtkrcv:latest`
`docker tag rtk-status-page_logger asia.gcr.io/<PROJECT-NAME>/rtk-status-page_logger:latest`
`docker push asia.gcr.io/<MY-PROJECT-NAME>/rtk-status-page_rtkrcv`
`docker push asia.gcr.io/<MY-PROJECT-NAME>/rtk-status-page_logger`

## create kubernetes secret
`kubectl create secret generic rtk-status-page-secret --from-literal=DBPW="my-db-password" --from-literal=AUSCORS="MY-USER:MY-PASSWORD@auscors.ga.gov.au:2101" --from-literal=DBHOST="MY-DB-HOST.ap-southeast-2.rds.amazonaws.com"`

## deploy containers to kubernetes
`kubectl create -f two-container-pod.yaml`
