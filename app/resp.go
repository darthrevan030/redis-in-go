package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	STRING = '+'
	ERROR = '-'
	INTEGER = ':'
	BULK = '$'
	ARRAY = '*'
)


/*
typ is used to determine the data type carried by the value.
str holds the value of the string received from the simple strings.
num holds the value of the integer received from the integers.
bulk is used to store the string received from the bulk strings.
array holds all the values received from the arrays.
*/
type Value struct {
	typ string
	str string
	num int
	bulk string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)
		if len(line) >= 2 && (line[len(line)-2] == '\r' && line[len(line)-1] == '\n') {
			break
		}
	}

	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, n, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 40)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.typ = "array"

	// read the number of elements in the array
	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	// for each element in the array, read the value and append it to the array
	v.array = make([]Value, length)
	for i := 0; i < length; i++ {
		value, err := r.Read()
		if err != nil {
			return v, err
		}
		v.array[i] = value
	}
	return v, nil
}

func (r *Resp) readBulk() (Value, error) {
	v := Value{}
	v.typ = "bulk"

	length, _ , err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, length)
	r.reader.Read(bulk)
	v.bulk = string(bulk)
	r.readLine() // to shift the pointer to the end so it can read the next bulk string correctly
	return v, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _type {
		case ARRAY:
			return r.readArray()
		case BULK:
			return r.readBulk()
		default:
			fmt.Printf("invalid type: %v\n", string(_type))
			return Value{}, err
		}
}
