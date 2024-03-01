package adapter

type Foo interface {
}

type Bar interface {
}

type Krog struct {
	Name string
}

func test() {

	var k Krog
	map_Foo_Bar := make(map[Foo]Bar)
	map_Foo_Bar[&struct{}{}] = &k

}
