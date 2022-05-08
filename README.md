# GO-JWT-MONGO-MYSQL-CRUD-GIN-SERVER

## This Repo consists of Go server with complete JWT, MongoDB and, MySQL CRUD APIs


### Basic Go built in packages:
  * fmt
  * os
  * log
  * time
  * bufio
  * strconv
  * sort (available in slice not in array)
  * crypto
  * math
  * strings...etc

### Basic Tips:
##### * Array vs Slice: Arrays will have fixed size where as slices have a dynamic size(can use append() to add new items).
##### * How to delete an item from slice?: we can use append() and slice operator (:) to delete an item from the slice.
        for example:
        var myslice = []string{"a", "b", "c"}
        Now to delete an item "b" from myslice, we can do like below
        myslice = append(myslice[:1], myslice[2:]...)
        Note**: three dots at the end(spread operator)
