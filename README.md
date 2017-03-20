# fmtc
The fmtc overrides the print functions of the fmt package, but with the ability to color output in bash using HTML tags

### Available text decoration tags

```
<b></b>               - make font Bold
<strong></strong>     - make font Bold
<u></u>               - make text Underlined
<dim></dim>           - make font Dim
<reverse></reverse>   - Reverse background anf font color
<blink></blink>       - make text Blink (not working on Mint/Ubuntu)
<black></black>       - set font color Black (Dark)
<red></red>           - set font color Red
<green></green>       - set font color Green
<yellow></yellow>     - set font color Yellow
<blue></blue>         - set font color Blue
<magenta></magenta>   - set font color Magenta
<cyan></cyan>         - set font color Cyan
<grey></grey>         - set font color Grey (Smokey)
<white></white>       - set font color White
```
### Available background colors tags:
```
<bg color="black"></bg> 	- black background color
<bg color="red"></bg> 		- red background color
<bg color="green"></bg> 	- green background color
<bg color="yellow"></bg> 	- yellow background color
<bg color="blue"></bg> 		- blue background color
<bg color="magenta"></bg> 	- magenta background color
<bg color="cyan"></bg> 		- cyan background color
<bg color="grey"></bg> 		- grey background color
<bg color="white"></bg> 	- white background color
```

### Install:
```bash
go get golang.org/x/net/html && go get github.com/Pronin1986/fmtc
```

### Usage:
```golang
fmtc.Print("<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
fmtc.Println("<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
fmtc.Printf("<b>%v <blue>%v</blue> <bg color=\"green\">%v</bg></b>", "HELLO", "BLUE", "TEXT")
```

```golang
got: = fmtc.Sprint("<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
got = fmtc.Sprintln("<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
got = fmtc.Sprintf("<b>%v <blue>%v</blue> <bg color=\"green\">%v</bg></b>", "HELLO", "BLUE", "TEXT")
```

```golang
buf := new(bytes.Buffer)
fmtc.Fprint(buf, "<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
fmtc.Fprintln(buf, "<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
fmtc.Fprintf(buf, "<b>%v <blue>%v</blue> <bg color=\"green\">%v</bg></b>", "HELLO", "BLUE", "TEXT")
```

You have already decorated text for console output like a picture above: 

![UsageExampleResult](http://www.pronin86.ru/git/fmtc/example.png)

### Example:

```golang
package main

import (
	"fmt"

	"github.com/Pronin1986/fmtc"
)

func main() {
	fmt.Println("HELLO WORLD")
	fmtc.Println("<b>HELLO <blue>WORLD</blue></b>")
}
```

![ExampleResult](http://www.pronin86.ru/git/fmtc/example2.png)

