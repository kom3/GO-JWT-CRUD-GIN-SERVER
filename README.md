# GO-JWT-MONGO-MYSQL-CRUD-GIN-SERVER

## This Repo consists of Go server with complete JWT, MongoDB and, MySQL CRUD APIs


### 1. Basic Go built in packages:
  ##### * fmt
  ##### * os        
      Ex: 
      file, err := os.Create("path/to/file.txt");
      myenv := os.Getenv("MYENV");
      
  ##### * log
  ##### * time
  ##### * bufio
  ##### * strconv
  ##### * sort (available in slice not in array)
  ##### * crypto
  ##### * math
  ##### * io         
      Ex: length, err := io.WriteString("Write this line in to a file");
      
  ##### * ioutil     
      Ex: databytes, err := ioutil.ReadFile("./path/to/file.txt");
      
  ##### * strings...etc



### 2. Basic Tips:
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

##### * Different ways of initializig a struct in GO.

        type mystruct struct{
        Name string
        Age int
        }
        
        Initialization can be done like below,
        
        Method: 1
        var mystructobj1 mystruct{"Name":"Manju", "Age":24}
        
        Method: 2
        var mystructobj2 mystruct{"Manju", 24}
        
        Method: 3
        var mystructobj3 mystruct;
        mystructobj3.Name = "Manju"
        mystructobj3.Age = 24
        
        
        
##### * Conflict between Global and Local variable
        We can redeclare same variable within the function, and which will dominate 
        within the function.

        The scoping closest to the use is what decides the value of variable. 
        So if you redeclare the variable inside your function, then for the 
        duration of that function you will have the value value assigned within the function.

        If you chose to use the same name for two things, 
        the consequence is that the inner name will always dominate.
        If you need both values then name the variables differently. same happens in other programming
        languages like Javascript and Python.
       
       


##### * Different ways of declaring/initializing variables in Go.

        Example:
        package main
        
        import "fmt"
        
        var e int8

        func main() {
          e = 200 //global
          var a int8
          a = 10
          b := 20
          var c = "I am a string 1"
          var d string = "I am a string 2"
          fmt.Println("Value of a", a) // 10
          fmt.Println("Value of b", b) // 20
          fmt.Println("Value of c", c) // I am a string 1
          fmt.Println("Value of d", d) // I am a string 2
          fmt.Println("Value global e", e) // 200
        }
        
        
##### * Different ways of initializing an array in Go.

        It does not support sort, better to use slice. Array will be having fixed length, but slices 
        are flexible(dynamic size)
	
	Example:
        package main

        import "fmt"

        func main() {

          //Array of Numbers
          var number_array = [4]int32{}
          fmt.Println("Array is", number_array) // Array is [0 0 0 0]
          fmt.Printf("Type of number_array is %T\n", number_array) //Type of number_array is [4]int32
          
          //(integer/float array default intializes all elements to zero)
          number_array[0] = 1 //skipping arr[1], so by default it will be initialized to 0
          number_array[2] = 2
          number_array[3] = 3
          fmt.Println("Array is", number_array) //Array is [1 0 2 3]

          //Array of Strings
          var strings_array = [4]string{}
          fmt.Println("Array is", strings_array) // Array is Array is [   ] (all elements are empty strings)
          
          //(string array default intializes all elements to empty string)
          strings_array[0] = "one" //skipping arr[1], so by default it will be initialized to ""(empty string)
          strings_array[2] = "two"
          strings_array[3] = "three"
          fmt.Println("Array is", strings_array) //Array is [one  two three]
          
          // Direct initialization
          var directly_initialized_array = [3]int{1, 2, 3}
	        fmt.Println("directly_initialized_array is", directly_initialized_array)
          
          
          // length of an array is always equals to the value given at the time of initialization irrespective of 
          // the number of elements it contains
          var my_new_array = [4]int{1, 2}  // initializing only two elemets but still lenght will be 4
          fmt.Println("length of my_new_array is", len(my_new_array)) //length of my_new_array is 4
  
        }
	
	
##### * Slices in Go.
	Slice will be having the same syntax like array but with empty brackets(no size will be mentioned)
	
	Example:
	package main

	import "fmt"

	func main() {

		var number_slice = []int32{1, 2, 3, 4} //there is another syntax to create slice, using make method
		// To add new element to above slice
		number_slice = append(number_slice, 5, 6, 7)
		fmt.Printf("Type of number_slice is %T\n", number_slice) //Type of number_slice is []int32
		fmt.Printf("number_slice is %v\n", number_slice)         // number_slice is [1 2 3 4 5 6]
		fmt.Printf("itmes are %v\n", number_slice[1:3])          // itmes are [2 3]

		//delete an item (deleting 2 from number_slice)
		//destructure operator (...)
		number_slice = append(number_slice[:1], number_slice[2:]...) //three dots are important
		fmt.Printf("number_slice is %v\n", number_slice) // number_slice is [1 3 4 5 6 7]
		
		//Creating slice using make method
		var my_another_slice = make([]int, 4)
		fmt.Println("my_another_slice", my_another_slice) //my_another_slice [0 0 0 0]
		my_another_slice[0] = 1
		my_another_slice[1] = 2
		my_another_slice[2] = 3
		my_another_slice[3] = 4
		fmt.Println("my_another_slice", my_another_slice) //my_another_slice [1 2 3 4]
		
		//my_another_slice[4] = 5                           // can't assign, throws index out of range error

		// Inorder to add more items than the size defined, we can use append, which will
		// internally reallocated the memory and adds the new elements
		my_another_slice = append(my_another_slice, 5, 6, 7)
		fmt.Println("my_another_slice", my_another_slice) //my_another_slice [1 2 3 4 5 6 7]

	}
	

##### * Pointers in Go.
        A pointer is  a variable which stores the reference of the other variable.
        Unlike other variables that hold values of a certain type, pointer holds the address of a variable. 

        Example:
        package main

        import "fmt"

        func main() {
          a := 1200
          b := 1500
          
          var ptr1 *int = &a  // we can explicitly mention the type of the pointer based on 
                              // what type of variable's address it's going to store. in our case
                              // it's storing the address of the variable which is of type int. so it will be *int
                              
          var ptr2 = &b  // type of ptr will be automatically inferred here
          
          fmt.Println("Reference of variable a is", ptr1) //0xc0000140d0
          fmt.Println("Value of variable a is", *ptr1)    //1200
          fmt.Println("Reference of variable b is", ptr2) //0xc0000140d8
          fmt.Println("Value of variable b is", *ptr2)    //1500
          
          *ptr2 = *ptr2 + 2 
          fmt.Println("Reference of variable b is", ptr2) //0xc0000140d8
          fmt.Println("New value of variable b is", b)    //1502
          fmt.Println("New value of variable b is", *ptr2)    //1502
        }
