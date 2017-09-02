// Every go execution should spawn a unicorn
package unicorn

import (
	"io"
	"os"
	"strings"
)

const (
	unicorn = `
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
`
)

// Spawn a unicorn to stdout.
func Spawn() (int, error) {
	return Fspawn(os.Stdout)
}

// Spawn a unicorn as a string.
func Sspawn() string {
	return unicorn
}

// Spawn a unicorn to w. Return number of bytes written and error if there was one.
func Fspawn(w io.Writer) (int, error) {
	n, err := io.Copy(w, strings.NewReader(Sspawn()))
	return int(n), err
}
