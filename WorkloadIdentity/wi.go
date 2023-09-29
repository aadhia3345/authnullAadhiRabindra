package wi

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

func AddIdentity() {
	/* store the name that the user enters
	store what type if user assigned option is chosen
	based on these two create an object that stores the name and type
	using struct??
	*/
}

func AddIdentityGroup() {
	/*
		if they choose add new then we have to create a new identity group
		with a name and type as an object
	*/
}

func GetIdentity() {
	/* return the identity corresponding with the tenant id and identity id */
}

func GetIdentityGroup() {
	/* return the identity group based on group id and tenant id */
}

func GetAllIdentityGroups() {
	/* prints out all identity groups that are currently stored */
}

func GetAllIdentities() {
	/* prints out all identities that are stored */
}

func AssignToWallet() {
	/*store wallet-id, assignment duration, rotation of credentials
	once they click save you will add it to the wallet */
}

func AssociateWorkloadRules() {
	/* 3 if statements that must match based on exclusion and inclusion, basically
	check the conditions and then based on that chosoe */
}
