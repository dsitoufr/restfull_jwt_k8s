pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''echo "Building project'''
                agent { docker { image 'maven:3.3.3' } }
            }
        }
    }
}     

