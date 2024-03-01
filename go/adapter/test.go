package adapter

type Foo struct {
}

type Bar struct {
}

func (*Bar) Get() map[*Foo]*Foo {
	return nil
}

type MapAnyToAny interface {
	Get() map[any]any
}

// var _ MapAnyToAny = &(Bar{})
