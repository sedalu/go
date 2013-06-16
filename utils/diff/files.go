package diff

type Files struct {
	lpath, rpath string
}

func NewFilesDiffer(lpath, rpath string) *Files {
	d := new(Files)
	d.lpath = lpath
	d.rpath = rpath
	return d
}
