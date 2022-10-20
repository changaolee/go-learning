package cuslog

type Formatter interface {
	Format(entry *Entry) error
}
