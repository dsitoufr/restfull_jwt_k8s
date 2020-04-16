pipeline {
   
   agent none
   
    environment {
        HELLO='restfull_jwt_k8s/hello'
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
               //sh 'cp -r ${WORKSPACE}/*  ${GOPATH/src/github}'
                //sh 'ls  ${WORKSPACE}'
               sh 'ls  ${GOPATH}/src/github}'
               
                //sh 'cd ${HELLO}'
                //sh 'go build'
               /* clean up our workspace */
                deleteDir() 
            }
        }
    }
}
