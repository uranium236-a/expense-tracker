package models

type Expense struct {
	ID      int
	Subject string
	Money   int
}

type ExpenseStore struct {
	Expenses map[int]*Expense
	NextID   int
}

func NewExpenseStore() *ExpenseStore {
	return &ExpenseStore{
		Expenses: map[int]*Expense{},
		NextID:   0,
	}
}

func (e *ExpenseStore) AddExpense(subject string, money int) {
	id := e.NextID

	e.Expenses[id] = &Expense{
		ID:      id,
		Subject: subject,
		Money:   money,
	}

	e.NextID++
}

func (e *ExpenseStore) DeleteExpense(ID int) {
	delete(e.Expenses, ID)
}

func (e *ExpenseStore) MaxExpense() Expense {
	maxExpense := Expense{Money: 0}

	for _, expense := range e.Expenses {
		if expense.Money > maxExpense.Money {
			maxExpense = *expense
		}
	}

	return maxExpense
}
func (e *ExpenseStore) MinExpense() Expense {
	minExpense := Expense{Money: 999999}

	for _, expense := range e.Expenses {
		if expense.Money < minExpense.Money {
			minExpense = *expense
		}
	}

	return minExpense
}
