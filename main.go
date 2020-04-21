package main

import (
	"fmt"
	"math"
)

var board= [9][9] int {
	{0,1,0,0,0,0,0,0,0},
	{0,0,2,0,0,0,0,0,0},
	{0,0,0,3,0,0,0,0,0},
	{1,0,0,0,4,0,0,0,0},
	{0,7,0,0,0,5,0,0,0},
	{0,0,8,0,0,0,6,0,0},
	{0,0,0,4,0,0,0,1,0},
	{0,0,0,0,0,0,0,0,0},
	{9,0,7,0,0,0,0,0,0},

}

func PrintGrid(board [9][9] int){
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func IsPossible(x int, y int , n int, board [9][9] int) bool{

	for i:=0;i<9;i++{
		if board[x][i]==n{
			return false
		}
	}
	for i:=0;i<9;i++{
		if board[i][y]==n{
			return false
		}
	}

	x0:= int(math.Floor(float64(x/3)))*3
	y0:= int(math.Floor(float64(y/3)))*3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
		if board[x0+i][y0+j]==n{return false}
		}
	}

	return true
}

func IsPossibleKnight(x int, y int , n int, board [9][9] int) bool{

	for i:=0;i<9;i++{
		if board[x][i]==n{
			return false
		}
	}
	for i:=0;i<9;i++{
		if board[i][y]==n{
			return false
		}
	}

	x0:= int(math.Floor(float64(x/3)))*3
	y0:= int(math.Floor(float64(y/3)))*3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if board[x0+i][y0+j]==n{return false}
		}
	}


	if x>1&&y>0{
	if board[x-2][y-1]==n{
		return false}
	}
	if x>1&&y<8{
	if board[x-2][y+1]==n{
		return false
	}
	}
	if x>0&&y<7 {
	if board[x-1][y+2]==n{
		return false
	}}
	if x<8 && y<7 {
	if board[x+1][y+2]==n{
		return false
	}}
	if x<7&&y<8 {
	if board[x+2][y+1]==n{
		return false
	}}
	if x<7&&y>0 {
	if board[x+2][y-1]==n{
		return false
	}}
	if x<8&&y>1 {
	if board[x+1][y-2]==n{
		return false
	}}
	if x>0&&y>1{
	if board[x-1][y-2]==n{
		return false
	}}

	return true
}

func Solver(board *[9][9] int,variant int) bool{

var	k bool =true

	for i:=0;i<9;i++{
		for j:=0;j<9;j++ {
			if board[i][j]==0 {
				k=false
			}
		}}
	if k==true {
		return true
	}
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]==0{
				for n:=1;n<10;n++{
					if variant==1{
					if IsPossibleKnight(i,j,n,*board)==true{
						board[i][j]=n
						if Solver(board,variant)==true{
							return true
						}
						board[i][j]=0
					} else {
						board[i][j]=0
					}}else if variant==2{
						if IsPossible(i,j,n,*board)==true{
							board[i][j]=n
							if Solver(board,variant)==true{
								return true
							}
							board[i][j]=0
						} else {
							board[i][j]=0
						}
					}

				}
				return false
			}
		}
	}

return false
}

func main(){

	success:= Solver(&board,2)
	if success==true{
		PrintGrid(board)
	} else {
		fmt.Printf("failure")

	}

	return
}