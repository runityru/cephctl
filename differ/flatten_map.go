package differ

func flattenMap(in map[string]map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range in {
		for j, m := range v {
			out[k+flattenMapSeparator+j] = m
		}
	}
	return out
}
