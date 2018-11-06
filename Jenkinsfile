node {
    label 'slave'
    stage('Checkout'){
        checkout scm
    }
    stage('Build') {
        sh "curl https://download.docker.com/linux/static/stable/x86_64/docker-18.06.1-ce.tgz | tar xvz docker/docker"
        sh "docker/docker build --no-cache -t registry:5000/cpuburn:latest ."
        sh "docker/docker push registry:5000/cpuburn:latest"
    }

    stage('List pods') {
    withKubeConfig([credentialsId: '${KUBE-CERT}',
                    caCertificate: '${CA-CERT}',
                    serverUrl: '${KUBE-URL}',
                    contextName: 'default'
                    ]) {
            sh 'kubectl get pods'
        }
    }

}