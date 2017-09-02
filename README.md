go-unicorn
----------
Because every Go execution should spawn a unicorn

## Install ##
```
$ go get -u https://github.com/orirawlings/go-unicorn
```

## Demo ##
```
package main

import (
	"github.com/orirawlings/go-unicorn"
)

func main() {
	unicorn.Spawn()
}
```

```
$ go run main.go

                                        ////))))))\          
                                       /////))))))))))=======
                                    ///////           ))     
                                ///////      \_    *) )      
                ________________//////      /  \    /        
           ////))                  \>       |   |  |         
     /////////                      \        \   |_o         
 ////////////                        )    _.  |              
  ///////   \       \                       \ |              
 ///         \       \                 \     \/\______.      
             /\       )________         \_    \____.__ \     
           _/  \     /         '--------- \________ \ \ \    
          /  _/ /   /                             / /  \ \   
         / /    \  <                             / /    > )  
        ( <       \ \                          _//      \/   
         \ \       \ \                        |_/            
          \\_       \\_                                      
           |_\       |_\ 
```

## Documentation ##
See https://godoc.org/github.com/orirawlings/go-unicorn
