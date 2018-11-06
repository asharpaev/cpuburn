node {
    label 'slave'
    stage('Checkout'){
        checkout scm
    }
    stage('Build') {
        sh "curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.8.11/bin/linux/amd64/kubectl"
        sh "chmod +x ./kubectl"
        sh "./kubectl get pods"
    }

}