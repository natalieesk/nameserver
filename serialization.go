package nameserver

import "errors"

const headerLength = 12

func (m *message) deserialize(b []byte) error {
	return nil
}

func (m *message) serialize() ([]byte, error) {
	return nil, nil
}

func extractLabels(b []byte) (l []label, remaining []byte, err error) {

	if b[0] == 0 {
		return nil, nil, errors.New("no question to extract")
	}

	for b[0] != 0 {
		length := b[0]
		lab := label(string(b[1 : length+1]))
		l = append(l, lab)
		b = b[length+1:]
	}

	return l, b[1:], nil // don't return the remaining null byte
}

func extractHeaders(in []byte) ([]byte, error) {
	if len(in) < headerLength {
		return nil, errors.New("Headers are too short")
	}
	return in[headerLength:], nil
}