package Nextcorev1Native

//EnumerationGroup ...
type EnumerationGroup struct {

    Name string `json:"name"`
    Documentation *Documentation `json:"documentation"`
    Items []string `json:"items"`
}

//EnumerationGroupReader ...
type EnumerationGroupReader struct {
    _enumerationgroup *EnumerationGroup
    documentation *DocumentationReader
    items *native.VectorStringReader
}

//Name ...
func (g *EnumerationGroupReader) Name() (value string) {
	value = g._enumerationgroup.Name
	return
}

//Documentation ...
func (g *EnumerationGroupReader) Documentation() Nextcorev1.DocumentationReader {

	if g.documentation != nil {
		return g.documentation
	}

	return nil
}

//Items ...
func (g *EnumerationGroupReader) Items() hybrids.VectorStringReader {
	if g.items != nil {
		return g.items
	}
	return nil
}

//NewEnumerationGroupReader ...
func NewEnumerationGroupReader(g *EnumerationGroupReader) Nextcorev1.EnumerationGroupReader{
	if g!=nil{
		return &EnumerationGroupReader{_group:g}
	}
	return nil
}

type VectorEnumerationGroupReader struct {
    _vector  []*EnumerationGroupReader
}

func (vg *VectorEnumerationGroupReader) Len() (size int) {
    size = len(vg._vector)
    return
}

func (vg *VectorEnumerationGroupReader) Get(i int) (item Nextcorev1.EnumerationGroupReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vector)}
		return
	}

	if i > len(vg._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vector)}
		return
	}

	item = vg._vector[i]
	return


}


func NewVectorEnumerationGroupReader(v hybrids.VectorTableReader) Nextcorev1.VectorEnumerationGroupReader {
    if v == nil {
        return nil
    }
    return &VectorEnumerationGroupReader{_vectorHybrid: v}
}
