pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''echo "Building project'''
                agent { docker { image 'golang' }}
            }
        }
    }
}     

