#----------------------------
#Cloud run on Anthos set up
#----------------------------
#knative-serving and gke-system namespaces are automatically created

#create cluster with Cloud Run enabled
gcloud container clusters create cr-cluster \
--zone=europe-west4-a \
--addons=HttpLoadBalancing,CloudRun \
--machine-type=n1-standard-2 \
--num-nodes=3 \
--enable-stackdriver-kubernetes

#General