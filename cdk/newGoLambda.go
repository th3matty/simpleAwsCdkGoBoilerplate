package main

/*
	// 	"github.com/aws/aws-cdk-go/awscdk/awslambdago"
	awslambdago.NewGoFunction(stack, jsii.String("myGoHandler"), &awslambdago.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./basic-api-app/handler-function-get"),
		Bundling: &awslambdago.BundlingOptions{
			GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
		},
	})

	// hasherservice

		lambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String(fmt.Sprintf("hasher-lambda-%s", functionName)), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String(fmt.Sprintf("hasher-%s", functionName)),
		Entry:        jsii.String(fmt.Sprintf("../functions/%s/main.go", functionName)),
		LogRetention: logRetention,
	})
*/

import (
	"fmt"
	//"os"

	//"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewGoLambda(stack constructs.Construct, functionName string) awscdklambdagoalpha.GoFunction {
	// logRetention := awslogs.RetentionDays_FIVE_DAYS

	// if env := os.Getenv("ENVIRONMENT"); env == "live" {
	// 	logRetention = awslogs.RetentionDays_TWO_WEEKS
	// }

	lambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String(fmt.Sprintf("string-container-lambda-%s", functionName)), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String(fmt.Sprintf("string-container-%s", functionName)),
		Entry:        jsii.String(fmt.Sprintf("../functions/%s/main.go", functionName)),
		//LogRetention: logRetention,
	})

	return lambda
}
