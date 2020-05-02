pipeline {
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
"""
    }
  }
  stages {
    stage('Install dependencies'){
        steps {
            container('golang') {
                ssh """
                    go version
                    go get -u github.com/rs/cors
                    go get -u github.com/braintree/manners 
                    go get -u github.com/dgrijalva/jwt-go 
                    go get -u github.com/gorilla/mux
                    go get -u golang.org/x/crypto/bcrypt
                """
            }
        }
    }
    stage('Unit tests') {
      steps {
        container('golang') {
          sh """
          ln -s `pwd` /go/src/restfull_jwt_k8s
          cd /go/src/restfull_jwt_k8s/hello
          go test
          """
        }
      }
    }
  }
}
