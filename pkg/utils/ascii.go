package utils

import "fmt"

// PrintASCII выводит ASCII‑арт с подписью и ссылкой на GitHub
func PrintASCII() {
	asciiArt := `
    ____   U  ___ u                ____           _      U  ___ __     __U _____ u 
U /"___|u  \/"_ \/         ___   / __"| u       |"|      \/"_ \\ \   /"/\| ___"|/ 
\| |  _ /  | | | |        |_"_| <\___ \/      U | | u    | | | |\ \ / // |  _|"   
 | |_| .-,_| |_| |         | |   u___) |       \| |/_.-,_| |_| |/\ V /_,-| |___   
  \____|\_)-\___/        U/| |\u |____/>>       |_____\_)-\___/U  \_/-(_/|_____|  
  _)(|_      \\       .-,_|___|_,-)(  (__)      //  \\     \\    //      <<   >>  
 (__)__)    (__)       \_)-' '-(_(__)          (_")("_)   (__)  (__)    (__) (__)
 created by Kudzeri
 -> GitHub: https://github.com/Kudzeri
 -> Linkedln: https://www.linkedin.com/in/kudzeri/
`
	fmt.Println(asciiArt)
}
