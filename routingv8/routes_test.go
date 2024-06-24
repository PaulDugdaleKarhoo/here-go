package routingv8_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"go.einride.tech/here/routingv8"
	"gotest.tools/v3/assert"
)

type RoutesMock struct {
	requestRawQuery string
	responseStatus  int
	responseBody    routingv8.RoutesResponse
	error           *routingv8.HereErrorResponse
}

func (c *RoutesMock) Do(req *http.Request) (*http.Response, error) {
	c.requestRawQuery = req.URL.RawQuery
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	b, err := json.Marshal(c.responseBody)
	if err != nil {
		return nil, err
	}
	if c.error != nil {
		b, err = json.Marshal(c.error)
		if err != nil {
			return nil, err
		}
	}
	r := bytes.NewReader(b)
	return &http.Response{
		StatusCode:    c.responseStatus,
		Header:        headers,
		Body:          io.NopCloser(r),
		ContentLength: int64(len(b)),
	}, nil
}

func TestRoutingervice_Routes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Einride Gothenburg.
	origin := routingv8.GeoWaypoint{
		Elevation: 1,
		Lat:       57.707752,
		Long:      11.949767,
	}
	// Einride Stockholm.
	destination := routingv8.GeoWaypoint{
		Elevation: 2,
		Lat:       59.337492,
		Long:      18.063672,
	}

	exp := routingv8.RoutesResponse{
		Routes: []routingv8.Route{
			{
				ID: "route-1",
				Sections: []routingv8.Section{
					{
						ID:   "section-1",
						Type: "veicle",
						Departure: routingv8.VehicleDeparture{
							Place: routingv8.Place{
								Type:             "place",
								Location:         origin,
								OriginalLocation: origin,
							},
						},
						Arrival: routingv8.VehicleDeparture{
							Place: routingv8.Place{
								Type:             "place",
								Location:         destination,
								OriginalLocation: destination,
							},
						},
						Summary: routingv8.Summary{
							Duration:     243,
							Length:       1206,
							BaseDuration: 136,
						},
					},
				},
			},
		},
	}
	httpClient := RoutesMock{responseBody: exp, responseStatus: 200}
	routingClient := routingv8.NewClient(&httpClient)

	got, err := routingClient.Routing.Routes(ctx, &routingv8.RoutesRequest{
		Origin:        origin,
		Destination:   destination,
		TransportMode: routingv8.TransportModeCar,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, &exp, got)
}

func TestRoutingervice_Routes_QueryParams(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Einride Gothenburg.
	origin := routingv8.GeoWaypoint{
		Lat:  57.707752,
		Long: 11.949767,
	}
	// Einride Stockholm.
	destination := routingv8.GeoWaypoint{
		Lat:  59.337492,
		Long: 18.063672,
	}

	for _, tt := range []struct {
		name     string
		request  *routingv8.RoutesRequest
		expected string
		errStr   string
	}{
		{
			name: "minimal",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary&transportMode=car",
		},
		{
			name: "multiple avoid areas",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				AvoidAreas: []routingv8.AreaFeature{
					routingv8.AreaFeatureFerry,
					routingv8.AreaFeatureTollRoad,
					routingv8.AreaFeatureTunnel,
					routingv8.AreaFeatureControlledAccessHighway,
				},
			},
			expected: "avoid%5Bfeatures%5D=ferry%2CtollRoad%2Ctunnel%2CcontrolledAccessHighway" +
				"&destination=59.337492%2C18.063672&origin=57.707752%2C11.949767&return=summary&transportMode=car",
		},
		{
			name: "multiple return attributes",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				Return: []routingv8.ReturnAttribute{
					routingv8.SummaryReturnAttribute,
					routingv8.PolylineReturnAttribute,
				},
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary%2Cpolyline&transportMode=car",
		},
		{
			name: "with spans",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				Return: []routingv8.ReturnAttribute{
					routingv8.SummaryReturnAttribute,
					routingv8.PolylineReturnAttribute,
				},
				Spans: []routingv8.SpanAttribute{
					routingv8.SpanAttributeNames,
					routingv8.SpanAttributeMaxSpeed,
				},
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary%2Cpolyline&spans=names%2CmaxSpeed&transportMode=car",
		},
		{
			name: "with spans without wanted polyline returned",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				Return: []routingv8.ReturnAttribute{
					routingv8.SummaryReturnAttribute,
				},
				Spans: []routingv8.SpanAttribute{
					routingv8.SpanAttributeNames,
					routingv8.SpanAttributeMaxSpeed,
				},
			},
			errStr: "spans parameter also requires that the polyline option is set in the return parameter",
		},
		{
			name: "with vias",
			request: &routingv8.RoutesRequest{
				Origin:      origin,
				Destination: destination,
				Via: []routingv8.GeoWaypoint{
					{
						Lat:  57.695538,
						Long: 11.992594,
					},
					{
						Lat:  59.323410,
						Long: 18.096137,
					},
				},
				TransportMode: routingv8.TransportModeCar,
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary&transportMode=car" +
				"&via=57.695538%2C11.992594&via=59.32341%2C18.096137",
		},
		{
			name: "with routingmode",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				RoutingMode:   routingv8.RoutingModeShort,
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary&routingMode=short&transportMode=car",
		},
		{
			name: "with trafficmode",
			request: &routingv8.RoutesRequest{
				Origin:        origin,
				Destination:   destination,
				TransportMode: routingv8.TransportModeCar,
				TrafficMode:   routingv8.TrafficModeDisabled,
			},
			expected: "destination=59.337492%2C18.063672&origin=57.707752%2C11.949767" +
				"&return=summary&trafficMode=disabled&transportMode=car",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := RoutesMock{}
			routingClient := routingv8.NewClient(&client)

			_, err := routingClient.Routing.Routes(ctx, tt.request)
			if tt.errStr != "" {
				assert.ErrorContains(t, err, tt.errStr)
			}
			assert.Equal(t, client.requestRawQuery, tt.expected)
		})
	}
}

func TestRoutingervice_Routes_Error(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Einride Gothenburg.
	origin := routingv8.GeoWaypoint{
		Lat:  57.707752,
		Long: 11.949767,
	}
	// Einride Stockholm.
	destination := routingv8.GeoWaypoint{
		Lat:  59.337492,
		Long: 18.063672,
	}

	exp := routingv8.HereErrorResponse{
		Title:  "Mocked Error",
		Status: 400,
	}

	httpClient := RoutesMock{responseStatus: 400, error: &exp}
	routingClient := routingv8.NewClient(&httpClient)

	_, err := routingClient.Routing.Routes(ctx, &routingv8.RoutesRequest{
		Origin:        origin,
		Destination:   destination,
		TransportMode: routingv8.TransportModeCar,
	})
	var responseError *routingv8.ResponseError
	assert.Check(t, errors.As(err, &responseError))
	assert.DeepEqual(t, responseError.Response, &exp)
	assert.Check(t, responseError.HTTPBody != "")
}
