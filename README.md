## build images (for now)
`docker-compose build`

## create kubernetes secret
`kubectl create secret generic rtk-status-page-secret --from-literal=DBPW="my-db-password" --from-literal=AUSCORS="MY-USER:MY-PASSWORD@auscors.ga.gov.au:2101" --from-literal=DBHOST="MY-DB-HOST.ap-southeast-2.rds.amazonaws.com"`

## deploy containers to kubernetes
`kubectl create -f two-container-pod.yaml`
