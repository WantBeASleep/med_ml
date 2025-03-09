package node_segment

type CreateNodeWithSegmentArg struct {
	Node struct {
		UziID     string
		Ai        bool
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}

	Segments []struct {
		ImageID   string
		Contor    []byte
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}
}