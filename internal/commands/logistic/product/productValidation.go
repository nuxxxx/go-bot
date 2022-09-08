package commands

import (
	"errors"
	"github/nuxxxcake/go-bot/internal/model/logistic"
	"reflect"
	"strconv"
	"strings"
)

type validProductFields struct {
	Title    string
	Price    float64
	Quantity int
}

func productValidation(fieldString string) (*validProductFields, error) {
	fields := strings.Split(fieldString, ",")

	if len(fields) != reflect.TypeOf(logistic.Product{}).NumField()-1 {
		return nil, errors.New("not enough fields")
	}

	// parsing Title
	title := fields[0]

	// parsing price
	price, err := strconv.ParseFloat(fields[1], 64)

	if err != nil {
		return nil, errors.New("invalid price")
	}

	// parsing quantity
	quantity, err := strconv.Atoi(fields[2])

	if err != nil {
		return nil, errors.New("invalid quantity")
	}

	return &validProductFields{
		Title:    title,
		Price:    price,
		Quantity: quantity,
	}, nil
}
