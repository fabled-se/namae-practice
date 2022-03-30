package namer

import (
	"fmt"
	"testing"
)

func TestFantasyName(t *testing.T){
	fmt.Println(ThemedName("FANTASY"))
}


func TestFantasyPeopleName(t *testing.T){
	fmt.Println(ThemedName("FANTASY_PEOPLE"))
}


func TestScifiName(t *testing.T){
	fmt.Println(ThemedName("SCIFI"))
}


func TestAnimalName(t *testing.T){
	fmt.Println(ThemedName("ANIMAL"))
}

//TODO: Add a test for PEOPLE