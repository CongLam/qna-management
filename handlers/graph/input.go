package graph

type InputPoint struct {
	StartX int `json:"start_x" binding:"gte=0,lte=9"`
	EndX   int `json:"end_x" binding:"gte=0,lte=9"`
	StartY int `json:"start_y" binding:"gte=0,lte=9"`
	EndY   int `json:"end_y" binding:"gte=0,lte=9"`
}
