package login

import (
	"fmt"
	"os"
	"strconv"
)

// NewLogin login to the selected instance.
func NewLogin(name string, region string, user string, silent bool, ssh bool, pushKey bool) {
	rawInstanceList := filterInstances(region, name, silent, ssh)
	instances := getInstancesInfo(rawInstanceList)

	if silent {
		showIPsList(instances)
	} else {
		showInstanceList(instances, user)
	}

	selectedInstance := selectInstanceIndex(instances)

	if ssh {
		launchSSH(instances[selectedInstance].IP, user)
	} else if pushKey {
		pushTempKeyPair(instances[selectedInstance].ID, instances[selectedInstance].AZ, instances[selectedInstance].IP, user)
	} else {
		launchSSM(instances[selectedInstance].ID)
	}
}

func showIPsList(instanceList []Instance) {
	if len(instanceList) == 0 {
		os.Exit(0)
	}
	for _, inst := range instanceList {
		fmt.Printf("%s\n", inst.IP)
	}
	os.Exit(0)
}

func showInstanceList(instanceList []Instance, user string) {
	if len(instanceList) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instanceList {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
}

func selectInstanceIndex(instanceList []Instance) int {
	var input string
	var index int
	var err error
	if len(instanceList) == 1 {
		fmt.Printf("\n\n")
		index = 0
	} else {
		fmt.Println("\nWhich one do you want to ssh in?")
		fmt.Scanln(&input)
		index, err = strconv.Atoi(input)
		if err != nil || index > len(instanceList)-1 || index < 0 {
			fmt.Println("Please enter a valid number.")
			index = selectInstanceIndex(instanceList)
		}
	}
	return index
}
