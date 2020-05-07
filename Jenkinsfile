pipeline {
  
  environment {
    PROJECT = "REPLACE_WITH_YOUR_PROJECT_ID"
    APP_NAME = "gceme"
    FE_SVC_NAME = "${APP_NAME}-frontend"
    CLUSTER = "jenkins-cd"
    CLUSTER_ZONE = "us-east1-d"
    IMAGE_TAG = "gcr.io/${PROJECT}/${APP_NAME}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
    JENKINS_CRED = "${PROJECT}"
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
                    go version
                    go get -u github.com/rs/cors
                    go get -u github.com/braintree/manners 
                    go get -u github.com/dgrijalva/jwt-go 
                    go get -u github.com/gorilla/mux
                    go get -u golang.org/x/crypto/bcrypt
                    
                   ln -s `pwd` /go/src/restfull_jwt_k8s
                   cd /go/src/restfull_jwt_k8s/hello
                   go test
                """
            }
        }
    }
    
    
    stage('Build Golang') {
      steps {
         sh """
           echo 'building golang'
           cd /go/src/restfull_jwt_k8s/hello
           go build .
          """
        
      }
    }
    
    stage('Build docker image and push to docker hub') {
      steps {
         sh 'echo build docker'
      }
    }
    
    stage('Deploy container image to k8s') {
      steps {
         sh 'echo deploying to k8s'
      }
    }
  }
}
