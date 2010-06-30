// test program for the readline package
package readline

import "testing"

func TestReadline(t *testing.T) {
	prompt := "by your command> ";

	//loop until ReadLine returns nil (signalling EOF)
	L: for {
		switch result := ReadLine(&prompt); true {
		case result == nil: break L //exit loop

		case *result != "": //ignore blank lines
			println(*result);
			AddHistory(*result); //allow user to recall this line
		}
	}
}
