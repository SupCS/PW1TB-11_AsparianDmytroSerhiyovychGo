package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Функція для округлення чисел до N знаків після коми
func roundFloat(value float64, precision int) string {
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, value)
}

// Функція для обробки запиту на калькулятор 1
func Task1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		H, _ := strconv.ParseFloat(r.FormValue("H"), 64)
		C, _ := strconv.ParseFloat(r.FormValue("C"), 64)
		S, _ := strconv.ParseFloat(r.FormValue("S"), 64)
		N, _ := strconv.ParseFloat(r.FormValue("N"), 64)
		O, _ := strconv.ParseFloat(r.FormValue("O"), 64)
		W, _ := strconv.ParseFloat(r.FormValue("W"), 64)
		A, _ := strconv.ParseFloat(r.FormValue("A"), 64)

		// Розрахунок коефіцієнтів переходу
		KRS := 100 / (100 - W)
		KRG := 100 / (100 - W - A)

		// Розрахунок нижчої теплоти згоряння для робочої маси
		QH := (339*C + 1030*H - 108.8*(O-S) - 25*W) / 1000

		// Розрахунок нижчої теплоти згоряння для сухої та горючої маси
		QH_dry := (QH + 0.025*W) * KRS
		QH_comb := (QH + 0.025*W) * KRG

		// Перерахунок складу сухої маси
		H_dry := H * KRS
		C_dry := C * KRS
		S_dry := S * KRS
		N_dry := N * KRS
		O_dry := O * KRS
		A_dry := A * KRS

		// Перерахунок складу горючої маси
		H_comb := H * KRG
		C_comb := C * KRG
		S_comb := S * KRG
		N_comb := N * KRG
		O_comb := O * KRG

		// Форматуємо числа до N знаків після коми
		precision := 2
		precision2 := 4
		tmpl, _ := template.ParseFiles("templates/task1.html")
		tmpl.Execute(w, map[string]string{
			"KRS":     roundFloat(KRS, precision),
			"KRG":     roundFloat(KRG, precision),
			"QH":      roundFloat(QH, precision2),
			"QH_dry":  roundFloat(QH_dry, precision2),
			"QH_comb": roundFloat(QH_comb, precision2),
			"H_dry":   roundFloat(H_dry, precision),
			"C_dry":   roundFloat(C_dry, precision),
			"S_dry":   roundFloat(S_dry, precision),
			"N_dry":   roundFloat(N_dry, precision),
			"O_dry":   roundFloat(O_dry, precision),
			"A_dry":   roundFloat(A_dry, precision),
			"H_comb":  roundFloat(H_comb, precision),
			"C_comb":  roundFloat(C_comb, precision),
			"S_comb":  roundFloat(S_comb, precision),
			"N_comb":  roundFloat(N_comb, precision),
			"O_comb":  roundFloat(O_comb, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/task1.html")
	tmpl.Execute(w, nil)
}
