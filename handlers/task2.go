package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

// Функція для обробки запиту на калькулятор 2
func Task2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		C, _ := strconv.ParseFloat(r.FormValue("C"), 64)
		H, _ := strconv.ParseFloat(r.FormValue("H"), 64)
		S, _ := strconv.ParseFloat(r.FormValue("S"), 64)
		O, _ := strconv.ParseFloat(r.FormValue("O"), 64)
		W, _ := strconv.ParseFloat(r.FormValue("W"), 64)
		A, _ := strconv.ParseFloat(r.FormValue("A"), 64)
		V, _ := strconv.ParseFloat(r.FormValue("V"), 64)
		Q_comb_value, _ := strconv.ParseFloat(r.FormValue("Q_comb"), 64)

		// Розрахунок коефіцієнту переходу до робочої маси
		KRS := (100 - W - A) / 100
		C_work := C * KRS
		H_work := H * KRS
		S_work := S * KRS
		O_work := O * KRS
		V_work := V * (100 - W) / 100
		Q_r := Q_comb_value*(100-W-A)/100 - 0.025*W

		// Форматуємо числа до 2 знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/task2.html")
		tmpl.Execute(w, map[string]string{
			"KRS":    roundFloat(KRS, precision),
			"C_work": roundFloat(C_work, precision),
			"H_work": roundFloat(H_work, precision),
			"S_work": roundFloat(S_work, precision),
			"O_work": roundFloat(O_work, precision),
			"V_work": roundFloat(V_work, 1),
			"A":      roundFloat(A, precision),
			"Q_r":    roundFloat(Q_r, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/task2.html")
	tmpl.Execute(w, nil)
}
