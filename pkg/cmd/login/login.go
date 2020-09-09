package login

func NewLogin(env string, app string, name string, region string, user string, silent bool, ssh bool) {
	rawInstanceList := filterInstances(region, env, app, name, silent, ssh)
	instances := getInstancesInfo(rawInstanceList)

	if silent {
		showIPsList(instances)
	} else {
		showInstanceList(instances, user)
	}

	selectedInstance := selectInstanceIndex(instances)

	if ssh {
		launchSSH(instances[selectedInstance].IP, user)
	} else {
		launchSSM(instances[selectedInstance].ID)
	}
}