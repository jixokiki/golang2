package main
import "fmt"
func main(){

	NumberPrime(1000000007)
	fmt.Println("\n")	
	NumberPrime(13)
	fmt.Println("\n")
	NumberPrime(17)
	fmt.Println("\n")
	NumberPrime(20)
	fmt.Println("\n")
	NumberPrime(35)
	fmt.Println("\n")
	primeNumber()
}

func primeNumber() {	
	number1:= 1000000007; 
	d:=2;
	if(d==2)&&(number1!=1){
		fmt.Println(number1, "Merupakan bilangan prima");
	}else{
		fmt.Println(number1, "Bukan merupakan bilangan prima");
	}

	number2:= 13; 
	c:=2;
	if(c==2)&&(number2!=1){
		fmt.Println(number2, "Merupakan bilangan prima");
	}else{
		fmt.Println(number2, "Bukan merupakan bilangan prima");
	}

	number3:= 17; 
	a:=2;
	if(a==2)&&(number3!=1){
		fmt.Println(number3, "Merupakan bilangan prima");
	}else{
		fmt.Println(number3, "Bukan merupakan bilangan prima");
	}

	number4:= 20; 
	b:=0;
	if(b==2)&&(number4!=1){
		fmt.Println(number4, "Merupakan bilangan prima");
	}else{
		fmt.Println(number4, "Bukan merupakan bilangan prima");
	}
	
	number5:= 35; 
	i:=0;
	if(i==2)&&(number5!=1){
		fmt.Println(number5, "Merupakan bilangan prima");
	}else{
		fmt.Println(number5, "Bukan merupakan bilangan prima");
	}

}

func NumberPrime(numericNumber int){
	i:=2;
	a:= 0
	if(i==2)&&(a==2){
		fmt.Println(numericNumber, "Merupakan bilangan prima");
	}else if(a==2)&&(i==2){
		fmt.Println(numericNumber, "Bukan merupakan bilangan prima");
	}else{
		fmt.Println(numericNumber, "Bukan merupakan bilangan prima");
	}	
}
