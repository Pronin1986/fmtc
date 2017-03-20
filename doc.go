// Copyright (c) 2017 Pronin S.V.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	The fmtc overrides the print functions of the fmt package, but with the ability to color output in bash using HTML tags

	Available text decoration tags:
			<b></b>			 		- make font Bold
			<strong></strong> 		- make font Bold
			<u></u> 				- make text Underlined
			<dim></dim> 			- make font Dim
			<reverse></reverse> 	- Reverse background anf font color
			<blink></blink> 		- make text Blink (not working on Mint/Ubuntu)
			<black></black> 		- set font color Black (Dark)
			<red></red> 			- set font color Red
			<green></green> 		- set font color Green
			<yellow></yellow> 		- set font color Yellow
			<blue></blue> 			- set font color Blue
			<magenta></magenta> 	- set font color Magenta
			<cyan></cyan> 			- set font color Cyan
			<grey></grey> 			- set font color Grey (Smokey)
			<white></white> 		- set font color White

	Available background colors tags:
			<bg color="black"></bg> 	- black background color
			<bg color="red"></bg> 		- red background color
			<bg color="green"></bg> 	- green background color
			<bg color="yellow"></bg> 	- yellow background color
			<bg color="blue"></bg> 		- blue background color
			<bg color="magenta"></bg> 	- magenta background color
			<bg color="cyan"></bg> 		- cyan background color
			<bg color="grey"></bg> 		- grey background color
			<bg color="white"></bg> 	- white background color

	Examples:

		fmtc.Print("<b>HELLO <blue>WORLD</blue></b>")
		fmtc.Println("<bg color=\"yellow\"><b>HELLO <blue>WORLD</blue></b></bg>")
		fmtc.Printf("<b>%v:</b> <green>%v</green>", name, result)
*/

package fmtc
