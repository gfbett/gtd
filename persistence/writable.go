package persistence

type Writable interface {
	ToStorableString() string
	LoadFromStorableString(data string) bool
	FileName() string
}
