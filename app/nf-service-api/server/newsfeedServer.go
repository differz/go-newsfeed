package server

type Mode int

const (
	ModeDebug Mode = iota
	ModeRelease
)

type NewsfeedServer interface {
	Run(httpAddr string)
}