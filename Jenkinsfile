pipeline {
    agent { docker { image 'golang:1.14.1' } }

    environment {
        HELLO='hello'
        AUTH='auth'
    }
    
    stages {
        stage('build') {
            steps {
                 sh 'go version'
                //create project directory
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github'

                //copy files from jenkins workspace to project directory
                sh 'cp -r ${WORKSPACE}/* ${GOPATH/src/github}'

                //sh 'go build'
            }
        }
    }
}
