package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	concurrent := []float64{1, 10, 100, 1000}

	// Data from the report
	latencyBeforeData := plotter.XYs{
		{concurrent[0], 16.2},
		{concurrent[1], 162.5},
		{concurrent[2], 1580.4},
		{concurrent[3], 14800.3},
	}
	latencyAfterData := plotter.XYs{
		{concurrent[0], 0.39},
		{concurrent[1], 3.6},
		{concurrent[2], 31.2},
		{concurrent[3], 300.5},
	}
	throughputBeforeData := plotter.XYs{
		{concurrent[0], 61.7},
		{concurrent[1], 61.5},
		{concurrent[2], 63.2},
		{concurrent[3], 67.6},
	}
	throughputAfterData := plotter.XYs{
		{concurrent[0], 2564.1},
		{concurrent[1], 2777.8},
		{concurrent[2], 3205.1},
		{concurrent[3], 3322.6},
	}


	titleFontSize := vg.Points(24)
	labelFontSize := vg.Points(18)
	tickFontSize := vg.Points(14)
	legendFontSize := vg.Points(16)
	markerSize := vg.Length(6)

	// Latency Before Index
	p1 := plot.New()
	p1.Title.Text = "Latency Before Index"
	p1.Title.TextStyle.Font.Size = titleFontSize
	p1.X.Label.Text = "Concurrent Connections"
	p1.X.Label.TextStyle.Font.Size = labelFontSize
	p1.Y.Label.Text = "Average Latency (ms)"
	p1.Y.Label.TextStyle.Font.Size = labelFontSize
	p1.X.Tick.Label.Font.Size = tickFontSize
	p1.Y.Tick.Label.Font.Size = tickFontSize
	line1, points1, err := plotter.NewLinePoints(latencyBeforeData)
	if err != nil {
		log.Fatal(err)
	}
	points1.Shape = draw.CircleGlyph{}
	points1.Radius = markerSize
	p1.Add(line1, points1)
	p1.Legend.Add("Latency", points1)
	p1.Legend.TextStyle.Font.Size = legendFontSize
	if err := p1.Save(20*vg.Inch, 10*vg.Inch, "latency_before.png"); err != nil {
		log.Fatal(err)
	}

	// Latency After Index
	p2 := plot.New()
	p2.Title.Text = "Latency After Index"
	p2.Title.TextStyle.Font.Size = titleFontSize
	p2.X.Label.Text = "Concurrent Connections"
	p2.X.Label.TextStyle.Font.Size = labelFontSize
	p2.Y.Label.Text = "Average Latency (ms)"
	p2.Y.Label.TextStyle.Font.Size = labelFontSize
	p2.X.Tick.Label.Font.Size = tickFontSize
	p2.Y.Tick.Label.Font.Size = tickFontSize
	line2, points2, err := plotter.NewLinePoints(latencyAfterData)
	if err != nil {
		log.Fatal(err)
	}
	points2.Shape = draw.CircleGlyph{}
	points2.Radius = markerSize
	p2.Add(line2, points2)
	p2.Legend.Add("Latency", points2)
	p2.Legend.TextStyle.Font.Size = legendFontSize
	if err := p2.Save(20*vg.Inch, 10*vg.Inch, "latency_after.png"); err != nil {
		log.Fatal(err)
	}

	// Throughput Before Index
	p3 := plot.New()
	p3.Title.Text = "Throughput Before Index"
	p3.Title.TextStyle.Font.Size = titleFontSize
	p3.X.Label.Text = "Concurrent Connections"
	p3.X.Label.TextStyle.Font.Size = labelFontSize
	p3.Y.Label.Text = "Requests per Second (Req/s)"
	p3.Y.Label.TextStyle.Font.Size = labelFontSize
	p3.X.Tick.Label.Font.Size = tickFontSize
	p3.Y.Tick.Label.Font.Size = tickFontSize
	line3, points3, err := plotter.NewLinePoints(throughputBeforeData)
	if err != nil {
		log.Fatal(err)
	}
	points3.Shape = draw.CircleGlyph{}
	points3.Radius = markerSize
	p3.Add(line3, points3)
	p3.Legend.Add("Throughput", points3)
	p3.Legend.TextStyle.Font.Size = legendFontSize
	if err := p3.Save(20*vg.Inch, 10*vg.Inch, "throughput_before.png"); err != nil {
		log.Fatal(err)
	}

	// Throughput After Index
	p4 := plot.New()
	p4.Title.Text = "Throughput After Index"
	p4.Title.TextStyle.Font.Size = titleFontSize
	p4.X.Label.Text = "Concurrent Connections"
	p4.X.Label.TextStyle.Font.Size = labelFontSize
	p4.Y.Label.Text = "Requests per Second (Req/s)"
	p4.Y.Label.TextStyle.Font.Size = labelFontSize
	p4.X.Tick.Label.Font.Size = tickFontSize
	p4.Y.Tick.Label.Font.Size = tickFontSize
	line4, points4, err := plotter.NewLinePoints(throughputAfterData)
	if err != nil {
		log.Fatal(err)
	}
	points4.Shape = draw.CircleGlyph{}
	points4.Radius = markerSize
	p4.Add(line4, points4)
	p4.Legend.Add("Throughput", points4)
	p4.Legend.TextStyle.Font.Size = legendFontSize
	if err := p4.Save(20*vg.Inch, 10*vg.Inch, "throughput_after.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("Graphs saved: latency_before.png, latency_after.png, throughput_before.png, throughput_after.png")
}
