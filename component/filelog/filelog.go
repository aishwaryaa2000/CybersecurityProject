package filelog
import(
	"os"
)
func WriteUserLog(username,action string){
	f, err := os.OpenFile("userlog.txt", os.O_APPEND|os.O_WRONLY, 0644) 
	if err != nil {
		panic(err)
	}
	
	defer f.Close()
	text := username + action + "\n"
	_, errr := f.WriteString(text) 
	if errr != nil {
		panic(err)
	}
	
	f.Close()
}
