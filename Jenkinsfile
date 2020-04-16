pipeline {
    agent { docker { image 'golang' } }

    environment {
        HELLO='hello'
        AUTH='auth'
    }
    
    stages {
        stage('build') {
            agent { docker { image 'golang' }}
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
