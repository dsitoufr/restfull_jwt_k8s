Requirements:

    Standard Kubernetes cluster 
    Kubernetes Secret/Credentials
    A build context
    Docker config file: Config.json
    Dockerfile

  Kubernetes Secret/Credentials
  =============================
  gcloud init
gcloud auth application-default login
 ==>
 Credentials saved to file: [/home/vagrant/.config/gcloud/application_default_credentials.json]

 cp -p /home/vagrant/.config/gcloud/application_default_credentials.json   kaniko-secret

kubectl --namespace ns-kaniko create secret generic aws-creds --from-file=kaniko-secret.json
cp -p $HOME/.docker/config.json  .
kubectl --namespace ns-kaniko create configmap config-json --from-file=config.json

kubectl --namespace ns-kaniko create configmap build-context --from-file=Dockerfile


++++ Cours
https://codeghar.com/blog/build-container-images-in-kubernetes-with-kaniko.html
https://hub.docker.com/r/csanchez/kaniko
https://support.cloudbees.com/hc/en-us/articles/360031223512-What-you-need-to-know-when-using-Kaniko-from-Kubernetes-Jenkins-Agents
https://github.com/GoogleContainerTools/kaniko


Save:
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
    stage('New stage') {
      steps {
         sh """
              echo 'new steps'
             """
      }
    }
    
  }
}
