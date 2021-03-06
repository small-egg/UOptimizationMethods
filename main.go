package main

import "fmt"
import "./methods"

var (
	x0 = 1.0
	y0 = 2.0

	eps = 0.001

	lambda = 1.0

	maxIt = 10
)

//func init() {
//	fmt.Println("Введите x0: ")
//	fmt.Scan(&x0)
//
//	fmt.Println("Введите y0: ")
//	fmt.Scan(&y0)
//
//	fmt.Println("Введите точность: ")
//	fmt.Scan(&eps)
//
//	fmt.Println("Введите максимальное количество итераций для некоторых методов")
//	fmt.Scan(&maxIt)
//}

func echo(method string, point methods.Point) {
	fmt.Printf("Метод %s:\nМинимум функции находится в точке (%f; %f)\nf(x,y) = %f\n\n",
		method, point.X, point.Y, methods.F(point))
}

func main() {
	start := methods.Point{x0, y0}

	p := methods.PatternSearch(start, eps, lambda)
	echo("Хука-Дживса", p)

	p = methods.PaulsMethod(start, eps)
	echo("Сопряженных напрвлений", p)

	p = methods.GradientDescent(maxIt, start, eps, lambda)
	echo("Градиентного спуска", p)

	p = methods.ConjugateGradients(3, start, eps)
	echo("Сопряженных градиентов", p)

	p = methods.Newton(maxIt, start, eps, lambda)
	echo("Метод Ньютона", p)

	p = methods.Dfp(start, eps)
	echo("Давидона-Флетчера-Пауэлла", p)
}
