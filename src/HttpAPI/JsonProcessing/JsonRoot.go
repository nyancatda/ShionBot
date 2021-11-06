package JsonProcessing

func JsonRoot(status int,message string) (map[string]interface{}) {
	data := map[string]interface{}{
		"status":  status,
		"message": message,
	}
	return data
}