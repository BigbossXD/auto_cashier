package services

import (
	"strconv"
	"testing"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/responses"
)

func TestCalulateChangeSuccess(t *testing.T) {

	priceTotal := float32(220.25)
	reciveTotal := float32(250.00)
	configs := []models.CashierConfigs{}
	moneyValue := []float32{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}
	maximumAmount := []int32{10, 20, 15, 20, 30, 20, 20, 20, 50}
	currentAmount := []int32{5, 12, 10, 10, 10, 10, 10, 10, 20}
	for i, v := range moneyValue {
		configs = append(configs, models.CashierConfigs{
			MachineId:     1,
			MoneyValue:    v,
			MaximumAmount: maximumAmount[i],
			CurrentAmount: currentAmount[i],
		})
	}

	want := []responses.ChangeItemRequest{
		{ID: 1, MoneyValue: 1000, Amount: 0},
		{ID: 2, MoneyValue: 500, Amount: 0},
		{ID: 3, MoneyValue: 100, Amount: 0},
		{ID: 4, MoneyValue: 50, Amount: 0},
		{ID: 5, MoneyValue: 20, Amount: 1},
		{ID: 6, MoneyValue: 10, Amount: 0},
		{ID: 7, MoneyValue: 5, Amount: 1},
		{ID: 8, MoneyValue: 1, Amount: 4},
		{ID: 9, MoneyValue: 0.25, Amount: 3},
	}

	got, b := calulateChange(priceTotal, reciveTotal, configs)

	if b {
		for i, v := range got {
			if v.Amount != want[i].Amount {
				t.Errorf("got %q, wanted %q", got[0].Amount, want[0].Amount)
			}
		}
	}

}

func TestCalulateChangeError(t *testing.T) {

	priceTotal := float32(220.25)
	reciveTotal := float32(250.00)
	configs := []models.CashierConfigs{}
	moneyValue := []float32{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}
	maximumAmount := []int32{10, 20, 15, 20, 30, 20, 20, 20, 50}
	currentAmount := []int32{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, v := range moneyValue {
		configs = append(configs, models.CashierConfigs{
			MachineId:     1,
			MoneyValue:    v,
			MaximumAmount: maximumAmount[i],
			CurrentAmount: currentAmount[i],
		})
	}

	want := false

	_, b := calulateChange(priceTotal, reciveTotal, configs)

	if b != want {
		t.Errorf("got %q, wanted %q", strconv.FormatBool(b), strconv.FormatBool(want))
	}

}
