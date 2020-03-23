package oop

import "fmt"

func RunPolymorphism() {
	fmt.Printf("\nBeginning of introduction...\n")
	fmt.Println("Polymorphism in Go is achieved with the help of interfaces. A type implements an interface if " +
		"it provides definitions for all the methods declared in the interface")

	fmt.Printf("\nBeginning of Polymorphism using interface...\n")
	fmt.Println("Any type which defines all the methods of an interface is said to implicitly implement " +
		"that interface. ")
	project1 := fixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := fixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := timeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := advertisement{adName: "Banner Ad", cpc: 2, noOfClicks: 500}
	popupAd := advertisement{adName: "Popup Ad", cpc: 5, noOfClicks: 750}
	incomeStreams := []income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}