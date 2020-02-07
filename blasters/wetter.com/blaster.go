package wetter_com

import (
	"fmt"
	"github.com/deletescape/noise/pkg/data"
	"github.com/deletescape/noise/pkg/entities"
	"github.com/deletescape/noise/pkg/runner"
	"github.com/deletescape/noise/pkg/wyrand"
	"github.com/valyala/fasthttp"
	"github.com/wI2L/jettison"
	"math/rand"
	"time"
)

func init() {
	runner.RegisterBlaster(&WetterComBlaster{})
}

type WetterComBlaster struct {
	entities.Blaster

	ctx   *entities.BlasterContext
	adIds []string
	err error
}

func (b *WetterComBlaster) Name() string {
	return "wetter.com/open-locate"
}

func (b *WetterComBlaster) Init(ctx *entities.BlasterContext) {
	b.ctx = ctx
}

func (b *WetterComBlaster) ReinitData(i uint64) {
	idCount := wyrand.IntN(31) + 20
	b.adIds = make([]string, idCount)
	for j := 0; j < idCount; j++ {
		b.adIds[j] = data.GetAdId()
	}
}

func (b *WetterComBlaster) Blast(i uint64) {
	recordCount := wyrand.IntN(10) + 3
	records := make([]Record, recordCount)

	adId := b.adIds[wyrand.IntN(len(b.adIds))]
	dev := data.GetAndroidDevice()
	os := data.GetAndroidVersion(entities.LOLLIPOP, entities.Q)
	apn := data.GetApn()
	wifi := data.GetWiFi()

	minTime := time.Now().UTC().AddDate(-4, 0, 0).Unix()
	maxTime := time.Now().UTC().AddDate(1, 0, 0).Unix()
	startTime := rand.Int63n(maxTime-minTime) + minTime

	for j := 0; j < recordCount; j++ {
		loc := data.GetLocation()
		now := startTime - int64(wyrand.Uint64())
		records[j] = Record{
			Latitude:           loc.Lat,
			Longitude:          loc.Lon,
			HorizontalAccuracy: loc.XAccuracy,
			VerticalAccuracy:   fmt.Sprintf("%f", loc.YAccuracy),
			UtcTimestamp:       now,
			Speed:              loc.Speed,
			SpeedAccuracy:      fmt.Sprintf("%f", loc.SpeedAccuracy),
			Course:             loc.Orientation,
			CourseAccuracy:     fmt.Sprintf("%f", loc.OrientationAccuracy),
			Altitude:           int(loc.Altitude),
			LocationSource:     sources[wyrand.IntN(len(sources))],
			AppVersionCode:     now - int64(wyrand.Uint64N(100000)+10000),
			DeviceManufacturer: dev.Manufacturer,
			DeviceModel:        dev.Model,
			OsVersion:          fmt.Sprintf("Android_%s", os.Release),
			QueryAccuracy:      accuracies[wyrand.IntN(len(accuracies))],
			CarrierName:        apn.Carrier,
			WifiSsid:           wifi.SSID,
			WifiBssid:          wifi.BSSID,
			ConfigHash:         configHash(),
			LocationMethod:     methods[wyrand.IntN(len(methods))],
			LocationContext:    contexts[wyrand.IntN(len(contexts))],
			AdID:               adId,
		}
	}
	var json []byte
	json, b.err = jettison.Marshal(records)
	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)
	args.SetBytesV("json", json)
	_, _, b.err = b.ctx.Client.Post(nil, "https://yes-sir.wetter.com/open-locate.json", args)
}

func (b *WetterComBlaster) Sleep(i uint64) {
	if b.err != nil {
		time.Sleep(500 * time.Millisecond)
	} else {
		time.Sleep(7 * time.Millisecond)
	}
}

func (b *WetterComBlaster) Error() error {
	defer func() {
		b.err = nil
	}()
	return b.err
}

// Seed data
func configHash() string {
	dispatch := wyrand.IntN(8)*100 + 21000
	gcmtask := (wyrand.IntN(6) + 1) * 100
	gcmquery := (wyrand.IntN(8) + 1) * 10
	gcmfastest := (wyrand.IntN(8) + 1) * 10
	latch := (wyrand.IntN(11) + 10) * 100
	return fmt.Sprintf("on|passive|Dispatch:%d|GcmTask:%d|GcmQuery:%d|GcmFastest:%d|Latch:%d|Gcm_OS_28_DP_5_DCP_-1_EXCL_11_26_4_9", dispatch, gcmtask, gcmquery, gcmfastest, latch)
}

var sources = []string{
	"OPEN_LOCATE_GCM",
	"OPEN_LOCATE_GCM_EXECUTE",
	"DISPATCH",
	"APP",
	"APP_UPDATE",
	"BOARDING",
	"OPEN_LOCATE_NO_PERMISSION",
	"OPEN_LOCATE_UPDATE",
	"PLOTPROJECTS_RUNNING",
	"UNDEFINED",
	"WIDGET",
}

var methods = []string{
	"gps",
	"network",
	"passive",
	"disabled",
	"service_not_found",
}

var contexts = []string{
	"foreground",
	"background",
}

var accuracies = []string{
	"passive",
	"low",
	"medium",
	"high",
	"not_set",
}
