package network

import (
	"github.com/Centny/gwf/routing"
	"me.car/db"
	"time"
)

func DriverAddRoute(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var startTime int64
	var sPlace string
	var ePlace string
	var sLat float64
	var sLng float64
	var eLat float64
	var eLng float64
	var sCity string
	var eCity string
	var distance float64
	t := time.Now()
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		startTime,R|I,R:0;
		sPlace,R|S,L:0;
		ePlace,R|S,L:0;
		sLat,R|F,R:0;
		sLng,R|F,R:0;
		eLat,R|F,R:0;
		eLng,R|F,R:0;
		sCity,R|S,L:0;
		eCity,R|S,L:0;
		distance,R|F,R:0;
		`, &uid, &startTime, &sPlace, &ePlace, &sLat, &sLng, &eLat, &eLng, &sCity, &eCity, &distance)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else if t.Unix()-startTime < 1800 {
		return hs.MsgResE(1, "invalid startTime , need > 30min from now")
	} else {
		s, err := db.DriverAddRoute(uid, startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity, distance)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}

}

// func PassengerFindDriver(hs *routing.HTTPSession) routing.HResult {

// }

// func PassengerJoinRoute(hs *routing.HTTPSession) routing.HResult {

// }

// func DriverCancelRoute(hs *routing.HTTPSession) routing.HResult {

// }

// func PassengerCancelRoute(hs *routing.HTTPSession) routing.HResult {

// }
