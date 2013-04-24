/*
You are given the following information, but you may prefer to do some research
for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century
unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1
Jan 1901 to 31 Dec 2000)?
*/

package main

import "fmt"

const (
    monday = iota
    tuesday
    wednesday
    thursday
    friday
    saturday
    sunday
)

var daynames = []string{ "Monday", "Tuesday", "Wednesday", "Thursday",
    "Friday", "Saturday", "Sunday" }

const (
    january = iota
    february
    march
    april
    may
    june
    july
    august
    september
    octobar
    november
    december
)

var monthlengths = []int{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 }
var monthnames = []string{ "January", "February", "March", "April", "May",
    "June", "July", "August", "Septeber", "October", "November", "December" }

func IsLeapYear(year int) (leap bool) {
    leap = false
    if year % 4 == 0  {
        if year % 100 != 0 {
            leap = true
        } else if year % 400 == 0 {
            leap = true
        }
    }
    return
}

func MonthLength(month, year int) (length int) {
    length = monthlengths[month]
    if month == february && IsLeapYear(year) {
        length+=1
    }
    return
}

func YearLength(year int) (length int) {
    length = 365
    if IsLeapYear(year) {
        length += 1
    }
    return
}

func main() {
    day := monday // date 1/1/1900
    day += YearLength(1900) % 7
    first_month_sunday := 0
    for year := 1901; year < 2001; year++ {
        for month := january; month <= december; month++ {
            if day == sunday  {
                first_month_sunday += 1
            }
            skip := (MonthLength(month,year) % 7)
            day = (day + skip)  % 7
        }
    }
    fmt.Println(first_month_sunday)
}
