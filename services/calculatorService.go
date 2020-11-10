package services

import (
	"errors"
	"math"
	"quasarproject/entities"
	"strings"
)

// CalculatorServiceInterface is an interface
type CalculatorServiceInterface interface {
	GetLocation(distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}

const (
	kenobiDistanceX    = -500.0
	kenobiDistanceY    = -200.0
	skywalkerDistanceX = 100.0
	skywalkerDistanceY = -100.0
	satoDistanceX      = 500.0
	satoDistanceY      = 100.0
	returnError        = 0.0
)

type calculatorService struct {
}

// NewCalculatorService implements service
func NewCalculatorService() CalculatorServiceInterface {
	return &calculatorService{}
}

func (c *calculatorService) GetLocation(distances ...float32) (x, y float32) {

	if len(distances) < 3 {
		return returnError, returnError
	}

	var kenobiSateliteDistance float32 = distances[0]
	var skywalkerSateliteDistance float32 = distances[1]
	var satoSateliteDistance float32 = distances[2]

	if validDistances := c.distanceValidator(kenobiSateliteDistance, skywalkerSateliteDistance, satoSateliteDistance); validDistances == false {
		return returnError, returnError
	}

	/** Trilateration algorithm

		Step 1:
		(x-x1)^2 + (y-y1)^2 = r1^2
		(x-x2)^2 + (y-y2)^2 = r2^2
		(x-x3)^2 + (y-y3)^2 = r3^2

		Step 2:
		x^2 - (2*x1*x) + x1^2 + y^2 - (2*y1*y) + y1^2 = r1^2
		x^2 - (2*x2*x) + x2^2 + y^2 - (2*y2*y) + y2^2 = r2^2
		x^2 - (2*x3*x) + x3^2 + y^2 - (2*y3*y) + y3^2 = r3^2

		Step 3:
		(-2*x1+2*x2)*x + (-2*y1+2*y2)*y = r1^2-r2^2-x1^2+x2^2-y1^2+y2^2
		(-2*x1+2*x2) or (-2*y1+2*y2)  -> (firstCalculation)
		r1^2-r2^2-x1^2+x2^2-y1^2+y2^2 -> (secondCalculation)

		Step 4:
		A*x + B*y = C
		D*x + E*y = F

		Step 5:
		x = (C*E - F*B) / (E*A - B*D)
		y = (C*D - A*F) / (B*D - A*E)
	**/

	A := c.firstCalculation(skywalkerDistanceX, kenobiDistanceX)                                                                                          // X2 - X1
	B := c.firstCalculation(skywalkerDistanceY, kenobiDistanceY)                                                                                          // Y2 - Y1
	C := c.secondCalculation(kenobiSateliteDistance, skywalkerSateliteDistance, kenobiDistanceX, skywalkerDistanceX, kenobiDistanceY, skywalkerDistanceY) // r1^2 - r2^2 - x1^2 + x2^2 - y1^2 + y2^2
	D := c.firstCalculation(satoDistanceX, skywalkerDistanceX)                                                                                            // X3 - X2
	E := c.firstCalculation(satoDistanceY, skywalkerDistanceY)                                                                                            // Y3 - Y2
	F := c.secondCalculation(skywalkerSateliteDistance, satoSateliteDistance, skywalkerDistanceX, satoDistanceX, skywalkerDistanceY, satoDistanceY)       // r2^2 - r3^2 - x2^2 + x3^2 - y2^2 + y3^2
	x = float32(math.Round((C*E - F*B) / (E*A - B*D)))
	y = float32(math.Round((C*D - A*F) / (B*D - A*E)))

	return x, y

}

func (c *calculatorService) firstCalculation(a float32, b float32) float64 {
	// We cast the result to float64
	return float64((2 * a) - (2 * b))
}

func (c *calculatorService) secondCalculation(r1 float32, r2 float32, x1 float32, x2 float32, y1 float32, y2 float32) float64 {
	// We cast each parameter to float64
	return math.Pow(float64(r1), 2) - math.Pow(float64(r2), 2) - math.Pow(float64(x1), 2) + math.Pow(float64(x2), 2) - math.Pow(float64(y1), 2) + math.Pow(float64(y2), 2)
}

func (c *calculatorService) getVectorNorm(x float32, y float32) float32 {
	// We cast each parameter to float64, and the result to float32
	return float32(math.Pow(float64(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)), 0.5))
}

func (c *calculatorService) distanceValidator(kenobiSateliteDistance float32, skywalkerSateliteDistance float32, satoSateliteDistance float32) bool {

	kenobiToSkywalker := c.getVectorNorm(kenobiDistanceX-skywalkerDistanceX, kenobiDistanceY-skywalkerDistanceY)
	skywalkerToSato := c.getVectorNorm(kenobiDistanceX-skywalkerDistanceX, kenobiDistanceY-skywalkerDistanceY)
	kenobiToSato := c.getVectorNorm(kenobiDistanceX-skywalkerDistanceX, kenobiDistanceY-skywalkerDistanceY)

	error1 := c.validator(kenobiToSkywalker, kenobiSateliteDistance, skywalkerSateliteDistance)
	error2 := c.validator(skywalkerToSato, skywalkerSateliteDistance, satoSateliteDistance)
	error3 := c.validator(kenobiToSato, kenobiSateliteDistance, satoSateliteDistance)

	return error1 == nil && error2 == nil && error3 == nil
}

func (c *calculatorService) validator(distance float32, r1 float32, r2 float32) error {

	if distance > (r1 + r2) {
		return errors.New("No se tocan")
	}

	if float64(distance) < math.Abs(float64(r1-r2)) {
		return errors.New("Estan incluidos")
	}
	return nil
}

func (c *calculatorService) GetMessage(messages ...[]string) (msg string) {
	var messageInfo entities.Message
	messageInfo.MessageInformation = make(map[int]string)

	for _, info := range messages {
		c.checkMessage(&messageInfo, info)
	}

	return c.mapMessage(messageInfo.MessageInformation)
}

func (c *calculatorService) checkMessage(data *entities.Message, info []string) {
	dataLen := len(info) - 1
	for key := range info {
		if data.MessageInformation[key] == "" || data.MessageInformation[key] == " " {
			data.MessageInformation[key] = info[dataLen-key]
		}
	}
}

func (c *calculatorService) mapMessage(finalMessage map[int]string) string {
	var auxMessage string

	finalLen := len(finalMessage) - 1
	for range finalMessage {
		if finalMessage[finalLen] != "" {
			auxMessage += finalMessage[finalLen] + " "
		}
		finalLen--
	}

	// We delete the last space
	return strings.TrimRight(auxMessage, " ")
}
