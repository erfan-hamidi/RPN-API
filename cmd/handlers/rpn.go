package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)


type RPNRequest struct {
    Expression string `json:"expression"`
}


type RPNResponse struct {
    Result float64 `json:"result"`
    Error  string  `json:"error,omitempty"`
}


func splitExpression(expression string) []string {
    return strings.Fields(expression)
}

func EvaluateRPN(c echo.Context) error {
    req := new(RPNRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, RPNResponse{Error: "Invalid request"})
    }

    result, err := calculateRPN(req.Expression)
    if err != nil {
        return c.JSON(http.StatusBadRequest, RPNResponse{Error: err.Error()})
    }

    return c.JSON(http.StatusOK, RPNResponse{Result: result})
}


func calculateRPN(expression string) (float64, error) {
    stack := []float64{}
    tokens := splitExpression(expression)

    for _, token := range tokens {
        switch token {
        case "+":
            if len(stack) < 2 {
                return 0, fmt.Errorf("invalid RPN expression")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a+b)
        case "-":
            if len(stack) < 2 {
                return 0, fmt.Errorf("invalid RPN expression")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a-b)
        case "*":
            if len(stack) < 2 {
                return 0, fmt.Errorf("invalid RPN expression")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a*b)
        case "/":
            if len(stack) < 2 {
                return 0, fmt.Errorf("invalid RPN expression")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a/b)
        default:
            num, err := strconv.ParseFloat(token, 64)
            if err != nil {
                return 0, fmt.Errorf("invalid token: %s", token)
            }
            stack = append(stack, num)
        }
    }

    if len(stack) != 1 {
        return 0, fmt.Errorf("invalid RPN expression")
    }

    return stack[0], nil
}
