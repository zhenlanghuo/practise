package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"practise/test/test_unmarshal_and_unsafe/pb/foody_base"
	"reflect"
	"testing"
	"unsafe"
)

func newStore(id uint64) *foody_base.Store {
	return &foody_base.Store{
		Id:               proto.Uint64(id),
		Name:             proto.String(fmt.Sprintf("%v", id)),
		MerchantId:       proto.Uint64(uint64(rand.Intn(9999999999))),
		TagId:            proto.Uint64(uint64(rand.Intn(9999999999))),
		BrandId:          proto.Uint64(uint64(rand.Intn(9999999999))),
		TobUserId:        proto.Uint64(uint64(rand.Intn(9999999999))),
		IsUseWallet:      foody_base.Store_WALLET_DISABLE.Enum(),
		IsUseMerchantApp: foody_base.Store_APP_DISABLE.Enum(),
		Location: &foody_base.Location{
			State:            proto.String("guangdong"),
			City:             proto.String("shenzhen"),
			District:         proto.String("nanshan"),
			Address:          proto.String("haiyiyuan"),
			Latitude:         proto.Float32(13.35),
			Longitude:        proto.Float32(12.34),
			Remark:           nil,
			PreciseLatitude:  proto.Float64(13.25),
			PreciseLongitude: proto.Float64(13.27),
		},
		PostalCode:           proto.String("1231234125"),
		RegisterPhone:        proto.String("327747270437"),
		Email:                proto.String("zhenlanghuo@qq.com"),
		EmailSource:          nil,
		RegisterTime:         proto.Uint64(12412511321451245),
		Logo:                 proto.String("xxxxxx.com"),
		Banner:               proto.String("xxxxxxxx.cn"),
		PartnerType:          foody_base.PartnerType_PARTNER_LISTED.Enum(),
		CommissionRate:       proto.Uint64(uint64(rand.Intn(999999999))),
		TaxRate:              proto.Uint64(uint64(rand.Intn(999999999))),
		ServiceFee:           proto.Uint64(uint64(rand.Intn(999999999))),
		MinSpend:             proto.Uint64(uint64(rand.Intn(999999999))),
		DeliveryDistance:     proto.Uint64(uint64(rand.Intn(999999999))),
		PreparationTime:      proto.Uint64(uint64(rand.Intn(999999999))),
		ContactPhone:         proto.String("21412513511"),
		Status:               foody_base.StoreStatus_STORE_ACTIVE.Enum(),
		AutoConfirmed:        foody_base.Store_Manual.Enum(),
		AutoConfirmedEnabled: foody_base.Store_Disable.Enum(),
		CreateTime:           proto.Uint64(12412511321451245),
		UpdateTime:           proto.Uint64(12412511321451245),
		RatingTotal:          proto.Uint32(uint32(rand.Intn(999999999))),
		RatingScore:          proto.Float32(4.3),
		OpeningStatus:        foody_base.OpeningStatus_Status_OPEN.Enum(),
		//SurchargeIntervals: &foody_base.Store_SurChargeIntervals{
		//	Intervals: []*foody_base.Store_SurChargeInterval{
		//		{
		//			OrderPriceEnd: proto.Uint64(uint64(rand.Intn(999999999))),
		//			Fee:           proto.Uint64(uint64(rand.Intn(999999999))),
		//		},
		//		{
		//			OrderPriceEnd: proto.Uint64(uint64(rand.Intn(999999999))),
		//			Fee:           proto.Uint64(uint64(rand.Intn(999999999))),
		//		},
		//	},
		//},
		//ServiceChargeFeeRate:     proto.Uint32(uint32(rand.Intn(999999999))),
		//DriverModifyOrderEnabled: proto.Uint32(1),
		//DeliveryDistanceMode:     foody_base.Store_DELIVERY_DISTANCE_MODE_DEFAULT.Enum(),
		//BusinessInfoAdded:        proto.Uint32(1),
		//IsInstantPrep:            proto.Uint32(1),
		//Flag: &foody_base.Store_Flag{
		//	OvertimeOrderMode: foody_base.Store_Flag_OVERTIME_ORDER_MODE_CONFIRM.Enum(),
		//},
		//ModifyOrderMode:    foody_base.Store_MODIFY_ORDER_MODE_NO_EDITING.Enum(),
		//StatusReason:       foody_base.Store_ON_BOARDING.Enum(),
		//StatusReasonRemark: proto.String("xxxxxxxxxxx"),
		//Overlay: &foody_base.Store_Overlay{
		//	LogoImage:   proto.String("xxxxxxxxxxx"),
		//	BannerImage: proto.String("xxxxxxxxxxx"),
		//},
		//ScheduledCommissions: &foody_base.Store_ScheduledCommissions{
		//	ScheduledCommissions: []*foody_base.Store_ScheduledCommission{
		//		{
		//			CommissionRate: proto.Uint32(uint32(rand.Intn(999999999))),
		//			Priority:       proto.Uint32(uint32(rand.Intn(999999999))),
		//			EffectiveTime:  proto.Uint64(uint64(rand.Intn(999999999))),
		//			ExpireTime:     proto.Uint64(uint64(rand.Intn(999999999))),
		//		},
		//	},
		//},
		//EffectiveCommissionRate: proto.Uint32(uint32(rand.Intn(999999999))),
		//VendorId:                proto.Uint64(uint64(rand.Intn(999999999))),
	}
}

func Benchmark_Unmarshal(b *testing.B) {
	v, err := proto.Marshal(newStore(1))
	if err != nil {
		b.Fatalf("marshal failed, err: %v", err)
	}
	store := &foody_base.Store{}
	err = proto.Unmarshal(v, store)
	if err != nil {
		b.Fatalf("marshal failed, err: %v", err)
	}
	b.Logf("%v", store)

	for i := 0; i < b.N; i++ {
		store = &foody_base.Store{}
		proto.Unmarshal(v, store)
	}
}

func Benchmark_Unmarshal2(b *testing.B) {
	v, err := newStore(1).Marshal()
	if err != nil {
		b.Fatalf("marshal failed, err: %v", err)
	}
	store := &foody_base.Store{}
	err = store.Unmarshal(v)
	if err != nil {
		b.Fatalf("marshal failed, err: %v", err)
	}
	b.Logf("%v", store)

	for i := 0; i < b.N; i++ {
		store = &foody_base.Store{}
		store.Unmarshal(v)
	}
}

func Benchmark_Unsafe(b *testing.B) {

	var store *foody_base.Store
	bytes := make([]byte, unsafe.Sizeof(*store))
	store = (*foody_base.Store)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
	store.Id = proto.Uint64(uint64(rand.Intn(100)))
	b.Logf("%v, %v", store, bytes)

	//bigbytes := make([]byte, 4*1024)

	for i := 0; i < b.N; i++ {
		//clone := make([]byte, 4*1024)
		//copy(clone, bigbytes)
		var store_ *foody_base.Store
		store_ = (*foody_base.Store)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
		store_.Id = proto.Uint64(1)
	}
}
