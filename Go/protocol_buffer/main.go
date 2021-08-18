package main

import (
"fmt"
"log"

"github.com/golang/protobuf/proto"
)

func main(){
	elliot := Person{
		Name: "Elliot",
		Age: 24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}
	data, err := proto.Marshal(&elliot)
	
	if err != nil{
		log.Fatal("Marshaling error: ",err)
	}

	newElliot := &Person{}
	err = proto.Unmarshal(data,newElliot)

	if err != nil {
		log.Fatal("unmarshaling error: ",err)
	}
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())
}
