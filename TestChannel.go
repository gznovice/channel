package main

/*
a mini race; 5 people attend; 8 rounds race; random speed  but keep same each round; show sequence for every round.
*/

import(
	"fmt"
	"math/rand"
	"time"
	"sort"	
)

const RoundLen = 20
const num = 5
const roundNum = 8

//every attendant has a channel.
var checkTime [num]chan int

type scoreObj struct{
	No int	
	RoundScore [roundNum]int
}

//key is no, value is scoreObj
var scoreMap map[int]scoreObj

func makeChannels(){	
	for i := range checkTime{
		checkTime[i] = make(chan int, roundNum)
	}
}



func createScoreMap(){
	scoreMap = make(map[int]scoreObj)
	for i := 0; i < num; i++{
		scoreMap[i] = scoreObj{
						No: i}
	}
}

func PrintRoundScore(round int){

	sequnce := make([]int, num)
	
	keys := make([]int, num)
	
	indexmaps := make(map[int]int)
	
	for i :=0; i < num; i++{
		keys[i] = scoreMap[i].RoundScore[round] + i
		indexmaps[keys[i]] = i
	}
	
	sort.Ints(keys)
	
	//fmt.Println("keys", keys)
	
	for i :=0; i < num; i++{
		sequnce[i] = indexmaps[keys[i]]
	}
	
	
	
	/*
	fmt.Println("sequence:", sequnce)
	
	for i :=0; i < num; i++{
		fmt.Printf("The Player %v, his score is %v.\n", sequnce[i], scoreMap[sequnce[i]].RoundScore[round])		
	}
	*/
	/*
	sort.Slice(sequnce, func(i, j int) bool {
       return scoreMap[i].RoundScore[round] < scoreMap[j].RoundScore[round]	//return is a must
    })
	
	fmt.Println("sequence:", sequnce)
	*///won't work , I have to implement my own...
	
	
	
	//Print the Player No. and score, sort by score...	
	fmt.Printf("Round %v\n", round + 1)
	for i :=0; i < num; i++{
		fmt.Printf("The %v is Player %v, his score is %v.\n", i + 1, sequnce[i] + 1, scoreMap[sequnce[i]].RoundScore[round])		
	}
}


func getRandSpeed() int{
	var speeds = []int{1, 2, 4 }
	return speeds[rand.Intn(len(speeds))]
}

func raceRun(myChan chan int){	
	for i := 1; i <= roundNum; i++ {
		timeUsed := RoundLen/getRandSpeed() 
		//fmt.Println("timeUsed:", timeUsed)
		time.Sleep(time.Duration(timeUsed/10) * time.Second)
		myChan <- timeUsed
	}	
}

func main(){
	rand.Seed(time.Now().Unix())
	
	makeChannels()
	createScoreMap()
	
	println("Race Starts");
	
	for i := 0; i < num; i++ {
		go raceRun(checkTime[i])
	}
	
	for i := 0; i < roundNum; i++ {
		for j := 0; j < num; j++ {
			  thisTime := <-checkTime[j]		//var thisTime = <-checkTime[i] is not corrent
			  //fmt.Printf("see index %v round %v time is %v\n", j, i, thisTime)
			 //scoreMap[j].RoundScore[i] = thisTime	//cannot assign map member directly
			 myStruct := scoreMap[j]
			 myStruct.RoundScore[i] = thisTime
			 scoreMap[j] = myStruct
				if( i != 0){
					myStruct := scoreMap[j]
					myStruct.RoundScore[i] += scoreMap[j].RoundScore[i - 1] 
					scoreMap[j] = myStruct
					//scoreMap[j].RoundScore[i] += scoreMap[j].RoundScore[i - 1]	//cannot assign map member directly
				}
		}
		
		PrintRoundScore(i)
	}	
	
	println("Race Ends");
}