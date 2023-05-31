package keeper

type Keeper struct {
	authority string
}

func NewKeeper(authority string) *Keeper {
	return &Keeper{
		authority: authority,
	}
}
