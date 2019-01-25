package generator

// GeneratePropertiesFile =  touch application.properties
func GeneratePropertiesFile(appName string) {
	writer := generateFile(appName + "/src/main/resources/application.properties")
	defer writer.Flush()
}
