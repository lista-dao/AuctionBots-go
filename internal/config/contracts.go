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
	FlashBuyContract() *common.Address
}

type contracts struct {
	contractsParams
	InteractionAddr *common.Address
	VatAddr         *common.Address
	HayAddr         *common.Address
	DogAddr         *common.Address
	SpotAddr        *common.Address
	Token0Addr      *common.Address
	FlashbuyAddr    *common.Address

	CollateralsAddr []common.Address

	populated bool
	o         sync.Once
}

type contractsParams struct {
	interaction string
	vat         string `populate:"required"`
	hay         string `populate:"required"`
	dog         string `populate:"required"`
	spot        string `populate:"required"`
	token0      string `populate:"required"`
	flashbuy    string
	collaterals string `populate:"required" type:"array"`
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
	flag.StringVar(&contr.flashbuy, "FLASHBUY", "", "flash buy contract")
	flag.Parse()
}

func (params contractsParams) check() error {
	val := reflect.ValueOf(params)
	typ := val.Type()
	for ind := 0; ind < val.NumField(); ind++ {
		field := val.Field(ind)
		fieldTyp := typ.Field(ind)
		populateTag := fieldTyp.Tag.Get("populate")
		if field.IsZero() && populateTag != "required" {
			continue
		}

		typeTag := fieldTyp.Tag.Get("type")
		switch typeTag {
		case "array":
			addresses := strings.Split(field.String(), ",")
			for ind := range addresses {
				if !common.IsHexAddress(addresses[ind]) {
					return fmt.Errorf("contract %s address is bad", fieldTyp.Name)
				}
			}

		default:
			if !common.IsHexAddress(field.String()) {
				return fmt.Errorf("contract %s address is bad", fieldTyp.Name)
			}
		}

	}

	return nil
}

// populateContracts - populate contracts struct
// using contractsParams
func (contr *contracts) populateContracts() {
	contr.o.Do(func() {
		contr.populated = true
		valParams := reflect.ValueOf(contr.contractsParams)
		typParams := valParams.Type()
		valContr := reflect.Indirect(reflect.ValueOf(contr))
		for ind := 0; ind < valParams.NumField(); ind++ {
			fieldTyp := typParams.Field(ind)
			fieldVal := valParams.Field(ind)
			contrFieldName := fmt.Sprintf("%sAddr", strings.Title(fieldTyp.Name))
			contrFieldVal := valContr.FieldByName(contrFieldName)

			typeTag := fieldTyp.Tag.Get("type")
			switch typeTag {
			case "array":
				if contrFieldVal.Kind() != reflect.Slice {
					panic(fmt.Sprintf("%s is not array", contrFieldName))
				}
				addresses := strings.Split(fieldVal.String(), ",")
				for ind := range addresses {
					val := reflect.New(reflect.TypeOf(common.Address{}))
					val.MethodByName("SetBytes").
						Call(
							[]reflect.Value{
								reflect.ValueOf(common.FromHex(addresses[ind])),
							},
						)
					reflect.Append(contrFieldVal, reflect.Indirect(val))
				}

			default:
				contrFieldVal.Set(reflect.New(reflect.TypeOf(common.Address{})))
				contrFieldVal.MethodByName("SetBytes").
					Call(
						[]reflect.Value{
							reflect.ValueOf(common.FromHex(fieldVal.String())),
						},
					)
			}
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

func (contr *contracts) FlashBuyContract() *common.Address {
	if !contr.populated {
		contr.populateContracts()
	}
	return contr.FlashbuyAddr
}
