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
    stage('Build') {
      steps {
        container('golang') {
          sh """
          ln -s `pwd` /go/src/restfull_jwt_k8s
          cd /go/src/restfull_jwt_k8s/hello
          go get github.com/rs/cors
          go test
          """
        }
      }
    }
  }
}
