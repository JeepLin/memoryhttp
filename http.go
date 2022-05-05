package memoryhttp

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/memory/var", VarSummary)
	http.HandleFunc("/memory/var/{var_name}", varMemory)
	http.HandleFunc("/memory/var/{var_name}/slice_index/{index}", varSliceIndexMemory)
}

type varInfo struct {
	VarName     string `json:"var_name"`
	VarType     string `json:"var_type"`
	VarTypeKind string `json:"var_type_kind"`
	VarSize     string `json:"var_size"`
}

func VarSummary(w http.ResponseWriter, r *http.Request) {
	ret := make([]*varInfo, 0)
	for _, v := range varMap {
		ret = append(ret, &varInfo{
			VarName:     v.VarName,
			VarType:     v.VarType,
			VarTypeKind: v.VarTypeKind,
			VarSize:     v.VarSize,
		})
	}

	b, _ := json.Marshal(ret)
	w.Write(b)
}

func varMemory(w http.ResponseWriter, r *http.Request) {

}

func varSliceIndexMemory(w http.ResponseWriter, r *http.Request) {

}
