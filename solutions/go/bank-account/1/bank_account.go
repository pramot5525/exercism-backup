package account
import "sync"

// Define the Account type here.
type Account struct {
	mu      sync.Mutex
    balance int64
    isClosed bool
}


func Open(amount int64) *Account {
    if amount < 0 {
        return nil
    }
    
    return &Account{
        balance: amount,
    }
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
	
    if a == nil || (a != nil && a.isClosed) {
        return 0, false
    }


    return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()

    if a == nil || (a != nil && a.isClosed) {
        return 0, false
    }
    newBalance := a.balance+ amount
    
    if newBalance < 0 {
        return 0, false
    }
    
	a.balance = newBalance
    
    return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
    
    if a == nil || (a != nil && a.isClosed) {
        return 0, false
    }
    a.isClosed = true

    return a.balance, true
}
