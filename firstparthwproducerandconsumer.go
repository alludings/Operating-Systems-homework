package main //package main is start with go language

import "fmt" //use to print on the codespace

func main() {
prodconvalue:= make(chan int) //channel for integer
	goempty:= make(chan struct{}) //channel for structure does nothing but indication 

		go func(){ //go function language
			for i:= 1;i<= 5; i++ { //for loop numbers 1 2 3 4 and 5
				fmt.Printf("Producer:%d\n", i)// prints the Producer number
				prodconvalue <- i      //makes the number to consumer
					<-goempty          //wait inducation and continues
}
close(prodconvalue) //after it says numbers it ends
}()

	for number:= range prodconvalue{ //starts stating numbers waiting clossed
		fmt.Printf("Consumer:%d\n", number) //consumer number prints
			goempty<- struct{}{}  //indicates to producer to continue
}
}