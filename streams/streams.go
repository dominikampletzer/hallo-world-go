package streams

type Any interface{}

type Mapper func(a Any) Any
type Predicate func(a Any) bool
type Accumulator func(a, b Any) Any
type Pair struct {
	name  string
	count int
}
type Iterator interface {
	HasNext() bool
	Next() Any
}
type Stream interface {
	Iterator() Iterator
	Map(m Mapper) Stream
	Filter(p Predicate) Stream
	Reduce(a Accumulator) Any
}

func ToStream(s []Any) Stream {
	return NewSliceStream(s)
}

func NewSliceStream(data []Any) *SliceStream {
	ss := new(SliceStream)
	ss.data = data
	return ss
}

type SliceStream struct {
	data []Any
}

func (s SliceStream) Iterator() Iterator {
	return nil
}
func (s *SliceStream) Map(mapper Mapper) Stream {
	for i, el := range s.data {
		s.data[i] = mapper(el)
		//strings.ToUpper(el.(string)))
	}
	return s
}
func (s *SliceStream) Filter(predicate Predicate) Stream {
	temp := new([]Any)
	for _, el := range s.data {
		if predicate(el) {
			*temp = append(*temp, el)
		}
		//strings.ToUpper(el.(string)))
	}
	s.data = *temp
	return s
}
func (s SliceStream) Reduce(reducer Accumulator) Any {

	var result interface{}
	for i, e := range s.data {
		if i == 0 {
			result = e
		} else {
			result = reducer(result, s.data[i])
		}
	}
	return result
}
