package Nextcorev1



//ResourceReader ...
type ResourceReader interface {

    //RID get resource id
    RID() ResourceIDReader

    //Metadata ...
    Metadata() (MetadataReader, error)

    //Fields ...
    Fields() VectorFieldReader

}

//VectorResourceReader ...
type VectorResourceReader interface {

    // Returns the current size of this vector
    Len() int

    //Get the item in the position i, if i < Len(),
    //if item does not exist should return the default value for the underlying data type
    //when i > Len() should return an VectorInvalidIndexError
    Get(i int) (item ResourceReader, err error)
}
