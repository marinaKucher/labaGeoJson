package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fogleman/gg"
	geojson "github.com/paulmach/go.geojson"
)

func main() {

	dc := gg.NewContext(1366, 1024)
	geojsonData, err := ioutil.ReadFile("2.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fc, err := geojson.UnmarshalFeatureCollection(geojsonData)
	if err != nil {
		fmt.Println(err.Error())
	}

	dc.InvertY()
	dc.Scale(10, 10)
	//dc.SetLineWidth(10)
	dc.SetRGB(120, 120, 120)

	//для линии или многих линий фаил первый
	for _, feature := range fc.Features {
		for i := 0; i < len(feature.Geometry.Polygon); i++ {
			for j := 0; j < len(feature.Geometry.Polygon[i]); j++ {
				dc.LineTo(feature.Geometry.Polygon[i][j][0], feature.Geometry.Polygon[i][j][1])
			}

		}
		fmt.Println("fdsd")
	}

	dc.SetHexColor("f00")
	dc.Fill()

	/*	//для линии или многих линий фаил первый
		for _, feature := range fc.Features {
			for i := 0; i < len(feature.Geometry.LineString)-1; i++ {
				dc.SetRGB(120, 120, 120)
				dc.DrawLine(feature.Geometry.LineString[i][0], feature.Geometry.LineString[i][1], feature.Geometry.LineString[i+1][0], feature.Geometry.LineString[i+1][1])
				dc.Stroke()
			}

		}*/

	/*для точки или нескольких точек фаил называется три
	for _, feature := range fc.Features {
		for i := 0; i < len(feature.Geometry.Point)-1; i++ {
			dc.DrawPoint(feature.Geometry.Point[i], feature.Geometry.Point[i+1], 20.0)
			dc.Stroke()
		}
	}*/
	dc.SavePNG("out.png")
}
