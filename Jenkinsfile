pipeline {
    agent { label 'kaniko'}

    stages {
        stage('Build without pushing to registry') {
            steps {
                container('kaniko') {
                    sh "pwd && ls -l"
                    sh '''executor \
                          --no-push \
                          --context=git://github.com/serbangilvitu/sample-go-http-app.git#refs/heads/master
                    '''
                }
            }
        }
    }
}