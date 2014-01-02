package eg

// パッケージ内のExampleインターフェース
type Example interface {
	// この型の名前を返す
	Name() string
	// この型の一意な識別子
	id() uint32
}

// Exampleインターフェースを実装した新たな値を生成
func NewExample() Example {
	return new(concreteType)
}