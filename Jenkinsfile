pipeline {
   agent none
    environment {
        HELLO='restfull_jwt_k8s/hello'
        AUTH='auth'
    }
   
    stages {
        stage('build') {
            //agent { docker { image 'golang' }}
           docker.image("golang:1.8.0-alpine").inside("-v ${pwd()}:${goPath}") {
              sh 'go version'
           }
        }
    }
}
            
