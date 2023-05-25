package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gogo/protobuf/proto"
	"practise/test/test_high_concurrent_read_from_redis/pb/foody_base"
	"runtime"
	"sync/atomic"

	//"github.com/golang/protobuf/proto"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var redisClient *redis.Client

func main() {
	runtime.GOMAXPROCS(4)
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	maxLifeTime := time.Minute * 1

	//redisClient = redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   0, // use default DB
	//
	//})
	//fmt.Println(redisClient.Get("1").Val())

	//writeStoreToRedis()

	store := newStore(uint64(1))
	value, _ := proto.Marshal(store)

	store = &foody_base.Store{}
	proto.Unmarshal(value, store)
	fmt.Println(store)

	fmt.Println("!!")
	quit := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	count := int64(0)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
				}

				//pipeline := redisClient.Pipeline()
				//for j := 1; j <= 2500; j++ {
				//	pipeline.Get(fmt.Sprintf("store:%v", j))
				//}
				//cmders, err := pipeline.Exec()
				//if err != nil {
				//	fmt.Printf("pipeline.Exec() faile, err: %v\n", err)
				//	return
				//}
				//for _, cmder := range cmders {
				//	if cmder.Err() != nil {
				//		fmt.Printf(fmt.Sprintf("pipeline.Exec() faile, err: %v\n", cmder.Err()))
				//		return
				//	}
				//	//val := cmder.(*redis.StringCmd).Val()
				//	//store := &foody_base.StoreMin{}
				//	//err = proto.Unmarshal([]byte(val), store)
				//	//if err != nil {
				//	//	fmt.Printf(fmt.Sprintf("json.Unmarshal faile, err: %v\n", err))
				//	//	return
				//	//}
				//	//fmt.Printf("store: %v\n", store)
				//}

				s := &foody_base.Store{}
				proto.Unmarshal(value, s)

				atomic.AddInt64(&count, 1)
			}
		}()
	}

	go func() {
		timer := time.NewTimer(maxLifeTime)
		select {
		case sig := <-sigs:
			fmt.Println()
			fmt.Println(sig)
		case <-timer.C:
			fmt.Println("到时间")
		}

		close(quit)
	}()

	fmt.Println("awaiting signal")
	<-quit
	wg.Wait()
	fmt.Println("exiting")
	fmt.Println("qps", float64(count)/float64(maxLifeTime/time.Second))
}

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

func writeStoreToRedis() {
	for i := 1; i <= 5000; i++ {
		store := newStore(uint64(i))
		value, err := proto.Marshal(store)
		if err != nil {
			fmt.Printf(fmt.Sprintf("writeStoreToRedis json.Marshal faile, err: %v\n", err))
			return
		}
		err = redisClient.Set(fmt.Sprintf("store:%v", i), value, time.Hour*24*30).Err()
		if err != nil {
			fmt.Printf(fmt.Sprintf("writeStoreToRedis redisClient.Set faile, err: %v\n", err))
			return
		}
	}
}
