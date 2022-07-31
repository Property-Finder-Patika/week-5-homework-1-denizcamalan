package main

type SoftwarePackeges struct{
}

// Request function. These are allowed pakages ,Each has 1 usage limit.
func (a *SoftwarePackeges) clientRequest(packageName string) string {
	if packageName == "package1" {
        return "Operation successful!"
    }
	if packageName == "package2" {
        return "Operation successful!"
    }
	if packageName == "package3" {
        return "Operation successful!"
    }
    return "There is any package with this name."
}