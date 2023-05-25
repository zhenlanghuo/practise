package pair

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"testing"
)

func TestBigInt(t *testing.T) {
	x0 := new(big.Int).SetInt64(350)
	// (x0*x0)/1e4
	a := new(big.Int).Div(new(big.Int).Mul(x0, x0), big1e4)
	// x0*(x0/1e4)
	b := new(big.Int).Mul(x0, new(big.Int).Div(x0, big1e4))

	t.Logf("a: %v", a)
	t.Logf("b: %v", b)
}

func Test__f(t *testing.T) {
	type args struct {
		x0 string
		y  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				x0: "23000000000000000000",
				y:  "13000000000000000000",
			},
			want: "208702000000000000000000",
		},
		{
			name: "test2",
			args: args{
				x0: "55489175361731",
				y:  "52677896467876",
			},
			want: "17",
		},
		{
			name: "test3",
			args: args{
				x0: "55489175361731413",
				y:  "52677896467876313",
			},
			want: "17111579821377",
		},
		{
			name: "test4",
			args: args{
				x0: "5548917",
				y:  "5267789",
			},
			want: "0",
		},
		{
			name: "test5",
			args: args{
				x0: "554891794727745542",
				y:  "52677894442455",
			},
			want: "9000224038308",
		},
		{
			name: "test6",
			args: args{
				x0: "4348823785333",
				y:  "3675860236763433",
			},
			want: "215997",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_x0, _ := new(big.Int).SetString(tt.args.x0, 10)
			_y, _ := new(big.Int).SetString(tt.args.y, 10)
			_want, _ := new(big.Int).SetString(tt.want, 10)
			if got := _f(_x0, _y); !reflect.DeepEqual(got, _want) {
				t.Errorf("_f() = %v, want %v", got, _want)
			}
		})
	}
}

func Test__d(t *testing.T) {
	type args struct {
		x0 string
		y  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				x0: "2312412412512312",
				y:  "124512312131242",
			},
			want: "12472600148",
		},
		{
			name: "test2",
			args: args{
				x0: "864576363",
				y:  "25963709433",
			},
			want: "0",
		},
		{
			name: "test3",
			args: args{
				x0: "864576363434232",
				y:  "25963709433",
			},
			want: "646264165",
		},
		{
			name: "test4",
			args: args{
				x0: "67451237663",
				y:  "5678631267345133323",
			},
			want: "6525270451020",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_x0, _ := new(big.Int).SetString(tt.args.x0, 10)
			_y, _ := new(big.Int).SetString(tt.args.y, 10)
			_want, _ := new(big.Int).SetString(tt.want, 10)
			if got := _d(_x0, _y); !reflect.DeepEqual(got, _want) {
				t.Errorf("_d() = %v, want %v", got, _want)
			}
		})
	}
}

func Test__get_y(t *testing.T) {
	type args struct {
		x0 string
		xy string
		y  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				x0: "235123412123215",
				xy: "3125141312",
				y:  "2342331251",
			},
			want: "23687014821409624",
		},
		{
			name: "test2",
			args: args{
				x0: "676462645123412",
				xy: "97561273",
				y:  "32176834521412",
			},
			want: "5215097524510427",
		},
		{
			name: "test3",
			args: args{
				x0: "676462645123412",
				xy: "97561273231242124",
				y:  "32176834521412",
			},
			want: "5244183474018809377",
		},
		{
			name: "test4",
			args: args{
				x0: "676462645123",
				xy: "9756127322124",
				y:  "321768345214434",
			},
			want: "2434134358860246339",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_x0, _ := new(big.Int).SetString(tt.args.x0, 10)
			_xy, _ := new(big.Int).SetString(tt.args.xy, 10)
			_y, _ := new(big.Int).SetString(tt.args.y, 10)
			_want, _ := new(big.Int).SetString(tt.want, 10)
			if got := _get_y(_x0, _xy, _y); !reflect.DeepEqual(got, _want) {
				t.Errorf("_get_y() = %v, want %v", got, _want)
			}
		})
	}
}

func TestPair__k(t *testing.T) {
	type fields struct {
		token0      common.Address
		token1      common.Address
		decimals0   uint64
		decimals1   uint64
		reserve0    string
		reserve1    string
		stable      bool
		stableFee   string
		volatileFee string
	}
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "1",
				reserve1:    "2",
				stable:      true,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				x: "874278685624",
				y: "235627321241256",
			},
			want: "11437544783326831712520860171883281036026717964846",
		},
		{
			name: "test2",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "1",
				reserve1:    "2",
				stable:      false,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				x: "874278685624",
				y: "235627321241256",
			},
			want: "206003944711909311882903744",
		},
		{
			name: "test3",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "1",
				reserve1:    "2",
				stable:      true,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				x: "4623412512",
				y: "2361623426464",
			},
			want: "60897049753710907243264082400743956698806",
		},
		{
			name: "test4",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "1",
				reserve1:    "2",
				stable:      false,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				x: "4623412512",
				y: "2361623426464",
			},
			want: "10918759298545969517568",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_decimals0 := new(big.Int).SetUint64(tt.fields.decimals0)
			_decimals1 := new(big.Int).SetUint64(tt.fields.decimals1)
			_reserve0, _ := new(big.Int).SetString(tt.fields.reserve0, 10)
			_reserve1, _ := new(big.Int).SetString(tt.fields.reserve1, 10)
			_stableFee, _ := new(big.Int).SetString(tt.fields.stableFee, 10)
			_volatileFee, _ := new(big.Int).SetString(tt.fields.volatileFee, 10)
			_x, _ := new(big.Int).SetString(tt.args.x, 10)
			_y, _ := new(big.Int).SetString(tt.args.y, 10)
			_want, _ := new(big.Int).SetString(tt.want, 10)
			p := &Pair{
				token0:      tt.fields.token0,
				token1:      tt.fields.token1,
				decimals0:   _decimals0,
				decimals1:   _decimals1,
				reserve0:    _reserve0,
				reserve1:    _reserve1,
				stable:      tt.fields.stable,
				stableFee:   _stableFee,
				volatileFee: _volatileFee,
			}
			if got := p._k(_x, _y); !reflect.DeepEqual(got, _want) {
				t.Errorf("_k() = %v, want %v", got, _want)
			}
		})
	}
}

func TestPair_getAmountOut(t *testing.T) {
	type fields struct {
		token0      common.Address
		token1      common.Address
		decimals0   uint64
		decimals1   uint64
		reserve0    string
		reserve1    string
		stable      bool
		stableFee   string
		volatileFee string
	}
	type args struct {
		amountIn string
		tokenIn  common.Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "42358668902742",
				reserve1:    "67342352335454",
				stable:      true,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				amountIn: "2312453",
				tokenIn:  common.BytesToAddress([]byte("123")),
			},
			want: "2367264",
		},
		{
			name: "test2",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "42358668902742",
				reserve1:    "67342352335454",
				stable:      true,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				amountIn: "2312453",
				tokenIn:  common.BytesToAddress([]byte("124")),
			},
			want: "2258008",
		},
		{
			name: "test3",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "42358668902742",
				reserve1:    "67342352335454",
				stable:      false,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				amountIn: "2312453",
				tokenIn:  common.BytesToAddress([]byte("124")),
			},
			want: "1454253",
		},
		{
			name: "test4",
			fields: fields{
				token0:      common.BytesToAddress([]byte("123")),
				token1:      common.BytesToAddress([]byte("124")),
				decimals0:   uint64(1e6),
				decimals1:   uint64(1e6),
				reserve0:    "42358668902742",
				reserve1:    "67342352335454",
				stable:      false,
				stableFee:   "2",
				volatileFee: "2",
			},
			args: args{
				amountIn: "2312453",
				tokenIn:  common.BytesToAddress([]byte("123")),
			},
			want: "3675632",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_decimals0 := new(big.Int).SetUint64(tt.fields.decimals0)
			_decimals1 := new(big.Int).SetUint64(tt.fields.decimals1)
			_reserve0, _ := new(big.Int).SetString(tt.fields.reserve0, 10)
			_reserve1, _ := new(big.Int).SetString(tt.fields.reserve1, 10)
			_stableFee, _ := new(big.Int).SetString(tt.fields.stableFee, 10)
			_volatileFee, _ := new(big.Int).SetString(tt.fields.volatileFee, 10)
			_amountIn, _ := new(big.Int).SetString(tt.args.amountIn, 10)
			_want, _ := new(big.Int).SetString(tt.want, 10)
			p := NewPair(tt.fields.token0, tt.fields.token1, _reserve0, _reserve1, _decimals0, _decimals1, tt.fields.stable, _stableFee, _volatileFee)
			if got := p.getAmountOut(_amountIn, tt.args.tokenIn); !reflect.DeepEqual(got, _want) {
				t.Errorf("getAmountOut() = %v, want %v", got, _want)
			}
		})
	}
}
