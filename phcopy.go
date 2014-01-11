// phcopy
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
)

func main() {
	cp("Q:\\common\\Glass Halogen\\Photometry\\Sphere Data.xlsm", "C:\\Sphere.net\\Data\\Sphere Data.xlsm")
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		fmt.Println("Problem with Source File")
		cpFailAlert()
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	//fmt.Printf("destination is %s:\n",d)
	if err != nil {
		fmt.Println("Problem with Destination")
		cpFailAlert()
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	fmt.Println("Your File has been succesfully copied to:", dst)
	cpSuccessAlert()
	return d.Close()
}

func cpFailAlert() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mailhost.consind.ge.com:25")
	if err != nil {
		log.Fatal(err)
	}
	// Set the sender and recipient.
	c.Mail("chris.trail@ge.com")
	c.Rcpt("chris.trail@ge.com")
	c.Rcpt("cstrail@gmail.com")
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("Sphere Data Excel file has failed to copy. Please check both source and destination paths")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}

func cpSuccessAlert() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mailhost.consind.ge.com:25")
	if err != nil {
		log.Fatal(err)
	}
	// Set the sender and recipient.
	c.Mail("chris.trail@ge.com")
	c.Rcpt("matspheredata@ge.com")
	c.Rcpt("cstrail@gmail.com")

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("Sphere Data File successfully copied to \"\\\\matilfs01cige\\geapps\\common\\Glass Halogen\\Photometry\\Sphere Data.xlsm\"")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
