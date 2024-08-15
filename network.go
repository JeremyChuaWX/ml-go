package mlgo

type network struct {
	layers []layer
}

func Network() *network {
	return &network{
		layers: []layer{},
	}
}

func (n *network) AddLayer(l layer) {
	n.layers = append(n.layers, l)
}

func (n *network) Forward(input *matrix) *matrix {
	result := input
	for _, layer := range n.layers {
		result = layer.forward(result)
	}
	return result
}

func (n *network) Backward(loss *matrix) {
	length := len(n.layers)
	y := loss
	for idx := range n.layers {
		pos := length - idx - 1
		layer := n.layers[pos]
		layer.backward(y)
	}
}

func (n *network) Update(lr float64) {
	for _, layer := range n.layers {
		layer.update(lr)
	}
}
