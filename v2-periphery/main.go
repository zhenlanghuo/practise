package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

var bigZero = big.NewInt(0)

func main() {
	//big1, _ := new(big.Int).SetString("1e15", 10)
	////big1, _ := new(big.Int).SetString("12334", 10)
	//fmt.Println("big1 is:", big1.Bits(), big1)
	//
	//115000000000000000000000000000000000000000000000000000000000000000000000000000
	//
	//big2 := big1.Uint64()
	//fmt.Println("big2 is: ", big2)
	//
	//fmt.Println("maxUint64: ", uint64(math.MaxUint64))

	//a, _ := new(big.Int).SetString("115000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	//b, _ := new(big.Int).SetString("15000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	//
	//fmt.Println(new(big.Int).Add(a, b))


	a, _ := new(big.Int).SetString("400", 10)
	b, _ := new(big.Int).SetString("401", 10)
	fmt.Println(new(big.Int).Sub(a, b))

	//amountIn, _ := new(big.Int).SetString("1844674407370955161500", 10)
	//reserveIn, _ := new(big.Int).SetString("1844674407370955161522", 10)
	//reserveOut, _ := new(big.Int).SetString("1844674407370955161533", 10)
	//log.Printf("amountIn: %v, reserveIn: %v, reserveOut: %v", amountIn, reserveIn, reserveOut)
	//amountOut, err := getAmountOut(amountIn, reserveIn, reserveOut)
	//log.Printf("amountOut: %v, err: %v", amountOut, err)

}

func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	if amountIn.Cmp(bigZero) != 1 {
		return nil, errors.New("UniswapV2Library: INSUFFICIENT_INPUT_AMOUNT")
	}

	if reserveIn.Cmp(bigZero) != 1 || reserveOut.Cmp(bigZero) != 1 {
		return nil, errors.New("UniswapV2Library: INSUFFICIENT_LIQUIDITY")
	}

	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(1000)), amountInWithFee)
	amountOut := numerator.Div(numerator, denominator)

	log.Printf("amountInWithFee: %v, numerator: %v, denominator: %v, amountOut: %v", amountInWithFee, numerator, denominator, amountOut)
	return amountOut, nil
}
