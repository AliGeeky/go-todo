package main

import (
	"fmt"
	"strconv" // برای تبدیل رشته به عدد
	"bufio"   // برای خواندن ورودی از کاربر به صورت خط به خط
	"os"      // برای کار با ورودی/خروجی سیستم (stdin)
	"strings" // برای عملیات روی رشته‌ها، مانند حذف فضای خالی
)

// Todo ساختاری برای نگهداری اطلاعات یک وظیفه (todo item)
type Todo struct {
	ID     int    // شناسه یکتای وظیفه
	Title  string // عنوان یا توضیحات وظیفه
	Done   bool   // وضعیت وظیفه: true اگر انجام شده باشد، false اگر نشده باشد
}

// یک اسلایس سراسری برای نگهداری تمام وظایف
// اسلایس یک نوع داده پویا در Go است که می‌تواند تعداد متغیری از عناصر را نگهداری کند.
var todos []Todo
// nextID برای تولید شناسه‌های یکتا برای وظایف جدید استفاده می‌شود.
var nextID int = 1

func main() {
	// حلقه اصلی برنامه برای نمایش منو و دریافت ورودی کاربر
	for {
		fmt.Println("\n--- Go Todo CLI ---")
		fmt.Println("1. افزودن وظیفه")
		fmt.Println("2. لیست وظایف")
		fmt.Println("3. علامت‌گذاری به عنوان انجام شده")
		fmt.Println("4. حذف وظیفه")
		fmt.Println("5. خروج")
		fmt.Print("انتخاب شما: ")

		reader := bufio.NewReader(os.Stdin) // ایجاد خواننده برای دریافت ورودی از کنسول
		input, _ := reader.ReadString('\n') // خواندن ورودی تا کاراکتر خط جدید
		input = strings.TrimSpace(input)    // حذف فاصله‌های اضافی (مانند خط جدید) از ابتدا و انتهای ورودی

		switch input { // بررسی انتخاب کاربر
		case "1":
			fmt.Print("عنوان وظیفه: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			addTodo(title) // فراخوانی تابع افزودن وظیفه
		case "2":
			listTodos() // فراخوانی تابع لیست وظایف
		case "3":
			fmt.Print("شناسه وظیفه برای علامت‌گذاری به عنوان انجام شده: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr)) // تبدیل رشته به عدد
			if err != nil { // بررسی خطا در تبدیل
				fmt.Println("خطا: شناسه معتبر نیست.")
				continue // ادامه به حلقه بعدی
			}
			markTodoDone(id) // فراخوانی تابع علامت‌گذاری
		case "4":
			fmt.Print("شناسه وظیفه برای حذف: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("خطا: شناسه معتبر نیست.")
				continue
			}
			deleteTodo(id) // فراخوانی تابع حذف
		case "5":
			fmt.Println("خروج از برنامه.")
			return // خروج از تابع main و پایان برنامه
		default:
			fmt.Println("انتخاب نامعتبر است. لطفا دوباره امتحان کنید.")
		}
	}
}