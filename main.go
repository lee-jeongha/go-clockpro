package main

import (
	"fmt"
	"flag"
	"math"
	"bufio"
	"bytes"
	"os"
)

func ClockProSimul(file string, cache_size int) {
	f, err := os.Open(file)

	if err != nil {
		//t.Fatal(err)
		fmt.Printf("error has type %T\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)

	//cache := clockpro.New(200)
	cache := New(cache_size)
	hit_cnt := 0
	miss_cnt := 0

	for scanner.Scan() {
		line := scanner.Bytes()
		fields := bytes.Split(line, []byte{','})
		key := string(fields[4])

		v := cache.Get(key)
		if v == nil {
			miss_cnt += 1
			cache.Set(key, key)
		} else {
			hit_cnt += 1
			if v.(string) != key {
				//t.Errorf("cache returned bad data: got %+v , want %+v\n", v, key)
				fmt.Printf("cache returned bad data: got %+v , want %+v\n", v, key)
				os.Exit(1)
			}
		}
	}
	fmt.Println(cache_size, hit_cnt, miss_cnt, hit_cnt+miss_cnt)
}


func main() {
	file := flag.String("file", "default.txt", "Input file")
	cache_size := flag.Int("cachesize", 100, "Cache Size")
	cs := 0.0

	flag.Parse()

	// Pointer variable --> Need to add * in front of them to dereference
	//fmt.Println(*file, *cache_size)
	fmt.Println("cache_size hit miss total")
	for i := 1; i <= 20; i += 1 {
		cs = math.Round(float64(*cache_size) * float64(i) * 0.05)
		ClockProSimul(*file, int(cs))
	}
}