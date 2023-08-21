package export

type Exporter interface {
	Export() error
}
