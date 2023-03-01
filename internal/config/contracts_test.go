package config

import (
	"strings"
	"testing"
)

func Test_contractsParams_check(t *testing.T) {
	type fields struct {
		Interaction string
		Vat         string
		Hay         string
		Dog         string
		Spot        string
		Token0      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "bad addresses",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "good addresses",
			fields: fields{
				Interaction: "0x0000000000000000000000000000000000000000",
				Vat:         "0x0000000000000000000000000000000000000000",
				Hay:         "0x0000000000000000000000000000000000000000",
				Dog:         "0x0000000000000000000000000000000000000000",
				Spot:        "0x0000000000000000000000000000000000000000",
				Token0:      "0x0000000000000000000000000000000000000000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &contractsParams{
				interaction: "",
				vat:         tt.fields.Vat,
				hay:         tt.fields.Hay,
				dog:         tt.fields.Dog,
				spot:        tt.fields.Spot,
				token0:      tt.fields.Token0,
				collaterals: "0x0000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000",
			}
			if err := params.check(); (err != nil) != tt.wantErr {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_contractsParams_populate(t *testing.T) {
	type fields struct {
		Interaction string
		Vat         string
		Hay         string
		Dog         string
		Spot        string
		Token0      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "good addresses",
			fields: fields{

				Interaction: "0x0000000000000000000000000000000000000000",
				Vat:         "0x1000000000000000000000000000000000000000",
				Hay:         "0x1000000000000000000000000000000000000000",
				Dog:         "0x1000000000000000000000000000000000000000",
				Spot:        "0x1000000000000000000000000000000000000000",
				Token0:      "0x1000000000000000000000000000000000000000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := contractsParams{
				interaction: "",
				vat:         tt.fields.Vat,
				hay:         tt.fields.Hay,
				dog:         tt.fields.Dog,
				spot:        tt.fields.Spot,
				token0:      tt.fields.Token0,
				collaterals: "0x0000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000",
			}

			contracts := contracts{contractsParams: params}
			contracts.populateContracts()
			if strings.Compare(strings.ToLower(contracts.VatContract().String()), strings.ToLower(tt.fields.Vat)) != 0 {
				t.Errorf("failed to populate contract: want %s have %s", params.vat, contracts.VatContract().String())
			}

			if strings.Compare(strings.ToLower(contracts.InteractionContract().String()), strings.ToLower(tt.fields.Interaction)) != 0 {
				t.Errorf("failed to populate contract: want %s have %s", params.vat, contracts.VatContract().String())
			}
		})
	}
}
