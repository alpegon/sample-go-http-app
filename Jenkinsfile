pipeline {
    agent { label 'kaniko'}

    stages {
        stage('Build and push to registry') {
            steps {
                container('kaniko') {
                    sh '''executor \
                          --destination=docker.io/serbangilvitu/kaniko-test:$(date -u +%Y-%m-%dT%H%M%S) \
                          --context=git://github.com/serbangilvitu/sample-go-http-app.git#refs/heads/master
                    '''
                }
            }
        }
    }
}