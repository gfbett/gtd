package persistence

type Writable interface {
	ToStorableString() string
	LoadFromStorableString(data string)
	FileName() string
}
