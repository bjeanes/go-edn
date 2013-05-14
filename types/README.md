# EDN Types

A library that contains the data types used by the EDN library.

## Considerations

* Currently the `Map` type is an alias to `map[Value]Value` (`Value` being an 
  interface provided by this library). In Go, `map`s do not support all type 
  as keys. 

  Since `Set`s are currently backed by `map[Value]bool`, the same restriction
  applies to set elements.

  Notably, `Vector`, `Set`, `Map`, etc., can not be map keys or set elements,
  despite being a valid expression in EDN. 

   * Consider using the hash-set and hash-map implementations from 
     [gohash][gohash]. The trade-off here will be further departure from
     native types as well as requiring all EDN types to meet the 
     [`Hasher` and `Equalser` interfaces][gohash interfaces].
