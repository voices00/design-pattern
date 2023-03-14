package main

//	责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。
//收到请求后， 每个处理者均可对请求进行处理， 或将其传递给链上的下个处理者。

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
