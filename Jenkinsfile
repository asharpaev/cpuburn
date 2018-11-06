node {
    stage('Checkout'){
        checkout scm
    }
    stage('Build') {
        sh "docker build --no-cache"
    }

}