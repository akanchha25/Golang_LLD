package permission

type Permission interface {
	IsPermitted() bool
}