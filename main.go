package main

import "github.com/carlonicolo/expert-session-finance/adapter/http"

func main() {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello Welcome in my page")
		})*/

	/*
		http.HandleFunc("/transactions", transaction.GetTransactions)
		http.HandleFunc("/transactions/create", transaction.CreateTransaction)

		http.ListenAndServe(":8888", nil)

	*/

	http.Init()
}

/*
type Transaction struct {
	Title    string    `json:"title"`
	Amount   float32   `json:"amount"`
	Type     int       `json:"type"`
	CreateAt time.Time `json:"create_at"`
}

type Transactions []Transaction
*/

/*
func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2022-01-05T11:40:25")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:    "Salario",
			Amount:   1200.0,
			Type:     0,
			CreateAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
*/
