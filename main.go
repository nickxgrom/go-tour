package main

func main() {
	//https://go.dev/tour/moretypes/23
	//fmt.Println(Exercises.WordCount("I am learning Go learning"))

	//https://go.dev/tour/moretypes/26'
	//fibonacci := Exercises.Fibonacci()
	//for i := 0; i < 50; i++ {
	//	fmt.Println(fibonacci())
	//}

	//https://go.dev/tour/methods/18
	//hosts := map[string]Exercises.IPAddr{
	//	"loopback":  {123, 0, 0, 1},
	//	"googleDNS": {8, 8, 8, 8},
	//}
	//
	//for name, ip := range hosts {
	//	fmt.Printf("%v: %v\n", name, ip)
	//}

	//https://go.dev/tour/methods/20
	//var val, err = Exercises.Sqrt(-2)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(val)
	//}

	//https://go.dev/tour/methods/23
	//s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//r := Exercises.Rot13Reader{R: s}
	//
	//io.Copy(os.Stdout, &r)

	//str := "Lbh penpxrq gur pbqr!"
	//byteArr := []byte(str)
	//
	//r := rot13{}
	//n, _ := r.Read(byteArr)
	//
	//fmt.Println(string(byteArr[:n]))
}

//type rot13 struct{}
//
//func (r rot13) Read(p []byte) (int, error) {
//	for i := range p {
//		if (p[i] >= 65 && p[i] <= 77) || (p[i] > 97 && p[i] <= 109) {
//			p[i] += 13
//		} else if (p[i] > 77 && p[i] <= 90) || (p[i] > 109 && p[i] <= 122) {
//			p[i] -= 13
//		}
//	}
//
//	return len(p), nil
//}
