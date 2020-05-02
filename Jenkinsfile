pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                    echo "Building project"
                '''
                docker.image("golang:1.8.0-alpine").inside("-v ${pwd()}:${goPath}") {
                     sh 'go version'
                }
            }
        }
    }
}     

