package shared

message StudioProfile {
	int64 id  = 1;
	string name = 2;
	int64 userId = 3;
	[]string listingIds = 4;
}
