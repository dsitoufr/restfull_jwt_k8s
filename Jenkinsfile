pipeline {
  
  environment {
    PROJECT = "my-fisrt-proj"
    APP_NAME = "myhello"
    FE_SVC_NAME = "${APP_NAME}-frontend"
    CLUSTER = "microk8s-cluster"
    CLUSTER_ZONE = "us-east1-d"
    IMAGE_TAG = "eu.gcr.io/my-fisrt-proj/mykaniko-k8s"
    JENKINS_CRED = "${PROJECT}"
    CLOUDSDK_CORE_PROJECT="my-fisrt-proj"
    ACCOUNT = "dsitoufr@gmail.com"
  }
  
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: some-label-value
spec:
  containers:
  - name: golang
    image: golang:1.10
    command:
    - cat
    tty: true
  - name: gcloud
    image: gcr.io/cloud-builders/gcloud
    command:
    - cat
    tty: true
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
    command:
    - cat
    tty: true
"""
    }
  }
  stages {
    stage('Install dependencies  and unit tests'){
        steps {
            container('golang') {
                sh """
                    echo 'building go sources'
                    go version
                    go get -u github.com/rs/cors
                    go get -u github.com/braintree/manners 
                    go get -u github.com/dgrijalva/jwt-go 
                    go get -u github.com/gorilla/mux
                    go get -u golang.org/x/crypto/bcrypt
                    
                    ln -s `pwd` /go/src/restfull_jwt_k8s
                    cd /go/src/restfull_jwt_k8s/hello
                    go build .
                """
            }
        }
    }
    stage('Build and push image to google container registry') {
      steps {
        container('gcloud') {
           sh """
               gcloud config set project ${PROJECT}
               gcloud config set account ${ACCOUNT}
               PYTHONUNBUFFERED=1 gcloud builds submit -t ${IMAGE_TAG} .
              """
        }
      }
    }
    
  }
}
