pipeline {
   
   agent { docker { image 'golang' }}
   
    environment {
        HELLO='restfull_jwt_k8s/hello'
        AUTH='auth'
    }
    
    stages {
        stage('build') {
            steps {
                 //sh 'go version'
                //create project directory
                // sh 'cd ${GOPATH}/src'
                //sh 'mkdir -p ${GOPATH}/src/github'
               
                bash '''#!/bin/bash
                 echo "hello world" 
                 go version
                 cd ${GOPATH}/src
                '''

                //copy files from jenkins workspace to project directory
               // sh 'cp -r ${WORKSPACE}/restfull_jwt_k8s  ${GOPATH/src/github}'
                //sh 'cd ${HELLO}'
                //sh 'go build'
            }
        }
    }
}
