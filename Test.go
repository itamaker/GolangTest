package main

import (
	//"fmt"
	//"os"
	//"strings"
	//"bufio"
)

func main() {

	//println("I am ", os.Args[0])

	//fmt.Print(123)

	//fmt.Println(strings.Join(os.Args[1:],""))
	//
	//var s,sep string
	//
	//for i := 1; i < len(os.Args); i++ {
	//
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)

	//s, sep := "", ""
	//
	//for _,arg := range os.Args[1:]{
	//
	//	s += sep + arg
	//
	//	sep = " "
	//}
	//
	//fmt.Println(s)


	//fmt.Println(strings.Join(os.Args[1:]," "))

	//fmt.Println(os.Args[1:])

	//counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//
	//for input.Scan() {
	//	counts[input.Text()]++
	//}
	//
	//for line,n := range counts {
	//	if n > 1 {
	//		fmt.Println("%d,\t%s\n",n,line)
	//	}
	//
	//}


	//counts := make(map[string]int)
	//files := os.Args[1:]
	//if len(files) == 0 {
	//	countLines(os.Stdin,counts)
	//}else {
	//	for _,arg := range  files{
	//		f,err := os.Open(arg)
	//		if err != nil {
	//			fmt.Fprintln(os.Stdin,"dup2:	%v\n",err)
	//			continue
	//		}
	//
	//		countLines(f,counts)
	//		f.Close()
	//
	//
	//	}
	//}
	//
	//for line,n := range  counts{
	//	if n > 1 {
	//		fmt.Println("%d\t%s\n",n,line)
	//	}
	//}



}


//func countLines(f *os.File,counts map[string]int)
//
//	input := bufio.NewScanner(f)
//
//	for input.Scan() {
//
//		counts[input.Text()]++
//
//	}
//
//}