package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/namsral/flag"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	BotToken = flag.String("bot-token", "", "Discord bot token")
)

func init() { flag.Parse() }

func main() {
	// creates k8s in-cluster config
	kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates k8s clientset
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates discord client
	discord, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("hello, world")
	_ = clientset
	_ = discord
}
