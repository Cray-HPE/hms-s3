@Library('dst-shared@master') _

dockerBuildPipeline {
        repository = "cray"
        imagePrefix = "hms"
        app = "s3"
        name = "hms-s3"
        description = "Cray HMS S3 code."
        dockerfile = "Dockerfile"
        slackNotification = ["", "", false, false, true, true]
        product = "internal"
}
