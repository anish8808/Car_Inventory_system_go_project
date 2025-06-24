package main

import "fmt"

/*
Beacuse of bad code writting these things can happen
	beacuse of rigidity(tightly coupled not able to make small changes)
	 or fragidity(kuch bhi change se code fat jaye)
	 or immobilty (reusable nahi hona) ,
	 or viscocity(hacking instaed of solution)


	if we will use design principle to we can restrict the above things

	1️⃣ Single Responsibility Principle (SRP)
	👉 	Definition: A class/type should have only one reason to change.

		✅ In the code:

		PaymentProcessor only coordinates — it doesn't handle validation, payment, or logging directly.

		MinAmountValidator handles validation logic only.

		StripeGateway handles payment logic only.

		ConsoleLogger handles logging only.

		📌 Why it follows SRP:
		If tomorrow you change validation rules → you modify MinAmountValidator, not PaymentProcessor.
		If you want to log to a file instead of console → change ConsoleLogger, no impact on processor.


	2️⃣ Open/Closed Principle (OCP)
		👉 Definition: Code should be open for extension, closed for modification.

		✅ In the code:

		Want a new payment gateway (e.g., PayPal)? Create PayPalGateway struct — no need to change PaymentProcessor.

		Want a new validator (e.g., MaxAmountValidator)? Add it without changing PaymentProcessor.

		📌 Why it follows OCP:
		We can extend behavior (add new gateways/validators/loggers) without touching or modifying the core processor logic.

	3️⃣ Liskov Substitution Principle (LSP)
		👉 Definition: Subtypes should be substitutable for base types without breaking behavior.

		✅ In the code:

		StripeGateway, PayPalGateway, CashGateway (if added) can all be passed where PaymentGateway is expected.

		MinAmountValidator, MaxAmountValidator can replace each other where PaymentValidator is used.

		📌 Why it follows LSP:
		PaymentProcessor doesn’t care what gateway/validator it gets as long as it satisfies the interface — no special case logic for each subtype.


	4️⃣ Interface Segregation Principle (ISP)
		👉 Definition: No client should be forced to depend on methods it doesn’t use.

		✅ In the code:

		PaymentValidator interface: only defines Validate.

		PaymentGateway interface: only defines Pay.

		Logger interface: only defines Log.

		📌 Why it follows ISP:
		We don’t force a gateway to implement Validate or Log.
		Each interface is small, focused — Go-style minimal interfaces.
	5️⃣ Dependency Inversion Principle (DIP)

👉 Definition: High-level modules should depend on abstractions, not concrete implementations.

		✅ In the code:

		PaymentProcessor depends on PaymentValidator, PaymentGateway, Logger interfaces.

		Main() provides concrete types (e.g., MinAmountValidator, StripeGateway, ConsoleLogger).

		📌 Why it follows DIP:
		PaymentProcessor has no hard dependency on Stripe, console, or any specific validator. You can inject mocks for testing, or replace components at runtime.



*/

// Interfaece
type PaymentValidator interface {
	Validate(float64) error
}

type PaymentGateway interface {
	Pay(float64) error
}

type logger interface {
	Log(string)
}

// Processor
type PaymetProcessor struct {
	validator PaymentValidator
	gateway   PaymentGateway
	logger    logger
}

func (p *PaymetProcessor) Process(amount float64) error {
	if err := p.validator.Validate(amount); err != nil {
		return err
	}

	if err := p.gateway.Pay(amount); err != nil {
		return err
	}

	p.logger.Log("Payment processed successfully")
	return nil
}

//implementation

type MinAmountValidator struct {
	min float64
}

func (m MinAmountValidator) Validate(amount float64) error {
	if amount < m.min {
		return fmt.Errorf("too low")
	} else {
		return nil
	}
}

type stripGateway struct {
}

func (s stripGateway) Pay(amount float64) error {
	fmt.Println("stripe amount Paid ", amount)

	return nil
}

type Consolelogger struct {
}

func (l Consolelogger) Log(msg string) {
	fmt.Println("Log message ", msg)
}

func main() {

	p := PaymetProcessor{
		validator: MinAmountValidator{min: 100},
		gateway:   stripGateway{},
		logger:    Consolelogger{},
	}

	_ = p.Process(1000)

}
