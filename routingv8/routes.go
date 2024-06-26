package routingv8

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Routes returns all possible routes between origin and destination.
// See https://developer.here.com/documentation/routing-api/dev_guide/topics/send-request.html#send-a-request
// for details about other parameters.
func (s *RoutingService) Routes(
	ctx context.Context,
	req *RoutesRequest,
) (_ *RoutesResponse, err error) {
	tm := req.TransportMode.String()
	if tm == invalid || tm == unspecified {
		return nil, fmt.Errorf("invalid transportmode")
	}

	u, err := s.URL.Parse("routes")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	returns := make([]string, 0, len(req.Return))
	if len(req.Return) > 0 {
		for _, attribute := range req.Return {
			returns = append(returns, string(attribute))
		}
	} else {
		returns = []string{string(SummaryReturnAttribute)}
	}
	values.Add("return", strings.Join(returns, ","))
	if req.DepartureTime != "" {
		values.Add("departureTime", req.DepartureTime)
	}
	values.Add("transportMode", tm)
	values.Add("origin", fmt.Sprintf("%v,%v", req.Origin.Lat, req.Origin.Long))
	values.Add("destination", fmt.Sprintf("%v,%v", req.Destination.Lat, req.Destination.Long))
	for _, via := range req.Via {
		values.Add("via", fmt.Sprintf("%v,%v", via.Lat, via.Long))
	}
	if len(req.Spans) > 0 {
		if !returnContains(req.Return, PolylineReturnAttribute) {
			return nil, errors.New("spans parameter also requires that the polyline option is set in the return parameter")
		}
		spanStrings := make([]string, 0, len(req.Spans))
		for _, span := range req.Spans {
			spanStrings = append(spanStrings, string(span))
		}
		values.Add("spans", strings.Join(spanStrings, ","))
	}
	if req.AvoidAreas != nil {
		areas := make([]string, 0, len(req.AvoidAreas))
		for _, area := range req.AvoidAreas {
			a := area.String()
			if a == invalid {
				return nil, fmt.Errorf("invalid avoid area")
			}
			if a != unspecified {
				areas = append(areas, a)
			}
		}
		values.Add("avoid[features]", strings.Join(areas, ","))
	}
	rm := req.RoutingMode.String()
	if rm == invalid {
		return nil, fmt.Errorf("invalid routingmode")
	}
	if rm != unspecified {
		values.Add("routingMode", rm)
	}
	trm := req.TrafficMode.String()
	if trm == invalid {
		return nil, fmt.Errorf("invalid trafficmode")
	}
	if trm != unspecified {
		values.Add("trafficMode", trm)
	}
	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp RoutesResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func returnContains(requested []ReturnAttribute, needle ReturnAttribute) bool {
	for _, attr := range requested {
		if attr == needle {
			return true
		}
	}
	return false
}
