/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import ("fmt"; "sync")

//тип данных в стеке
type stk_type interface{}

type stack struct{
	lock sync.Mutex

	buff []stk_type //данные стека (массив)
}

func sinit() *stack{
	return &stack { sync.Mutex{}, make([]stk_type, 0)}	
}


func (stk *stack) push(value stk_type){
	stk.lock.Lock()
	defer stk.lock.Unlock()

	stk.buff = append(stk.buff, value)
}

func (stk *stack) pop() stk_type {
	stk.lock.Lock()
	defer stk.lock.Unlock()
	
	
	length := len(stk.buff)

	if length == 0 { 
	    return nil
	}

	lastValue := stk.buff[length - 1]
	stk.buff = stk.buff[:length - 1]
	
	return lastValue
}

func (stk *stack) peek() stk_type {
	stk.lock.Lock()
	defer stk.lock.Unlock()
	
	length := len(stk.buff)

	if length == 0 { 
	    return nil
	}

	return stk.buff[length - 1]
}

func (stk *stack) empty() bool{
	stk.lock.Lock()
	defer stk.lock.Unlock()

	return (len(stk.buff) == 0)
}

func (stk *stack) get() []stk_type{
    stk.lock.Lock()
	defer stk.lock.Unlock()

	var cpy = make([]stk_type, len(stk.buff))
	copy(cpy, stk.buff)

	return cpy
}

func main() {
    s := sinit() 
    
    s.push(34)
    s.push(5)
    s.push(45)
    
    fmt.Println(s.get())
    fmt.Println()
    
    for !s.empty() {
        fmt.Println(s.pop())
    }
    fmt.Println()
    
    s.push(6)
    fmt.Println(s.pop())
    fmt.Println(s.peek())
    
    s.push(6784)
    fmt.Println(s.peek())
    
    s.push(5)
    s.push(4)
    s.push(3)
    s.push(2)
    s.push(9999)
    
    fmt.Println(s.get())
}
