package generate

//go:generate ../../../build/scripts/generate-aws-interfaces.sh sts STS
//go:generate ../../../build/scripts/generate-aws-interfaces.sh autoscaling ASG
//go:generate ../../../build/scripts/generate-aws-interfaces.sh cloudwatchlogs CloudWatchLogs
