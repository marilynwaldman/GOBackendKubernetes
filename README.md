
Push to Google Cloud without CORS first.  Check with Curl.

### Tutorial from :
    https://medium.com/google-cloud/running-a-simple-kubernetes-frontend-backend-service-on-google-cloud-platform-85eb0346f600

This works!!!!


gcloud config set project testkube-187517
gcloud config set compute/zone us-west1-b 

gcloud container clusters create kubetest --num-nodes=3

git clone https://github.com/marilynwaldman/kubefrontback.git

cd kub*


cd ../be
   docker build -t be .

cd ..

kubectl create -f  be-rc.yaml
kubectl create -f  be-srv.yaml

kubectl get service


docker tag be gcr.io/testkube-187517/be


gcloud docker -- push gcr.io/testkube-187517/be

curl http://....

kubectl logs -f POD-NAME

kubectl delete -f  be-srv.yaml
kubectl delete -f  be-rc.yaml







docker build -t be .
docker run --rm -p 3000:3000 be

Utilities

###!/bin/bash
##### Delete all containers
docker rm $(docker ps -a -q)
### Delete all images
docker rmi $(docker images -q)