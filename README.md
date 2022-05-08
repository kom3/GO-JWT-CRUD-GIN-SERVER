# GO-JWT-MONGO-MYSQL-CRUD-GIN-SERVER

## This Repo consists of Go server with complete JWT, MongoDB and, MySQL CRUD APIs


### Basic Go built in packages:
  * fmt
  * os        Ex: file, err := os.Create("path/to/file.txt");    myenv := os.Getenv("MYENV")
  * log
  * time
  * bufio
  * strconv
  * sort (available in slice not in array)
  * crypto
  * math
  * io         Ex: length, err := io.WriteString("Write this line in to a file");
  * ioutil     Ex: databytes, err := ioutil.ReadFile("./path/to/file.txt");
  * strings...etc

### Basic Tips:
##### * Array vs Slice: Arrays will have fixed size where as slices have a dynamic size(can use append() to add new items).
##### * How to delete an item from slice?: we can use append() and slice operator (:) to delete an item from the slice.
        for example:
        var myslice = []string{"a", "b", "c"}
        Now to delete an item "b" from myslice, we can do like below
        myslice = append(myslice[:1], myslice[2:]...)
        Note**: three dots at the end(spread operator)
        
##### * What are the available data types in GO?
        1. Basic type: Numbers, strings, and booleans come under this category.
        2. Aggregate type: Array and structs come under this category.
        3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
        4. Interface type

##### * For strings we should always use double quotes("mystring"), single quotes not allowed (~~'mystring'~~).
