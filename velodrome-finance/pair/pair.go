package pair

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var big1e18 = new(big.Int).SetUint64(uint64(1e18))
var big1e4 = new(big.Int).SetUint64(uint64(1e4))
var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

type Pair struct {
	token0      common.Address
	token1      common.Address
	decimals0   *big.Int
	decimals1   *big.Int
	reserve0    *big.Int
	reserve1    *big.Int
	stable      bool
	stableFee   *big.Int
	volatileFee *big.Int
}

func NewPair(token0 common.Address, token1 common.Address, reserve0, reserve1, decimals0, decimals1 *big.Int, stable bool, stableFee, volatileFee *big.Int) *Pair {
	return &Pair{
		token0:      token0,
		token1:      token1,
		decimals0:   decimals0, // 根据 token0 设置
		decimals1:   decimals1, // 根据 token1 设置
		reserve0:    reserve0,
		reserve1:    reserve1,
		stable:      stable,
		stableFee:   stableFee,
		volatileFee: volatileFee,
	}
}

func _f(x0, y *big.Int) *big.Int {

	// x0*x0
	_x02 := new(big.Int).Mul(x0, x0)
	// (x0*x0)/1e18
	_x02Div1e18 := new(big.Int).Div(_x02, big1e18)
	// x0*x0/1e18*x0/1e18 -> (((x0*x0)/1e18)*x0)/1e18
	_x0 := new(big.Int).Div(new(big.Int).Mul(_x02Div1e18, x0), big1e18)

	// y*y
	_y2 := new(big.Int).Mul(y, y)
	// (y*y)/1e18
	_y2Div1e18 := new(big.Int).Div(_y2, big1e18)
	// y*y/1e18*y/1e18 -> (((y*y)/1e18)*y)/1e18
	_y := new(big.Int).Div(new(big.Int).Mul(_y2Div1e18, y), big1e18)

	// x0*(y*y/1e18*y/1e18)/1e18+(x0*x0/1e18*x0/1e18)*y/1e18
	return new(big.Int).Add(
		new(big.Int).Div(new(big.Int).Mul(x0, _y), big1e18),
		new(big.Int).Div(new(big.Int).Mul(_x0, y), big1e18),
	)
}

func _d(x0, y *big.Int) *big.Int {

	// x0*x0
	_x02 := new(big.Int).Mul(x0, x0)
	// (x0*x0)/1e18
	_x02Div1e18 := new(big.Int).Div(_x02, big1e18)
	// x0*x0/1e18*x0/1e18 -> (((x0*x0)/1e18)*x0)/1e18
	_x0 := new(big.Int).Div(new(big.Int).Mul(_x02Div1e18, x0), big1e18)

	// y*y
	_y2 := new(big.Int).Mul(y, y)
	// (y*y)/1e18
	_y2Div1e18 := new(big.Int).Div(_y2, big1e18)

	// 3*x0
	_3x0 := new(big.Int).Mul(new(big.Int).SetInt64(3), x0)

	// 3*x0*(y*y/1e18)/1e18+(x0*x0/1e18*x0/1e18)
	return new(big.Int).Add(
		new(big.Int).Div(new(big.Int).Mul(_3x0, _y2Div1e18), big1e18),
		_x0,
	)
}

func _get_y(x0, xy, y *big.Int) *big.Int {
	for i := 0; i < 255; i++ {
		y_prev := new(big.Int).Set(y)
		k := _f(x0, y)
		// k < xy
		if k.Cmp(xy) == -1 {
			// dy = (xy - k)*1e18/_d(x0, y)
			dy := new(big.Int).Div(new(big.Int).Mul(new(big.Int).Sub(xy, k), big1e18), _d(x0, y))
			// y = y + dy
			y = new(big.Int).Add(y, dy)
		} else {
			// dy = (k - xy)*1e18/_d(x0, y)
			dy := new(big.Int).Div(new(big.Int).Mul(new(big.Int).Sub(k, xy), big1e18), _d(x0, y))
			// y = y - dy
			y = new(big.Int).Sub(y, dy)
		}
		// y > y_prev
		if y.Cmp(y_prev) == 1 {
			// y - y_prev <= 1
			if new(big.Int).Sub(y, y_prev).Cmp(bigOne) != 1 {
				return y
			}
		} else {
			// y_prev - y <= 1
			if new(big.Int).Sub(y_prev, y).Cmp(bigOne) != 1 {
				return y
			}
		}
	}

	return y
}

func (p *Pair) getAmountOut(amountIn *big.Int, tokenIn common.Address) *big.Int {
	_reserve0, _reserve1 := p.reserve0, p.reserve1
	// amountIn -= amountIn * PairFactory(factory).getFee(stable) / 10000; // remove fee from amount received
	// PairFactory(factory).getFee(stable) -> stable ? stableFee : volatileFee
	fee := p.volatileFee
	if p.stable {
		fee = p.stableFee
	}
	amountIn = new(big.Int).Sub(amountIn, new(big.Int).Div(new(big.Int).Mul(amountIn, fee), big1e4))
	//log.Printf("amountIn=%v", amountIn)
	// return _getAmountOut(amountIn, tokenIn, _reserve0, _reserve1);
	return p._getAmountOut(amountIn, tokenIn, _reserve0, _reserve1)
}

func (p *Pair) _getAmountOut(amountIn *big.Int, tokenIn common.Address, _reserve0, _reserve1 *big.Int) *big.Int {
	if p.stable {
		xy := p._k(_reserve0, _reserve1)
		//log.Printf("xy=%v", xy)

		// _reserve0 = _reserve0 * 1e18 / decimals0
		_reserve0 = new(big.Int).Div(new(big.Int).Mul(_reserve0, big1e18), p.decimals0)
		//log.Printf("_reserve0=%v", _reserve0)

		// _reserve1 = _reserve1 * 1e18 / decimals1
		_reserve1 = new(big.Int).Div(new(big.Int).Mul(_reserve1, big1e18), p.decimals1)
		//log.Printf("_reserve1=%v", _reserve1)

		// (uint reserveA, uint reserveB) = tokenIn == token0 ? (_reserve0, _reserve1) : (_reserve1, _reserve0);
		reserveA, reserveB := _reserve1, _reserve0
		if tokenIn == p.token0 {
			reserveA, reserveB = _reserve0, _reserve1
		}
		//log.Printf("reserveA=%v", reserveA)
		//log.Printf("reserveB=%v", reserveB)

		// amountIn = tokenIn == token0 ? amountIn * 1e18 / decimals0 : amountIn * 1e18 / decimals1;
		_amountIn := new(big.Int).Div(new(big.Int).Mul(amountIn, big1e18), p.decimals1)
		if tokenIn == p.token0 {
			_amountIn = new(big.Int).Div(new(big.Int).Mul(amountIn, big1e18), p.decimals0)
		}
		amountIn = _amountIn
		//log.Printf("amountIn=%v", amountIn)

		// uint y = reserveB - _get_y(amountIn+reserveA, xy, reserveB);
		y := new(big.Int).Sub(reserveB, _get_y(new(big.Int).Add(amountIn, reserveA), xy, reserveB))
		//log.Printf("y=%v", y)

		// return y * (tokenIn == token0 ? decimals1 : decimals0) / 1e18;
		decimals := p.decimals0
		if tokenIn == p.token0 {
			decimals = p.decimals1
		}
		return new(big.Int).Div(new(big.Int).Mul(y, decimals), big1e18)
	} else {
		// (uint reserveA, uint reserveB) = tokenIn == token0 ? (_reserve0, _reserve1) : (_reserve1, _reserve0);
		reserveA, reserveB := _reserve1, _reserve0
		if tokenIn == p.token0 {
			reserveA, reserveB = _reserve0, _reserve1
		}
		// return amountIn * reserveB / (reserveA + amountIn);
		return new(big.Int).Div(new(big.Int).Mul(amountIn, reserveB), new(big.Int).Add(reserveA, amountIn))
	}
}

func (p *Pair) _k(x, y *big.Int) *big.Int {
	if p.stable {
		_x := new(big.Int).Div(new(big.Int).Mul(x, big1e18), p.decimals0)
		_y := new(big.Int).Div(new(big.Int).Mul(y, big1e18), p.decimals1)
		_a := new(big.Int).Div(new(big.Int).Mul(_x, _y), big1e18)
		_b := new(big.Int).Add(new(big.Int).Div(new(big.Int).Mul(_x, _x), big1e18), new(big.Int).Div(new(big.Int).Mul(_y, _y), big1e18))
		return new(big.Int).Div(new(big.Int).Mul(_a, _b), big1e18) // x3y+y3x >= k
	} else {
		return new(big.Int).Mul(x, y) // xy >= k
	}
}
