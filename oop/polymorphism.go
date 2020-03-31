package oop

import "fmt"

func RunPolymorphism() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	Polymorphism in Go is achieved with the help of interfaces. As we have already discussed, interfaces can be implicitly
	implemented in Go. A type implements an interface if it provides definitions for all the methods declared in the interface.
	 */

	fmt.Printf("\nBeginning of Polymorphism using interfaces...\n")
	/*
	Any type which defines all the methods of an interface is said to implicitly implement that interface.
	A variable of type interface can hold any value which implements the interface. This property of interfaces is used
	to achieve polymorphism in Go.
	 */
	project1 := fixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := fixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := timeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	/*
	Let's say the organisation has found a new income stream through advertisements. Let's see how simple it is to add
	this new income stream and calculate the total income without making any changes to the calculateNetIncome function.
	This becomes possible because of polymorphism.
	 */
	bannerAd := advertisement{adName: "Banner Ad", cpc: 2, noOfClicks: 500}
	popupAd := advertisement{adName: "Popup Ad", cpc: 5, noOfClicks: 750}
	incomeStreams := []income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
	/*
	You would have noticed that we did not make any changes to the calculateNetIncome function though we added a new
	income stream. It just worked because of polymorphism. Since the new Advertisement type also implemented the Income
	interface, we were able to add it to the incomeStreams slice. The calculateNetIncome function also worked without
	any changes as it was able to call the calculate() and source() methods of the Advertisement type.
	 */
}