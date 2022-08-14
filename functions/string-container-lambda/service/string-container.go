package service

import (
	"github.com/th3matty/simpleAwsCdkGoBoilerplate/functions/string-container-lambda/model/"
	"go.uber.org/zap"
)

type stringContainer struct {
	logger *zap.Logger
}

type StringContainer interface {
	StringContainerInput(input model.InputPost)
}
