package component

var St *Storage

type Storage struct {
	Map map[string]string
}

func init() {
	St = new(Storage)
}

func (s *Storage) Receive(data interface{}) error {
	return nil
}