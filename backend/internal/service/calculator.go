package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"calculator/api"

	"connectrpc.com/connect"
)

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) Calculate(
	ctx context.Context,
	req *connect.Request[api.CalculateRequest],
) (*connect.Response[api.CalculateResponse], error) {
	request := req.Msg
	var result float64
	var err error

	switch request.Operation {
	case "add":
		result = request.A + request.B
		log.Printf("执行加法运算: %f + %f = %f", request.A, request.B, result)
	case "subtract":
		result = request.A - request.B
		log.Printf("执行减法运算: %f - %f = %f", request.A, request.B, result)
	case "multiply":
		result = request.A * request.B
		log.Printf("执行乘法运算: %f * %f = %f", request.A, request.B, result)
	case "divide":
		if request.B == 0 {
			err = errors.New("除数不能为零")
			log.Printf("错误: 尝试除以零 (被除数: %f)", request.A)
		} else {
			result = request.A / request.B
			log.Printf("执行除法运算: %f / %f = %f", request.A, request.B, result)
		}
	default:
		err = fmt.Errorf("不支持的操作: %s", request.Operation)
		log.Printf("错误: %v", err)
	}

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&api.CalculateResponse{
		Result: result,
	})
	return res, nil
}
