package main

import "fmt"
import "os"
import "bufio"
//import "strings"
//import "strconv"

/**
 * Don't let the machines win. You are humanity's last hope...
 **/
 
type Vertex struct {
	X int
	Y int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // width: the number of cells on the X axis
    var width int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&width)
    
    // height: the number of cells on the Y axis
    var height int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&height)
    
    arr := make(map[Vertex]int)
    
    var i,j int
    var res = 0
    for i = 0; i < height; i++ {
        scanner.Scan()
        line := scanner.Text() // width characters, each either 0 or .
    
        for j =0 ; j< width; j ++ {
          
            if line[j] == '.' {
                res = 2
            } else {
                res = 1
            }
            arr[Vertex{j, i}] = res
        }
        
        fmt.Fprintln(os.Stderr, line)
    }

    var ki, kj int
    

    for i = 0; i < height; i++ {
        for j = 0; j < width; j++ {
            
            if arr[Vertex{j, i}] == 1 {
                
                fmt.Print( j, i, " ")

            	    for kj=j+1; kj<width; kj++ {
                        if arr[Vertex{kj, i}] == 1 {
                            break
                        }
                    } 
                    
                    if kj < width {
                        fmt.Print(kj, i, " ")
            	    } else {
                        fmt.Print(-1,-1, " ")
    
            	    }

                    for ki=i+1; ki<height; ki++ {
                        if arr[Vertex{j, ki}] == 1 {
                            break
                        }
                    } 
                    if ki < height {
                        fmt.Print(j, ki, " ")
            	    } else {
                        fmt.Print(-1,-1," ")
    
            	    }

                fmt.Println()
            
            }
        }
            
    }
        
}
