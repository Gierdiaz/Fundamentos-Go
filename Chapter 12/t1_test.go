package main

import "testing"

//- struct test, fields: data []int, answer int
//- tests := []test{[]int{}, int}
//- range tests


type test struct {
	data   []int
	answer int
}

//teste por tabela

func TestSomaTabela (t *testing.T) {
	tests := []test {
		test{data : []int {1,2,3,4}, answer: 10},
		test{[]int {10,11,12}, 33},
		test{[]int {-5,0,5,10}, 10},
	}

	for _, v := range tests {
		z :=(Soma(v.data ...))
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}


//(Soma(i ... int) int


func TestSoma(t *testing.T) {
	teste :=(Soma(3,2,1))
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}


func TestMultiplicar(t *testing.T) {
	teste := Multiplica(10,10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

//BenchMark

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(2,2)
	}

	//go test -bench Multiplica
}
 