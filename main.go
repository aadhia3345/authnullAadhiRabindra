package main

type workloadIdentity struct {
	id        int
	tenant_id int
	name      string
	usertype  string
	createdby string
}

type workloadIdentityGroup struct {
	id   int
	name string
}

type serviceAccount struct {
	id        int
	tenant_id int
	username  string
	usertype  string
}

type endpointGroup struct {
	id   int
	name string
}

func addIdentity() {
	/* store the name that the user enters
	store what type if user assigned option is chosen
	based on these two create an object that stores the name and type
	using struct??
	*/
}

func addIdentityGroup() {
	/*
		if they choose add new then we have to create a new identity group
		with a name and type as an object
	*/
}

func getIdentity() {
	/* return the identity corresponding with the tenant id and identity id */
}

func getIdentityGroup() {
	/* return the identity group based on group id and tenant id */
}

func getAllIdentityGroups() {
	/* prints out all identity groups that are currently stored */
}

func getAllIdentities() {
	/* prints out all identities that are stored */
}

func assignToWallet() {
	/*store wallet-id, assignment duration, rotation of credentials
	once they click save you will add it to the wallet */
}

func associateWorkloadRules() {
	/* 3 if statements that must match based on exclusion and inclusion, basically
	check the conditions and then based on that chosoe */
}

func addServiceAccount() {
	/* get id, tenant_id, username, type
	then add the object with these attributes */
}

func addEnpointGroup() {
	/*get id and name of and then add it */
}

func getServiceAccount() {
	/*from tenantid and identityid get the service account
	corresponding to it */
}

func getEndpointGroup() {
	/*from groupid and identityid get the endpoint group
	corresponding to it */
}

func getServiceAccounts() {
	//return all the service accounts stored
}

func getEndpointGroups() {
	//return all endpoint groups that are stored
}

func assignSAtoWallet() {
	//send the service account to the wallet
}

func associateServiceAccountRules() {
	//check the if statements and then decide what permissions should be granted

}
func main() {

}
