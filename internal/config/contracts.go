package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"sync"
)

var _ configer = (*contracts)(nil)

type contracter interface {
	InteractionContract() common.Address
	SpotContract() common.Address
	VatContract() common.Address
	DogContract() common.Address
	HayContract() common.Address
	Token0Contract() common.Address
}

type contracts struct {
	contractsParams
	InteractionAddr *common.Address
	VatAddr         *common.Address
	HayAddr         *common.Address
	DogAddr         *common.Address
	SpotAddr        *common.Address
	Token0Addr      *common.Address

	populated bool
	o         sync.Once
}

type contractsParams struct {
	interaction string
	vat         string
	hay         string
	dog         string
	spot        string
	token0      string
}

func (contr *contracts) validate() error {
	return errors.Wrap(contr.check(), "failed to validate webhook")
}

func (contr *contracts) populate() {
	flag.StringVar(&contr.dog, "DOG", "", "dog contract address")
	flag.StringVar(&contr.interaction, "INTERACTION", "", "interaction contract address")
	flag.StringVar(&contr.vat, "VAT", "", "vat contract")
	flag.StringVar(&contr.hay, "HAY", "", "hay contract")
	flag.StringVar(&contr.spot, "SPOT", "", "spot contract")
	flag.StringVar(&contr.token0, "TOKEN0", "", "token 0 contract")
	flag.Parse()
}

func (params contractsParams) check() error {
	val := reflect.ValueOf(params)
	typ := val.Type()
	for ind := 0; ind < val.NumField(); ind++ {
		field := val.Field(ind)
		if !common.IsHexAddress(field.String()) {
			return fmt.Errorf("contract %s address is bad", typ.Field(ind).Name)
		}
	}

	return nil
}

// populateContracts - populate contracts into contracts struct
// using contracts params
func (contr *contracts) populateContracts() {
	contr.o.Do(func() {
		contr.populated = true
		valParams := reflect.ValueOf(contr.contractsParams)
		typParams := valParams.Type()
		val := reflect.Indirect(reflect.ValueOf(contr))
		for ind := 0; ind < valParams.NumField(); ind++ {
			fieldName := fmt.Sprintf("%sAddr", strings.Title(typParams.Field(ind).Name))
			valField := val.FieldByName(fieldName)
			valField.Set(reflect.New(reflect.TypeOf(common.Address{})))
			valField.MethodByName("SetBytes").
				Call(
					[]reflect.Value{
						reflect.ValueOf(common.FromHex(valParams.Field(ind).String())),
					},
				)
		}
	})
}

func (contr *contracts) InteractionContract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.InteractionAddr
}

func (contr *contracts) VatContract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.VatAddr
}

func (contr *contracts) DogContract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.DogAddr
}

func (contr *contracts) HayContract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.HayAddr
}

func (contr *contracts) SpotContract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.SpotAddr
}

func (contr *contracts) Token0Contract() common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return *contr.Token0Addr
}
