package manifest

func (env *Environment) IsWhitelistedAccount(acctID string) bool {
	if len(env.AccountIDs) == 0 {
		// If no account IDs are specified, assume allow all
		return true
	}

	for _, id := range env.AccountIDs {
		if acctID == id {
			return true
		}
	}
	return false
}
