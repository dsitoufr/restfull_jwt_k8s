pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }

            steps {
                //create project directory
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github'

                //copy files from jenkins workspace to project directory
                sh 'rsync -ax "*tmp" ${WORKSPACE}/* ${GOPATH/src/github}
                  ${GOPATH}/src/github/'

                //sh 'go build'
            }
        }
    }
}
