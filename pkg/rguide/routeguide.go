package rguide

import (
	context "context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	sync "sync"
	"time"

	"google.golang.org/protobuf/proto"
)

//go:embed sample_data.json
var jsonDBFile []byte

func NewServer() *RGServer {
	s := &RGServer{
		routeNotes: make(map[string][]*RouteNote),
	}
	s.loadFeatures("")
	return s
}

type RGServer struct {
	UnimplementedRouteGuideServer
	savedFeatures []*Feature // read-only after initialized

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*RouteNote
}

// GetFeature returns the feature at the given point.
func (s *RGServer) GetFeature(ctx context.Context, point *Point) (*Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	// No feature was found, return an unnamed feature
	return &Feature{Location: point}, nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *RGServer) ListFeatures(rect *Rectangle, stream RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points,  number of known features visited, total distance traveled, and
// total time spent.
func (s *RGServer) RecordRoute(stream RouteGuide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *Point
	startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if lastPoint != nil {
			distance += calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *RGServer) RouteChat(stream RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)

		s.mu.Lock()
		s.routeNotes[key] = append(s.routeNotes[key], in)
		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		rn := make([]*RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

// loadFeatures loads features from a JSON file.
func (s *RGServer) loadFeatures(filePath string) {
	// var data []byte
	// if filePath != "" {
	// 	var err error
	// 	data, err = ioutil.ReadFile(filePath)
	// 	if err != nil {
	// 		log.Fatalf("Failed to load default features: %v", err)
	// 	}
	// } else {
	// 	data = exampleData
	// }
	data := jsonDBFile
	if err := json.Unmarshal(data, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// calcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func calcDistance(p1 *Point, p2 *Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Lat) / CordFactor)
	lat2 := toRadians(float64(p2.Lat) / CordFactor)
	lng1 := toRadians(float64(p1.Long) / CordFactor)
	lng2 := toRadians(float64(p2.Long) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func inRange(point *Point, rect *Rectangle) bool {
	left := math.Min(float64(rect.Lo.Long), float64(rect.Hi.Long))
	right := math.Max(float64(rect.Lo.Long), float64(rect.Hi.Long))
	top := math.Max(float64(rect.Lo.Lat), float64(rect.Hi.Lat))
	bottom := math.Min(float64(rect.Lo.Lat), float64(rect.Hi.Lat))

	if float64(point.Long) >= left &&
		float64(point.Long) <= right &&
		float64(point.Lat) >= bottom &&
		float64(point.Lat) <= top {
		return true
	}
	return false
}

func serialize(point *Point) string {
	return fmt.Sprintf("%d %d", point.Lat, point.Long)
}
