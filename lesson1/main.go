package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// var files []*os.File

// ImplicitPanicError struct
type ImplicitPanicError struct {
	Time time.Time
	Err  error
}

func (err *ImplicitPanicError) Error() string {
	return fmt.Sprintf("%s: %v", err.Time.Format(time.UnixDate), err.Err)
}

func main() {
	err := implicitPanic()
	if err != nil {
		fmt.Println(err)
	}

	// err := createFiles(1000000)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func implicitPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%w: %v", newErrorImplicitPanic(), v)
		}
	}()

	var test interface{}

	s := test.(string)
	fmt.Println(s)

	return nil
}

func createFiles(filesNumber int) error {
	if filesNumber < 1 {
		return errors.New("filesNumber must be positive number")
	}

	// Сколько не пытался, не смог добиться того, чтобы возникла паника
	// пробовал складывать указатели в слайс, но тоже всё отработало
	for i := 1; i <= filesNumber; i++ {
		filename := "./files/file_" + strconv.Itoa(i)
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		// files = append(files, file)
		fmt.Println(filename)

		// С отложенным закрытием файла, тоже не очень понял, чем это поможет внутри цикла...
		// решил сделать простое закрытие
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("close file error: %s\n", closeErr)
		}
	}

	return nil
}

func newErrorImplicitPanic() *ImplicitPanicError {
	return &ImplicitPanicError{
		Time: time.Now(),
		Err:  errors.New("implicit error"),
	}
}
