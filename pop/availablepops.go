package pop

var availablePops = []POP{}

func init() {
	availablePops = append(availablePops, NewPOP("US West", "https://api.system.aws-usw02-pr.ice.predix.io"))
	availablePops = append(availablePops, NewPOP("US East", "https://api.system.asv-pr.ice.predix.io"))
	availablePops = append(availablePops, NewPOP("Japan", "https://api.system.aws-jp01-pr.ice.predix.io"))
	availablePops = append(availablePops, NewPOP("UK", "https://api.system.dc-uk01-pr.ice.predix.io"))
}
