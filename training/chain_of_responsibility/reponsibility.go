package chain_of_responsibility

type Department interface {
	execute(*Patient)
	setNext(Department)
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Reception struct {
}
