package main

import "fmt"
//собственно сама структура(имя баланс и какой банк)
type BankSystem struct {
	Name string
	Balance float64
	Bank string
}
//функция по поиску человека в структуре 
func FindPerson(banking []BankSystem, tempname string) int {
	idxPerson := -1
	for i := range banking {
		if banking[i].Name == tempname {
			idxPerson = i
			break
		}
	} 
	return idxPerson
}
//функция по просмотрю баланса
func ShowBalance(banking []BankSystem, tempname string) {
	idxPerson := FindPerson(banking, tempname)
	if idxPerson == -1 {
		fmt.Println("There is not such user(Такого пользователя не существует)")
	} else {
		fmt.Println("Your balance(Ваш баланс):",banking[idxPerson].Balance)
	} 
}
//функция по переводу денежных средств
func Transfer(banking []BankSystem, tempname1, tempname2 string, commission float64) {
	idxFrom := FindPerson(banking, tempname1)
	idxTo := FindPerson(banking, tempname2)
	if idxFrom == -1 {
		fmt.Println("There is not such sender(Такого отправителя не существует)")
	} else if idxTo == -1 {
		fmt.Println("There is not such recipient(Такого получателя не существует)")
	} else {
		fmt.Println("How much money do you want to transfer(Сколько денег хотите перевести)?")
		var sum float64
		fmt.Scanln(&sum)
		if sum > banking[idxFrom].Balance {
			fmt.Println("There are not enogh funds on the balance(недостаточно средств на счете)")
		} else if sum < 0 {
			fmt.Println("The amount is not correct(Сумма не может быть меньше нуля)")
		} else if sum <= banking[idxFrom].Balance {
			if banking[idxFrom].Bank != banking[idxTo].Bank {
				banking[idxFrom].Balance -= sum
				banking[idxTo].Balance += sum - (sum * commission)
				fmt.Println("The commission:",sum * commission)
				fmt.Println("Funds have been successfully transferred with a commission of 1%(Средства успешно переведены с комиссией в 1%)")
				fmt.Println(banking[idxFrom].Name,"on your balance(на балансе):",banking[idxFrom].Balance)
				fmt.Println(banking[idxTo].Name,"on your balance(на балансе):",banking[idxTo].Balance)
			} else {
				banking[idxFrom].Balance -= sum
				banking[idxTo].Balance += sum
				fmt.Println("The commission: 0")
				fmt.Println("Funds have been successfully transfer(Средства успешно переведены)")
				fmt.Println(banking[idxFrom].Name,"on your balance(на балансе):",banking[idxFrom].Balance)
				fmt.Println(banking[idxTo].Name,"on your balance(на балансе):",banking[idxTo].Balance)
			}
		}
	}
}
// функция по пополнению баланса
func Deposit(banking []BankSystem, tempname string) {
	idxPerson := FindPerson(banking, tempname)
	if idxPerson == -1 {
		fmt.Println("There is not such user(Такого пользователя не существует)")
	} else {
		var sum float64
		fmt.Println("How much do you want to top up(Сколько хотите пополнить)?")
		fmt.Scanln(&sum)
		if sum <= 0 {
			fmt.Println("The amount is not correct(Сумма не может быть меньше нуля)")
		} else {
			banking[idxPerson].Balance += sum
			fmt.Println("The balance has been successfully replenished(Баланс успешно пополнен)")
			fmt.Println("On your balance(На вашем балансе):",banking[idxPerson].Balance)
		}
	}
}
// функция по снятию денежных средств
func Withdraw(banking []BankSystem, tempname string) {
	idxPerson := FindPerson(banking, tempname)
	if idxPerson == -1 {
		fmt.Println("There is not such user(Такого пользователя не существует)")
	} else {
		var sum float64
		fmt.Println("How much do you want to withdraw(Сколько хотите снять)?")
		fmt.Scanln(&sum)
		if sum > banking[idxPerson].Balance {
			fmt.Println("There are not enogh funds on the balance(недостаточно средств на счете)")
		} else if sum < 0 {
			fmt.Println("The amount is not correct(Сумма не может быть меньше нуля)")
		} else if sum <= banking[idxPerson].Balance {
			banking[idxPerson].Balance -= sum
			fmt.Println("The amount is successfully withdraw(Сумма успешно снята)")
			fmt.Println("On your balance(На вашем балансе):",banking[idxPerson].Balance)
		}
	}
}
//функция по просмотрю всей информации о клиентах
func (b BankSystem) FullInfo() {
	fmt.Println("Your name:",b.Name)
	fmt.Println("Your balance:",b.Balance)
	fmt.Println("Your bank:",b.Bank)
}



func main() {
// задаем нашу структуру
	banking := []BankSystem{}
//выбираем количество клиентов
	var number int
	fmt.Println("How many accounts do you want to create?")
	fmt.Scanln(&number)
//указываем комисию за перевод в разные банки
	commission := 0.01
//создаем наших клиентов
	for i := 0; i < number; i++ {
		var name string
		var balance float64
		var bank string

		fmt.Println("What is your name?")
		fmt.Scanln(&name)
		fmt.Println("How much money do you want to deposit?")
		fmt.Scanln(&balance)
		fmt.Println("What kind of bank do you have?")
		fmt.Scanln(&bank)

		newBank := BankSystem {
			Name: name,
			Balance: balance,
			Bank: bank,
		}

		banking = append(banking, newBank)
	}

	flag := false
//начинается веселье и интерактив
	for !flag {

		var value int
//выборы для интерактива
		fmt.Println("1.Balance(Баланс)")
		fmt.Println("2.Transfer(Перевод)")
		fmt.Println("3.Replenishment(Пополнение)")
		fmt.Println("4.Withdrawal(Снятие)")
		fmt.Println("5.Full info(Полная информация)")
		fmt.Println("6.Exit(Выход)")

		fmt.Scanln(&value)
//переходы в функции в зависимости от выбора
		switch value {
		 	case 1:
				var tempname string
				fmt.Println("Whose balance do you want to see(Кого хотите посмотреть)?")
				fmt.Scanln(&tempname)

				ShowBalance(banking, tempname)

		 	case 2:
				var tempname1 string
				fmt.Println("Who will transfer the funds(Кто будет переводить деньги)?")
				fmt.Scanln(&tempname1)
				var tempname2 string
				fmt.Println("Who will receive the funds(Кто будет получать деньги)?")
				fmt.Scanln(&tempname2)

				Transfer(banking, tempname1, tempname2, commission)

		 	case 3:
				var tempname string 
				fmt.Println("Who will replenish the balance(Кто будет пополнять баланс)?")
				fmt.Scanln(&tempname)

				Deposit(banking, tempname)
				
		 	case 4:
				var tempname string
				fmt.Println("Who will withdraw from the balance(Кто будет снимать деньги)?")
				fmt.Scanln(&tempname)
				
				Withdraw(banking, tempname)

		 	case 5:
			 	for _, b := range banking {
				 	b.FullInfo()
			 	}
		 	case 6:
			 	fmt.Println("All the best, goodbye(Всего доброго, до свидания)")
				flag = true	
			default:
				fmt.Println("There is no such option(Такого варианта нет)")
		}
	}
}